package lexer

// Token represents a lexical token.
type Token struct {
	Type   string
	Value  string
	Line   int
	Column int
}
