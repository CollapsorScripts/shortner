package utils

import "testing"

func TestEncodeBase62(t *testing.T) {
	url := "https://example.com"

	short_url := GenerateRandomString(6)

	t.Logf("Результат кодирования: \nИсходные данные: %s\nВыходные данные: %s", url, short_url)
}
