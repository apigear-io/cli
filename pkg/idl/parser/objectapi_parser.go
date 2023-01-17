// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // ObjectApi

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ObjectApiParser struct {
	*antlr.BaseParser
}

var objectapiParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func objectapiParserInit() {
	staticData := &objectapiParserStaticData
	staticData.literalNames = []string{
		"", "'module'", "'import'", "'interface'", "'{'", "'}'", "':'", "'('",
		"')'", "','", "'signal'", "'struct'", "'enum'", "'='", "'['", "']'",
		"'bool'", "'int'", "'float'", "'string'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "WHITESPACE", "INTEGER", "HEX", "TYPE_IDENTIFIER", "IDENTIFIER",
		"VERSION", "DOCLINE", "TAGLINE", "COMMENT",
	}
	staticData.ruleNames = []string{
		"documentRule", "headerRule", "moduleRule", "importRule", "declarationsRule",
		"interfaceRule", "interfaceMembersRule", "propertyRule", "operationRule",
		"operationReturnRule", "operationParamRule", "signalRule", "structRule",
		"structFieldRule", "enumRule", "enumMemberRule", "schemaRule", "arrayRule",
		"primitiveSchema", "symbolSchema", "metaRule",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 28, 232, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 1,
		0, 1, 0, 5, 0, 45, 8, 0, 10, 0, 12, 0, 48, 9, 0, 1, 1, 1, 1, 5, 1, 52,
		8, 1, 10, 1, 12, 1, 55, 9, 1, 1, 2, 5, 2, 58, 8, 2, 10, 2, 12, 2, 61, 9,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 3,
		4, 74, 8, 4, 1, 5, 5, 5, 77, 8, 5, 10, 5, 12, 5, 80, 9, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 5, 5, 86, 8, 5, 10, 5, 12, 5, 89, 9, 5, 1, 5, 1, 5, 1, 6, 1,
		6, 1, 6, 3, 6, 96, 8, 6, 1, 7, 5, 7, 99, 8, 7, 10, 7, 12, 7, 102, 9, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 5, 8, 109, 8, 8, 10, 8, 12, 8, 112, 9, 8,
		1, 8, 1, 8, 1, 8, 5, 8, 117, 8, 8, 10, 8, 12, 8, 120, 9, 8, 1, 8, 1, 8,
		3, 8, 124, 8, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 133,
		8, 10, 1, 11, 5, 11, 136, 8, 11, 10, 11, 12, 11, 139, 9, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 5, 11, 145, 8, 11, 10, 11, 12, 11, 148, 9, 11, 1, 11,
		1, 11, 1, 12, 5, 12, 153, 8, 12, 10, 12, 12, 12, 156, 9, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 5, 12, 162, 8, 12, 10, 12, 12, 12, 165, 9, 12, 1, 12,
		1, 12, 1, 13, 5, 13, 170, 8, 13, 10, 13, 12, 13, 173, 9, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 14, 5, 14, 180, 8, 14, 10, 14, 12, 14, 183, 9, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 189, 8, 14, 10, 14, 12, 14, 192, 9,
		14, 1, 14, 1, 14, 1, 15, 5, 15, 197, 8, 15, 10, 15, 12, 15, 200, 9, 15,
		1, 15, 1, 15, 1, 15, 3, 15, 205, 8, 15, 1, 15, 3, 15, 208, 8, 15, 1, 16,
		1, 16, 3, 16, 212, 8, 16, 1, 16, 3, 16, 215, 8, 16, 1, 17, 1, 17, 1, 17,
		1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 224, 8, 18, 1, 19, 1, 19, 1, 20, 1,
		20, 3, 20, 230, 8, 20, 1, 20, 0, 0, 21, 0, 2, 4, 6, 8, 10, 12, 14, 16,
		18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 0, 0, 240, 0, 42, 1, 0,
		0, 0, 2, 49, 1, 0, 0, 0, 4, 59, 1, 0, 0, 0, 6, 66, 1, 0, 0, 0, 8, 73, 1,
		0, 0, 0, 10, 78, 1, 0, 0, 0, 12, 95, 1, 0, 0, 0, 14, 100, 1, 0, 0, 0, 16,
		110, 1, 0, 0, 0, 18, 125, 1, 0, 0, 0, 20, 128, 1, 0, 0, 0, 22, 137, 1,
		0, 0, 0, 24, 154, 1, 0, 0, 0, 26, 171, 1, 0, 0, 0, 28, 181, 1, 0, 0, 0,
		30, 198, 1, 0, 0, 0, 32, 211, 1, 0, 0, 0, 34, 216, 1, 0, 0, 0, 36, 223,
		1, 0, 0, 0, 38, 225, 1, 0, 0, 0, 40, 229, 1, 0, 0, 0, 42, 46, 3, 2, 1,
		0, 43, 45, 3, 8, 4, 0, 44, 43, 1, 0, 0, 0, 45, 48, 1, 0, 0, 0, 46, 44,
		1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 1, 1, 0, 0, 0, 48, 46, 1, 0, 0, 0,
		49, 53, 3, 4, 2, 0, 50, 52, 3, 6, 3, 0, 51, 50, 1, 0, 0, 0, 52, 55, 1,
		0, 0, 0, 53, 51, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 3, 1, 0, 0, 0, 55,
		53, 1, 0, 0, 0, 56, 58, 3, 40, 20, 0, 57, 56, 1, 0, 0, 0, 58, 61, 1, 0,
		0, 0, 59, 57, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 62, 1, 0, 0, 0, 61, 59,
		1, 0, 0, 0, 62, 63, 5, 1, 0, 0, 63, 64, 5, 24, 0, 0, 64, 65, 5, 25, 0,
		0, 65, 5, 1, 0, 0, 0, 66, 67, 5, 2, 0, 0, 67, 68, 5, 24, 0, 0, 68, 69,
		5, 25, 0, 0, 69, 7, 1, 0, 0, 0, 70, 74, 3, 10, 5, 0, 71, 74, 3, 24, 12,
		0, 72, 74, 3, 28, 14, 0, 73, 70, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 73, 72,
		1, 0, 0, 0, 74, 9, 1, 0, 0, 0, 75, 77, 3, 40, 20, 0, 76, 75, 1, 0, 0, 0,
		77, 80, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 81, 1,
		0, 0, 0, 80, 78, 1, 0, 0, 0, 81, 82, 5, 3, 0, 0, 82, 83, 5, 24, 0, 0, 83,
		87, 5, 4, 0, 0, 84, 86, 3, 12, 6, 0, 85, 84, 1, 0, 0, 0, 86, 89, 1, 0,
		0, 0, 87, 85, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 90, 1, 0, 0, 0, 89, 87,
		1, 0, 0, 0, 90, 91, 5, 5, 0, 0, 91, 11, 1, 0, 0, 0, 92, 96, 3, 14, 7, 0,
		93, 96, 3, 16, 8, 0, 94, 96, 3, 22, 11, 0, 95, 92, 1, 0, 0, 0, 95, 93,
		1, 0, 0, 0, 95, 94, 1, 0, 0, 0, 96, 13, 1, 0, 0, 0, 97, 99, 3, 40, 20,
		0, 98, 97, 1, 0, 0, 0, 99, 102, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 100, 101,
		1, 0, 0, 0, 101, 103, 1, 0, 0, 0, 102, 100, 1, 0, 0, 0, 103, 104, 5, 24,
		0, 0, 104, 105, 5, 6, 0, 0, 105, 106, 3, 32, 16, 0, 106, 15, 1, 0, 0, 0,
		107, 109, 3, 40, 20, 0, 108, 107, 1, 0, 0, 0, 109, 112, 1, 0, 0, 0, 110,
		108, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 113, 1, 0, 0, 0, 112, 110,
		1, 0, 0, 0, 113, 114, 5, 24, 0, 0, 114, 118, 5, 7, 0, 0, 115, 117, 3, 20,
		10, 0, 116, 115, 1, 0, 0, 0, 117, 120, 1, 0, 0, 0, 118, 116, 1, 0, 0, 0,
		118, 119, 1, 0, 0, 0, 119, 121, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 121,
		123, 5, 8, 0, 0, 122, 124, 3, 18, 9, 0, 123, 122, 1, 0, 0, 0, 123, 124,
		1, 0, 0, 0, 124, 17, 1, 0, 0, 0, 125, 126, 5, 6, 0, 0, 126, 127, 3, 32,
		16, 0, 127, 19, 1, 0, 0, 0, 128, 129, 5, 24, 0, 0, 129, 130, 5, 6, 0, 0,
		130, 132, 3, 32, 16, 0, 131, 133, 5, 9, 0, 0, 132, 131, 1, 0, 0, 0, 132,
		133, 1, 0, 0, 0, 133, 21, 1, 0, 0, 0, 134, 136, 3, 40, 20, 0, 135, 134,
		1, 0, 0, 0, 136, 139, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 137, 138, 1, 0,
		0, 0, 138, 140, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 140, 141, 5, 10, 0, 0,
		141, 142, 5, 24, 0, 0, 142, 146, 5, 7, 0, 0, 143, 145, 3, 20, 10, 0, 144,
		143, 1, 0, 0, 0, 145, 148, 1, 0, 0, 0, 146, 144, 1, 0, 0, 0, 146, 147,
		1, 0, 0, 0, 147, 149, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 149, 150, 5, 8,
		0, 0, 150, 23, 1, 0, 0, 0, 151, 153, 3, 40, 20, 0, 152, 151, 1, 0, 0, 0,
		153, 156, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155,
		157, 1, 0, 0, 0, 156, 154, 1, 0, 0, 0, 157, 158, 5, 11, 0, 0, 158, 159,
		5, 24, 0, 0, 159, 163, 5, 4, 0, 0, 160, 162, 3, 26, 13, 0, 161, 160, 1,
		0, 0, 0, 162, 165, 1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 163, 164, 1, 0, 0,
		0, 164, 166, 1, 0, 0, 0, 165, 163, 1, 0, 0, 0, 166, 167, 5, 5, 0, 0, 167,
		25, 1, 0, 0, 0, 168, 170, 3, 40, 20, 0, 169, 168, 1, 0, 0, 0, 170, 173,
		1, 0, 0, 0, 171, 169, 1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172, 174, 1, 0,
		0, 0, 173, 171, 1, 0, 0, 0, 174, 175, 5, 24, 0, 0, 175, 176, 5, 6, 0, 0,
		176, 177, 3, 32, 16, 0, 177, 27, 1, 0, 0, 0, 178, 180, 3, 40, 20, 0, 179,
		178, 1, 0, 0, 0, 180, 183, 1, 0, 0, 0, 181, 179, 1, 0, 0, 0, 181, 182,
		1, 0, 0, 0, 182, 184, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 184, 185, 5, 12,
		0, 0, 185, 186, 5, 24, 0, 0, 186, 190, 5, 4, 0, 0, 187, 189, 3, 30, 15,
		0, 188, 187, 1, 0, 0, 0, 189, 192, 1, 0, 0, 0, 190, 188, 1, 0, 0, 0, 190,
		191, 1, 0, 0, 0, 191, 193, 1, 0, 0, 0, 192, 190, 1, 0, 0, 0, 193, 194,
		5, 5, 0, 0, 194, 29, 1, 0, 0, 0, 195, 197, 3, 40, 20, 0, 196, 195, 1, 0,
		0, 0, 197, 200, 1, 0, 0, 0, 198, 196, 1, 0, 0, 0, 198, 199, 1, 0, 0, 0,
		199, 201, 1, 0, 0, 0, 200, 198, 1, 0, 0, 0, 201, 204, 5, 24, 0, 0, 202,
		203, 5, 13, 0, 0, 203, 205, 5, 21, 0, 0, 204, 202, 1, 0, 0, 0, 204, 205,
		1, 0, 0, 0, 205, 207, 1, 0, 0, 0, 206, 208, 5, 9, 0, 0, 207, 206, 1, 0,
		0, 0, 207, 208, 1, 0, 0, 0, 208, 31, 1, 0, 0, 0, 209, 212, 3, 36, 18, 0,
		210, 212, 3, 38, 19, 0, 211, 209, 1, 0, 0, 0, 211, 210, 1, 0, 0, 0, 212,
		214, 1, 0, 0, 0, 213, 215, 3, 34, 17, 0, 214, 213, 1, 0, 0, 0, 214, 215,
		1, 0, 0, 0, 215, 33, 1, 0, 0, 0, 216, 217, 5, 14, 0, 0, 217, 218, 5, 15,
		0, 0, 218, 35, 1, 0, 0, 0, 219, 224, 5, 16, 0, 0, 220, 224, 5, 17, 0, 0,
		221, 224, 5, 18, 0, 0, 222, 224, 5, 19, 0, 0, 223, 219, 1, 0, 0, 0, 223,
		220, 1, 0, 0, 0, 223, 221, 1, 0, 0, 0, 223, 222, 1, 0, 0, 0, 224, 37, 1,
		0, 0, 0, 225, 226, 5, 24, 0, 0, 226, 39, 1, 0, 0, 0, 227, 230, 5, 27, 0,
		0, 228, 230, 5, 26, 0, 0, 229, 227, 1, 0, 0, 0, 229, 228, 1, 0, 0, 0, 230,
		41, 1, 0, 0, 0, 26, 46, 53, 59, 73, 78, 87, 95, 100, 110, 118, 123, 132,
		137, 146, 154, 163, 171, 181, 190, 198, 204, 207, 211, 214, 223, 229,
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

// ObjectApiParserInit initializes any static state used to implement ObjectApiParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewObjectApiParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ObjectApiParserInit() {
	staticData := &objectapiParserStaticData
	staticData.once.Do(objectapiParserInit)
}

// NewObjectApiParser produces a new parser instance for the optional input antlr.TokenStream.
func NewObjectApiParser(input antlr.TokenStream) *ObjectApiParser {
	ObjectApiParserInit()
	this := new(ObjectApiParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &objectapiParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// ObjectApiParser tokens.
const (
	ObjectApiParserEOF             = antlr.TokenEOF
	ObjectApiParserT__0            = 1
	ObjectApiParserT__1            = 2
	ObjectApiParserT__2            = 3
	ObjectApiParserT__3            = 4
	ObjectApiParserT__4            = 5
	ObjectApiParserT__5            = 6
	ObjectApiParserT__6            = 7
	ObjectApiParserT__7            = 8
	ObjectApiParserT__8            = 9
	ObjectApiParserT__9            = 10
	ObjectApiParserT__10           = 11
	ObjectApiParserT__11           = 12
	ObjectApiParserT__12           = 13
	ObjectApiParserT__13           = 14
	ObjectApiParserT__14           = 15
	ObjectApiParserT__15           = 16
	ObjectApiParserT__16           = 17
	ObjectApiParserT__17           = 18
	ObjectApiParserT__18           = 19
	ObjectApiParserWHITESPACE      = 20
	ObjectApiParserINTEGER         = 21
	ObjectApiParserHEX             = 22
	ObjectApiParserTYPE_IDENTIFIER = 23
	ObjectApiParserIDENTIFIER      = 24
	ObjectApiParserVERSION         = 25
	ObjectApiParserDOCLINE         = 26
	ObjectApiParserTAGLINE         = 27
	ObjectApiParserCOMMENT         = 28
)

// ObjectApiParser rules.
const (
	ObjectApiParserRULE_documentRule         = 0
	ObjectApiParserRULE_headerRule           = 1
	ObjectApiParserRULE_moduleRule           = 2
	ObjectApiParserRULE_importRule           = 3
	ObjectApiParserRULE_declarationsRule     = 4
	ObjectApiParserRULE_interfaceRule        = 5
	ObjectApiParserRULE_interfaceMembersRule = 6
	ObjectApiParserRULE_propertyRule         = 7
	ObjectApiParserRULE_operationRule        = 8
	ObjectApiParserRULE_operationReturnRule  = 9
	ObjectApiParserRULE_operationParamRule   = 10
	ObjectApiParserRULE_signalRule           = 11
	ObjectApiParserRULE_structRule           = 12
	ObjectApiParserRULE_structFieldRule      = 13
	ObjectApiParserRULE_enumRule             = 14
	ObjectApiParserRULE_enumMemberRule       = 15
	ObjectApiParserRULE_schemaRule           = 16
	ObjectApiParserRULE_arrayRule            = 17
	ObjectApiParserRULE_primitiveSchema      = 18
	ObjectApiParserRULE_symbolSchema         = 19
	ObjectApiParserRULE_metaRule             = 20
)

// IDocumentRuleContext is an interface to support dynamic dispatch.
type IDocumentRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDocumentRuleContext differentiates from other interfaces.
	IsDocumentRuleContext()
}

type DocumentRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDocumentRuleContext() *DocumentRuleContext {
	var p = new(DocumentRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ObjectApiParserRULE_documentRule
	return p
}

func (*DocumentRuleContext) IsDocumentRuleContext() {}

func NewDocumentRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentRuleContext {
	var p = new(DocumentRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectApiParserRULE_documentRule

	return p
}

func (s *DocumentRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *DocumentRuleContext) HeaderRule() IHeaderRuleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHeaderRuleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHeaderRuleContext)
}

func (s *DocumentRuleContext) AllDeclarationsRule() []IDeclarationsRuleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclarationsRuleContext); ok {
			len++
		}
	}

	tst := make([]IDeclarationsRuleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclarationsRuleContext); ok {
			tst[i] = t.(IDeclarationsRuleContext)
			i++
		}
	}

	return tst
}

