/*
Copyright 2020 The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/btwiuse/k8cc/pkg/apis/k8cc.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DistccLister helps list Distccs.
type DistccLister interface {
	// List lists all Distccs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Distcc, err error)
	// Distccs returns an object that can list and get Distccs.
	Distccs(namespace string) DistccNamespaceLister
	DistccListerExpansion
}

// distccLister implements the DistccLister interface.
type distccLister struct {
	indexer cache.Indexer
}

// NewDistccLister returns a new DistccLister.
func NewDistccLister(indexer cache.Indexer) DistccLister {
	return &distccLister{indexer: indexer}
}

// List lists all Distccs in the indexer.
func (s *distccLister) List(selector labels.Selector) (ret []*v1alpha1.Distcc, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Distcc))
	})
	return ret, err
}

// Distccs returns an object that can list and get Distccs.
func (s *distccLister) Distccs(namespace string) DistccNamespaceLister {
	return distccNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DistccNamespaceLister helps list and get Distccs.
type DistccNamespaceLister interface {
	// List lists all Distccs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Distcc, err error)
	// Get retrieves the Distcc from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Distcc, error)
	DistccNamespaceListerExpansion
}

// distccNamespaceLister implements the DistccNamespaceLister
// interface.
type distccNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Distccs in the indexer for a given namespace.
func (s distccNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Distcc, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Distcc))
	})
	return ret, err
}

// Get retrieves the Distcc from the indexer for a given namespace and name.
func (s distccNamespaceLister) Get(name string) (*v1alpha1.Distcc, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("distcc"), name)
	}
	return obj.(*v1alpha1.Distcc), nil
}
