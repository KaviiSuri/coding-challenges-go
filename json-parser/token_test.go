package main

import "testing"

func TestLexer(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []Token
	}{
		{
			name:     "Empty Input",
			input:    "",
			expected: []Token{{Type: EOF, Literal: ""}},
		},
		{
			name:     "STRING",
			input:    `"hello"`,
			expected: []Token{{Type: STRING, Literal: "hello"}},
		},
		{
			name:     "NUMBER",
			input:    `42`,
			expected: []Token{{Type: NUMBER, Literal: "42"}},
		},
		{
			name:     "Boolean",
			input:    `true`,
			expected: []Token{{Type: TRUE, Literal: "true"}},
		},
		{
			name:     "Empty Object",
			input:    "{}",
			expected: []Token{{Type: LEFT_CURLY_BRACE, Literal: "{"}, {Type: RIGHT_CURLY_BRACE, Literal: "}"}, {Type: EOF, Literal: ""}},
		},
		{
			name:     "Empty Array",
			input:    "[]",
			expected: []Token{{Type: LEFT_SQUARE_BRACKET, Literal: "["}, {Type: RIGHT_SQUARE_BRACKET, Literal: "]"}, {Type: EOF, Literal: ""}},
		},
		{
			name:  "Object with STRING Value",
			input: `{ "key": "value" }`,
			expected: []Token{
				{Type: LEFT_CURLY_BRACE, Literal: "{"},
				{Type: STRING, Literal: "key"},
				{Type: COLON, Literal: ":"},
				{Type: STRING, Literal: "value"},
				{Type: RIGHT_CURLY_BRACE, Literal: "}"},
				{Type: EOF, Literal: ""},
			},
		},
		{
			name:  "Object with NUMBER Value (Carriage Returns)",
			input: `{\r"key": 42\r}`,
			expected: []Token{
				{Type: LEFT_CURLY_BRACE, Literal: "{"},
				{Type: STRING, Literal: "key"},
				{Type: COLON, Literal: ":"},
				{Type: NUMBER, Literal: "42"},
				{Type: RIGHT_CURLY_BRACE, Literal: "}"},
				{Type: EOF, Literal: ""},
			},
		},
		{
			name:  "Object with Boolean Value (New Line Characters)",
			input: `{\n"a": true,\n"b": false,\n"c": null\n}`,
			expected: []Token{
				{Type: LEFT_CURLY_BRACE, Literal: "{"},
				{Type: STRING, Literal: "a"},
				{Type: COLON, Literal: ":"},
				{Type: TRUE, Literal: "true"},
				{Type: COMMA, Literal: ","},
				{Type: STRING, Literal: "b"},
				{Type: COLON, Literal: ":"},
				{Type: FALSE, Literal: "false"},
				{Type: COMMA, Literal: ","},
				{Type: STRING, Literal: "c"},
				{Type: COLON, Literal: ":"},
				{Type: NULL, Literal: "null"},
				{Type: RIGHT_CURLY_BRACE, Literal: "}"},
				{Type: EOF, Literal: ""},
			},
		},
		{
			name:  "Object with Boolean, NULL, and NUMBER Values (Decimal and Negative)",
			input: `{ "key1": true, "key2": false, "key3": null, "key4": "value", "key5": 101.2, "key6": -42 }`,
			expected: []Token{
				{Type: LEFT_CURLY_BRACE, Literal: "{"},
				{Type: STRING, Literal: "key1"},
				{Type: COLON, Literal: ":"},
				{Type: TRUE, Literal: "true"},
				{Type: COMMA, Literal: ","},
				{Type: STRING, Literal: "key2"},
				{Type: COLON, Literal: ":"},
				{Type: FALSE, Literal: "false"},
				{Type: COMMA, Literal: ","},
				{Type: STRING, Literal: "key3"},
				{Type: COLON, Literal: ":"},
				{Type: NULL, Literal: "null"},
				{Type: COMMA, Literal: ","},
				{Type: STRING, Literal: "key4"},
				{Type: COLON, Literal: ":"},
				{Type: STRING, Literal: "value"},
				{Type: COMMA, Literal: ","},
				{Type: STRING, Literal: "key5"},
				{Type: COLON, Literal: ":"},
				{Type: NUMBER, Literal: "101.2"},
				{Type: COMMA, Literal: ","},
				{Type: STRING, Literal: "key6"},
				{Type: COLON, Literal: ":"},
				{Type: NUMBER, Literal: "-42"},
				{Type: RIGHT_CURLY_BRACE, Literal: "}"},
				{Type: EOF, Literal: ""},
			},
		},
		{
			name:  "Object with Empty Nested Object and Array Values",
			input: `{ "key": "value", "key-n": 101, "key-o": {}, "key-l": [] }`,
			expected: []Token{
				{Type: LEFT_CURLY_BRACE, Literal: "{"},
				{Type: STRING, Literal: "key"},
				{Type: COLON, Literal: ":"},
				{Type: STRING, Literal: "value"},
				{Type: COMMA, Literal: ","},
				{Type: STRING, Literal: "key-n"},
				{Type: COLON, Literal: ":"},
				{Type: NUMBER, Literal: "101"},
				{Type: COMMA, Literal: ","},
				{Type: STRING, Literal: "key-o"},
				{Type: COLON, Literal: ":"},
				{Type: LEFT_CURLY_BRACE, Literal: "{"},
				{Type: RIGHT_CURLY_BRACE, Literal: "}"},
				{Type: COMMA, Literal: ","},
				{Type: STRING, Literal: "key-l"},
				{Type: COLON, Literal: ":"},
				{Type: LEFT_SQUARE_BRACKET, Literal: "["},
				{Type: RIGHT_SQUARE_BRACKET, Literal: "]"},
				{Type: RIGHT_CURLY_BRACE, Literal: "}"},
				{Type: EOF, Literal: ""},
			},
		},
		{
			name:  "Object with Nested Object and Array Values",
			input: `{ "key": "value", "key-n": 101, "key-o": { "inner key": "inner value" }, "key-l": [ "list value" ] }`,
			expected: []Token{
				{Type: LEFT_CURLY_BRACE, Literal: "{"},
				{Type: STRING, Literal: "key"},
				{Type: COLON, Literal: ":"},
				{Type: STRING, Literal: "value"},
				{Type: COMMA, Literal: ","},
				{Type: STRING, Literal: "key-n"},
				{Type: COLON, Literal: ":"},
				{Type: NUMBER, Literal: "101"},
				{Type: COMMA, Literal: ","},
				{Type: STRING, Literal: "key-o"},
				{Type: COLON, Literal: ":"},
				{Type: LEFT_CURLY_BRACE, Literal: "{"},
				{Type: STRING, Literal: "inner key"},
				{Type: COLON, Literal: ":"},
				{Type: STRING, Literal: "inner value"},
				{Type: RIGHT_CURLY_BRACE, Literal: "}"},
				{Type: COMMA, Literal: ","},
				{Type: STRING, Literal: "key-l"},
				{Type: COLON, Literal: ":"},
				{Type: LEFT_SQUARE_BRACKET, Literal: "["},
				{Type: STRING, Literal: "list value"},
				{Type: RIGHT_SQUARE_BRACKET, Literal: "]"},
				{Type: RIGHT_CURLY_BRACE, Literal: "}"},
				{Type: EOF, Literal: ""},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewTokenizer(tt.input)
			for _, expected := range tt.expected {
				actual := l.NextToken()
				if actual.Type != expected.Type {
					t.Errorf("expected %+v, got %+v", tokens[expected.Type], tokens[actual.Type])
				}
				if actual.Literal != expected.Literal {
					t.Errorf("expected \"%+v\", got \"%+v\" %d %d",
						expected.Literal,
						actual.Literal,
						len(expected.Literal),
						len(actual.Literal),
					)
				}
			}
		})
	}
}
