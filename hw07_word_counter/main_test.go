package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]int
	}{
		{
			name:     "Simple text",
			input:    "Hello world! This is a simple text, hello world.",
			expected: map[string]int{"hello": 2, "world": 2, "this": 1, "is": 1, "a": 1, "simple": 1, "text": 1},
		},
		{
			name:     "Empty input",
			input:    "",
			expected: map[string]int{},
		},
		{
			name:     "Repeated words",
			input:    "apple apple banana banana banana apple",
			expected: map[string]int{"apple": 3, "banana": 3},
		},
		{
			name:     "Mixed case",
			input:    "Go Golang golang GO",
			expected: map[string]int{"go": 2, "golang": 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := CountWords(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
