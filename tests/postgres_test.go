package tests

import (
	"testing"
	"triva/configs"
)

func TestConnectPostgres(t *testing.T) {
	pq := configs.ConnectPostgresSQL()

	t.Log(`connected to PostgreSQL`)

	pq.Close()
}