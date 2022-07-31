package repository

import (
	"log"

	"gitlab.com/mr687/privy-be-test-go/entity"
	"gorm.io/gorm"
)

// Define contract what's this repository can do
type UserRepository interface {
	Insert() (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	FindByID(id uint64) (*entity.User, error)
	IsDuplicateEmail() bool
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (repo *userRepository) Insert(user *entity.User) (*entity.User, error) {
	// Doing something before insert
	err := user.BeforeInsert()
	if err != nil {
		log.Fatal(err)
	}

	err = repo.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repo *userRepository) Save(user *entity.User) error {
	err := user.BeforeInsert()
	if err != nil {
		log.Fatal(err)
	}

	err = repo.db.Save(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *userRepository) FindByID(id uint64) (*entity.User, error) {
	var user *entity.User
	err := repo.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repo *userRepository) FindByEmail(email string) (*entity.User, error) {
	var user *entity.User
	err := repo.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repo *userRepository) IsDuplicateEmail(email string) bool {
	var user *entity.User
	err := repo.db.Where("email = ?", email).First(&user).Error
	return err == nil
}
