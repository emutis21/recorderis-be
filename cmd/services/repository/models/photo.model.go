package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Photo struct {
	ID        int       `gorm:"primaryKey"`
	PhotoID   string    `gorm:"type:uuid;unique;not null;default:uuid_generate_v4()"`
	MemoryID  string    `gorm:"type:uuid;not null"`
	Index     int       `gorm:"not null"`
	FileName  string    `gorm:"type:varchar(255);not null"`
	URL       string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt

	Memory *Memory `gorm:"foreignKey:MemoryID;references:MemoryID"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) error {
	if p.PhotoID == "" {
		p.PhotoID = uuid.New().String()
	}

	return nil
}

func (Photo) TableName() string {
	return "photos"
}
