package models

type User struct {
	ID           int    `json:"id"`
	UserID       string `json:"user_id"`
	Username     string `json:"username"`
	DisplayName  string `json:"display_name"`
	AvatarURL    string `json:"avatar_url"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	Role         string `json:"role"`
	CreatedAt    string `json:"created_at"`
	DeletedAt    string `json:"deleted_at"`
	UpdatedAt    string `json:"updated_at"`
}

type UserResponse struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
	AvatarURL   string `json:"avatar_url"`
	Email       string `json:"email"`
	Role        string `json:"role"`
}

type CreateUserRequest struct {
	Username    string `json:"username" validate:"required"`
	DisplayName string `json:"display_name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=8"`
	Role        string `json:"role" validate:"omitempty,oneof=admin user"`
}

type UpdateUserRequest struct {
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Role        string `json:"role" validate:"omitempty,oneof=admin user"`
}
