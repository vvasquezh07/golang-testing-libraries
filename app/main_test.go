package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnAMessage(t *testing.T) {
	assert := assert.New(t)

	var expected string = "Hello World"
	var current = PrintMessage()

	assert.Equal(expected, current, "The two messages should be the same.")
	assert.NotNil(current)
}

// Table driven test
func TestShouldReturnASum(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		x           int
		y           int
		expectedSum int
	}{
		{2, 3, 5},
		{-1, 1, 0},
		{0, 10, 10},
		{-5, -3, -8},
	}

	for _, test := range tests {
		assert.Equal(Sum(test.x, test.y), test.expectedSum)
	}
}