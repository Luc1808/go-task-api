package models

import "time"

type TokenPair struct {
	AcessToken   string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshToken struct {
	ID        int       `gorm:"primaryKey"`
	Token     string    `gorm:"type:text;not null;index"`
	UserID    uint      `gorm:"not null;index"`
	User      User      `gorm:"foreignKey:UserID"`
	ExpiresAt time.Time `gorm:"not null"`
	CreatedAt time.Time
}
