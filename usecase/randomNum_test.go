package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomNumbers(t *testing.T) {
	n := RandomNumbers()

	assert.GreaterOrEqual(t, n, 0)
	assert.LessOrEqual(t, n, 10000000000)

}
