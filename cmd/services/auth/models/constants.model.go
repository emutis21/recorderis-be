package models

import "time"

const (
	// Token Durations
	AccessTokenDuration            = 15 * time.Minute
	DefaultRefreshTokenDuration    = 24 * time.Hour
	RememberMeRefreshTokenDuration = 30 * 24 * time.Hour

	// Device Types
	DeviceTypeWeb    = "web"
	DeviceTypeMobile = "mobile"
	DeviceTypeTablet = "tablet"

	// Error Messages
	ErrInvalidCredentials = "credenciales inválidas"
	ErrUserNotFound       = "usuario no encontrado"
	ErrUserExists         = "usuario ya existe"
	ErrInvalidToken       = "token inválido"
	ErrExpiredToken       = "token expirado"
)
