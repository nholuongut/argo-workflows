package executor

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"

	wfv1 "github.com/nholuongut/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	"github.com/nholuongut/argo-workflows/v3/workflow/common"
	"github.com/nholuongut/argo-workflows/v3/workflow/executor/mocks"
)

const (
	fakePodName       = "fake-test-pod-1234567890"
	fakeNamespace     = "default"
	fakeContainerName = "main"
)

func TestWorkflowExecutor_LoadArtifacts(t *testing.T) {
	tests := []struct {
		name     string
		artifact wfv1.Artifact
		error    string
	}{
		{"ErrNotSupplied", wfv1.Artifact{Name: "foo"}, "required artifact 'foo' not supplied"},
		{"ErrFailedToLoad", wfv1.Artifact{
			Name: "foo",
			ArtifactLocation: wfv1.ArtifactLocation{
				S3: &wfv1.S3Artifact{
					Key: "my-key",
				},
			},
		}, "failed to load artifact 'foo': template artifact location not set"},
		{"ErrNoPath", wfv1.Artifact{
			Name: "foo",
			ArtifactLocation: wfv1.ArtifactLocation{
				S3: &wfv1.S3Artifact{
					S3Bucket: wfv1.S3Bucket{Endpoint: "my-endpoint", Bucket: "my-bucket"},
					Key:      "my-key",
				},
			},
		}, "Artifact foo did not specify a path"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			we := WorkflowExecutor{
				Template: wfv1.Template{
					Inputs: wfv1.Inputs{
						Artifacts: []wfv1.Artifact{test.artifact},
					},
				},
			}
			err := we.LoadArtifacts(context.Background())
			assert.EqualError(t, err, test.error)
		})
	}
}

func TestSaveParameters(t *testing.T) {
	fakeClientset := fake.NewSimpleClientset()
	mockRuntimeExecutor := mocks.ContainerRuntimeExecutor{}
	templateWithOutParam := wfv1.Template{
		Outputs: wfv1.Outputs{
			Parameters: []wfv1.Parameter{
				{
					Name: "my-out",
					ValueFrom: &wfv1.ValueFrom{
						Path: "/path",
					},
				},
			},
		},
	}
	we := WorkflowExecutor{
		PodName:         fakePodName,
		Template:        templateWithOutParam,
		ClientSet:       fakeClientset,
		Namespace:       fakeNamespace,
		RuntimeExecutor: &mockRuntimeExecutor,
	}
	mockRuntimeExecutor.On("GetFileContents", fakeContainerName, "/path").Return("has a newline\n", nil)

	ctx := context.Background()
	err := we.SaveParameters(ctx)
	assert.NoError(t, err)
	assert.Equal(t, "has a newline", we.Template.Outputs.Parameters[0].Value.String())
}

// TestIsBaseImagePath tests logic of isBaseImagePath which determines if a path is coming from a
// base image layer versus a shared volumeMount.
func TestIsBaseImagePath(t *testing.T) {
	templateWithSameDir := wfv1.Template{
		Inputs: wfv1.Inputs{
			Artifacts: []wfv1.Artifact{
				{
					Name: "samedir",
					Path: "/samedir",
				},
			},
		},
		Container: &corev1.Container{},
		Outputs: wfv1.Outputs{
			Artifacts: []wfv1.Artifact{
				{
					Name: "samedir",
					Path: "/samedir",
				},
			},
		},
	}

	we := WorkflowExecutor{
		Template: templateWithSameDir,
	}
	// 1. unrelated dir/file should be captured from base image layer
	assert.True(t, we.isBaseImagePath("/foo"))

	// 2. when input and output directory is same, it should be captured from shared emptyDir
	assert.False(t, we.isBaseImagePath("/samedir"))

	// 3. when output is a sub path of input dir, it should be captured from shared emptyDir
	we.Template.Outputs.Artifacts[0].Path = "/samedir/inner"
	assert.False(t, we.isBaseImagePath("/samedir/inner"))

	// 4. when output happens to overlap with input (in name only), it should be captured from base image layer
	we.Template.Inputs.Artifacts[0].Path = "/hello.txt"
	we.Template.Outputs.Artifacts[0].Path = "/hello.txt-COINCIDENCE"
	assert.True(t, we.isBaseImagePath("/hello.txt-COINCIDENCE"))

	// 5. when output is under a user specified volumeMount, it should be captured from shared mount
	we.Template.Inputs.Artifacts = nil
	we.Template.Container.VolumeMounts = []corev1.VolumeMount{
		{
			Name:      "workdir",
			MountPath: "/user-mount",
		},
	}
	we.Template.Outputs.Artifacts[0].Path = "/user-mount/some-path"
	assert.False(t, we.isBaseImagePath("/user-mount"))
	assert.False(t, we.isBaseImagePath("/user-mount/some-path"))
	assert.False(t, we.isBaseImagePath("/user-mount/some-path/foo"))
	assert.True(t, we.isBaseImagePath("/user-mount-coincidence"))
}

