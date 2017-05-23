// This file was automatically generated by informer-gen

package v1

import (
	api_v1 "github.com/openshift/origin/pkg/authorization/api/v1"
	clientset "github.com/openshift/origin/pkg/authorization/generated/clientset"
	internalinterfaces "github.com/openshift/origin/pkg/authorization/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/openshift/origin/pkg/authorization/generated/listers/authorization/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// ClusterPolicyInformer provides access to a shared informer and lister for
// ClusterPolicies.
type ClusterPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ClusterPolicyLister
}

type clusterPolicyInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newClusterPolicyInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				return client.AuthorizationV1().ClusterPolicies().List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				return client.AuthorizationV1().ClusterPolicies().Watch(options)
			},
		},
		&api_v1.ClusterPolicy{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *clusterPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&api_v1.ClusterPolicy{}, newClusterPolicyInformer)
}

func (f *clusterPolicyInformer) Lister() v1.ClusterPolicyLister {
	return v1.NewClusterPolicyLister(f.Informer().GetIndexer())
}
