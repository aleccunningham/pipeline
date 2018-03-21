package v1alpha1

import (
	v1alpha1 "github.com/marjoram/pipeline/apis/agent.cncd.io/v1alpha1"
	scheme "github.com/marjoram/pipeline/client/k8s/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AgentsGetter has a method to return a AgentInterface.
// A group's client should implement this interface.
type AgentsGetter interface {
	Agents() AgentInterface
}

// AgentInterface has methods to work with Agent resources.
type AgentInterface interface {
	Create(*v1alpha1.Agent) (*v1alpha1.Agent, error)
	Update(*v1alpha1.Agent) (*v1alpha1.Agent, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Agent, error)
	List(opts v1.ListOptions) (*v1alpha1.AgentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Agent, err error)
	AgentExpansion
}

// agents implements AgentInterface
type agents struct {
	client rest.Interface
}

// newAgents returns a Agents
func newAgents(c *AgentV1alpha1Client) *agents {
	return &agents{
		client: c.RESTClient(),
	}
}

// Get takes name of the agent, and returns the corresponding agent object, and an error if there is any.
func (c *agents) Get(name string, options v1.GetOptions) (result *v1alpha1.Agent, err error) {
	result = &v1alpha1.Agent{}
	err = c.client.Get().
		Resource("agents").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Agents that match those selectors.
func (c *agents) List(opts v1.ListOptions) (result *v1alpha1.AgentList, err error) {
	result = &v1alpha1.AgentList{}
	err = c.client.Get().
		Resource("agents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested agents.
func (c *agents) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("agents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a agent and creates it.  Returns the server's representation of the agent, and an error, if there is any.
func (c *agents) Create(agent *v1alpha1.Agent) (result *v1alpha1.Agent, err error) {
	result = &v1alpha1.Agent{}
	err = c.client.Post().
		Resource("agents").
		Body(agent).
		Do().
		Into(result)
	return
}

// Update takes the representation of a agent and updates it. Returns the server's representation of the agent, and an error, if there is any.
func (c *agents) Update(agent *v1alpha1.Agent) (result *v1alpha1.Agent, err error) {
	result = &v1alpha1.Agent{}
	err = c.client.Put().
		Resource("agents").
		Name(agent.Name).
		Body(agent).
		Do().
		Into(result)
	return
}

// Delete takes name of the agent and deletes it. Returns an error if one occurs.
func (c *agents) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("agents").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *agents) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Resource("agents").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched agent.
func (c *agents) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Agent, err error) {
	result = &v1alpha1.Agent{}
	err = c.client.Patch(pt).
		Resource("agents").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
