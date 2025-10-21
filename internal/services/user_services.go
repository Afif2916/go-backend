package services


import (
	"github.com/Afif2916/go-backend/internal/models"
	"github.com/Afif2916/go-backend/internal/repositories"
)

func GetAllUsers() ([]models.User, error) {
	return repositories.GetAllUsers()
}

func GetUserByID(id uint) (models.User, error) {
	return repositories.GetUserByID(id)
}

func CreateUser(user *models.User) error {
	return repositories.CreateUser(user)
}

func UpdateUser(user *models.User) error {
	return repositories.UpdateUser(user)
}

func DeleteUser(id uint) error {
	return repositories.DeleteUser(id)
}
