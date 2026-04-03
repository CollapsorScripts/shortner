package server

import (
	"encoding/json"
	"fmt"
	"shortner/internal/bootstrap"
	"shortner/internal/config"
	"sync"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	srvLog "github.com/gofiber/fiber/v3/middleware/logger"
)

const apiStr = "/api/v1"
const contextTimeout = time.Second * 30

var endpointList map[string]fiber.Handler

type Router struct {
	cfg *config.Config
	r   *fiber.App
	mu  sync.Mutex
	//сервисы
	services *bootstrap.Services
}

// NewRouter - создает новый роутер для маршрутизации
func NewRouter(cfg *config.Config, services *bootstrap.Services) *fiber.App {
	fiberCFG := fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	}

	router := &Router{
		r:        fiber.New(fiberCFG),
		mu:       sync.Mutex{},
		cfg:      cfg,
		services: services,
	}

	router.r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowMethods: []string{"GET", "HEAD", "PUT", "POST", "DELETE"},
	}))

	//router.r.Use(auth.AuthMiddleware)

	return router.loadEndpoints()
}

func (r *Router) createEndpoints() {
	list := map[string]fiber.Handler{
		"/shorten":          r.Shorten,
		"/:short_url":       r.GetOriginalURL,
		"/stats/:short_url": r.GetStats,
	}

	endpointList = list
}

func getHandler(endpoint string) (string, fiber.Handler) {
	if _, ok := endpointList[endpoint]; ok {
		return endpoint, endpointList[endpoint]
	}

	return endpoint, func(c fiber.Ctx) error {
		return SetHTTPError(c, "Данный метод не реализован", fiber.StatusNotImplemented)
	}
}

func (route *Router) loadEndpoints() *fiber.App {
	//Инициализируем эндпоинты и обработчики
	{
		route.createEndpoints()
	}

	route.r.Use(srvLog.New())

	endpoint := route.r.Group(fmt.Sprintf("%s/", apiStr))

	shortEndpoint := route.r.Group("/")
	shortEndpoint.Use(middleware)

	// Сокращение URL
	endpoint.Post(getHandler("/shorten"))
	// Доступ по сокращенному URL
	shortEndpoint.Get(getHandler("/:short_url"))
	// Статистика
	endpoint.Get(getHandler("/stats/:short_url"))

	return route.r
}
