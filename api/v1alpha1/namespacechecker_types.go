package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// NamespaceCheckerSpec defines the desired state of NamespaceChecker
type NamespaceCheckerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Namespaces         []string `json:"namespaces"`
	ConfigMapNames     []string `json:"configMapNames"`
	ConfigMapNamespace string   `json:"configMapsNamespace"`
	SecretsNames       []string `json:"secretNames"`
	SecretsNamespace   string   `json:"secretsNamespace"`
}

// NamespaceCheckerStatus defines the observed state of NamespaceChecker
type NamespaceCheckerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	NamespacesExist  map[string]bool              `json:"namespacesExist"`
	ConfigMapsExists map[string]bool              `json:"configMapsExists"`
	SecretsExists    map[string]bool              `json:"secretsExists"`
	ConfigMapsData   map[string]map[string]string `json:"configMapsData,omitempty"`
	SecretsData      map[string]map[string][]byte `json:"secretsData,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// NamespaceChecker is the Schema for the namespacecheckers API
type NamespaceChecker struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NamespaceCheckerSpec   `json:"spec,omitempty"`
	Status NamespaceCheckerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NamespaceCheckerList contains a list of NamespaceChecker
type NamespaceCheckerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []NamespaceChecker `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NamespaceChecker{}, &NamespaceCheckerList{})
}
