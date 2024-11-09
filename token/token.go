package token

const (
	ILLEGAL = "ILLEFAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // x, y, add, hello
	INT   = "INT"   // 123456

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	BANG  = "!"
	MINUS = "-"
	SLASH = "/"

	EQ       = "=="
	NOT_EQ   = "!="
	LT       = "<"
	GT       = ">"
	ASTERISK = "*"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	RETURN   = "RETURN"
	IF       = "IF"
	ELSE     = "ELSE"
	FALSE    = "FALSE"
	TRUE     = "TRUE"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
