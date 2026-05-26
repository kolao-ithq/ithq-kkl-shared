package entities

import "time"

type InsertOnly struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}
