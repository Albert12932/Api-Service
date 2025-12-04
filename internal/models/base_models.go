package models

import "time"

type Question struct {
	Id        int64     `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

type Answer struct {
	Id         int64     `json:"id"`
	QuestionId int64     `json:"question_id"`
	UserId     int64     `json:"user_id"`
	Text       string    `json:"text"`
	CreatedAt  time.Time `json:"created_at"`
}
