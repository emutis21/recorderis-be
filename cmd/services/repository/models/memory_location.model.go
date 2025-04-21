package models

import (
	"time"

	"gorm.io/gorm"
)

type MemoryLocation struct {
	MemoryID   string    `gorm:"type:uuid;not null;primaryKey"`
	LocationID string    `gorm:"type:uuid;not null;primaryKey"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	DeletedAt  gorm.DeletedAt
}

func (MemoryLocation) TableName() string {
	return "memory_locations"
}
