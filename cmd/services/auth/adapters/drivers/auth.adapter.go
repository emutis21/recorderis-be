package drivers

import (
	"context"
	"fmt"
	api_models "recorderis/cmd/services/api/models"
	"recorderis/cmd/services/auth/models"
	"recorderis/cmd/services/auth/ports/drivens"
	"recorderis/internals/errors"
	"time"

	"github.com/google/uuid"
)

type AuthAdapter struct {
	userRepo    drivens.ForUserRepository
	tokenMgr    drivens.ForTokenManager
	tokenRepo   drivens.ForTokenRepository
	passwordMgr drivens.ForPasswordManager
}

func NewAuthAdapter(
	userRepo drivens.ForUserRepository,
	tokenMgr drivens.ForTokenManager,
	tokenRepo drivens.ForTokenRepository,
	passwordMgr drivens.ForPasswordManager,
) *AuthAdapter {
	return &AuthAdapter{
		userRepo:    userRepo,
		tokenMgr:    tokenMgr,
		tokenRepo:   tokenRepo,
		passwordMgr: passwordMgr,
	}
}

func (a *AuthAdapter) Register(ctx context.Context, req *models.RegisterRequest) (*models.TokenResponse, error) {
	if _, err := a.userRepo.FindUserByEmail(ctx, req.Email); err == nil {
		return nil, fmt.Errorf("email already registered")
	}

	hashedPassword, err := a.passwordMgr.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %w", err)
	}

	user, err := a.userRepo.CreateUser(ctx, &api_models.CreateUserRequest{
		Username:    req.Username,
		DisplayName: req.DisplayName,
		Email:       req.Email,
		Password:    hashedPassword,
	})
	if err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
	}

	accessToken, err := a.tokenMgr.GenerateToken(user.UserID, user.Role)
	if err != nil {
		return nil, fmt.Errorf("error generating access token: %w", err)
	}

	refreshToken, err := a.tokenMgr.GenerateRefreshToken(user.UserID)
	if err != nil {
		return nil, fmt.Errorf("error generating refresh token: %w", err)
	}

	return &models.TokenResponse{
		AccessToken:  accessToken,
		TokenType:    "Bearer",
		ExpiresIn:    int64(time.Hour.Seconds() * 24),
		RefreshToken: refreshToken,
	}, nil
}

func (a *AuthAdapter) Login(ctx context.Context, req *models.LoginRequest) (*models.TokenResponse, error) {
	user, err := a.userRepo.FindUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	if err := a.passwordMgr.ValidatePassword(req.Password, user.PasswordHash); err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	accessToken, err := a.tokenMgr.GenerateToken(user.UserID, user.Role)
	if err != nil {
		return nil, fmt.Errorf("error generating access token: %w", err)
	}

	refreshTokenDetails := &models.RefreshTokenDetails{
		TokenID:      uuid.New().String(),
		UserID:       user.UserID,
		DeviceType:   req.DeviceType,
		DeviceName:   req.DeviceName,
		IsRememberMe: req.RememberMe,
		ExpiresAt:    time.Now().Add(a.getRefreshTokenDuration(req.RememberMe)),
		LastUsedAt:   time.Now(),
	}

	if err := a.tokenRepo.SaveRefreshToken(ctx, refreshTokenDetails); err != nil {
		return nil, fmt.Errorf("error saving refresh token: %w", err)
	}

	session := &models.SessionInfo{
		SessionID:      uuid.New().String(),
		UserID:         user.UserID,
		RefreshTokenID: refreshTokenDetails.TokenID,
		IPAddress:      req.IPAddress,
		UserAgent:      req.UserAgent,
		LastActivity:   time.Now(),
	}

	if err := a.tokenRepo.CreateSession(ctx, session); err != nil {
		return nil, fmt.Errorf("error creating session: %w", err)
	}

	return &models.TokenResponse{
		AccessToken:  accessToken,
		TokenType:    "Bearer",
		ExpiresIn:    int64(models.AccessTokenDuration.Seconds()),
		RefreshToken: refreshTokenDetails.TokenID,
	}, nil
}

func (a *AuthAdapter) RefreshToken(ctx context.Context, refreshToken string) (*models.TokenResponse, error) {
	tokenDetails, err := a.tokenRepo.GetRefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, errors.NewUnauthorizedError("Invalid refresh token", err)
	}

	if time.Now().After(tokenDetails.ExpiresAt) {
		a.tokenRepo.RevokeRefreshToken(ctx, tokenDetails.TokenID)
		return nil, fmt.Errorf("refresh token expired")
	}

	user, err := a.userRepo.FindUserById(ctx, tokenDetails.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	accessToken, err := a.tokenMgr.GenerateToken(user.UserID, user.Role)
	if err != nil {
		return nil, fmt.Errorf("error generating access token: %w", err)
	}

	if err := a.tokenRepo.UpdateLastUsed(ctx, tokenDetails.TokenID); err != nil {
		return nil, fmt.Errorf("error updating token: %w", err)
	}

	return &models.TokenResponse{
		AccessToken: accessToken,
		TokenType:   "Bearer",
		ExpiresIn:   int64(models.AccessTokenDuration.Seconds()),
	}, nil
}

func (a *AuthAdapter) Logout(ctx context.Context, tokenID string) error {
	return a.tokenRepo.RevokeRefreshToken(ctx, tokenID)
}

func (a *AuthAdapter) getRefreshTokenDuration(rememberMe bool) time.Duration {
	if rememberMe {
		return models.RememberMeRefreshTokenDuration
	}
	return models.DefaultRefreshTokenDuration
}

func (a *AuthAdapter) GetUserById(ctx context.Context, userID string) (*api_models.User, error) {
	return a.userRepo.FindUserById(ctx, userID)
}
