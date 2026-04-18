package schema

import (
	"DBDock/models"
	"database/sql"
	"fmt"
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

func (p *PostgresProvider) GetPrimaryKey(db *sql.DB, table string) (string, error) {
	query := `
		SELECT a.attname
		FROM pg_index i
		JOIN pg_attribute a
		  ON a.attrelid = i.indrelid
		 AND a.attnum = ANY(i.indkey)
		WHERE i.indrelid = $1::regclass
		AND i.indisprimary;
	`

	var key string
	err := db.QueryRow(query, table).Scan(&key)
	return key, err
}

func (p *PostgresProvider) GetTableData(db *sql.DB, table string, limit int, offset int, orderBy string, order models.Order) (models.Result[[]models.Column], error) {
	query := fmt.Sprintf(`
		SELECT *
		FROM public.%s
		ORDER BY %s %s
		LIMIT $1 OFFSET $2
	`, table, orderBy, order)

	rows, err := db.Query(query, limit, offset)
	if err != nil {
		return models.Result[[]models.Column]{}, err
	}
	defer rows.Close()

	var total int64

	countQuery := fmt.Sprintf(`SELECT COUNT(*) FROM public.%s`, table)
	err = db.QueryRow(countQuery).Scan(&total)
	if err != nil {
		return models.Result[[]models.Column]{}, err
	}

	cols, err := rows.Columns()
	if err != nil {
		return models.Result[[]models.Column]{}, err
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
			return models.Result[[]models.Column]{}, err
		}

		for i, val := range values {
			results[i].Rows = append(results[i].Rows, val)
		}
	}

	return models.Result[[]models.Column]{
		Result: results,
		Total:  total,
	}, nil
}
