package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hexagonal-test-v2/internal/core/domain"
	"hexagonal-test-v2/internal/core/port"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type oauthService struct {
	config       *oauth2.Config
	userService  port.UserService
	stateStorage map[string]*domain.OAuthState
}

func NewOAuthService(userService port.UserService) port.OAuthService {
	return &oauthService{
		config: &oauth2.Config{
			ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
			RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
			Scopes:       []string{"openid", "profile", "email"},
			Endpoint:     google.Endpoint,
		},
		userService:  userService,
		stateStorage: make(map[string]*domain.OAuthState),
	}
}

func (s *oauthService) GenerateState() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	state := hex.EncodeToString(bytes)

	// Store state with expiration
	s.stateStorage[state] = &domain.OAuthState{
		State:     state,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(10 * time.Minute),
	}

	return state
}

func (s *oauthService) GetAuthURL(state string) string {
	return s.config.AuthCodeURL(state)
}

func (s *oauthService) ValidateState(state string) error {
	storedState, exists := s.stateStorage[state]
	if !exists {
		return fmt.Errorf("invalid state parameter")
	}

	if time.Now().After(storedState.ExpiresAt) {
		delete(s.stateStorage, state)
		return fmt.Errorf("state parameter expired")
	}

	// Remove used state
	delete(s.stateStorage, state)
	return nil
}

func (s *oauthService) HandleCallback(code, state string) (*domain.User, error) {
	// Validate state
	if err := s.ValidateState(state); err != nil {
		return nil, err
	}

	// Exchange code for token
	token, err := s.config.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange code for token: %v", err)
	}

	// Get user info from Google
	oauthUser, err := s.getUserInfo(token)
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %v", err)
	}

	// Find or create user
	user, err := s.userService.FindByEmail(oauthUser.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %v", err)
	}

	if user == nil {
		// Create new user
		user, err = s.userService.CreateOAuthUser(oauthUser)
		if err != nil {
			return nil, fmt.Errorf("failed to create user: %v", err)
		}
	} else {
		// Update existing user
		user, err = s.userService.UpdateOAuthUser(user, oauthUser)
		if err != nil {
			return nil, fmt.Errorf("failed to update user: %v", err)
		}
	}

	return user, nil
}

func (s *oauthService) getUserInfo(token *oauth2.Token) (*domain.OAuthUser, error) {
	client := s.config.Client(context.Background(), token)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get user info: status %d", resp.StatusCode)
	}

	var googleUser struct {
		ID            string `json:"id"`
		Email         string `json:"email"`
		Name          string `json:"name"`
		Picture       string `json:"picture"`
		VerifiedEmail bool   `json:"verified_email"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&googleUser); err != nil {
		return nil, err
	}

	return &domain.OAuthUser{
		ID:       googleUser.ID,
		Email:    googleUser.Email,
		Name:     googleUser.Name,
		Picture:  googleUser.Picture,
		Provider: domain.ProviderGoogle,
		Verified: googleUser.VerifiedEmail,
	}, nil
}
