package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Afif2916/go-backend/internal/controllers"
)

func SetupRouter() *gin.Engine {
	r:= gin.Default()

	
	userGroup := r.Group("/api/users")
	{
		userGroup.GET("/", controllers.GetUsers)
		userGroup.POST("/", controllers.CreateUser)
		userGroup.GET("/:id", controllers.GetUser)
		userGroup.PUT("/:id", controllers.UpdateUser)
		userGroup.DELETE("/:id", controllers.DeleteUser)
	}

	return r
}