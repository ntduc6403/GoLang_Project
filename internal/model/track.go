package model

import "time"

type Track struct {
	ID	int `json: "id"`
	AlbumId int `json: "album_id"`
	Title string `json: "title"`
	Duration int `json: "duration"`
	TrackNumber int `json: "track_number"`
	CreatedAt time.Time `json: "created_at"`
	UpdatedAt time.Time `json: "updated_at"`
}
