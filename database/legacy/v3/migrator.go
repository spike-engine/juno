package v3

import (
	"github.com/jmoiron/sqlx"

	"github.com/spike-engine/juno/database"
	"github.com/spike-engine/juno/database/postgresql"
)

var _ database.Migrator = &Migrator{}

// Migrator represents the database migrator that should be used to migrate from v2 of the database to v3
type Migrator struct {
	Sql *sqlx.DB
}

func NewMigrator(db *postgresql.Database) *Migrator {
	return &Migrator{
		Sql: sqlx.NewDb(db.Sql, "postgres"),
	}
}
