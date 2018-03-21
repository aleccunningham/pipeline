package k8s

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"

	pipelinev1alpha1 "github.com/marjoram/pipeline-operator/apis/pipeline.cncd.io/v1alpha1"
	pipelineclientset "github.com/marjoram/pipeline-operator/client/k8s/clientset/versioned"
	"github.com/marjoram/pipeline-operator/log"
)

// Pipeline the RF service that knows how to interact with k8s to get them
type Pipeline interface {
	// ListRedisFailovers lists the pipelines on a cluster.
	ListPipelines(namespace string, opts metav1.ListOptions) (*pipelinev1alpha1.PipelineList, error)
	// WatchPipelines watches the pipelines on a cluster.
	WatchPipelines(namespace string, opts metav1.ListOptions) (watch.Interface, error)
}

// PipelineService is the Pipeline service implementation using API calls to kubernetes.
type PipelineService struct {
	crdClient pipelineclientset.Interface
	logger    log.Logger
}

// NewRedisFailoverService returns a new Workspace KubeService.
func NewPipelineService(crdcli pipelineclientset.Interface, logger log.Logger) *PipelineService {
	logger = logger.With("service", "k8s.pipeline")
	return &RedisFailoverService{
		crdClient: crdcli,
		logger:    logger,
	}
}

// ListRedisFailovers satisfies redisfailover.Service interface.
func (r *PipelineService) ListPipelines(namespace string, opts metav1.ListOptions) (*pipelinev1alpha1.PipelineList, error) {
	return r.crdClient.Storage().Pipelines(namespace).List(opts)
}

// WatchRedisFailovers satisfies redisfailover.Service interface.
func (r *PipelineService) WatchPipelines(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return r.crdClient.Storage().Pipelines(namespace).Watch(opts)
}
