package user

import (
	"gorm.io/gorm"
	"os/user"
)

type UserRepository interface {
	CreateUser(user *User) error
	// ... (add other required repository methods)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(user *User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetUserByEmail(email string) (*user.User, error) {
	var userCol user.User
	result := r.db.Where("email = ?", email).First(&userCol)
	return &userCol, result.Error
}
