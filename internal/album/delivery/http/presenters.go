package http

import (
	"strings"

	albumUseCase "github.com/thinhhuy997/go-windows/internal/album/usecase"
	"github.com/thinhhuy997/go-windows/internal/model"
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
	Title string `json:"title"`
	Artist string `json:"artist"`
	Genre string `json:"genre"`
}

func (r addRequest) toInput() (albumUseCase.CreateInput, error) {
	vErrCollect := pkgErrors.NewValidationErrorCollector()

	if strings.TrimSpace(r.Title) == "" {
		vErrCollect.Add(pkgErrors.NewValidationError("name", errMsgNameIsRequired))
	}

	if vErrCollect.HasError() {
		return albumUseCase.CreateInput{}, vErrCollect
	}

	return albumUseCase.CreateInput{
		Title:        r.Title,
		Artist: r.Artist,
	}, nil
}

type albumResponse struct {
	ID	int `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Genre string `json:"genre"`
	ReleaseDate response.DateTime `json:"release_date"`
}

func newListResponse(albums []model.Album) []albumResponse {
	data := []albumResponse{}
	for _, album := range albums {
		data = append(data, albumResponse{
			ID:          album.ID,
			Title:        album.Title,
			Artist: album.Artist,
			Genre: album.Genre,
			ReleaseDate:   response.DateTime(album.ReleaseDate),
		})
	}

	return data
}

func newDetailResponse(album model.Album) albumResponse {
	data := albumResponse{
		ID: album.ID,
		Title: album.Title,
		Artist: album.Artist,
		Genre: album.Genre,
		ReleaseDate: response.DateTime(album.ReleaseDate),
	}

	return data
}
