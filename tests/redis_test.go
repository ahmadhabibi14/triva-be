package tests

import (
	"testing"
	"triva/configs"
	"triva/helper"
	"triva/internal/repository/users"

	"github.com/stretchr/testify/assert"
)

func TestConnectRedis(t *testing.T) {
	rd := configs.NewRedisClient()
	pong, err := rd.Ping().Result()

	assert.Nil(t, err, `failed to connect redis`)

	t.Log(`connected to redis:`, pong)

	t.Run(`insertSessionMustSucceed`, func(t *testing.T) {
		session, err := users.NewSessionJSON(`77b44fc9-eef1-487d-965f-17a2ef1f99f4`, `habibi14`, true)

		assert.Nil(t, err, `failed to marshal json`)

		sessionKey := helper.RandString(35)
		status := rd.Set(users.SESSION_PREFIX + sessionKey, session, users.SESSION_EXPIRED)

		t.Log(`redis status cmd:`, status)
	})
}