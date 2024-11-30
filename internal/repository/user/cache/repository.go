package cache

import (
	"context"

	"github.com/Mobo140/microservices/auth/internal/client/cache"
	"github.com/Mobo140/microservices/auth/internal/model"
	"github.com/Mobo140/microservices/auth/internal/repository"
	"github.com/Mobo140/microservices/auth/internal/repository/user/cache/converter"
	modelRepo "github.com/Mobo140/microservices/auth/internal/repository/user/cache/model"
	redigo "github.com/gomodule/redigo/redis"
)

var _ repository.UserCacheRepository = (*userCache)(nil)

type userCache struct {
	cache cache.Client
}

func NewRepository(cache cache.Client) *userCache { //nolint:revive // it's ok
	return &userCache{cache: cache}
}

func (r *userCache) GetHashAndRoleByUsername(ctx context.Context, username string) (*model.UserAuthData, error) {
	values, err := r.cache.Cache().HGetAll(ctx, username)
	if err != nil {
		return nil, err
	}

	if len(values) == 0 {
		return nil, model.ErrUserNotFound
	}

	var data modelRepo.UserAuthData

	err = redigo.ScanStruct(values, &data)
	if err != nil {
		return nil, err
	}

	return converter.ToUserAuthDataFromRepo(&data), nil
}

func (r *userCache) SetHashAndRole(ctx context.Context, username string, data *model.UserAuthData) error {
	err := r.cache.Cache().HashSet(ctx,
		username,
		modelRepo.UserAuthData{
			HashedPassword: data.HashedPassword,
			Role:           data.Role,
		})
	if err != nil {
		return err
	}

	return nil
}
