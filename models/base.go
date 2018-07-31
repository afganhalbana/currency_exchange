package models

import "time"

type Base struct {
	ID        int64      `gorm:"primary_key"`
	CreatedAt time.Time  `gorm:"not null;type:timestamp"`
	UpdatedAt time.Time  `gorm:"type:timestamp;DEFAULT:current_timestamp"`
	DeletedAt *time.Time `gorm:"index"`
}
