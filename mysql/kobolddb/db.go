package kobolddb

import (
	"database/sql"
)

// KoboldDB is a wrapper struct to allow us to add additional functionality
// to the mysql database struct that handles mysql database connections.
type KoboldDB struct {
	*sql.DB
}

// New takes in an *sql.DB that has already been initialized and uses it for
// the form manipulation functions
func New(db *sql.DB) *KoboldDB {
	return &KoboldDB{db}
}
