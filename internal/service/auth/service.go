package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/Mobo140/microservices/auth/internal/config"
	"github.com/Mobo140/microservices/auth/internal/model"
	"github.com/Mobo140/microservices/auth/internal/repository"
	"github.com/Mobo140/microservices/auth/internal/service"
	"github.com/Mobo140/microservices/auth/internal/utils"
)

var _ service.AuthService = (*serv)(nil)

type serv struct {
	userDBRepository    repository.UserDBRepository
	userCacheRepository repository.UserCacheRepository
	logRepository       repository.LogRepository
	cfg                 config.SecretConfig
}

func NewService(
	userDBrepository repository.UserDBRepository,
	cacheDBrepository repository.UserCacheRepository,
	logRepository repository.LogRepository,
	cfg config.SecretConfig,
) *serv {
	return &serv{
		userDBRepository:    userDBrepository,
		userCacheRepository: cacheDBrepository,
		logRepository:       logRepository,
		cfg:                 cfg,
	}
}

func (s *serv) Login(ctx context.Context, data *model.LoginData) (*string, error) {
	userData, err := s.userDBRepository.GetHashAndRoleByUsername(ctx, data.Username)
	if err != nil {
		return nil, err
	}

	valid := utils.VerifyPassword(userData.HashedPassword, data.Password)
	if !valid {
		return nil, fmt.Errorf("password is invalid")
	}

	refreshToken, err := utils.GenerateToken(
		&model.Info{
			Username: data.Username,
			Role:     userData.Role,
		},
		s.cfg.RefreshKey(),
		s.cfg.RefreshExpire(),
	)

	if err != nil {
		return nil, errors.New("failed to generate refresh token")
	}

	return &refreshToken, nil
}

func (s *serv) GetRefreshToken(ctx context.Context, refreshTokenStr string) (*string, error) {
	claims, err := utils.VerifyToken(refreshTokenStr, s.cfg.RefreshKey())
	if err != nil {
		return nil, err
	}

	role, err := s.userDBRepository.GetRoleByUsername(ctx, claims.Username)
	if err != nil {
		return nil, err
	}

	refreshToken, err := utils.GenerateToken(
		&model.Info{
			Username: claims.Username,
			Role:     *role,
		},
		s.cfg.RefreshKey(),
		s.cfg.RefreshExpire(),
	)

	if err != nil {
		return nil, err
	}

	return &refreshToken, nil
}

func (s *serv) GetAccessToken(ctx context.Context, refreshTokenStr string) (*string, error) {
	claims, err := utils.VerifyToken(refreshTokenStr, s.cfg.RefreshKey())
	if err != nil {
		return nil, err
	}

	role, err := s.userDBRepository.GetRoleByUsername(ctx, claims.Username)
	if err != nil {
		return nil, err
	}

	accessToken, err := utils.GenerateToken(
		&model.Info{
			Username: claims.Username,
			Role:     *role,
		},
		s.cfg.AccessKey(),
		s.cfg.AccessExpire(),
	)

	if err != nil {
		return nil, err
	}

	return &accessToken, nil
}
