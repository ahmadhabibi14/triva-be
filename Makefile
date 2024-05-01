setup:
	go install github.com/cosmtrek/air@latest
	go install -tags "postgres,mysql" github.com/golang-migrate/migrate/v4/cmd/migrate@latest

migrate:
	migrate create -ext sql -dir database/migration $(state)

migrate-up:
	migrate -path database/migration -database "postgres://habi:habi123@localhost:5432/triva?sslmode=disable" -verbose up

migrate-down:
	migrate -path database/migration -database "postgres://habi:habi123@localhost:5432/triva?sslmode=disable" -verbose down

migrate-fix:
	migrate -path database/migration -database "postgres://habi:habi123@localhost:5432/triva?sslmode=disable" force $(version)

migrate-go:
	go run cmd/database/migrate.go

build:
	go build -o bin/triva cmd/triva/main.go

run:
	go run cmd/triva/main.go