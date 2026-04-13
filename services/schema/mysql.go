package schema

import (
	"DBDock/models"
	"database/sql"
)

type MySQLProvider struct{}

func (m *MySQLProvider) ListTables(db *sql.DB) ([]models.Table, error) {
	query := `
		SELECT table_name
		FROM information_schema.tables
		WHERE table_schema = DATABASE()
		  AND table_type = 'BASE TABLE'
		ORDER BY table_name;
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []models.Table
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		tables = append(tables, models.Table{
			Name:   name,
			Schema: nil,
		})
	}

	return tables, nil
}
