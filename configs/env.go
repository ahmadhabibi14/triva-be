package configs

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv(envFile string) {
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalln("cannot load .env files")
	}
}