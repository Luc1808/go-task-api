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

// CREATE TABLE comments (
//     id SERIAL PRIMARY KEY,
//     content TEXT NOT NULL,
//     created_at TIMESTAMP NOT NULL,
//     task_id VARCHAR REFERENCES tasks(id),
//     user_id INTEGER REFERENCES users(id)
// );
