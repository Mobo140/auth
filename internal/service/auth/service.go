package auth

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/Mobo140/auth/internal/config"
	"github.com/Mobo140/auth/internal/model"
	"github.com/Mobo140/auth/internal/repository"
	"github.com/Mobo140/auth/internal/service"
	"github.com/Mobo140/auth/internal/utils"
	"github.com/Mobo140/platform_common/pkg/db"
)

var _ service.AuthService = (*serv)(nil)

type serv struct {
	userDBRepository    repository.UserDBRepository
	userCacheRepository repository.UserCacheRepository
	logRepository       repository.LogRepository
	txManager           db.TxManager
	cfg                 config.SecretConfig
}

func NewService(
	userDBrepository repository.UserDBRepository,
	cacheDBrepository repository.UserCacheRepository,
	logRepository repository.LogRepository,
	txManager db.TxManager,
	cfg config.SecretConfig,
) *serv { //nolint:revive //it's ok
	return &serv{
		userDBRepository:    userDBrepository,
		userCacheRepository: cacheDBrepository,
		logRepository:       logRepository,
		txManager:           txManager,
		cfg:                 cfg,
	}
}

func (s *serv) Login(ctx context.Context, data *model.LoginData) (*string, error) {
	var userData *model.UserAuthData
	var refreshToken string

	err := s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
		var errTx error
		var errCache error

		// check data in cache
		userData, errCache = s.userCacheRepository.GetHashAndRoleByUsername(ctx, data.Username)
		if errCache == nil && userData != nil {
			return nil
		}

		// no data in cache. Check in db
		userData, errTx = s.userDBRepository.GetHashAndRoleByUsername(ctx, data.Username)
		if errTx != nil {
			return errTx
		}

		// write data to cache
		errCache = s.userCacheRepository.SetHashAndRole(ctx, data.Username, userData)
		if errCache != nil {
			log.Printf("failed to set user data in cache: %v", errCache)
		}

		valid := utils.VerifyPassword(userData.HashedPassword, data.Password)
		if !valid {
			return fmt.Errorf("password is invalid")
		}

		logEntry := model.LogEntryAuth{
			Username: data.Username,
			Action:   fmt.Sprintf("Login user: Name=%s, Role=%d", data.Username, userData.Role),
		}

		errTx = s.logRepository.CreateLogAuth(ctx, &logEntry)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	refreshToken, err = utils.GenerateToken(
		&model.Info{
			Username: data.Username,
			Role:     strconv.FormatInt(userData.Role, 10),
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
	var claims *model.UserClaims
	var refreshToken string

	err := s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
		var errTx error

		claims, errTx = utils.VerifyToken(refreshTokenStr, s.cfg.RefreshKey())
		if errTx != nil {
			return errors.New("failed to verify token")
		}

		logEntry := model.LogEntryAuth{
			Username: claims.Username,
			Action:   fmt.Sprintf("Get refresh token: Name=%s Role=%s", claims.Username, claims.Role),
		}

		errTx = s.logRepository.CreateLogAuth(ctx, &logEntry)
		if errTx != nil {
			return fmt.Errorf("failed to create log entry, Role: %q", claims.Role)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	refreshToken, err = utils.GenerateToken(
		&model.Info{
			Username: claims.Username,
			Role:     claims.Role,
		},
		s.cfg.RefreshKey(),
		s.cfg.RefreshExpire(),
	)

	if err != nil {
		return nil, errors.New("failed to generate refresh token")
	}

	return &refreshToken, nil
}

func (s *serv) GetAccessToken(ctx context.Context, refreshTokenStr string) (*string, error) {
	var claims *model.UserClaims
	var accessToken string

	err := s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
		var errTx error

		claims, errTx = utils.VerifyToken(refreshTokenStr, s.cfg.RefreshKey())
		if errTx != nil {
			return errTx
		}

		logEntry := model.LogEntryAuth{
			Username: claims.Username,
			Action:   fmt.Sprintf("Get access token: Name=%s, Role=%s", claims.Username, claims.Role),
		}

		errTx = s.logRepository.CreateLogAuth(ctx, &logEntry)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	accessToken, err = utils.GenerateToken(
		&model.Info{
			Username: claims.Username,
			Role:     claims.Role,
		},
		s.cfg.AccessKey(),
		s.cfg.AccessExpire(),
	)

	if err != nil {
		return nil, errors.New("failed to generate access token")
	}

	return &accessToken, nil
}
