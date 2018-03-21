package fake

import (
	v1alpha1 "github.com/marjoram/pipeline/client/k8s/clientset/versioned/typed/agent.cncd.io/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeAgentV1alpha1 struct {
	*testing.Fake
}

func (c *FakeAgentV1alpha1) Agents() v1alpha1.AgentInterface {
	return &FakeAgents{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeAgentV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
