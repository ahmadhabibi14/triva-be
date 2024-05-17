package logger

import (
	"io"
	"os"
	"strconv"
	"triva/configs"

	"github.com/rs/zerolog"
)

var Log zerolog.Logger

func InitLogger() {
	logLevel, err := strconv.Atoi(os.Getenv(`LOG_LEVEL`))
	if err != nil {
		logLevel = int(zerolog.InfoLevel)
	}

	var logOutput io.Writer

	if os.Getenv(`WEB_ENV`) == `prod` {
		file, _ := os.OpenFile(configs.PATH_APPLICATION_LOG, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		logOutput = file
	} else {
		var output io.Writer = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: `03:04 PM`,
			PartsOrder: []string{
				zerolog.TimestampFieldName,
				zerolog.LevelFieldName,
				zerolog.CallerFieldName,
				zerolog.MessageFieldName,
			},
		}
		logOutput = output
	}

	l := zerolog.New(logOutput).
		Level(zerolog.Level(logLevel)).
		With().
		Timestamp().
		Caller().
		Logger()

	Log = l
}
