package tests

import (
	"os"
	"testing"
	"triva/configs"
	"triva/internal/bootstrap/database"
	"triva/internal/bootstrap/logger"

	"github.com/rs/zerolog"
)

var (
	DB *database.Database
	LOG   *zerolog.Logger
)

func TestMain(m *testing.M) {
	configs.LoadEnv()
	logger.InitLogger()

	Pg, err := configs.ConnectPostgresSQL()
	if err != nil {
		panic(err)
	}

	Rd := configs.NewRedisClient()
	_, err = Rd.Ping().Result()
	if err != nil {
		panic(err)
	}

	DB = database.NewDatabase(Pg, Rd)

	os.Exit(m.Run())
}
