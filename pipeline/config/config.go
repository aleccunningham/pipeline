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

import (
	"fmt"

	"github.com/marjoram/pipeline/apis/pipeline.cncd.io/v1alpha1"
)

// PipelineAnnotationOption returns a spark-submit option for a driver annotation of the given key and value.
func GetPipelineAnnotationOption(key string, value string) string {
	return fmt.Sprintf("%s%s=%s", PipelineAnnotationKeyPrefix, key, value)
}

// GetAgentAnnotationOption returns a spark-submit option for an executor annotation of the given key and value.
func GetAgentAnnotationOption(key string, value string) string {
	return fmt.Sprintf("%s%s=%s", AgentAnnotationKeyPrefix, key, value)
}

// GetPipelineEnvVarConfOptions returns a list of spark-submit options for setting driver environment variables.
func GetDriverEnvVarConfOptions(app *v1alpha1.Pipeline) []string {
	var envVarConfOptions []string
	for key, value := range app.Spec.Pipeline.EnvVars {
		envVar := fmt.Sprintf("%s%s=%s", PipelineEnvVarConfigKeyPrefix, key, value)
		envVarConfOptions = append(envVarConfOptions, envVar)
	}
	return envVarConfOptions
}

// GetExecutorEnvVarConfOptions returns a list of spark-submit options for setting executor environment variables.
func GetExecutorEnvVarConfOptions(app *v1alpha1.Pipeline) []string {
	var envVarConfOptions []string
	for key, value := range app.Spec.Agent.EnvVars {
		envVar := fmt.Sprintf("%s%s=%s", AgentEnvVarConfigKeyPrefix, key, value)
		envVarConfOptions = append(envVarConfOptions, envVar)
	}
	return envVarConfOptions
}
