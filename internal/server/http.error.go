package server

import "github.com/gofiber/fiber/v3"

type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// SetHTTPError - возвращает ошибку в формате HTTPError (Code, Message)
func SetHTTPError(c fiber.Ctx, errStr string, code int) error {
	h := HTTPError{
		Code:    code,
		Message: errStr,
	}

	return c.Status(code).JSON(h)
}
