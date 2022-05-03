package idl

import (
	"objectapi/pkg/idl/parser"
	"objectapi/pkg/model"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Parser defines the parser data
type Parser struct {
	System *model.System
}

// NewParser creates a new parser with a named system
func NewParser(s *model.System) *Parser {
	return &Parser{
		System: s,
	}
}

// ParseFile parses a file containing idl document
func (p *Parser) ParseFile(file string) error {
	input, err := antlr.NewFileStream(file)
	if err != nil {
		return err
	}
	return p.parseStream(input)
}

// ParseString parses a string containing idl document
func (p *Parser) ParseString(str string) error {
	input := antlr.NewInputStream(str)
	return p.parseStream(input)
}

// parse idl from antlr file stream
func (p *Parser) parseStream(input antlr.CharStream) error {
	// create the lexer
	lexer := parser.NewObjectApiLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// create the parser
	parser := parser.NewObjectApiParser(tokens)
	listener := NewObjectApiListener(p.System)
	start := parser.DocumentRule()
	antlr.ParseTreeWalkerDefault.Walk(listener, start)
	return nil
}
