package services

import (
	"goGuestThePlace/models"
	"gorm.io/gorm"
)

type UserService interface {
	GetUser(ID int) (models.Users, error)
	CreateUser(user models.Users) (models.Users, error)
	UpdateUser(user models.Users) (models.Users, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetUser(ID int) (models.Users, error) {
	var users models.Users

	err := r.db.Find(&users).Error

	return users, err

}

func (r *repository) CreateUser(user models.Users) (models.Users, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *repository) UpdateUser(user models.Users) (models.Users, error) {
	err := r.db.Save(&user).Error
	return user, err	
}
