package models

type Column struct {
	Name string        `json:"name"`
	Rows []interface{} `json:"rows"`
}
