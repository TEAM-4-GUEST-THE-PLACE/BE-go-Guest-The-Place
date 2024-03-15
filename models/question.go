package models

import (
	"time"
)

type Question struct {
	ID         	uint   `json:"id" gorm:"primaryKey"`
	Question   	string `json:"question"`
	Options 	string `json:"options"`
	Answer		string `json:"answer" db:"answer"`
	Image      	string `json:"image"`

	Created_ct time.Time
	Updated_ct time.Time
}
