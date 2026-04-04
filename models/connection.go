package models

import "time"

type DBConnection struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	Type      DatabaseDriver `json:"type"`
	Host      string         `json:"host"`
	Port      int            `json:"port"`
	Username  string         `json:"username"`
	Password  string         `json:"password"`
	Database  string         `json:"database"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"UpdatedAt"`
}
