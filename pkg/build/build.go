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
	"fmt"
	"io"

	"github.com/marjoram/skaffold/pkg/build/docker"
	"github.com/marjoram/skaffold/pkg/config"
	"github.com/marjoram/skaffold/pkg/constants"
	"github.com/pkg/errors"
)

// LocalBuilder uses the host docker daemon to build and tag the image
type Builder struct {
	*config.BuildConfig

	api          docker.DockerAPIClient
	localCluster bool
	kubeContext  string
}

// NewBuilder returns an new instance of a Builder
func NewBuilder(cfg *config.BuildConfig, kubeContext string) (*Builder, error) {
	api, err := docker.NewDockerAPIClient(kubeContext)
	if err != nil {
		return nil, errors.Wrap(err, "getting docker client")
	}

	builder := &Builder{
		BuildConfig: cfg,

		kubeContext:  kubeContext,
		api:          api,
		localCluster: kubeContext == constants.DefaultMinikubeContext || kubeContext == constants.DefaultDockerForDesktopContext,
	}
	return builder, nil
}

// Build runs a docker build on the host and tags the resulting image with
// its checksum. It streams build progress to the writer argument.
func (b *Builder) Build(out io.Writer, tagger tag.Tagger, artifacts []*config.Artifact) (*BuildResult, error) {
	if b.localCluster {
		if _, err := fmt.Fprintf(out, "Found [%s] context, using local docker daemon.\n", b.kubeContext); err != nil {
			return nil, errors.Wrap(err, "writing status")
		}
	}
	defer b.api.Close()
	res := &BuildResult{
		Builds: []Build{},
	}
	for _, artifact := range artifacts {
		if artifact.DockerfilePath == "" {
			artifact.DockerfilePath = constants.DefaultDockerfilePath
		}
		initialTag := util.RandomID()
		err := docker.RunBuild(b.api, &docker.BuildOptions{
			ImageName:   initialTag,
			Dockerfile:  artifact.DockerfilePath,
			ContextDir:  artifact.Workspace,
			ProgressBuf: out,
			BuildBuf:    out,
			BuildArgs:   artifact.BuildArgs,
		})
		if err != nil {
			return nil, errors.Wrap(err, "running build")
		}
		digest, err := docker.Digest(b.api, initialTag)
		if err != nil {
			return nil, errors.Wrap(err, "build and tag")
		}
		if digest == "" {
			return nil, fmt.Errorf("digest not found")
		}
		/*
			tag, err := tagger.GenerateFullyQualifiedImageName(".", &tag.TagOptions{
				ImageName: artifact.ImageName,
				Digest:    digest,
			})
			if err != nil {
				return nil, errors.Wrap(err, "generating tag")
			}
		*/
		if err := b.api.ImageTag(context.Background(), fmt.Sprintf("%s:latest", initialTag), tag); err != nil {
			return nil, errors.Wrap(err, "tagging image")
		}
		if _, err := io.WriteString(out, fmt.Sprintf("Successfully tagged %s\n", tag)); err != nil {
			return nil, errors.Wrap(err, "writing tag status")
		}
		res.Builds = append(res.Builds, Build{
			ImageName: artifact.ImageName,
			Tag:       tag,
			Artifact:  artifact,
		})
	}

	return res, nil
}
