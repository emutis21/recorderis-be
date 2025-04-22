package models

import (
	"time"

	"gorm.io/gorm"
)

type MemoryTag struct {
	MemoryID string    `gorm:"type:uuid;not null;primaryKey"`
	TagID    string    `gorm:"type:uuid;not null;primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt
}

func (MemoryTag) TableName() string {
	return "memory_tags"
}
