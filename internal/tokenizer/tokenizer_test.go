package tokenizer

import (
	"errors"
	"fmt"
	"testing"
)

type TestTokenizerStruct struct {
	Input                      string
	Output                     []Token
	ExpectedTokenizationErrors []error
	Error                      error
}

func TestTokenize(t *testing.T) {
	tests := []TestTokenizerStruct{
		{
			Input: "10 + 55",
			Output: []Token{
				{Type: INT, Position: 0, Len: 2},
				{Type: PLUS, Position: 3, Len: 1},
				{Type: INT, Position: 5, Len: 2},
				{Type: EOF, Position: 7, Len: 1},
			},
			Error: nil,
		},
		{
			Input: "10.3 - 4 ",
			Output: []Token{
				{Type: FLOAT, Position: 0, Len: 4},
				{Type: MINUS, Position: 5, Len: 1},
				{Type: INT, Position: 7, Len: 1},
				{Type: EOF, Position: 9, Len: 1},
			},
			Error: nil,
		},
		{
			Input: "(10 * 5^3) / 4.5",
			Output: []Token{
				{Type: LPAREN, Position: 0, Len: 1},
				{Type: INT, Position: 1, Len: 2},
				{Type: MUL, Position: 4, Len: 1},
				{Type: INT, Position: 6, Len: 1},
				{Type: POW, Position: 7, Len: 1},
				{Type: INT, Position: 8, Len: 1},
				{Type: RPAREN, Position: 9, Len: 1},
				{Type: DIV, Position: 11, Len: 1},
				{Type: FLOAT, Position: 13, Len: 3},
				{Type: EOF, Position: 16, Len: 1},
			},
			Error: nil,
		},
		{
			Input: "4..5",
			Output: []Token{
				{Type: ERR, Position: 0, Len: 4},
				{Type: EOF, Position: 4, Len: 1},
			},
			ExpectedTokenizationErrors: []error{ErrMultipleDecimalPoints},
			Error:                      nil,
		},
		{
			Input: "4.5%",
			Output: []Token{
				{Type: FLOAT, Position: 0, Len: 3},
				{Type: ERR, Position: 3, Len: 1},
				{Type: EOF, Position: 4, Len: 1},
			},
			ExpectedTokenizationErrors: []error{ErrIllegalCharacter},
			Error:                      nil,
		},
		{
			Input:  "",
			Output: []Token{{Type: EOF, Position: 0, Len: 1}},
			Error:  ErrEmptyString,
		},
	}

	for _, test := range tests {
		tokenizer, err := NewTokenizer(test.Input)

		if !errors.Is(err, test.Error) {
			t.Error("differing errors between test and function result")
		}

		tokenizedInput := tokenizer.Tokenize()
		if len(tokenizedInput) != len(test.Output) {
			fmt.Println(tokenizedInput)
			t.Errorf(
				`len of tokenizer's Tokenize output: %d
				len of expected Tokenize output: %d\n`,
				len(tokenizedInput),
				len(test.Output),
			)
		} else {
			for j, token := range tokenizedInput {
				if test.Output[j] != token {
					t.Errorf("differing tokens between test and function result")
				}
			}
		}

		if len(test.ExpectedTokenizationErrors) > 0 {
			if len(test.ExpectedTokenizationErrors) != len(tokenizer.Errors) {
				t.Errorf(
					`len of tokenizer's Error output: %d
					len of expected Error output: %d`,
					len(test.ExpectedTokenizationErrors),
					len(tokenizer.Errors),
				)
			} else {
				for i := range tokenizer.Errors {
					if !errors.Is(tokenizer.Errors[i], test.ExpectedTokenizationErrors[i]) {
						t.Errorf("differing errors between test and function result")
					}
				}
			}
		}
	}
}
