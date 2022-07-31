package repository

import (
	"gitlab.com/mr687/privy-be-test-go/entity"
	"gorm.io/gorm"
)

// Define contract what's this repository can do
type UserRepository interface {
	InsertUser(user *entity.User) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (repo *userRepository) InsertUser(newUser *entity.User) (*entity.User, error) {
	if err := repo.db.Create(newUser).Error; err != nil {
		return nil, err
	}
	return newUser, nil
}

func (repo *userRepository) FindByEmail(email string) (*entity.User, error) {
	var user *entity.User
	if err := repo.db.Where("email = ?", email).Take(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
