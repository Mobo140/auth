package config

import (
	"time"

	"github.com/joho/godotenv"
)

type GRPCConfig interface {
	Address() string
}

type HTTPConfig interface {
	Address() string
}

type PGConfig interface {
	DSN() string
}

type RedisConfig interface {
	Address() string
	ConnectionTimeout() time.Duration
	MaxIdle() int
	IdleTimeout() time.Duration
}

type SecretConfig interface {
	RefreshKey() []byte
	RefreshExpire() time.Duration
	AccessKey() []byte
	AccessExpire() time.Duration
}
type StorageConfig interface {
	Mode() string
}

type SwaggerConfig interface {
	Address() string
}

type PrometheusConfig interface {
	Address() string
}

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}
