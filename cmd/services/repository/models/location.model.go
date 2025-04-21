package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Location struct {
	ID         int       `gorm:"primaryKey"`
	LocationID string    `gorm:"type:uuid;unique;not null;default:uuid_generate_v4()"`
	Location   string    `gorm:"type:varchar(255);not null"`
	Longitude  float64   `gorm:"type:decimal(10,7);not null"`
	Latitude   float64   `gorm:"type:decimal(10,7);not null"`
	City       string    `gorm:"type:varchar(255);not null"`
	Country    string    `gorm:"type:varchar(255);not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	DeletedAt  gorm.DeletedAt

	Memories []Memory `gorm:"many2many:memory_locations;foreignKey:LocationID;joinForeignKey:LocationID;References:MemoryID;joinReferences:MemoryID"`
}

func (l *Location) BeforeCreate(tx *gorm.DB) error {
	if l.LocationID == "" {
		l.LocationID = uuid.New().String()
	}

	return nil
}

func (Location) TableName() string {
	return "locations"
}
