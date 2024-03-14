package models

import (
	"time"
)

type Question struct {
	ID          	 uint   `json:"id" gorm:"primaryKey"`
	question      	 string `json:"question"`
	answer       	 string `json:"answer"`
	answerTrue  	 string `json:"answerTrue"`
	answerFalse 	 string `json:"answerFalse"`
	image    		 string `json:"image"`

	created_ct  time.Time
	updated_ct  time.Time
	
}
