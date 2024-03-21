package services

import (
	"goGuestThePlace/models"
	"gorm.io/gorm"
)

type TransactionService interface {
	GetTransaction() ([]models.Transaction, error)
	GetTransactionById(TransactionId int) (models.Transaction, error)
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	UpadateTransaction(status string, orderId int) (models.Transaction, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetTransaction() ([]models.Transaction, error) {
	var transactions []models.Transaction

	err := r.db.Find(&transactions).Error

	return transactions, err

}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	var transactions models.Transaction

	err := r.db.Create(&transaction).Error

	return transactions, err
}
