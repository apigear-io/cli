// Code generated from pkg/idl/parser/ObjectApi.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 26, 168,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7,
	3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20,
	3, 21, 6, 21, 136, 10, 21, 13, 21, 14, 21, 137, 3, 21, 3, 21, 3, 22, 5,
	22, 143, 10, 22, 3, 22, 6, 22, 146, 10, 22, 13, 22, 14, 22, 147, 3, 23,
	3, 23, 3, 23, 3, 23, 6, 23, 154, 10, 23, 13, 23, 14, 23, 155, 3, 24, 3,
	24, 7, 24, 160, 10, 24, 12, 24, 14, 24, 163, 11, 24, 3, 25, 3, 25, 3, 25,
	3, 25, 2, 2, 26, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19,
	11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37,
	20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 3, 2, 8, 5, 2, 11,
	12, 15, 15, 34, 34, 4, 2, 45, 45, 47, 47, 5, 2, 50, 59, 67, 72, 99, 104,
	5, 2, 67, 92, 97, 97, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 3,
	2, 50, 59, 2, 172, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2,
	2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2,
	2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2,
	2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3,
	2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39,
	3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2,
	47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 3, 51, 3, 2, 2, 2, 5, 58, 3, 2, 2, 2,
	7, 65, 3, 2, 2, 2, 9, 75, 3, 2, 2, 2, 11, 77, 3, 2, 2, 2, 13, 79, 3, 2,
	2, 2, 15, 81, 3, 2, 2, 2, 17, 83, 3, 2, 2, 2, 19, 85, 3, 2, 2, 2, 21, 87,
	3, 2, 2, 2, 23, 94, 3, 2, 2, 2, 25, 101, 3, 2, 2, 2, 27, 106, 3, 2, 2,
	2, 29, 108, 3, 2, 2, 2, 31, 113, 3, 2, 2, 2, 33, 117, 3, 2, 2, 2, 35, 123,
	3, 2, 2, 2, 37, 130, 3, 2, 2, 2, 39, 132, 3, 2, 2, 2, 41, 135, 3, 2, 2,
	2, 43, 142, 3, 2, 2, 2, 45, 149, 3, 2, 2, 2, 47, 157, 3, 2, 2, 2, 49, 164,
	3, 2, 2, 2, 51, 52, 7, 111, 2, 2, 52, 53, 7, 113, 2, 2, 53, 54, 7, 102,
	2, 2, 54, 55, 7, 119, 2, 2, 55, 56, 7, 110, 2, 2, 56, 57, 7, 103, 2, 2,
	57, 4, 3, 2, 2, 2, 58, 59, 7, 107, 2, 2, 59, 60, 7, 111, 2, 2, 60, 61,
	7, 114, 2, 2, 61, 62, 7, 113, 2, 2, 62, 63, 7, 116, 2, 2, 63, 64, 7, 118,
	2, 2, 64, 6, 3, 2, 2, 2, 65, 66, 7, 107, 2, 2, 66, 67, 7, 112, 2, 2, 67,
	68, 7, 118, 2, 2, 68, 69, 7, 103, 2, 2, 69, 70, 7, 116, 2, 2, 70, 71, 7,
	104, 2, 2, 71, 72, 7, 99, 2, 2, 72, 73, 7, 101, 2, 2, 73, 74, 7, 103, 2,
	2, 74, 8, 3, 2, 2, 2, 75, 76, 7, 125, 2, 2, 76, 10, 3, 2, 2, 2, 77, 78,
	7, 127, 2, 2, 78, 12, 3, 2, 2, 2, 79, 80, 7, 60, 2, 2, 80, 14, 3, 2, 2,
	2, 81, 82, 7, 42, 2, 2, 82, 16, 3, 2, 2, 2, 83, 84, 7, 43, 2, 2, 84, 18,
	3, 2, 2, 2, 85, 86, 7, 46, 2, 2, 86, 20, 3, 2, 2, 2, 87, 88, 7, 117, 2,
	2, 88, 89, 7, 107, 2, 2, 89, 90, 7, 105, 2, 2, 90, 91, 7, 112, 2, 2, 91,
	92, 7, 99, 2, 2, 92, 93, 7, 110, 2, 2, 93, 22, 3, 2, 2, 2, 94, 95, 7, 117,
	2, 2, 95, 96, 7, 118, 2, 2, 96, 97, 7, 116, 2, 2, 97, 98, 7, 119, 2, 2,
	98, 99, 7, 101, 2, 2, 99, 100, 7, 118, 2, 2, 100, 24, 3, 2, 2, 2, 101,
	102, 7, 103, 2, 2, 102, 103, 7, 112, 2, 2, 103, 104, 7, 119, 2, 2, 104,
	105, 7, 111, 2, 2, 105, 26, 3, 2, 2, 2, 106, 107, 7, 63, 2, 2, 107, 28,
	3, 2, 2, 2, 108, 109, 7, 100, 2, 2, 109, 110, 7, 113, 2, 2, 110, 111, 7,
	113, 2, 2, 111, 112, 7, 110, 2, 2, 112, 30, 3, 2, 2, 2, 113, 114, 7, 107,
	2, 2, 114, 115, 7, 112, 2, 2, 115, 116, 7, 118, 2, 2, 116, 32, 3, 2, 2,
	2, 117, 118, 7, 104, 2, 2, 118, 119, 7, 110, 2, 2, 119, 120, 7, 113, 2,
	2, 120, 121, 7, 99, 2, 2, 121, 122, 7, 118, 2, 2, 122, 34, 3, 2, 2, 2,
	123, 124, 7, 117, 2, 2, 124, 125, 7, 118, 2, 2, 125, 126, 7, 116, 2, 2,
	126, 127, 7, 107, 2, 2, 127, 128, 7, 112, 2, 2, 128, 129, 7, 105, 2, 2,
	129, 36, 3, 2, 2, 2, 130, 131, 7, 93, 2, 2, 131, 38, 3, 2, 2, 2, 132, 133,
	7, 95, 2, 2, 133, 40, 3, 2, 2, 2, 134, 136, 9, 2, 2, 2, 135, 134, 3, 2,
	2, 2, 136, 137, 3, 2, 2, 2, 137, 135, 3, 2, 2, 2, 137, 138, 3, 2, 2, 2,
	138, 139, 3, 2, 2, 2, 139, 140, 8, 21, 2, 2, 140, 42, 3, 2, 2, 2, 141,
	143, 9, 3, 2, 2, 142, 141, 3, 2, 2, 2, 142, 143, 3, 2, 2, 2, 143, 145,
	3, 2, 2, 2, 144, 146, 4, 50, 59, 2, 145, 144, 3, 2, 2, 2, 146, 147, 3,
	2, 2, 2, 147, 145, 3, 2, 2, 2, 147, 148, 3, 2, 2, 2, 148, 44, 3, 2, 2,
	2, 149, 150, 7, 50, 2, 2, 150, 151, 7, 122, 2, 2, 151, 153, 3, 2, 2, 2,
	152, 154, 9, 4, 2, 2, 153, 152, 3, 2, 2, 2, 154, 155, 3, 2, 2, 2, 155,
	153, 3, 2, 2, 2, 155, 156, 3, 2, 2, 2, 156, 46, 3, 2, 2, 2, 157, 161, 9,
	5, 2, 2, 158, 160, 9, 6, 2, 2, 159, 158, 3, 2, 2, 2, 160, 163, 3, 2, 2,
	2, 161, 159, 3, 2, 2, 2, 161, 162, 3, 2, 2, 2, 162, 48, 3, 2, 2, 2, 163,
	161, 3, 2, 2, 2, 164, 165, 9, 7, 2, 2, 165, 166, 7, 48, 2, 2, 166, 167,
	9, 7, 2, 2, 167, 50, 3, 2, 2, 2, 8, 2, 137, 142, 147, 155, 161, 3, 8, 2,
	2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'module'", "'import'", "'interface'", "'{'", "'}'", "':'", "'('",
	"')'", "','", "'signal'", "'struct'", "'enum'", "'='", "'bool'", "'int'",
	"'float'", "'string'", "'['", "']'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "WHITESPACE", "INTEGER", "HEX", "IDENTIFIER", "VERSION",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "WHITESPACE", "INTEGER", "HEX", "IDENTIFIER", "VERSION",
}

type ObjectApiLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewObjectApiLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *ObjectApiLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewObjectApiLexer(input antlr.CharStream) *ObjectApiLexer {
	l := new(ObjectApiLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "ObjectApi.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ObjectApiLexer tokens.
const (
	ObjectApiLexerT__0       = 1
	ObjectApiLexerT__1       = 2
	ObjectApiLexerT__2       = 3
	ObjectApiLexerT__3       = 4
	ObjectApiLexerT__4       = 5
	ObjectApiLexerT__5       = 6
	ObjectApiLexerT__6       = 7
	ObjectApiLexerT__7       = 8
	ObjectApiLexerT__8       = 9
	ObjectApiLexerT__9       = 10
	ObjectApiLexerT__10      = 11
	ObjectApiLexerT__11      = 12
	ObjectApiLexerT__12      = 13
	ObjectApiLexerT__13      = 14
	ObjectApiLexerT__14      = 15
	ObjectApiLexerT__15      = 16
	ObjectApiLexerT__16      = 17
	ObjectApiLexerT__17      = 18
	ObjectApiLexerT__18      = 19
	ObjectApiLexerWHITESPACE = 20
	ObjectApiLexerINTEGER    = 21
	ObjectApiLexerHEX        = 22
	ObjectApiLexerIDENTIFIER = 23
	ObjectApiLexerVERSION    = 24
)
