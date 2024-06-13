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

func ReadBody[T any](c *fiber.Ctx, b []byte) (T, error) {
	var body T
	err := c.BodyParser(&body)
	if err != nil {
		return body, errors.New(`invalid payload, please check your request and try again`)
	}

	err = ValidateStruct(body)
	if err != nil {
		return body, err
	}

	return body, nil
}