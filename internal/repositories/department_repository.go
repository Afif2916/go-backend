package repositories

import (
	"github.com/Afif2916/go-backend/internal/config"
	"github.com/Afif2916/go-backend/internal/models"
	preload "github.com/Afif2916/go-backend/internal/repositories/preloads"
)

func GetAllDepartements() ([]models.Department, error) {
	var departments []models.Department

	result := preload.PreloadDepartmentRelations(config.DB).Find(&departments)

	return departments, result.Error
}

func GetAllDepartementById(id uint) ([]models.Department, error) {
	var department []models.Department
	result := preload.PreloadDepartmentRelations(config.DB).First(&department, id)

	return department, result.Error
}
