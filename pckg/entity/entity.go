package entity

import "time"

type Question struct {
	ID          string    `json:"id"`
	Content     string    `json:"content"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UserCreated string    `json:"user_created"`
	UpdatedAt   time.Time `json:"updated_at"`
	UserUpdated string    `json:"user_updated"`
}

type Answer struct {
	ID          string    `json:"id"`
	QuestionId  string    `json:"question_id"`
	Content     string    `json:"content"`
	CreatedAt   time.Time `json:"created_at"`
	UserCreated string    `json:"user_created"`
}
