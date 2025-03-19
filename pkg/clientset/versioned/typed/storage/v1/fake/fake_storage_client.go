// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/loft-sh/api/v4/pkg/clientset/versioned/typed/storage/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeStorageV1 struct {
	*testing.Fake
}

func (c *FakeStorageV1) AccessKeys() v1.AccessKeyInterface {
	return newFakeAccessKeys(c)
}

func (c *FakeStorageV1) Apps() v1.AppInterface {
	return newFakeApps(c)
}

func (c *FakeStorageV1) Clusters() v1.ClusterInterface {
	return newFakeClusters(c)
}

func (c *FakeStorageV1) ClusterAccesses() v1.ClusterAccessInterface {
	return newFakeClusterAccesses(c)
}

func (c *FakeStorageV1) ClusterRoleTemplates() v1.ClusterRoleTemplateInterface {
	return newFakeClusterRoleTemplates(c)
}

func (c *FakeStorageV1) DevPodEnvironmentTemplates() v1.DevPodEnvironmentTemplateInterface {
	return newFakeDevPodEnvironmentTemplates(c)
}

func (c *FakeStorageV1) DevPodWorkspaceInstances(namespace string) v1.DevPodWorkspaceInstanceInterface {
	return newFakeDevPodWorkspaceInstances(c, namespace)
}

func (c *FakeStorageV1) DevPodWorkspacePresets() v1.DevPodWorkspacePresetInterface {
	return newFakeDevPodWorkspacePresets(c)
}

func (c *FakeStorageV1) DevPodWorkspaceTemplates() v1.DevPodWorkspaceTemplateInterface {
	return newFakeDevPodWorkspaceTemplates(c)
}

func (c *FakeStorageV1) NetworkPeers() v1.NetworkPeerInterface {
	return newFakeNetworkPeers(c)
}

func (c *FakeStorageV1) Projects() v1.ProjectInterface {
	return newFakeProjects(c)
}

func (c *FakeStorageV1) SharedSecrets(namespace string) v1.SharedSecretInterface {
	return newFakeSharedSecrets(c, namespace)
}

func (c *FakeStorageV1) SpaceInstances(namespace string) v1.SpaceInstanceInterface {
	return newFakeSpaceInstances(c, namespace)
}

func (c *FakeStorageV1) SpaceTemplates() v1.SpaceTemplateInterface {
	return newFakeSpaceTemplates(c)
}

func (c *FakeStorageV1) Tasks() v1.TaskInterface {
	return newFakeTasks(c)
}

func (c *FakeStorageV1) Teams() v1.TeamInterface {
	return newFakeTeams(c)
}

func (c *FakeStorageV1) Users() v1.UserInterface {
	return newFakeUsers(c)
}

func (c *FakeStorageV1) VirtualClusterInstances(namespace string) v1.VirtualClusterInstanceInterface {
	return newFakeVirtualClusterInstances(c, namespace)
}

func (c *FakeStorageV1) VirtualClusterTemplates() v1.VirtualClusterTemplateInterface {
	return newFakeVirtualClusterTemplates(c)
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeStorageV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
