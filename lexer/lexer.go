package lexer

import "burpee/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	// Skip whitespace
	l.skipWhitespace()

	switch l.ch {
	case '=':
		// Check if the next character is an equal sign
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			// If so, we have an equality operator
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}

	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '!':
		// Check if the next character is an equal sign
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			// If so, we have a not-equal operator
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.Type, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// readNumber reads in a number and advances the lexer's position until it encounters a non-digit character.
// e.g., 12345 -> 12345
func (l *Lexer) readNumber() string {
	position := l.position

	// Read until we hit a non-digit character
	for isDigit(l.ch) {
		l.readChar()
	}

	// Return the number
	return l.input[position:l.position]
}

// readIdentifier reads in an identifier and advances the lexer's position until it encounters a non-letter-character.
// e.g. let five = 5; -> let and five are identifiers
func (l *Lexer) readIdentifier() string {
	position := l.position

	// Read until we hit a non-letter character
	for isLetter(l.ch) {
		l.readChar()
	}

	// Return the identifier
	return l.input[position:l.position]
}

// peekChar returns the next character in the input without advancing the lexer's position.
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

// readChar advances the lexer's position and reads the next character in the input.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

// isDigit checks if a character is a digit.
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// isLetter checks if a character is a letter or underscore.
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// skipWhitespace advances the lexer's position until it encounters a non-whitespace character.
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		// Advance the lexer's position
		l.readChar()
	}
}
