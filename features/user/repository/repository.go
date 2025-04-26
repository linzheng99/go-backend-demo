package userRepository

import (
	userModel "github.com/linzheng99/go-backend-demo/features/user/model"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

type UserRepository interface {
	CreateUser(user *userModel.User) error
	ListUsers() ([]userModel.User, error)
}

// create a new user repository
func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) CreateUser(user *userModel.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepo) ListUsers() ([]userModel.User, error) {
	var users []userModel.User
	err := r.db.Find(&users).Error
	return users, err
}
