package cep

// Service provides the city of a given CEP.
type Service struct {
	brasilAPI string
}

// NewService creates a new CEP service.
func NewService() *Service {
	return &Service{
		brasilAPI: "https://brasilapi.com.br",
	}
}

// Get returns the city of the given CEP.
func (s Service) Get(queryCEP string) (string, error) {
	brasilapi := NewBrasilApi(s.brasilAPI)

	resp, err := brasilapi.GetCep(queryCEP)
	if err != nil {
		return "", err
	}

	return resp.Cidade, nil
}
