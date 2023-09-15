package parser

import (
	"testing"

	"burpee/ast"

	"burpee/lexer"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLetStatements(t *testing.T) {
	input := `
		let x = 5;
     	let y = 10;
		let foobar = 838383;

		let err - 123;
    `

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	assert.NotNil(t, program)
	require.Len(t, p.Errors(), 0)

	// We have three let statements in the input
	assert.Len(t, program.Statements, 3)

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		testLetStatement(t, stmt, tt.expectedIdentifier)
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) {
	assert.Equal(t, s.TokenLiteral(), "let")
	assert.IsType(t, &ast.LetStatement{}, s)

	letStmt, ok := s.(*ast.LetStatement)
	require.True(t, ok)

	assert.Equal(t, name, letStmt.Name.Value)
	assert.Equal(t, name, letStmt.Name.TokenLiteral())
}
