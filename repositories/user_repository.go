// repositories/user_repository.go
package repositories

import (
    "fmt"
    "project/models"
)

// شبیه‌سازی پایگاه‌داده با یک نقشه
var userDB = map[string]models.User{}

// افزودن کاربر جدید
func CreateUser(user models.User) error {
	if _, exists := userDB[user.Username]; exists {
		return fmt.Errorf("کاربر قبلاً ثبت شده است")
	}
	userDB[user.Username] = user
	return nil
}

// دریافت کاربر با نام کاربری
func GetUserByUsername(username string) (models.User, error) {
	user, exists := userDB[username]
	if !exists {
		return models.User{}, fmt.Errorf("کاربر پیدا نشد")
	}
	return user, nil
}
