// Code generated from pkg/idl/parser/ObjectApi.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ObjectApi

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
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
		"VERSION",
	}
	staticData.ruleNames = []string{
		"documentRule", "headerRule", "moduleRule", "importRule", "declarationsRule",
		"interfaceRule", "interfaceMembersRule", "propertyRule", "methodRule",
		"outputRule", "inputRule", "signalRule", "structRule", "structFieldRule",
		"enumRule", "enumMemberRule", "schemaRule", "arrayRule", "primitiveSchema",
		"symbolSchema",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 25, 172, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 1, 0, 1, 0, 5,
		0, 43, 8, 0, 10, 0, 12, 0, 46, 9, 0, 1, 1, 1, 1, 5, 1, 50, 8, 1, 10, 1,
		12, 1, 53, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 4, 3, 4, 66, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 72, 8, 5, 10,
		5, 12, 5, 75, 9, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 3, 6, 82, 8, 6, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 5, 8, 91, 8, 8, 10, 8, 12, 8, 94, 9,
		8, 1, 8, 1, 8, 3, 8, 98, 8, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1,
		10, 3, 10, 107, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 113, 8, 11, 10,
		11, 12, 11, 116, 9, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12,
		124, 8, 12, 10, 12, 12, 12, 127, 9, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 139, 8, 14, 10, 14, 12, 14,
		142, 9, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 3, 15, 149, 8, 15, 1, 15,
		3, 15, 152, 8, 15, 1, 16, 1, 16, 3, 16, 156, 8, 16, 1, 16, 3, 16, 159,
		8, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 168, 8,
		18, 1, 19, 1, 19, 1, 19, 0, 0, 20, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
		22, 24, 26, 28, 30, 32, 34, 36, 38, 0, 0, 171, 0, 40, 1, 0, 0, 0, 2, 47,
		1, 0, 0, 0, 4, 54, 1, 0, 0, 0, 6, 58, 1, 0, 0, 0, 8, 65, 1, 0, 0, 0, 10,
		67, 1, 0, 0, 0, 12, 81, 1, 0, 0, 0, 14, 83, 1, 0, 0, 0, 16, 87, 1, 0, 0,
		0, 18, 99, 1, 0, 0, 0, 20, 102, 1, 0, 0, 0, 22, 108, 1, 0, 0, 0, 24, 119,
		1, 0, 0, 0, 26, 130, 1, 0, 0, 0, 28, 134, 1, 0, 0, 0, 30, 145, 1, 0, 0,
		0, 32, 155, 1, 0, 0, 0, 34, 160, 1, 0, 0, 0, 36, 167, 1, 0, 0, 0, 38, 169,
		1, 0, 0, 0, 40, 44, 3, 2, 1, 0, 41, 43, 3, 8, 4, 0, 42, 41, 1, 0, 0, 0,
		43, 46, 1, 0, 0, 0, 44, 42, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 1, 1, 0,
		0, 0, 46, 44, 1, 0, 0, 0, 47, 51, 3, 4, 2, 0, 48, 50, 3, 6, 3, 0, 49, 48,
		1, 0, 0, 0, 50, 53, 1, 0, 0, 0, 51, 49, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0,
		52, 3, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 54, 55, 5, 1, 0, 0, 55, 56, 5, 24,
		0, 0, 56, 57, 5, 25, 0, 0, 57, 5, 1, 0, 0, 0, 58, 59, 5, 2, 0, 0, 59, 60,
		5, 24, 0, 0, 60, 61, 5, 25, 0, 0, 61, 7, 1, 0, 0, 0, 62, 66, 3, 10, 5,
		0, 63, 66, 3, 24, 12, 0, 64, 66, 3, 28, 14, 0, 65, 62, 1, 0, 0, 0, 65,
		63, 1, 0, 0, 0, 65, 64, 1, 0, 0, 0, 66, 9, 1, 0, 0, 0, 67, 68, 5, 3, 0,
		0, 68, 69, 5, 24, 0, 0, 69, 73, 5, 4, 0, 0, 70, 72, 3, 12, 6, 0, 71, 70,
		1, 0, 0, 0, 72, 75, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0,
		74, 76, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 76, 77, 5, 5, 0, 0, 77, 11, 1,
		0, 0, 0, 78, 82, 3, 14, 7, 0, 79, 82, 3, 16, 8, 0, 80, 82, 3, 22, 11, 0,
		81, 78, 1, 0, 0, 0, 81, 79, 1, 0, 0, 0, 81, 80, 1, 0, 0, 0, 82, 13, 1,
		0, 0, 0, 83, 84, 5, 24, 0, 0, 84, 85, 5, 6, 0, 0, 85, 86, 3, 32, 16, 0,
		86, 15, 1, 0, 0, 0, 87, 88, 5, 24, 0, 0, 88, 92, 5, 7, 0, 0, 89, 91, 3,
		20, 10, 0, 90, 89, 1, 0, 0, 0, 91, 94, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0,
		92, 93, 1, 0, 0, 0, 93, 95, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 95, 97, 5,
		8, 0, 0, 96, 98, 3, 18, 9, 0, 97, 96, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98,
		17, 1, 0, 0, 0, 99, 100, 5, 6, 0, 0, 100, 101, 3, 32, 16, 0, 101, 19, 1,
		0, 0, 0, 102, 103, 5, 24, 0, 0, 103, 104, 5, 6, 0, 0, 104, 106, 3, 32,
		16, 0, 105, 107, 5, 9, 0, 0, 106, 105, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0,
		107, 21, 1, 0, 0, 0, 108, 109, 5, 10, 0, 0, 109, 110, 5, 24, 0, 0, 110,
		114, 5, 7, 0, 0, 111, 113, 3, 20, 10, 0, 112, 111, 1, 0, 0, 0, 113, 116,
		1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 117, 1, 0,
		0, 0, 116, 114, 1, 0, 0, 0, 117, 118, 5, 8, 0, 0, 118, 23, 1, 0, 0, 0,
		119, 120, 5, 11, 0, 0, 120, 121, 5, 24, 0, 0, 121, 125, 5, 4, 0, 0, 122,
		124, 3, 26, 13, 0, 123, 122, 1, 0, 0, 0, 124, 127, 1, 0, 0, 0, 125, 123,
		1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 128, 1, 0, 0, 0, 127, 125, 1, 0,
		0, 0, 128, 129, 5, 5, 0, 0, 129, 25, 1, 0, 0, 0, 130, 131, 5, 24, 0, 0,
		131, 132, 5, 6, 0, 0, 132, 133, 3, 32, 16, 0, 133, 27, 1, 0, 0, 0, 134,
		135, 5, 12, 0, 0, 135, 136, 5, 24, 0, 0, 136, 140, 5, 4, 0, 0, 137, 139,
		3, 30, 15, 0, 138, 137, 1, 0, 0, 0, 139, 142, 1, 0, 0, 0, 140, 138, 1,
		0, 0, 0, 140, 141, 1, 0, 0, 0, 141, 143, 1, 0, 0, 0, 142, 140, 1, 0, 0,
		0, 143, 144, 5, 5, 0, 0, 144, 29, 1, 0, 0, 0, 145, 148, 5, 24, 0, 0, 146,
		147, 5, 13, 0, 0, 147, 149, 5, 21, 0, 0, 148, 146, 1, 0, 0, 0, 148, 149,
		1, 0, 0, 0, 149, 151, 1, 0, 0, 0, 150, 152, 5, 9, 0, 0, 151, 150, 1, 0,
		0, 0, 151, 152, 1, 0, 0, 0, 152, 31, 1, 0, 0, 0, 153, 156, 3, 36, 18, 0,
		154, 156, 3, 38, 19, 0, 155, 153, 1, 0, 0, 0, 155, 154, 1, 0, 0, 0, 156,
		158, 1, 0, 0, 0, 157, 159, 3, 34, 17, 0, 158, 157, 1, 0, 0, 0, 158, 159,
		1, 0, 0, 0, 159, 33, 1, 0, 0, 0, 160, 161, 5, 14, 0, 0, 161, 162, 5, 15,
		0, 0, 162, 35, 1, 0, 0, 0, 163, 168, 5, 16, 0, 0, 164, 168, 5, 17, 0, 0,
		165, 168, 5, 18, 0, 0, 166, 168, 5, 19, 0, 0, 167, 163, 1, 0, 0, 0, 167,
		164, 1, 0, 0, 0, 167, 165, 1, 0, 0, 0, 167, 166, 1, 0, 0, 0, 168, 37, 1,
		0, 0, 0, 169, 170, 5, 24, 0, 0, 170, 39, 1, 0, 0, 0, 16, 44, 51, 65, 73,
		81, 92, 97, 106, 114, 125, 140, 148, 151, 155, 158, 167,
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
	this.GrammarFileName = "ObjectApi.g4"

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
	ObjectApiParserRULE_methodRule           = 8
	ObjectApiParserRULE_outputRule           = 9
	ObjectApiParserRULE_inputRule            = 10
	ObjectApiParserRULE_signalRule           = 11
	ObjectApiParserRULE_structRule           = 12
	ObjectApiParserRULE_structFieldRule      = 13
	ObjectApiParserRULE_enumRule             = 14
	ObjectApiParserRULE_enumMemberRule       = 15
	ObjectApiParserRULE_schemaRule           = 16
	ObjectApiParserRULE_arrayRule            = 17
	ObjectApiParserRULE_primitiveSchema      = 18
	ObjectApiParserRULE_symbolSchema         = 19
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
		p.SetState(40)
		p.HeaderRule()
	}
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ObjectApiParserT__2)|(1<<ObjectApiParserT__10)|(1<<ObjectApiParserT__11))) != 0 {
		{
			p.SetState(41)
			p.DeclarationsRule()
		}

		p.SetState(46)
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
		p.SetState(47)
		p.ModuleRule()
	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ObjectApiParserT__1 {
		{
			p.SetState(48)
			p.ImportRule()
		}

		p.SetState(53)
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
		p.SetState(54)
		p.Match(ObjectApiParserT__0)
	}
	{
		p.SetState(55)

		var _m = p.Match(ObjectApiParserIDENTIFIER)

		localctx.(*ModuleRuleContext).name = _m
	}
	{
		p.SetState(56)

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
		p.SetState(58)
		p.Match(ObjectApiParserT__1)
	}
	{
		p.SetState(59)

		var _m = p.Match(ObjectApiParserIDENTIFIER)

		localctx.(*ImportRuleContext).name = _m
	}
	{
		p.SetState(60)

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

	p.SetState(65)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ObjectApiParserT__2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(62)
			p.InterfaceRule()
		}

	case ObjectApiParserT__10:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(63)
			p.StructRule()
		}

	case ObjectApiParserT__11:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(64)
			p.EnumRule()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	{
		p.SetState(67)
		p.Match(ObjectApiParserT__2)
	}
	{
		p.SetState(68)

		var _m = p.Match(ObjectApiParserIDENTIFIER)

		localctx.(*InterfaceRuleContext).name = _m
	}
	{
		p.SetState(69)
		p.Match(ObjectApiParserT__3)
	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ObjectApiParserT__9 || _la == ObjectApiParserIDENTIFIER {
		{
			p.SetState(70)
			p.InterfaceMembersRule()
		}

		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(76)
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

func (s *InterfaceMembersRuleContext) MethodRule() IMethodRuleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodRuleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethodRuleContext)
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

	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(78)
			p.PropertyRule()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(79)
			p.MethodRule()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(80)
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
		p.SetState(83)

		var _m = p.Match(ObjectApiParserIDENTIFIER)

		localctx.(*PropertyRuleContext).name = _m
	}
	{
		p.SetState(84)
		p.Match(ObjectApiParserT__5)
	}
	{
		p.SetState(85)

		var _x = p.SchemaRule()

		localctx.(*PropertyRuleContext).schema = _x
	}

	return localctx
}

