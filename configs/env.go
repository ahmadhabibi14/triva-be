package configs

import (
	"github.com/joho/godotenv"
)

func LoadEnv() {
	dirRetryList := []string{``, `../`, `../../`, `../../../`}
	var err error
	for _, dirPrefix := range dirRetryList {
		envFile := dirPrefix + `.env`
		err = godotenv.Overload(envFile)
		if err == nil {
			return
		}
	}
	panic("cannot load .env file")
}
