// Package services provides services like connection, database, theme etc.
package services

import (
	"DBDock/db"
	"DBDock/models"
	"database/sql"
	"fmt"
	"sync"
)

type ConnectionManager struct {
	mu          sync.Mutex
	connections map[int]*sql.DB
}

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		connections: make(map[int]*sql.DB),
	}
}

func (cm *ConnectionManager) Connect(conn models.DBConnection) error {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	if _, exists := cm.connections[conn.ID]; exists {
		dbConn := cm.connections[conn.ID]
		err := dbConn.Close()
		if err != nil {
			return err
		}
	}

	var driver, dsn string

	switch conn.DatabaseDriver.Name {
	case "postgres":
		driver = "postgres"
		dsn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=prefer", conn.Host, conn.Port, conn.Username, conn.Password, conn.DatabaseName)
	case "mysql":
		driver = "mysql"
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conn.Username, conn.Password, conn.Host, conn.Port, conn.DatabaseName)
	default:
		return fmt.Errorf("unsupported driver")
	}

	dbConn, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}

	if err := dbConn.Ping(); err != nil {
		return err
	}

	query := `
	UPDATE connections SET last_used_at=DATETIME('now'), is_connected=TRUE where id=?
	`

	_, err = db.DB.Exec(query, conn.ID)
	if err != nil {
		return err
	}

	cm.connections[conn.ID] = dbConn
	return nil
}

func (cm *ConnectionManager) Disconnect(id int) error {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	if _, exists := cm.connections[id]; !exists {
		return nil
	}

	dbConn := cm.connections[id]

	query := `
		UPDATE connections SET last_used_at=DATETIME('now'), is_connected=FALSE where id=?
	`

	db.DB.Exec(query, id)
	return dbConn.Close()

}

func (cm *ConnectionManager) DisconnectAll() {
	for id, dbConn := range cm.connections {
		query := `
			UPDATE connections SET last_used_at=DATETIME('now'), is_connected=FALSE where id=?
		`

		fmt.Println("closing connection with id=", id)

		db.DB.Exec(query, id)
		dbConn.Close()
	}
}
