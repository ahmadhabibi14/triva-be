package service

import (
	"errors"
	"triva/helper"
	"triva/internal/bootstrap/database"
	"triva/internal/bootstrap/logger"
	"triva/internal/repository/users"
	"triva/internal/request"
	"triva/internal/response"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Db *database.Database
}

func NewAuthService(Db *database.Database) *AuthService {
	return &AuthService{Db: Db}
}

func (as *AuthService) Login(in request.LoginIn) (out response.LoginOut, err error) {
	user := users.NewUserMutator(as.Db)
	user.Username = in.Username
	if err = user.FindByUsername(); err != nil {
		return
	}

	if passwordMatch := helper.VerifyPassword(in.Password, user.Password); passwordMatch != nil {
		errMsg := errors.New(`password does not match`)
		logger.Log.Err(passwordMatch).Msg(errMsg.Error())
		err = errMsg
		return
	}

	out.SessionKey = helper.RandString(35)

	user.HideSecrets()
	out.User = user

	session := users.NewSessionMutator(as.Db)
	err = session.SetSession(out.SessionKey, user.Id, user.Username)

	return
}

func (as *AuthService) Register(in request.RegisterIn) (out response.RegisterOut, err error) {
	user := users.NewUserMutator(as.Db)
	user.Username = in.Username
	user.FullName = in.FullName
	user.Email = in.Email

	hashedPassword, errGen := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if errGen != nil {
		err = errors.New(`failed to set password`)
		logger.Log.Err(errGen).Msg(err.Error())
		return
	}

	user.Password = string(hashedPassword)

	err = user.CreateUser()
	if err != nil {
		return
	}

	user.HideSecrets()
	out.User = user

	return
}