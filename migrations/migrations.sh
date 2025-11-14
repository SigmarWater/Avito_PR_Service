#!/bin/bash
source .env

sleep 2 && goose -dir "${MIGRATIONS_DIR}" postgres "${POSTGRES_DSN}" up -v