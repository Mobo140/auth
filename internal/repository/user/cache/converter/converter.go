package converter

import (
	"database/sql"
	"time"

	"github.com/Mobo140/microservices/auth/internal/model"
	modelRepo "github.com/Mobo140/microservices/auth/internal/repository/user/cache/model"
)

func ToUserInfoFromRepo(info *modelRepo.UserInfo) *model.UserInfo {
	var updatedAt sql.NullTime
	if info.UpdatedAtNs != nil {
		updatedAt = sql.NullTime{
			Time:  time.Unix(0, *info.UpdatedAtNs),
			Valid: true,
		}
	}

	return &model.UserInfo{
		ID:        info.ID,
		Name:      info.Name,
		Email:     info.Email,
		Role:      info.Role,
		CreatedAt: time.Unix(0, info.CreatedAtNs),
		UpdatedAt: updatedAt,
	}

}

func ToUserFromService(user *model.User) *modelRepo.User {
	return &modelRepo.User{
		Name:           user.Name,
		Email:          user.Email,
		Role:           user.Role,
		HashedPassword: user.HashedPassword,
	}
}
