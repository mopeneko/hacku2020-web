package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeartRateRepository_Get(t *testing.T) {
	repo := NewHeartRateRepository()

	var value uint = 120

	repo.count = value

	assert.Equal(t, repo.count, value)
}

func TestHeartRateRepository_Set(t *testing.T) {
	repo := NewHeartRateRepository()

	var value uint = 120

	repo.Set(value)

	assert.Equal(t, repo.count, value)
}

func TestHeartRateRepository(t *testing.T) {
	repo := NewHeartRateRepository()

	var value uint = 120

	repo.Set(value)
	count := repo.Get()

	assert.Equal(t, count, value)
}
