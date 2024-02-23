package controller

type IdSchema struct {
	Id string `json:"id"`
}
type EmptySchema struct{}

type WithMetadata[T any] struct {
	Id       string `json:"id"`
	Data     T      `json:"data"`
	Created  string `json:"created"`
	Modified string `json:"modified"`
}
