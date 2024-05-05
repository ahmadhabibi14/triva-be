package helper

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type HTTPResponse struct {
	Code   int		`json:"code"`
	Status string	`json:"status"`
	Errors string	`json:"errors"`
	Data   any		`json:"data"`
}

func NewHTTPResponse(code int, errors string, data any) HTTPResponse {
	return HTTPResponse{
		Code:   code,
		Status: http.StatusText(code),
		Errors: errors,
		Data:   data,
	}
}

func ReadJSON[T any](c *fiber.Ctx, b []byte) (T, error) {
	var body T
	err := c.BodyParser(&body)
	if err != nil {
		return body, errors.New(`invalid payload, please check your request and try again`)
	}

	errvalid := ValidateStruct(body)
	if errvalid != nil {
		return body, errvalid
	}

	return body, nil
}
