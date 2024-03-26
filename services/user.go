package services

import (
	"goGuestThePlace/models"

	"gorm.io/gorm"
)

type UserService interface {
	GetUser() ([]models.Users, error)
	GetUserById(ID int) (models.Users, error)
	GetUserByEmail (email string) (models.Users, error)
	CreatedUser(users models.Users) (models.Users,error)
	UpadateUser(users models.Users) (models.Users,error)
	DeleteUsers(users models.Users) (models.Users,error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetUser() ([]models.Users, error) {
	var users []models.Users

	err := r.db.Preload("Avatars").Find(&users).Error

	return users, err
}

func (r *repository) GetUserById(ID int) (models.Users, error) {
	var users models.Users

	err := r.db.First(&users, ID).Error

	return users, err
}

func (r *repository) GetUserByEmail(email string) (models.Users, error) {
	var users models.Users

	err := r.db.First(&users, "email = ?",email).Error

	return users, err
}

func (r *repository) CreatedUser(users models.Users) (models.Users, error) {
	err := r.db.Create(&users).Error

	return users, err
}


func (r *repository) UpadateUser(users models.Users) (models.Users, error) {
	err := r.db.Save(&users).Error

	return users, err
}

func (r *repository) DeleteUsers(users models.Users) (models.Users, error) {
	err := r.db.Delete(&users).Error

	return users, err
}
