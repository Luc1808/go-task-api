package models

type Project struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Description string

	// Owner relationship
	OwnerID uint
	Owner   User `gorm:"foreignKey:OwnerID"`

	// Members of the project (many-to-many relationship)
	Members []ProjectMember `gorm:"foreignKey:ProjectID"`

	// Tasks in the project
	Tasks []Task `gorm:"foreignKey:ProjectID"`
}

// CREATE TABLE projects (
//     id SERIAL PRIMARY KEY,
//     name VARCHAR NOT NULL,
//     description TEXT,
//     owner_id INTEGER REFERENCES users(id)
// );
