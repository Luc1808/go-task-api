package models

type Project struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Description string
	OwnerID     uint
	Owner       User `gorm:"foreignKey:OwnerID"`
	Members     []ProjectMember
	Tasks       []Task
}
