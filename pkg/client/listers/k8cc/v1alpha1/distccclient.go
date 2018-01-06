/*
Copyright 2018 The Kubernetes Authors.

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

// This file was automatically generated by lister-gen

package v1alpha1

import (
	v1alpha1 "github.com/mbrt/k8cc/pkg/apis/k8cc.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DistccClientLister helps list DistccClients.
type DistccClientLister interface {
	// List lists all DistccClients in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DistccClient, err error)
	// DistccClients returns an object that can list and get DistccClients.
	DistccClients(namespace string) DistccClientNamespaceLister
	DistccClientListerExpansion
}

// distccClientLister implements the DistccClientLister interface.
type distccClientLister struct {
	indexer cache.Indexer
}

// NewDistccClientLister returns a new DistccClientLister.
func NewDistccClientLister(indexer cache.Indexer) DistccClientLister {
	return &distccClientLister{indexer: indexer}
}

// List lists all DistccClients in the indexer.
func (s *distccClientLister) List(selector labels.Selector) (ret []*v1alpha1.DistccClient, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DistccClient))
	})
	return ret, err
}

// DistccClients returns an object that can list and get DistccClients.
func (s *distccClientLister) DistccClients(namespace string) DistccClientNamespaceLister {
	return distccClientNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DistccClientNamespaceLister helps list and get DistccClients.
type DistccClientNamespaceLister interface {
	// List lists all DistccClients in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DistccClient, err error)
	// Get retrieves the DistccClient from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DistccClient, error)
	DistccClientNamespaceListerExpansion
}

// distccClientNamespaceLister implements the DistccClientNamespaceLister
// interface.
type distccClientNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DistccClients in the indexer for a given namespace.
func (s distccClientNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DistccClient, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DistccClient))
	})
	return ret, err
}

// Get retrieves the DistccClient from the indexer for a given namespace and name.
func (s distccClientNamespaceLister) Get(name string) (*v1alpha1.DistccClient, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("distccclient"), name)
	}
	return obj.(*v1alpha1.DistccClient), nil
}
