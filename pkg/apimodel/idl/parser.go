package idl

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/foundation"
	"github.com/apigear-io/cli/pkg/apimodel/idl/parser"
	"github.com/apigear-io/cli/pkg/foundation/logging"
	"github.com/apigear-io/cli/pkg/apimodel"

	"github.com/antlr4-go/antlr/v4"
)

// Parser defines the parser data
type Parser struct {
	System *apimodel.System
}

// NewParser creates a new parser with a named system
func NewParser(s *apimodel.System) *Parser {
	return &Parser{
		System: s,
	}
}

// TODO: ParseFile is called 3 times (e.g. during solution check, run solution and ...)
// ParseFile parses a file containing idl document
func (p *Parser) ParseFile(file string) error {
	if !foundation.IsFile(file) {
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
	logging.Info().Msgf("parse idl from input stream")
	lexer := parser.NewObjectApiLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// create the parser
	parser := parser.NewObjectApiParser(tokens)
	listener := NewObjectApiListener(p.System)
	start := parser.DocumentRule()
	antlr.ParseTreeWalkerDefault.Walk(listener, start)
	return nil
}
