package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Description struct {
	ID            int       `gorm:"primaryKey"`
	DescriptionID string    `gorm:"type:uuid;unique;not null;default:uuid_generate_v4()"`
	MemoryID      string    `gorm:"type:uuid;not null"`
	Index         int       `gorm:"not null"`
	Text          string    `gorm:"type:text;not null"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
	DeletedAt     gorm.DeletedAt
	Version       int `gorm:"default:1;not null"`

	Memory *Memory `gorm:"foreignKey:MemoryID;references:MemoryID"`
}

func (d *Description) BeforeCreate(tx *gorm.DB) error {
	if d.DescriptionID == "" {
		d.DescriptionID = uuid.New().String()
	}

	return nil
}

func (Description) TableName() string {
	return "descriptions"
}
