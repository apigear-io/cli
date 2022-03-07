package idl

import (
	"objectapi/pkg/idl/parser"
	"objectapi/pkg/model"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Parser struct {
	System *model.System
}

// parse idl from file
func (p *Parser) ParseFile(file string) error {
	input, err := antlr.NewFileStream(file)
	if err != nil {
		return err
	}
	return p.ParseStream(input)
}

func (p *Parser) ParseString(str string) error {
	input := antlr.NewInputStream(str)
	return p.ParseStream(input)
}

// parse idl from antlr file stream
func (p *Parser) ParseStream(input antlr.CharStream) error {
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

func NewIDLParser(s *model.System) *Parser {
	return &Parser{
		System: s,
	}
}