func (s *DocumentRuleContext) DeclarationsRule(i int) IDeclarationsRuleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationsRuleContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationsRuleContext)
}

func (s *DocumentRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocumentRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocumentRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.EnterDocumentRule(s)
	}
}

func (s *DocumentRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.ExitDocumentRule(s)
	}
}

func (p *ObjectApiParser) DocumentRule() (localctx IDocumentRuleContext) {
	this := p
	_ = this

	localctx = NewDocumentRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ObjectApiParserRULE_documentRule)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(42)
		p.HeaderRule()
	}
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&201332744) != 0 {
		{
			p.SetState(43)
			p.DeclarationsRule()
		}

		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IHeaderRuleContext is an interface to support dynamic dispatch.
type IHeaderRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHeaderRuleContext differentiates from other interfaces.
	IsHeaderRuleContext()
}

type HeaderRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHeaderRuleContext() *HeaderRuleContext {
	var p = new(HeaderRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ObjectApiParserRULE_headerRule
	return p
}

func (*HeaderRuleContext) IsHeaderRuleContext() {}

func NewHeaderRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HeaderRuleContext {
	var p = new(HeaderRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectApiParserRULE_headerRule

	return p
}

func (s *HeaderRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *HeaderRuleContext) ModuleRule() IModuleRuleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModuleRuleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IModuleRuleContext)
}

