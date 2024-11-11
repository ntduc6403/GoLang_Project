package usecase

import (
	"context"
	"time"

	"github.com/thinhhuy997/go-windows/internal/model"
)

type CreateInput struct {
	AlbumID 		int
	Title			string
	Duration		int
	TrackNumber 	int
}


// func (uc implUsecase) List(ctx context.Context) ([]model.Album, error) {
// 	albums, err := uc.repo.List(ctx)
// 	if err != nil {
// 		uc.l.Errorf(ctx, "todo.usecase.All.repo.All(ctx): %s", err)
// 		return nil, err
// 	}

// 	return albums, nil
// }

func (uc implUsecase) All(ctx context.Context) ([]model.Track, error) {
	tracks, err := uc.repo.All(ctx)
	if err != nil {
		uc.l.Errorf(ctx, "track.usecase.All.repo.All(ctx): %s", err)
		return nil, err
	}

	return tracks, nil
}

func (uc implUsecase) Create(ctx context.Context, input CreateInput) error {
	now := time.Now()
	m := model.Track{
		AlbumId: input.AlbumID,
		Title: input.Title,
		Duration: input.Duration,
		TrackNumber: input.TrackNumber,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := uc.repo.Create(ctx, m); err != nil {
		uc.l.Errorf(ctx, "track.usecase.Create.repo.Create(ctx, %+v): %s", m, err)
		return err
	}

	return nil
}

func (uc implUsecase) Detail(ctx context.Context, id int) (model.Track, error) {
	track, err := uc.repo.Detail(ctx, id)

	if err != nil {
		uc.l.Errorf(ctx, "todo.usecase.Detail.repo.Detail(ctx, %d): %", id, err)
		return model.Track{}, err
	}

	return track, nil
}

func (uc implUsecase) Update(ctx context.Context, track model.Track) error {
    return uc.repo.Update(ctx, track) // Gọi phương thức Update từ repository
}
func (uc implUsecase) Delete(ctx context.Context, id int) error {
    return uc.repo.Delete(ctx, id) // Gọi phương thức Delete từ repository
}