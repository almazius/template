package http

import "github.com/gofiber/fiber/v3"

func (handler *AuthHandler) LoginHandler() fiber.Handler {
	return func(c fiber.Ctx) error {
		///
		return c.JSON("popa")
	}
}
