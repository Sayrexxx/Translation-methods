package lexer

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
