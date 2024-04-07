package configs

import (
	"github.com/joho/godotenv"
)

func LoadEnv(envFile string) error {
	err := godotenv.Load(envFile)
	return err
}