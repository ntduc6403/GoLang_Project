package pg

import (
	"context"

	"github.com/thinhhuy997/go-windows/internal/dbmodel"
	"github.com/thinhhuy997/go-windows/internal/model"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (repo implRepository) All(ctx context.Context, queryParams) ([]model.Track, error) {
	// dbtracks, err := dbmodel.Tracks().All(ctx, repo.db)
	dbtracks, err := dbmodel.Tracks(qm.Where("id > ?", 16)).All(ctx, repo.db)
	if err != nil {
		repo.l.Errorf(ctx, "track.repository.All.dbm.All: %s", err)
		return nil, err
	}

	var tracks []model.Track
	for _, dbtrack := range dbtracks {
		tracks = append(tracks, model.Track{
			ID: dbtrack.ID,
			AlbumId: dbtrack.AlbumID,
			Title: dbtrack.Title,
			Duration: dbtrack.Duration.Int,
			TrackNumber: dbtrack.TrackNumber.Int,
			CreatedAt: dbtrack.CreatedAt.Time,
			UpdatedAt: dbtrack.UpdatedAt.Time,
		})
	}

	return tracks, nil
}

func (repo implRepository) Create(ctx context.Context, track model.Track) error {
	dbm := dbmodel.Track{
		AlbumID: track.AlbumId,
		Title: track.Title,
		Duration: null.IntFrom(track.Duration),
		TrackNumber: null.IntFrom(track.TrackNumber),
	}

	wl := boil.Whitelist(
		dbmodel.TrackColumns.AlbumID,
		dbmodel.TrackColumns.Title,
		dbmodel.TrackColumns.Duration,
		dbmodel.TrackColumns.TrackNumber,
	)

	if err := dbm.Insert(ctx, repo.db, wl); err != nil {
		repo.l.Errorf(ctx, "track.repository.Create.dbm.Insert %+v: %s", dbm, err)
		return err
	}

	return nil
} 


func (repo implRepository) Detail(ctx context.Context, id int) (model.Track, error) {
	dbtrack, err := dbmodel.Tracks(dbmodel.TrackWhere.ID.EQ(id)).One(ctx, repo.db)

	if err != nil {
		repo.l.Errorf(ctx, "track.repository.Detail.dbm.Detail: %s", err)
		return model.Track{}, err
	}

	var track = model.Track{}

	track = model.Track{
		ID: dbtrack.ID,
		AlbumId: dbtrack.AlbumID,
		Title: dbtrack.Title,
		Duration: dbtrack.Duration.Int,
		TrackNumber: dbtrack.TrackNumber.Int,
	}

	return track, nil
}
