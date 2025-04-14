package repository

import (
	"fiber-auth-app/config"
	"fiber-auth-app/models"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmail(email string) (*models.User, error)
	FindById(id int) (*models.User, error)
	Create(user *models.User) error
}

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository() UserRepository {
	if config.DB == nil {
		log.Fatal("Database connection is nil")
	}
	return &userRepository{
		db: config.DB, // DB is your GORM database instance
	}
}

// FindByEmail finds a user by email
func (r *userRepository) FindById(id int) (*models.User, error) {
	var user models.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}
	return &user, nil
}
func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}
	return &user, nil
}

// Create creates a new user in the database
func (r *userRepository) Create(user *models.User) error {
	if r.db == nil {
		log.Fatal("Database connection is nil")
	}
	if user == nil {
		return fmt.Errorf("user is nil")
	}
	err := r.db.Create(user).Error
	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}
	return nil
}