func TestDefaultParameters(t *testing.T) {
	fakeClientset := fake.NewSimpleClientset()
	mockRuntimeExecutor := mocks.ContainerRuntimeExecutor{}
	templateWithOutParam := wfv1.Template{
		Outputs: wfv1.Outputs{
			Parameters: []wfv1.Parameter{
				{
					Name: "my-out",
					ValueFrom: &wfv1.ValueFrom{
						Default: wfv1.AnyStringPtr("Default Value"),
						Path:    "/path",
					},
				},
			},
		},
	}
	we := WorkflowExecutor{
		PodName:         fakePodName,
		Template:        templateWithOutParam,
		ClientSet:       fakeClientset,
		Namespace:       fakeNamespace,
		RuntimeExecutor: &mockRuntimeExecutor,
	}
	mockRuntimeExecutor.On("GetFileContents", fakeContainerName, "/path").Return("", fmt.Errorf("file not found"))

	ctx := context.Background()
	err := we.SaveParameters(ctx)
	assert.NoError(t, err)
	assert.Equal(t, we.Template.Outputs.Parameters[0].Value.String(), "Default Value")
}

func TestDefaultParametersEmptyString(t *testing.T) {
	fakeClientset := fake.NewSimpleClientset()
	mockRuntimeExecutor := mocks.ContainerRuntimeExecutor{}
	templateWithOutParam := wfv1.Template{
		Outputs: wfv1.Outputs{
			Parameters: []wfv1.Parameter{
				{
					Name: "my-out",
					ValueFrom: &wfv1.ValueFrom{
						Default: wfv1.AnyStringPtr(""),
						Path:    "/path",
					},
				},
			},
		},
	}
	we := WorkflowExecutor{
		PodName:         fakePodName,
		Template:        templateWithOutParam,
		ClientSet:       fakeClientset,
		Namespace:       fakeNamespace,
		RuntimeExecutor: &mockRuntimeExecutor,
	}
	mockRuntimeExecutor.On("GetFileContents", fakeContainerName, "/path").Return("", fmt.Errorf("file not found"))

	ctx := context.Background()
	err := we.SaveParameters(ctx)
	assert.NoError(t, err)
	assert.Equal(t, "", we.Template.Outputs.Parameters[0].Value.String())
}

func TestIsTarball(t *testing.T) {
	tests := []struct {
		path      string
		isTarball bool
		expectErr bool
	}{
		{"testdata/file", false, false},
		{"testdata/file.zip", false, false},
		{"testdata/file.tar", false, false},
		{"testdata/file.gz", false, false},
		{"testdata/file.tar.gz", true, false},
		{"testdata/file.tgz", true, false},
		{"testdata/not-found", false, true},
	}

	for _, test := range tests {
		ok, err := isTarball(test.path)
		if test.expectErr {
			assert.Error(t, err, test.path)
		} else {
			assert.NoError(t, err, test.path)
		}
		assert.Equal(t, test.isTarball, ok, test.path)
	}
}

func TestUnzip(t *testing.T) {
	zipPath := "testdata/file.zip"
	destPath := "testdata/unzippedFile"

	// test
	err := unzip(zipPath, destPath)
	assert.NoError(t, err)

	// check unzipped file
	fileInfo, err := os.Stat(destPath)
	assert.NoError(t, err)
	assert.True(t, fileInfo.Mode().IsRegular())

	// cleanup
	err = os.Remove(destPath)
	assert.NoError(t, err)
}

func TestUntar(t *testing.T) {
	tarPath := "testdata/file.tar.gz"
	destPath := "testdata/untarredFile"

	// test
	err := untar(tarPath, destPath)
	assert.NoError(t, err)

	// check untarred file
	fileInfo, err := os.Stat(destPath)
	assert.NoError(t, err)
	assert.True(t, fileInfo.Mode().IsRegular())

	// cleanup
	err = os.Remove(destPath)
	assert.NoError(t, err)
}

