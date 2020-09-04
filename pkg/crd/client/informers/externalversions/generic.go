// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	v1 "k8c.io/kubermatic/v2/pkg/crd/kubermatic/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=kubermatic.k8s.io, Version=v1
	case v1.SchemeGroupVersion.WithResource("addons"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubermatic().V1().Addons().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("addonconfigs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubermatic().V1().AddonConfigs().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("clusters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubermatic().V1().Clusters().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("externalclusters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubermatic().V1().ExternalClusters().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("kubermaticsettings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubermatic().V1().KubermaticSettings().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("projects"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubermatic().V1().Projects().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("users"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubermatic().V1().Users().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("userprojectbindings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubermatic().V1().UserProjectBindings().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("usersshkeies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubermatic().V1().UserSSHKeys().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}