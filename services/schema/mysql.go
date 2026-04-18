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

func (m *MySQLProvider) GetPrimaryKey(db *sql.DB, table string) (string, error) {
	query := `
		SELECT COLUMN_NAME
		FROM INFORMATION_SCHEMA.KEY_COLUMN_USAGE
		WHERE TABLE_SCHEMA = DATABASE()
		AND TABLE_NAME = ?
		AND CONSTRAINT_NAME = 'PRIMARY'
		LIMIT 1;
	`

	var key string
	err := db.QueryRow(query, table).Scan(&key)
	return key, err
}

func (m *MySQLProvider) GetTableData(db *sql.DB, table string, limit int, offset int, orderBy string, order models.Order) ([]models.Column, error) {
	query := `
		SELECT *
		FROM ?
		ORDER BY ? ?
		LIMIT ? OFFSET ?
	`

	rows, err := db.Query(query, table, orderBy, order, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	results := make([]models.Column, len(cols))
	for i, c := range cols {
		results[i] = models.Column{
			Name: c,
			Rows: []interface{}{},
		}
	}

	for rows.Next() {
		values := make([]interface{}, len(cols))
		ptrs := make([]interface{}, len(cols))

		for i := range values {
			ptrs[i] = &values[i]
		}

		if err := rows.Scan(ptrs...); err != nil {
			return nil, err
		}

		for i, val := range values {
			results[i].Rows = append(results[i].Rows, val)
		}
	}

	return results, nil
}
