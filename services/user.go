package services

import (
	"goGuestThePlace/models"
	"gorm.io/gorm"
)

type UserService interface {
	GetUser() (models.Users, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetUser() ([]models.Users, error) {
	var users []models.Users

	err := r.db.Find(&users).Error

	return users, err
}


