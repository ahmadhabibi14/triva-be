package tests

import (
	"testing"
	"triva/configs"

	"github.com/stretchr/testify/assert"
)

func TestConnectRedis(t *testing.T) {
	rd := configs.NewRedisClient()
	_, err := rd.Ping().Result()
	assert.NotNil(t, err, `failed to connect redis`)
}