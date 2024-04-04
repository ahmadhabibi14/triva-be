setup:
	go install -tags "postgres,mysql" github.com/golang-migrate/migrate/v4/cmd/migrate@latest

migrate:
	migrate create -ext sql -dir database/migration $(state)

migrate-up:
	migrate -path database/migration -database "postgres://habi:habi123@localhost:5432/bwizz?sslmode=disable" -verbose up

migrate-down:
	migrate -path database/migration -database "postgres://habi:habi123@localhost:5432/bwizz?sslmode=disable" -verbose down