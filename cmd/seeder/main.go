package main

import (
	"flag"
	"log"

	"github.com/Afif2916/go-backend/internal/config"
	"github.com/Afif2916/go-backend/internal/database/seeder"
	"github.com/Afif2916/go-backend/internal/models"
)

func main() {
	clear := flag.Bool("clear", false, "Clear all data before seeding")
	flag.Parse()

	// Initialize config and database
	config.InitConfig()
	config.ConnectDatabase()

	// Run migrations first
	err := config.DB.AutoMigrate(
		&models.User{},
		&models.Announcement{},
		&models.Attendance{},
		&models.Leave{},
		&models.AttendanceLocation{},
		&models.Division{},
		&models.Department{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Run seeder dengan opsi
	if err := seeder.RunWithOptions(*clear); err != nil {
		log.Fatal("Failed to run seeders:", err)
	}

	log.Println("Database seeding completed successfully!")
}
