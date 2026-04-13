package models

type Table struct {
	Name   string  `json:"name"`
	Schema *string `json:"schema,omitempty"`
}
