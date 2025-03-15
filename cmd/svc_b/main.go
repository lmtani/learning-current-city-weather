package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/lmtani/learning-current-city-weather/internal/entity"
	"github.com/lmtani/learning-current-city-weather/internal/infra/cep"
	"github.com/lmtani/learning-current-city-weather/internal/infra/otel"
	"github.com/lmtani/learning-current-city-weather/internal/usecase"
	"github.com/lmtani/learning-current-city-weather/pkg/weather"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func main() {
	// Handle SIGINT (CTRL+C) gracefully.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	otelShutdown, err := otel.SetupOTelSDK(ctx, "service-b")
	if err != nil {
		return
	}
	// Handle shutdown properly so nothing leaks.
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	fmt.Println("Service B is running on port 8080")
	// Start HTTP server.
	srv := &http.Server{
		Addr:         ":8080",
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
		ReadTimeout:  time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      newHTTPHandler(),
	}
	srvErr := make(chan error, 1)
	go func() {
		srvErr <- srv.ListenAndServe()
	}()

	// Wait for interruption.
	select {
	case err = <-srvErr:
		// Error when starting HTTP server.
		return
	case <-ctx.Done():
		// Wait for first CTRL+C.
		// Stop receiving signal notifications as soon as possible.
		stop()
	}

	// When Shutdown is called, ListenAndServe immediately returns ErrServerClosed.
	err = srv.Shutdown(context.Background())
	return
}

func newHTTPHandler() http.Handler {
	mux := http.NewServeMux()

	// handleFunc is a replacement for mux.HandleFunc
	// which enriches the handler's HTTP instrumentation with the pattern as the http.route.
	handleFunc := func(pattern string, handlerFunc func(http.ResponseWriter, *http.Request)) {
		// Configure the "http.route" for the HTTP instrumentation.
		handler := otelhttp.WithRouteTag(pattern, http.HandlerFunc(handlerFunc))
		mux.Handle(pattern, handler)
	}

	// Register handlers.
	handleFunc("/", GetTemperature)

	// Add HTTP instrumentation for the whole server.
	handler := otelhttp.NewHandler(mux, "/")
	return handler
}

// GetTemperature returns the temperature of a city.
func GetTemperature(w http.ResponseWriter, r *http.Request) {
	cepService := cep.NewService()
	weatherService := weather.NewService(os.Getenv("WEATHER_API_KEY"))
	getTemperature := usecase.NewGetTemperature(weatherService, cepService)
	cep := r.URL.Query().Get("cep")
	if cep == "" {
		http.Error(w, "Missing 'cep' query parameter", http.StatusBadRequest)
		return
	}

	output, err := getTemperature.Execute(cep)
	if err != nil {
		switch err {
		case entity.ErrCEPNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Return the temperature in JSON format.
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
