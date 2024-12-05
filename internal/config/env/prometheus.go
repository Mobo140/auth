package env

import (
	"errors"
	"net"
	"os"

	"github.com/Mobo140/auth/internal/config"
)

var _ config.PrometheusConfig = (*prometheusConfig)(nil)

const (
	prometheusHostEnvName = "PROMETHEUS_HOST"
	prometheusPortEnvName = "PROMETHEUS_PORT"
)

type prometheusConfig struct {
	host string
	port string
}

func NewPrometheusConfig() (*prometheusConfig, error) { //nolint:revive // it's ok
	host := os.Getenv(prometheusHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("prometheus host not found")
	}

	port := os.Getenv(prometheusPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("prometheus port not found")
	}

	return &prometheusConfig{
		host: host,
		port: port,
	}, nil
}

func (c *prometheusConfig) Address() string {
	return net.JoinHostPort(c.host, c.port)
}
