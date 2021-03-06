package lexer

import "github.com/elkdanger/monkey/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

// CreateLexer creates a new lexer instance given some input
func CreateLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// NextToken returns the next available token in the stream
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
		break
	case '+':
		tok = newToken(token.PLUS, l.ch)
		break
	case '(':
		tok = newToken(token.LPAREN, l.ch)
		break
	case ')':
		tok = newToken(token.RPAREN, l.ch)
		break
	case '{':
		tok = newToken(token.LBRACE, l.ch)
		break
	case '}':
		tok = newToken(token.RBRACE, l.ch)
		break
	case ',':
		tok = newToken(token.COMMA, l.ch)
		break
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
		break
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NEQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
		break
	case '-':
		tok = newToken(token.MINUS, l.ch)
		break
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
		break
	case '/':
		if l.peekChar() == '/' {
			l.eatRestOfLine()
			return l.NextToken()
		}

		tok = newToken(token.SLASH, l.ch)

		break
	case '<':
		tok = newToken(token.LESSTHAN, l.ch)
		break
	case '>':
		tok = newToken(token.GREATERTHAN, l.ch)
		break

	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		}

		tok = newToken(token.ILLEGAL, l.ch)
	}

	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) peekChar() byte {
	if l.position >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) eatRestOfLine() {
	for l.ch != '\n' {
		l.readChar()
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
