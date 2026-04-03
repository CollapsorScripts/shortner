package utils

import (
	"bytes"
	"math/rand"
	"shortner/pkg/logger"

	"github.com/goccy/go-json"
)

const alphabet = "QOS4rT08Dm7dZVOPwucfM2haFiNyEjBK3UtC9IqYlzv6XpWgWsAJebG5H1RxnLbK"

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

// EncodeBase62 - кодирует число в строку Base62
func EncodeBase62(num int64) string {
	if num == 0 {
		return "0"
	}

	base := int64(len(alphabet))
	result := ""

	for num > 0 {
		result = string(alphabet[num%base]) + result
		num /= base
	}

	return result
}

// GenerateRandomString - генерирует случайный набор символов (англ алфавит, case uppercase + цифры от 0 до 9)
func GenerateRandomString(length int) string {
	var result = make([]byte, length)
	for i := range length {
		result[i] = alphabet[rand.Intn(len(alphabet))]
	}

	return string(result)
}
