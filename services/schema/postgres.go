package schema

import (
	"DBDock/models"
	"database/sql"
)

type PostgresProvider struct{}

func (p *PostgresProvider) ListTables(db *sql.DB) ([]models.Table, error) {
	query := `
		SELECT table_schema, table_name
		FROM information_schema.tables
		WHERE table_type = 'BASE TABLE'
			AND table_schema NOT IN ('pg_catalog', 'information_schema')
		ORDER BY table_schema, table_name;
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []models.Table
	for rows.Next() {
		var schema string
		var name string
		if err := rows.Scan(&schema, &name); err != nil {
			return nil, err
		}
		tables = append(tables, models.Table{
			Name:   name,
			Schema: &schema,
		})
	}

	return tables, nil
}
