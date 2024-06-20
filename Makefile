setup:
	go install github.com/cosmtrek/air@latest
	go install -tags "postgres,mysql" github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/swaggo/swag/cmd/swag@latest

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

restore:
	docker cp ./database/backups/$(filename) triva-db:/$(filename)
	docker exec -it triva-db psql -U habi -d triva -a -f ./$(filename)

backup:
	docker exec -it triva-db pg_dump -U habi -F t triva > database/backups/db-$(date +%Y-%m-%d).sql

backup-compressed:
	docker exec -it triva-db pg_dump -U habi -F t triva | gzip > database/backups/db-$(date +%Y-%m-%d).tar.gz

swagger:
	swag init -g cmd/triva/main.go --output ./docs --parseDependency --parseInternal

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