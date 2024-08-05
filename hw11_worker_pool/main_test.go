package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncremetation(t *testing.T) {
	counter := Counter{}
	counter.C = 10

	counter.Incrementation(100)

	assert.Equal(t, counter.C, 110)
}
