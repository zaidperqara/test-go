package user

import (
	"errors"
	"gorm.io/gorm"
	"os/user"
)

type UserService interface {
	Register(user *User) (*User, error)
	// ... (add other required service methods)
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) Register(user *User) (*User, error) {
	// Input validation (add more thorough validation as needed)
	if user.Name == "" || user.Email == "" || user.Password == "" {
		return nil, errors.New("all fields are required")
	}

	if err := s.repo.CreateUser(user); err != nil {
		return nil, err
	}

	user.Password = "" // Don't return the hashed password
	return user, nil
}

func (s *userService) Login(email string, password string) (*user.User, error) {
	// 1. Retrieve user by email
	existingUser, err := s.repo.GetUserByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid username or password")
		}
		return nil, err // Handle other database errors
	}

	// 2. Compare password
	if err := existingUser.ComparePassword(password); err != nil {
		return nil, errors.New("invalid username or password")
	}

	// 3. If successful, return the user (consider what data to include)
	return existingUser, nil
}
