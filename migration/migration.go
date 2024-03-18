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
		
		
	)

	if err != nil {
		panic(err)
	}
}