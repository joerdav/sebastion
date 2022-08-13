package sebastion_test

import (
	"testing"

	"github.com/joerdav/sebastion"
)

func TestGetInt(t *testing.T) {
	tests := []struct {
		name     string
		values   sebastion.InputValues
		idx      int
		expected int
	}{
		{
			name:     "given an empty map return 0",
			values:   map[int]interface{}{},
			idx:      0,
			expected: 0,
		},
		{
			name: "given an idx that doesn't exist, return 0",
			values: map[int]interface{}{
				1: 2,
			},
			idx:      2,
			expected: 0,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if a := tt.values.GetInt(tt.idx); a != tt.expected {
				t.Errorf("values.GetInt(%v) = %v, want: %v", tt.idx, a, tt.expected)
			}
		})
	}
}
