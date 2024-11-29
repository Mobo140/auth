package repository

import (
	"context"

	"github.com/Mobo140/microservices/auth/internal/model"
)

type UserDBRepository interface {
	Create(ctx context.Context, user *model.User) (int64, error)
	Get(ctx context.Context, id int64) (*model.UserInfo, error)
	Update(ctx context.Context, id int64, userInfo *model.UpdateUserInfo) error
	Delete(ctx context.Context, id int64) error
	GetUsers(ctx context.Context, params *model.GetUsersRequest) ([]*model.UserInfo, error)
	GetHashAndRoleByUsername(ctx context.Context, username string) (*model.UserAuthData, error)
	GetRoleByUsername(ctx context.Context, username string) (int64, error)
}

type UserCacheRepository interface {
	GetHashAndRoleByUsername(ctx context.Context, username string) (*model.UserAuthData, error)
	SetHashAndRole(ctx context.Context, username string, data *model.UserAuthData) error
}

type LogRepository interface {
	CreateLogUser(ctx context.Context, logEntry *model.LogEntryUser) error
	CreateLogAuth(ctx context.Context, logEntry *model.LogEntryAuth) error
}

type AccessDBRepository interface {
	GetEndpointsAccess(ctx context.Context) ([]*model.AccessEndpoint, error)
}

type AccessCacheRepository interface {
	GetEndpoints(ctx context.Context) ([]*model.AccessEndpoint, error)
	SetEndpoints(ctx context.Context, endpoints []*model.AccessEndpoint) error
}
