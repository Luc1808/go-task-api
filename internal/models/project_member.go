package models

// ProjectMember model (join table)
type ProjectMember struct {
	UserID    uint `gorm:"primaryKey"`
	ProjectID uint `gorm:"primaryKey"`
	Role      string

	// User    User    `gorm:"foreignKey:UserID"`
	// Project Project `gorm:"foreignKey:ProjectID"`
}

// CREATE TABLE project_members (
//     user_id INTEGER NOT NULL,
//     project_id INTEGER NOT NULL,
//     role VARCHAR,
//     PRIMARY KEY (user_id, project_id),
//     FOREIGN KEY (user_id) REFERENCES users(id),
//     FOREIGN KEY (project_id) REFERENCES projects(id)
// );
