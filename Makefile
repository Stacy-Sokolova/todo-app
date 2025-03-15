include .env

run:
	go run ./cmd/main.go

migrate:
	migrate -path ./schema -database 'postgres://${POSTGRES_USER}:${DB_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=${SSLMODE}' up