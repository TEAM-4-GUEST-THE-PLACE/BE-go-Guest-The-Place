package services


import (
	"goGuestThePlace/models"
	"gorm.io/gorm"

)

type UserService interface {
	GetUser() ([]models.User, error)

}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetUser() ([]models.User, error) {
	var users []models.User


err := r.db.Find(&users).Error

return users, err

}