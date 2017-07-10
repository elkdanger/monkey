package parser

import (
	"github.com/elkdanger/monkey/ast"
	"github.com/elkdanger/monkey/lexer"
	"github.com/elkdanger/monkey/token"
)

type Parser struct {
	l *lexer.Lexer

	currentToken token.Token
	peekToken    token.Token
}

// New creates a new parser, given a lexer
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram parses the input from the given lexer
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
