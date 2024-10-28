package usecase

import (
	"context"
	"time"

	"github.com/thinhhuy997/go-windows/internal/model"
)

type CreateInput struct {
	Title	string
	Artist	string
	genre string
}


func (uc implUsecase) List(ctx context.Context) ([]model.Album, error) {
	albums, err := uc.repo.List(ctx)
	if err != nil {
		uc.l.Errorf(ctx, "todo.usecase.All.repo.All(ctx): %s", err)
		return nil, err
	}

	return albums, nil
}

// Create a new album
func (uc implUsecase) Create(ctx context.Context, input CreateInput) error {
	now := time.Now()
	m := model.Album{
		Title: input.Title,
		Artist: input.Artist,
		Genre: input.genre,
		ReleaseDate: now,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := uc.repo.Create(ctx, m); err != nil {
		uc.l.Errorf(ctx, "todo.usecase.Create.repo.Create(ctx, %+v): %s", m, err)
	}

	return nil
}


// Detail a album
func (uc implUsecase) Detail(ctx context.Context, id int) (model.Album, error) {
	// Call to repository
	album, err := uc.repo.Detail(ctx, id)

	if err != nil {
		uc.l.Errorf(ctx, "album.usecase.Detail.repo.Detail(ctx, %d): %s", id, err)
		return model.Album{}, nil
	}

	return album, nil
}