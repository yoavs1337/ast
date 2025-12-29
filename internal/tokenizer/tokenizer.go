package tokenizer

import (
	"fmt"
	"unicode"
)

type TokenType string

type Token struct {
	Type     TokenType
	Literal  string
	Position int
}

const (
	// numerical values
	INT   TokenType = "INT"
	FLOAT TokenType = "FLOAT"

	// mathematical operations
	PLUS  TokenType = "PLUS"
	MINUS TokenType = "MINUS"
	MUL   TokenType = "MUL"
	DIV   TokenType = "DIV"
	POW   TokenType = "POW"

	// parentheses
	LPAREN TokenType = "LPARAN"
	RPAREN TokenType = "RPARAN"

	// special tokens
	EOF TokenType = "EOF"
	ERR TokenType = "ERR"
)

type Tokenizer struct {
	Input        string
	CurPosition  int
	NextPosition int
	ch           byte
}

func (t *Tokenizer) readChar() {
	if t.CurPosition == len(t.Input)-1 {
		t.ch = 0
	} else {
		t.CurPosition = t.NextPosition
		t.NextPosition++
		t.ch = t.Input[t.CurPosition]
	}
}

func (t *Tokenizer) skipWhitespace() {
	for unicode.IsSpace(rune(t.ch)) {
		t.readChar()
	}
}

func (t *Tokenizer) peekChar() byte {
	if t.CurPosition == len(t.Input)-1 {
		return 0
	}
	return t.Input[t.NextPosition]
}

func (t *Tokenizer) readNumber() Token {
	token := Token{Type: INT, Literal: "", Position: t.CurPosition}
	value := string(t.ch)

	for unicode.IsDigit(rune(t.peekChar())) {
		t.readChar()
		value = value + string(t.ch)

		if t.peekChar() == '.' && token.Type == INT {
			token.Type = FLOAT
			value = value + "."
			t.readChar()
		} else if t.peekChar() == '.' {
			return Token{
				Type:     ERR,
				Literal:  "two decimal points in a single number",
				Position: t.NextPosition,
			}
		}
	}

	token.Literal = value
	return token
}

func (t *Tokenizer) NextToken() Token {
	t.skipWhitespace()
	nextToken := Token{Type: EOF, Literal: "", Position: len(t.Input)}

	switch t.ch {
	case 0:
		return nextToken
	case '+':
		nextToken = Token{Type: PLUS, Literal: "", Position: t.CurPosition}
	case '-':
		nextToken = Token{Type: MINUS, Literal: "", Position: t.CurPosition}
	case '*':
		nextToken = Token{Type: MUL, Literal: "", Position: t.CurPosition}
	case '/':
		nextToken = Token{Type: DIV, Literal: "", Position: t.CurPosition}
	case '^':
		nextToken = Token{Type: POW, Literal: "", Position: t.CurPosition}
	case '(':
		nextToken = Token{Type: LPAREN, Literal: "", Position: t.CurPosition}
	case ')':
		nextToken = Token{Type: RPAREN, Literal: "", Position: t.CurPosition}
	default:
		if unicode.IsDigit(rune(t.ch)) {
			nextToken = t.readNumber()
		} else {
			nextToken = Token{
				Type:     ERR,
				Literal:  "illegal character",
				Position: t.CurPosition,
			}
		}
	}

	t.readChar()

	return nextToken
}

func (t *Tokenizer) Tokenize() []Token {
	result := make([]Token, 0)
	token := t.NextToken()
	result = append(result, token)

	for token.Type != EOF {
		token = t.NextToken()
		result = append(result, token)
	}
	return result
}

func NewTokenizer(input string) (Tokenizer, error) {
	if len(input) == 0 {
		return Tokenizer{}, fmt.Errorf("tokenizer cannot be created for an empty string")
	}
	return Tokenizer{
		Input:        input,
		CurPosition:  0,
		NextPosition: 1,
		ch:           input[0],
	}, nil
}
