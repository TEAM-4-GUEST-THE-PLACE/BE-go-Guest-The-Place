package models

import (
	"time"
)

type Question struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Options string `json:"options"`
	Correct_option  string `json:"correct_option" db:"correct_option"`
	Image   string `json:"image"`

	Created_ct time.Time
	Updated_ct time.Time
}
