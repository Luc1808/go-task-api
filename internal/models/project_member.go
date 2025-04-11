package models

// ProjectMember model (join table)
type ProjectMember struct {
	UserID    uint `gorm:"primaryKey"`
	ProjectID uint `gorm:"primaryKey"`
	Role      string

	// User    User    `gorm:"foreignKey:UserID"`
	// Project Project `gorm:"foreignKey:ProjectID"`
}
