package controller

import (
	"errors"
	"triva/configs"
	"triva/internal/bootstrap/database"
	"triva/internal/bootstrap/logger"
	"triva/internal/repository/users"

	"github.com/gofiber/fiber/v2"
)

const (
	errMsgUnauthorized	= `you are unauthorized to process this operation`
	errMsgInvalidKey		= `invalid session key`
)

func getSession(db *database.Database, c *fiber.Ctx) (session *users.Session, err error) {
	sessionId := c.Cookies(configs.AUTH_COOKIE, ``)
	apiKey := c.Get("X-API-KEY", ``)

	var KEY string = sessionId
	if sessionId == `` { KEY = apiKey }
	
	if KEY == `` {
		err = errors.New(errMsgUnauthorized)
		return
	}

	session = users.NewSessionMutator(db)
	
	err = session.GetSession(KEY);
	if err != nil {
		logger.Log.Err(err).Msg("cannot get session data for " + KEY)
		
		err = errors.New(errMsgInvalidKey)
		return
	}

	return
}