func (s *HeaderRuleContext) AllImportRule() []IImportRuleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IImportRuleContext); ok {
			len++
		}
	}

	tst := make([]IImportRuleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IImportRuleContext); ok {
			tst[i] = t.(IImportRuleContext)
			i++
		}
	}

	return tst
}

func (s *HeaderRuleContext) ImportRule(i int) IImportRuleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportRuleContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImportRuleContext)
}

func (s *HeaderRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HeaderRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HeaderRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.EnterHeaderRule(s)
	}
}

func (s *HeaderRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.ExitHeaderRule(s)
	}
}

func (p *ObjectApiParser) HeaderRule() (localctx IHeaderRuleContext) {
	this := p
	_ = this

	localctx = NewHeaderRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ObjectApiParserRULE_headerRule)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(49)
		p.ModuleRule()
	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ObjectApiParserT__1 {
		{
			p.SetState(50)
			p.ImportRule()
		}

		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IModuleRuleContext is an interface to support dynamic dispatch.
type IModuleRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// GetVersion returns the version token.
	GetVersion() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// SetVersion sets the version token.
	SetVersion(antlr.Token)

	// IsModuleRuleContext differentiates from other interfaces.
	IsModuleRuleContext()
}

type ModuleRuleContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	name    antlr.Token
	version antlr.Token
}

func NewEmptyModuleRuleContext() *ModuleRuleContext {
	var p = new(ModuleRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ObjectApiParserRULE_moduleRule
	return p
}

func (*ModuleRuleContext) IsModuleRuleContext() {}

func NewModuleRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleRuleContext {
	var p = new(ModuleRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectApiParserRULE_moduleRule

	return p
}

func (s *ModuleRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *ModuleRuleContext) GetName() antlr.Token { return s.name }

func (s *ModuleRuleContext) GetVersion() antlr.Token { return s.version }

func (s *ModuleRuleContext) SetName(v antlr.Token) { s.name = v }

func (s *ModuleRuleContext) SetVersion(v antlr.Token) { s.version = v }

func (s *ModuleRuleContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ObjectApiParserIDENTIFIER, 0)
}

func (s *ModuleRuleContext) VERSION() antlr.TerminalNode {
	return s.GetToken(ObjectApiParserVERSION, 0)
}

func (s *ModuleRuleContext) AllMetaRule() []IMetaRuleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMetaRuleContext); ok {
			len++
		}
	}

	tst := make([]IMetaRuleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMetaRuleContext); ok {
			tst[i] = t.(IMetaRuleContext)
			i++
		}
	}

	return tst
}

func (s *ModuleRuleContext) MetaRule(i int) IMetaRuleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMetaRuleContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMetaRuleContext)
}

func (s *ModuleRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModuleRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.EnterModuleRule(s)
	}
}

func (s *ModuleRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.ExitModuleRule(s)
	}
}

func (p *ObjectApiParser) ModuleRule() (localctx IModuleRuleContext) {
	this := p
	_ = this

	localctx = NewModuleRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ObjectApiParserRULE_moduleRule)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ObjectApiParserDOCLINE || _la == ObjectApiParserTAGLINE {
		{
			p.SetState(56)
			p.MetaRule()
		}

		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(62)
		p.Match(ObjectApiParserT__0)
	}
	{
		p.SetState(63)

		var _m = p.Match(ObjectApiParserIDENTIFIER)

		localctx.(*ModuleRuleContext).name = _m
	}
	{
		p.SetState(64)

		var _m = p.Match(ObjectApiParserVERSION)

		localctx.(*ModuleRuleContext).version = _m
	}

	return localctx
}

// IImportRuleContext is an interface to support dynamic dispatch.
type IImportRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// GetVersion returns the version token.
	GetVersion() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// SetVersion sets the version token.
	SetVersion(antlr.Token)

	// IsImportRuleContext differentiates from other interfaces.
	IsImportRuleContext()
}

type ImportRuleContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	name    antlr.Token
	version antlr.Token
}

func NewEmptyImportRuleContext() *ImportRuleContext {
	var p = new(ImportRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ObjectApiParserRULE_importRule
	return p
}

func (*ImportRuleContext) IsImportRuleContext() {}

func NewImportRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportRuleContext {
	var p = new(ImportRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectApiParserRULE_importRule

	return p
}

func (s *ImportRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportRuleContext) GetName() antlr.Token { return s.name }

func (s *ImportRuleContext) GetVersion() antlr.Token { return s.version }

func (s *ImportRuleContext) SetName(v antlr.Token) { s.name = v }

func (s *ImportRuleContext) SetVersion(v antlr.Token) { s.version = v }

func (s *ImportRuleContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ObjectApiParserIDENTIFIER, 0)
}

func (s *ImportRuleContext) VERSION() antlr.TerminalNode {
	return s.GetToken(ObjectApiParserVERSION, 0)
}

func (s *ImportRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.EnterImportRule(s)
	}
}

