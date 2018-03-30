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

package build

import (
	"context"
	"io"

	"github.com/GoogleCloudPlatform/skaffold/pkg/skaffold/build/tag"
	"github.com/GoogleCloudPlatform/skaffold/pkg/skaffold/config"
)

// BuildResult holds the results of builds
type BuildResult struct {
	Builds []Build
}

// Build is the result corresponding to each Artifact built.
type Build struct {
	ImageName string
	Tag       string
	Artifact  *config.Artifact // The artifact used in the build.
}

// Builder is an interface to the Build API of Skaffold.
// It must build and make the resulting image accesible to the cluster.
// This could include pushing to a authorized repository or loading the nodes with the image.
// If artifacts is supplied, the builder should only rebuild those artifacts.
type Builder interface {
	Build(ctx context.Context, out io.Writer, tagger tag.Tagger, artifacts []*config.Artifact) (*BuildResult, error)
}
