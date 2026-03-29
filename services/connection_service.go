package services

import (
	"DBDock/models"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" // ⚠ blank import to register driver
	_ "github.com/lib/pq"
)

type ConnectionService struct{}

func NewConnectionService() *ConnectionService {
	return &ConnectionService{}
}

func (cs *ConnectionService) Test(conn models.DBConnection) error {
	var dsn string
	var driver string

	switch conn.Type.ID {
	case "postgres":
		driver = "postgres"
		dsn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", conn.Host, conn.Port, conn.Username, conn.Password, conn.Database)
	case "mysql":
		driver = "mysql"
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conn.Username, conn.Password, conn.Host, conn.Port, conn.Database)
	default:
		return fmt.Errorf("unsupported db")
	}

	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer db.Close()
	db.SetConnMaxLifetime(5 * time.Second)
	return db.Ping()
}
