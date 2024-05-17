package database

import (
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

type Database struct {
	DB *sqlx.DB
	RD *redis.Client
	LOG *zerolog.Logger
}

func NewDatabase(DB *sqlx.DB, RD *redis.Client, LOG *zerolog.Logger) *Database {
	return &Database{DB, RD, LOG}
}