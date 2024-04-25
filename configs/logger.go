package configs

import (
	"io"
	"os"
	"strconv"

	"github.com/rs/zerolog"
)

func NewLogger() *zerolog.Logger {
	logLevel, err := strconv.Atoi(os.Getenv(`LOG_LEVEL`))
	if err != nil {
		logLevel = int(zerolog.InfoLevel)
	}

	var logOutput io.Writer
	
	if os.Getenv(`WEB_ENV`) == `prod` {
		file, _ := os.OpenFile(`logs/application.log`, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
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

	return &l
}
