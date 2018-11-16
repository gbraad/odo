// Code generated by informer-gen. DO NOT EDIT.

package internalversion

import (
	time "time"

	template "github.com/openshift/origin/pkg/template/apis/template"
	internalinterfaces "github.com/openshift/origin/pkg/template/generated/informers/internalversion/internalinterfaces"
	internalclientset "github.com/openshift/origin/pkg/template/generated/internalclientset"
	internalversion "github.com/openshift/origin/pkg/template/generated/listers/template/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// TemplateInstanceInformer provides access to a shared informer and lister for
// TemplateInstances.
type TemplateInstanceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.TemplateInstanceLister
}

type templateInstanceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewTemplateInstanceInformer constructs a new informer for TemplateInstance type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTemplateInstanceInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTemplateInstanceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredTemplateInstanceInformer constructs a new informer for TemplateInstance type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTemplateInstanceInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Template().TemplateInstances(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Template().TemplateInstances(namespace).Watch(options)
			},
		},
		&template.TemplateInstance{},
		resyncPeriod,
		indexers,
	)
}

func (f *templateInstanceInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTemplateInstanceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *templateInstanceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&template.TemplateInstance{}, f.defaultInformer)
}

func (f *templateInstanceInformer) Lister() internalversion.TemplateInstanceLister {
	return internalversion.NewTemplateInstanceLister(f.Informer().GetIndexer())
}