package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AgentList is a list of Agents.
type AgentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Agent `json:"items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Agent is a duke resource defining a CI lifecycle
type Agent struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the ddesired behaviour of the pod terminator.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status
	// +optional
	Spec AgentSpec `json:"spec"`
}

// PipelineSpec is the spec for a Agent resource
type AgentSpec struct {
	// selector is how the target will be selected
	selector map[string]string `json:"selector,omitempty"`
	// sourceRepository is the location of a Deployments code
	sourceRepository string `json:"sourceRepository,omitempty"`
	// image is the Dockerfile used by the agent
	image   string  `json:"image,omitempty"`
	secrets Secrets `json:"secrets,omitempty"`
	worker  Worker  `json:"worker,omitempty"`
}

// Secrets defines the location, kind, and type of a kubernetes secret object
type Secrets struct {
	// name is the label for a single secret
	name  string `json:"name,omitempty"`
	alias string `json:"alias,omitempty"`
	// TODO
}

// Worker defines the configuration for a pipeline executor
type Worker struct {
	// replicas controllers how many executors the operator will run at one given time
	replicas int32 `json:"replicas,omitempty"`
	// resources define k8 style cpu and memory limits & requests for executors
	resources Resources `json:"resources,omitempty"`
}

// Resources sets the limits and requests for a container
type Resources struct {
	Requests CPUAndMem `json:"requests,omitempty"`
	Limits   CPUAndMem `json:"limits,omitempty"`
}

// CPUAndMem defines how many cpu and ram the container will request/limit
type CPUAndMem struct {
	CPU    string `json:"cpu,omitempty"`
	Memory string `json:"memory,omitempty"`
}

// Status has the status of the cluster
type AgentStatus struct {
	Phase      Phase       `json:"phase"`
	Conditions []Condition `json:"conditions"`
}

// Phase of the pipeline status
type Phase string

// Condition saves the state information of the pipeline
type Condition struct {
	Type           ConditionType `json:"type"`
	Reason         string        `json:"reason"`
	TransitionTime string        `json:"transitionTime"`
}

// ConditionType defines the condition that the pipelinne can have
type ConditionType string
