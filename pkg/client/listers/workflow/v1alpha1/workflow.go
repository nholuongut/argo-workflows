// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/nholuongut/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// WorkflowLister helps list Workflows.
// All objects returned here must be treated as read-only.
type WorkflowLister interface {
	// List lists all Workflows in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Workflow, err error)
	// Workflows returns an object that can list and get Workflows.
	Workflows(namespace string) WorkflowNamespaceLister
	WorkflowListerExpansion
}

// workflowLister implements the WorkflowLister interface.
type workflowLister struct {
	indexer cache.Indexer
}

// NewWorkflowLister returns a new WorkflowLister.
func NewWorkflowLister(indexer cache.Indexer) WorkflowLister {
	return &workflowLister{indexer: indexer}
}

// List lists all Workflows in the indexer.
func (s *workflowLister) List(selector labels.Selector) (ret []*v1alpha1.Workflow, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Workflow))
	})
	return ret, err
}

// Workflows returns an object that can list and get Workflows.
func (s *workflowLister) Workflows(namespace string) WorkflowNamespaceLister {
	return workflowNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// WorkflowNamespaceLister helps list and get Workflows.
// All objects returned here must be treated as read-only.
type WorkflowNamespaceLister interface {
	// List lists all Workflows in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Workflow, err error)
	// Get retrieves the Workflow from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Workflow, error)
	WorkflowNamespaceListerExpansion
}

// workflowNamespaceLister implements the WorkflowNamespaceLister
// interface.
type workflowNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Workflows in the indexer for a given namespace.
func (s workflowNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Workflow, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Workflow))
	})
	return ret, err
}

// Get retrieves the Workflow from the indexer for a given namespace and name.
func (s workflowNamespaceLister) Get(name string) (*v1alpha1.Workflow, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("workflow"), name)
	}
	return obj.(*v1alpha1.Workflow), nil
}
