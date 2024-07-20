package main

import (
	"fmt"
	"strconv"
)

func NewParser(input string) *Parser {
	return &Parser{
		tokenizer: NewTokenizer(input),
	}
}

type Parser struct {
	tokenizer Tokenizer
}

func (p *Parser) Parse() (interface{}, error) {
	var output interface{}
	var err error

	tok := p.tokenizer.NextToken()

	output, err = p.ParseToken(tok)

	tok = p.tokenizer.NextToken()

	if tok.Type != EOF {
		err = fmt.Errorf("expected end of input but found %s", tok.Literal)
	}
	return output, err
}

func (p *Parser) ParseObject(obj map[string]interface{}) (interface{}, error) {
	var err error
	// The LEFT_CURLY_BRACE token is read in Parse, and given it has no use in our parsing
	// other than to identify an object value, we immediately advance to the next token
	tok := p.tokenizer.NextToken()

	if tok.Type == RIGHT_CURLY_BRACE {
		return obj, err
	}

	// To iteratively parse multi-key objects, we loop until we don't find encounter
	// the COMMA token
	for {
		if tok.Type != STRING {
			return obj, fmt.Errorf("expected key but found %s", tok.Literal)
		}

		key := tok.Literal

		tok = p.tokenizer.NextToken()

		if tok.Type != COLON {
			return obj, fmt.Errorf("expected name separate but found %s", tok.Literal)
		}

		tok = p.tokenizer.NextToken()

		value, err := p.ParseToken(tok)
		if err != nil {
			return obj, err
		}

		obj[key] = value

		tok = p.tokenizer.NextToken()

		if tok.Type != COMMA {
			break
		}

		// We advance past the COMMA token to arrive at the next
		// key in the object
		tok = p.tokenizer.NextToken()
	}

	if tok.Type != RIGHT_CURLY_BRACE {
		err = fmt.Errorf("expected } but found %s", tok.Literal)
	}

	return obj, err
}

func (p *Parser) ParseArray(arr []interface{}) (interface{}, error) {
	var err error

	tok := p.tokenizer.NextToken()

	if tok.Type == RIGHT_SQUARE_BRACKET {
		return arr, err
	}

	// To iteratively parse arrays, we loop until we don't encounter
	// the COMMA token
	for {
		value, err := p.ParseToken(tok)
		if err != nil {
			return arr, err
		}

		arr = append(arr, value)

		tok = p.tokenizer.NextToken()

		if tok.Type != COMMA {
			break
		}

		// We advance past the COMMA token to arrive at the next
		// value in the array
		tok = p.tokenizer.NextToken()
	}

	if tok.Type != RIGHT_SQUARE_BRACKET {
		err = fmt.Errorf("expected end of input but found %s", tok.Literal)
	}

	return arr, err
}

func (p *Parser) ParseToken(tok Token) (interface{}, error) {
	var value interface{}
	var err error
	switch tok.Type {
	case TRUE:
		value = true
	case FALSE:
		value = false
	case NULL:
		value = nil
	case STRING:
		value = tok.Literal
	case NUMBER:
		value, err = strconv.ParseFloat(tok.Literal, 64)
		if err != nil {
			return value, err
		}
	case LEFT_CURLY_BRACE:
		value, err = p.ParseObject(make(map[string]interface{}))
	case LEFT_SQUARE_BRACKET:
		value, err = p.ParseArray(make([]interface{}, 0))
	case ILLEGAL:
		err = fmt.Errorf("illegal token %s encountered", tok.Literal)
	case EOF:
		err = fmt.Errorf("unexpected end of input")
	default:
		err = fmt.Errorf("unknown token %s", tok.Literal)
	}
	return value, err
}
