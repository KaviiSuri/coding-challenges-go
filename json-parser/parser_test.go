package main

import (
	"reflect"
	"testing"
)

func TestParser(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedValue interface{}
		expectError   bool
	}{
		{
			name:          "Empty Input",
			input:         "",
			expectedValue: nil,
			expectError:   true,
		},
		{
			name:          "String",
			input:         `"hello"`,
			expectedValue: "hello",
			expectError:   false,
		},
		{
			name:          "Number",
			input:         `42`,
			expectedValue: float64(42),
			expectError:   false,
		},
		{
			name:          "Invalid Number",
			input:         `42..5`,
			expectedValue: float64(0),
			expectError:   true,
		},
		{
			name:          "Boolean",
			input:         `true`,
			expectedValue: true,
			expectError:   false,
		},
		{
			name:          "Empty Object",
			input:         "{}",
			expectedValue: make(map[string]interface{}),
			expectError:   false,
		},
		{
			name:          "Invalid Object (Unbalanced Brace)",
			input:         "{",
			expectedValue: make(map[string]interface{}),
			expectError:   true,
		},
		{
			name:          "Invalid Object (Unbalanced Extra Brace)",
			input:         "{}}",
			expectedValue: make(map[string]interface{}),
			expectError:   true,
		},
		{
			name:          "Simple Object with String Value",
			input:         `{"key":"value"}`,
			expectedValue: map[string]interface{}{"key": "value"},
			expectError:   false,
		},
		{
			name:          "Invalid Simple Object with String Value",
			input:         `{"key':"value"}`,
			expectedValue: make(map[string]interface{}),
			expectError:   true,
		},
		{
			name:          "Simple Object with Numeric, Boolean and Null Values",
			input:         `{ "keyA": "value", "keyB": 42.5, "keyC": true, "keyD": null }`,
			expectedValue: map[string]interface{}{"keyA": "value", "keyB": float64(42.5), "keyC": true, "keyD": nil},
			expectError:   false,
		},
		{
			name:          "Simple Object Containing Nested Objects",
			input:         `{"key": {"key2": "value"}, "key3": []}`,
			expectedValue: map[string]interface{}{"key": map[string]interface{}{"key2": "value"}, "key3": []interface{}{}},
			expectError:   false,
		},
		{
			name:          "Empty Array",
			input:         "[]",
			expectedValue: make([]interface{}, 0),
			expectError:   false,
		},
		{
			name:          "Invalid Array (Unbalanced Brace)",
			input:         "[",
			expectedValue: make([]interface{}, 0),
			expectError:   true,
		},
		{
			name:          "Invalid Array (Unbalanced Extra Brace)",
			input:         "[]]",
			expectedValue: make([]interface{}, 0),
			expectError:   true,
		},
		{
			name:          "Simple Array with String Values",
			input:         `[ "a", "b", "c" ]`,
			expectedValue: []interface{}{"a", "b", "c"},
			expectError:   false,
		},
		{
			name:          "Simple Array with Numeric, Boolean, and Null Values",
			input:         `[ true, false, null, -42.5, "a" ]`,
			expectedValue: []interface{}{true, false, nil, float64(-42.5), "a"},
			expectError:   false,
		},
		{
			name:  "Array with Object Values",
			input: `[ { "key": "a" }, { "key": "b"}, {"key": "c"} ]`,
			expectedValue: []interface{}{
				map[string]interface{}{"key": "a"},
				map[string]interface{}{"key": "b"},
				map[string]interface{}{"key": "c"},
			},
			expectError: false,
		},
		{
			name:  "Array with Nested Arrays",
			input: `[ ["a"], ["b"], ["c"] ]`,
			expectedValue: []interface{}{
				[]interface{}{"a"},
				[]interface{}{"b"},
				[]interface{}{"c"},
			},
			expectError: false,
		},
		{
			name:  "Complex Object",
			input: `{"key": {"key2": "value", "key3": { "key4": [-84.25] } }, "key5": [{"key6": false, "key7": null }] }`,
			expectedValue: map[string]interface{}{
				"key": map[string]interface{}{
					"key2": "value",
					"key3": map[string]interface{}{
						"key4": []interface{}{float64(-84.25)},
					},
				},
				"key5": []interface{}{
					map[string]interface{}{"key6": false, "key7": nil},
				},
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewParser(tt.input)
			actualValue, actualError := p.Parse()
			if actualError != nil && !tt.expectError {
				t.Fatalf("Expected no error but received: %v", actualError)
			}
			if !reflect.DeepEqual(actualValue, tt.expectedValue) {
				t.Fatalf("Expected %T but received %T", tt.expectedValue, actualValue)
			}
		})
	}
}
