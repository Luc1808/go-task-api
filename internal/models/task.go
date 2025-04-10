package models

import "time"

type Task struct {
	ID          string `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string
	Status      string
	DueDate     time.Time

	ProjectID uint
	Project   Project

	Comments []Comment
}
