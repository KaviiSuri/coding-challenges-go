package main

import "fmt"

type TokenType int

const (
	ILLEGAL TokenType = iota
	EOF

	LEFT_CURLY_BRACE  // '{'
	RIGHT_CURLY_BRACE // '}'

	LEFT_SQUARE_BRACKET  // '['
	RIGHT_SQUARE_BRACKET // ']'

	COMMA // ','
	COLON // ':'

	STRING //

	NUMBER
	TRUE
	FALSE
	NULL
)

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",

	LEFT_CURLY_BRACE:  "LEFT_CURLY_BRACE",
	RIGHT_CURLY_BRACE: "RIGHT_CURLY_BRACE",

	LEFT_SQUARE_BRACKET:  "LEFT_SQUARE_BRACKET",
	RIGHT_SQUARE_BRACKET: "RIGHT_SQUARE_BRACKET",

	COMMA: "COMMA",
	COLON: "COLON",

	STRING: "STRING",
	NUMBER: "NUMBER",
	TRUE:   "true",
	FALSE:  "false",
	NULL:   "null",
}

type Token struct {
	Type    TokenType
	Literal string
}

type Tokenizer struct {
	str     string
	pos     int
	readPos int
	ch      byte
}

func NewTokenizer(input string) Tokenizer {
	t := Tokenizer{
		str: input,
	}
	t.readChar()
	return t
}

func (t *Tokenizer) readChar() {
	if t.readPos >= len(t.str) {
		t.ch = 0
	} else {
		t.ch = t.str[t.readPos]
	}
	t.pos = t.readPos
	t.readPos += 1
}

func (t *Tokenizer) peekChar() byte {
	if t.readPos >= len(t.str) {
		return 0
	}
	return t.str[t.readPos]
}

func (t *Tokenizer) skipWhitespace() {
	for {
		if t.ch == ' ' {
			t.readChar()
			continue
		}

		if t.ch == '\\' && (t.peekChar() == 't' || t.peekChar() == 'n' || t.peekChar() == 'r') {
			t.readChar()
			t.readChar()
			continue
		}
		break
	}
}

func (t *Tokenizer) NextToken() (tok Token) {
	t.skipWhitespace()

	switch t.ch {
	case '{':
		tok.Type = LEFT_CURLY_BRACE
		tok.Literal = string(t.ch)
	case '}':
		tok.Type = RIGHT_CURLY_BRACE
		tok.Literal = string(t.ch)
	case '[':
		tok.Type = LEFT_SQUARE_BRACKET
		tok.Literal = string(t.ch)
	case ']':
		tok.Type = RIGHT_SQUARE_BRACKET
		tok.Literal = string(t.ch)
	case ',':
		tok.Type = COMMA
		tok.Literal = string(t.ch)
	case ':':
		tok.Type = COLON
		tok.Literal = string(t.ch)
	case '"':
		tok = t.readString()
	case 0:
		tok.Type = EOF
	default:
		if isDigit(t.ch) || t.ch == '-' {
			tok = t.readNumber()
		} else if isLiteralName(t.ch) {
			tok = t.readLiteral()
		} else {
			tok.Type = ILLEGAL
			tok.Literal = string(t.ch)
		}
	}
	t.readChar()

	return tok
}

func (t *Tokenizer) readNumber() Token {
	start := t.pos
	for isDigit(t.peekChar()) || t.peekChar() == '.' || t.peekChar() == '-' {
		t.readChar()
	}
	return Token{
		Type:    NUMBER,
		Literal: t.str[start:t.readPos],
	}
}

func (t *Tokenizer) readLiteral() (tok Token) {
	start := t.pos

	switch t.ch {
	case 't':
		for i, c := range tokens[TRUE][1:] {
			if c != rune(t.peekChar()) {
				fmt.Printf("%v %v", c, t.peekChar())
				return Token{ILLEGAL, t.str[start : start+i]}
			}
			tok.Literal = "true"
			tok.Type = TRUE
			t.readChar()
		}
	case 'f':
		for i, c := range tokens[FALSE][1:] {
			if c != rune(t.peekChar()) {
				return Token{ILLEGAL, t.str[start : start+i]}
			}
			tok.Type = FALSE
			tok.Literal = "false"
			t.readChar()
		}
	case 'n':
		for i, c := range tokens[NULL][1:] {
			if c != rune(t.peekChar()) {
				return Token{ILLEGAL, t.str[start : start+i]}
			}
			tok.Type = NULL
			tok.Literal = "null"
			t.readChar()
		}
	default:
		return Token{ILLEGAL, string(t.ch)}
	}

	return tok
}

func (t *Tokenizer) readString() (tok Token) {
	start := t.pos + 1

	for {
		t.readChar()
		if t.ch == '"' {
			tok.Type = STRING
			tok.Literal = t.str[start:t.pos]
			break
		}
		if t.ch == 0 {
			if t.str[t.pos-1] != '"' {
				tok.Type = ILLEGAL
				tok.Literal = t.str[start:t.pos]
			}
			break
		}
	}

	return tok
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isLiteralName(ch byte) bool {
	return ch == 't' || ch == 'f' || ch == 'n'
}
