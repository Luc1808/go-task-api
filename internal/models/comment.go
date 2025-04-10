package models

import "time"

type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	Content   string `gorm:"not null"`
	CreatedAt time.Time

	TaskID uint
	Task   Task

	UserID uint
	User   User
}
