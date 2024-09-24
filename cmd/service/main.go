//go:build !test

package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/s21platform/gateway-service/internal/config"
	"github.com/s21platform/gateway-service/internal/handlers/api"
	authhandler "github.com/s21platform/gateway-service/internal/handlers/auth"
	"github.com/s21platform/gateway-service/internal/middlewares"
	"github.com/s21platform/gateway-service/internal/rpc/auth"
	authusecase "github.com/s21platform/gateway-service/internal/useCase/auth"
	"github.com/s21platform/metrics-lib/pkg"
	"log"
	"net/http"
)

func main() {
	cfg := config.MustLoad()

	metrics, err := pkg.NewMetrics(cfg.Metrics.Host, cfg.Metrics.Port)
	if err != nil {
		log.Fatalf("failed to init metrics: %v", err)
	}

	// rpc clients
	authClient := auth.NewService(cfg)

	// usecases declaration
	authUseCase := authusecase.New(authClient)

	// handlers declaration
	authHandlers := authhandler.New(authUseCase)
	apiHandlers := api.New()

	r := chi.NewRouter()

	r.Use(func(next http.Handler) http.Handler {
		return middlewares.MetricMiddleware(next, metrics)
	})

	authhandler.AttachAuthRoutes(r, authHandlers)
	api.AttachApiRoutes(r, apiHandlers, cfg)

	fmt.Println(fmt.Sprintf(":%s", cfg.Service.Port))

	http.ListenAndServe(fmt.Sprintf(":%s", cfg.Service.Port), r)
}
