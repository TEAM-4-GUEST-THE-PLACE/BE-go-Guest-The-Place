package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConfig()  {
	var err error
	dsn := "host=roundhouse.proxy.rlwy.net user=postgres password=kqaUZaPvstRpAiHSUDeBbVxVPtkgpEhv dbname=railway port=32190 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed connect to database")
	}
}