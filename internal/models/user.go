package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"uniqueIndex;not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	// Projects owned by this user
	OwnedProjects []Project `gorm:"foreignKey:OwnerID"`
	// Comments created by this user
	Comments []Comment `gorm:"foreignKey:UserID"`
}
