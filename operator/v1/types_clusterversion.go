package v1 // +genclient
import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// MyOperatorResource is an example operator configuration type
//
// Compatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.
// +openshift:compatibility-gen:internal
type ClusterVersion struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is the standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata"`

	// +kubebuilder:validation:Required
	// +required
	Spec   ClusterVersionSpec   `json:"spec"`
	Status ClusterVersionStatus `json:"status"`
}

type ClusterVersionSpec struct {
	OperatorSpec `json:",inline"`
}

type ClusterVersionStatus struct {
	OperatorStatus `json:",inline"`
}
