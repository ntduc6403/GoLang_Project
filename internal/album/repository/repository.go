package repository

import (
	"context"

	"github.com/thinhhuy997/go-windows/internal/model"
)

// Repository is the interface for todo repository
//
//go:generate mockery --name=Repository
type Repository interface {
	List(ctx context.Context) ([]model.Album, error)
	Create(ctx context.Context, album model.Album) error
	Detail(ctx context.Context, id int) (model.Album, error)
	Update(ctx context.Context, album model.Album) error
	Delete(ctx context.Context, id int) error
}
