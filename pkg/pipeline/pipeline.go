package pipeline

import (
	"context"
	"time"

	"golang.org/x/sync/errgroup"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/marjoram/pipeline/pipeline/backend"
	"github.com/marjoram/pipeline/pipeline/multipart"
)

type (
	// State defines the pipeline and process state.
	State struct {
		// Global state of the pipeline.
		Pipeline struct {
			// Pipeline time started
			Time int64 `json:"time"`
			// Current pipeline step
			Step *backend.Step `json:"step"`
			// Current pipeline error state
			Error error `json:"error"`
		}

		// Current process state.
		Process *backend.State
	}
)

// Runtime is a configuration runtime.
type Runtime struct {
	err     error
	spec    *backend.Config
	engine  backend.Engine
	started int64
	k8Cli   kubernetes.Interface

	ctx    context.Context
	tracer Tracer
	logger Logger
}

// New returns a new runtime using the specified runtime
// configuration and runtime engine.
func New(spec *backend.Config, opts ...Option) *Runtime {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	k8Cli, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	r := new(Runtime)
	r.k8Cli = k8Cli
	r.spec = spec
	r.ctx = context.Background()
	for _, opts := range opts {
		opts(r)
	}
	return r
}

// Run starts the runtime and waits for it to complete.
func (r *Runtime) Run() error {
	defer func() {
		r.engine.Destroy(r.spec)
	}()

	r.started = time.Now().Unix()
	if err := r.engine.Setup(r.spec); err != nil {
		return err
	}

	for _, stage := range r.spec.Stages {
		select {
		case <-r.ctx.Done():
			return ErrCancel
		case err := <-r.execAll(stage.Steps):
			if err != nil {
				r.err = err
			}
		}
	}

	return r.err
}

//
//
//

func (r *Runtime) execAll(procs []*backend.Step) <-chan error {
	var g errgroup.Group
	done := make(chan error)

	for _, proc := range procs {
		proc := proc
		g.Go(func() error {
			return r.exec(proc)
		})
	}

	go func() {
		done <- g.Wait()
		close(done)
	}()
	return done
}

//
//
//

func (r *Runtime) exec(proc *backend.Step) error {
	switch {
	case r.err != nil && proc.OnFailure == false:
		return nil
	case r.err == nil && proc.OnSuccess == false:
		return nil
	}

	if r.tracer != nil {
		state := new(State)
		state.Pipeline.Time = r.started
		state.Pipeline.Error = r.err
		state.Pipeline.Step = proc
		state.Process = new(backend.State) // empty
		if err := r.tracer.Trace(state); err == ErrSkip {
			return nil
		} else if err != nil {
			return err
		}
	}

	if err := r.engine.Exec(proc); err != nil {
		return err
	}

	if r.logger != nil {
		rc, err := r.engine.Tail(proc)
		if err != nil {
			return err
		}

		go func() {
			r.logger.Log(proc, multipart.New(rc))
			rc.Close()
		}()
	}

	if proc.Detached {
		return nil
	}

	wait, err := r.engine.Wait(proc)
	if err != nil {
		return err
	}

	if r.tracer != nil {
		state := new(State)
		state.Pipeline.Time = r.started
		state.Pipeline.Error = r.err
		state.Pipeline.Step = proc
		state.Process = wait
		if err := r.tracer.Trace(state); err != nil {
			return err
		}
	}

	if wait.OOMKilled {
		return &OomError{
			Name: proc.Name,
			Code: wait.ExitCode,
		}
	} else if wait.ExitCode != 0 {
		return &ExitError{
			Name: proc.Name,
			Code: wait.ExitCode,
		}
	}
	return nil
}
