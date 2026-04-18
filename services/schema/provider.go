// Package schema provides factory functions for creating schema providers.
package schema

import (
	"DBDock/models"
	"database/sql"
)

type SchemaProvider interface {
	ListTables(db *sql.DB) ([]models.Table, error)
	GetPrimaryKey(db *sql.DB, table string) (string, error)
	GetTableData(db *sql.DB, table string, limit int, offset int, orderBy string, order models.Order) ([]models.Column, error)
}
