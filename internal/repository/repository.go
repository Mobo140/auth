package repository

import (
	"context"

	"github.com/Mobo140/microservices/auth/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) (int64, error)
	Get(ctx context.Context, id int64) (*model.UserInfo, error)
	Update(ctx context.Context, id int64, userInfo *model.UpdateUserInfo) error
	Delete(ctx context.Context, id int64) error
	GetUsers(ctx context.Context, params *model.GetUsersRequest) ([]*model.UserInfo, error)
}

type LogRepository interface {
	Create(ctx context.Context, logEntry *model.LogEntry) error
}
