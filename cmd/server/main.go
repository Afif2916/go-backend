package main

import (
	"log"
	"os"

	"github.com/Afif2916/go-backend/internal/config"
	"github.com/Afif2916/go-backend/internal/models"
	"github.com/Afif2916/go-backend/internal/routes"
)

func main() {
	config.InitConfig()
	config.ConnectDatabase()

	// Auto migrate
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

	// Run seeder jika environment development
	if os.Getenv("APP_ENV") == "development" {
		log.Println("Running in development mode, seeding database...")
		// Anda bisa panggil seeder di sini jika mau
		// atau biarkan manual melalui cmd/seeder
	}

	r := routes.SetupRouter()
	r.Run(":8080")
}
