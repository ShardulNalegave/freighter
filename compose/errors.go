package compose

import "fmt"

type UnknownCharacterLexerError struct {
	Line int
	Char string
}

func (t *UnknownCharacterLexerError) Error() string {
	return fmt.Sprintf("Unknown character ('%s') at line %d", t.Char, t.Line)
}
