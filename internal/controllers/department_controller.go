package controllers

import (
	"net/http"
	//"strconv"

	//"github.com/Afif2916/go-backend/internal/models"
	"github.com/Afif2916/go-backend/internal/services"
	"github.com/Afif2916/go-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

func GetAllDepartements(c *gin.Context) {
	departments, err := services.GetAllDepartement()
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, http.StatusOK, "Successfully retrieved Departments", departments)
}

// func GetDepartmentById(c *gin.Context) {
// 	idParam := c.Param("id")
// 	id, err := strconv.Atoi(idparam)
// }
