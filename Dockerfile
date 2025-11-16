FROM golang:1.25-alpine3.22  AS builder
WORKDIR /app

# Копируем go.mod и go.sum для auth
COPY go.mod go.sum /app/
# Скачиваем зависимости (platform уже доступен через replace)
RUN go mod download
# Копируем исходный код auth
COPY . /app/
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o pull_request_service ./cmd/server/main.go

FROM alpine:3.22.1
WORKDIR /app
COPY --from=builder /app/pull_request_service /app/
COPY .env /app/.env
EXPOSE 8081 8081
CMD ["./pull_request_service"]
