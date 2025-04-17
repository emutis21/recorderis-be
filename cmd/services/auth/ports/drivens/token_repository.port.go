package drivens

import (
	"context"
	"recorderis/cmd/services/auth/models"
)

type ForTokenRepository interface {
	// Métodos para refresh tokens
	SaveRefreshToken(ctx context.Context, token *models.RefreshTokenDetails) error
	GetRefreshToken(ctx context.Context, tokenHash string) (*models.RefreshTokenDetails, error)
	UpdateLastUsed(ctx context.Context, tokenID string) error
	RevokeRefreshToken(ctx context.Context, tokenID string) error
	RevokeAllUserTokens(ctx context.Context, userID string) error

	// Métodos para sesiones activas
	CreateSession(ctx context.Context, session *models.SessionInfo) error
	UpdateSessionActivity(ctx context.Context, sessionID string) error
	GetUserSessions(ctx context.Context, userID string) ([]models.SessionInfo, error)
	DeleteSession(ctx context.Context, sessionID string) error
}
