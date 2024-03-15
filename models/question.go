package models

import (
	"time"
)

type Question struct {
	ID         uint     `json:"id" gorm:"primaryKey"`
	Question   string   `json:"question"`
	Answer     string   `json:"answer"`
	AnswerTrue string   `json:"answerTrue"`
	Image      string   `json:"image"`

	Created_ct time.Time
	Updated_ct time.Time
}
