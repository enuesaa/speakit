package controller

// common schemas
type EmptySchema struct {}
type ListSchema[T any] struct {
	Items []T `json:"items"`
}

type WithMetadata[T any] struct {
	Id string `json:"id"`
	Data T `json:"data"`
	Created string `json:"created"`
	Modified string `json:"modified"`
}
