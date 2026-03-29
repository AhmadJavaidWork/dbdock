package services

import "DBDock/models"

type DatabaseService struct{}

func NewDatabaseService() *DatabaseService {
	return &DatabaseService{}
}

func (ds *DatabaseService) ListSupported() []models.DatabaseDriver {
	return []models.DatabaseDriver{
		{
			ID:          "postgres",
			Label:       "PostgreSQL",
			DefaultPort: 5432,
		},
		{
			ID:          "mysql",
			Label:       "MySQL",
			DefaultPort: 3306,
		},
	}
}
