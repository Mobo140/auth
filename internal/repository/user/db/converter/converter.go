package converter

import (
	model "github.com/Mobo140/microservices/auth/internal/model"
	modelRepo "github.com/Mobo140/microservices/auth/internal/repository/user/db/model"
)

func ToUserFromRepo(user *modelRepo.User) *model.User {
	return &model.User{
		Name:           user.Name,
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
		Role:           user.Role,
	}
}

func ToUserInfoFromRepo(info *modelRepo.UserInfo) *model.UserInfo {
	return &model.UserInfo{
		ID:        info.ID,
		Name:      info.Name,
		Email:     info.Email,
		Role:      info.Role,
		CreatedAt: info.CreatedAt,
		UpdatedAt: info.UpdatedAt,
	}
}

func ToUsersInfoFromRepo(users []*modelRepo.UserInfo) []*model.UserInfo {
	var usersList = make([]*model.UserInfo, len(users))
	for i, user := range users {
		usersList[i] = ToUserInfoFromRepo(user)
	}

	return usersList
}

func ToUserAuthDataFromRepo(user *modelRepo.UserAuthData) *model.UserAuthData {
	var hashedPassword string
	if user.HashedPassword.Valid {
		hashedPassword = user.HashedPassword.String
	}

	return &model.UserAuthData{
		HashedPassword: hashedPassword,
		Role:           user.Role,
	}
}
