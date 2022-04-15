// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	versioned "github.com/GoogleContainerTools/kpt/porch/api/generated/clientset/versioned"
	internalinterfaces "github.com/GoogleContainerTools/kpt/porch/api/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/GoogleContainerTools/kpt/porch/api/generated/listers/porch/v1alpha1"
	porchv1alpha1 "github.com/GoogleContainerTools/kpt/porch/api/porch/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// FunctionInformer provides access to a shared informer and lister for
// Functions.
type FunctionInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.FunctionLister
}

type functionInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewFunctionInformer constructs a new informer for Function type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFunctionInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFunctionInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredFunctionInformer constructs a new informer for Function type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFunctionInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PorchV1alpha1().Functions(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PorchV1alpha1().Functions(namespace).Watch(context.TODO(), options)
			},
		},
		&porchv1alpha1.Function{},
		resyncPeriod,
		indexers,
	)
}

func (f *functionInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFunctionInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *functionInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&porchv1alpha1.Function{}, f.defaultInformer)
}

func (f *functionInformer) Lister() v1alpha1.FunctionLister {
	return v1alpha1.NewFunctionLister(f.Informer().GetIndexer())
}