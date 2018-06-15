// This file was automatically generated by informer-gen

package v1

import (
	project_v1 "github.com/openshift/origin/pkg/project/apis/project/v1"
	clientset "github.com/openshift/origin/pkg/project/generated/clientset"
	internalinterfaces "github.com/openshift/origin/pkg/project/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/openshift/origin/pkg/project/generated/listers/project/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// ProjectInformer provides access to a shared informer and lister for
// Projects.
type ProjectInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ProjectLister
}

type projectInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newProjectInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				return client.ProjectV1().Projects().List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				return client.ProjectV1().Projects().Watch(options)
			},
		},
		&project_v1.Project{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *projectInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&project_v1.Project{}, newProjectInformer)
}

func (f *projectInformer) Lister() v1.ProjectLister {
	return v1.NewProjectLister(f.Informer().GetIndexer())
}
