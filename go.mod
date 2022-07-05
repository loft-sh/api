module github.com/loft-sh/api/v2

go 1.13

require (
	github.com/ghodss/yaml v1.0.0
	github.com/loft-sh/agentapi/v2 v2.2.1
	github.com/loft-sh/apiserver v0.0.0-20220507140345-294e3e3117e3
	github.com/loft-sh/jspolicy v0.1.0
	k8s.io/api v0.24.0
	k8s.io/apimachinery v0.24.0
	k8s.io/apiserver v0.24.0
	k8s.io/client-go v0.24.0
	k8s.io/kube-openapi v0.0.0-20220328201542-3ee0da9b0b42
	sigs.k8s.io/controller-runtime v0.11.2
	sigs.k8s.io/controller-tools v0.8.0
)

replace github.com/loft-sh/agentapi/v2 => ../../agentapi/v2
