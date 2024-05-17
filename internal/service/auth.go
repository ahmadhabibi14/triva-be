package service

import (
	"errors"
	"triva/helper"
	"triva/internal/bootstrap/database"
	"triva/internal/repository/users"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Db *database.Database
}

func NewAuthService(Db *database.Database) *AuthService {
	return &AuthService{Db: Db}
}

func (as *AuthService) Login(username, password string) (sessionKey string, err error) {
	user := users.NewUserMutator(as.Db)
	user.Username = username
	if !user.FindUsernamePassword() {
		err = errors.New(`username not found`)
		return
	}

	if passwordMatch := helper.VerifyPassword(password, user.Password); passwordMatch != nil {
		err = errors.New(`password does not match`)
		return
	}

	sessionKey = helper.RandString(35)

	session := users.NewSessionMutator(as.Db)
	err = session.SetSession(sessionKey, user.Id, user.Username)

	return
}

func (as *AuthService) Register(username, fullName, email, password string) error {
	user := users.NewUserMutator(as.Db)
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