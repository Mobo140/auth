package converter

import (
	model "github.com/Mobo140/microservices/auth/internal/model"
	modelRepo "github.com/Mobo140/microservices/auth/internal/repository/user/model"
)

func ToUserFromRepo(user *modelRepo.User) *model.User {
	return &model.User{
		Name:            user.Name,
		Email:           user.Email,
		Password:        user.Password,
		PasswordConfirm: user.PasswordConfirm,
		Role:            user.Role,
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
