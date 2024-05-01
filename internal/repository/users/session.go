package users

import (
	"errors"
	"time"

	"github.com/go-redis/redis"
	"github.com/goccy/go-json"
)

const (
	SESSION_PREFIX  = `session:`
	SESSION_EXPIRED = (24 * time.Hour) * 60
)

type Session struct {
	UserID        string `json:"user_id"`
	Username      string `json:"username"`
	Authenticated bool   `json:"authenticated"`
}

func NewSession(userId, username string, authenticated bool) *Session {
	return &Session{
		UserID:        userId,
		Username:      username,
		Authenticated: authenticated,
	}
}

func NewSessionJSON(userId, username string, authenticated bool) ([]byte, error) {
	session := &Session{
		UserID:        userId,
		Username:      username,
		Authenticated: authenticated,
	}

	sessionJson, err := json.Marshal(session)

	if err != nil {
		return nil, err
	}

	return sessionJson, nil
}

func GetSessionData(sessionId string, rds *redis.Client) (userSession Session, err error) {
	sessionKey := SESSION_PREFIX + sessionId

	userData, err := rds.Get(sessionKey).Result()
	if err != nil {
		err = errors.New(`session not found`)
		return
	}

	err = json.Unmarshal([]byte(userData), &userSession)
	if err != nil {
		err = errors.New(`invalid session data`)
		return
	}

	return
}
