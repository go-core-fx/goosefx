package goosefx

import (
	"database/sql"
	"fmt"
	"io/fs"

	"github.com/pressly/goose/v3"
)

type Storage fs.FS

func NewProvider(db *sql.DB, storage Storage, dialect goose.Dialect) (*goose.Provider, error) {
	provider, err := goose.NewProvider(dialect, db, storage)
	if err != nil {
		return nil, fmt.Errorf("init provider: %w", err)
	}

	return provider, nil
}
