
// Table tags {
//   id integer [primary key, increment]
//   tag_id uuid [unique, not null, default: 'uuid_generate_v4()']
//   name varchar(255) [unique, not null]
//   created_at timestamp [default: 'now()', not null]
//   deleted_at timestamp [null]
// }

// Table memory_tags {
//   memory_id uuid [ref: > memories.memory_id, not null]
//   tag_id uuid [ref: > tags.tag_id, not null]
//   created_at timestamp [default: 'now()', not null]
//   deleted_at timestamp [null]

//   indexes {
//     (memory_id, tag_id) [unique]
//   }
// }

package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tag struct {
	ID        int       `gorm:"primaryKey"`
	TagID     string    `gorm:"type:uuid;unique;not null;default:uuid_generate_v4()"`
	Name      string    `gorm:"type:varchar(255);unique;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt

	Memories []Memory `gorm:"many2many:memory_tags;foreignKey:TagID;joinForeignKey:TagID;References:MemoryID;joinReferences:MemoryID"`
}

func (t *Tag) BeforeCreate(tx *gorm.DB) error {
	if t.TagID == "" {
		t.TagID = uuid.New().String()
	}

	return nil
}

func (Tag) TableName() string {
	return "tags"
}
