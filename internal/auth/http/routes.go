package http

import (
	"github.com/gofiber/fiber/v3"
	"template/internal/auth/usecase"
)

type AuthHandler struct {
	app         *fiber.App
	authService *usecase.AuthService
}

func NewAuthHandler(app *fiber.App, authService *usecase.AuthService) *AuthHandler {
	return &AuthHandler{
		app:         app,
		authService: authService,
	}
}

func MapAuthHandlers(handler *AuthHandler) {
	handler.app.Get("/login", handler.LoginHandler())
	/// ...
}
