package cep

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/lmtani/learning-current-city-weather/internal/entity"
)

// Cep represents a Brazilian postal code
type Cep struct {
	Cep         string `json:"cep"`
	Bairro      string `json:"neighborhood"`
	Rua         string `json:"street"`
	Cidade      string `json:"city"`
	Uf          string `json:"state"`
	TimeElapsed int64
}

type BrasilApiOutput struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

type BrasilApi struct {
	url     string
	timeout time.Duration
}

func NewBrasilApi(host string) *BrasilApi {
	return &BrasilApi{url: host}
}

func (b *BrasilApi) GetCep(cep string) (*Cep, error) {
	start := time.Now()
	route := fmt.Sprintf("%s/api/cep/v1/%s", b.url, cep)

	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return nil, entity.ErrCEPNotFound
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("request failed")
	}

	var p *BrasilApiOutput
	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		return nil, err
	}

	c := &Cep{
		Cep:         p.Cep,
		Bairro:      p.Neighborhood,
		Rua:         p.Street,
		Cidade:      p.City,
		Uf:          p.State,
		TimeElapsed: time.Since(start).Milliseconds(),
	}
	return c, nil
}
