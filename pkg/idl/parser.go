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
func (p *Parser) ParseFromFile(file string) error {
	input, err := antlr.NewFileStream(file)
	if err != nil {
		return err
	}
	return p.parseFromStream(input)
}

func (p *Parser) ParseFromString(str string) error {
	input := antlr.NewInputStream(str)
	return p.parseFromStream(input)
}

// parse idl from antlr file stream
func (p *Parser) parseFromStream(input antlr.CharStream) error {
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

func NewParser(name string) *Parser {
	return &Parser{
		System: model.NewSystem(name),
	}
}
