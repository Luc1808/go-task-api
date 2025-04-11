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
