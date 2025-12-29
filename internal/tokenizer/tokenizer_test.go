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
				{Type: INT, Literal: "10", Position: 0},
				{Type: PLUS, Literal: "", Position: 3},
				{Type: INT, Literal: "55", Position: 5},
				{Type: EOF, Literal: "", Position: 7},
			},
			Error: nil,
		},
		{
			Input: "10.3 - 4 ",
			Output: []Token{
				{Type: FLOAT, Literal: "10.3", Position: 0},
				{Type: MINUS, Literal: "", Position: 5},
				{Type: INT, Literal: "4", Position: 7},
				{Type: EOF, Literal: "", Position: 9},
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
