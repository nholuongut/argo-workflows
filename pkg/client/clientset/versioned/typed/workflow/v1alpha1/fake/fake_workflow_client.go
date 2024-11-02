// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/nholuongut/argo-workflows/v3/pkg/client/clientset/versioned/typed/workflow/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakenholuongutV1alpha1 struct {
	*testing.Fake
}

func (c *FakenholuongutV1alpha1) ClusterWorkflowTemplates() v1alpha1.ClusterWorkflowTemplateInterface {
	return &FakeClusterWorkflowTemplates{c}
}

func (c *FakenholuongutV1alpha1) CronWorkflows(namespace string) v1alpha1.CronWorkflowInterface {
	return &FakeCronWorkflows{c, namespace}
}

func (c *FakenholuongutV1alpha1) Workflows(namespace string) v1alpha1.WorkflowInterface {
	return &FakeWorkflows{c, namespace}
}

func (c *FakenholuongutV1alpha1) WorkflowEventBindings(namespace string) v1alpha1.WorkflowEventBindingInterface {
	return &FakeWorkflowEventBindings{c, namespace}
}

func (c *FakenholuongutV1alpha1) WorkflowTaskSets(namespace string) v1alpha1.WorkflowTaskSetInterface {
	return &FakeWorkflowTaskSets{c, namespace}
}

func (c *FakenholuongutV1alpha1) WorkflowTemplates(namespace string) v1alpha1.WorkflowTemplateInterface {
	return &FakeWorkflowTemplates{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakenholuongutV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
