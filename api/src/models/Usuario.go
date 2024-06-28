package models

import "time"

type User struct {
	Id        uint64    `json:"id,omitempty"`
	Nome      string    `json:"nome,omitempty"`
	Nickname  string    `json:"nickname,omitempty"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
