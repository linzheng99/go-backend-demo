package userService

import (
	userModel "github.com/linzheng99/go-backend-demo/features/user/model"
	userRepository "github.com/linzheng99/go-backend-demo/features/user/repository"
)

type UserService struct {
	userRepository *userRepository.UserRepo
}

// create a new user service
func NewUserService(userRepository *userRepository.UserRepo) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) CreateUser(user *userModel.User) error {
	return s.userRepository.CreateUser(user)
}

func (s *UserService) GetAllUsers() ([]userModel.User, error) {
	return s.userRepository.ListUsers()
}
