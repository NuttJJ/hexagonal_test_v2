package handler

import (
	"hexagonal-test-v2/internal/core/port"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type OAuthHandler struct {
	oauthService port.OAuthService
	userService  port.UserService
}

func NewOAuthHandler(oauthService port.OAuthService, userService port.UserService) *OAuthHandler {
	return &OAuthHandler{
		oauthService: oauthService,
		userService:  userService,
	}
}

func (h *OAuthHandler) GoogleLogin(c *fiber.Ctx) error {
	state := h.oauthService.GenerateState()
	url := h.oauthService.GetAuthURL(state)

	return c.Redirect(url)
}

func (h *OAuthHandler) GoogleCallback(c *fiber.Ctx) error {
	code := c.Query("code")
	state := c.Query("state")

	if code == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Authorization code not provided",
		})
	}

	user, err := h.oauthService.HandleCallback(code, state)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "OAuth authentication failed",
		})
	}

	// Generate JWT token
	// token, err := generateJWT(user)
	// if err != nil {
	// 	return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": "Failed to generate token",
	// 	})
	// }

	return c.JSON(fiber.Map{
		"message": "Login successful",
		// "token":   token,
		"user": user,
	})
}
