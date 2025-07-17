package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (r *Router) AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization header required",
		})
	}

	// tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	// // Validate JWT token
	// claims, err := validateJWT(tokenString)
	// if err != nil {
	// 	return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
	// 		"error": "Invalid token",
	// 	})
	// }

	// c.Locals("user", claims)
	return c.Next()
}
