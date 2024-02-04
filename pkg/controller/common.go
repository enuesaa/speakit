package controller

type EmptySchema struct{}

type ListSchema[T any] struct {
	Items []T `json:"items"`
}

type WithMetadata[T any] struct {
	Id       string `json:"id"`
	Data     T      `json:"data"`
	Created  string `json:"created"`
	Modified string `json:"modified"`
}

type IdSchema struct {
	Id string `json:"id"`
}

func createListResponse[T any]() ListSchema[WithMetadata[T]] {
	return ListSchema[WithMetadata[T]]{
		Items: make([]WithMetadata[T], 0),
	}
}