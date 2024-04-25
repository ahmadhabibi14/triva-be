package service

import (
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