/*
Copyright 2023.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// UrlResource1Spec defines the desired state of UrlResource1
type UrlResource1Spec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of UrlResource1. Edit urlresource1_types.go to remove/update
	// +optional
	//TargetURL string `json:"targetURL"`
	TargetURLs []string `json:"targetURLs"`
}

// UrlResource1Status defines the observed state of UrlResource1
type UrlResource1Status struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// ShortURL string `json:"shortURL"`
	URLMap map[string]string `json:"shortURL"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// UrlResource1 is the Schema for the urlresource1s API
type UrlResource1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UrlResource1Spec   `json:"spec,omitempty"`
	Status UrlResource1Status `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// UrlResource1List contains a list of UrlResource1
type UrlResource1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UrlResource1 `json:"items"`
}

func init() {
	SchemeBuilder.Register(&UrlResource1{}, &UrlResource1List{})
}
