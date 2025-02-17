package commands

import (
	"bytes"
	"fmt"
	"testing"
	"text/tabwriter"
	"time"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	wfv1 "github.com/nholuongut/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
)

func testPrintNodeImpl(t *testing.T, expected string, node wfv1.NodeStatus, getArgs getFlags) {
	var result bytes.Buffer
	w := tabwriter.NewWriter(&result, 0, 8, 1, '\t', 0)
	filtered, _ := filterNode(node, getArgs)
	if !filtered {
		printNode(w, node, "", getArgs)
	}
	err := w.Flush()
	assert.NoError(t, err)
	assert.Equal(t, expected, result.String())
}

// TestPrintNode
func TestPrintNode(t *testing.T) {
	nodeName := "testNode"
	kubernetesNodeName := "testKnodeName"
	nodeTemplateName := "testTemplate"
	nodeTemplateRefName := "testTemplateRef"
	nodeID := "testID"
	nodeMessage := "test"
	getArgs := getFlags{
		output: "",
	}
	timestamp := metav1.Time{
		Time: time.Now(),
	}
	node := wfv1.NodeStatus{
		Name:        nodeName,
		Phase:       wfv1.NodeRunning,
		DisplayName: nodeName,
		Type:        wfv1.NodeTypePod,
		ID:          nodeID,
		StartedAt:   timestamp,
		FinishedAt:  timestamp,
		Message:     nodeMessage,
	}
	node.HostNodeName = kubernetesNodeName
	testPrintNodeImpl(t, fmt.Sprintf("%s %s\t%s\t%s\t%s\t%s\t%s\n", jobStatusIconMap[wfv1.NodeRunning], nodeName, "", nodeID, "0s", nodeMessage, ""), node, getArgs)

	// Compatibility test
	getArgs.status = "Running"
	testPrintNodeImpl(t, fmt.Sprintf("%s %s\t\t%s\t%s\t%s\t\n", jobStatusIconMap[wfv1.NodeRunning], nodeName, nodeID, "0s", nodeMessage), node, getArgs)

	getArgs.status = ""
	getArgs.nodeFieldSelectorString = "phase=Running"
	testPrintNodeImpl(t, fmt.Sprintf("%s %s\t\t%s\t%s\t%s\t\n", jobStatusIconMap[wfv1.NodeRunning], nodeName, nodeID, "0s", nodeMessage), node, getArgs)

	getArgs.nodeFieldSelectorString = "phase!=foobar"
	testPrintNodeImpl(t, fmt.Sprintf("%s %s\t\t%s\t%s\t%s\t\n", jobStatusIconMap[wfv1.NodeRunning], nodeName, nodeID, "0s", nodeMessage), node, getArgs)

	getArgs.nodeFieldSelectorString = "phase!=Running"
	testPrintNodeImpl(t, "", node, getArgs)

	// Compatibility test
	getArgs.nodeFieldSelectorString = ""
	getArgs.status = "foobar"
	testPrintNodeImpl(t, "", node, getArgs)

	getArgs.status = ""
	getArgs.nodeFieldSelectorString = "phase=foobar"
	testPrintNodeImpl(t, "", node, getArgs)

	getArgs = getFlags{
		output: "",
	}

	node.TemplateName = nodeTemplateName
	testPrintNodeImpl(t, fmt.Sprintf("%s %s\t%s\t%s\t%s\t%s\t%s\n", jobStatusIconMap[wfv1.NodeRunning], nodeName, nodeTemplateName, nodeID, "0s", nodeMessage, ""), node, getArgs)

	node.Type = wfv1.NodeTypeSuspend
	testPrintNodeImpl(t, fmt.Sprintf("%s %s\t%s\t%s\t%s\t%s\t%s\n", nodeTypeIconMap[wfv1.NodeTypeSuspend], nodeName, nodeTemplateName, "", "", nodeMessage, ""), node, getArgs)

	node.TemplateRef = &wfv1.TemplateRef{
		Name:     nodeTemplateRefName,
		Template: nodeTemplateRefName,
	}
	testPrintNodeImpl(t, fmt.Sprintf("%s %s\t%s/%s\t%s\t%s\t%s\t%s\n", nodeTypeIconMap[wfv1.NodeTypeSuspend], nodeName, nodeTemplateRefName, nodeTemplateRefName, "", "", nodeMessage, ""), node, getArgs)

	getArgs.output = "wide"
	testPrintNodeImpl(t, fmt.Sprintf("%s %s\t%s/%s\t%s\t%s\t%s\t%s\t%s\t\n", nodeTypeIconMap[wfv1.NodeTypeSuspend], nodeName, nodeTemplateRefName, nodeTemplateRefName, "", "", getArtifactsString(node), nodeMessage, ""), node, getArgs)

	node.Type = wfv1.NodeTypePod
	testPrintNodeImpl(t, fmt.Sprintf("%s %s\t%s/%s\t%s\t%s\t%s\t%s\t%s\t%s\n", jobStatusIconMap[wfv1.NodeRunning], nodeName, nodeTemplateRefName, nodeTemplateRefName, nodeID, "0s", getArtifactsString(node), nodeMessage, "", kubernetesNodeName), node, getArgs)

	getArgs.output = "short"
	testPrintNodeImpl(t, fmt.Sprintf("%s %s\t%s/%s\t%s\t%s\t%s\t%s\n", nodeTypeIconMap[wfv1.NodeTypeSuspend], nodeName, nodeTemplateRefName, nodeTemplateRefName, nodeID, "0s", nodeMessage, kubernetesNodeName), node, getArgs)

	getArgs.status = "foobar"
	testPrintNodeImpl(t, "", node, getArgs)
}

