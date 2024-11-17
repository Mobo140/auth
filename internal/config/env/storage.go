package env

import (
	"os"

	"github.com/Mobo140/microservices/auth/internal/config"
)

var _ config.StorageConfig = (*storageConfig)(nil)

const storageModeEnvName = "STORAGE_MODE"

type storageConfig struct {
	mode string
}

func NewStorageConfig() (*storageConfig, error) {
	storageMode := os.Getenv(storageModeEnvName)
	if len(storageMode) == 0 {
		return nil, nil
		//return nil, errors.New("storage mode not found")
	}

	return &storageConfig{
		mode: storageMode,
	}, nil
}

func (cfg *storageConfig) Mode() string {
	return cfg.mode
}