func TestChmod(t *testing.T) {
	type perm struct {
		dir  string
		file string
	}

	tests := []struct {
		mode        int32
		recurse     bool
		permissions perm
	}{
		{
			0o777,
			false,
			perm{
				"drwxrwxrwx",
				"-rw-------",
			},
		},
		{
			0o777,
			true,
			perm{
				"drwxrwxrwx",
				"-rwxrwxrwx",
			},
		},
	}

	for _, test := range tests {
		// Setup directory and file for testing
		tempDir, err := ioutil.TempDir("testdata", "chmod-dir-test")
		assert.NoError(t, err)

		tempFile, err := ioutil.TempFile(tempDir, "chmod-file-test")
		assert.NoError(t, err)

		// TearDown test by removing directory and file
		defer os.RemoveAll(tempDir)

		// Run chmod function
		err = chmod(tempDir, test.mode, test.recurse)
		assert.NoError(t, err)

		// Check directory mode if set
		dirPermission, err := os.Stat(tempDir)
		assert.NoError(t, err)
		assert.Equal(t, dirPermission.Mode().String(), test.permissions.dir)

		// Check file mode mode if set
		filePermission, err := os.Stat(tempFile.Name())
		assert.NoError(t, err)
		assert.Equal(t, filePermission.Mode().String(), test.permissions.file)
	}
}

func TestSaveArtifacts(t *testing.T) {
	fakeClientset := fake.NewSimpleClientset()
	mockRuntimeExecutor := mocks.ContainerRuntimeExecutor{}
	templateWithOutParam := wfv1.Template{
		Inputs: wfv1.Inputs{
			Artifacts: []wfv1.Artifact{
				{
					Name: "samedir",
					Path: "/samedir",
				},
			},
		},
		Outputs: wfv1.Outputs{
			Artifacts: []wfv1.Artifact{
				{
					Name:     "samedir",
					Path:     "/samedir",
					Optional: true,
				},
			},
		},
	}
	we := WorkflowExecutor{
		PodName:         fakePodName,
		Template:        templateWithOutParam,
		ClientSet:       fakeClientset,
		Namespace:       fakeNamespace,
		RuntimeExecutor: &mockRuntimeExecutor,
	}

	ctx := context.Background()
	err := we.SaveArtifacts(ctx)
	assert.NoError(t, err)

	we.Template.Outputs.Artifacts[0].Optional = false
	err = we.SaveArtifacts(ctx)
	assert.Error(t, err)
}

func TestMonitorProgress(t *testing.T) {
	deadline, ok := t.Deadline()
	if !ok {
		deadline = time.Now().Add(time.Second)
	}
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	fakeClientset := fake.NewSimpleClientset(&corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fakePodName,
			Namespace: fakeNamespace,
		},
		Spec: corev1.PodSpec{},
		Status: corev1.PodStatus{
			ContainerStatuses: []corev1.ContainerStatus{
				{
					Name: "main",
					State: corev1.ContainerState{
						Running: &corev1.ContainerStateRunning{
							StartedAt: metav1.Now(),
						},
					},
				},
			},
		},
	})

	f, err := os.CreateTemp("", "")
	require.NoError(t, err)
	defer func() {
		name := f.Name()
		err := f.Close()
		assert.NoError(t, err)
		err = os.Remove(name)
		assert.NoError(t, err)
	}()
	annotationPackTickDuration := 5 * time.Millisecond
	readProgressFileTickDuration := time.Millisecond
	progressFile := f.Name()

	mockRuntimeExecutor := mocks.ContainerRuntimeExecutor{}
	we := NewExecutor(fakeClientset, nil, fakePodName, fakeNamespace, &mockRuntimeExecutor, wfv1.Template{}, false, deadline, annotationPackTickDuration, readProgressFileTickDuration)

	go we.monitorProgress(ctx, progressFile)

	go func(ctx context.Context) {
		progress := 0
		maxProgress := 10
		tickDuration := time.Millisecond
		ticker := time.After(tickDuration)
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker:
				t.Logf("tick progress=%d", progress)
				_, err := fmt.Fprintf(f, "%d/100\n", progress*10)
				assert.NoError(t, err)
				if progress >= maxProgress {
					return
				}
				progress += 1
				ticker = time.After(tickDuration)
			}
		}
	}(ctx)

	ticker := time.After(annotationPackTickDuration)
	for {
		select {
		case <-ctx.Done():
			t.Error(ctx.Err())
			return
		case <-ticker:
			pod, err := we.getPod(ctx)
			assert.NoError(t, err)
			progress, ok := pod.Annotations[common.AnnotationKeyProgress]
			if ok && progress == "100/100" {
				t.Log("success reaching 100/100 progress")
				return
			}
			ticker = time.After(annotationPackTickDuration)
		}
	}
}
