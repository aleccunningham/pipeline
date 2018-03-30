package pipeline

import (
	"context"
	"fmt"
	"io"
)

type Result struct{}

// Deployer is the deploy API for the pipeline agent, providing an interface
// that takes built images and deploys them to a kubernetes cluster
type Deployer interface {
	// Deploy should take a build and deploy it to a 
	// kubernetes cluster 
	Deploy(context.Context, io.Writer, *build.BuildResult) (*Result, error)
}
