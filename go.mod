module github.com/loft-sh/api/v2

go 1.13

require (
	github.com/ghodss/yaml v1.0.0
	github.com/loft-sh/agentapi/v2 v2.1.5-0.20220429080136-946c36ec7229
	github.com/loft-sh/apiserver v0.0.0-20211216225656-cafe09adc0b8
	github.com/loft-sh/jspolicy v0.1.0
	k8s.io/api v0.23.5
	k8s.io/apimachinery v0.23.5
	k8s.io/apiserver v0.23.5
	k8s.io/client-go v0.23.5
	k8s.io/kube-openapi v0.0.0-20211115234752-e816edb12b65
	sigs.k8s.io/controller-runtime v0.11.2
	sigs.k8s.io/controller-tools v0.7.0
)

replace github.com/loft-sh/agentapi/v2 => ../../agentapi/v2