// IMethodRuleContext is an interface to support dynamic dispatch.
type IMethodRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// GetInputs returns the inputs rule contexts.
	GetInputs() IInputRuleContext

	// SetInputs sets the inputs rule contexts.
	SetInputs(IInputRuleContext)

	// IsMethodRuleContext differentiates from other interfaces.
	IsMethodRuleContext()
}

type MethodRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	inputs IInputRuleContext
}

func NewEmptyMethodRuleContext() *MethodRuleContext {
	var p = new(MethodRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ObjectApiParserRULE_methodRule
	return p
}

func (*MethodRuleContext) IsMethodRuleContext() {}

func NewMethodRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodRuleContext {
	var p = new(MethodRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectApiParserRULE_methodRule

	return p
}

func (s *MethodRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodRuleContext) GetName() antlr.Token { return s.name }

func (s *MethodRuleContext) SetName(v antlr.Token) { s.name = v }

func (s *MethodRuleContext) GetInputs() IInputRuleContext { return s.inputs }

func (s *MethodRuleContext) SetInputs(v IInputRuleContext) { s.inputs = v }

func (s *MethodRuleContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ObjectApiParserIDENTIFIER, 0)
}

func (s *MethodRuleContext) OutputRule() IOutputRuleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOutputRuleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOutputRuleContext)
}

