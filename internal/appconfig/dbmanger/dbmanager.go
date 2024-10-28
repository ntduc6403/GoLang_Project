package dbmanger

import (
	"database/sql"
	"fmt"

	"github.com/thinhhuy997/go-windows/config"
	appPG "github.com/thinhhuy997/go-windows/internal/appconfig/pg"
)

const (
	dbDriverPG = "postgres"
)

// DBConfig is the database configuration
type DBConfig struct {
	driver string
	pgDB   *sql.DB
}

// NewDBConfig returns a new DBConfig
func NewDBConfig(cfg *config.Config) (DBConfig, error) {
	switch cfg.Database.Driver {
	case dbDriverPG:
		db, err := appPG.Connect(cfg.PG.URL)
		if err != nil {
			return DBConfig{}, err
		}

		return newPGDBConfig(db), nil
	default:
		return DBConfig{}, fmt.Errorf("database driver %s is not supported", cfg.Database.Driver)
	}
}

func newPGDBConfig(db *sql.DB) DBConfig {
	return DBConfig{
		driver: dbDriverPG,
		pgDB:   db,
	}
}

// IsPG returns true if the database driver is PostgreSQL
func (c DBConfig) IsPG() bool {
	return c.driver == dbDriverPG
}

// GetPG returns the PostgreSQL database
func (c DBConfig) GetPG() *sql.DB {
	return c.pgDB
}
