package utils

import (
	"bytes"
	"shortner/pkg/logger"

	"github.com/goccy/go-json"
)

// jsonPrettyPrint - форматирует JSON строку с отступами
func jsonPrettyPrint(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "   ")
	if err != nil {
		return in
	}
	return out.String()
}

// ToBytesJSON - конвертирует объект в JSON байт
func ToBytesJSON[T any](object T) []byte {
	jsonByte, err := json.Marshal(object)
	if err != nil {
		logger.Error("Ошибка при получении JSON: ", err.Error())
	}
	n := len(jsonByte)
	result := string(jsonByte[:n])

	return []byte(jsonPrettyPrint(result))
}

// ToJSON - конвертирует объект в JSON строку
func ToJSON[T any](object T) string {
	jsonByte, err := json.Marshal(object)
	if err != nil {
		logger.Error("Ошибка при получении JSON: %v", err)
	}
	n := len(jsonByte)
	result := string(jsonByte[:n])

	return jsonPrettyPrint(result)
}
