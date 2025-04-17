package drivers

import (
	"context"
	"recorderis/cmd/services/auth/models"
	api_models "recorderis/cmd/services/api/models"
)

type ForAuth interface {
	/* Public */
	Login(ctx context.Context, req *models.LoginRequest) (*models.TokenResponse, error)
	Register(ctx context.Context, req *models.RegisterRequest) (*models.TokenResponse, error)
	RefreshToken(ctx context.Context, refreshToken string) (*models.TokenResponse, error)

	/* Private */
	Logout(ctx context.Context, tokenID string) error
	GetUserById(ctx context.Context, userID string) (*api_models.User, error)
}
