package converter

import (
	"github.com/Mobo140/microservices/auth/internal/model"
	modelCache "github.com/Mobo140/microservices/auth/internal/repository/user/cache/model"
)

func ToUserAuthDataFromRepo(data *modelCache.UserAuthData) *model.UserAuthData {
	return &model.UserAuthData{
		HashedPassword: data.HashedPassword,
		Role:           data.Role,
	}
}

func ToUserAuthDataFromService(data *model.UserAuthData) *modelCache.UserAuthData {
	return &modelCache.UserAuthData{
		HashedPassword: data.HashedPassword,
		Role:           data.Role,
	}
}
