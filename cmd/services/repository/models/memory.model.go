package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Memory struct {
	ID          int       `gorm:"primaryKey"`
	MemoryID    string    `gorm:"type:uuid;unique;not null;default:uuid_generate_v4()"`
	Index       int       `gorm:"not null"`
	IndexPtr    *int      `gorm:"-"`
	UserID      string    `gorm:"type:uuid;not null"`
	Title       string    `gorm:"type:varchar(255);not null"`
	Date        time.Time `gorm:"not null"`
	IsPublic    bool      `gorm:"default:false;not null"`
	IsPublicPtr *bool     `gorm:"-"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt

	User         *User          `gorm:"foreignKey:UserID;references:UserID"`
	Descriptions []*Description `gorm:"foreignKey:MemoryID;references:MemoryID"`
	Photos       []*Photo       `gorm:"foreignKey:MemoryID;references:MemoryID"`
}

func (m *Memory) BeforeCreate(tx *gorm.DB) error {
	if m.MemoryID == "" {
		m.MemoryID = uuid.New().String()
	}

	return nil
}

func (Memory) TableName() string {
	return "memories"
}
