// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
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
		"'bool'", "'int'", "'int32'", "'int64'", "'float'", "'float32'", "'float64'",
		"'string'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "WHITESPACE", "INTEGER", "HEX", "TYPE_IDENTIFIER",
		"IDENTIFIER", "VERSION", "DOCLINE", "TAGLINE", "COMMENT",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "WHITESPACE",
		"INTEGER", "HEX", "TYPE_IDENTIFIER", "IDENTIFIER", "VERSION", "DOCLINE",
		"TAGLINE", "COMMENT",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 32, 246, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7,
		1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 23, 4, 23, 178, 8, 23, 11, 23, 12, 23, 179, 1, 23,
		1, 23, 1, 24, 3, 24, 185, 8, 24, 1, 24, 4, 24, 188, 8, 24, 11, 24, 12,
		24, 189, 1, 25, 1, 25, 1, 25, 1, 25, 4, 25, 196, 8, 25, 11, 25, 12, 25,
		197, 1, 26, 1, 26, 5, 26, 202, 8, 26, 10, 26, 12, 26, 205, 9, 26, 1, 26,
		1, 26, 3, 26, 209, 8, 26, 1, 27, 1, 27, 5, 27, 213, 8, 27, 10, 27, 12,
		27, 216, 9, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29,
		5, 29, 226, 8, 29, 10, 29, 12, 29, 229, 9, 29, 1, 30, 1, 30, 5, 30, 233,
		8, 30, 10, 30, 12, 30, 236, 9, 30, 1, 31, 1, 31, 5, 31, 240, 8, 31, 10,
		31, 12, 31, 243, 9, 31, 1, 31, 1, 31, 0, 0, 32, 1, 1, 3, 2, 5, 3, 7, 4,
		9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
		29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23,
		47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32,
		1, 0, 9, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 43, 43, 45, 45, 3, 0, 48, 57,
		65, 70, 97, 102, 2, 0, 65, 90, 95, 95, 3, 0, 48, 57, 65, 90, 95, 95, 3,
		0, 65, 90, 95, 95, 97, 122, 5, 0, 46, 46, 48, 57, 65, 90, 95, 95, 97, 122,
		1, 0, 48, 57, 2, 0, 10, 10, 13, 13, 255, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0,
		0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0,
		0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1,
		0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27,
		1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0,
		35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0,
		0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0,
		0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0,
		0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 1, 65, 1,
		0, 0, 0, 3, 72, 1, 0, 0, 0, 5, 79, 1, 0, 0, 0, 7, 89, 1, 0, 0, 0, 9, 91,
		1, 0, 0, 0, 11, 93, 1, 0, 0, 0, 13, 95, 1, 0, 0, 0, 15, 97, 1, 0, 0, 0,
		17, 99, 1, 0, 0, 0, 19, 101, 1, 0, 0, 0, 21, 108, 1, 0, 0, 0, 23, 115,
		1, 0, 0, 0, 25, 120, 1, 0, 0, 0, 27, 122, 1, 0, 0, 0, 29, 124, 1, 0, 0,
		0, 31, 126, 1, 0, 0, 0, 33, 131, 1, 0, 0, 0, 35, 135, 1, 0, 0, 0, 37, 141,
		1, 0, 0, 0, 39, 147, 1, 0, 0, 0, 41, 153, 1, 0, 0, 0, 43, 161, 1, 0, 0,
		0, 45, 169, 1, 0, 0, 0, 47, 177, 1, 0, 0, 0, 49, 184, 1, 0, 0, 0, 51, 191,
		1, 0, 0, 0, 53, 199, 1, 0, 0, 0, 55, 210, 1, 0, 0, 0, 57, 217, 1, 0, 0,
		0, 59, 221, 1, 0, 0, 0, 61, 230, 1, 0, 0, 0, 63, 237, 1, 0, 0, 0, 65, 66,
		5, 109, 0, 0, 66, 67, 5, 111, 0, 0, 67, 68, 5, 100, 0, 0, 68, 69, 5, 117,
		0, 0, 69, 70, 5, 108, 0, 0, 70, 71, 5, 101, 0, 0, 71, 2, 1, 0, 0, 0, 72,
		73, 5, 105, 0, 0, 73, 74, 5, 109, 0, 0, 74, 75, 5, 112, 0, 0, 75, 76, 5,
		111, 0, 0, 76, 77, 5, 114, 0, 0, 77, 78, 5, 116, 0, 0, 78, 4, 1, 0, 0,
		0, 79, 80, 5, 105, 0, 0, 80, 81, 5, 110, 0, 0, 81, 82, 5, 116, 0, 0, 82,
		83, 5, 101, 0, 0, 83, 84, 5, 114, 0, 0, 84, 85, 5, 102, 0, 0, 85, 86, 5,
		97, 0, 0, 86, 87, 5, 99, 0, 0, 87, 88, 5, 101, 0, 0, 88, 6, 1, 0, 0, 0,
		89, 90, 5, 123, 0, 0, 90, 8, 1, 0, 0, 0, 91, 92, 5, 125, 0, 0, 92, 10,
		1, 0, 0, 0, 93, 94, 5, 58, 0, 0, 94, 12, 1, 0, 0, 0, 95, 96, 5, 40, 0,
		0, 96, 14, 1, 0, 0, 0, 97, 98, 5, 41, 0, 0, 98, 16, 1, 0, 0, 0, 99, 100,
		5, 44, 0, 0, 100, 18, 1, 0, 0, 0, 101, 102, 5, 115, 0, 0, 102, 103, 5,
		105, 0, 0, 103, 104, 5, 103, 0, 0, 104, 105, 5, 110, 0, 0, 105, 106, 5,
		97, 0, 0, 106, 107, 5, 108, 0, 0, 107, 20, 1, 0, 0, 0, 108, 109, 5, 115,
		0, 0, 109, 110, 5, 116, 0, 0, 110, 111, 5, 114, 0, 0, 111, 112, 5, 117,
		0, 0, 112, 113, 5, 99, 0, 0, 113, 114, 5, 116, 0, 0, 114, 22, 1, 0, 0,
		0, 115, 116, 5, 101, 0, 0, 116, 117, 5, 110, 0, 0, 117, 118, 5, 117, 0,
		0, 118, 119, 5, 109, 0, 0, 119, 24, 1, 0, 0, 0, 120, 121, 5, 61, 0, 0,
		121, 26, 1, 0, 0, 0, 122, 123, 5, 91, 0, 0, 123, 28, 1, 0, 0, 0, 124, 125,
		5, 93, 0, 0, 125, 30, 1, 0, 0, 0, 126, 127, 5, 98, 0, 0, 127, 128, 5, 111,
		0, 0, 128, 129, 5, 111, 0, 0, 129, 130, 5, 108, 0, 0, 130, 32, 1, 0, 0,
		0, 131, 132, 5, 105, 0, 0, 132, 133, 5, 110, 0, 0, 133, 134, 5, 116, 0,
		0, 134, 34, 1, 0, 0, 0, 135, 136, 5, 105, 0, 0, 136, 137, 5, 110, 0, 0,
		137, 138, 5, 116, 0, 0, 138, 139, 5, 51, 0, 0, 139, 140, 5, 50, 0, 0, 140,
		36, 1, 0, 0, 0, 141, 142, 5, 105, 0, 0, 142, 143, 5, 110, 0, 0, 143, 144,
		5, 116, 0, 0, 144, 145, 5, 54, 0, 0, 145, 146, 5, 52, 0, 0, 146, 38, 1,
		0, 0, 0, 147, 148, 5, 102, 0, 0, 148, 149, 5, 108, 0, 0, 149, 150, 5, 111,
		0, 0, 150, 151, 5, 97, 0, 0, 151, 152, 5, 116, 0, 0, 152, 40, 1, 0, 0,
		0, 153, 154, 5, 102, 0, 0, 154, 155, 5, 108, 0, 0, 155, 156, 5, 111, 0,
		0, 156, 157, 5, 97, 0, 0, 157, 158, 5, 116, 0, 0, 158, 159, 5, 51, 0, 0,
		159, 160, 5, 50, 0, 0, 160, 42, 1, 0, 0, 0, 161, 162, 5, 102, 0, 0, 162,
		163, 5, 108, 0, 0, 163, 164, 5, 111, 0, 0, 164, 165, 5, 97, 0, 0, 165,
		166, 5, 116, 0, 0, 166, 167, 5, 54, 0, 0, 167, 168, 5, 52, 0, 0, 168, 44,
		1, 0, 0, 0, 169, 170, 5, 115, 0, 0, 170, 171, 5, 116, 0, 0, 171, 172, 5,
		114, 0, 0, 172, 173, 5, 105, 0, 0, 173, 174, 5, 110, 0, 0, 174, 175, 5,
		103, 0, 0, 175, 46, 1, 0, 0, 0, 176, 178, 7, 0, 0, 0, 177, 176, 1, 0, 0,
		0, 178, 179, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180,
		181, 1, 0, 0, 0, 181, 182, 6, 23, 0, 0, 182, 48, 1, 0, 0, 0, 183, 185,
		7, 1, 0, 0, 184, 183, 1, 0, 0, 0, 184, 185, 1, 0, 0, 0, 185, 187, 1, 0,
		0, 0, 186, 188, 2, 48, 57, 0, 187, 186, 1, 0, 0, 0, 188, 189, 1, 0, 0,
		0, 189, 187, 1, 0, 0, 0, 189, 190, 1, 0, 0, 0, 190, 50, 1, 0, 0, 0, 191,
		192, 5, 48, 0, 0, 192, 193, 5, 120, 0, 0, 193, 195, 1, 0, 0, 0, 194, 196,
		7, 2, 0, 0, 195, 194, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0, 197, 195, 1, 0,
		0, 0, 197, 198, 1, 0, 0, 0, 198, 52, 1, 0, 0, 0, 199, 203, 7, 3, 0, 0,
		200, 202, 7, 4, 0, 0, 201, 200, 1, 0, 0, 0, 202, 205, 1, 0, 0, 0, 203,
		201, 1, 0, 0, 0, 203, 204, 1, 0, 0, 0, 204, 208, 1, 0, 0, 0, 205, 203,
		1, 0, 0, 0, 206, 207, 5, 91, 0, 0, 207, 209, 5, 93, 0, 0, 208, 206, 1,
		0, 0, 0, 208, 209, 1, 0, 0, 0, 209, 54, 1, 0, 0, 0, 210, 214, 7, 5, 0,
		0, 211, 213, 7, 6, 0, 0, 212, 211, 1, 0, 0, 0, 213, 216, 1, 0, 0, 0, 214,
		212, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 56, 1, 0, 0, 0, 216, 214, 1,
		0, 0, 0, 217, 218, 7, 7, 0, 0, 218, 219, 5, 46, 0, 0, 219, 220, 7, 7, 0,
		0, 220, 58, 1, 0, 0, 0, 221, 222, 5, 47, 0, 0, 222, 223, 5, 47, 0, 0, 223,
		227, 1, 0, 0, 0, 224, 226, 8, 8, 0, 0, 225, 224, 1, 0, 0, 0, 226, 229,
		1, 0, 0, 0, 227, 225, 1, 0, 0, 0, 227, 228, 1, 0, 0, 0, 228, 60, 1, 0,
		0, 0, 229, 227, 1, 0, 0, 0, 230, 234, 5, 64, 0, 0, 231, 233, 8, 8, 0, 0,
		232, 231, 1, 0, 0, 0, 233, 236, 1, 0, 0, 0, 234, 232, 1, 0, 0, 0, 234,
		235, 1, 0, 0, 0, 235, 62, 1, 0, 0, 0, 236, 234, 1, 0, 0, 0, 237, 241, 5,
		35, 0, 0, 238, 240, 8, 8, 0, 0, 239, 238, 1, 0, 0, 0, 240, 243, 1, 0, 0,
		0, 241, 239, 1, 0, 0, 0, 241, 242, 1, 0, 0, 0, 242, 244, 1, 0, 0, 0, 243,
		241, 1, 0, 0, 0, 244, 245, 6, 31, 0, 0, 245, 64, 1, 0, 0, 0, 11, 0, 179,
		184, 189, 197, 203, 208, 214, 227, 234, 241, 1, 6, 0, 0,
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
	ObjectApiLexerT__19           = 20
	ObjectApiLexerT__20           = 21
	ObjectApiLexerT__21           = 22
	ObjectApiLexerT__22           = 23
	ObjectApiLexerWHITESPACE      = 24
	ObjectApiLexerINTEGER         = 25
	ObjectApiLexerHEX             = 26
	ObjectApiLexerTYPE_IDENTIFIER = 27
	ObjectApiLexerIDENTIFIER      = 28
	ObjectApiLexerVERSION         = 29
	ObjectApiLexerDOCLINE         = 30
	ObjectApiLexerTAGLINE         = 31
	ObjectApiLexerCOMMENT         = 32
)
