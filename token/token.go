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

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

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

	STRING = "STRING"
	COLON  = ":"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"return": RETURN,
	"false":  FALSE,
	"true":   TRUE,
	"else":   ELSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
