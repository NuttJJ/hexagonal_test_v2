package handler

import (
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	*fiber.App
}

type RouterParams struct {
	UserHandler  *UserHandler
	OAuthHandler *OAuthHandler
}

func NewRouter(p RouterParams) (*Router, error) {
	app := fiber.New(fiber.Config{
		AppName: "Hexagonal Architecture API v1.0.0",
	})

	// Middleware
	// app.Use(logger.New())
	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins: "*",
	// 	AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	// 	AllowHeaders: "Origin,Content-Type,Accept,Authorization",
	// }))

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "OK",
			"message": "Server is running",
		})
	})

	api := app.Group("/api")
	v1 := api.Group("/v1")

	user := v1.Group("/user")
	{
		user.Post("/register", p.UserHandler.RegisterEndpoint)
		user.Get("/list", p.UserHandler.ListUsersEndpoint)
	}

	// OAuth routes
	auth := app.Group("/auth")
	auth.Get("/google", p.OAuthHandler.GoogleLogin)
	auth.Get("/google/callback", p.OAuthHandler.GoogleCallback)

	// Protected routes (middleware to be implemented)
	// protected := api.Group("/", authMiddleware)
	// protected.Get("/users", p.UserHandler.ListUsersEndpoint)

	return &Router{App: app}, nil
}
