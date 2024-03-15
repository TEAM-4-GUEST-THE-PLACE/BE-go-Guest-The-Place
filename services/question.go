package services

import (
	"goGuestThePlace/models"

	"gorm.io/gorm"
)

type QuestionService interface {
	GetQuestion() ([]models.Question, error)
}

func RepositoryQuestion(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetQuestion() ([]models.Question, error) {
	var questions []models.Question

	// err := r.db.Find(&questions).Error
	r.db.Find(&questions)

	return questions, nil	
}
	