func (s *ImportRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.ExitImportRule(s)
	}
}

func (p *ObjectApiParser) ImportRule() (localctx IImportRuleContext) {
	this := p
	_ = this

	localctx = NewImportRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ObjectApiParserRULE_importRule)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		p.Match(ObjectApiParserT__1)
	}
	{
		p.SetState(67)

		var _m = p.Match(ObjectApiParserIDENTIFIER)

		localctx.(*ImportRuleContext).name = _m
	}
	{
		p.SetState(68)

		var _m = p.Match(ObjectApiParserVERSION)

		localctx.(*ImportRuleContext).version = _m
	}

	return localctx
}

// IDeclarationsRuleContext is an interface to support dynamic dispatch.
type IDeclarationsRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationsRuleContext differentiates from other interfaces.
	IsDeclarationsRuleContext()
}

type DeclarationsRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationsRuleContext() *DeclarationsRuleContext {
	var p = new(DeclarationsRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ObjectApiParserRULE_declarationsRule
	return p
}

func (*DeclarationsRuleContext) IsDeclarationsRuleContext() {}

func NewDeclarationsRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationsRuleContext {
	var p = new(DeclarationsRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectApiParserRULE_declarationsRule

	return p
}

func (s *DeclarationsRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationsRuleContext) InterfaceRule() IInterfaceRuleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInterfaceRuleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInterfaceRuleContext)
}

func (s *DeclarationsRuleContext) StructRule() IStructRuleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructRuleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructRuleContext)
}

func (s *DeclarationsRuleContext) EnumRule() IEnumRuleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumRuleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumRuleContext)
}

func (s *DeclarationsRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationsRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationsRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.EnterDeclarationsRule(s)
	}
}

func (s *DeclarationsRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.ExitDeclarationsRule(s)
	}
}

func (p *ObjectApiParser) DeclarationsRule() (localctx IDeclarationsRuleContext) {
	this := p
	_ = this

	localctx = NewDeclarationsRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ObjectApiParserRULE_declarationsRule)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(70)
			p.InterfaceRule()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(71)
			p.StructRule()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(72)
			p.EnumRule()
		}

	}

	return localctx
}

// IInterfaceRuleContext is an interface to support dynamic dispatch.
type IInterfaceRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsInterfaceRuleContext differentiates from other interfaces.
	IsInterfaceRuleContext()
}

type InterfaceRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyInterfaceRuleContext() *InterfaceRuleContext {
	var p = new(InterfaceRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ObjectApiParserRULE_interfaceRule
	return p
}

func (*InterfaceRuleContext) IsInterfaceRuleContext() {}

func NewInterfaceRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InterfaceRuleContext {
	var p = new(InterfaceRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectApiParserRULE_interfaceRule

	return p
}

func (s *InterfaceRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *InterfaceRuleContext) GetName() antlr.Token { return s.name }

func (s *InterfaceRuleContext) SetName(v antlr.Token) { s.name = v }

func (s *InterfaceRuleContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ObjectApiParserIDENTIFIER, 0)
}

func (s *InterfaceRuleContext) AllMetaRule() []IMetaRuleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMetaRuleContext); ok {
			len++
		}
	}

	tst := make([]IMetaRuleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMetaRuleContext); ok {
			tst[i] = t.(IMetaRuleContext)
			i++
		}
	}

	return tst
}

func (s *InterfaceRuleContext) MetaRule(i int) IMetaRuleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMetaRuleContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMetaRuleContext)
}

func (s *InterfaceRuleContext) AllInterfaceMembersRule() []IInterfaceMembersRuleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInterfaceMembersRuleContext); ok {
			len++
		}
	}

	tst := make([]IInterfaceMembersRuleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInterfaceMembersRuleContext); ok {
			tst[i] = t.(IInterfaceMembersRuleContext)
			i++
		}
	}

	return tst
}

func (s *InterfaceRuleContext) InterfaceMembersRule(i int) IInterfaceMembersRuleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInterfaceMembersRuleContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInterfaceMembersRuleContext)
}

func (s *InterfaceRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InterfaceRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InterfaceRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.EnterInterfaceRule(s)
	}
}

func (s *InterfaceRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.ExitInterfaceRule(s)
	}
}

func (p *ObjectApiParser) InterfaceRule() (localctx IInterfaceRuleContext) {
	this := p
	_ = this

	localctx = NewInterfaceRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ObjectApiParserRULE_interfaceRule)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ObjectApiParserDOCLINE || _la == ObjectApiParserTAGLINE {
		{
			p.SetState(75)
			p.MetaRule()
		}

		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(81)
		p.Match(ObjectApiParserT__2)
	}
	{
		p.SetState(82)

		var _m = p.Match(ObjectApiParserIDENTIFIER)

		localctx.(*InterfaceRuleContext).name = _m
	}
	{
		p.SetState(83)
		p.Match(ObjectApiParserT__3)
	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&218104832) != 0 {
		{
			p.SetState(84)
			p.InterfaceMembersRule()
		}

		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(90)
		p.Match(ObjectApiParserT__4)
	}

	return localctx
}

// IInterfaceMembersRuleContext is an interface to support dynamic dispatch.
type IInterfaceMembersRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInterfaceMembersRuleContext differentiates from other interfaces.
	IsInterfaceMembersRuleContext()
}

type InterfaceMembersRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInterfaceMembersRuleContext() *InterfaceMembersRuleContext {
	var p = new(InterfaceMembersRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ObjectApiParserRULE_interfaceMembersRule
	return p
}

func (*InterfaceMembersRuleContext) IsInterfaceMembersRuleContext() {}

func NewInterfaceMembersRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InterfaceMembersRuleContext {
	var p = new(InterfaceMembersRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectApiParserRULE_interfaceMembersRule

	return p
}

func (s *InterfaceMembersRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *InterfaceMembersRuleContext) PropertyRule() IPropertyRuleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropertyRuleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropertyRuleContext)
}

func (s *InterfaceMembersRuleContext) OperationRule() IOperationRuleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperationRuleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperationRuleContext)
}

func (s *InterfaceMembersRuleContext) SignalRule() ISignalRuleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISignalRuleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISignalRuleContext)
}

func (s *InterfaceMembersRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InterfaceMembersRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InterfaceMembersRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.EnterInterfaceMembersRule(s)
	}
}

func (s *InterfaceMembersRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.ExitInterfaceMembersRule(s)
	}
}

func (p *ObjectApiParser) InterfaceMembersRule() (localctx IInterfaceMembersRuleContext) {
	this := p
	_ = this

	localctx = NewInterfaceMembersRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ObjectApiParserRULE_interfaceMembersRule)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(92)
			p.PropertyRule()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(93)
			p.OperationRule()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(94)
			p.SignalRule()
		}

	}

	return localctx
}

// IPropertyRuleContext is an interface to support dynamic dispatch.
type IPropertyRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// GetSchema returns the schema rule contexts.
	GetSchema() ISchemaRuleContext

	// SetSchema sets the schema rule contexts.
	SetSchema(ISchemaRuleContext)

	// IsPropertyRuleContext differentiates from other interfaces.
	IsPropertyRuleContext()
}

type PropertyRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	schema ISchemaRuleContext
}

func NewEmptyPropertyRuleContext() *PropertyRuleContext {
	var p = new(PropertyRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ObjectApiParserRULE_propertyRule
	return p
}

func (*PropertyRuleContext) IsPropertyRuleContext() {}

func NewPropertyRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyRuleContext {
	var p = new(PropertyRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectApiParserRULE_propertyRule

	return p
}

func (s *PropertyRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyRuleContext) GetName() antlr.Token { return s.name }

func (s *PropertyRuleContext) SetName(v antlr.Token) { s.name = v }

func (s *PropertyRuleContext) GetSchema() ISchemaRuleContext { return s.schema }

func (s *PropertyRuleContext) SetSchema(v ISchemaRuleContext) { s.schema = v }

func (s *PropertyRuleContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ObjectApiParserIDENTIFIER, 0)
}

func (s *PropertyRuleContext) SchemaRule() ISchemaRuleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISchemaRuleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISchemaRuleContext)
}

