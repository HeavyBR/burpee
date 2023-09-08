package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x ,y...
	INT   = "INT"   // 1, 2, 3, 4...

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

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]Type{
	"fn":  FUNCTION,
	"let": LET,
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
