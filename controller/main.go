package controller

// common schemas
type EmptySchema struct {}
type ListSchema[T any] struct {
	Items []T `json:"items"`
}
