package tests

import (
	"testing"
	"triva/internal/repository/users"

	"github.com/stretchr/testify/assert"
)

const (
	sessionKey	= `thiskey`
	userId			= `user-id-uuid-valid`
	username		= `johndoe998` 
)

func TestSetSession(t *testing.T) {
	session := users.NewSessionMutator(DB)
	err := session.SetSession(sessionKey, userId, username)

	assert.Nil(t, err, `failed to set session`)
}

func TestGetSession(t *testing.T) {
	session := users.NewSessionMutator(DB)
	err := session.GetSession(sessionKey)

	assert.Nil(t, err, `session not found`)

	t.Log(`User ID:`, session.UserID)
	t.Log(`Username:`, session.Username)
}