func (s *PropertyRuleContext) AllMetaRule() []IMetaRuleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMetaRuleContext); ok {
			len++
		}
	}

	tst := make([]IMetaRuleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMetaRuleContext); ok {
			tst[i] = t.(IMetaRuleContext)
			i++
		}
	}

	return tst
}

func (s *PropertyRuleContext) MetaRule(i int) IMetaRuleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMetaRuleContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMetaRuleContext)
}

func (s *PropertyRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.EnterPropertyRule(s)
	}
}

func (s *PropertyRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.ExitPropertyRule(s)
	}
}

func (p *ObjectApiParser) PropertyRule() (localctx IPropertyRuleContext) {
	this := p
	_ = this

	localctx = NewPropertyRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ObjectApiParserRULE_propertyRule)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ObjectApiParserDOCLINE || _la == ObjectApiParserTAGLINE {
		{
			p.SetState(97)
			p.MetaRule()
		}

		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(103)

		var _m = p.Match(ObjectApiParserIDENTIFIER)

		localctx.(*PropertyRuleContext).name = _m
	}
	{
		p.SetState(104)
		p.Match(ObjectApiParserT__5)
	}
	{
		p.SetState(105)

		var _x = p.SchemaRule()

		localctx.(*PropertyRuleContext).schema = _x
	}

	return localctx
}

// IOperationRuleContext is an interface to support dynamic dispatch.
type IOperationRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// GetParams returns the params rule contexts.
	GetParams() IOperationParamRuleContext

	// SetParams sets the params rule contexts.
	SetParams(IOperationParamRuleContext)

	// IsOperationRuleContext differentiates from other interfaces.
	IsOperationRuleContext()
}

type OperationRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	params IOperationParamRuleContext
}

func NewEmptyOperationRuleContext() *OperationRuleContext {
	var p = new(OperationRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ObjectApiParserRULE_operationRule
	return p
}

func (*OperationRuleContext) IsOperationRuleContext() {}

func NewOperationRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationRuleContext {
	var p = new(OperationRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectApiParserRULE_operationRule

	return p
}

func (s *OperationRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *OperationRuleContext) GetName() antlr.Token { return s.name }

func (s *OperationRuleContext) SetName(v antlr.Token) { s.name = v }

func (s *OperationRuleContext) GetParams() IOperationParamRuleContext { return s.params }

func (s *OperationRuleContext) SetParams(v IOperationParamRuleContext) { s.params = v }

func (s *OperationRuleContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ObjectApiParserIDENTIFIER, 0)
}

func (s *OperationRuleContext) AllMetaRule() []IMetaRuleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMetaRuleContext); ok {
			len++
		}
	}

	tst := make([]IMetaRuleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMetaRuleContext); ok {
			tst[i] = t.(IMetaRuleContext)
			i++
		}
	}

	return tst
}

func (s *OperationRuleContext) MetaRule(i int) IMetaRuleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMetaRuleContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMetaRuleContext)
}

func (s *OperationRuleContext) OperationReturnRule() IOperationReturnRuleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperationReturnRuleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperationReturnRuleContext)
}

func (s *OperationRuleContext) AllOperationParamRule() []IOperationParamRuleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOperationParamRuleContext); ok {
			len++
		}
	}

	tst := make([]IOperationParamRuleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOperationParamRuleContext); ok {
			tst[i] = t.(IOperationParamRuleContext)
			i++
		}
	}

	return tst
}

func (s *OperationRuleContext) OperationParamRule(i int) IOperationParamRuleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperationParamRuleContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperationParamRuleContext)
}

func (s *OperationRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperationRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.EnterOperationRule(s)
	}
}

func (s *OperationRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.ExitOperationRule(s)
	}
}

func (p *ObjectApiParser) OperationRule() (localctx IOperationRuleContext) {
	this := p
	_ = this

	localctx = NewOperationRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ObjectApiParserRULE_operationRule)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ObjectApiParserDOCLINE || _la == ObjectApiParserTAGLINE {
		{
			p.SetState(107)
			p.MetaRule()
		}

		p.SetState(112)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(113)

		var _m = p.Match(ObjectApiParserIDENTIFIER)

		localctx.(*OperationRuleContext).name = _m
	}
	{
		p.SetState(114)
		p.Match(ObjectApiParserT__6)
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ObjectApiParserIDENTIFIER {
		{
			p.SetState(115)

			var _x = p.OperationParamRule()

			localctx.(*OperationRuleContext).params = _x
		}

		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(121)
		p.Match(ObjectApiParserT__7)
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ObjectApiParserT__5 {
		{
			p.SetState(122)
			p.OperationReturnRule()
		}

	}

	return localctx
}

// IOperationReturnRuleContext is an interface to support dynamic dispatch.
type IOperationReturnRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSchema returns the schema rule contexts.
	GetSchema() ISchemaRuleContext

	// SetSchema sets the schema rule contexts.
	SetSchema(ISchemaRuleContext)

	// IsOperationReturnRuleContext differentiates from other interfaces.
	IsOperationReturnRuleContext()
}

type OperationReturnRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	schema ISchemaRuleContext
}

func NewEmptyOperationReturnRuleContext() *OperationReturnRuleContext {
	var p = new(OperationReturnRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ObjectApiParserRULE_operationReturnRule
	return p
}

func (*OperationReturnRuleContext) IsOperationReturnRuleContext() {}

func NewOperationReturnRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationReturnRuleContext {
	var p = new(OperationReturnRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectApiParserRULE_operationReturnRule

	return p
}

func (s *OperationReturnRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *OperationReturnRuleContext) GetSchema() ISchemaRuleContext { return s.schema }

func (s *OperationReturnRuleContext) SetSchema(v ISchemaRuleContext) { s.schema = v }

func (s *OperationReturnRuleContext) SchemaRule() ISchemaRuleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISchemaRuleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISchemaRuleContext)
}

func (s *OperationReturnRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationReturnRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperationReturnRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.EnterOperationReturnRule(s)
	}
}

func (s *OperationReturnRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.ExitOperationReturnRule(s)
	}
}

func (p *ObjectApiParser) OperationReturnRule() (localctx IOperationReturnRuleContext) {
	this := p
	_ = this

	localctx = NewOperationReturnRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ObjectApiParserRULE_operationReturnRule)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(125)
		p.Match(ObjectApiParserT__5)
	}
	{
		p.SetState(126)

		var _x = p.SchemaRule()

		localctx.(*OperationReturnRuleContext).schema = _x
	}

	return localctx
}

// IOperationParamRuleContext is an interface to support dynamic dispatch.
type IOperationParamRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// GetSchema returns the schema rule contexts.
	GetSchema() ISchemaRuleContext

	// SetSchema sets the schema rule contexts.
	SetSchema(ISchemaRuleContext)

	// IsOperationParamRuleContext differentiates from other interfaces.
	IsOperationParamRuleContext()
}

type OperationParamRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	schema ISchemaRuleContext
}

func NewEmptyOperationParamRuleContext() *OperationParamRuleContext {
	var p = new(OperationParamRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ObjectApiParserRULE_operationParamRule
	return p
}

func (*OperationParamRuleContext) IsOperationParamRuleContext() {}

func NewOperationParamRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationParamRuleContext {
	var p = new(OperationParamRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectApiParserRULE_operationParamRule

	return p
}

func (s *OperationParamRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *OperationParamRuleContext) GetName() antlr.Token { return s.name }

func (s *OperationParamRuleContext) SetName(v antlr.Token) { s.name = v }

func (s *OperationParamRuleContext) GetSchema() ISchemaRuleContext { return s.schema }

func (s *OperationParamRuleContext) SetSchema(v ISchemaRuleContext) { s.schema = v }

func (s *OperationParamRuleContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ObjectApiParserIDENTIFIER, 0)
}

func (s *OperationParamRuleContext) SchemaRule() ISchemaRuleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISchemaRuleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISchemaRuleContext)
}

