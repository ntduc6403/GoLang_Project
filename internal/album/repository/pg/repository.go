package pg

import (
	"context"

	"github.com/thinhhuy997/go-windows/internal/dbmodel"
	"github.com/thinhhuy997/go-windows/internal/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// All returns all todos
func (repo implRepository) List(ctx context.Context) ([]model.Album, error) {
	dbAlbums, err := dbmodel.Albums().All(ctx, repo.db)
	if err != nil {
		repo.l.Errorf(ctx, "album.repository.All.dbm.All: %s", err)
		return nil, err
	}

	var albums []model.Album
	for _, dbalbum := range dbAlbums {
		albums = append(albums, model.Album{
			ID: dbalbum.ID,
			Title: dbalbum.Title,
			Artist: dbalbum.Artist,
			Genre: dbalbum.Genre.String,
			ReleaseDate: dbalbum.ReleaseDate.Time,
			CreatedAt: dbalbum.CreatedAt.Time,
			UpdatedAt: dbalbum.UpdatedAt.Time,
		})
	}

	return albums, nil
}

// Create a new album
func (repo implRepository) Create(ctx context.Context, album model.Album) error {
	dbm := dbmodel.Album{
		Title: album.Title,
		Artist: album.Artist,
	}

	wl := boil.Whitelist(
		dbmodel.AlbumColumns.Title,
		dbmodel.AlbumColumns.Artist,
	)

	if err := dbm.Insert(ctx, repo.db, wl); err != nil {
		repo.l.Errorf(ctx, "album.repository.Create.dbm.Insert %+v: %s", dbm, err)
		return err
	}

	return nil
}

// Detail a album
func (repo implRepository) Detail(ctx context.Context, id int) (model.Album, error) {
	dbAlbum, err := dbmodel.Albums(dbmodel.AlbumWhere.ID.EQ(id)).One(ctx, repo.db)

	if err != nil {
		repo.l.Errorf(ctx, "todo.repository.Detail.dbm.Detail: %s", err)
		return model.Album{}, err
	}

	var album = model.Album{}
	album = model.Album{
		ID: dbAlbum.ID,
		Title: dbAlbum.Title,
		Artist: dbAlbum.Artist,
		Genre: dbAlbum.Genre.String,
		CreatedAt: dbAlbum.CreatedAt.Time,
		UpdatedAt: dbAlbum.UpdatedAt.Time,
	}

	return album, nil
}

