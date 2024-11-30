package cache

import (
	"context"

	"github.com/Mobo140/microservices/auth/internal/client/cache"
	"github.com/Mobo140/microservices/auth/internal/model"
	"github.com/Mobo140/microservices/auth/internal/repository"
	"github.com/Mobo140/microservices/auth/internal/repository/access/cache/converter"
	modelCache "github.com/Mobo140/microservices/auth/internal/repository/access/cache/model"
	redigo "github.com/gomodule/redigo/redis"
)

var _ repository.AccessCacheRepository = (*accessCache)(nil)

const (
	keyEndpoints = "endpoints"
)

type accessCache struct {
	cache cache.Client
}

func NewRepository(cache cache.Client) *accessCache { //nolint:revive //it's ok
	return &accessCache{cache: cache}
}

func (r *accessCache) GetEndpoints(ctx context.Context) ([]*model.AccessEndpoint, error) {
	endpoints, err := r.cache.Cache().HGetAll(ctx, keyEndpoints)
	if err != nil {
		return nil, err
	}

	if len(endpoints) == 0 {
		return nil, model.ErrEndpointsNotFound
	}

	var accessEndpoints []*modelCache.AccessEndpoint

	err = redigo.ScanStruct(endpoints, &accessEndpoints)
	if err != nil {
		return nil, err
	}

	return converter.ToEndpointsAccessFromRepo(accessEndpoints), nil
}

func (r *accessCache) SetEndpoints(ctx context.Context, endpoints []*model.AccessEndpoint) error {
	data := converter.ToEndpointsAccessFromService(endpoints)

	err := r.cache.Cache().HashSet(ctx,
		keyEndpoints,
		data)
	if err != nil {
		return err
	}

	return nil
}
