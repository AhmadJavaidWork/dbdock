package services

import (
	"DBDock/db"
	"DBDock/models"
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

func (cs *ConnectionService) Test(conn models.CreateDBConnection, driver models.DatabaseDriver) error {
	var dsn string

	switch driver.Name {
	case "postgres":
		dsn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=prefer", conn.Host, conn.Port, conn.Username, conn.Password, conn.DatabaseName)
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conn.Username, conn.Password, conn.Host, conn.Port, conn.DatabaseName)
	default:
		return fmt.Errorf("unsupported db")
	}

	db, err := sql.Open(driver.Name, dsn)
	if err != nil {
		return err
	}
	defer db.Close()
	db.SetConnMaxLifetime(5 * time.Second)
	return db.Ping()
}

func (cs *ConnectionService) Create(conn models.CreateDBConnection) (models.DBConnection, error) {
	query := `
		INSERT INTO connections (database_driver_id, name, host, port, username, password, database_name)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	result, err := db.DB.Exec(
		query,
		conn.DatabaseDriverID,
		conn.Name,
		conn.Host,
		conn.Port,
		conn.Username,
		conn.Password,
		conn.DatabaseName,
	)
	if err != nil {
		return models.DBConnection{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.DBConnection{}, err
	}

	query = getConnectionByIDQuery
	row := db.DB.QueryRow(query, id)
	var created models.DBConnection
	err = row.Scan(
		&created.ID,
		&created.DatabaseDriverID,
		&created.Name,
		&created.Host,
		&created.Port,
		&created.Username,
		&created.Password,
		&created.DatabaseName,
		&created.CreatedAt,
		&created.UpdatedAt,
		&created.DatabaseDriver.ID,
		&created.DatabaseDriver.Name,
		&created.DatabaseDriver.Label,
		&created.DatabaseDriver.DefaultPort,
		&created.DatabaseDriver.CreatedAt,
		&created.DatabaseDriver.UpdatedAt,
	)

	if err != nil {
		return models.DBConnection{}, err
	}

	return created, nil
}

func (cs *ConnectionService) GetAll() ([]models.DBConnection, error) {
	query := `
		SELECT
			c.id,
			c.database_driver_id,
			c.name,
			c.host,
			c.port,
			c.username,
			c.password,
			c.database_name,
			c.created_at,
			c.updated_at,
			d.id AS database_driver_id,
			d.name AS database_driver_name,
			d.label AS database_driver_label,
			d.default_port AS database_driver_default_port,
			d.created_at AS default_port_created_at,
			d.updated_at AS default_port_updated_at
		FROM connections c
		LEFT JOIN database_drivers d
		ON c.database_driver_id = d.id
		WHERE c.deleted_at IS NULL AND d.deleted_at IS NULL
	`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.DBConnection
	for rows.Next() {
		var c models.DBConnection
		err := rows.Scan(
			&c.ID,
			&c.DatabaseDriverID,
			&c.Name,
			&c.Host,
			&c.Port,
			&c.Username,
			&c.Password,
			&c.DatabaseName,
			&c.CreatedAt,
			&c.UpdatedAt,
			&c.DatabaseDriver.ID,
			&c.DatabaseDriver.Name,
			&c.DatabaseDriver.Label,
			&c.DatabaseDriver.DefaultPort,
			&c.DatabaseDriver.CreatedAt,
			&c.DatabaseDriver.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, c)
	}

	return result, nil
}

func (cs *ConnectionService) Update(conn models.DBConnection) (models.DBConnection, error) {
	query := `
		UPDATE connections SET
		database_driver_id = ?, 
		name = ?, 
		host = ?, 
		port = ? , 
		username = ?, 
		password = ?, 
		database_name = ? 
		WHERE id = ?
	`
	_, err := db.DB.Exec(
		query,
		conn.DatabaseDriverID,
		conn.Name,
		conn.Host,
		conn.Port,
		conn.Username,
		conn.Password,
		conn.DatabaseName,
		conn.ID,
	)

	if err != nil {
		return models.DBConnection{}, err
	}

	query = getConnectionByIDQuery

	row := db.DB.QueryRow(query, conn.ID)
	var updated models.DBConnection
	err = row.Scan(
		&updated.ID,
		&updated.DatabaseDriverID,
		&updated.Name,
		&updated.Host,
		&updated.Port,
		&updated.Username,
		&updated.Password,
		&updated.DatabaseName,
		&updated.CreatedAt,
		&updated.UpdatedAt,
		&updated.DatabaseDriver.ID,
		&updated.DatabaseDriver.Name,
		&updated.DatabaseDriver.Label,
		&updated.DatabaseDriver.DefaultPort,
		&updated.DatabaseDriver.CreatedAt,
		&updated.DatabaseDriver.UpdatedAt,
	)

	if err != nil {
		return models.DBConnection{}, err
	}

	return updated, nil
}

const getConnectionByIDQuery = `
	SELECT
		c.id,
		c.database_driver_id,
		c.name,
		c.host,
		c.port,
		c.username,
		c.password,
		c.database_name,
		c.created_at,
		c.updated_at,
		d.id AS database_driver_id,
		d.name AS database_driver_name,
		d.label AS database_driver_label,
		d.default_port AS database_driver_default_port,
		d.created_at AS default_port_created_at,
		d.updated_at AS default_port_updated_at
	FROM connections c
	LEFT JOIN database_drivers d
	ON c.database_driver_id = d.id
	WHERE c.deleted_at IS NULL AND d.deleted_at IS NULL AND c.id=?
`
