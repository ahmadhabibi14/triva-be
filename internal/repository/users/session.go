package users

import (
	"errors"
	"fmt"
	"time"
	"triva/internal/database"

	"github.com/goccy/go-json"
)

const (
	SESSION_PREFIX  = `session:`
	SESSION_EXPIRED = (24 * time.Hour) * 60
)

type Session struct {
	Db						*database.Database `json:"-"` 
	UserID        string				`json:"user_id"`
	Username      string				`json:"username"`
}

func NewSessionMutator(Db *database.Database) *Session {
	return &Session{Db: Db}
}

func (s *Session) SetSession(sessionKey, userId, username string) error {
	s.UserID = userId
	s.Username = username

	sessionJSON, err := json.Marshal(&s)
	if err != nil {
		return errors.New(`failed to marshal session data`)
	}

	err = s.Db.RD.Set(SESSION_PREFIX + sessionKey, sessionJSON, SESSION_EXPIRED).Err()
	if err != nil {
		return fmt.Errorf(`failed to set session: %v`, err)
	}

	return nil
}

func (s *Session) GetSession(sessionKey string) error {
	sessionData, err := s.Db.RD.Get(SESSION_PREFIX + sessionKey).Result()
	if err != nil {
		return errors.New(`session not found`)
	}

	err = json.Unmarshal([]byte(sessionData), s)
	if err != nil {
		return errors.New(`invalid session data`)
	}

	return nil
}