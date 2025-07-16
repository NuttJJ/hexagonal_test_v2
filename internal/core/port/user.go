package port

import (
	"hexagonal-test-v2/internal/core/domain"
)

type UserRepository interface {
	CreateUser(user *domain.User) (*domain.User, error)
	ListUsers() ([]*domain.User, error)
}

type UserService interface {
	Register(user *domain.User) (*domain.User, error)
	ListUsers() ([]*domain.User, error)
}
