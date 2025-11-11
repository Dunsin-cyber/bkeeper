package services

import (
	"errors"
	"fmt"

	"github.com/Dunsin-cyber/bkeeper/cmd/api/requests"
	"github.com/Dunsin-cyber/bkeeper/common"
	"github.com/Dunsin-cyber/bkeeper/internal/models"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserSrvice(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (userService UserService) RegisterUser(userRequest requests.RegisterUserRequest) (*models.UserModel, error) {
	hashedPassword, err := common.HashPassword(userRequest.Password)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return nil, errors.New("user registration failed")
	}

	createdUser := models.UserModel{
		FirstName: &userRequest.FirstName,
		LastName:  &userRequest.LastName,
		Email:     userRequest.Email,
		Password:  hashedPassword,
	}

	result := userService.db.Create(&createdUser)
	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, errors.New("user Registration failed")
	}

	return &createdUser, nil

}

func (userService UserService) GetUserByEmail(email string) (*models.UserModel, error) {
	var user models.UserModel
	result := userService.db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (userService UserService) ChangeAuthUserPassword(user *models.UserModel, newPassword string) error {

	//5 hash the new password and update the password filed in the database
	hashedNewPassword, err := common.HashPassword(newPassword)
	if err != nil {
		return errors.New("password change failed")

	}

	if err := userService.db.Where("id = ?", user.ID).First(&user).Error; err != nil {
		return err
	}

	user.Password = hashedNewPassword

	if err := userService.db.Save(&user).Error; err != nil {
		return errors.New("password change failed")
	}

	return nil

}
