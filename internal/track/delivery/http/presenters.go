package http

import (
	"strings"
	"time"

	"github.com/thinhhuy997/go-windows/internal/model"
	trackUseCase "github.com/thinhhuy997/go-windows/internal/track/usecase"
	pkgErrors "github.com/thinhhuy997/go-windows/pkg/errors"
	"github.com/thinhhuy997/go-windows/pkg/response"
)

// type addRequest struct {
// 	Name        string `json:"name"`
// 	Description string `json:"description"`
// }

// func (r addRequest) toInput() (albumUseCase.CreateInput, error) {
// 	vErrCollect := pkgErrors.NewValidationErrorCollector()

// 	if strings.TrimSpace(r.Name) == "" {
// 		vErrCollect.Add(pkgErrors.NewValidationError("name", errMsgNameIsRequired))
// 	}

// 	if vErrCollect.HasError() {
// 		return albumUseCase.CreateInput{}, vErrCollect
// 	}

// 	return albumUseCase.CreateInput{
// 		Title:        r.Title,
// 		Description: r.Description,
// 	}, nil
// }

type todoResponse struct {
	ID          int               `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	CreatedAt   response.DateTime `json:"created_at"`
}

type addRequest struct {
	AlbumID     int    `json:"album_id"`
	Title       string `json:"title"`
	Duration    int    `json:"duration"`
	TrackNumber int    `json:"track_number"`
}

type filterParams struct {
	AlbumID int `json:"album_id"`
}

func (r addRequest) toInput() (trackUseCase.CreateInput, error) {
	vErrCollect := pkgErrors.NewValidationErrorCollector()

	if strings.TrimSpace(r.Title) == "" {
		vErrCollect.Add(pkgErrors.NewValidationError("name", errMsgNameIsRequired))
	}

	if vErrCollect.HasError() {
		return trackUseCase.CreateInput{}, vErrCollect
	}

	// return trackUseCase.CreateInput{
	// 	Title:        r.Title,
	// 	Artist: r.Artist,
	// }, nil

	return trackUseCase.CreateInput{
		AlbumID: r.AlbumID,
		Title: r.Title,
		Duration: r.Duration,
		TrackNumber: r.TrackNumber,
	}, nil
}

type albumResponse struct {
	ID	int `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Genre string `json:"genre"`
	ReleaseDate response.DateTime `json:"release_date"`
}

type trackResponse struct {
	ID int `json:"id"`
	AlbumId int `json: "album_id"`
	Title string `json: "title"`
	Duration int `json: "duration"`
	TrackNumber int `json: "track_number"`
	Created_at time.Time `json: "created_at"`
	UpdatedAt time.Time `json: "updated_at"`
}

func newListResponse(albums []model.Track) []trackResponse {
	data := []trackResponse{}
	for _, track := range albums {
		data = append(data, trackResponse{
			ID: track.ID,
			AlbumId: track.AlbumId,
			Title: track.Title,
			Duration: track.Duration,
			TrackNumber: track.TrackNumber,
			Created_at: track.CreatedAt,
			UpdatedAt: track.UpdatedAt,
		})
	}

	return data
}

func newDetailResponse(track model.Track) trackResponse {
	data := trackResponse{
		ID: track.ID,
		AlbumId: track.AlbumId,
		Title: track.Title,
		Duration: track.Duration,
		TrackNumber: track.TrackNumber,
		Created_at: track.CreatedAt,
		UpdatedAt: track.UpdatedAt,
	}

	return data
}
