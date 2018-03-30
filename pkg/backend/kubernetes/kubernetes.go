package kubernetes

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"os"
	"time"

	"path"

	. "github.com/bannzaicloud/banzai-types/components"
	"github.com/marjoram/pipeline/pipeline/backend"
	log "github.com/sirupsen/logrus"
	validator "gopkg.in/go-playground/validator.v8"
)

type engine struct {
	client *client.Client
}

// New returns a new Kubernetes Engine.
func New(k8sCli *client.Client) backend.Engine {
	return &engine{
		client: k8sCli,
	}
}

// Setup the pipeline environment.
func (e *engine) Setup(*backend.Config) error {
	// POST /api/v1/namespaces
	return nil
}

// Start the pipeline step.
func (e *engine) Exec(*backend.Step) error {
	ctx := context.Background()
	// Start agent with a Step as the cmd
	return e.client.CreatePod(ctx, agent.Name, *backend.Step)
}

// Wait for the pipeline step to complete and returns
// the completion results.
func (e *engine) Wait(*backend.Step) (*backend.State, error) {
	ctx := context.Background()
	// Start agent and sleep
	agent, err := e.client.CreateAndWaitPod()
	if err != nil {
		return nil, err
	}
	
	info, err := e.client.PodInspect(ctx, agent.Name)
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

// Tail the pipeline step logs.
func (e *engine) Tail(*backend.Step) (io.ReadCloser, error) {
	// GET /api/v1/namespaces/{namespace}/pods/{name}/log
	return nil, nil
}

// Destroy the pipeline environment.
func (e *engine) Destroy(*backend.Config) error {
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
	return nil
}
