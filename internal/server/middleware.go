package server

import (
	"shortner/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

// middleware — сбор данных
func middleware(c fiber.Ctx) error {
	// Простейшие CORS-заголовки (лучше вынести в отдельный middleware)
	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("Content-Type", "application/json")

	logger.Info("middleware")

	ip := c.IP()
	ua := c.Get("User-Agent")

	c.Locals("ip", ip)
	c.Locals("ua", ua)

	return c.Next()
}
