package services

import (
	"goGuestThePlace/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

// GetTransactionById implements TransactionService.
func (*repository) GetTransactionById(TransactionId int) (models.Transaction, error) {
	panic("unimplemented")
}

// UpadateTransaction implements TransactionService.
func (*repository) UpadateTransaction(status string, orderId int) (models.Transaction, error) {
	panic("unimplemented")
}

// DeleteUser implements UserService.
func (*repository) DeleteUser(users models.Users) (models.Users, error) {
	panic("unimplemented")
}
