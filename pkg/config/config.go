/*
Copyright 2018 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	yaml "gopkg.in/yaml.v2"
)

// PipelineConfig is the top level config object
// that is parsed from a skaffold.yaml
//
// APIVersion and Kind are currently reserved for future use.
type PipelineConfig struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`

	Build  BuildConfig  `yaml:"build"`
	Deploy DeployConfig `yaml:"deploy"`
}

// BuildConfig contains all the configuration for the build steps
type BuildConfig struct {
	Artifacts []*Artifact `yaml:"artifacts"`
}

// DeployConfig contains all the configuration needed by the deploy steps
type DeployConfig struct {
	Name     string   `yaml:"name"`
	Selector []string `yaml:"selectors"`
}

// Artifact represents items that need should be built, along with the context in which
// they should be built.
type Artifact struct {
	ImageName      string             `yaml:"imageName"`
	DockerfilePath string             `yaml:"dockerfilePath"`
	Workspace      string             `yaml:"workspace"`
	BuildArgs      map[string]*string `yaml:"buildArgs"`
}

// Parse reads from an io.Reader and unmarshals the result into a SkaffoldConfig.
// The default config argument provides default values for the config,
// which can be overridden if present in the config file.
func Parse(config []byte, defaultConfig *SkaffoldConfig) (*SkaffoldConfig, error) {
	if err := yaml.Unmarshal(config, defaultConfig); err != nil {
		return nil, err
	}

	return defaultConfig, nil
}
