// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	workflowv1alpha1 "github.com/nholuongut/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	versioned "github.com/nholuongut/argo-workflows/v3/pkg/client/clientset/versioned"
	internalinterfaces "github.com/nholuongut/argo-workflows/v3/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/nholuongut/argo-workflows/v3/pkg/client/listers/workflow/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterWorkflowTemplateInformer provides access to a shared informer and lister for
// ClusterWorkflowTemplates.
type ClusterWorkflowTemplateInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ClusterWorkflowTemplateLister
}

type clusterWorkflowTemplateInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterWorkflowTemplateInformer constructs a new informer for ClusterWorkflowTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterWorkflowTemplateInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterWorkflowTemplateInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterWorkflowTemplateInformer constructs a new informer for ClusterWorkflowTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterWorkflowTemplateInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.nholuongutV1alpha1().ClusterWorkflowTemplates().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.nholuongutV1alpha1().ClusterWorkflowTemplates().Watch(context.TODO(), options)
			},
		},
		&workflowv1alpha1.ClusterWorkflowTemplate{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterWorkflowTemplateInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterWorkflowTemplateInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterWorkflowTemplateInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&workflowv1alpha1.ClusterWorkflowTemplate{}, f.defaultInformer)
}

func (f *clusterWorkflowTemplateInformer) Lister() v1alpha1.ClusterWorkflowTemplateLister {
	return v1alpha1.NewClusterWorkflowTemplateLister(f.Informer().GetIndexer())
}
