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

package fake

import (
	v1alpha1 "github.com/mbrt/k8cc/pkg/client/clientset/versioned/typed/k8cc/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeK8ccV1alpha1 struct {
	*testing.Fake
}

func (c *FakeK8ccV1alpha1) Distccs(namespace string) v1alpha1.DistccInterface {
	return &FakeDistccs{c, namespace}
}

func (c *FakeK8ccV1alpha1) DistccClients(namespace string) v1alpha1.DistccClientInterface {
	return &FakeDistccClients{c, namespace}
}

func (c *FakeK8ccV1alpha1) DistccClientClaims(namespace string) v1alpha1.DistccClientClaimInterface {
	return &FakeDistccClientClaims{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeK8ccV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
