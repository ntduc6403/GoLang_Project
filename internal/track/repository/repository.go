package repository

import (
	"context"

	"github.com/thinhhuy997/go-windows/internal/model"
)

// Repository is the interface for todo repository
//
//go:generate mockery --name=Repository
type Repository interface {
	// // Create creates a new todo
	// // Create(ctx context.Context, todo model.Todo) error
	// // All returns all todos
	// List(ctx context.Context) ([]model.Album, error)
	// Create(ctx context.Context, album model.Album) error
	// // Delete deletes a todo
	// // Delete(ctx context.Context, id int) error

	// // Detail a todo
	// Detail(ctx context.Context, id int) (model.Album, error)

	All(ctx context.Context) ([]model.Track, error)
	Create(ctx context.Context, track model.Track) error
	Detail(ctx context.Context, id int) (model.Track, error)
	Update(ctx context.Context, track model.Track) error
	Delete(ctx context.Context, id int) error
} 
