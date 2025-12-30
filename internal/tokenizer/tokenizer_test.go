package tokenizer

import (
	"testing"
)

type TestTokenizerStruct struct {
	Input  string
	Output []Token
	Error  error
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
	}

	for _, test := range tests {
		tokenizer, err := NewTokenizer(test.Input)
		if err != test.Error {
			t.Error("differing errors between test and function result")
		}
		tokenizedInput := tokenizer.Tokenize()
		if len(tokenizedInput) != len(test.Output) {
			t.Errorf(`len of tokenizedInput: %d
			len of output: %d\n`, len(tokenizedInput), len(test.Output))
		} else {
			for j, token := range tokenizedInput {
				if test.Output[j] != token {
					t.Errorf("differing tokens between test and function result")
				}
			}
		}

	}
}
