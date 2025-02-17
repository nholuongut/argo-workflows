// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/nholuongut/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeWorkflowTemplates implements WorkflowTemplateInterface
type FakeWorkflowTemplates struct {
	Fake *FakenholuongutV1alpha1
	ns   string
}

var workflowtemplatesResource = schema.GroupVersionResource{Group: "nholuongut.io", Version: "v1alpha1", Resource: "workflowtemplates"}

var workflowtemplatesKind = schema.GroupVersionKind{Group: "nholuongut.io", Version: "v1alpha1", Kind: "WorkflowTemplate"}

// Get takes name of the workflowTemplate, and returns the corresponding workflowTemplate object, and an error if there is any.
func (c *FakeWorkflowTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.WorkflowTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(workflowtemplatesResource, c.ns, name), &v1alpha1.WorkflowTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WorkflowTemplate), err
}

// List takes label and field selectors, and returns the list of WorkflowTemplates that match those selectors.
func (c *FakeWorkflowTemplates) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.WorkflowTemplateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(workflowtemplatesResource, workflowtemplatesKind, c.ns, opts), &v1alpha1.WorkflowTemplateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.WorkflowTemplateList{ListMeta: obj.(*v1alpha1.WorkflowTemplateList).ListMeta}
	for _, item := range obj.(*v1alpha1.WorkflowTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested workflowTemplates.
func (c *FakeWorkflowTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(workflowtemplatesResource, c.ns, opts))

}

// Create takes the representation of a workflowTemplate and creates it.  Returns the server's representation of the workflowTemplate, and an error, if there is any.
func (c *FakeWorkflowTemplates) Create(ctx context.Context, workflowTemplate *v1alpha1.WorkflowTemplate, opts v1.CreateOptions) (result *v1alpha1.WorkflowTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(workflowtemplatesResource, c.ns, workflowTemplate), &v1alpha1.WorkflowTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WorkflowTemplate), err
}

// Update takes the representation of a workflowTemplate and updates it. Returns the server's representation of the workflowTemplate, and an error, if there is any.
func (c *FakeWorkflowTemplates) Update(ctx context.Context, workflowTemplate *v1alpha1.WorkflowTemplate, opts v1.UpdateOptions) (result *v1alpha1.WorkflowTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(workflowtemplatesResource, c.ns, workflowTemplate), &v1alpha1.WorkflowTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WorkflowTemplate), err
}

// Delete takes name of the workflowTemplate and deletes it. Returns an error if one occurs.
func (c *FakeWorkflowTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(workflowtemplatesResource, c.ns, name), &v1alpha1.WorkflowTemplate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeWorkflowTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(workflowtemplatesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.WorkflowTemplateList{})
	return err
}

// Patch applies the patch and returns the patched workflowTemplate.
func (c *FakeWorkflowTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.WorkflowTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(workflowtemplatesResource, c.ns, name, pt, data, subresources...), &v1alpha1.WorkflowTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WorkflowTemplate), err
}
