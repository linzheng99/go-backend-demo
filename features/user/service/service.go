package user

import (
	user "github.com/linzheng99/go-backend-demo/features/user/model"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *user.User) error {
	return db.Create(user).Error
}

func GetAllUsers(db *gorm.DB) ([]user.User, error) {
	var users []user.User
	err := db.Find(&users).Error
	return users, err
}
