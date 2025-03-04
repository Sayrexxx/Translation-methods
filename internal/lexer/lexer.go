package lexer

import (
	"fmt"
	"translation_methods/internal/utils"
)

// Lexer represents a lexical analyzer.
type Lexer struct {
	code         string          // Source code
	pos          int             // Current position in the code
	line         int             // Current line
	column       int             // Current column
	Tokens       []Token         // List of tokens
	keywords     map[string]bool // Keywords table
	operators    map[string]bool // Operators table
	punctuations map[string]bool // Punctuations table
}

// NewLexer creates a new Lexer instance.
func NewLexer(code string) *Lexer {
	return &Lexer{
		code:         code,
		pos:          0,
		line:         1,
		column:       1,
		Tokens:       []Token{},
		keywords:     map[string]bool{},
		operators:    map[string]bool{},
		punctuations: map[string]bool{},
	}
}

// addToken adds a token to the list.
func (l *Lexer) addToken(tokenType string, value string) {
	l.Tokens = append(l.Tokens, Token{
		Type:   tokenType,
		Value:  value,
		Line:   l.line,
		Column: l.column,
	})
	l.column += len(value)
}

// peek returns current character
func (l *Lexer) peek() byte {
	if l.pos >= len(l.code) {
		return 0
	}
	return l.code[l.pos]
}

// peekNext return next character
func (l *Lexer) peekNext() byte {
	if l.pos+1 >= len(l.code) {
		return 0
	}
	return l.code[l.pos+1]
}

// error represents critical message
func (l *Lexer) error(message string) {
	fmt.Printf("Error on line %d, column %d: %s\n", l.line, l.column, message)
}

// readIdentifier reads an identifier or keyword.
func (l *Lexer) readIdentifier() {
	startPos := l.pos
	for l.pos < len(l.code) && (utils.IsLetter(l.peek()) || utils.IsDigit(l.peek()) || l.peek() == '\'') {
		l.pos++
	}
	value := l.code[startPos:l.pos]

	// Check if the identifier is a keyword
	if _, ok := l.keywords[value]; ok {
		l.addToken("keyword", value)
	} else {
		l.addToken("identifier", value)
	}
}
