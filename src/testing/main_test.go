package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Ega")
	assert.Equal(t, "Hello Ega", result, "result is not 'Hello Ega'")
}
