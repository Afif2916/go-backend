package seeder

import (
	"fmt"
	"log"

	"github.com/Afif2916/go-backend/internal/models"
	"gorm.io/gorm"
)

type DepartmentSeeder struct {
	db *gorm.DB
}

func NewDepartmentSeeder(db *gorm.DB) *DepartmentSeeder {
	return &DepartmentSeeder{db: db}
}

func (s *DepartmentSeeder) Run() error {
	// Get division IDs first
	var divisions []models.Division
	if err := s.db.Find(&divisions).Error; err != nil {
		return fmt.Errorf("failed to get divisions: %v", err)
	}

	divisionMap := make(map[string]uint)
	for _, div := range divisions {
		divisionMap[div.Name] = div.ID
	}

	departments := []models.Department{
		{DepartmentName: "Software Development", DivisionID: divisionMap["Technology"]},
		{DepartmentName: "IT Infrastructure", DivisionID: divisionMap["Technology"]},
		{DepartmentName: "Recruitment", DivisionID: divisionMap["Human Resources"]},
		{DepartmentName: "Training & Development", DivisionID: divisionMap["Human Resources"]},
		{DepartmentName: "Accounting", DivisionID: divisionMap["Finance"]},
		{DepartmentName: "Tax", DivisionID: divisionMap["Finance"]},
		{DepartmentName: "Digital Marketing", DivisionID: divisionMap["Marketing"]},
		{DepartmentName: "Brand Management", DivisionID: divisionMap["Marketing"]},
		{DepartmentName: "Supply Chain", DivisionID: divisionMap["Operations"]},
		{DepartmentName: "Quality Control", DivisionID: divisionMap["Operations"]},
	}

	for _, dept := range departments {
		var existing models.Department
		if err := s.db.Where("department_name = ? AND division_id = ?", dept.DepartmentName, dept.DivisionID).First(&existing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := s.db.Create(&dept).Error; err != nil {
					return fmt.Errorf("failed to create department %s: %v", dept.DepartmentName, err)
				}
				log.Printf("Created department: %s", dept.DepartmentName)
			} else {
				return err
			}
		} else {
			log.Printf("Department already exists: %s", dept.DepartmentName)
		}
	}

	log.Println("Department seeder completed successfully!")
	return nil
}
