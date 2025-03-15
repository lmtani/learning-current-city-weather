package cep

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

// Service provides the city of a given CEP.
type Service struct {
	brasilAPI string
	Tracer    trace.Tracer
}

// NewService creates a new CEP service.
func NewService() *Service {
	return &Service{
		brasilAPI: "https://brasilapi.com.br",
	}
}

// Get returns the city of the given CEP.
func (s Service) Get(ctx context.Context, queryCEP string) (string, error) {
	// if tracer is provided, use it
	if s.Tracer != nil {
		_, span := s.Tracer.Start(ctx, "request-brasilapi")
		defer span.End()
	}

	brasilapi := NewBrasilApi(s.brasilAPI)

	resp, err := brasilapi.GetCep(queryCEP)
	if err != nil {
		return "", err
	}

	return resp.Cidade, nil
}
