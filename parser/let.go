package parser

import (
	"burpee/ast"
	"burpee/token"
)

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	// After a let we MUST have an identifier
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}

	// the next peek should be an assign, because after an identifier we should have an assignment
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: We're skipping the expressions until we encounter a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}
