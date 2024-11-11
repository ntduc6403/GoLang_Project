package model

import "time"

type Album struct {
	ID          int       `json:"id"`
	Title       string    `json: "title"`
	Artist      string    `json: "artist"`
	Genre       string    `json: "genre"`
	ReleaseDate time.Time `json:"release_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
