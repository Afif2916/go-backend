package repositories

import (
	"github.com/Afif2916/go-backend/internal/config"
	"github.com/Afif2916/go-backend/internal/models"
	"gorm.io/gorm"
)

func GetAllAnnouncements() ([]models.Announcement, error) {
	var announcements []models.Announcement

	result := config.DB.
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name, email, role, created_at, updated_at")
		}).
		Order("created_at DESC").
		Find(&announcements)

	if result.Error != nil {
		return nil, result.Error
	}

	return announcements, result.Error
}

func GetAnnouncementByID(id uint) (models.Announcement, error) {
	var announcement models.Announcement

	result := config.DB.
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name, email, role, created_at, updated_at")
		}).
		First(&announcement, id)

	if result.Error != nil {
		return announcement, result.Error
	}

	return announcement, result.Error
}