func (s *MethodRuleContext) AllInputRule() []IInputRuleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInputRuleContext); ok {
			len++
		}
	}

	tst := make([]IInputRuleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInputRuleContext); ok {
			tst[i] = t.(IInputRuleContext)
			i++
		}
	}

	return tst
}

func (s *MethodRuleContext) InputRule(i int) IInputRuleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInputRuleContext); ok {
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

	return t.(IInputRuleContext)
}

func (s *MethodRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.EnterMethodRule(s)
	}
}

func (s *MethodRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.ExitMethodRule(s)
	}
}

func (p *ObjectApiParser) MethodRule() (localctx IMethodRuleContext) {
	this := p
	_ = this

	localctx = NewMethodRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ObjectApiParserRULE_methodRule)
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
		p.SetState(87)

		var _m = p.Match(ObjectApiParserIDENTIFIER)

		localctx.(*MethodRuleContext).name = _m
	}
	{
		p.SetState(88)
		p.Match(ObjectApiParserT__6)
	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ObjectApiParserIDENTIFIER {
		{
			p.SetState(89)

			var _x = p.InputRule()

			localctx.(*MethodRuleContext).inputs = _x
		}

		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(95)
		p.Match(ObjectApiParserT__7)
	}
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ObjectApiParserT__5 {
		{
			p.SetState(96)
			p.OutputRule()
		}

	}

	return localctx
}

