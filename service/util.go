package service

import (
	"encoding/json"

	"github.com/google/uuid"
)

func createId() string {
	uid, _ := uuid.NewUUID()
	return uid.String()
}

func toJson(value interface{}) string {
	b, _ := json.Marshal(value)
	return string(b)
}

func parseJson[T any](value string) T {
	var t T
	json.Unmarshal([]byte(value), &t)
	return t
}