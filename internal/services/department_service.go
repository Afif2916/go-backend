package services

import (
	"github.com/Afif2916/go-backend/internal/models"
	"github.com/Afif2916/go-backend/internal/repositories"
)

func GetAllDepartement() ([]models.Department, error) {
	return repositories.GetAllDepartements()
}

func GetDepartementById(id uint) ([]models.Department, error) {
	return repositories.GetAllDepartementById(id)
}
