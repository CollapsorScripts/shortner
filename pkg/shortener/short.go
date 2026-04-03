package shortener

import (
	"fmt"
	"net/url"
	"shortner/internal/config"
	"shortner/utils"
)

func checkUrl(input string) bool {
	u, err := url.ParseRequestURI(input)
	if err != nil {
		return false
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}

	if u.Host == "" {
		return false
	}

	return true
}

func GenerateShortUrl(cfg *config.Config, url string) (string, error) {
	if !checkUrl(url) {
		return "", fmt.Errorf("некорректная ссылка")
	}

	return utils.GenerateRandomString(cfg.ShortUrlLength), nil
}
