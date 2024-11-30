package env

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/Mobo140/microservices/auth/internal/config"
)

var _ config.SecretConfig = (*secretConfig)(nil)

const (
	secretRefreshPath = "REFRESH_SECRET_PATH"
	secretAccessPath  = "ACCESS_SECRET_PATH"
	expireRefresh     = "REFRESH_TOKEN_EXPIRE"
	expireAccess      = "ACCESS_TOKEN_EXPIRE"
)

type secretConfig struct {
	refreshKey    []byte
	refreshExpire time.Duration
	accessKey     []byte
	accessExpire  time.Duration
}

func NewSecretConfig() (*secretConfig, error) { //nolint:revive // it's ok
	refreshPath := os.Getenv(secretRefreshPath)
	if refreshPath == "" {
		return nil, errors.New("refresh secret path is not set")
	}

	refreshKey, err := os.ReadFile(refreshPath)
	if err != nil {
		return nil, errors.New("failed to read refresh secret file")
	}

	accessPath := os.Getenv(secretAccessPath)
	if accessPath == "" {
		return nil, errors.New("access secret path is not set")
	}

	accessKey, err := os.ReadFile(accessPath)
	if err != nil {
		return nil, errors.New("failed to read access secret file")
	}

	refreshExpireStr := os.Getenv(expireRefresh)

	if refreshPath == "" {
		return nil, errors.New("refresh expire is not set")
	}

	refreshExpireInt, err := strconv.Atoi(refreshExpireStr)
	if err != nil {
		return nil, errors.New("invalid refresh expire value")
	}

	accessExpireStr := os.Getenv(expireAccess)

	if accessPath == "" {
		return nil, errors.New("access expire is not set")
	}

	accessExpireInt, err := strconv.Atoi(accessExpireStr)
	if err != nil {
		return nil, errors.New("invalid access expire value")
	}

	return &secretConfig{
		refreshKey:    refreshKey,
		accessKey:     accessKey,
		refreshExpire: time.Duration(refreshExpireInt) * time.Minute,
		accessExpire:  time.Duration(accessExpireInt) * time.Minute,
	}, nil
}

func (s *secretConfig) RefreshKey() []byte {
	return s.refreshKey
}

func (s *secretConfig) RefreshExpire() time.Duration {
	return s.refreshExpire
}

func (s *secretConfig) AccessKey() []byte {
	return s.accessKey
}

func (s *secretConfig) AccessExpire() time.Duration {
	return s.accessExpire
}
