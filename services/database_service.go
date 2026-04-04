package services

import (
	"DBDock/db"
	"DBDock/models"
)

type DatabaseService struct{}

func NewDatabaseService() *DatabaseService {
	return &DatabaseService{}
}

func (ds *DatabaseService) ListSupported() ([]models.DatabaseDriver, error) {
	query := `
		SELECT
			id, name, label, default_port, created_at, updated_at
		FROM database_drivers
	`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.DatabaseDriver

	for rows.Next() {
		var c models.DatabaseDriver
		err := rows.Scan(&c.ID, &c.Name, &c.Label, &c.DefaultPort, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, c)
	}

	return result, nil
}
