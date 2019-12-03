package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func shouldReturnAMessageUsingTestify(t *testing.T) {
	assert := assert.New(t)

	var expected string = "Hello World"
	var current = PrintMessage()

	assert.Equal(expected, current, "The two messages should be the same.")
}
