package compose

type TokenType string

const (
	IDENT_TokenType  TokenType = "IDENT"
	STRING_TokenType TokenType = "STRING"
	NUMBER_TokenType TokenType = "NUMBER"

	LISTEN_AT_TokenType   TokenType = "LISTEN-AT"
	SPAWN_TokenType       TokenType = "SPAWN"
	ADD_BACKEND_TokenType TokenType = "ADD-BACKEND"

	SEMICOLON_TokenType TokenType = "SEMICOLON"

	EOF_TokenType TokenType = "END-OF-FILE"
)

type Token struct {
	Type    TokenType
	Line    int
	Lexeme  string
	Literal interface{}
}
