/*
Copyright 2022.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BucketSpec defines the desired state of Bucket
type BucketSpec struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Format:=string

	// Enabling versioning
	Versioning string `json:"versioning"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Format:=string

	// Enabling encryption
	Encrypt string `json:"encryption"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Format:=string

	// Enabling public access
	Public string `json:"public"`
}

// BucketStatus defines the observed state of Bucket
type BucketStatus struct {
	// +optional
	NumberOfObjects int32 `json:"NumberOfObjects"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Bucket is the Schema for the buckets API
type Bucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BucketSpec   `json:"spec,omitempty"`
	Status BucketStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// BucketList contains a list of Bucket
type BucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Bucket `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Bucket{}, &BucketList{})
}
