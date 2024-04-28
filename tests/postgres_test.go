package tests

import (
	"testing"
	"triva/configs"

	"github.com/stretchr/testify/assert"
)

func TestConnectPostgres(t *testing.T) {
	_, err := configs.ConnectPostgresSQL()

	assert.Nil(t, err, `failed to connect PostgreSQL`)

	t.Log(`connected to PostgreSQL`)
}