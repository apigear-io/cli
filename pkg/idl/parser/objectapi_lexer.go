// Code generated from pkg/idl/parser/ObjectApi.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type ObjectApiLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var objectapilexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func objectapilexerLexerInit() {
	staticData := &objectapilexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'module'", "'import'", "'interface'", "'{'", "'}'", "':'", "'('",
		"')'", "','", "'signal'", "'struct'", "'enum'", "'='", "'['", "']'",
		"'bool'", "'int'", "'float'", "'string'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "WHITESPACE", "INTEGER", "HEX", "TYPE_IDENTIFIER", "IDENTIFIER",
		"VERSION",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "WHITESPACE", "INTEGER", "HEX", "TYPE_IDENTIFIER",
		"IDENTIFIER", "VERSION",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 25, 179, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1,
		14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 19, 4, 19, 136, 8, 19, 11, 19, 12, 19, 137, 1, 19,
		1, 19, 1, 20, 3, 20, 143, 8, 20, 1, 20, 4, 20, 146, 8, 20, 11, 20, 12,
		20, 147, 1, 21, 1, 21, 1, 21, 1, 21, 4, 21, 154, 8, 21, 11, 21, 12, 21,
		155, 1, 22, 1, 22, 5, 22, 160, 8, 22, 10, 22, 12, 22, 163, 9, 22, 1, 22,
		1, 22, 3, 22, 167, 8, 22, 1, 23, 1, 23, 5, 23, 171, 8, 23, 10, 23, 12,
		23, 174, 9, 23, 1, 24, 1, 24, 1, 24, 1, 24, 0, 0, 25, 1, 1, 3, 2, 5, 3,
		7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13,
		27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22,
		45, 23, 47, 24, 49, 25, 1, 0, 8, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 43,
		43, 45, 45, 3, 0, 48, 57, 65, 70, 97, 102, 2, 0, 65, 90, 95, 95, 3, 0,
		48, 57, 65, 90, 95, 95, 3, 0, 65, 90, 95, 95, 97, 122, 5, 0, 46, 46, 48,
		57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 185, 0, 1, 1, 0, 0, 0, 0, 3,
		1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11,
		1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0,
		19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0,
		0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0,
		0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0,
		0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1,
		0, 0, 0, 1, 51, 1, 0, 0, 0, 3, 58, 1, 0, 0, 0, 5, 65, 1, 0, 0, 0, 7, 75,
		1, 0, 0, 0, 9, 77, 1, 0, 0, 0, 11, 79, 1, 0, 0, 0, 13, 81, 1, 0, 0, 0,
		15, 83, 1, 0, 0, 0, 17, 85, 1, 0, 0, 0, 19, 87, 1, 0, 0, 0, 21, 94, 1,
		0, 0, 0, 23, 101, 1, 0, 0, 0, 25, 106, 1, 0, 0, 0, 27, 108, 1, 0, 0, 0,
		29, 110, 1, 0, 0, 0, 31, 112, 1, 0, 0, 0, 33, 117, 1, 0, 0, 0, 35, 121,
		1, 0, 0, 0, 37, 127, 1, 0, 0, 0, 39, 135, 1, 0, 0, 0, 41, 142, 1, 0, 0,
		0, 43, 149, 1, 0, 0, 0, 45, 157, 1, 0, 0, 0, 47, 168, 1, 0, 0, 0, 49, 175,
		1, 0, 0, 0, 51, 52, 5, 109, 0, 0, 52, 53, 5, 111, 0, 0, 53, 54, 5, 100,
		0, 0, 54, 55, 5, 117, 0, 0, 55, 56, 5, 108, 0, 0, 56, 57, 5, 101, 0, 0,
		57, 2, 1, 0, 0, 0, 58, 59, 5, 105, 0, 0, 59, 60, 5, 109, 0, 0, 60, 61,
		5, 112, 0, 0, 61, 62, 5, 111, 0, 0, 62, 63, 5, 114, 0, 0, 63, 64, 5, 116,
		0, 0, 64, 4, 1, 0, 0, 0, 65, 66, 5, 105, 0, 0, 66, 67, 5, 110, 0, 0, 67,
		68, 5, 116, 0, 0, 68, 69, 5, 101, 0, 0, 69, 70, 5, 114, 0, 0, 70, 71, 5,
		102, 0, 0, 71, 72, 5, 97, 0, 0, 72, 73, 5, 99, 0, 0, 73, 74, 5, 101, 0,
		0, 74, 6, 1, 0, 0, 0, 75, 76, 5, 123, 0, 0, 76, 8, 1, 0, 0, 0, 77, 78,
		5, 125, 0, 0, 78, 10, 1, 0, 0, 0, 79, 80, 5, 58, 0, 0, 80, 12, 1, 0, 0,
		0, 81, 82, 5, 40, 0, 0, 82, 14, 1, 0, 0, 0, 83, 84, 5, 41, 0, 0, 84, 16,
		1, 0, 0, 0, 85, 86, 5, 44, 0, 0, 86, 18, 1, 0, 0, 0, 87, 88, 5, 115, 0,
		0, 88, 89, 5, 105, 0, 0, 89, 90, 5, 103, 0, 0, 90, 91, 5, 110, 0, 0, 91,
		92, 5, 97, 0, 0, 92, 93, 5, 108, 0, 0, 93, 20, 1, 0, 0, 0, 94, 95, 5, 115,
		0, 0, 95, 96, 5, 116, 0, 0, 96, 97, 5, 114, 0, 0, 97, 98, 5, 117, 0, 0,
		98, 99, 5, 99, 0, 0, 99, 100, 5, 116, 0, 0, 100, 22, 1, 0, 0, 0, 101, 102,
		5, 101, 0, 0, 102, 103, 5, 110, 0, 0, 103, 104, 5, 117, 0, 0, 104, 105,
		5, 109, 0, 0, 105, 24, 1, 0, 0, 0, 106, 107, 5, 61, 0, 0, 107, 26, 1, 0,
		0, 0, 108, 109, 5, 91, 0, 0, 109, 28, 1, 0, 0, 0, 110, 111, 5, 93, 0, 0,
		111, 30, 1, 0, 0, 0, 112, 113, 5, 98, 0, 0, 113, 114, 5, 111, 0, 0, 114,
		115, 5, 111, 0, 0, 115, 116, 5, 108, 0, 0, 116, 32, 1, 0, 0, 0, 117, 118,
		5, 105, 0, 0, 118, 119, 5, 110, 0, 0, 119, 120, 5, 116, 0, 0, 120, 34,
		1, 0, 0, 0, 121, 122, 5, 102, 0, 0, 122, 123, 5, 108, 0, 0, 123, 124, 5,
		111, 0, 0, 124, 125, 5, 97, 0, 0, 125, 126, 5, 116, 0, 0, 126, 36, 1, 0,
		0, 0, 127, 128, 5, 115, 0, 0, 128, 129, 5, 116, 0, 0, 129, 130, 5, 114,
		0, 0, 130, 131, 5, 105, 0, 0, 131, 132, 5, 110, 0, 0, 132, 133, 5, 103,
		0, 0, 133, 38, 1, 0, 0, 0, 134, 136, 7, 0, 0, 0, 135, 134, 1, 0, 0, 0,
		136, 137, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138,
		139, 1, 0, 0, 0, 139, 140, 6, 19, 0, 0, 140, 40, 1, 0, 0, 0, 141, 143,
		7, 1, 0, 0, 142, 141, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 145, 1, 0,
		0, 0, 144, 146, 2, 48, 57, 0, 145, 144, 1, 0, 0, 0, 146, 147, 1, 0, 0,
		0, 147, 145, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 42, 1, 0, 0, 0, 149,
		150, 5, 48, 0, 0, 150, 151, 5, 120, 0, 0, 151, 153, 1, 0, 0, 0, 152, 154,
		7, 2, 0, 0, 153, 152, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 153, 1, 0,
		0, 0, 155, 156, 1, 0, 0, 0, 156, 44, 1, 0, 0, 0, 157, 161, 7, 3, 0, 0,
		158, 160, 7, 4, 0, 0, 159, 158, 1, 0, 0, 0, 160, 163, 1, 0, 0, 0, 161,
		159, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162, 166, 1, 0, 0, 0, 163, 161,
		1, 0, 0, 0, 164, 165, 5, 91, 0, 0, 165, 167, 5, 93, 0, 0, 166, 164, 1,
		0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 46, 1, 0, 0, 0, 168, 172, 7, 5, 0,
		0, 169, 171, 7, 6, 0, 0, 170, 169, 1, 0, 0, 0, 171, 174, 1, 0, 0, 0, 172,
		170, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 48, 1, 0, 0, 0, 174, 172, 1,
		0, 0, 0, 175, 176, 7, 7, 0, 0, 176, 177, 5, 46, 0, 0, 177, 178, 7, 7, 0,
		0, 178, 50, 1, 0, 0, 0, 8, 0, 137, 142, 147, 155, 161, 166, 172, 1, 6,
		0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ObjectApiLexerInit initializes any static state used to implement ObjectApiLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewObjectApiLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func ObjectApiLexerInit() {
	staticData := &objectapilexerLexerStaticData
	staticData.once.Do(objectapilexerLexerInit)
}

// NewObjectApiLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewObjectApiLexer(input antlr.CharStream) *ObjectApiLexer {
	ObjectApiLexerInit()
	l := new(ObjectApiLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &objectapilexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "ObjectApi.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ObjectApiLexer tokens.
const (
	ObjectApiLexerT__0            = 1
	ObjectApiLexerT__1            = 2
	ObjectApiLexerT__2            = 3
	ObjectApiLexerT__3            = 4
	ObjectApiLexerT__4            = 5
	ObjectApiLexerT__5            = 6
	ObjectApiLexerT__6            = 7
	ObjectApiLexerT__7            = 8
	ObjectApiLexerT__8            = 9
	ObjectApiLexerT__9            = 10
	ObjectApiLexerT__10           = 11
	ObjectApiLexerT__11           = 12
	ObjectApiLexerT__12           = 13
	ObjectApiLexerT__13           = 14
	ObjectApiLexerT__14           = 15
	ObjectApiLexerT__15           = 16
	ObjectApiLexerT__16           = 17
	ObjectApiLexerT__17           = 18
	ObjectApiLexerT__18           = 19
	ObjectApiLexerWHITESPACE      = 20
	ObjectApiLexerINTEGER         = 21
	ObjectApiLexerHEX             = 22
	ObjectApiLexerTYPE_IDENTIFIER = 23
	ObjectApiLexerIDENTIFIER      = 24
	ObjectApiLexerVERSION         = 25
)
