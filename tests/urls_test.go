package tests

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"shortner/internal/server"
	"shortner/tests/suite"
	"shortner/utils"
	"sync"
	"testing"

	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v3"
)

func Test_CreateUrl_Good(t *testing.T) {
	_, st := suite.New(t)

	endpoint := fmt.Sprintf("%s%s", st.BaseURL, "shorten")

	requestGoodBody := struct {
		Url string `json:"url"`
	}{
		Url: "https://example.com",
	}

	req, err := http.NewRequest(fiber.MethodPost, endpoint, bytes.NewBuffer(utils.ToBytesJSON(requestGoodBody)))
	if err != nil {
		t.Fatalf("Ошибка при создании запроса: %v", err)
		return
	}

	st.SetStandardHeaders(req)

	response, err := st.Client.Do(req)
	if err != nil {
		t.Errorf("Ошибка выполнения запроса: %v", err)
		return
	}
	defer response.Body.Close()

	respBytes, err := io.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Ошибка чтения тела ответа: %v", err)
		return
	}

	t.Logf("Статус: %s", response.Status)

	if response.StatusCode != http.StatusOK {
		httpError := new(server.HTTPError)
		if err := json.Unmarshal(respBytes, httpError); err != nil {
			t.Errorf("Ошибка при анмаршлинге тела ошибки: %v; body=%q", err, string(respBytes))
			return
		}
		t.Errorf("Ошибка: %s", utils.ToJSON(httpError))
		return
	}

	shortenResponse := struct {
		ShortUrl string `json:"short_url"`
	}{}

	if err := json.Unmarshal(respBytes, &shortenResponse); err != nil {
		t.Errorf("Ошибка при анмаршлинге тела ответа: %v; body=%q", err, string(respBytes))
		return
	}

	t.Logf("Ответ: %s", utils.ToJSON(shortenResponse))
}

func Test_CreateUrl_Bad(t *testing.T) {
	_, st := suite.New(t)

	endpoint := fmt.Sprintf("%s%s", st.BaseURL, "shorten")

	requestGoodBody := struct {
		Url string `json:"url"`
	}{
		Url: "12321321",
	}

	req, err := http.NewRequest(fiber.MethodPost, endpoint, bytes.NewBuffer(utils.ToBytesJSON(requestGoodBody)))
	if err != nil {
		t.Fatalf("Ошибка при создании запроса: %v", err)
		return
	}

	st.SetStandardHeaders(req)

	response, err := st.Client.Do(req)
	if err != nil {
		t.Errorf("Ошибка выполнения запроса: %v", err)
		return
	}
	defer response.Body.Close()

	respBytes, err := io.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Ошибка чтения тела ответа: %v", err)
		return
	}

	t.Logf("Статус: %s", response.Status)

	if response.StatusCode != http.StatusOK {
		httpError := new(server.HTTPError)
		if err := json.Unmarshal(respBytes, httpError); err != nil {
			t.Errorf("Ошибка при анмаршлинге тела ошибки: %v; body=%q", err, string(respBytes))
			return
		}
		t.Errorf("Ошибка: %s", utils.ToJSON(httpError))
		return
	}

	shortenResponse := struct {
		ShortUrl string `json:"short_url"`
	}{}

	if err := json.Unmarshal(respBytes, &shortenResponse); err != nil {
		t.Errorf("Ошибка при анмаршлинге тела ответа: %v; body=%q", err, string(respBytes))
		return
	}

	t.Logf("Ответ: %s", utils.ToJSON(shortenResponse))
}

func Test_Clicks(t *testing.T) {
	_, st := suite.New(t)
	endpoint := fmt.Sprintf("http://localhost:8010/1BbKte")

	req, err := http.NewRequest(fiber.MethodGet, endpoint, nil)
	if err != nil {
		t.Fatalf("Ошибка при создании запроса: %v", err)
		return
	}

	var wg sync.WaitGroup

	for i := 0; i < 48; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, _ = st.Client.Do(req)
			if err != nil {
				t.Errorf("Ошибка выполнения запроса: %v", err)
				return
			}
		}()
	}

	wg.Wait()
	t.Logf("Все запросы были выполнены")
}
