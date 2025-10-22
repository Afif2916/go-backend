package seeder

import (
	"fmt"
	"log"

	"github.com/Afif2916/go-backend/internal/config"
	"gorm.io/gorm"
)

type Seeder interface {
	Run() error
}

func RunAll(db *gorm.DB) error {
	seeders := []Seeder{
		NewDivisionSeeder(db),
		NewDepartmentSeeder(db),
		NewUserSeeder(db),
	}

	for _, seeder := range seeders {
		if err := seeder.Run(); err != nil {
			return fmt.Errorf("seeder error: %v", err)
		}
	}

	log.Println("All seeders completed successfully!")
	return nil
}

// ClearAll untuk menghapus semua data
func ClearAll(db *gorm.DB) error {
	// Hapus dalam urutan yang benar untuk menghindari constraint violation
	tables := []string{"users", "leaves", "attendances", "announcements", "attendance_locations", "departments", "divisions"}

	for _, table := range tables {
		if err := db.Exec("DELETE FROM " + table).Error; err != nil {
			log.Printf("Warning: Failed to delete from %s: %v", table, err)
		}
		if err := db.Exec("ALTER TABLE " + table + " AUTO_INCREMENT = 1").Error; err != nil {
			log.Printf("Warning: Failed to reset auto increment for %s: %v", table, err)
		}
	}

	log.Println("All tables cleared successfully!")
	return nil
}

// Run dengan opsi clear
func RunWithOptions(clear bool) error {
	db := config.DB

	if clear {
		if err := ClearAll(db); err != nil {
			return err
		}
	}

	return RunAll(db)
}
