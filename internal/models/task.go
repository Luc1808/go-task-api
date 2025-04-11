package models

import "time"

type Task struct {
	ID          string `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string
	Status      string
	DueDate     time.Time

	ProjectID uint
	// Project   Project

	// Comments on this task
	Comments []Comment `gorm:"foreignKey:TaskID"`
}

// CREATE TABLE tasks (
//     id VARCHAR PRIMARY KEY,
//     title VARCHAR NOT NULL,
//     description TEXT,
//     status VARCHAR,
//     due_date TIMESTAMP,
//     project_id INTEGER REFERENCES projects(id)
// );
