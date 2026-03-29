package models

type DBConnection struct {
	Name     string         `json:"name"`
	Type     DatabaseDriver `json:"type"`
	Host     string         `json:"host"`
	Port     int            `json:"port"`
	Username string         `json:"username"`
	Password string         `json:"password"`
	Database string         `json:"database"`
}
