package internal

import (
	"reflect"
	"testing"
)

// TestBuildCode tests the BuildCode method using a table-based approach
func TestBuildCode(t *testing.T) {
	tests := []struct {
		name     string
		root     *Node
		expected Code
	}{
		{
			name: "Basic",
			root: &Node{
				Left: &Node{
					Ch:    'a',
					Left:  nil,
					Right: nil,
				},
				Right: &Node{
					Ch:    'b',
					Left:  nil,
					Right: nil,
				},
			},
			expected: Code{
				'a': {false},
				'b': {true},
			},
		},
		{
			name: "Complex",
			root: &Node{
				Left: &Node{
					Ch: 0,
					Left: &Node{
						Ch:    'a',
						Left:  nil,
						Right: nil,
					},
					Right: &Node{
						Ch:    'b',
						Left:  nil,
						Right: nil,
					},
				},
				Right: &Node{
					Ch:    'c',
					Left:  nil,
					Right: nil,
				},
			},
			expected: Code{
				'a': {false, false},
				'b': {false, true},
				'c': {true},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code := tt.root.BuildCode()

			for k, v := range tt.expected {
				if !reflect.DeepEqual(code[k], v) {
					t.Errorf("For character '%c', expected %v but got %v", k, v, code[k])
				}
			}

			if len(code) != len(tt.expected) {
				t.Errorf("Expected code map length %d but got %d", len(tt.expected), len(code))
			}
		})
	}
}
