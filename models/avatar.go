package models

import (
	"time"
)


type Avatar struct {
	ID         	uint   		`json:"id" gorm:"primaryKey"`
	Image 		string 		`json:"image"`
	Users		[]Users		`gorm:"foreignKey:Avatars"`

	Created_at time.Time
	Updated_at time.Time
}