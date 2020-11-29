package lexer

import "interpreterlesson/token"

// Lexer - tokenizer
type Lexer struct {
	input        string
	position     int
	readPosition int
	currentChar  byte
}

// New - Create new lexer
func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

func (lexer *Lexer) readChar() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.currentChar = 0
	} else {
		lexer.currentChar = lexer.input[lexer.readPosition]
	}

	lexer.position = lexer.readPosition
	lexer.readPosition++
}

// NextToken - get next token
func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token

	lexer.skipWhitespace()

	switch lexer.currentChar {
	case '=':
		if lexer.peekChar() == '=' {
			ch := lexer.currentChar
			lexer.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(lexer.currentChar)}
		} else {
			tok = newToken(token.ASSIGN, lexer.currentChar)
		}
	case '+':
		tok = newToken(token.PLUS, lexer.currentChar)
	case '(':
		tok = newToken(token.LPAREN, lexer.currentChar)
	case ')':
		tok = newToken(token.RPAREN, lexer.currentChar)
	case '{':
		tok = newToken(token.LBRACE, lexer.currentChar)
	case '}':
		tok = newToken(token.RBRACE, lexer.currentChar)
	case ',':
		tok = newToken(token.COMMA, lexer.currentChar)
	case ';':
		tok = newToken(token.SEMICOLON, lexer.currentChar)
	case '-':
		tok = newToken(token.MINUS, lexer.currentChar)
	case '!':
		if lexer.peekChar() == '=' {
			ch := lexer.currentChar
			lexer.readChar()
			tok = token.Token{Type: token.NOTEQ, Literal: string(ch) + string(lexer.currentChar)}
		} else {
			tok = newToken(token.BANG, lexer.currentChar)
		}
	case '/':
		tok = newToken(token.SLASH, lexer.currentChar)
	case '*':
		tok = newToken(token.ASTERISK, lexer.currentChar)
	case '<':
		tok = newToken(token.LT, lexer.currentChar)
	case '>':
		tok = newToken(token.GT, lexer.currentChar)

	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(lexer.currentChar) {
			tok.Literal = lexer.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isDigit(lexer.currentChar) {
			tok.Type = token.INT
			tok.Literal = lexer.readNumber()
			return tok
		}
		tok = newToken(token.ILLEGAL, lexer.currentChar)

	}

	lexer.readChar()
	return tok
}

func (lexer *Lexer) peekChar() byte {
	if lexer.readPosition >= len(lexer.input) {
		return 0
	}
	return lexer.input[lexer.readPosition]

}

func (lexer *Lexer) skipWhitespace() {
	for lexer.currentChar == ' ' || lexer.currentChar == '\t' || lexer.currentChar == '\n' || lexer.currentChar == '\r' {
		lexer.readChar()
	}
}

func (lexer *Lexer) readNumber() string {
	position := lexer.position
	for isDigit(lexer.currentChar) {
		lexer.readChar()
	}

	return lexer.input[position:lexer.position]
}

func (lexer *Lexer) readIdentifier() string {
	position := lexer.position

	for isLetter(lexer.currentChar) {
		lexer.readChar()
	}

	return lexer.input[position:lexer.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && 'z' >= ch || 'A' <= ch && 'Z' >= ch || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType token.Type, currentChar byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(currentChar)}
}
