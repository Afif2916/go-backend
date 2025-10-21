package main

import (
	"github.com/Afif2916/go-backend/internal/routes"
	"github.com/Afif2916/go-backend/internal/models"
	"github.com/Afif2916/go-backend/internal/config"
)

func main(){

	config.InitConfig()
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.User{})
	config.DB.AutoMigrate(&models.Product{})

	r:= routes.SetupRouter()
	r.Run(":8080")
}