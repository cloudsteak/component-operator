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
	ConfigMapNamespace string   `json:"configMapNamespace"`
	SecretsNames       []string `json:"secretsNames"`
	SecretsNamespace   string   `json:"secretsNamespace"`
}

// NamespaceCheckerStatus defines the observed state of NamespaceChecker
type NamespaceCheckerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	NamespacesExist map[string]bool `json:"namespacesExist"`
	ConfigMapsExist map[string]bool `json:"configMapsExist"`
	SecretsExist    map[string]bool `json:"secretsExist"`
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

	Spec   NamespaceCheckerSpec   `json:"spec,omitempty"`
	Status NamespaceCheckerStatus `json:"status,omitempty"`
}

func init() {
	SchemeBuilder.Register(&NamespaceChecker{}, &NamespaceCheckerList{})
}