func (s *OperationParamRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationParamRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperationParamRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.EnterOperationParamRule(s)
	}
}

func (s *OperationParamRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.ExitOperationParamRule(s)
	}
}

func (p *ObjectApiParser) OperationParamRule() (localctx IOperationParamRuleContext) {
	this := p
	_ = this

	localctx = NewOperationParamRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ObjectApiParserRULE_operationParamRule)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)

		var _m = p.Match(ObjectApiParserIDENTIFIER)

		localctx.(*OperationParamRuleContext).name = _m
	}
	{
		p.SetState(129)
		p.Match(ObjectApiParserT__5)
	}
	{
		p.SetState(130)

		var _x = p.SchemaRule()

		localctx.(*OperationParamRuleContext).schema = _x
	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ObjectApiParserT__8 {
		{
			p.SetState(131)
			p.Match(ObjectApiParserT__8)
		}

	}

	return localctx
}

// ISignalRuleContext is an interface to support dynamic dispatch.
type ISignalRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// GetParams returns the params rule contexts.
	GetParams() IOperationParamRuleContext

	// SetParams sets the params rule contexts.
	SetParams(IOperationParamRuleContext)

	// IsSignalRuleContext differentiates from other interfaces.
	IsSignalRuleContext()
}

type SignalRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	params IOperationParamRuleContext
}

func NewEmptySignalRuleContext() *SignalRuleContext {
	var p = new(SignalRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ObjectApiParserRULE_signalRule
	return p
}

func (*SignalRuleContext) IsSignalRuleContext() {}

func NewSignalRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SignalRuleContext {
	var p = new(SignalRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectApiParserRULE_signalRule

	return p
}

func (s *SignalRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *SignalRuleContext) GetName() antlr.Token { return s.name }

func (s *SignalRuleContext) SetName(v antlr.Token) { s.name = v }

func (s *SignalRuleContext) GetParams() IOperationParamRuleContext { return s.params }

func (s *SignalRuleContext) SetParams(v IOperationParamRuleContext) { s.params = v }

func (s *SignalRuleContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ObjectApiParserIDENTIFIER, 0)
}

func (s *SignalRuleContext) AllMetaRule() []IMetaRuleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMetaRuleContext); ok {
			len++
		}
	}

	tst := make([]IMetaRuleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMetaRuleContext); ok {
			tst[i] = t.(IMetaRuleContext)
			i++
		}
	}

	return tst
}

func (s *SignalRuleContext) MetaRule(i int) IMetaRuleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMetaRuleContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMetaRuleContext)
}

func (s *SignalRuleContext) AllOperationParamRule() []IOperationParamRuleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOperationParamRuleContext); ok {
			len++
		}
	}

	tst := make([]IOperationParamRuleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOperationParamRuleContext); ok {
			tst[i] = t.(IOperationParamRuleContext)
			i++
		}
	}

	return tst
}

func (s *SignalRuleContext) OperationParamRule(i int) IOperationParamRuleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperationParamRuleContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperationParamRuleContext)
}

func (s *SignalRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SignalRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SignalRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.EnterSignalRule(s)
	}
}

func (s *SignalRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.ExitSignalRule(s)
	}
}

func (p *ObjectApiParser) SignalRule() (localctx ISignalRuleContext) {
	this := p
	_ = this

	localctx = NewSignalRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ObjectApiParserRULE_signalRule)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ObjectApiParserDOCLINE || _la == ObjectApiParserTAGLINE {
		{
			p.SetState(134)
			p.MetaRule()
		}

		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(140)
		p.Match(ObjectApiParserT__9)
	}
	{
		p.SetState(141)

		var _m = p.Match(ObjectApiParserIDENTIFIER)

		localctx.(*SignalRuleContext).name = _m
	}
	{
		p.SetState(142)
		p.Match(ObjectApiParserT__6)
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ObjectApiParserIDENTIFIER {
		{
			p.SetState(143)

			var _x = p.OperationParamRule()

			localctx.(*SignalRuleContext).params = _x
		}

		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(149)
		p.Match(ObjectApiParserT__7)
	}

	return localctx
}

// IStructRuleContext is an interface to support dynamic dispatch.
type IStructRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsStructRuleContext differentiates from other interfaces.
	IsStructRuleContext()
}

type StructRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyStructRuleContext() *StructRuleContext {
	var p = new(StructRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ObjectApiParserRULE_structRule
	return p
}

func (*StructRuleContext) IsStructRuleContext() {}

func NewStructRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructRuleContext {
	var p = new(StructRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectApiParserRULE_structRule

	return p
}

func (s *StructRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *StructRuleContext) GetName() antlr.Token { return s.name }

func (s *StructRuleContext) SetName(v antlr.Token) { s.name = v }

func (s *StructRuleContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ObjectApiParserIDENTIFIER, 0)
}

func (s *StructRuleContext) AllMetaRule() []IMetaRuleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMetaRuleContext); ok {
			len++
		}
	}

	tst := make([]IMetaRuleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMetaRuleContext); ok {
			tst[i] = t.(IMetaRuleContext)
			i++
		}
	}

	return tst
}

func (s *StructRuleContext) MetaRule(i int) IMetaRuleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMetaRuleContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMetaRuleContext)
}

func (s *StructRuleContext) AllStructFieldRule() []IStructFieldRuleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStructFieldRuleContext); ok {
			len++
		}
	}

	tst := make([]IStructFieldRuleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStructFieldRuleContext); ok {
			tst[i] = t.(IStructFieldRuleContext)
			i++
		}
	}

	return tst
}

func (s *StructRuleContext) StructFieldRule(i int) IStructFieldRuleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructFieldRuleContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructFieldRuleContext)
}

func (s *StructRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.EnterStructRule(s)
	}
}

func (s *StructRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.ExitStructRule(s)
	}
}

func (p *ObjectApiParser) StructRule() (localctx IStructRuleContext) {
	this := p
	_ = this

	localctx = NewStructRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ObjectApiParserRULE_structRule)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ObjectApiParserDOCLINE || _la == ObjectApiParserTAGLINE {
		{
			p.SetState(151)
			p.MetaRule()
		}

		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(157)
		p.Match(ObjectApiParserT__10)
	}
	{
		p.SetState(158)

		var _m = p.Match(ObjectApiParserIDENTIFIER)

		localctx.(*StructRuleContext).name = _m
	}
	{
		p.SetState(159)
		p.Match(ObjectApiParserT__3)
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&218103808) != 0 {
		{
			p.SetState(160)
			p.StructFieldRule()
		}

		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(166)
		p.Match(ObjectApiParserT__4)
	}

	return localctx
}

// IStructFieldRuleContext is an interface to support dynamic dispatch.
type IStructFieldRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// GetSchema returns the schema rule contexts.
	GetSchema() ISchemaRuleContext

	// SetSchema sets the schema rule contexts.
	SetSchema(ISchemaRuleContext)

	// IsStructFieldRuleContext differentiates from other interfaces.
	IsStructFieldRuleContext()
}

type StructFieldRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	schema ISchemaRuleContext
}

func NewEmptyStructFieldRuleContext() *StructFieldRuleContext {
	var p = new(StructFieldRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ObjectApiParserRULE_structFieldRule
	return p
}

func (*StructFieldRuleContext) IsStructFieldRuleContext() {}

func NewStructFieldRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructFieldRuleContext {
	var p = new(StructFieldRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectApiParserRULE_structFieldRule

	return p
}

func (s *StructFieldRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *StructFieldRuleContext) GetName() antlr.Token { return s.name }

func (s *StructFieldRuleContext) SetName(v antlr.Token) { s.name = v }

func (s *StructFieldRuleContext) GetSchema() ISchemaRuleContext { return s.schema }

func (s *StructFieldRuleContext) SetSchema(v ISchemaRuleContext) { s.schema = v }

func (s *StructFieldRuleContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ObjectApiParserIDENTIFIER, 0)
}

func (s *StructFieldRuleContext) SchemaRule() ISchemaRuleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISchemaRuleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISchemaRuleContext)
}