func TestStatusToNodeFieldSelector(t *testing.T) {
	one := statusToNodeFieldSelector("Running")
	assert.Equal(t, "phase=Running", one)
}

func Test_printWorkflowHelper(t *testing.T) {
	t.Run("Progress", func(t *testing.T) {
		var wf wfv1.Workflow
		wfv1.MustUnmarshal(`
status:
  phase: Running
  progress: 1/2
`, &wf)
		output := printWorkflowHelper(&wf, getFlags{})
		assert.Regexp(t, `Progress: *1/2`, output)
	})
	t.Run("EstimatedDuration", func(t *testing.T) {
		var wf wfv1.Workflow
		wfv1.MustUnmarshal(`
status:
  estimatedDuration: 1
  phase: Running
`, &wf)
		output := printWorkflowHelper(&wf, getFlags{})
		assert.Regexp(t, `EstimatedDuration: *1 second`, output)
	})
	t.Run("IndexOrdering", func(t *testing.T) {
		var wf wfv1.Workflow
		wfv1.MustUnmarshal(`apiVersion: nholuongut.io/v1alpha1
kind: Workflow
metadata:
  creationTimestamp: "2020-06-02T16:04:21Z"
  generateName: many-items-
  generation: 32
  labels:
    workflows.nholuongut.io/completed: "true"
    workflows.nholuongut.io/phase: Succeeded
  name: many-items-z26lj
  namespace: argo
  resourceVersion: "5102"
  selfLink: /apis/nholuongut.io/v1alpha1/namespaces/argo/workflows/many-items-z26lj
  uid: d21f092a-f659-4300-bd69-983a9912a379
spec:
  entrypoint: parallel-sleep
  templates:
  - name: parallel-sleep
    steps:
    - - name: sleep
        template: sleep
        withItems:
        - zero
        - one
        - two
        - three
        - four
        - five
        - six
        - seven
        - eight
        - nine
        - ten
        - eleven
        - twelve
  - container:
      command:
      - sh
      - -c
      - sleep 10
      image: alpine:latest
    name: sleep
status:
  conditions:
  - status: "True"
    type: Completed
  finishedAt: "2020-06-02T16:05:01Z"
  nodes:
    many-items-z26lj:
      children:
      - many-items-z26lj-1414877240
      displayName: many-items-z26lj
      finishedAt: "2020-06-02T16:05:01Z"
      id: many-items-z26lj
      name: many-items-z26lj
      outboundNodes:
      - many-items-z26lj-1939921510
      - many-items-z26lj-2156977535
      - many-items-z26lj-3409403178
      - many-items-z26lj-1774150289
      - many-items-z26lj-3491220632
      - many-items-z26lj-1942531647
      - many-items-z26lj-3178865096
      - many-items-z26lj-3031375822
      - many-items-z26lj-753834747
      - many-items-z26lj-2619926859
      - many-items-z26lj-1052882686
      - many-items-z26lj-3011405271
      - many-items-z26lj-3126938806
      phase: Succeeded
      startedAt: "2020-06-02T16:04:21Z"
      templateName: parallel-sleep
      type: Steps
    many-items-z26lj-753834747:
      boundaryID: many-items-z26lj
      displayName: sleep(8:eight)
      finishedAt: "2020-06-02T16:04:42Z"
      id: many-items-z26lj-753834747
      name: many-items-z26lj[0].sleep(8:eight)
      startedAt: "2020-06-02T16:04:21Z"
      templateName: sleep
      type: Pod
    many-items-z26lj-1052882686:
      boundaryID: many-items-z26lj
      displayName: sleep(10:ten)
      finishedAt: "2020-06-02T16:04:45Z"
      id: many-items-z26lj-1052882686
      name: many-items-z26lj[0].sleep(10:ten)
      startedAt: "2020-06-02T16:04:22Z"
      templateName: sleep
      type: Pod
    many-items-z26lj-1414877240:
      boundaryID: many-items-z26lj
      children:
      - many-items-z26lj-1939921510
      - many-items-z26lj-2156977535
      - many-items-z26lj-3409403178
      - many-items-z26lj-1774150289
      - many-items-z26lj-3491220632
      - many-items-z26lj-1942531647
      - many-items-z26lj-3178865096
      - many-items-z26lj-3031375822
      - many-items-z26lj-753834747
      - many-items-z26lj-2619926859
      - many-items-z26lj-1052882686
      - many-items-z26lj-3011405271
      - many-items-z26lj-3126938806
      displayName: '[0]'
      finishedAt: "2020-06-02T16:05:01Z"
      id: many-items-z26lj-1414877240
      name: many-items-z26lj[0]
      phase: Succeeded
      startedAt: "2020-06-02T16:04:21Z"
      templateName: parallel-sleep
      type: StepGroup
    many-items-z26lj-1774150289:
      boundaryID: many-items-z26lj
      displayName: sleep(3:three)
      finishedAt: "2020-06-02T16:04:54Z"
      id: many-items-z26lj-1774150289
      name: many-items-z26lj[0].sleep(3:three)
      startedAt: "2020-06-02T16:04:21Z"
      templateName: sleep
      type: Pod
    many-items-z26lj-1939921510:
      boundaryID: many-items-z26lj
      displayName: sleep(0:zero)
      finishedAt: "2020-06-02T16:04:48Z"
      id: many-items-z26lj-1939921510
      name: many-items-z26lj[0].sleep(0:zero)
      startedAt: "2020-06-02T16:04:21Z"
      templateName: sleep
      type: Pod
    many-items-z26lj-1942531647:
      boundaryID: many-items-z26lj
      displayName: sleep(5:five)
      finishedAt: "2020-06-02T16:04:47Z"
      id: many-items-z26lj-1942531647
      name: many-items-z26lj[0].sleep(5:five)
      phase: Succeeded
      startedAt: "2020-06-02T16:04:21Z"
      templateName: sleep
      type: Pod
    many-items-z26lj-2156977535:
      boundaryID: many-items-z26lj
      displayName: sleep(1:one)
      finishedAt: "2020-06-02T16:04:53Z"
      id: many-items-z26lj-2156977535
      name: many-items-z26lj[0].sleep(1:one)
      startedAt: "2020-06-02T16:04:21Z"
      templateName: sleep
      type: Pod
    many-items-z26lj-2619926859:
      boundaryID: many-items-z26lj
      displayName: sleep(9:nine)
      finishedAt: "2020-06-02T16:04:40Z"
      id: many-items-z26lj-2619926859
      name: many-items-z26lj[0].sleep(9:nine)
      startedAt: "2020-06-02T16:04:21Z"
      templateName: sleep
      type: Pod
    many-items-z26lj-3011405271:
      boundaryID: many-items-z26lj
      displayName: sleep(11:eleven)
      finishedAt: "2020-06-02T16:04:44Z"
      id: many-items-z26lj-3011405271
      name: many-items-z26lj[0].sleep(11:eleven)
      startedAt: "2020-06-02T16:04:22Z"
      templateName: sleep
      type: Pod
    many-items-z26lj-3031375822:
      boundaryID: many-items-z26lj
      displayName: sleep(7:seven)
      finishedAt: "2020-06-02T16:04:57Z"
      id: many-items-z26lj-3031375822
      name: many-items-z26lj[0].sleep(7:seven)
      startedAt: "2020-06-02T16:04:21Z"
      templateName: sleep
      type: Pod
    many-items-z26lj-3126938806:
      boundaryID: many-items-z26lj
      displayName: sleep(12:twelve)
      finishedAt: "2020-06-02T16:04:59Z"
      id: many-items-z26lj-3126938806
      name: many-items-z26lj[0].sleep(12:twelve)
      startedAt: "2020-06-02T16:04:22Z"
      templateName: sleep
      type: Pod
    many-items-z26lj-3178865096:
      boundaryID: many-items-z26lj
      displayName: sleep(6:six)
      finishedAt: "2020-06-02T16:04:56Z"
      id: many-items-z26lj-3178865096
      name: many-items-z26lj[0].sleep(6:six)
      startedAt: "2020-06-02T16:04:21Z"
      templateName: sleep
      type: Pod
    many-items-z26lj-3409403178:
      boundaryID: many-items-z26lj
      displayName: sleep(2:two)
      finishedAt: "2020-06-02T16:04:51Z"
      id: many-items-z26lj-3409403178
      name: many-items-z26lj[0].sleep(2:two)
      startedAt: "2020-06-02T16:04:21Z"
      templateName: sleep
      type: Pod
    many-items-z26lj-3491220632:
      boundaryID: many-items-z26lj
      displayName: sleep(4:four)
      finishedAt: "2020-06-02T16:04:50Z"
      id: many-items-z26lj-3491220632
      name: many-items-z26lj[0].sleep(4:four)
      phase: Succeeded
      startedAt: "2020-06-02T16:04:21Z"
      templateName: sleep
      type: Pod
  phase: Succeeded
  startedAt: "2020-06-02T16:04:21Z"
`, &wf)
		output := printWorkflowHelper(&wf, getFlags{})
		assert.Contains(t, output, `         
   ├─ sleep(9:nine)     sleep           many-items-z26lj-2619926859  19s         
   ├─ sleep(10:ten)     sleep           many-items-z26lj-1052882686  23s         
   ├─ sleep(11:eleven)  sleep           many-items-z26lj-3011405271  22s`)
		assert.Contains(t, output, "This workflow does not have security context set. "+
			"You can run your workflow pods more securely by setting it.\n"+
			"Learn more at https://nholuongut.github.io/argo-workflows/workflow-pod-security-context/\n")
	})
}

func Test_printWorkflowHelperNudges(t *testing.T) {
	securedWf := wfv1.Workflow{
		ObjectMeta: metav1.ObjectMeta{},
		Spec: wfv1.WorkflowSpec{
			SecurityContext: &corev1.PodSecurityContext{},
		},
	}

	insecureWf := securedWf
	insecureWf.Spec.SecurityContext = nil

	securityNudges := "This workflow does not have security context set. " +
		"You can run your workflow pods more securely by setting it.\n" +
		"Learn more at https://nholuongut.github.io/argo-workflows/workflow-pod-security-context/\n"

	t.Run("SecuredWorkflow", func(t *testing.T) {
		output := printWorkflowHelper(&securedWf, getFlags{})
		assert.NotContains(t, output, securityNudges)
	})
	t.Run("InsecureWorkflow", func(t *testing.T) {
		output := printWorkflowHelper(&insecureWf, getFlags{})
		assert.Contains(t, output, securityNudges)
	})
}
