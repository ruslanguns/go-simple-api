package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/ruslanguns/go-simple-api/internal/api/handlers"
	"github.com/ruslanguns/go-simple-api/internal/config"
	"github.com/ruslanguns/go-simple-api/internal/services"
	"github.com/ruslanguns/go-simple-api/pkg/logger"
)

func NewServer(log *logger.Logger, cfg *config.Config) http.Handler {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	// r.Use(customMiddleware.Logging(log))

	// Services
	userService := services.NewUserService()
	productService := services.NewProductService()

	// Handlers
	userHandler := handlers.NewUserHandler(log, userService)
	productHandler := handlers.NewProductHandler(log, productService)

	// Routes /api/v1/users & /api/v1/products
	r.Route("/api/v1", func(r chi.Router) {
		r.Mount("/users", userHandler.Routes())
		r.Mount("/products", productHandler.Routes())
	})

	r.Get("/healthz", handlers.HandleHealthCheck(log))

	return r
}
