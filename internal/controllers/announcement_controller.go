package controllers

import (
	"net/http"
	//"strconv"

	//"github.com/Afif2916/go-backend/internal/models"
	"github.com/Afif2916/go-backend/internal/services"
	"github.com/Afif2916/go-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

func GetAnnouncements(c *gin.Context) {
	announcements, err := services.GetAllAnnouncements()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	utils.Success(c, http.StatusOK, "Announcements retrieved successfully", announcements)
}
