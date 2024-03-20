package services

import (
	"goGuestThePlace/models"
	"gorm.io/gorm"
)

type UserService interface {
	GetUser() ([]models.Users, error)
	UpadateUser(ID int, users models.Users) (models.Users,error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetUser() ([]models.Users, error) {
	var users []models.Users

	err := r.db.Find(&users).Error

	return users, err
}

func (r *repository) UpadateUser(ID int, users models.Users) (models.Users, error) {
	var user models.Users

	err := r.db.Save(&user).Error

	return user, err
}
