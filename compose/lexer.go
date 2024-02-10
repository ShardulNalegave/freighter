package compose

type Lexer struct {
	line    int
	current int
	source  string
	tokens  []*Token
}

func (l *Lexer) ScanTokens() ([]*Token, error) {
	for l.current < len(l.source) {
		c := l.advance()
		switch c {
		case ";":
			l.tokens = append(l.tokens, &Token{
				Line:    l.line,
				Type:    SEMICOLON_TokenType,
				Literal: ";",
				Lexeme:  ";",
			})
		case " ":
		case "\t":
		case "\n":
			l.line += 1
		default:
			return nil, &UnknownCharacterLexerError{
				Char: c,
				Line: l.line,
			}
		}
	}

	l.tokens = append(l.tokens, &Token{
		Line: l.line,
		Type: EOF_TokenType,
	})

	return l.tokens, nil
}

func (l *Lexer) advance() string {
	l.current += 1
	return string(l.source[l.current-1])
}

func (l *Lexer) peek() string {
	return string(l.source[l.current])
}

func NewLexer(source string) *Lexer {
	return &Lexer{
		source:  source,
		current: 0,
		line:    1,
		tokens:  make([]*Token, 0),
	}
}
