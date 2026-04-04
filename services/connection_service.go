package services

import (
	"DBDock/db"
	"DBDock/models"
	"DBDock/utils"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type ConnectionService struct{}

func NewConnectionService() *ConnectionService {
	return &ConnectionService{}
}

func (cs *ConnectionService) Test(conn models.DBConnection) error {
	var dsn string
	var driver string

	switch conn.Type.Name {
	case "postgres":
		driver = "postgres"
		dsn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=prefer", conn.Host, conn.Port, conn.Username, conn.Password, conn.Database)
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

func (cs *ConnectionService) Create(conn models.DBConnection) error {
	query := `
		INSERT INTO connections (database_driver_id, name, host, port, username, password, database)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	hashedPassword, err := utils.HashPassword(conn.Password)
	if err != nil {
		return err
	}
	_, err = db.DB.Exec(
		query,
		conn.Type.ID,
		conn.Name,
		conn.Host,
		conn.Port,
		conn.Username,
		hashedPassword,
		conn.Database,
	)
	return err
}