func (s *StructFieldRuleContext) AllMetaRule() []IMetaRuleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMetaRuleContext); ok {
			len++
		}
	}

	tst := make([]IMetaRuleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMetaRuleContext); ok {
			tst[i] = t.(IMetaRuleContext)
			i++
		}
	}

	return tst
}

func (s *StructFieldRuleContext) MetaRule(i int) IMetaRuleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMetaRuleContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMetaRuleContext)
}

func (s *StructFieldRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructFieldRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructFieldRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.EnterStructFieldRule(s)
	}
}

func (s *StructFieldRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.ExitStructFieldRule(s)
	}
}

func (p *ObjectApiParser) StructFieldRule() (localctx IStructFieldRuleContext) {
	this := p
	_ = this

	localctx = NewStructFieldRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ObjectApiParserRULE_structFieldRule)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ObjectApiParserDOCLINE || _la == ObjectApiParserTAGLINE {
		{
			p.SetState(168)
			p.MetaRule()
		}

		p.SetState(173)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(174)

		var _m = p.Match(ObjectApiParserIDENTIFIER)

		localctx.(*StructFieldRuleContext).name = _m
	}
	{
		p.SetState(175)
		p.Match(ObjectApiParserT__5)
	}
	{
		p.SetState(176)

		var _x = p.SchemaRule()

		localctx.(*StructFieldRuleContext).schema = _x
	}

	return localctx
}

// IEnumRuleContext is an interface to support dynamic dispatch.
type IEnumRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsEnumRuleContext differentiates from other interfaces.
	IsEnumRuleContext()
}

type EnumRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyEnumRuleContext() *EnumRuleContext {
	var p = new(EnumRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ObjectApiParserRULE_enumRule
	return p
}

func (*EnumRuleContext) IsEnumRuleContext() {}

func NewEnumRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumRuleContext {
	var p = new(EnumRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectApiParserRULE_enumRule

	return p
}

func (s *EnumRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumRuleContext) GetName() antlr.Token { return s.name }

func (s *EnumRuleContext) SetName(v antlr.Token) { s.name = v }

func (s *EnumRuleContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ObjectApiParserIDENTIFIER, 0)
}

func (s *EnumRuleContext) AllMetaRule() []IMetaRuleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMetaRuleContext); ok {
			len++
		}
	}

	tst := make([]IMetaRuleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMetaRuleContext); ok {
			tst[i] = t.(IMetaRuleContext)
			i++
		}
	}

	return tst
}

func (s *EnumRuleContext) MetaRule(i int) IMetaRuleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMetaRuleContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMetaRuleContext)
}

func (s *EnumRuleContext) AllEnumMemberRule() []IEnumMemberRuleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEnumMemberRuleContext); ok {
			len++
		}
	}

	tst := make([]IEnumMemberRuleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEnumMemberRuleContext); ok {
			tst[i] = t.(IEnumMemberRuleContext)
			i++
		}
	}

	return tst
}

func (s *EnumRuleContext) EnumMemberRule(i int) IEnumMemberRuleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumMemberRuleContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumMemberRuleContext)
}

func (s *EnumRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.EnterEnumRule(s)
	}
}

func (s *EnumRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.ExitEnumRule(s)
	}
}

func (p *ObjectApiParser) EnumRule() (localctx IEnumRuleContext) {
	this := p
	_ = this

	localctx = NewEnumRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ObjectApiParserRULE_enumRule)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ObjectApiParserDOCLINE || _la == ObjectApiParserTAGLINE {
		{
			p.SetState(178)
			p.MetaRule()
		}

		p.SetState(183)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(184)
		p.Match(ObjectApiParserT__11)
	}
	{
		p.SetState(185)

		var _m = p.Match(ObjectApiParserIDENTIFIER)

		localctx.(*EnumRuleContext).name = _m
	}
	{
		p.SetState(186)
		p.Match(ObjectApiParserT__3)
	}
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&218103808) != 0 {
		{
			p.SetState(187)
			p.EnumMemberRule()
		}

		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(193)
		p.Match(ObjectApiParserT__4)
	}

	return localctx
}

// IEnumMemberRuleContext is an interface to support dynamic dispatch.
type IEnumMemberRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsEnumMemberRuleContext differentiates from other interfaces.
	IsEnumMemberRuleContext()
}

type EnumMemberRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	value  antlr.Token
}

func NewEmptyEnumMemberRuleContext() *EnumMemberRuleContext {
	var p = new(EnumMemberRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ObjectApiParserRULE_enumMemberRule
	return p
}

func (*EnumMemberRuleContext) IsEnumMemberRuleContext() {}

func NewEnumMemberRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumMemberRuleContext {
	var p = new(EnumMemberRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectApiParserRULE_enumMemberRule

	return p
}

func (s *EnumMemberRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumMemberRuleContext) GetName() antlr.Token { return s.name }

func (s *EnumMemberRuleContext) GetValue() antlr.Token { return s.value }

func (s *EnumMemberRuleContext) SetName(v antlr.Token) { s.name = v }

func (s *EnumMemberRuleContext) SetValue(v antlr.Token) { s.value = v }

func (s *EnumMemberRuleContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ObjectApiParserIDENTIFIER, 0)
}

func (s *EnumMemberRuleContext) AllMetaRule() []IMetaRuleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMetaRuleContext); ok {
			len++
		}
	}

	tst := make([]IMetaRuleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMetaRuleContext); ok {
			tst[i] = t.(IMetaRuleContext)
			i++
		}
	}

	return tst
}

func (s *EnumMemberRuleContext) MetaRule(i int) IMetaRuleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMetaRuleContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMetaRuleContext)
}

func (s *EnumMemberRuleContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(ObjectApiParserINTEGER, 0)
}

func (s *EnumMemberRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumMemberRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumMemberRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.EnterEnumMemberRule(s)
	}
}

func (s *EnumMemberRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.ExitEnumMemberRule(s)
	}
}

func (p *ObjectApiParser) EnumMemberRule() (localctx IEnumMemberRuleContext) {
	this := p
	_ = this

	localctx = NewEnumMemberRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ObjectApiParserRULE_enumMemberRule)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ObjectApiParserDOCLINE || _la == ObjectApiParserTAGLINE {
		{
			p.SetState(195)
			p.MetaRule()
		}

		p.SetState(200)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(201)

		var _m = p.Match(ObjectApiParserIDENTIFIER)

		localctx.(*EnumMemberRuleContext).name = _m
	}
	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ObjectApiParserT__12 {
		{
			p.SetState(202)
			p.Match(ObjectApiParserT__12)
		}
		{
			p.SetState(203)

			var _m = p.Match(ObjectApiParserINTEGER)

			localctx.(*EnumMemberRuleContext).value = _m
		}

	}
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ObjectApiParserT__8 {
		{
			p.SetState(206)
			p.Match(ObjectApiParserT__8)
		}

	}

	return localctx
}

// ISchemaRuleContext is an interface to support dynamic dispatch.
type ISchemaRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSchemaRuleContext differentiates from other interfaces.
	IsSchemaRuleContext()
}

type SchemaRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySchemaRuleContext() *SchemaRuleContext {
	var p = new(SchemaRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ObjectApiParserRULE_schemaRule
	return p
}

func (*SchemaRuleContext) IsSchemaRuleContext() {}

func NewSchemaRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SchemaRuleContext {
	var p = new(SchemaRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectApiParserRULE_schemaRule

	return p
}

func (s *SchemaRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *SchemaRuleContext) PrimitiveSchema() IPrimitiveSchemaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitiveSchemaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitiveSchemaContext)
}

