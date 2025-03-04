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

// readNumber reads a numeric literal.
func (l *Lexer) readNumber() {
	startPos := l.pos
	for l.pos < len(l.code) && utils.IsDigit(l.peek()) {
		l.pos++
	}
	if l.peek() == '.' {
		l.pos++
		for l.pos < len(l.code) && utils.IsDigit(l.peek()) {
			l.pos++
		}
	}
	value := l.code[startPos:l.pos]
	l.addToken("number", value)
}

// readChar reads a character literal
func (l *Lexer) readChar() {
	l.pos++ // Skip opening quote
	startPos := l.pos
	for l.pos < len(l.code) && l.peek() != '\'' {
		if l.peek() == '\\' {
			l.pos++ // Skip escape character
		}
		l.pos++
	}
	if l.peek() != '\'' {
		l.error("Unclosed character literal")
		return
	}
	value := l.code[startPos:l.pos]
	l.addToken("char", value)
	l.pos++ // Skip closing quote
}

// readString reads a string literal.
func (l *Lexer) readString() {
	l.pos++ // Skip opening quote
	startPos := l.pos
	for l.pos < len(l.code) && l.peek() != '"' {
		if l.peek() == '\\' {
			l.pos++ // Skip escape character
		}
		l.pos++
	}
	if l.peek() != '"' {
		l.error("Unclosed string literal")
		return
	}
	value := l.code[startPos:l.pos]
	l.addToken("string", value)
	l.pos++ // Skip closing quote
}

// readLambda reads lambda expression
func (l *Lexer) readLambda() {
	l.pos++ // Skip '\'
	startPos := l.pos
	for l.pos < len(l.code) && l.peek() != '>' {
		l.pos++
	}
	if l.peek() != '>' {
		l.error("Unclosed lambda expression")
		return
	}
	value := l.code[startPos:l.pos]
	l.addToken("lambda", value)
	l.pos++ // Skip '>'
}

// readOperator reads operator
func (l *Lexer) readOperator() {
	startPos := l.pos
	for l.pos < len(l.code) && utils.IsOperator(l.peek()) {
		l.pos++
	}
	value := l.code[startPos:l.pos]
	l.addToken("operator", value)
}

// readPunctuation reads punctuation construct
func (l *Lexer) readPunctuation() {
	value := string(l.peek())
	l.addToken("punctuation", value)
	l.pos++
}

// readPragma reads GHC extension
func (l *Lexer) readPragma() {
	l.pos += 2 // Skip '{#'
	startPos := l.pos
	for l.pos < len(l.code) && !(l.peek() == '#' && l.peekNext() == '}') {
		l.pos++
	}
	if l.peek() != '#' || l.peekNext() != '}' {
		l.error("Unclosed GHC extension")
		return
	}
	value := l.code[startPos:l.pos]
	l.addToken("pragma", value)
	l.pos += 2 // Skip '#}'
}
