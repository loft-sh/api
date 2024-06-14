// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	managementv1 "github.com/loft-sh/api/v4/pkg/apis/management/v1"
	versioned "github.com/loft-sh/api/v4/pkg/clientset/versioned"
	internalinterfaces "github.com/loft-sh/api/v4/pkg/informers/externalversions/internalinterfaces"
	v1 "github.com/loft-sh/api/v4/pkg/listers/management/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// AnnouncementInformer provides access to a shared informer and lister for
// Announcements.
type AnnouncementInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.AnnouncementLister
}

type announcementInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewAnnouncementInformer constructs a new informer for Announcement type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAnnouncementInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAnnouncementInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredAnnouncementInformer constructs a new informer for Announcement type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAnnouncementInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ManagementV1().Announcements().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ManagementV1().Announcements().Watch(context.TODO(), options)
			},
		},
		&managementv1.Announcement{},
		resyncPeriod,
		indexers,
	)
}

func (f *announcementInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAnnouncementInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *announcementInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&managementv1.Announcement{}, f.defaultInformer)
}

func (f *announcementInformer) Lister() v1.AnnouncementLister {
	return v1.NewAnnouncementLister(f.Informer().GetIndexer())
}
