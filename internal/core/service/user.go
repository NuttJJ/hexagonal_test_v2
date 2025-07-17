package service

import (
	"hexagonal-test-v2/internal/core/domain"
	"hexagonal-test-v2/internal/core/port"
	"time"
)

type UserService struct {
	repo port.UserRepository
}

// FindByGoogleID implements port.UserService.
func (us *UserService) FindByGoogleID(googleID string) (*domain.User, error) {
	return us.repo.FindByGoogleID(googleID)
}

// UpdateOAuthUser implements port.UserService.
func (us *UserService) UpdateOAuthUser(user *domain.User, oauthUser *domain.OAuthUser) (*domain.User, error) {
	// อัพเดทข้อมูลจาก OAuth
	user.Name = oauthUser.Name
	user.Avatar = oauthUser.Picture
	user.IsVerified = oauthUser.Verified
	user.UpdatedAt = time.Now()

	// ถ้ายังไม่มี GoogleID ให้เพิ่ม
	if user.GoogleID == "" {
		user.GoogleID = oauthUser.ID
	}

	return us.repo.UpdateUser(user)
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

func (us *UserService) FindByEmail(email string) (*domain.User, error) {
	return us.repo.FindByEmail(email)
}

func (us *UserService) CreateOAuthUser(oauthUser *domain.OAuthUser) (*domain.User, error) {
	user := &domain.User{
		Name:       oauthUser.Name,
		Email:      oauthUser.Email,
		GoogleID:   oauthUser.ID,
		Avatar:     oauthUser.Picture,
		Provider:   string(oauthUser.Provider),
		IsVerified: oauthUser.Verified,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	return us.repo.CreateUser(user)
}
