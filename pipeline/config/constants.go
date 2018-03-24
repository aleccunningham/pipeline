/*
Copyright 2017 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

const (
	// DefaultPipelineConfDir is the default directory for Pipeline configuration files if not specified.
	// This directory is where the Pipeline ConfigMap is mounted in the driver and executor containers.
	DefaultPipelineConfDir = "/etc/pipeline/conf"
	// PipelineConfigMapVolumeName is the name of the ConfigMap volume of Pipeline configuration files.
	PipelineConfigMapVolumeName = "pipeline-configmap-volume"
	// PipelineConfDirEnvVar is the environment variable to add to the driver and executor Pods that point
	// to the directory where the Pipeline ConfigMap is mounted.
	PipelineConfDirEnvVar = "PIPELINE_CONF_DIR"
)

const (
	// LabelAnnotationPrefix is the prefix of every labels and annotations added by the controller.
	LabelAnnotationPrefix = "pipeline.cncd.io/"
	// PipelineConfigMapAnnotation is the name of the annotation added to the driver and executor Pods
	// that indicates the presence of a Pipeline ConfigMap that should be mounted to the driver and
	// executor Pods with the environment variable PIPELINE_CONF_DIR set to point to the mount path.
	PipelineConfigMapAnnotation = LabelAnnotationPrefix + "pipeline-configmap"
	// AgentConfigMapAnnotation is the name of the annotation added to the driver and executor Pods
	// that indicates the presence of a Hadoop ConfigMap that should be mounted to the driver and
	// executor Pods with the environment variable AGENT_CONF_DIR set to point to the mount path.
	AgentConfigMapAnnotation = LabelAnnotationPrefix + "agent-configmap"
	// GeneralConfigMapsAnnotationPrefix is the prefix of general annotations that specifies the name
	// and mount paths of additional ConfigMaps to be mounted.
	GeneralConfigMapsAnnotationPrefix = LabelAnnotationPrefix + "configmap."
	// VolumesAnnotationPrefix is the prefix of annotations that specify a Volume.
	VolumesAnnotationPrefix = LabelAnnotationPrefix + "volumes."
	// VolumeMountsAnnotationPrefix is the prefix of annotations that specify a VolumeMount.
	VolumeMountsAnnotationPrefix = LabelAnnotationPrefix + "volumemounts."
	// OwnerReferenceAnnotation is the name of the annotation added to the driver and executor Pods
	// that specifies the OwnerReference of the owning SparkApplication.
	OwnerReferenceAnnotation = LabelAnnotationPrefix + "ownerreference"
	// PipelineAppIDLabel is the name of the label used to group API objects, e.g., Pipeline UI service, Pods,
	// ConfigMaps, etc., belonging to the same Pipeline application.
	PipelineAppIDLabel = LabelAnnotationPrefix + "app-id"
	// PipelineAppNameLabel is the name of the label for the PipelineApplication object name.
	PipelineAppNameLabel = LabelAnnotationPrefix + "app-name"
	// LaunchedByPipelineOperatorLabel is a label on agent pods launched through the Pipeline Operator.
	LaunchedByPipelineOperatorLabel = LabelAnnotationPrefix + "launched-by-pipeline-operator"
)

const (
	// PipelineContainerImageKey is the configuration property for specifying the unified container image.
	PipelineContainerImageKey = "pipeline.kubernetes.container.image"
	// PipelineContainerImageKey is the configuration property for specifying the container image pull policy.
	PipelineContainerImagePullPolicyKey = "pipeline.kubernetes.container.image.pullPolicy"
	// PipelineNodeSelectorKeyPrefix is the configuration property prefix for specifying node selector for the pods.
	PipelineNodeSelectorKeyPrefix = "pipeline.kubernetes.node.selector."
	// PipelineDriverContainerImageKey is the configuration property for specifying a custom driver container image.
	PipelineDriverContainerImageKey = "pipeline.kubernetes.driver.container.image"
	// PipelineExecutorContainerImageKey is the configuration property for specifying a custom executor container image.
	PipelineAgentContainerImageKey = "pipeline.kubernetes.agent.container.image"
	// PipelineDriverCoreLimitKey is the configuration property for specifying the hard CPU limit for the driver pod.
	PipelineDriverCoreLimitKey = "pipeline.kubernetes.driver.limit.cores"
	// PipelineDriverCoreLimitKey is the configuration property for specifying the hard CPU limit for the executor pods.
	PipelineAgentCoreLimitKey = "pipeline.kubernetes.agent.limit.cores"
	// PipelineDriverSecretKeyPrefix is the configuration property prefix for specifying secrets to be mounted into the
	// driver.
	PipelineDriverSecretKeyPrefix = "pipeline.kubernetes.driver.secrets."
	// PipelineDriverSecretKeyPrefix is the configuration property prefix for specifying secrets to be mounted into the
	// executors.
	PipelineAgentSecretKeyPrefix = "pipeline.kubernetes.agent.secrets."
	// PipelineDriverEnvVarConfigKeyPrefix is the Pipeline configuration prefix for setting environment variables
	// into the driver.
	PipelineDriverEnvVarConfigKeyPrefix = "pipeline.kubernetes.driverEnv."
	// PipelineAgentEnvVarConfigKeyPrefix is the Pipeline configuration prefix for setting environment variables into the executor.
	PipelineAgentEnvVarConfigKeyPrefix = "pipeline.agentEnv."
	// PipelineDriverAnnotationKeyPrefix is the Pipeline configuration key prefix for annotations on the driver Pod.
	PipelineDriverAnnotationKeyPrefix = "pipeline.kubernetes.driver.annotation."
	// PipelineAgentAnnotationKeyPrefix is the Pipeline configuration key prefix for annotations on the executor Pods.
	PipelineAgentAnnotationKeyPrefix = "pipeline.kubernetes.agent.annotation."
	// PipelineDriverLabelKeyPrefix is the Pipeline configuration key prefix for labels on the driver Pod.
	PipelineDriverLabelKeyPrefix = "pipeline.kubernetes.driver.label."
	// PipelineAgentLabelKeyPrefix is the Pipeline configuration key prefix for labels on the executor Pods.
	PipelineAgentLabelKeyPrefix = "pipeline.kubernetes.agent.label."
	// PipelineDriverPodNameKey is the Pipeline configuration key for driver pod name.
	PipelineDriverPodNameKey = "pipeline.kubernetes.driver.pod.name"
	// PipelineDriverServiceAccountName is the Pipeline configuration key for specifying name of the Kubernetes service
	// account used by the driver pod.
	PipelineDriverServiceAccountName = "pipeline.kubernetes.authenticate.driver.serviceAccountName"
	// PipelineInitContainerImage is the Pipeline configuration key for specifying a custom init-container image.
	PipelineInitContainerImage = "pipeline.kubernetes.initContainer.image"
)

const (
	// GoogleApplicationCredentialsEnvVar is the environment variable used by the
	// Application Default Credentials mechanism. More details can be found at
	// https://developers.google.com/identity/protocols/application-default-credentials.
	GoogleApplicationCredentialsEnvVar = "GOOGLE_APPLICATION_CREDENTIALS"
	// ServiceAccountJSONKeyFileName is the assumed name of the service account
	// Json key file. This name is added to the service account secret mount path to
	// form the path to the Json key file referred to by GOOGLE_APPLICATION_CREDENTIALS.
	ServiceAccountJSONKeyFileName = "key.json"
)
