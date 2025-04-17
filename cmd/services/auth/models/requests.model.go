package models

type LoginRequest struct {
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required"`
	RememberMe bool   `json:"remember_me"`
	DeviceType string `json:"device_type" validate:"required,oneof=web mobile tablet"`
	DeviceName string `json:"device_name,omitempty"`

	IPAddress string `json:"-"`
	UserAgent string `json:"-"`
}

type RegisterRequest struct {
	Username    string `json:"username" validate:"required"`
	DisplayName string `json:"display_name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=8"`
	DeviceType  string `json:"device_type" validate:"required,oneof=web mobile tablet"`
	DeviceName  string `json:"device_name,omitempty"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type LogoutRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}
