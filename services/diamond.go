package services


import (
	"goGuestThePlace/models"
	"gorm.io/gorm"
)


type DiamondService interface {
	GetDiamond() ([]models.Diamond, error)
}


func RepositoryDiamond(db *gorm.DB) *repository {
	return &repository{db}
}


func (r *repository) GetDiamond() ([]models.Diamond, error) {
	var diamonds []models.Diamond

	err := r.db.Find(&diamonds).Error

	return diamonds, err

}