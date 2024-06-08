setup:
	go install github.com/cosmtrek/air@latest
	go install -tags "postgres,mysql" github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/swaggo/swag/cmd/swag@latest

migrate:
	migrate create -ext sql -dir databases/migration $(state)

migrate-up:
	migrate -path databases/migration -database "postgres://habi:habi123@localhost:5432/triva?sslmode=disable" -verbose up

migrate-down:
	migrate -path databases/migration -database "postgres://habi:habi123@localhost:5432/triva?sslmode=disable" -verbose down

migrate-fix:
	migrate -path databases/migration -database "postgres://habi:habi123@localhost:5432/triva?sslmode=disable" force $(version)

migrate-go:
	go run cmd/database/migrate.go

swagger:
	swag init -g cmd/triva/main.go --output ./docs

build:
	go build -o bin/triva cmd/triva/main.go

run:
	go run cmd/triva/main.go

docker-dev-up:
	docker-compose -f docker-compose.dev.yml up -d

docker-dev-down:
	docker-compose -f docker-compose.dev.yml down

docker-prod-up:
	docker-compose -f docker-compose.prod.yml up -d

docker-prod-down:
	docker-compose -f docker-compose.prod.yml down