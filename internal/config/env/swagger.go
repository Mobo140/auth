package env

import (
	"errors"
	"net"
	"os"

	"github.com/Mobo140/auth/internal/config"
)

var _ config.SwaggerConfig = (*swaggerConfig)(nil)

const (
	swaggerHost = "SWAGGER_HOST"
	swaggerPort = "SWAGGER_PORT"
)

type swaggerConfig struct {
	host string
	port string
}

func NewSwaggerConfig() (*swaggerConfig, error) { //nolint:revive // it's ok
	host := os.Getenv(swaggerHost)
	if len(host) == 0 {
		return nil, errors.New("swagger host not found")
	}

	port := os.Getenv(swaggerPort)
	if len(port) == 0 {
		return nil, errors.New("swagger port not found")
	}

	return &swaggerConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg *swaggerConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
