// Package schema provides factory functions for creating schema providers.
package schema

import (
	"DBDock/models"
	"database/sql"
)

type SchemaProvider interface {
	ListTables(db *sql.DB) ([]models.Table, error)
}
