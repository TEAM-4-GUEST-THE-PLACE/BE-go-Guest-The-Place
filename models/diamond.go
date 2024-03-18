package models

import (
	"time"
)

type Diamond struct {
	ID 			uint 	`json:"id" gorm:"primaryKey"`
	Amount 		string 	`json:"amount"`
	Image 		string 	`json:"image"`
	Price 		string 	`json:"price"`
	
	
	Created_ct time.Time
	Updated_ct time.Time

}
