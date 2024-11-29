// services/auth_service.go
package services

import (
	"fmt"
	"test/models"
	"test/repositories"
	"test/utils"
)

// ثبت‌نام کاربر جدید
func RegisterUser(username, password string) error {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	user := models.User{
		Username: username,
		Password: hashedPassword,
	}
	return repositories.CreateUser(user)
}

// ورود کاربر
func LoginUser(username, password string) (bool, error) {
	user, err := repositories.GetUserByUsername(username)
	if err != nil {
		return false, err
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return false, fmt.Errorf("رمز عبور اشتباه است")
	}
	return true, nil
}
