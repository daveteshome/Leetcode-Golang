package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_main(t *testing.T) {
	test := []struct {
		name     string
		values   []int
		expected bool
	}{
		{
			name:     "check contain duplicate 1",
			values:   []int{1, 2, 3, 4},
			expected: false,
		},
		{
			name:     "check contain duplicate 1",
			values:   []int{1, 2, 3, 1},
			expected: true,
		},
		{
			name:     "check contain duplicate 1",
			values:   []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			expected: true,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			result := containsDuplicate(tt.values)
			assert.Equal(t, result, tt.expected)
		})
	}
}
