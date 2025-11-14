package config

import (
	"github.com/joho/godotenv"
	"log"
)

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		log.Printf("failed to load .evn file: %v\n", err.Error())
		return err
	}
	return nil
}

type PostgresConfig interface {
	DSN() string
}
