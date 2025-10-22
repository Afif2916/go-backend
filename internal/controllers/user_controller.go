package controllers

import (
	"net/http"
	"strconv"

	"github.com/Afif2916/go-backend/internal/config"
	"github.com/Afif2916/go-backend/internal/models"
	"github.com/Afif2916/go-backend/internal/services"
	"github.com/Afif2916/go-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users, err := services.GetAllUsers()
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, http.StatusOK, "Successfully retrieved users", users)
}

func GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := services.GetUserByID(uint(id))
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, http.StatusOK, "Successfully retrieved users", user)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := services.CreateUser(&user); err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	var createdUser models.User
	if err := config.DB.
		Preload("Division").
		Preload("Department").
		Preload("Department.Division").
		First(&createdUser, user.ID).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "User created but failed to load relations: "+err.Error())
		return
	}
	utils.Success(c, http.StatusCreated, "Successfully Created user", createdUser)
}

func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = uint(id)
	if err := services.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := services.DeleteUser(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
