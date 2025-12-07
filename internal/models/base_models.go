package models

import "time"

type Question struct {
	Id        int64     `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`

	Answers []Answer `json:"answers" gorm:"foreignKey:QuestionId"`
}

func (Question) TableName() string {
	return "questions"
}

type Answer struct {
	Id         int64     `json:"id"`
	QuestionId int64     `json:"question_id"`
	UserId     string    `json:"user_id" gorm:"not null"`
	Text       string    `json:"text" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at"`
}

type CreateAnswerRequest struct {
	Text   string `json:"text"`
	UserId string `json:"user_id"`
}
