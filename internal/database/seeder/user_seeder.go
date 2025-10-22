package seeder

import (
	"fmt"
	"log"

	"github.com/Afif2916/go-backend/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserSeeder struct {
	db *gorm.DB
}

func NewUserSeeder(db *gorm.DB) *UserSeeder {
	return &UserSeeder{db: db}
}

func (s *UserSeeder) Run() error {
	// Get department IDs
	var departments []models.Department
	if err := s.db.Find(&departments).Error; err != nil {
		return fmt.Errorf("failed to get departments: %v", err)
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}

	users := []models.User{
		{
			Name:         "Admin User",
			Email:        "admin@company.com",
			PasswordHash: string(hashedPassword),
			Role:         models.RoleAdmin,
			DivisionID:   departments[0].DivisionID, // Technology division
			DepartmentID: departments[0].ID,         // Software Development
		},
		{
			Name:         "John Doe",
			Email:        "john.doe@company.com",
			PasswordHash: string(hashedPassword),
			Role:         models.RoleEmployee,
			DivisionID:   departments[0].DivisionID, // Technology division
			DepartmentID: departments[0].ID,         // Software Development
		},
		{
			Name:         "Jane Smith",
			Email:        "jane.smith@company.com",
			PasswordHash: string(hashedPassword),
			Role:         models.RoleEmployee,
			DivisionID:   departments[2].DivisionID, // HR division
			DepartmentID: departments[2].ID,         // Recruitment
		},
		{
			Name:         "Bob Johnson",
			Email:        "bob.johnson@company.com",
			PasswordHash: string(hashedPassword),
			Role:         models.RoleEmployee,
			DivisionID:   departments[4].DivisionID, // Finance division
			DepartmentID: departments[4].ID,         // Accounting
		},
	}

	for _, user := range users {
		var existing models.User
		if err := s.db.Where("email = ?", user.Email).First(&existing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := s.db.Create(&user).Error; err != nil {
					return fmt.Errorf("failed to create user %s: %v", user.Email, err)
				}
				log.Printf("Created user: %s (%s)", user.Name, user.Email)
			} else {
				return err
			}
		} else {
			log.Printf("User already exists: %s", user.Email)
		}
	}

	log.Println("User seeder completed successfully!")
	return nil
}
