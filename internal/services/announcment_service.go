package services

import (
	 "github.com/Afif2916/go-backend/internal/models"
	"github.com/Afif2916/go-backend/internal/repositories"
)

// func CreateAnnouncment(announcement *models.Announcement) error {
// 	return repositories.CreateAnnouncement(announcements)
// }

func GetAllAnnouncements() ([]models.Announcement, error) {
	return repositories.GetAllAnnouncements()
}

// func GetAnnouncmenetByID(id uint) (*models.Announcement, error) {
// 	return repositories.GetAnnouncementByID(id)
// }



