package port

import (
	"hexagonal-test-v2/internal/core/domain"
)

// internal/core/port/oauth.go
type OAuthService interface {
	GetAuthURL(state string) string
	HandleCallback(code, state string) (*domain.User, error)
	ValidateState(state string) error
	GenerateState() string
}