func (s *SchemaRuleContext) SymbolSchema() ISymbolSchemaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISymbolSchemaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISymbolSchemaContext)
}

func (s *SchemaRuleContext) ArrayRule() IArrayRuleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayRuleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayRuleContext)
}

func (s *SchemaRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SchemaRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SchemaRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.EnterSchemaRule(s)
	}
}

func (s *SchemaRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.ExitSchemaRule(s)
	}
}

func (p *ObjectApiParser) SchemaRule() (localctx ISchemaRuleContext) {
	this := p
	_ = this

	localctx = NewSchemaRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ObjectApiParserRULE_schemaRule)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(211)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ObjectApiParserT__15, ObjectApiParserT__16, ObjectApiParserT__17, ObjectApiParserT__18:
		{
			p.SetState(209)
			p.PrimitiveSchema()
		}

	case ObjectApiParserIDENTIFIER:
		{
			p.SetState(210)
			p.SymbolSchema()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ObjectApiParserT__13 {
		{
			p.SetState(213)
			p.ArrayRule()
		}

	}

	return localctx
}

// IArrayRuleContext is an interface to support dynamic dispatch.
type IArrayRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayRuleContext differentiates from other interfaces.
	IsArrayRuleContext()
}

type ArrayRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayRuleContext() *ArrayRuleContext {
	var p = new(ArrayRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ObjectApiParserRULE_arrayRule
	return p
}

func (*ArrayRuleContext) IsArrayRuleContext() {}

func NewArrayRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayRuleContext {
	var p = new(ArrayRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectApiParserRULE_arrayRule

	return p
}

func (s *ArrayRuleContext) GetParser() antlr.Parser { return s.parser }
func (s *ArrayRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.EnterArrayRule(s)
	}
}

func (s *ArrayRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.ExitArrayRule(s)
	}
}

func (p *ObjectApiParser) ArrayRule() (localctx IArrayRuleContext) {
	this := p
	_ = this

	localctx = NewArrayRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ObjectApiParserRULE_arrayRule)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(216)
		p.Match(ObjectApiParserT__13)
	}
	{
		p.SetState(217)
		p.Match(ObjectApiParserT__14)
	}

	return localctx
}

// IPrimitiveSchemaContext is an interface to support dynamic dispatch.
type IPrimitiveSchemaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsPrimitiveSchemaContext differentiates from other interfaces.
	IsPrimitiveSchemaContext()
}

type PrimitiveSchemaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyPrimitiveSchemaContext() *PrimitiveSchemaContext {
	var p = new(PrimitiveSchemaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ObjectApiParserRULE_primitiveSchema
	return p
}

func (*PrimitiveSchemaContext) IsPrimitiveSchemaContext() {}

func NewPrimitiveSchemaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitiveSchemaContext {
	var p = new(PrimitiveSchemaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectApiParserRULE_primitiveSchema

	return p
}

func (s *PrimitiveSchemaContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitiveSchemaContext) GetName() antlr.Token { return s.name }

func (s *PrimitiveSchemaContext) SetName(v antlr.Token) { s.name = v }

func (s *PrimitiveSchemaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitiveSchemaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitiveSchemaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.EnterPrimitiveSchema(s)
	}
}

func (s *PrimitiveSchemaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.ExitPrimitiveSchema(s)
	}
}

func (p *ObjectApiParser) PrimitiveSchema() (localctx IPrimitiveSchemaContext) {
	this := p
	_ = this

	localctx = NewPrimitiveSchemaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ObjectApiParserRULE_primitiveSchema)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(223)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ObjectApiParserT__15:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(219)

			var _m = p.Match(ObjectApiParserT__15)

			localctx.(*PrimitiveSchemaContext).name = _m
		}

	case ObjectApiParserT__16:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(220)

			var _m = p.Match(ObjectApiParserT__16)

			localctx.(*PrimitiveSchemaContext).name = _m
		}

	case ObjectApiParserT__17:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(221)

			var _m = p.Match(ObjectApiParserT__17)

			localctx.(*PrimitiveSchemaContext).name = _m
		}

	case ObjectApiParserT__18:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(222)

			var _m = p.Match(ObjectApiParserT__18)

			localctx.(*PrimitiveSchemaContext).name = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISymbolSchemaContext is an interface to support dynamic dispatch.
type ISymbolSchemaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsSymbolSchemaContext differentiates from other interfaces.
	IsSymbolSchemaContext()
}

type SymbolSchemaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptySymbolSchemaContext() *SymbolSchemaContext {
	var p = new(SymbolSchemaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ObjectApiParserRULE_symbolSchema
	return p
}

func (*SymbolSchemaContext) IsSymbolSchemaContext() {}

func NewSymbolSchemaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SymbolSchemaContext {
	var p = new(SymbolSchemaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectApiParserRULE_symbolSchema

	return p
}

func (s *SymbolSchemaContext) GetParser() antlr.Parser { return s.parser }

func (s *SymbolSchemaContext) GetName() antlr.Token { return s.name }

func (s *SymbolSchemaContext) SetName(v antlr.Token) { s.name = v }

func (s *SymbolSchemaContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ObjectApiParserIDENTIFIER, 0)
}

func (s *SymbolSchemaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SymbolSchemaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SymbolSchemaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.EnterSymbolSchema(s)
	}
}

func (s *SymbolSchemaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.ExitSymbolSchema(s)
	}
}

func (p *ObjectApiParser) SymbolSchema() (localctx ISymbolSchemaContext) {
	this := p
	_ = this

	localctx = NewSymbolSchemaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ObjectApiParserRULE_symbolSchema)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(225)

		var _m = p.Match(ObjectApiParserIDENTIFIER)

		localctx.(*SymbolSchemaContext).name = _m
	}

	return localctx
}

// IMetaRuleContext is an interface to support dynamic dispatch.
type IMetaRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTagLine returns the tagLine token.
	GetTagLine() antlr.Token

	// GetDocLine returns the docLine token.
	GetDocLine() antlr.Token

	// SetTagLine sets the tagLine token.
	SetTagLine(antlr.Token)

	// SetDocLine sets the docLine token.
	SetDocLine(antlr.Token)

	// IsMetaRuleContext differentiates from other interfaces.
	IsMetaRuleContext()
}

type MetaRuleContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	tagLine antlr.Token
	docLine antlr.Token
}

func NewEmptyMetaRuleContext() *MetaRuleContext {
	var p = new(MetaRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ObjectApiParserRULE_metaRule
	return p
}

func (*MetaRuleContext) IsMetaRuleContext() {}

func NewMetaRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MetaRuleContext {
	var p = new(MetaRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectApiParserRULE_metaRule

	return p
}

func (s *MetaRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *MetaRuleContext) GetTagLine() antlr.Token { return s.tagLine }

func (s *MetaRuleContext) GetDocLine() antlr.Token { return s.docLine }

func (s *MetaRuleContext) SetTagLine(v antlr.Token) { s.tagLine = v }

func (s *MetaRuleContext) SetDocLine(v antlr.Token) { s.docLine = v }

func (s *MetaRuleContext) TAGLINE() antlr.TerminalNode {
	return s.GetToken(ObjectApiParserTAGLINE, 0)
}

func (s *MetaRuleContext) DOCLINE() antlr.TerminalNode {
	return s.GetToken(ObjectApiParserDOCLINE, 0)
}

func (s *MetaRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MetaRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MetaRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.EnterMetaRule(s)
	}
}

func (s *MetaRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.ExitMetaRule(s)
	}
}

func (p *ObjectApiParser) MetaRule() (localctx IMetaRuleContext) {
	this := p
	_ = this

	localctx = NewMetaRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ObjectApiParserRULE_metaRule)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(229)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ObjectApiParserTAGLINE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(227)

			var _m = p.Match(ObjectApiParserTAGLINE)

			localctx.(*MetaRuleContext).tagLine = _m
		}

	case ObjectApiParserDOCLINE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(228)

			var _m = p.Match(ObjectApiParserDOCLINE)

			localctx.(*MetaRuleContext).docLine = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
