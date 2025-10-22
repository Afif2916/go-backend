package repositories

import (
	"errors"

	"github.com/Afif2916/go-backend/internal/config"
	"github.com/Afif2916/go-backend/internal/models"
	preload "github.com/Afif2916/go-backend/internal/repositories/preloads"
	"github.com/Afif2916/go-backend/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User

	result := preload.PreloadUserRelations(config.DB).Find(&users)

	return users, result.Error
}

func GetUserByID(id uint) (models.User, error) {
	var user models.User
	result := preload.PreloadUserRelations(config.DB).First(&user, id)
	return user, result.Error
}

func CreateUser(user *models.User) error {
	if err := utils.ValidateUserFields(user.Name, user.Email, user.PasswordHash, user.DivisionID, user.DepartmentID); err != nil {
		return err
	}

	var existing models.User
	if err := config.DB.Where("email = ?", user.Email).First(&existing).Error; err == nil {
		return errors.New("email already in use")
	}

	var division models.Division
	if err := config.DB.First(&division, user.DivisionID).Error; err != nil {
		return errors.New("division tidak ditemukan")
	}
	var department models.Department
	if err := config.DB.First(&department, user.DepartmentID).Error; err != nil {
		return errors.New("department tidak ditemukan")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("gagal mengenkripsi password")
	}
	user.PasswordHash = string(hashedPassword)

	return config.DB.Create(user).Error
}

func UpdateUser(user *models.User) error {
	return config.DB.Save(user).Error
}

func DeleteUser(id uint) error {
	return config.DB.Delete(&models.User{}, id).Error
}
