package models

import (
	"time"
//	"gorm.io/gorm"
)

type Announcement struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"type:varchar(255);not null" json:"title"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	CreatedBy uint      `json:"created_by"`
	User      User      `gorm:"foreignKey:CreatedBy"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
