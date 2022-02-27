package lexer

type TokenType int

const (
	TokenUnknown TokenType = iota
	TokenEOF
)

type Token struct {
	Rule   string
	Value  string
	Line   int
	Column int
}

func NewToken(rule, value string, line, column int) *Token {
	return &Token{
		Rule:   rule,
		Value:  value,
		Line:   line,
		Column: column,
	}
}

func NewUnknownToken() *Token {
	return NewToken("Unknown", "", -1, -1)
}