// IOutputRuleContext is an interface to support dynamic dispatch.
type IOutputRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSchema returns the schema rule contexts.
	GetSchema() ISchemaRuleContext

	// SetSchema sets the schema rule contexts.
	SetSchema(ISchemaRuleContext)

	// IsOutputRuleContext differentiates from other interfaces.
	IsOutputRuleContext()
}

type OutputRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	schema ISchemaRuleContext
}

func NewEmptyOutputRuleContext() *OutputRuleContext {
	var p = new(OutputRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ObjectApiParserRULE_outputRule
	return p
}

func (*OutputRuleContext) IsOutputRuleContext() {}

func NewOutputRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OutputRuleContext {
	var p = new(OutputRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectApiParserRULE_outputRule

	return p
}

func (s *OutputRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *OutputRuleContext) GetSchema() ISchemaRuleContext { return s.schema }

func (s *OutputRuleContext) SetSchema(v ISchemaRuleContext) { s.schema = v }

func (s *OutputRuleContext) SchemaRule() ISchemaRuleContext {
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

func (s *OutputRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OutputRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OutputRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.EnterOutputRule(s)
	}
}

func (s *OutputRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.ExitOutputRule(s)
	}
}

func (p *ObjectApiParser) OutputRule() (localctx IOutputRuleContext) {
	this := p
	_ = this

	localctx = NewOutputRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ObjectApiParserRULE_outputRule)

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
		p.SetState(99)
		p.Match(ObjectApiParserT__5)
	}
	{
		p.SetState(100)

		var _x = p.SchemaRule()

		localctx.(*OutputRuleContext).schema = _x
	}

	return localctx
}

// IInputRuleContext is an interface to support dynamic dispatch.
type IInputRuleContext interface {
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

	// IsInputRuleContext differentiates from other interfaces.
	IsInputRuleContext()
}

type InputRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	schema ISchemaRuleContext
}

func NewEmptyInputRuleContext() *InputRuleContext {
	var p = new(InputRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ObjectApiParserRULE_inputRule
	return p
}

func (*InputRuleContext) IsInputRuleContext() {}

func NewInputRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputRuleContext {
	var p = new(InputRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectApiParserRULE_inputRule

	return p
}

func (s *InputRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *InputRuleContext) GetName() antlr.Token { return s.name }

func (s *InputRuleContext) SetName(v antlr.Token) { s.name = v }

func (s *InputRuleContext) GetSchema() ISchemaRuleContext { return s.schema }

func (s *InputRuleContext) SetSchema(v ISchemaRuleContext) { s.schema = v }

func (s *InputRuleContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ObjectApiParserIDENTIFIER, 0)
}

func (s *InputRuleContext) SchemaRule() ISchemaRuleContext {
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

func (s *InputRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.EnterInputRule(s)
	}
}

func (s *InputRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.ExitInputRule(s)
	}
}

