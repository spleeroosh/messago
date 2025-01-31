package entity

import "time"

//go:generate go run github.com/mailru/easyjson/easyjson

//easyjson:json
type Message struct {
	ID        int       `json:"id"`
	Type      string    `json:"type"`    // тип словаря (страны, должности и т.п.)
	Sender    string    `json:"sender"`  // значение элемента 	// язык ('ru', 'en')
	Content   string    `json:"content"` // название элемента
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
