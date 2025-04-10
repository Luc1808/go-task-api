package models

import "time"

type User struct {
	ID          uint   `gorm:"primaryKey"`
	Username    string `gorm:"uniqueIndex;not null"`
	Password    string `gorm:"not null"`
	CreatedAt   time.Time
	Projects    []Project
	Memberships []ProjectMember
	Comments    []Comment
}
