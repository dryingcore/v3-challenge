package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	handler "github.com/dryingcore/v3-challenge/internal/adapter/handler/http"
	"github.com/dryingcore/v3-challenge/internal/core/usecase"
	"github.com/dryingcore/v3-challenge/pkg/queue"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	publisher := queue.NewFakePublisher()

	_handler := handler.NewTelemetryHandler(
		usecase.NewGyroscopeUC(publisher),
		usecase.NewGPSUseCase(publisher),
	)

	r.Post("/telemetry/gyroscope", _handler.HandleGyroscope)
	r.Post("/telemetry/gps", _handler.HandleGPS)

	log.Println("Servidor iniciado na porta :3000")
	http.ListenAndServe(":3000", r)
}
