package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PipelineList is a list of Pipelines.
type PipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Pipeline `json:"items"`
}

// Different types of deployments
const (
	ClusterMode         DeployMode = "cluster"
	ClientMode          DeployMode = "client"
	InClusterClientMode DeployMode = "in-cluster-client"
)

// An enumeration of the secret types supported
const (
	// GCPServiceAccountSecret is for secrets sourced from a GCP SA Json key file, which also needs to
	// be defined as an Env variable GOOGLE_APPLICATION_CREDENTIALS
	GCPServiceAccountSecret SecretType = "GCPServiceAcount"
	// SlackTokenSecret is for slack webhook tokens to enabled notifications
	SlackTokenSecret SecretType = "SlackToken"
	// GenericType is for secrets that need no special handling
	GenericType SecretType = "Generic"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Pipeline is a duke resource defining a CI lifecycle
type Pipeline struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the ddesired behaviour of the pod terminator.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status
	// +optional
	Spec   PipelineSpec   `json:"spec"`
	Status PipelineStatus `json:"status,omitempty"`
}

// PipelineSpec describes the specification for a Cloud Native Continous Delivery pipeline using Kubernetes as a build manager
// It carries all information a pipeline.Run() command uses to run a Pipeline
type PipelineSpec struct {
	// Selector is how the target will be selected
	Selector map[string]string `json:"selector,omitempty"`
	// Environment is an array of environment variables that are ingested by the Pipeline executor and agents
	Environment map[string]string `json:"environment"`
	// Volumes is the list of Kubernetes volumes that can be mounted by the driver and/or executors
	Volumes []apiv1.Volume `json:"volumes,omitempty"`
	// Agent defines a Cloud Native Continous Delivery build executor
	Agent AgentSpec `json:"agent,omitempty"`
}

// PipelinePodSpec defines common things that can be customized for a Pipeline driver or executor pod
type PipelinePodSpec struct {
	// Image is the container image to use
	Image *string `json:"image,omitempty"`
	// ConfigMaps carries information of other ConfigMaps to add to the pod
	ConfigMaps []NamePath `json:"configMaps,omitempty"`
	// Secrets carries information of secrets to add to the pod
	Secrets []Secret `json:"secrets,omitempty"`
	// EnvVars carries the environment variables to add to the pod
	EnvVars map[string]string `json:"envVars,omitempty"`
	// Labels are the Kubernetes labels to be added to the pod
	Labels map[string]string `json:"labels,omitempty"`
	// Annotations are the Kubernetes annotationsn to be added to the pod
	Annotations map[string]string `json:"annotations,omitempty"`
	// VolumeMounts specifies the volumes listed in ".spec.volumes" to mount into the main container's filesystem
	VolumeMounts []apiv1.VolumeMount `json:"volumeMounts,omitempty"`
}

// NamePath is a pair of a name and a path to which the named objects should be mounted to
type NamePath struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

// SecretType tells the type of a secret
type SecretType string

// Secret defines the location, kind, and type of a kubernetes secret object
type Secret struct {
	// Name is the label for a single secret
	Name string     `json:"name"`
	Path string     `json:"path"`
	Type SecretType `json:"secretType"`
}

// AgentSpec is the specification of the pipeline executor; mirroring many
type AgentSpec struct {
	// PipelinePodSpec references the base spec for all pods
	PipelinePodSpec
	// State is the current state of the pipeline being executed
	State State `json:"state,omitempty"`
	// Pipeline carries the build stages for the agent to complete
	Stages []Stage `json:"pipeline,omitempty"`
	// Services define sidecar pods to run with Agents (i.e. databases)
	Services []Service `json:"services,omitempty"`
	// Detached defines whether to run the container with tty or not
	Detached *bool `json:"detached,omitempty"`
	// Privileged allows for the executor to use the docker daemon to build images
	Privileged *bool `json:"privileged,omitempty"`
	// ExtraHosts define external hosts
	ExtraHosts []*string `json:"extraHosts,omitempty"`
}

// Service is a sidecar container to run while agent executes pipeline stages
type Service struct {
	// PipelinePodSpec references the base spec for all pods
	PipelinePodSpec
}

// Stage denotes a collection of one or more steps.
type Stage struct {
	// Name defines the collection of steps for a given stage
	Name string `json:"name,omitempty"`
	// Labels defines labels to group pipeline steps
	Labels map[string]string `json:"labels,omitempty"`
	// Steps is a list of step objects that are each run in isolated pods in sequence
	Steps []*Step `json:"steps,omitempty"`
}

// Steps defines pipeline steps (pipeline execution)
type Step struct {
	// name is the label for a single pipeline step
	Name *string `json:"name,omitempty"`
	// Image is the container image to use
	Image *string `json:"image,omitempty"`
	// Pull defines whether to pull the executors container image
	Pull *bool `json:"pull,omitempty"`
	// Command is an array of strings defining a custom command to execute
	Command []*string `json:"commands,omitempty"`
	// Entrypoint is a list of commands the executor runs as the image entrypoint
	Entrypoint []*string `json:"entrypoint,omitempty"`
	// Labels defines labels to group pipeline steps
	Labels map[string]string `json:"labels,omitempty"`
	// EnvVars defines environment variables ingested by the build executor
	EnvVars map[string]string `json:"envVars,omitempty"`
	// Secrets define step-specific secrets
	Secrets []Secret `json:"secrets,omitempty"`
	// WorkingDir is the directory on the agent host to execute commands from
	WorkingDir *string `json:"workingDir,omitempty"`
}

type State struct {
	// ExitCode is the container exit code
	ExitCode *int `json:"exitCode,omitempy"`
	// Container exited, true or false
	Exited *bool `json:"exited"`
	// Container is oom killed, true or false
	OOMKilled *bool `json:"oom_killed"`
}
