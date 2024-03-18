package services

import (
	"goGuestThePlace/models"
	"gorm.io/gorm"
)


type AvatarService interface {
	GetAvatar() ([]models.Avatar, error)
}


func RepositoryAvatar(db *gorm.DB) *repository {
	return &repository{db}
}


func (r *repository) GetAvatar() ([]models.Avatar, error) {
	var avatars []models.Avatar

	err := r.db.Find(&avatars).Error

	return avatars, err
}