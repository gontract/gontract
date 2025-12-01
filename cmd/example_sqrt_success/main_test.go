package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateSqrt(t *testing.T) {
	var num float64 = 1
	root := CalculateSqrt(num)
	assert.Equal(t, root, num)
}
