package configs

import (
	"database/sql"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"
	rotateLogs "github.com/sutantodadang/go-rotate-logs"
)

func ConnectPostgresSQL() *sqlx.DB {
	postgresDbName := os.Getenv("POSTGRES_DB")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")

	postgresURL := fmt.Sprintf(
		"user=%v password=%v dbname=%v sslmode=disable",
		postgresUser, postgresPassword, postgresDbName,
	)

	driverName := "postgres"
	db, err := sql.Open(driverName, postgresURL)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	zlogLevel, err := strconv.Atoi(os.Getenv(`LOG_LEVEL`))
	if err != nil {
		zlogLevel = int(zerolog.InfoLevel)
	}
	
	var zlog zerolog.Logger

	if os.Getenv(`PROJECT_ENV`) == `prod` {
		logFile := &rotateLogs.RotateLogsWriter{
			Config: rotateLogs.Config{
				Directory:     OS_PATH_SQL_LOG,
				Filename:      "sql.log",
				MaxSize:       10,
				UsingTime:     true,
				FormatTime:    "02-01-2006",
				CleanOldFiles: true,
				MaxAge:        60,
			},
		}
		zlog = zerolog.New(logFile).Level(zerolog.Level(zlogLevel)).With().Logger()
	} else {
		var output io.Writer = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: `2006/01/02 03:04 PM`,
			PartsOrder: []string{
				zerolog.TimestampFieldName,
				zerolog.LevelFieldName,
				zerolog.CallerFieldName,
				zerolog.MessageFieldName,
			},
		}
		zlog = zerolog.New(output).
			Level(zerolog.Level(zlogLevel)).
			With().Timestamp().Logger()
	}

	loggerOptions := []sqldblogger.Option{
		sqldblogger.WithWrapResult(false),
		sqldblogger.WithExecerLevel(sqldblogger.LevelDebug),
		sqldblogger.WithQueryerLevel(sqldblogger.LevelDebug),
		sqldblogger.WithPreparerLevel(sqldblogger.LevelDebug),
		sqldblogger.WithTimeFormat(sqldblogger.TimeFormatRFC3339),
	}
	
	db = sqldblogger.OpenDriver(postgresURL, db.Driver(), zerologadapter.New(zlog), loggerOptions...)

	sqlxDb := sqlx.NewDb(db, driverName)

	return sqlxDb
}