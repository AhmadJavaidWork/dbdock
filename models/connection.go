package models

import "time"

type DBConnectionBase struct {
	Name             string `json:"name"`
	DatabaseDriverID int    `json:"databaseDriverId"`
	Host             string `json:"host"`
	Port             int    `json:"port"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	DatabaseName     string `json:"databaseName"`
}

type CreateDBConnection struct {
	DBConnectionBase
}

type DBConnectionEntity struct {
	DBConnectionBase
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type DBConnection struct {
	DBConnectionEntity
	DatabaseDriver DatabaseDriver `json:"databaseDriver"`
}

type SaveDBConnectionResponse struct {
	Connection DBConnection `json:"connection"`
	Message    string       `json:"message"`
}
