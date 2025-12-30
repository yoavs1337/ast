package tokenizer

import (
	"fmt"
	"unicode"
)

type TokenType string

type Token struct {
	Type     TokenType
	Position int
	Len      int
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
	Errors       []error
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
	token := Token{Type: INT, Position: t.CurPosition, Len: 1}
	tokenLen := 1

	for unicode.IsDigit(rune(t.peekChar())) {
		t.readChar()
		tokenLen++

		if t.peekChar() == '.' && token.Type == INT {
			token.Type = FLOAT
			t.readChar()
			tokenLen++
		} else if t.peekChar() == '.' {
			token.Type = ERR
			t.readChar()
			tokenLen++
		}
	}

	token.Len = tokenLen
	if token.Type == ERR {
		errorMsg := fmt.Errorf(
			"number with multiple decimal points at [%d:%d]",
			token.Position,
			token.Position+token.Len,
		)
		t.Errors = append(t.Errors, errorMsg)
	}

	return token
}

func (t *Tokenizer) NextToken() Token {
	t.skipWhitespace()
	nextToken := Token{Type: EOF, Position: len(t.Input), Len: 1}

	switch t.ch {
	case 0:
		return nextToken
	case '+':
		nextToken = Token{Type: PLUS, Position: t.CurPosition, Len: 1}
	case '-':
		nextToken = Token{Type: MINUS, Position: t.CurPosition, Len: 1}
	case '*':
		nextToken = Token{Type: MUL, Position: t.CurPosition, Len: 1}
	case '/':
		nextToken = Token{Type: DIV, Position: t.CurPosition, Len: 1}
	case '^':
		nextToken = Token{Type: POW, Position: t.CurPosition, Len: 1}
	case '(':
		nextToken = Token{Type: LPAREN, Position: t.CurPosition, Len: 1}
	case ')':
		nextToken = Token{Type: RPAREN, Position: t.CurPosition, Len: 1}
	default:
		if unicode.IsDigit(rune(t.ch)) {
			nextToken = t.readNumber()
		} else {
			nextToken = Token{
				Type:     ERR,
				Position: t.CurPosition,
				Len:      1,
			}
			errorMsg := fmt.Errorf(
				"illegal character at [%d,%d]",
				nextToken.Position,
				nextToken.Position+nextToken.Len,
			)
			t.Errors = append(t.Errors, errorMsg)
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
