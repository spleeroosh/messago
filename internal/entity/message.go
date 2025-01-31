package entity

import "time"

//go:generate go run github.com/mailru/easyjson/easyjson

//easyjson:json
type Message struct {
	ID        int       `json:"id"`
	Sender    string    `json:"sender"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
