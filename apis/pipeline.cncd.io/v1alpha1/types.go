package v1alpha1

import (
	"k8s.io/api/core/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PipelineList is a list of Pipelines.
type PipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Pipeline `json:"items"`
}

// Image pull policies
const (
	IfNotPresent ImagePullPolicy = "IfNotPresent"
	Always       ImagePullPolicy = "Always"
)

// Secret types
const (
	// GCPServiceAccountSecret is for secrets sourced from a GCP SA Json key file, which also needs to
	// be defined as an Env variable GOOGLE_APPLICATION_CREDENTIALS
	GCPServiceAccountSecret SecretType = "GCPServiceAcount"
	// SlackTokenSecret is for slack webhook tokens to enabled notifications
	SlackTokenSecret SecretType = "SlackToken"
	// GenericType is for secrets that need no special handling
	GenericType SecretType = "Generic"
)

// Pipeline services
const (
	// MySQL is a PipelineServiceType that runs as a sidecar in an executors pod
	// allowing steps to defines interactions with the database
	MySQL PipelineServiceType = "mysql"
	// Redis is a PipelineServiceType that runs as a sidecar in an executors pod, enabling
	// caching capabilities in pipeline steps
	Redis PipelineServiceType = "cache"
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
	Spec   PipelineSpec `json:"spec"`
	Status State        `json:"state,omitempty"`
}

// PipelineSpec describes the specification for a Cloud Native Continous Delivery pipeline using Kubernetes as a build manager
// A PipelineSpec closely reflects the structure of a drone.yaml pipeline, along with the server's docker-compose configuration
type PipelineSpec struct {
	// Selector is how the target will be selected
	Selector map[string]string `json:"selector,omitempty"`
	// Environment is an array of environment variables that are ingested by the Pipeline executor and agents
	EnvVars map[string]string `json:"env"`
	// Volumes is the list of Kubernetes volumes that can be mounted by the driver and/or executors
	Volumes []apiv1.Volume `json:"volumes,omitempty"`
	// Pipeline defines how the Pipeline daemon should run
	Pipeline DriverSpec `json:"pipeline"`
	// Agent defines a Cloud Native Continous Delivery build executor
	Agent AgentSpec `json:"agent,omitempty"`
}

// DriverSpec is the specification for a pipeline daemon server
// It is referred to as a Pipeline instance
type DriverSpec struct {
	// PipelinePodSpec references the base spec for all pods
	PipelinePodSpec
	// Pipeline carries the build stages for the agent to complete
	Steps []Step `json:"steps,omitempty"`
	// Services define sidecar pods to run with Agents (i.e. databases)
	Services []ServiceSideCar `json:"services,omitempty"`
}

// AgentSpec is the specification of the pipeline executor; mirroring many
type AgentSpec struct {
	// PipelinePodSpec references the base spec for all pods
	PipelinePodSpec
	// Number of instances to deploy for a Prometheus deployment.
	Replicas *int32 `json:"replicas,omitempty"`
	// Privileged allows for the executor to use the docker daemon to build images
	Privileged bool `json:"privileged,omitempty"`
}

// PipelinePodSpec defines common things that can be customized for a Pipeline driver or executor pod
type PipelinePodSpec struct {
	// Standard object’s metadata. More info:
	// http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
	// Metadata Labels and Annotations gets propagated to the prometheus pods.
	PodMetadata *metav1.ObjectMeta `json:"metadata,omitempty"`
	// Version of Prometheus to be deployed.
	Version string `json:"version,omitempty"`
	// Image is the container image to use
	Image string `json:"defaultImage,omitempty"`
	// Define resources requests and limits for single Pods.
	Resources v1.ResourceRequirements `json:"resources,omitempty"`
	// ConfigMaps carries information of other ConfigMaps to add to the pod
	ConfigMaps []NamePath `json:"configMaps,omitempty"`
	// Secrets carries information of secrets to add to the pod
	Secrets []Secret `json:"secrets,omitempty"`
	// EnvVars carries the environment variables to add to the pod
	EnvVars map[string]string `json:"env,omitempty"`
	// VolumeMounts specifies the volumes listed in ".spec.volumes" to mount into the main container's filesystem
	VolumeMounts []apiv1.VolumeMount `json:"volumeMounts,omitempty"`
	// ServiceAccountName is the name of the ServiceAccount to use to run the
	// Prometheus Pods.
	ServiceAccountName NamePath `json:"serviceAccountName,omitempty"`
}

// NamePath is a pair of a name and a path to which the named objects should be mounted to
type NamePath struct {
	Name      string `json:"name"`
	mountPath string `json:"mountPath"`
}

// ImagePullPolicy defines the policy of which to handle images
type ImagePullPolicy string

// SecretType tells the type of a secret
type SecretType string

// Secret defines the location, kind, and type of a kubernetes secret object
type Secret struct {
	// Name is the label for a single secret
	Name string     `json:"name"`
	Path string     `json:"path"`
	Type SecretType `json:"secretType"`
}

// PipelineServiceType is the type of sidecar to run in the executors pod
// common uses include databases, caching, and dependent microservices
type PipelineServiceType string

// PipelineService is a sidecar container to run while agent executes pipeline stages
type PipelineService struct {
	// PipelinePodSpec references the base spec for all pods
	PipelinePodSpec
	// ServiceType is the type of service that will run as a sidecar during pipeline execution
	Type []PipelineServiceType `json:serviceType"`
}

type ServiceSideCar struct {
	// PipelinePodSpec references the base spec for all pods
	PipelinePodSpec
	// Name is the container name
	Name string              `json:"name"`
	Type PipelineServiceType `json:"serviceType"`
}

// Steps defines pipeline steps (pipeline execution)
type Step struct {
	// name is the label for a single pipeline step
	Name string `json:"name,omitempty"`
	// Image is the container image to use
	Image string `json:"image,omitempty"`
	// ImagePullPolicy defines the pull policy
	ImagePullPolicy ImagePullPolicy `json:"imagePullPolicy,omitempty"`
	// Command is an array of strings defining a custom command to execute
	Command []string `json:"commands,omitempty"`
	// Entrypoint is a list of commands the executor runs as the image entrypoint
	Entrypoint []string `json:"entrypoint,omitempty"`
	// EnvVars defines environment variables ingested by the build executor
	EnvVars map[string]string `json:"env,omitempty"`
	// Secrets define step-specific secrets
	Secrets []Secret `json:"secrets,omitempty"`
	// WorkingDir is the directory on the agent host to execute commands from
	WorkingDir string `json:"workingDir,omitempty"`
	// OnSuccess determines whether to take action on success
	OnSuccess bool `json:"onSuccess,omitempty"`
	// OnFailure determines whether to take action on failure
	OnFailure bool `json:"onFailure,omitempty"`
}

// Auth defines registry authentication credentials.
type Auth struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
}

// Conn defines a container network connection.
type Conn struct {
	Name    string   `json:"name"`
	Aliases []string `json:"aliases"`
}

type State struct {
	// ExitCode is the container exit code
	ExitCode int `json:"exitCode,omitempy"`
	// Container exited, true or false
	Exited bool `json:"exited"`
	// Container is oom killed, true or false
	OOMKilled bool `json:"oom_killed"`
}
