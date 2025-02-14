package v1alpha1

import (
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/yaml"
)

func validateContainerSetTemplate(yamlStr string) error {
	var cst ContainerSetTemplate
	err := yaml.Unmarshal([]byte(yamlStr), &cst)
	if err != nil {
		panic(err)
	}
	return cst.Validate()
}

func TestContainerSetTemplate(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		x := &ContainerSetTemplate{}
		assert.Empty(t, x.GetGraph())
		assert.Empty(t, x.GetContainers())
		assert.False(t, x.HasSequencedContainers())
	})
	t.Run("Single", func(t *testing.T) {
		x := &ContainerSetTemplate{Containers: []ContainerNode{{}}}
		assert.Len(t, x.GetGraph(), 1)
		assert.Len(t, x.GetContainers(), 1)
		assert.False(t, x.HasSequencedContainers())
	})
	t.Run("Parallel", func(t *testing.T) {
		x := &ContainerSetTemplate{Containers: []ContainerNode{{}, {}}}
		assert.Len(t, x.GetGraph(), 2)
		assert.Len(t, x.GetContainers(), 2)
		assert.False(t, x.HasSequencedContainers())
	})
	t.Run("Graph", func(t *testing.T) {
		x := &ContainerSetTemplate{Containers: []ContainerNode{{Container: corev1.Container{Name: "a"}}, {Dependencies: []string{"a"}}}}
		assert.Len(t, x.GetGraph(), 2)
		assert.Len(t, x.GetContainers(), 2)
		assert.True(t, x.HasSequencedContainers())
		assert.True(t, x.HasContainerNamed("a"))
	})
}

func TestInvalidContainerSetEmpty(t *testing.T) {
	invalidContainerSetEmpty := `
volumeMounts:
  - name: workspace
    mountPath: /workspace
`
	err := validateContainerSetTemplate(invalidContainerSetEmpty)
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "containers must have at least one container")
	}
}

func TestInvalidContainerSetDuplicateVolumeMounting(t *testing.T) {
	invalidContainerSetDuplicateNames := `
volumeMounts:
  - name: workspace
    mountPath: /workspace
  - name: workspace2
    mountPath: /workspace
containers:
  - name: a
    image: nholuongut/argosay:v2
`
	err := validateContainerSetTemplate(invalidContainerSetDuplicateNames)
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "volumeMounts[1].mountPath '/workspace' already mounted in volumeMounts.workspace")
	}
}

func TestInvalidContainerSetDuplicateNames(t *testing.T) {
	invalidContainerSetDuplicateNames := `
volumeMounts:
  - name: workspace
    mountPath: /workspace
containers:
  - name: a
    image: nholuongut/argosay:v2
  - name: a
    image: nholuongut/argosay:v2
`
	err := validateContainerSetTemplate(invalidContainerSetDuplicateNames)
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "containers[1].name 'a' is not unique")
	}
}

func TestInvalidContainerSetDependencyNotFound(t *testing.T) {
	invalidContainerSetDependencyNotFound := `
volumeMounts:
  - name: workspace
    mountPath: /workspace
containers:
  - name: a
    image: nholuongut/argosay:v2
  - name: b
    image: nholuongut/argosay:v2
    dependencies:
      - c
`
	err := validateContainerSetTemplate(invalidContainerSetDependencyNotFound)
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "containers.b dependency 'c' not defined")
	}
}

func TestInvalidContainerSetDependencyCycle(t *testing.T) {
	invalidContainerSetDependencyCycle := `
volumeMounts:
  - name: workspace
    mountPath: /workspace
containers:
  - name: a
    image: nholuongut/argosay:v2
    dependencies:
      - b
  - name: b
    image: nholuongut/argosay:v2
    dependencies:
      - a
`
	err := validateContainerSetTemplate(invalidContainerSetDependencyCycle)
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "containers dependency cycle detected: b->a->b")
	}
}

func TestValidContainerSet(t *testing.T) {
	validContainerSet := `
volumeMounts:
  - name: workspace
    mountPath: /workspace
containers:
  - name: a
    image: nholuongut/argosay:v2
  - name: b
    image: nholuongut/argosay:v2
    dependencies:
      - a
  - name: c
    image: nholuongut/argosay:v2
    dependencies:
      - a
  - name: d
    image: nholuongut/argosay:v2
    dependencies:
      - b
      - c
`
	err := validateContainerSetTemplate(validContainerSet)
	assert.NoError(t, err)
}
