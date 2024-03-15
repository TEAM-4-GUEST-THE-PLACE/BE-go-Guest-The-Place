package migration


import (
	"goGuestThePlace/config"
	"goGuestThePlace/models"

)

func RunMigration()  {
	err := config.DB.AutoMigrate(
		&models.Question{},
		&models.User{},
		
	)

	if err != nil {
		panic(err)
	}
}