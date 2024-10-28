package pg

import (
	"database/sql"

	"github.com/thinhhuy997/go-windows/internal/track/repository"
	"github.com/thinhhuy997/go-windows/pkg/log"
)

type implRepository struct {
	l  log.Logger
	db *sql.DB
}

func NewRepository(l log.Logger, db *sql.DB) repository.Repository {
	return &implRepository{
		l:  l,
		db: db,
	}
}
