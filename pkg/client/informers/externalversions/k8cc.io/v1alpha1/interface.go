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

// This file was automatically generated by informer-gen

package v1alpha1

import (
	internalinterfaces "github.com/mbrt/k8cc/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Distccs returns a DistccInformer.
	Distccs() DistccInformer
	// DistccClaims returns a DistccClaimInformer.
	DistccClaims() DistccClaimInformer
	// DistccClients returns a DistccClientInformer.
	DistccClients() DistccClientInformer
	// DistccClientClaims returns a DistccClientClaimInformer.
	DistccClientClaims() DistccClientClaimInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Distccs returns a DistccInformer.
func (v *version) Distccs() DistccInformer {
	return &distccInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DistccClaims returns a DistccClaimInformer.
func (v *version) DistccClaims() DistccClaimInformer {
	return &distccClaimInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DistccClients returns a DistccClientInformer.
func (v *version) DistccClients() DistccClientInformer {
	return &distccClientInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DistccClientClaims returns a DistccClientClaimInformer.
func (v *version) DistccClientClaims() DistccClientClaimInformer {
	return &distccClientClaimInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}