package server

import (
	"errors"
	"shortner/pkg/logger"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type statsResponse struct {
	OriginalURL    string `json:"original_url"`
	Clicks         int64  `json:"clicks"`
	CreatedAt      string `json:"created_at"`
	LastAccessedAt string `json:"last_accessed_at"`
}

func formatTimestamp(ts pgtype.Timestamptz) string {
	if !ts.Valid {
		return ""
	}

	return ts.Time.Format("02/01/2006 15:04")
}

func (r *Router) GetStats(c fiber.Ctx) error {
	shortURL := c.Params("short_url")

	if shortURL == "" {
		logger.Error("Указана пустая сокращенная ссылка")
		return SetHTTPError(c, "Сокращенная ссылка не может быть пустой", fiber.StatusBadRequest)
	}

	urlModel, err := r.services.URLS.GetUrlByShortUrl(c.RequestCtx(), shortURL)
	if err != nil {
		logger.Error("Ошибка при получении ссылки: %v", err)
		if errors.Is(err, pgx.ErrNoRows) {
			return SetHTTPError(c, "Сокращенная ссылка не найдена", fiber.StatusNotFound)
		}
		return SetHTTPError(c, "Ошибка при получении оригинальной ссылки", fiber.StatusInternalServerError)
	}

	stats, err := r.services.Statistics.GetStatisticsByUrlId(c.RequestCtx(), urlModel.ID)
	if err != nil {
		logger.Error("Ошибка при поиске статистики: %v", err)
		if errors.Is(err, pgx.ErrNoRows) {
			return SetHTTPError(c, "Статистика не найдена", fiber.StatusNotFound)
		}
		return SetHTTPError(c, "Ошибка при получении статистики", fiber.StatusInternalServerError)
	}

	response := statsResponse{
		OriginalURL:    urlModel.OriginalUrl,
		Clicks:         stats.Clicks,
		CreatedAt:      formatTimestamp(stats.CreatedAt),
		LastAccessedAt: formatTimestamp(stats.LastAccessed),
	}

	return c.JSON(response)
}