func (p *ObjectApiParser) InputRule() (localctx IInputRuleContext) {
	this := p
	_ = this

	localctx = NewInputRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ObjectApiParserRULE_inputRule)
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
		p.SetState(102)

		var _m = p.Match(ObjectApiParserIDENTIFIER)

		localctx.(*InputRuleContext).name = _m
	}
	{
		p.SetState(103)
		p.Match(ObjectApiParserT__5)
	}
	{
		p.SetState(104)

		var _x = p.SchemaRule()

		localctx.(*InputRuleContext).schema = _x
	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ObjectApiParserT__8 {
		{
			p.SetState(105)
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

	// GetInputs returns the inputs rule contexts.
	GetInputs() IInputRuleContext

	// SetInputs sets the inputs rule contexts.
	SetInputs(IInputRuleContext)

	// IsSignalRuleContext differentiates from other interfaces.
	IsSignalRuleContext()
}

type SignalRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	inputs IInputRuleContext
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

func (s *SignalRuleContext) GetInputs() IInputRuleContext { return s.inputs }

func (s *SignalRuleContext) SetInputs(v IInputRuleContext) { s.inputs = v }

func (s *SignalRuleContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ObjectApiParserIDENTIFIER, 0)
}

func (s *SignalRuleContext) AllInputRule() []IInputRuleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInputRuleContext); ok {
			len++
		}
	}

	tst := make([]IInputRuleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInputRuleContext); ok {
			tst[i] = t.(IInputRuleContext)
			i++
		}
	}

	return tst
}

func (s *SignalRuleContext) InputRule(i int) IInputRuleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInputRuleContext); ok {
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

	return t.(IInputRuleContext)
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
	{
		p.SetState(108)
		p.Match(ObjectApiParserT__9)
	}
	{
		p.SetState(109)

		var _m = p.Match(ObjectApiParserIDENTIFIER)

		localctx.(*SignalRuleContext).name = _m
	}
	{
		p.SetState(110)
		p.Match(ObjectApiParserT__6)
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ObjectApiParserIDENTIFIER {
		{
			p.SetState(111)

			var _x = p.InputRule()

			localctx.(*SignalRuleContext).inputs = _x
		}

		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(117)
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
	{
		p.SetState(119)
		p.Match(ObjectApiParserT__10)
	}
	{
		p.SetState(120)

		var _m = p.Match(ObjectApiParserIDENTIFIER)

		localctx.(*StructRuleContext).name = _m
	}
	{
		p.SetState(121)
		p.Match(ObjectApiParserT__3)
	}
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ObjectApiParserIDENTIFIER {
		{
			p.SetState(122)
			p.StructFieldRule()
		}

		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(128)
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
		p.SetState(130)

		var _m = p.Match(ObjectApiParserIDENTIFIER)

		localctx.(*StructFieldRuleContext).name = _m
	}
	{
		p.SetState(131)
		p.Match(ObjectApiParserT__5)
	}
	{
		p.SetState(132)

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
	{
		p.SetState(134)
		p.Match(ObjectApiParserT__11)
	}
	{
		p.SetState(135)

		var _m = p.Match(ObjectApiParserIDENTIFIER)

		localctx.(*EnumRuleContext).name = _m
	}
	{
		p.SetState(136)
		p.Match(ObjectApiParserT__3)
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ObjectApiParserIDENTIFIER {
		{
			p.SetState(137)
			p.EnumMemberRule()
		}

		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(143)
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
	{
		p.SetState(145)

		var _m = p.Match(ObjectApiParserIDENTIFIER)

		localctx.(*EnumMemberRuleContext).name = _m
	}
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ObjectApiParserT__12 {
		{
			p.SetState(146)
			p.Match(ObjectApiParserT__12)
		}
		{
			p.SetState(147)

			var _m = p.Match(ObjectApiParserINTEGER)

			localctx.(*EnumMemberRuleContext).value = _m
		}

	}
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ObjectApiParserT__8 {
		{
			p.SetState(150)
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
	p.SetState(155)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ObjectApiParserT__15, ObjectApiParserT__16, ObjectApiParserT__17, ObjectApiParserT__18:
		{
			p.SetState(153)
			p.PrimitiveSchema()
		}

	case ObjectApiParserIDENTIFIER:
		{
			p.SetState(154)
			p.SymbolSchema()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ObjectApiParserT__13 {
		{
			p.SetState(157)
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
		p.SetState(160)
		p.Match(ObjectApiParserT__13)
	}
	{
		p.SetState(161)
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

	p.SetState(167)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ObjectApiParserT__15:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(163)

			var _m = p.Match(ObjectApiParserT__15)

			localctx.(*PrimitiveSchemaContext).name = _m
		}

	case ObjectApiParserT__16:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(164)

			var _m = p.Match(ObjectApiParserT__16)

			localctx.(*PrimitiveSchemaContext).name = _m
		}

	case ObjectApiParserT__17:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(165)

			var _m = p.Match(ObjectApiParserT__17)

			localctx.(*PrimitiveSchemaContext).name = _m
		}

	case ObjectApiParserT__18:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(166)

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
		p.SetState(169)

		var _m = p.Match(ObjectApiParserIDENTIFIER)

		localctx.(*SymbolSchemaContext).name = _m
	}

	return localctx
}
