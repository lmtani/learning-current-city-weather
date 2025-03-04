package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemperature_GetKelvin(t *testing.T) {
	tc := []struct {
		name     string
		celsius  float64
		expected float64
	}{
		{
			name:     "should return 273 when celsius is 0",
			celsius:  0,
			expected: 273,
		},
		{
			name:     "should return 274 when celsius is 1",
			celsius:  1,
			expected: 274,
		},
		{
			name:     "should return 272 when celsius is -1",
			celsius:  -1,
			expected: 272,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			tm := Temperature{Celsius: tt.celsius}
			assert.Equal(t, tm.GetKelvin(), tt.expected)
		})
	}
}

func TestTemperature_GetCelsius(t *testing.T) {
	tc := []struct {
		name     string
		celsius  float64
		expected float64
	}{
		{
			name:     "should return 0 when celsius is 0",
			celsius:  0,
			expected: 0,
		},
		{
			name:     "should return 1 when celsius is 1",
			celsius:  1,
			expected: 1,
		},
		{
			name:     "should return -1 when celsius is -1",
			celsius:  -1,
			expected: -1,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			tm := Temperature{Celsius: tt.celsius}
			assert.Equal(t, tm.GetCelsius(), tt.expected)
		})
	}
}

func TestTemperature_GetFahrenheit(t *testing.T) {
	tc := []struct {
		name     string
		celsius  float64
		expected float64
	}{
		{
			name:     "should return 32 when celsius is 0",
			celsius:  0,
			expected: 32,
		},
		{
			name:     "should return 33.8 when celsius is 1",
			celsius:  1,
			expected: 33.8,
		},
		{
			name:     "should return 30.2 when celsius is -1",
			celsius:  -1,
			expected: 30.2,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			tm := Temperature{Celsius: tt.celsius}
			assert.Equal(t, tm.GetFahrenheit(), tt.expected)
		})
	}
}
