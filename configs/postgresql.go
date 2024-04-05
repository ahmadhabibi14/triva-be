package configs

import (
	"fmt"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var PostgresDB *sqlx.DB

func ConnectPostgresSQL() error {
	postgresDbName := os.Getenv("POSTGRES_DB")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresInfo := fmt.Sprintf(
		"user=%v password=%v dbname=%v sslmode=disable",
		postgresUser, postgresPassword, postgresDbName,
	)

	db, err := sqlx.Connect("postgres", postgresInfo)
	if err != nil {
		return err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	PostgresDB = db

	return nil
}