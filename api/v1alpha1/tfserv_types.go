/*
Copyright 2019 Dan Rusei.

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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TfservSpec defines the desired state of Tfserv
type TfservSpec struct {
	GrpcPort int32 `json:"grpcPort,omitempty"`
	//RestPort is the port number listening REST
	RestPort int32 `json:"restPort,omitempty"`
	//Replicas is the number of pod replicas
	Replicas int32 `json:"replicas,omitempty"`
	//GrpcPort is the port number listening RPC
	//ConfigMap is the name of the ConfigMap used for configuration
	ConfigMap string `json:"configMap,omitempty"`
	//ConfigFileName is the name of the config file
	ConfigFileName string `json:"configFileName,omitempty"`
	//ConfigFileLocation is the path to config file
	ConfigFileLocation string `json:"configFileLocation,omitempty"`
	//SecretFileName is the name of the secret file
	SecretFileName string `json:"secretFileName,omitempty"`
	//SecretFileLocation is the path to the Secret file
	SecretFileLocation string `json:"secretFileLocation,omitempty"`
}

// TfservStatus defines the observed state of Tfserv
type TfservStatus struct {
	// A list of pointers to currently running objects.
	Active []corev1.ObjectReference `json:"active,omitempty"`
}

// +kubebuilder:object:root=true

// Tfserv is the Schema for the tfservs API
type Tfserv struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TfservSpec   `json:"spec,omitempty"`
	Status TfservStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TfservList contains a list of Tfserv
type TfservList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Tfserv `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Tfserv{}, &TfservList{})
}
