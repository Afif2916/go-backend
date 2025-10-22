package models

import (
	//	"gorm.io/gorm"
	"time"
)

type UserRole string

const (
	RoleAdmin    UserRole = "ADMIN"
	RoleEmployee UserRole = "EMPLOYEE"
)

type User struct {
	ID           uint     `gorm:"primaryKey" json:"id"`
	Name         string   `gorm:"type:varchar(100);not null" json:"name"`
	Email        string   `gorm:"uniqueIndex;not null" json:"email"`
	PasswordHash string   `gorm:"not null" json:"password"`
	Role         UserRole `gorm:"type:enum('ADMIN','EMPLOYEE');default:'EMPLOYEE'" json:"role"`

	// Relasi dengan Division dan Department
	DivisionID uint     `gorm:"not null;index" json:"division_id"`
	Division   Division `gorm:"foreignKey:DivisionID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"division"`

	DepartmentID uint       `gorm:"not null;index" json:"department_id"`
	Department   Department `gorm:"foreignKey:DepartmentID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"department"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Attendances   []Attendance   `gorm:"foreignKey:UserID"`
	Leaves        []Leave        `gorm:"foreignKey:UserID"`
	Announcements []Announcement `gorm:"foreignKey:CreatedBy"`
}
