package token

// Type token type
type Type string

const (
	// ILLEGAL illegal code
	ILLEGAL = "ILLEGAL"
	// EOF End of file
	EOF = "EOF"

	// IDENT Identifier
	IDENT = "IDENT"
	// INT number
	INT = "INT"

	// ASSIGN assign sign
	ASSIGN = "="
	// PLUS addition sign
	PLUS = "+"
	// MINUS minus sign
	MINUS = "-"
	// BANG exclamation mark
	BANG = "!"
	// ASTERISK asterisk
	ASTERISK = "*"
	// SLASH slash
	SLASH = "/"

	// LT Lesser Than sign
	LT = "<"
	// GT Greater Than sign
	GT = ">"

	// EQ equal sign
	EQ = "=="
	// NotEQ not equal
	NotEQ = "!="

	// COMMA operator
	COMMA = ","
	// SEMICOLON operator
	SEMICOLON = ";"

	// LPAREN left parenthesis
	LPAREN = "("
	// RPAREN right parenthesis
	RPAREN = ")"
	// LBRACE left brace
	LBRACE = "{"
	// RBRACE right brace
	RBRACE = "}"

	// FUNCTION keyword
	FUNCTION = "FUNCTION"
	// LET keyword
	LET = "LET"
	// TRUE keyword
	TRUE = "TRUE"
	// FALSE keyword
	FALSE = "FALSE"
	// IF keyword
	IF = "IF"
	// ELSE keyword
	ELSE = "ELSE"
	// RETURN keyword
	RETURN = "RETURN"
)

// Token type
type Token struct {
	Type    Type
	Literal string
}

var keywords = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent function
func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
