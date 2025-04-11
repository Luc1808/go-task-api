package models

import "time"

type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	Content   string `gorm:"not null"`
	CreatedAt time.Time

	// Task relationship
	TaskID uint
	// Task   Task `gorm:"foreignKey:TaskID"`

	UserID uint
	// User   User `gorm:"foreignKey:UserID"`
}
