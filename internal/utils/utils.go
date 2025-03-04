package utils

import (
	"strings"
	"unicode"
)

// IsDigit checks if a character is a digit.
func IsDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

// IsLetter checks if a character is a letter.
func IsLetter(ch byte) bool {
	return unicode.IsLetter(rune(ch)) || ch == '_'
}

// IsOperator checks if a character is an operator.
func IsOperator(ch byte) bool {
	operators := "+-*/=<>!&|^~:.@"
	return strings.ContainsRune(operators, rune(ch))
}

// IsPunctuation checks if a character is a punctuation.
func IsPunctuation(ch byte) bool {
	punctuations := "(),;[]{}"
	return strings.ContainsRune(punctuations, rune(ch))
}
