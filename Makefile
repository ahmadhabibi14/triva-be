setup:
	go install -tags "postgres,mysql" github.com/golang-migrate/migrate/v4/cmd/migrate@latest

migrate:
	migrate create -ext sql -dir database/migration $(state)

migrate-up:
	migrate -path database/migration -database "postgres://habi:habi123@localhost:5432/triva?sslmode=disable" -verbose up

migrate-down:
	migrate -path database/migration -database "postgres://habi:habi123@localhost:5432/triva?sslmode=disable" -verbose down

build:
	go build -o bin/triva cmd/triva.go

run:
	go run cmd/triva.go