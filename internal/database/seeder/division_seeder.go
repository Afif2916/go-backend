package seeder

import (
	"fmt"
	"log"

	"github.com/Afif2916/go-backend/internal/models"
	"gorm.io/gorm"
)

type DivisionSeeder struct {
	db *gorm.DB
}

func NewDivisionSeeder(db *gorm.DB) *DivisionSeeder {
	return &DivisionSeeder{db: db}
}

func (s *DivisionSeeder) Run() error {
	divisions := []models.Division{
		{Name: "Technology"},
		{Name: "Human Resources"},
		{Name: "Finance"},
		{Name: "Marketing"},
		{Name: "Operations"},
	}
	for _, division := range divisions {
		var existing models.Division
		if err := s.db.Where("name = ?", division.Name).First(&existing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := s.db.Create(&division).Error; err != nil {
					return fmt.Errorf("failed to create division %s: %v", division.Name, err)
				}
				log.Printf("Created division: %s", division.Name)
			} else {
				return err
			}
		} else {
			log.Printf("Division already exists: %s", division.Name)
		}
	}

	log.Println("Division seeder completed successfully!")
	return nil

}
