package service

import (
	"encoding/json"
	"strings"

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

func renameKeyToId(key string) string {
	list := strings.Split(key, ":")
	if len(list) == 0 {
		return ""
	}
	return list[1]
}
