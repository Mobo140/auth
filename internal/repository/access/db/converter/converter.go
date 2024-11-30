package converter

import (
	"github.com/Mobo140/auth/internal/model"
	modelRepo "github.com/Mobo140/auth/internal/repository/access/db/model"
)

func ToEndpointsAccessFromRepo(endpoints []*modelRepo.AccessEndpoint) []*model.AccessEndpoint {
	var endpointsList = make([]*model.AccessEndpoint, len(endpoints))
	for i, endpoint := range endpoints {
		endpointsList[i] = ToEndpointAccessFromRepo(endpoint)
	}

	return endpointsList
}

func ToEndpointAccessFromRepo(endpoint *modelRepo.AccessEndpoint) *model.AccessEndpoint {
	return &model.AccessEndpoint{
		Endpoint: endpoint.Endpoint,
		Role:     endpoint.Role,
	}
}
