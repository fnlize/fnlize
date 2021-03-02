/*
Copyright The Fission Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	corev1 "github.com/fnlize/fnlize/pkg/apis/core/v1"
	versioned "github.com/fnlize/fnlize/pkg/apis/genclient/clientset/versioned"
	internalinterfaces "github.com/fnlize/fnlize/pkg/apis/genclient/informers/externalversions/internalinterfaces"
	v1 "github.com/fnlize/fnlize/pkg/apis/genclient/listers/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// HTTPTriggerInformer provides access to a shared informer and lister for
// HTTPTriggers.
type HTTPTriggerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.HTTPTriggerLister
}

type _hTTPTriggerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewHTTPTriggerInformer constructs a new informer for HTTPTrigger type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewHTTPTriggerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredHTTPTriggerInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredHTTPTriggerInformer constructs a new informer for HTTPTrigger type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredHTTPTriggerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().HTTPTriggers(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().HTTPTriggers(namespace).Watch(options)
			},
		},
		&corev1.HTTPTrigger{},
		resyncPeriod,
		indexers,
	)
}

func (f *_hTTPTriggerInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredHTTPTriggerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *_hTTPTriggerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&corev1.HTTPTrigger{}, f.defaultInformer)
}

func (f *_hTTPTriggerInformer) Lister() v1.HTTPTriggerLister {
	return v1.NewHTTPTriggerLister(f.Informer().GetIndexer())
}
