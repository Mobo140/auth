package service

import (
	"context"

	"github.com/Mobo140/microservices/auth/internal/model"
)

type UserService interface {
	Create(ctx context.Context, user *model.User) (int64, error)
	Get(ctx context.Context, id int64) (*model.UserInfo, error)
	Update(ctx context.Context, id int64, userInfo *model.UpdateUserInfo) error
	Delete(ctx context.Context, id int64) error
	GetUsers(ctx context.Context, params *model.GetUsersRequest) ([]*model.UserInfo, error)
}

type AuthService interface {
	Login(ctx context.Context, data *model.LoginData) (*string, error)
	GetRefreshToken(ctx context.Context, refreshToken string) (*string, error)
	GetAccessToken(ctx context.Context, refreshToken string) (*string, error)
}
