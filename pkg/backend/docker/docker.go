package docker

import (
	"context"
	"io"
	"io/ioutil"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/marjoram/pipeline/pipeline/backend"
	"github.com/GoogleCloudPlatform/skaffold/pkg/skaffold/docker"
	"github.com/GoogleCloudPlatform/skaffold/pkg/skaffold/constants"
	"github.com/GoogleCloudPlatform/skaffold/pkg/skaffold/build/tag"
	"github.com/GoogleCloudPlatform/skaffold/pkg/skaffold/util"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type engine struct {
	client 		*client.Client
	
	api		docker.DockerAPIClient
	kubeContext 	string
}

// New returns a new Docker Engine using the given client.
func New(cli *client.Client, kubeContext string) backend.Engine {
	return &engine{
		client: cli,
		api:	api,
		kubeContext: kubeContext == constants.DefaultMinikubeContext || kubeContext == DefaultDockerForDesktopContext,
	}
}

// NewEnv returns a new Docker Engine using the client connection
// environment variables.
func NewEnv() (backend.Engine, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}
	return New(cli), nil
}

func (e *engine) Build(ctx context.Context, e.api) (*BuildResult, error) {
	err := docker.RunBuild(ctx, e.api, &docker.BuildOptions{
		ImageName:	initialTag,
		Dockerfile:	artifact.DockerfilePath,
		ContextDir:	artifact.Workspace,
	})
	if err != nil {
		return nil, errors.Wrap(err, "running build")
	}
	digest, err := docker.Digest(ctx, l.api, initialTag)
	if err != nil {
		return nil, errors.Wrap(err, "running build and tag")
	}
	if digest == "" {
			return nil, fmt.Errorf("digest not found")
	}
	tag, err := tagger.GenerateFullyQualifedImageName(".", &tag.TagOptions{
		ImageName:	artifact.ImageName,
		Digest:		digest,
	})
	if err != nil {
			return nil, errors.Wrap(err, "generating tag")
		}
		if err := l.api.ImageTag(ctx, fmt.Sprintf("%s:latest", initialTag), tag); err != nil {
			return nil, errors.Wrap(err, "tagging image")
		}
		if _, err := io.WriteString(out, fmt.Sprintf("Successfully tagged %s\n", tag)); err != nil {
			return nil, errors.Wrap(err, "writing tag status")
		}
		if !*l.LocalBuild.SkipPush {
			if err := docker.RunPush(ctx, l.api, tag, out); err != nil {
				return nil, errors.Wrap(err, "running push")
			}
		}
		res.Builds = append(res.Builds, Build{
			ImageName: artifact.ImageName,
			Tag:       tag,
			Artifact:  artifact,
		})
	}

	return res, nil
}

func (e *engine) Setup(conf *backend.Config) error {
	for _, vol := range conf.Volumes {
		_, err := e.client.VolumeCreate(noContext, volume.VolumesCreateBody{
			Name:       vol.Name,
			Driver:     vol.Driver,
			DriverOpts: vol.DriverOpts,
			// Labels:     defaultLabels,
		})
		if err != nil {
			return err
		}
	}
	for _, network := range conf.Networks {
		_, err := e.client.NetworkCreate(noContext, network.Name, types.NetworkCreate{
			Driver:  network.Driver,
			Options: network.DriverOpts,
			// Labels:  defaultLabels,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *engine) Exec(proc *backend.Step) error {
	ctx := context.Background()

	config := toConfig(proc)
	hostConfig := toHostConfig(proc)

	// create pull options with encoded authorization credentials.
	pullopts := types.ImagePullOptions{}
	if proc.AuthConfig.Username != "" && proc.AuthConfig.Password != "" {
		pullopts.RegistryAuth, _ = encodeAuthToBase64(proc.AuthConfig)
	}

	// automatically pull the latest version of the image if requested
	// by the process configuration.
	if proc.Pull {
		rc, perr := e.client.ImagePull(ctx, config.Image, pullopts)
		if perr == nil {
			io.Copy(ioutil.Discard, rc)
			rc.Close()
		}
		// fix for drone/drone#1917
		if perr != nil && proc.AuthConfig.Password != "" {
			return perr
		}
	}

	_, err := e.client.ContainerCreate(ctx, config, hostConfig, nil, proc.Name)
	if client.IsErrImageNotFound(err) {
		// automatically pull and try to re-create the image if the
		// failure is caused because the image does not exist.
		rc, perr := e.client.ImagePull(ctx, config.Image, pullopts)
		if perr != nil {
			return perr
		}
		io.Copy(ioutil.Discard, rc)
		rc.Close()

		_, err = e.client.ContainerCreate(ctx, config, hostConfig, nil, proc.Name)
	}
	if err != nil {
		return err
	}

	if len(proc.NetworkMode) == 0 {
		for _, net := range proc.Networks {
			err = e.client.NetworkConnect(ctx, net.Name, proc.Name, &network.EndpointSettings{
				Aliases: net.Aliases,
			})
			if err != nil {
				return err
			}
		}
	}

	// if proc.Network != "host" { // or bridge, overlay, none, internal, container:<name> ....
	// 	err = e.client.NetworkConnect(ctx, proc.Network, proc.Name, &network.EndpointSettings{
	// 		Aliases: proc.NetworkAliases,
	// 	})
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	return e.client.ContainerStart(ctx, proc.Name, startOpts)
}

func (e *engine) Kill(proc *backend.Step) error {
	return e.client.ContainerKill(noContext, proc.Name, "9")
}

func (e *engine) Wait(proc *backend.Step) (*backend.State, error) {
	_, err := e.client.ContainerWait(noContext, proc.Name)
	if err != nil {
		// todo
	}

	info, err := e.client.ContainerInspect(noContext, proc.Name)
	if err != nil {
		return nil, err
	}
	if info.State.Running {
		// todo
	}

	return &backend.State{
		Exited:    true,
		ExitCode:  info.State.ExitCode,
		OOMKilled: info.State.OOMKilled,
	}, nil
}

func (e *engine) Tail(proc *backend.Step) (io.ReadCloser, error) {
	logs, err := e.client.ContainerLogs(noContext, proc.Name, logsOpts)
	if err != nil {
		return nil, err
	}
	rc, wc := io.Pipe()

	go func() {
		stdcopy.StdCopy(wc, wc, logs)
		logs.Close()
		wc.Close()
		rc.Close()
	}()
	return rc, nil
}

func (e *engine) Destroy(conf *backend.Config) error {
	for _, stage := range conf.Stages {
		for _, step := range stage.Steps {
			e.client.ContainerKill(noContext, step.Name, "9")
			e.client.ContainerRemove(noContext, step.Name, removeOpts)
		}
	}
	for _, volume := range conf.Volumes {
		e.client.VolumeRemove(noContext, volume.Name, true)
	}
	for _, network := range conf.Networks {
		e.client.NetworkRemove(noContext, network.Name)
	}
	return nil
}

func (e *engine) Close() error {
	return e.client.Close()
}

var (
	noContext = context.Background()

	startOpts = types.ContainerStartOptions{}

	removeOpts = types.ContainerRemoveOptions{
		RemoveVolumes: true,
		RemoveLinks:   false,
		Force:         false,
	}

	logsOpts = types.ContainerLogsOptions{
		Follow:     true,
		ShowStdout: true,
		ShowStderr: true,
		Details:    false,
		Timestamps: false,
	}
)
