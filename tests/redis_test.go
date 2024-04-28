package tests

import (
	"testing"
	"triva/configs"

	"github.com/stretchr/testify/assert"
)

func TestConnectRedis(t *testing.T) {
	rd := configs.NewRedisClient()
	pong, err := rd.Ping().Result()

	assert.Nil(t, err, `failed to connect redis`)

	t.Log(`connected to redis:`, pong)
}