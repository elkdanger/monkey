package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	IDENT   = "IDENTIFIER"
	INT     = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"
	EQ     = "=="
	NEQ    = "!="

	// Delimiters
	COMMA       = ","
	SEMICOLON   = ";"
	LPAREN      = "("
	RPAREN      = ")"
	LBRACE      = "{"
	RBRACE      = "}"
	BANG        = "!"
	MINUS       = "-"
	SLASH       = "/"
	ASTERISK    = "*"
	LESSTHAN    = "<"
	GREATERTHAN = ">"
	FUNCTION    = "FUNCTION"
	LET         = "LET"
	IF          = "IF"
	RETURN      = "RETURN"
	ELSE        = "ELSE"
	TRUE        = "TRUE"
	FALSE       = "FALSE"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
