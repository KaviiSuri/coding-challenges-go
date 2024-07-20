package internal

import (
	"reflect"
	"testing"
)

func TestAnalyseFreqs(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected Freqs
	}{
		{
			name:  "test",
			input: "abbcc",
			expected: Freqs{
				'a': 1,
				'b': 2,
				'c': 2,
			},
		},
		{
			name:  "test",
			input: "aabbcc",
			expected: Freqs{
				'a': 2,
				'b': 2,
				'c': 2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := AnalyseFreq([]byte(tt.input))
			if !reflect.DeepEqual(f, tt.expected) {
				t.Errorf("Mismatch in %v", tt.input)
				t.Error("Expected:", PrettyPrentFreqs(tt.expected))
				t.Error("Got:", PrettyPrentFreqs(f))
			}
		})
	}
}
