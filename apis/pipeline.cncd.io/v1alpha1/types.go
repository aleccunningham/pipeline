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
	Spec PipelineSpec `json:"spec"`
}

// PipelineSpec is the spec for a Duked resource
type PipelineSpec struct {
	// Selector is how the target will be selected
	Selector map[string]string `json:"selector,omitempty"`
	// sourceRepository is the location of a Deployments code
	sourceRepository string `json:"sourceRepository,omitempty"`
	// Dockerfile specifies which dockerfile to use from sourceRepository
	Dockerfile string      `json:"dockerfile,omitempty"`
	Volume     Volume      `json:"volume,omitempty"`
	Volumes    Volumes     `json:"volumes,omitempty"`
	SSH        SSHSettings `json:"ssh,omitempty"`
	Steps      Steps       `json:"steps,omitempty"`
	Notify     Notify      `json:"notify,omitempty"`
	When       RunWhen     `json:"when,omitempty"`
	Config     Config      `json:"config,omitempty"`
}

// SSHSettings defines operations for using SSH keys with duke
type SSHSettings struct {
	// hostPath specifies the location of the .ssh/ directory
	hostPath string `json:"hostPath,omitempty"`
}

// Steps defines pipeline steps (pipeline execution)
type Steps struct {
	// name is the label for a single pipeline step
	name string `json:"name,omitempty"`
	// image is the docker builder image to run steps in
	image string `json:"image,omitempty"`
	// commands is an array of strings defining a custom command to execute
	commands map[string]string `json:"commands,omitempty"`
	// env defines environment variables ingested by the build executor
	env     map[string]string `json:"env,omitempty"`
	secrets Secrets           `json:"secrets,omitempty"`
}

// Secrets defines the location, kind, and type of a kubernetes secret object
type Secrets struct {
	// name is the label for a single secret
	name string `json:"name,omitempty"`
	// TODO
}

// Notify
type Notify struct {
	When  []NotifyWhen  `json:"when,omitempty"`
	Where []NotifyWhere `json:"where,omitempty"`
}

// NotifyWhen defines conditionals for notifications
type NotifyWhen struct {
	// event holds all possible event-based conditionals i.e. "on_success"
	Events []Event `json:"event,omitempty"`
}

// NotifyWhere defines the external services to POST to for notifications
// New providers can be added via satisfying notify.Interface
type NotifyWhere struct {
	// slack integration
	Slack NotifyWhereSlack `json:"slack,omitempty"`
	// email integration
	Email NotifyWhereEmail `json:"email,omitempty"`
}

// NotifyWhereSlack defines notifications via slack
type NotifyWhereSlack struct {
	// channel is a list of slack channels
	Channel map[string]string `json:"channel,omitempty"`
	// TODO allow fromSecret
	// token is the workspace-specific slack token for authentication
	Token string `json:"token,omitempty"`
}

// NotifyWhereEmail defines notifications via email
type NotifyWhereEmail struct {
	// from_address is the outgoing email address
	fromAddress string `json:"from_address,omitempty"`
	// to_address is a list of addresses to notify
	toAddress map[string]string `json:"to_address,omitempty"`
}

// Event is an event used as a when/where conditional
type Event struct {
	Name  string `json:"name,omitempty"`
	Alias string `json:"name,omitempty"`
}

// RunWhen defines what events trigger the operator to execute the pipeline
type RunWhen struct {
	// event holds all possible event-based conditionals i.e. "on_commit"
	Event []Event `json:"events,omitempty"`
	// branch defines a git branch conditional
	Branch string `json:"branch,omitempty"`
}

// Status has the status of the cluster
type PipelineStatus struct {
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

// Config defines the runtime configuration of a pipeline.
type Config struct {
	Stages   []*Stage   `json:"pipeline"` // pipeline stages
	Networks []*Network `json:"networks"` // network definitions
	Volumes  []*Volume  `json:"volumes"`  // volume definitions
	Secrets  []*Secret  `json:"secrets"`  // secret definitions
}

// Stage denotes a collection of one or more steps.
type Stage struct {
	Name  string  `json:"name,omitempty"`
	Alias string  `json:"alias,omitempty"`
	Steps []*Step `json:"steps,omitempty"`
}

// Step defines a container process.
type Step struct {
	Name         string            `json:"name"`
	Alias        string            `json:"alias,omitempty"`
	Image        string            `json:"image,omitempty"`
	Pull         bool              `json:"pull,omitempty"`
	Detached     bool              `json:"detach,omitempty"`
	Privileged   bool              `json:"privileged,omitempty"`
	WorkingDir   string            `json:"working_dir,omitempty"`
	Environment  map[string]string `json:"environment,omitempty"`
	Labels       map[string]string `json:"labels,omitempty"`
	Entrypoint   []string          `json:"entrypoint,omitempty"`
	Command      []string          `json:"command,omitempty"`
	ExtraHosts   []string          `json:"extra_hosts,omitempty"`
	Volumes      []string          `json:"volumes,omitempty"`
	Tmpfs        []string          `json:"tmpfs,omitempty"`
	Devices      []string          `json:"devices,omitempty"`
	Networks     []Conn            `json:"networks,omitempty"`
	DNS          []string          `json:"dns,omitempty"`
	DNSSearch    []string          `json:"dns_search,omitempty"`
	MemSwapLimit int64             `json:"memswap_limit,omitempty"`
	MemLimit     int64             `json:"mem_limit,omitempty"`
	ShmSize      int64             `json:"shm_size,omitempty"`
	CPUQuota     int64             `json:"cpu_quota,omitempty"`
	CPUShares    int64             `json:"cpu_shares,omitempty"`
	CPUSet       string            `json:"cpu_set,omitempty"`
	OnFailure    bool              `json:"on_failure,omitempty"`
	OnSuccess    bool              `json:"on_success,omitempty"`
	AuthConfig   Auth              `json:"auth_config,omitempty"`
	NetworkMode  string            `json:"network_mode,omitempty"`
	IpcMode      string            `json:"ipc_mode,omitempty"`
	Sysctls      map[string]string `json:"sysctls,omitempty"`
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

// Network defines a container network.
type Network struct {
	Name       string            `json:"name,omitempty"`
	Driver     string            `json:"driver,omitempty"`
	DriverOpts map[string]string `json:"driver_opts,omitempty"`
}

// Volume defines a container volume.
type Volume struct {
	Name       string            `json:"name,omitempty"`
	Driver     string            `json:"driver,omitempty"`
	DriverOpts map[string]string `json:"driver_opts,omitempty"`
}

// Volumes define storage objects used in a pipeline
type Volumes struct {
	Volumes []*Volume
}

// Secret defines a runtime secret
type Secret struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
	Mount string `json:"mount,omitempty"`
	Mask  bool   `json:"mask,omitempty"`
}

// State defines a container state.
type State struct {
	// Container exit code
	ExitCode int `json:"exit_code"`
	// Container exited, true or false
	Exited bool `json:"exited"`
	// Container is oom killed, true or false
	OOMKilled bool `json:"oom_killed"`
}
