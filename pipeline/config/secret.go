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
	"path/filepath"

	"github.com/marjoram/pipeline/apis/pipeline.cncd.io/v1alpha1"
)

// GetPipelineSecretConfOptions returns a list of spark-submit options for mounting driver secrets.
func GetPipelineSecretConfOptions(app *v1alpha1.Pipeline) []string {
	var secretConfOptions []string
	for _, s := range app.Spec.Pipeline.Secrets {
		conf := fmt.Sprintf("%s%s=%s", PipelineDriverSecretKeyPrefix, s.Name, s.Path)
		secretConfOptions = append(secretConfOptions, conf)
		if s.Type == v1alpha1.GCPServiceAccountSecret {
			conf = fmt.Sprintf(
				"%s%s=%s",
				PipelineDriverEnvVarConfigKeyPrefix,
				GoogleApplicationCredentialsEnvVar,
				filepath.Join(s.Path, ServiceAccountJSONKeyFileName))
			secretConfOptions = append(secretConfOptions, conf)
	}
	return secretConfOptions
}

// GetExecutorSecretConfOptions returns a list of spark-submit options for mounting executor secrets.
func GetAgentSecretConfOptions(app *v1alpha1.Pipeline) []string {
	var secretConfOptions []string
	for _, s := range app.Spec.Agent.Secrets {
		conf := fmt.Sprintf("%s%s=%s", AgentSecretKeyPrefix, s.Name, s.Path)
		secretConfOptions = append(secretConfOptions, conf)
		if s.Type == v1alpha1.GCPServiceAccountSecret {
			conf = fmt.Sprintf(
				"%s%s=%s",
				AgentEnvVarConfigKeyPrefix,
				GoogleApplicationCredentialsEnvVar,
				filepath.Join(s.Path, ServiceAccountJSONKeyFileName))
			secretConfOptions = append(secretConfOptions, conf)
	}
	return secretConfOptions
}
