package users

import (
	"time"

	"github.com/goccy/go-json"
)

const (
	SESSION_PREFIX 	= `session:`
	SESSION_EXPIRED = (24 * time.Hour) * 60
)

type Session struct {
	UserID string `json:"user_id"`
	Username string `json:"username"`
	Authenticated bool `json:"authenticated"`
}

func NewSession(userId, username string, authenticated bool) *Session {
	return &Session{
		UserID: userId,
		Username: username,
		Authenticated: authenticated,
	}
}

func NewSessionJSON(userId, username string, authenticated bool) ([]byte, error) {
	session := &Session{
		UserID: userId,
		Username: username,
		Authenticated: authenticated,
	}

	sessionJson, err := json.Marshal(session)

	if err != nil {
		return nil, err
	}

	return sessionJson, nil
}