package migration


import (
	"goGuestThePlace/config"
	"goGuestThePlace/models"

)

func RunMigration()  {
	err := config.DB.AutoMigrate(
		&models.Question{},
		&models.Users{},
		&models.Diamond{},
		&models.Transaction{},
		
		
		
	)

	if err != nil {
		panic(err)
	}
}