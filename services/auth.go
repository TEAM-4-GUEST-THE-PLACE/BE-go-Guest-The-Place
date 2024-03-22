package services

import (
	"goGuestThePlace/models"
	"gorm.io/gorm"
)

type AuthService interface {
	Login(fullname string) (models.Users, error)
}

func RepositoryAuth(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Login(fullname string) (models.Users, error) {
	var users models.Users
	err := r.db.First(&users, "fullname = ?",fullname ).Error
	return users, err
}
