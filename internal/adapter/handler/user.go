package handler

import (
	"hexagonal-test-v2/internal/core/domain"
	"hexagonal-test-v2/internal/core/port"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService port.UserService
}

func NewUserHandler(userService port.UserService) *UserHandler {
	return &UserHandler{userService}
}

type registerRequest struct {
	Name                 string `json:"name" validate:"required"`
	Email                string `json:"email,omitempty" validate:"required,email"`
	Password             string `json:"password,omitempty" validate:"required,min=8"`
	PasswordConfirmation string `json:"password_confirmation,omitempty" validate:"required,min=8"`
}

func (h *UserHandler) RegisterEndpoint(c *fiber.Ctx) error {
	var req registerRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Basic validation
	if req.Password != req.PasswordConfirmation {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Password and password confirmation do not match",
		})
	}

	user := domain.User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := h.userService.Register(&user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to register user",
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
		"user": fiber.Map{
			"id":         user.ID,
			"name":       user.Name,
			"email":      user.Email,
			"created_at": user.CreatedAt,
			"updated_at": user.UpdatedAt,
		},
	})
}

func (h *UserHandler) ListUsersEndpoint(c *fiber.Ctx) error {
	users, err := h.userService.ListUsers()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch users",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"users": users,
	})
}
