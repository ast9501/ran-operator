package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RansliceSpec defines the desired state of Ranslice
type RansliceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	CUdomainName 	string		`json:"cuDomainName"`
	DUdomainName 	string 		`json:"duDomainName"`
	AMFdomainName	string		`json:"amfDomainName"`

}

// RansliceStatus defines the observed state of Ranslice
type RansliceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	
	// State of RAN slice
	State 	string 	`json:"state"`

	// CU address
	CUAddr 	string	`json:"cuAddr"`
	// DU address
	DUAddr	string	`json:"duAddr"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Ranslice is the Schema for the ranslice API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=ranslice,scope=Namespaced
type Ranslice struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RansliceSpec   `json:"spec,omitempty"`
	Status RansliceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RansliceList contains a list of Ranslice
type RansliceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ranslice `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Ranslice{}, &RansliceList{})
}
