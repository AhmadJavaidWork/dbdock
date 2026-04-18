package models

type Order string

const (
	ASC  Order = "ASC"
	DESC Order = "DESC"
)

type Result[T any] struct {
	Result T     `json:"result"`
	Total  int64 `json:"total"`
}
