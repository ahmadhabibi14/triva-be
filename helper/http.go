package helper

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type HTTPResponse struct {
	Errors 	string	`json:"errors"`
	Data   	any			`json:"data,omitempty"`
}

func NewHTTPResponse(errors string, data any) HTTPResponse {
	return HTTPResponse{
		Errors: errors,
		Data:   data,
	}
}

func ReadBody[T any](c *fiber.Ctx) (out T, err error) {
	err = c.BodyParser(&out)
	if err != nil {
		err = errors.New(`invalid payload, please check your request and try again`)
		return
	}

	err = ValidateStruct(out)
	if err != nil {
		return
	}

	return
}