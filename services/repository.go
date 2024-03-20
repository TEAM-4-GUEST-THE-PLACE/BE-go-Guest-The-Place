package services

import (
	"goGuestThePlace/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

// DeleteUser implements UserService.
func (*repository) DeleteUser(users models.Users) (models.Users, error) {
	panic("unimplemented")
}
