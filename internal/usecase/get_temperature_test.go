package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/lmtani/learning-current-city-weather/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockWeatherService struct {
	mock.Mock
}

func (m *MockWeatherService) GetTemperature(city string) (float64, error) {
	args := m.Called(city)
	return args.Get(0).(float64), args.Error(1)
}

type MockCepService struct {
	mock.Mock
}

func (m *MockCepService) Get(cep string) (string, error) {
	args := m.Called(cep)
	return args.String(0), args.Error(1)
}

func TestGetTemperature_Execute(t *testing.T) {
	t.Run("should return temperature when CEP is valid", func(t *testing.T) {
		mockWeather := &MockWeatherService{}
		mockCep := &MockCepService{}

		mockCep.On("Get", "12345678").Return("São Paulo", nil).Times(1)
		mockWeather.On("GetTemperature", "São Paulo").Return(25.0, nil).Times(1)

		usecase := NewGetTemperature(mockWeather, mockCep)
		result, err := usecase.Execute("12345678")

		assert.Nil(t, err)
		assert.Equal(t, 25.0, result.Celsius)
		assert.Equal(t, 77.0, result.Fahrenheit)
		assert.Equal(t, 298.0, result.Kelvin)

		mockCep.AssertExpectations(t)
		mockWeather.AssertExpectations(t)
	})

	t.Run("should return ErrInvalidCEP when CEP is invalid", func(t *testing.T) {
		mockWeather := &MockWeatherService{}
		mockCep := &MockCepService{}

		usecase := NewGetTemperature(mockWeather, mockCep)
		_, err := usecase.Execute("123")

		assert.ErrorIs(t, err, entity.ErrCEPInvalid)
		mockCep.AssertNotCalled(t, "Get")
		mockWeather.AssertNotCalled(t, "GetTemperature")
	})

	t.Run("should return ErrNotFound when CEP not found", func(t *testing.T) {
		mockWeather := &MockWeatherService{}
		mockCep := &MockCepService{}

		mockCep.On("Get", "12345678").Return("", entity.ErrCEPNotFound).Times(1)

		usecase := NewGetTemperature(mockWeather, mockCep)
		_, err := usecase.Execute("12345678")

		assert.ErrorIs(t, err, entity.ErrCEPNotFound)
		mockCep.AssertExpectations(t)
		mockWeather.AssertNotCalled(t, "GetTemperature")
	})

	t.Run("should return ErrWeatherAPI when weather service fails", func(t *testing.T) {
		mockWeather := &MockWeatherService{}
		mockCep := &MockCepService{}

		mockCep.On("Get", "12345678").Return("São Paulo", nil)
		mockWeather.On("GetTemperature", "São Paulo").Return(0.0, errors.New("external error")).Times(1)

		usecase := NewGetTemperature(mockWeather, mockCep)
		_, err := usecase.Execute("12345678")

		assert.ErrorIs(t, err, entity.ErrWeatherAPI)
		mockCep.AssertExpectations(t)
		mockWeather.AssertExpectations(t)
	})

	t.Run("should return ErrNotFound when CEP not found after 3 retries", func(t *testing.T) {
		mockWeather := &MockWeatherService{}
		mockCep := &MockCepService{}

		mockCep.On("Get", "12345678").Return("", errors.New("request failed")).Times(3)

		usecase := NewGetTemperature(mockWeather, mockCep)
		usecase.TimeToSleep = 2 * time.Millisecond
		_, err := usecase.Execute("12345678")

		assert.ErrorIs(t, err, entity.ErrCEPNotFound)
		mockCep.AssertExpectations(t)
		mockWeather.AssertNotCalled(t, "GetTemperature")
	})
}
