package idl

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/apigear-io/cli/pkg/idl/parser"
	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/model"

	"github.com/antlr4-go/antlr/v4"
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

// TODO: ParseFile is called 3 times (e.g. during solution check, run solution and ...)
// ParseFile parses a file containing idl document
func (p *Parser) ParseFile(file string) error {
	if !helper.IsFile(file) {
		return fmt.Errorf("file %s does not exist", file)
	}

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
	log.Info().Msgf("parse idl from input stream")
	lexer := parser.NewObjectApiLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// create the parser
	parser := parser.NewObjectApiParser(tokens)
	listener := NewObjectApiListener(p.System)
	start := parser.DocumentRule()
	antlr.ParseTreeWalkerDefault.Walk(listener, start)
	return nil
}
