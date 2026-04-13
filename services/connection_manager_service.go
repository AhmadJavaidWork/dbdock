// Package services provides services like connection, database, theme etc.
package services

import (
	"DBDock/db"
	"DBDock/models"
	"DBDock/services/schema"
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
		UPDATE connections SET last_used_at=DATETIME('now') where id=?
	`

	_, err = db.DB.Exec(query, conn.ID)
	if err != nil {
		return err
	}

	cm.connections[conn.ID] = dbConn
	return nil
}

func (cm *ConnectionManager) IsConnected(connectionID int) bool {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	_, exists := cm.connections[connectionID]
	return exists
}

func (cm *ConnectionManager) Disconnect(connectionID int) error {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	if _, exists := cm.connections[connectionID]; !exists {
		return nil
	}

	dbConn := cm.connections[connectionID]
	delete(cm.connections, connectionID)

	query := `
		UPDATE connections SET last_used_at=DATETIME('now') where id=?
	`

	db.DB.Exec(query, connectionID)
	return dbConn.Close()

}

func (cm *ConnectionManager) DisconnectAll() {
	for connectionID, dbConn := range cm.connections {
		query := `
			UPDATE connections SET last_used_at=DATETIME('now') where id=?
		`

		fmt.Println("closing connection with id=", connectionID)

		db.DB.Exec(query, connectionID)
		dbConn.Close()
	}
}

func (cm *ConnectionManager) Query(connectionID int, query string) ([]map[string]interface{}, error) {
	cm.mu.Lock()
	dbConn, exists := cm.connections[connectionID]
	cm.mu.Unlock()

	if !exists {
		return nil, fmt.Errorf("connection not found")
	}

	rows, err := dbConn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cols, _ := rows.Columns()
	results := []map[string]interface{}{}

	for rows.Next() {
		values := make([]interface{}, len(cols))
		ptrs := make([]interface{}, len(cols))
		for i := range values {
			ptrs[i] = &values[i]
		}

		rows.Scan(ptrs...)

		rowMap := map[string]interface{}{}
		for i, col := range cols {
			rowMap[col] = values[i]
		}

		results = append(results, rowMap)
	}

	return results, nil
}

func (cm *ConnectionManager) ListTables(connectionID int, driver string) ([]models.Table, error) {
	cm.mu.Lock()
	dbConn, exists := cm.connections[connectionID]
	cm.mu.Unlock()

	if !exists {
		return nil, fmt.Errorf("connection not found")
	}

	provider, err := schema.GetSchemaProvider(driver)
	if err != nil {
		return nil, err
	}

	return provider.ListTables(dbConn)
}
