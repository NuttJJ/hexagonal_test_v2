package main

import (
	"log"
	"log/slog"

	"hexagonal-test-v2/internal/adapter/handler"
	"hexagonal-test-v2/internal/adapter/storage/mongo"
	"hexagonal-test-v2/internal/adapter/storage/mongo/repository"
	"hexagonal-test-v2/internal/core/service"
)

func main() {
	// Connect to MongoDB
	db, err := mongo.ConnectMongoDB()
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	defer db.Close()

	// Create dependencies
	userRepo := repository.NewMongoUserRepository(db.Database)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Create OAuth dependencies
	oauthService := service.NewOAuthService(userService)
	oauthHandler := handler.NewOAuthHandler(oauthService, userService)

	// Create router
	router, err := handler.NewRouter(handler.RouterParams{
		UserHandler:  userHandler,
		OAuthHandler: oauthHandler,
	})
	if err != nil {
		log.Fatal("Failed to create router:", err)
	}

	// Start server
	slog.Info("🚀 Server starting on :8080")
	if err := router.Listen("127.0.0.1:8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
