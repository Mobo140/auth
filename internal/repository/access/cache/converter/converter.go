package converter

import (
	"github.com/Mobo140/auth/internal/model"
	modelCache "github.com/Mobo140/auth/internal/repository/access/cache/model"
)

func ToEndpointsAccessFromRepo(endpoints []*modelCache.AccessEndpoint) []*model.AccessEndpoint {
	var endpointsList = make([]*model.AccessEndpoint, len(endpoints))
	for i, endpoint := range endpoints {
		endpointsList[i] = ToEndpointAccessFromRepo(endpoint)
	}

	return endpointsList
}

func ToEndpointAccessFromRepo(endpoint *modelCache.AccessEndpoint) *model.AccessEndpoint {
	return &model.AccessEndpoint{
		Endpoint: endpoint.Endpoint,
		Role:     endpoint.Role,
	}
}

func ToEndpointsAccessFromService(endpoints []*model.AccessEndpoint) []*modelCache.AccessEndpoint {
	var endpointsList = make([]*modelCache.AccessEndpoint, len(endpoints))
	for i, endpoint := range endpoints {
		endpointsList[i] = ToEndpointAccessFromService(endpoint)
	}

	return endpointsList
}

func ToEndpointAccessFromService(endpoint *model.AccessEndpoint) *modelCache.AccessEndpoint {
	return &modelCache.AccessEndpoint{
		Endpoint: endpoint.Endpoint,
		Role:     endpoint.Role,
	}
}
