package tests

import (
	"os"
	"testing"
	"triva/configs"
)

func TestMain(m *testing.M) {
	configs.LoadEnv()

	os.Exit(m.Run()) 
}