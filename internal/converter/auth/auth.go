package auth

import (
	"github.com/Mobo140/microservices/auth/internal/model"
	desc "github.com/Mobo140/microservices/auth/pkg/auth_v1"
)

func ToLoginDataFromDesc(req *desc.LoginRequest) *model.LoginData {
	return &model.LoginData{
		Username: req.Name,
		Password: req.Password,
	}
}
