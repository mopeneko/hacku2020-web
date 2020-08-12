package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeartRateRepository(t *testing.T) {
	repo := NewHeartRateRepository()

	var value uint = 120

	repo.Set(value)
	count := repo.Get()

	assert.Equal(t, count, value)
}
