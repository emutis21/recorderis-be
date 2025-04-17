package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type User struct {
	ID           int       `gorm:"primaryKey"`
	UserID       string    `gorm:"type:uuid;unique;not null;default:uuid_generate_v4()"`
	Username     string    `gorm:"type:varchar(255);not null"`
	DisplayName  string    `gorm:"type:varchar(255);not null"`
	AvatarURL    string    `gorm:"type:varchar(255)"`
	Email        string    `gorm:"type:varchar(255);unique;not null"`
	PasswordHash string    `gorm:"type:varchar(255);not null"`
	Role         Role      `gorm:"type:user_role;not null;default:'user'"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.UserID == "" {
		u.UserID = uuid.New().String()
	}

	return nil
}

func (User) TableName() string {
	return "users"
}
