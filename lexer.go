package lexer

type Lexer struct {
	rules     Rules
	buf       []byte
	tokens    []*Token
	nextToken int
}

func New(rules Rules) (*Lexer, error) {
	if err := rules.Init(); err != nil {
		return nil, err
	}
	l := &Lexer{
		rules:     rules,
		tokens:    make([]*Token, 0, 128),
		nextToken: 0,
	}
	return l, nil
}

func (l *Lexer) ParseString(str string) {
	l.Parse([]byte(str))
}

func (l *Lexer) Parse(buf []byte) {
	l.buf = buf
	l.tokens = l.tokens[:0]
	l.nextToken = 0

	line, column := 1, 1

	iStart := 0
	for iStart < len(l.buf) {
		stack := l.buf[iStart:]
		r, result := l.rules.Find(stack)
		if result != nil && r != nil {
			for _, b := range result {
				switch b {
				case '\n':
					line++
					column = 1
				default:
					column++
				}
			}

			if !r.DeleteToken {
				t := NewToken(
					r.Name,
					string(r.Replace(result)),
					line, column,
				)
				l.tokens = append(l.tokens, t)
			}
			iStart += len(result)
		} else {
			var result byte
			for _, b := range stack {
				iStart++
				switch b {
				case '\n':
					line++
					column = 1
				case '\t', '\r', ' ':
					line++
				default:
					result = b
				}
				if result != 0 {
					break
				}
			}

			t := NewToken(
				"Unknown",
				string(result),
				line, column,
			)
			l.tokens = append(l.tokens, t)
		}
	}
	l.tokens = append(l.tokens, NewToken(
		"EOF",
		"",
		line, column,
	))
}

func (l *Lexer) NextToken() *Token {
	if l.nextToken >= len(l.tokens) {
		return nil
	}
	t := l.tokens[l.nextToken]
	l.nextToken++
	return t
}
