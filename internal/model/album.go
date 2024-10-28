package model

import "time"

// type test struct {
// 	ID          int       `json:"id"`
// 	Name        string    `json:"name"`
// 	Description string    `json:"description"`
// 	CreatedAt   time.Time `json:"created_at"`
// 	UpdatedAt   time.Time `json:"updated_at"`
// }

type Album struct {
	ID	int `json: "id"`
	Title string `json: "title"`
	Artist string `json: "artist"`
	Genre string `json: "genre"`
	ReleaseDate time.Time `json:"release_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

