package usecase

import (
	"context"

	"github.com/thinhhuy997/go-windows/internal/album/repository"
	"github.com/thinhhuy997/go-windows/internal/model"
	"github.com/thinhhuy997/go-windows/pkg/log"
)

// Usecase is the interface for todo usecase
//
//go:generate mockery --name=Usecase
type Usecase interface {
	// Create creates a new todo
	// Create(ctx context.Context, int CreateInput) error
	// All returns all todos
	List(ctx context.Context) ([]model.Album, error)
	Create(ctx context.Context, int CreateInput) error
	Detail(ctx context.Context, id int) (model.Album, error)

	Update(ctx context.Context, album model.Album) error
	

}

type implUsecase struct {
	l    log.Logger
	repo repository.Repository
}

func New(l log.Logger, repo repository.Repository) Usecase {
	return &implUsecase{
		l:    l,
		repo: repo,
	}
}
