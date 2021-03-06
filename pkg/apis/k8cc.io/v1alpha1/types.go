/*
Copyright 2017 The Kubernetes Authors.

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

package v1alpha1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=distcc

// Distcc is a specification for a Distcc resource
type Distcc struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DistccSpec   `json:"spec"`
	Status DistccStatus `json:"status"`
}

// DistccSpec is the spec for a Distcc resource
type DistccSpec struct {
	DeploymentName  string                `json:"deploymentName"`
	ServiceName     string                `json:"serviceName"`
	Selector        *metav1.LabelSelector `json:"selector,omitempty"`
	MinReplicas     *int32                `json:"minReplicas,omitempty"`
	MaxReplicas     int32                 `json:"maxReplicas"`
	UserReplicas    int32                 `json:"userReplicas"`
	LeaseDuration   metav1.Duration       `json:"leaseDuration"`
	DownscaleWindow *metav1.Duration      `json:"downscaleWindow"`
	Template        v1.PodTemplateSpec    `json:"template"`
}

// DistccStatus is the status for a Distcc resource
type DistccStatus struct {
	LastUpdateTime *metav1.Time `json:"lastUpdateTime,omitempty"`
	LastScaleTime  *metav1.Time `json:"lastScaleTime,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=distccs

// DistccList is a list of Distcc resources
type DistccList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Distcc `json:"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=distccclaim

// DistccClaim represents a request for some Distcc resources from a user
type DistccClaim struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DistccClaimSpec   `json:"spec"`
	Status DistccClaimStatus `json:"status"`
}

// DistccClaimSpec is the spec for a DistccClaim resource
type DistccClaimSpec struct {
	DistccName string `json:"distccName"`
	UserName   string `json:"userName"`
}

// DistccClaimStatus is the status for a DistccClaim resource
type DistccClaimStatus struct {
	ExpirationTime *metav1.Time `json:"expirationTime"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=distccclaims

// DistccClaimList is a list of DistccClaim resources
type DistccClaimList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []DistccClaim `json:"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=distccclient

// DistccClient is a specification for a DistccClient resource
type DistccClient struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec DistccClientSpec `json:"spec"`
}

// DistccClientSpec is the spec for a DistccClient resource
type DistccClientSpec struct {
	LeaseDuration metav1.Duration       `json:"leaseDuration"`
	Selector      *metav1.LabelSelector `json:"selector,omitempty"`
	Template      v1.PodTemplateSpec    `json:"template"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=distccclients

// DistccClientList is a list of DistccClient resources
type DistccClientList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []DistccClient `json:"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=distccclientclaim

// DistccClientClaim represents a request for a client POD from a user
type DistccClientClaim struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DistccClientClaimSpec   `json:"spec"`
	Status DistccClientClaimStatus `json:"status"`
}

// DistccClientClaimSpec contains info about a user lease
type DistccClientClaimSpec struct {
	DistccClientName string   `json:"distccClientName"`
	UserName         string   `json:"userName"`
	Secrets          []Secret `json:"secrets"`
}

// Secret contains the name and the reference to a secret in the namespace
type Secret struct {
	Name         string                `json:"name"`
	VolumeSource v1.SecretVolumeSource `json:"volumeSource"`
}

// DistccClientClaimStatus is the status for a DistccClientClaim resource
type DistccClientClaimStatus struct {
	ExpirationTime *metav1.Time             `json:"expirationTime"`
	Deployment     *v1.LocalObjectReference `json:"deployment"`
	Service        *v1.LocalObjectReference `json:"service"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=distccclientclaims

// DistccClientClaimList is a list of DistccClientClaim resources
type DistccClientClaimList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []DistccClientClaim `json:"items"`
}
