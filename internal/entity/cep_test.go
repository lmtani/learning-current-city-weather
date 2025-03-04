package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCEP_IsValid(t *testing.T) {
	t.Run("should return true when CEP is valid", func(t *testing.T) {
		c := CEP("12345678")
		assert.True(t, c.IsValid())
	})

	t.Run("should return false when CEP has less than 8 numbers", func(t *testing.T) {
		c := CEP("123")
		assert.False(t, c.IsValid())
	})

	t.Run("should return false when CEP has more than 8 numbers", func(t *testing.T) {
		c := CEP("123456789")
		assert.False(t, c.IsValid())
	})

	t.Run("should return false when CEP has letters", func(t *testing.T) {
		c := CEP("1234567A")
		assert.False(t, c.IsValid())
	})

	t.Run("should return false when CEP has hifen", func(t *testing.T) {
		c := CEP("12345-678")
		assert.False(t, c.IsValid())
	})
}
