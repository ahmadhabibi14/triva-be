package tests

import (
	"os"
	"testing"
	"triva/configs"
	"triva/internal/database"

	"github.com/rs/zerolog"
)

var (
	DB *database.Database
	LOG   *zerolog.Logger
)

func TestMain(m *testing.M) {
	configs.LoadEnv()
	LOG = configs.NewLogger()

	Pg, err := configs.ConnectPostgresSQL()
	if err != nil {
		panic(err)
	}

	Rd := configs.NewRedisClient()
	_, err = Rd.Ping().Result()
	if err != nil {
		panic(err)
	}

	DB = database.NewDatabase(Pg, Rd, LOG)

	os.Exit(m.Run())
}
