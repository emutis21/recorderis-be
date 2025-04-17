package drivens

import (
	"context"
	"errors"
	"recorderis/cmd/services/auth/models"
	"recorderis/cmd/services/auth/ports/drivens"
	"time"

	"gorm.io/gorm"
)

var _ drivens.ForTokenRepository = (*GormTokenRepository)(nil)

type GormTokenRepository struct {
	db *gorm.DB
}

func NewGormTokenRepository(db *gorm.DB) *GormTokenRepository {
	err := db.AutoMigrate(&models.RefreshTokenDetails{}, &models.SessionInfo{})
	if err != nil {
		panic("Failed to migrate token/session tables: " + err.Error())
	}
	return &GormTokenRepository{db: db}
}

func (r *GormTokenRepository) SaveRefreshToken(ctx context.Context, token *models.RefreshTokenDetails) error {
	return r.db.WithContext(ctx).Create(token).Error
}

func (r *GormTokenRepository) GetRefreshToken(ctx context.Context, tokenID string) (*models.RefreshTokenDetails, error) {
	var tokenDetails models.RefreshTokenDetails
	err := r.db.WithContext(ctx).Where("token_id = ?", tokenID).First(&tokenDetails).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("refresh token not found")
		}
		return nil, err
	}
	return &tokenDetails, nil
}

func (r *GormTokenRepository) UpdateLastUsed(ctx context.Context, tokenID string) error {
	return r.db.WithContext(ctx).Model(&models.RefreshTokenDetails{}).Where("token_id = ?", tokenID).Update("last_used_at", time.Now()).Error
}

func (r *GormTokenRepository) RevokeRefreshToken(ctx context.Context, tokenID string) error {
	return r.db.WithContext(ctx).Where("token_id = ?", tokenID).Delete(&models.RefreshTokenDetails{}).Error
}

func (r *GormTokenRepository) RevokeAllUserTokens(ctx context.Context, userID string) error {
	return r.db.WithContext(ctx).Where("user_id = ?", userID).Delete(&models.RefreshTokenDetails{}).Error
}

func (r *GormTokenRepository) CreateSession(ctx context.Context, session *models.SessionInfo) error {
	return r.db.WithContext(ctx).Create(session).Error
}

func (r *GormTokenRepository) UpdateSessionActivity(ctx context.Context, sessionID string) error {
	return r.db.WithContext(ctx).Model(&models.SessionInfo{}).Where("session_id = ?", sessionID).Update("last_activity", time.Now()).Error
}

func (r *GormTokenRepository) GetUserSessions(ctx context.Context, userID string) ([]models.SessionInfo, error) {
	var sessions []models.SessionInfo
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&sessions).Error
	if err != nil {
		return nil, err
	}
	return sessions, nil
}

func (r *GormTokenRepository) DeleteSession(ctx context.Context, sessionID string) error {
	return r.db.WithContext(ctx).Where("session_id = ?", sessionID).Delete(&models.SessionInfo{}).Error
}
