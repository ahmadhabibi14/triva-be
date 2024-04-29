package service

import (
	"errors"
	"triva/helper"
	"triva/internal/repository/users"

	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
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

func (as *AuthService) setSession(userId, username string) (string, error) {
	session := users.NewSession(userId, username, true)

	sessionKey := helper.RandString(35)
	as.rd.Set(users.SESSION_PREFIX + sessionKey, session, users.SESSION_EXPIRED)
	return sessionKey, nil
}

func (as *AuthService) Login(username, password string) (sessionKey string, err error) {
	user := users.NewUserMutator(as.db)
	user.Username = username
	if !user.FindUsernamePassword() {
		err = errors.New(`username not found`)
		return
	}

	if passwordMatch := helper.VerifyPassword(password, user.Password); passwordMatch != nil {
		err = errors.New(`password does not match`)
		return
	}

	sessionKey, err = as.setSession(user.Id, username)
	return
}

func (as *AuthService) Register(username, fullName, email, password string) error {
	user := users.NewUserMutator(as.db)
	user.Username = username
	user.FullName = fullName
	user.Email = email

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New(`failed to set password`)
	}

	user.Password = string(hashedPassword)

	return user.CreateUser()
}