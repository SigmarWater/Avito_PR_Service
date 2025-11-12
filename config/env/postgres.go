package env

import (
	"errors"
	"os"
)

const (
	postgresDSNEnvName = "POSTGRES_DSN"
)

type RepositoryConfig struct {
	dsn string
}

func NewRepositoryConfig() (*RepositoryConfig, error) {
	dsn := os.Getenv(postgresDSNEnvName)
	if len(dsn) == 0 {
		return nil, errors.New("not exist dsn")
	}
	return &RepositoryConfig{
		dsn: dsn,
	}, nil
}

func (c *RepositoryConfig) DSN() string {
	return c.dsn
}
