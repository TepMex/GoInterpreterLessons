package token

// Type - type of token
type Type string

// Token for lexer
type Token struct {
	Type    Type
	Literal string
}

// tokens
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	IDENT   = "IDENT"
	INT     = "INT"

	ASSIGN    = "="
	PLUS      = "+"
	COMMA     = ","
	SEMICOLON = ";"
	BANG      = "!"
	MINUS     = "-"
	SLASH     = "/"
	ASTERISK  = "*"
	LT        = "<"
	GT        = ">"
	EQ        = "=="
	NOTEQ     = "!="

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	IF       = "IF"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	ELSE     = "ELSE"
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
	"else":   ELSE,
}

// LookupIdentifier - looking up identifiers
func LookupIdentifier(identifier string) Type {
	if token, ok := keywords[identifier]; ok {
		return token
	}
	return IDENT
}
