package routes

import (
	"github.com/Afif2916/go-backend/internal/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	userGroup := r.Group("/api/users")
	{
		userGroup.GET("/", controllers.GetUsers)
		userGroup.POST("/", controllers.CreateUser)
		userGroup.GET("/:id", controllers.GetUser)
		userGroup.PUT("/:id", controllers.UpdateUser)
		userGroup.DELETE("/:id", controllers.DeleteUser)

	}

	announcementGroup := r.Group("/api/announcements")
	{
		announcementGroup.GET("/", controllers.GetAnnouncements)
	}

	divisionGroup := r.Group("/api/departments")
	{
		divisionGroup.GET("/", controllers.GetAllDepartements)
	}

	return r
}
