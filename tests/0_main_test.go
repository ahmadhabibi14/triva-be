package tests

import (
	"testing"
	"triva/configs"
)

func TestMain(m *testing.M) {
	configs.LoadEnv()
}