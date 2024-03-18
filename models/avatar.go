package models

import (
	"time"
)


type Avatar struct {
	ID         	uint   `json:"id" gorm:"primaryKey"`
	Image 		string `json:"image"`


	Created_ct time.Time
	Updated_ct time.Time
}