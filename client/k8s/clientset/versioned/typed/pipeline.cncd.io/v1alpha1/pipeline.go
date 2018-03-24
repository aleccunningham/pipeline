package v1alpha1

import (
	v1alpha1 "github.com/marjoram/pipeline/apis/pipeline.cncd.io/v1alpha1"
	scheme "github.com/marjoram/pipeline/client/k8s/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PipelinesGetter has a method to return a PipelineInterface.
// A group's client should implement this interface.
type PipelinesGetter interface {
	Pipelines() PipelineInterface
}

// PipelineInterface has methods to work with Pipeline resources.
type PipelineInterface interface {
	Create(*v1alpha1.Pipeline) (*v1alpha1.Pipeline, error)
	Update(*v1alpha1.Pipeline) (*v1alpha1.Pipeline, error)
	UpdateStatus(*v1alpha1.Pipeline) (*v1alpha1.Pipeline, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Pipeline, error)
	List(opts v1.ListOptions) (*v1alpha1.PipelineList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Pipeline, err error)
	PipelineExpansion
}

// pipelines implements PipelineInterface
type pipelines struct {
	client rest.Interface
}

// newPipelines returns a Pipelines
func newPipelines(c *PipelineV1alpha1Client) *pipelines {
	return &pipelines{
		client: c.RESTClient(),
	}
}

// Get takes name of the pipeline, and returns the corresponding pipeline object, and an error if there is any.
func (c *pipelines) Get(name string, options v1.GetOptions) (result *v1alpha1.Pipeline, err error) {
	result = &v1alpha1.Pipeline{}
	err = c.client.Get().
		Resource("pipelines").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Pipelines that match those selectors.
func (c *pipelines) List(opts v1.ListOptions) (result *v1alpha1.PipelineList, err error) {
	result = &v1alpha1.PipelineList{}
	err = c.client.Get().
		Resource("pipelines").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested pipelines.
func (c *pipelines) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("pipelines").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a pipeline and creates it.  Returns the server's representation of the pipeline, and an error, if there is any.
func (c *pipelines) Create(pipeline *v1alpha1.Pipeline) (result *v1alpha1.Pipeline, err error) {
	result = &v1alpha1.Pipeline{}
	err = c.client.Post().
		Resource("pipelines").
		Body(pipeline).
		Do().
		Into(result)
	return
}

// Update takes the representation of a pipeline and updates it. Returns the server's representation of the pipeline, and an error, if there is any.
func (c *pipelines) Update(pipeline *v1alpha1.Pipeline) (result *v1alpha1.Pipeline, err error) {
	result = &v1alpha1.Pipeline{}
	err = c.client.Put().
		Resource("pipelines").
		Name(pipeline.Name).
		Body(pipeline).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *pipelines) UpdateStatus(pipeline *v1alpha1.Pipeline) (result *v1alpha1.Pipeline, err error) {
	result = &v1alpha1.Pipeline{}
	err = c.client.Put().
		Resource("pipelines").
		Name(pipeline.Name).
		SubResource("status").
		Body(pipeline).
		Do().
		Into(result)
	return
}

// Delete takes name of the pipeline and deletes it. Returns an error if one occurs.
func (c *pipelines) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("pipelines").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *pipelines) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Resource("pipelines").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched pipeline.
func (c *pipelines) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Pipeline, err error) {
	result = &v1alpha1.Pipeline{}
	err = c.client.Patch(pt).
		Resource("pipelines").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
