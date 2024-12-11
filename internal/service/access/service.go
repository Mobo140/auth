package access

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/Mobo140/auth/internal/config"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"

	"github.com/Mobo140/auth/internal/model"
	"github.com/Mobo140/auth/internal/repository"
	"github.com/Mobo140/auth/internal/service"
	"github.com/Mobo140/auth/internal/utils"
	"github.com/Mobo140/platform_common/pkg/db"
	"github.com/Mobo140/platform_common/pkg/logger"
)

var (
	_               service.AccessService = (*serv)(nil)
	accessibleRoles map[string]string
)

type serv struct {
	accessDBRepository    repository.AccessDBRepository
	accessCacheRepository repository.AccessCacheRepository
	logRepository         repository.LogRepository
	txManager             db.TxManager
	cfg                   config.SecretConfig
}

func NewService(
	accessDBRepository repository.AccessDBRepository,
	accessCacheRepository repository.AccessCacheRepository,
	logRepository repository.LogRepository,
	txManager db.TxManager,
	cfg config.SecretConfig,
) *serv { //nolint:revive //it's ok
	return &serv{
		accessDBRepository:    accessDBRepository,
		accessCacheRepository: accessCacheRepository,
		logRepository:         logRepository,
		txManager:             txManager,
		cfg:                   cfg,
	}
}

func (s *serv) Check(ctx context.Context, accessToken string, endpoint string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Check/Service")
	defer span.Finish()

	logger.Info("Verify token...")

	claims, err := utils.VerifyToken(accessToken, s.cfg.AccessKey())
	if err != nil {
		err = fmt.Errorf("access token is invalid: %w", err)
		logger.Error("Failed to verify token: %w", zap.Error(err))

		return err
	}

	logger.Info("Token verifyed")

	accessibleMap, err := s.accessibleRoles(ctx, claims.Username)
	if err != nil {
		logger.Error("Failed to get accessible roles: %w", zap.Error(err))

		return errors.New("failed to get accessible roles")
	}

	role, ok := accessibleMap[endpoint]
	if !ok {
		return nil
	}

	if role == claims.Role {
		return nil
	}

	return errors.New("access denied")
}

func (s *serv) accessibleRoles(ctx context.Context, username string) (map[string]string, error) {
	if accessibleRoles == nil {
		var accessData []*model.AccessEndpoint
		accessibleRoles = make(map[string]string)

		err := s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
			var errTx error
			var errCache error

			// check data in cache
			accessData, errCache = s.accessCacheRepository.GetEndpoints(ctx)
			if errCache == nil && accessData != nil {
				return nil
			}

			// no data in cache. Check in db
			accessData, errTx = s.accessDBRepository.GetEndpointsAccess(ctx)
			if errTx != nil {
				return errTx
			}

			// write data to cache
			errCache = s.accessCacheRepository.SetEndpoints(ctx, accessData)
			if errCache != nil {
				log.Printf("failed to set endpoints in cache: %v", errCache)
			}

			logEntry := model.LogEntryAuth{
				Username: username,
				Action:   fmt.Sprintf("Get endpoints for access: Name=%s", username),
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

		for _, access := range accessData {
			accessibleRoles[access.Endpoint] = strconv.FormatInt(access.Role, 10)
		}
	}

	return accessibleRoles, nil
}
