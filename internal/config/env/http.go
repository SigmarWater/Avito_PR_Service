package env

import (
	"errors"
	"github.com/SigmarWater/Avito_PR_Service/internal/config"
	"net"
	"os"
)

var _ config.HTTPConfig = (*httpConfig)(nil)

const (
	httpHostEnvName = "HTTP_HOST"
	httpPortEnvName = "HTTP_PORT"
)

type httpConfig struct {
	host string
	port string
}

func NewHTTPConfig() (*httpConfig, error) {
	host := os.Getenv(httpHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("not found host key")
	}

	port := os.Getenv(httpPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("not found port key")
	}

	return &httpConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg *httpConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
