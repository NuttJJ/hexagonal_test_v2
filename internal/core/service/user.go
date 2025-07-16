package service

import (
	"hexagonal-test-v2/internal/core/domain"
	"hexagonal-test-v2/internal/core/port"
)

type UserService struct {
	repo port.UserRepository
}

func NewUserService(userRepository port.UserRepository) *UserService {
	return &UserService{userRepository}
}

func (us *UserService) Register(user *domain.User) (*domain.User, error) {
	user, err := us.repo.CreateUser(user)
	if err != nil {
		if err == domain.ErrConflictingData {
			return nil, err
		}
		return nil, domain.ErrInternal
	}
	return user, nil
}

func (us *UserService) ListUsers() ([]*domain.User, error) {
	users, err := us.repo.ListUsers()
	if err != nil {
		return nil, domain.ErrInternal
	}

	return users, nil
}
