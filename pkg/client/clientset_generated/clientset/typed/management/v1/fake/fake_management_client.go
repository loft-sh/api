// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/loft-sh/api/v3/pkg/client/clientset_generated/clientset/typed/management/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeManagementV1 struct {
	*testing.Fake
}

func (c *FakeManagementV1) Announcements() v1.AnnouncementInterface {
	return &FakeAnnouncements{c}
}

func (c *FakeManagementV1) Apps() v1.AppInterface {
	return &FakeApps{c}
}

func (c *FakeManagementV1) Clusters() v1.ClusterInterface {
	return &FakeClusters{c}
}

func (c *FakeManagementV1) ClusterAccesses() v1.ClusterAccessInterface {
	return &FakeClusterAccesses{c}
}

func (c *FakeManagementV1) ClusterConnects() v1.ClusterConnectInterface {
	return &FakeClusterConnects{c}
}

func (c *FakeManagementV1) ClusterRoleTemplates() v1.ClusterRoleTemplateInterface {
	return &FakeClusterRoleTemplates{c}
}

func (c *FakeManagementV1) Configs() v1.ConfigInterface {
	return &FakeConfigs{c}
}

func (c *FakeManagementV1) DirectClusterEndpointTokens() v1.DirectClusterEndpointTokenInterface {
	return &FakeDirectClusterEndpointTokens{c}
}

func (c *FakeManagementV1) Events() v1.EventInterface {
	return &FakeEvents{c}
}

func (c *FakeManagementV1) Features() v1.FeatureInterface {
	return &FakeFeatures{c}
}

func (c *FakeManagementV1) IngressAuthTokens() v1.IngressAuthTokenInterface {
	return &FakeIngressAuthTokens{c}
}

func (c *FakeManagementV1) Licenses() v1.LicenseInterface {
	return &FakeLicenses{c}
}

func (c *FakeManagementV1) LicenseTokens() v1.LicenseTokenInterface {
	return &FakeLicenseTokens{c}
}

func (c *FakeManagementV1) LoftUpgrades() v1.LoftUpgradeInterface {
	return &FakeLoftUpgrades{c}
}

func (c *FakeManagementV1) OwnedAccessKeys() v1.OwnedAccessKeyInterface {
	return &FakeOwnedAccessKeys{c}
}

func (c *FakeManagementV1) PolicyViolations() v1.PolicyViolationInterface {
	return &FakePolicyViolations{c}
}

func (c *FakeManagementV1) Projects() v1.ProjectInterface {
	return &FakeProjects{c}
}

func (c *FakeManagementV1) ProjectSecrets(namespace string) v1.ProjectSecretInterface {
	return &FakeProjectSecrets{c, namespace}
}

func (c *FakeManagementV1) ResetAccessKeys() v1.ResetAccessKeyInterface {
	return &FakeResetAccessKeys{c}
}

func (c *FakeManagementV1) Selves() v1.SelfInterface {
	return &FakeSelves{c}
}

func (c *FakeManagementV1) SelfSubjectAccessReviews() v1.SelfSubjectAccessReviewInterface {
	return &FakeSelfSubjectAccessReviews{c}
}

func (c *FakeManagementV1) SharedSecrets(namespace string) v1.SharedSecretInterface {
	return &FakeSharedSecrets{c, namespace}
}

func (c *FakeManagementV1) SpaceConstraints() v1.SpaceConstraintInterface {
	return &FakeSpaceConstraints{c}
}

func (c *FakeManagementV1) SpaceInstances(namespace string) v1.SpaceInstanceInterface {
	return &FakeSpaceInstances{c, namespace}
}

func (c *FakeManagementV1) SpaceTemplates() v1.SpaceTemplateInterface {
	return &FakeSpaceTemplates{c}
}

func (c *FakeManagementV1) SubjectAccessReviews() v1.SubjectAccessReviewInterface {
	return &FakeSubjectAccessReviews{c}
}

func (c *FakeManagementV1) Tasks() v1.TaskInterface {
	return &FakeTasks{c}
}

func (c *FakeManagementV1) Teams() v1.TeamInterface {
	return &FakeTeams{c}
}

func (c *FakeManagementV1) Users() v1.UserInterface {
	return &FakeUsers{c}
}

func (c *FakeManagementV1) VirtualClusterInstances(namespace string) v1.VirtualClusterInstanceInterface {
	return &FakeVirtualClusterInstances{c, namespace}
}

func (c *FakeManagementV1) VirtualClusterTemplates() v1.VirtualClusterTemplateInterface {
	return &FakeVirtualClusterTemplates{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeManagementV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
