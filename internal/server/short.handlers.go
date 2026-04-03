package server

import (
	"errors"
	"shortner/internal/database/service"
	db "shortner/internal/database/sqlc"
	"shortner/pkg/logger"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5"
)

type shortenRequest struct {
	URL string `json:"url"`
}

type shortenResponse struct {
	ShortURL string `json:"short_url"`
}

// Shorten - сокращает URL
func (r *Router) Shorten(c fiber.Ctx) error {
	var req shortenRequest
	if err := c.Bind().JSON(&req); err != nil {
		logger.Error("Ошибка при анмаршлинге json: %v", err)
		return SetHTTPError(c, "Указаны неверные аргументы", fiber.StatusBadRequest)
	}

	createdURL, err := r.services.URLS.CreateUrl(c.RequestCtx(), r.cfg, req.URL)
	if err != nil {
		logger.Error("Ошибка при сокращении url: %v", err)
		return SetHTTPError(c, "Ошибка при сокращении URL", fiber.StatusInternalServerError)
	}

	return c.JSON(shortenResponse{
		ShortURL: createdURL.ShortUrl,
	})
}

func (r *Router) statisticsCollection(c fiber.Ctx, urlModel *db.Url) {
	stats, err := r.services.Statistics.IncrementClicksCountByUrlId(c.RequestCtx(), urlModel.ID)
	if err != nil {
		logger.Error("Ошибка при увеличении счетчика кликов: %v", err)
		return
	}

	ip := c.Locals("ip").(string)
	ua := c.Locals("ua").(string)

	fingerprint, err := r.services.FingerPrint.GetFingerPrintByIp(c.RequestCtx(), ip)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		logger.Error("Ошибка при получении отпечатка: %v", err)
		return
	}

	var fpua *service.CreateFullFingerPrintTx

	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		fpua, err = r.services.FingerPrint.CreateFullFingerPrint(c.RequestCtx(), stats.ID, ip, ua)
		if err != nil {
			logger.Error("Ошибка при создании отпечатка: %v", err)
			return
		}

		return
	}

	if fpua == nil && fingerprint != nil {
		userAgent, err := r.services.UserAgent.GetUserAgentByFpIdAgent(c.RequestCtx(), fingerprint.ID, ua)
		if err != nil && !errors.Is(err, pgx.ErrNoRows) {
			logger.Error("Ошибка при получении user_agent: %v", err)
			return
		}

		if err != nil && errors.Is(err, pgx.ErrNoRows) {
			_, err := r.services.UserAgent.CreateUserAgent(c.RequestCtx(), fingerprint.ID, ua)
			if err != nil {
				logger.Error("Ошибка при создании user_agent: %v", err)
				return
			}

			return
		}

		if userAgent != nil {
			r.services.UserAgent.UpdateUserAgentLastAccessedById(c.RequestCtx(), userAgent.ID)
		}
	}

	return
}

// GetOriginalURL - получает оригинальный URL по сокращенной ссылке
func (r *Router) GetOriginalURL(c fiber.Ctx) error {
	shortURL := c.Params("short_url")

	if shortURL == "" {
		logger.Error("Указана пустая сокращенная ссылка")
		return SetHTTPError(c, "Сокращенная ссылка не может быть пустой", fiber.StatusBadRequest)
	}

	urlModel, err := r.services.URLS.GetUrlByShortUrl(c.RequestCtx(), shortURL)
	if err != nil {
		logger.Error("Ошибка при получении оригинального url: %v", err)
		if errors.Is(err, pgx.ErrNoRows) {
			return SetHTTPError(c, "Сокращенная ссылка не найдена", fiber.StatusNotFound)
		}
		return SetHTTPError(c, "Ошибка при получении оригинального URL", fiber.StatusInternalServerError)
	}

	// Сбор статистики
	{
		r.statisticsCollection(c, urlModel)
	}

	return c.Redirect().To(urlModel.OriginalUrl)
}
