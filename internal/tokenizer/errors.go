package tokenizer

import "errors"

var (
	ErrMultipleDecimalPoints = errors.New("number with multiple decimal points")
	ErrIllegalCharacter      = errors.New("")
	ErrEmptyString           = errors.New("tokenizer cannot be created for an empty string")
)
