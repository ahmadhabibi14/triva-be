package tests

import (
	"os"
	"testing"
	"triva/configs"

	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

var (
	PG_DB *sqlx.DB
	RD_DB *redis.Client
	LOG   *zerolog.Logger
)

func TestMain(m *testing.M) {
	configs.LoadEnv()

	var err error

	PG_DB, err = configs.ConnectPostgresSQL()
	if err != nil {
		panic(err)
	}

	RD_DB = configs.NewRedisClient()
	_, err = RD_DB.Ping().Result()
	if err != nil {
		panic(err)
	}

	LOG = configs.NewLogger()

	os.Exit(m.Run())
}
