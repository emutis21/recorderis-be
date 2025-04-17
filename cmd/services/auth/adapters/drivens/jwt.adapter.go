package drivens

import (
	"fmt"
	"recorderis/cmd/services/auth/models"
	"recorderis/cmd/services/auth/ports/drivens"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var _ drivens.ForTokenManager = (*JWTAdapter)(nil)


type JWTAdapter struct {
	config models.TokenConfig
}

func NewJWTAdapter(config models.TokenConfig) *JWTAdapter {
	return &JWTAdapter{
		config: config,
	}
}

func (j *JWTAdapter) GenerateToken(userID string, role string) (string, error) {
	return j.createToken(userID, role, models.AccessToken)
}

func (j *JWTAdapter) GenerateRefreshToken(userID string) (string, error) {
	return j.createToken(userID, "", models.RefreshToken)
}

func (j *JWTAdapter) ValidateToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return j.config.SigningKey, nil
	})

	if err != nil {
		if err == jwt.ErrTokenExpired {
			return "", fmt.Errorf(models.ErrExpiredToken)
		}
		return "", fmt.Errorf(models.ErrInvalidToken)
	}

	claims, ok := token.Claims.(*models.Claims)
	if !ok || !token.Valid {
		return "", fmt.Errorf(models.ErrInvalidToken)
	}

	if claims.Type != models.AccessToken {
		return "", fmt.Errorf("invalid token type")
	}

	return claims.UserID, nil
}

func (j *JWTAdapter) InvalidateToken(token string) error {
	return nil
}

func (j *JWTAdapter) createToken(userID string, role string, tokenType models.TokenType) (string, error) {
	var expirationTime time.Time

	if tokenType == models.AccessToken {
		expirationTime = time.Now().Add(j.config.AccessTokenDuration)
	} else {
		expirationTime = time.Now().Add(j.config.RefreshTokenDuration)
	}

	claims := &models.Claims{
		UserID: userID,
		Role:   role,
		Type:   tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    j.config.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(j.config.SigningKey)
	if err != nil {
		return "", fmt.Errorf("error signing token: %w", err)
	}

	return tokenString, nil
}
