package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConfig()  {
	var err error
	dsn := "host=postgresql-167179-0.cloudclusters.net user=admin password=admin123 dbname=be-GTP port=19727 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed connect to database")
	}
}