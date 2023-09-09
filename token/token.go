package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x ,y...
	INT   = "INT"   // 1, 2, 3, 4...

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	BANG     = "!"
	MINUS    = "-"
	SLASH    = "/"
	ASTERISK = "*"

	LT     = "<"
	GT     = ">"
	EQ     = "=="
	NOT_EQ = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	FALSE    = "FALSE"
	TRUE     = "TRUE"
)

var keywords = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"false":  FALSE,
	"true":   TRUE,
}

type Type string

type Token struct {
	Type    Type
	Literal string
}

func LookupIdent(ident string) Type {
	// Check if the identifier is a keyword
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	// Otherwise, it's just an identifier
	return IDENT
}
