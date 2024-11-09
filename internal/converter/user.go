package converter

import (
	"errors"

	"github.com/Mobo140/microservices/auth/internal/model"
	desc "github.com/Mobo140/microservices/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToUserInfoFromService(info *model.UserInfo) (*desc.UserInfo, error) {
	var updatedAtTime *timestamppb.Timestamp
	if info.UpdatedAt.Valid {
		updatedAtTime = timestamppb.New(info.UpdatedAt.Time)
	}

	role, err := mapRoleFromIntToDesc(info.Role)
	if err != nil {
		return nil, err
	}

	return &desc.UserInfo{
		Id:        info.ID,
		Name:      info.Name,
		Email:     info.Email,
		Role:      role,
		CreatedAt: timestamppb.New(info.CreatedAt),
		UpdatedAt: updatedAtTime,
	}, nil
}

func ToUserFromDesc(user *desc.User) (*model.User, error) {
	role, err := mapRoleFromDescToInt(user.Role)
	if err != nil {
		return nil, err
	}

	return &model.User{
		Name:            user.Name,
		Email:           user.Email,
		Password:        user.Password,
		PasswordConfirm: user.PasswordConfirm,
		Role:            role,
	}, nil
}

func ToUpdateUserInfoFromDesc(user *desc.UpdateUserInfo) (*model.UpdateUserInfo, error) {
	if user == nil {
		return nil, errors.New("update user info is empty")
	}

	return &model.UpdateUserInfo{
		Name:  mapNameFromDescToString(user.Name),
		Email: mapEmailFromDescToString(user.Email),
	}, nil
}

func ToGetUsersParamsFromDesc(limit int64, offset int64) *model.GetUsersRequest {
	return &model.GetUsersRequest{
		Limit:  limit,
		Offset: offset,
	}
}

func ToUsersListFromService(users []*model.UserInfo) ([]*desc.UserInfo, error) {
	var err error
	usersList := make([]*desc.UserInfo, len(users))

	for i, user := range users {
		usersList[i], err = ToUserInfoFromService(user)
		if err != nil {
			return nil, err
		}
	}

	return usersList, nil
}
