package service

import (
	"triva/helper"
	"triva/internal/repository/users"

	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
)

type AuthService struct {
	db *sqlx.DB
	rd *redis.Client
}

func NewAuthService(db *sqlx.DB, rd *redis.Client) *AuthService {
	return &AuthService{
		db: db,
		rd: rd,
	}
}

func (as *AuthService) SetSession(userId, username string) (string, error) {
	session := users.NewSession(userId, username, true)

	sessionKey := helper.RandString(35)
	as.rd.Set(users.SESSION_PREFIX + sessionKey, session, users.SESSION_EXPIRED)
	return sessionKey, nil
}