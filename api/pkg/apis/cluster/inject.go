package cluster

import (
	"github.com/loft-sh/api/pkg/managerfactory"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var Client client.Client

var Factory managerfactory.SharedManagerFactory
