package cache

import (
	"context"
	"strconv"

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

func NewRepository(cache cache.Client) *userCache {
	return &userCache{cache: cache}
}

func (r *userCache) Get(ctx context.Context, id int64) (*model.UserInfo, error) {
	idStr := strconv.FormatInt(id, 10)
	values, err := r.cache.Cache().HGetAll(ctx, idStr)
	if err != nil {
		return nil, err
	}

	if len(values) == 0 {
		return nil, model.ErrorUserNotFound
	}

	var info modelRepo.UserInfo

	err = redigo.ScanStruct(values, &info)
	if err != nil {
		return nil, err
	}

	return converter.ToUserInfoFromRepo(&info), nil
}

func (r *userCache) Create(ctx context.Context, id int64, user *model.User) (int64, error) {
	idStr := strconv.FormatInt(id, 10)
	userRepo := converter.ToUserFromService(user)

	err := r.cache.Cache().HashSet(ctx, idStr, userRepo)
	if err != nil {
		return 0, err
	}

	return id, nil
}
