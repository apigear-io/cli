// Code generated from pkg/idl/parser/ObjectApi.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // ObjectApi

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 27, 174,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 3, 2, 3, 2, 7, 2, 45, 10,
	2, 12, 2, 14, 2, 48, 11, 2, 3, 3, 3, 3, 7, 3, 52, 10, 3, 12, 3, 14, 3,
	55, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6,
	3, 6, 5, 6, 68, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 74, 10, 7, 12, 7,
	14, 7, 77, 11, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 5, 8, 84, 10, 8, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 7, 10, 93, 10, 10, 12, 10, 14, 10,
	96, 11, 10, 3, 10, 3, 10, 5, 10, 100, 10, 10, 3, 11, 3, 11, 3, 11, 3, 12,
	3, 12, 3, 12, 3, 12, 5, 12, 109, 10, 12, 3, 13, 3, 13, 3, 13, 3, 13, 7,
	13, 115, 10, 13, 12, 13, 14, 13, 118, 11, 13, 3, 13, 3, 13, 3, 14, 3, 14,
	3, 14, 3, 14, 7, 14, 126, 10, 14, 12, 14, 14, 14, 129, 11, 14, 3, 14, 3,
	14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 7, 16, 141,
	10, 16, 12, 16, 14, 16, 144, 11, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17,
	5, 17, 151, 10, 17, 3, 17, 5, 17, 154, 10, 17, 3, 18, 3, 18, 5, 18, 158,
	10, 18, 3, 18, 5, 18, 161, 10, 18, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3,
	20, 3, 20, 5, 20, 170, 10, 20, 3, 21, 3, 21, 3, 21, 2, 2, 22, 2, 4, 6,
	8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 2, 2,
	2, 173, 2, 42, 3, 2, 2, 2, 4, 49, 3, 2, 2, 2, 6, 56, 3, 2, 2, 2, 8, 60,
	3, 2, 2, 2, 10, 67, 3, 2, 2, 2, 12, 69, 3, 2, 2, 2, 14, 83, 3, 2, 2, 2,
	16, 85, 3, 2, 2, 2, 18, 89, 3, 2, 2, 2, 20, 101, 3, 2, 2, 2, 22, 104, 3,
	2, 2, 2, 24, 110, 3, 2, 2, 2, 26, 121, 3, 2, 2, 2, 28, 132, 3, 2, 2, 2,
	30, 136, 3, 2, 2, 2, 32, 147, 3, 2, 2, 2, 34, 157, 3, 2, 2, 2, 36, 162,
	3, 2, 2, 2, 38, 169, 3, 2, 2, 2, 40, 171, 3, 2, 2, 2, 42, 46, 5, 4, 3,
	2, 43, 45, 5, 10, 6, 2, 44, 43, 3, 2, 2, 2, 45, 48, 3, 2, 2, 2, 46, 44,
	3, 2, 2, 2, 46, 47, 3, 2, 2, 2, 47, 3, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2,
	49, 53, 5, 6, 4, 2, 50, 52, 5, 8, 5, 2, 51, 50, 3, 2, 2, 2, 52, 55, 3,
	2, 2, 2, 53, 51, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 5, 3, 2, 2, 2, 55,
	53, 3, 2, 2, 2, 56, 57, 7, 3, 2, 2, 57, 58, 7, 26, 2, 2, 58, 59, 7, 27,
	2, 2, 59, 7, 3, 2, 2, 2, 60, 61, 7, 4, 2, 2, 61, 62, 7, 26, 2, 2, 62, 63,
	7, 27, 2, 2, 63, 9, 3, 2, 2, 2, 64, 68, 5, 12, 7, 2, 65, 68, 5, 26, 14,
	2, 66, 68, 5, 30, 16, 2, 67, 64, 3, 2, 2, 2, 67, 65, 3, 2, 2, 2, 67, 66,
	3, 2, 2, 2, 68, 11, 3, 2, 2, 2, 69, 70, 7, 5, 2, 2, 70, 71, 7, 26, 2, 2,
	71, 75, 7, 6, 2, 2, 72, 74, 5, 14, 8, 2, 73, 72, 3, 2, 2, 2, 74, 77, 3,
	2, 2, 2, 75, 73, 3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 78, 3, 2, 2, 2, 77,
	75, 3, 2, 2, 2, 78, 79, 7, 7, 2, 2, 79, 13, 3, 2, 2, 2, 80, 84, 5, 16,
	9, 2, 81, 84, 5, 18, 10, 2, 82, 84, 5, 24, 13, 2, 83, 80, 3, 2, 2, 2, 83,
	81, 3, 2, 2, 2, 83, 82, 3, 2, 2, 2, 84, 15, 3, 2, 2, 2, 85, 86, 7, 26,
	2, 2, 86, 87, 7, 8, 2, 2, 87, 88, 5, 34, 18, 2, 88, 17, 3, 2, 2, 2, 89,
	90, 7, 26, 2, 2, 90, 94, 7, 9, 2, 2, 91, 93, 5, 22, 12, 2, 92, 91, 3, 2,
	2, 2, 93, 96, 3, 2, 2, 2, 94, 92, 3, 2, 2, 2, 94, 95, 3, 2, 2, 2, 95, 97,
	3, 2, 2, 2, 96, 94, 3, 2, 2, 2, 97, 99, 7, 10, 2, 2, 98, 100, 5, 20, 11,
	2, 99, 98, 3, 2, 2, 2, 99, 100, 3, 2, 2, 2, 100, 19, 3, 2, 2, 2, 101, 102,
	7, 8, 2, 2, 102, 103, 5, 34, 18, 2, 103, 21, 3, 2, 2, 2, 104, 105, 7, 26,
	2, 2, 105, 106, 7, 8, 2, 2, 106, 108, 5, 34, 18, 2, 107, 109, 7, 11, 2,
	2, 108, 107, 3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109, 23, 3, 2, 2, 2, 110,
	111, 7, 12, 2, 2, 111, 112, 7, 26, 2, 2, 112, 116, 7, 9, 2, 2, 113, 115,
	5, 22, 12, 2, 114, 113, 3, 2, 2, 2, 115, 118, 3, 2, 2, 2, 116, 114, 3,
	2, 2, 2, 116, 117, 3, 2, 2, 2, 117, 119, 3, 2, 2, 2, 118, 116, 3, 2, 2,
	2, 119, 120, 7, 10, 2, 2, 120, 25, 3, 2, 2, 2, 121, 122, 7, 13, 2, 2, 122,
	123, 7, 26, 2, 2, 123, 127, 7, 6, 2, 2, 124, 126, 5, 28, 15, 2, 125, 124,
	3, 2, 2, 2, 126, 129, 3, 2, 2, 2, 127, 125, 3, 2, 2, 2, 127, 128, 3, 2,
	2, 2, 128, 130, 3, 2, 2, 2, 129, 127, 3, 2, 2, 2, 130, 131, 7, 7, 2, 2,
	131, 27, 3, 2, 2, 2, 132, 133, 7, 26, 2, 2, 133, 134, 7, 8, 2, 2, 134,
	135, 5, 34, 18, 2, 135, 29, 3, 2, 2, 2, 136, 137, 7, 14, 2, 2, 137, 138,
	7, 26, 2, 2, 138, 142, 7, 6, 2, 2, 139, 141, 5, 32, 17, 2, 140, 139, 3,
	2, 2, 2, 141, 144, 3, 2, 2, 2, 142, 140, 3, 2, 2, 2, 142, 143, 3, 2, 2,
	2, 143, 145, 3, 2, 2, 2, 144, 142, 3, 2, 2, 2, 145, 146, 7, 7, 2, 2, 146,
	31, 3, 2, 2, 2, 147, 150, 7, 26, 2, 2, 148, 149, 7, 15, 2, 2, 149, 151,
	7, 23, 2, 2, 150, 148, 3, 2, 2, 2, 150, 151, 3, 2, 2, 2, 151, 153, 3, 2,
	2, 2, 152, 154, 7, 11, 2, 2, 153, 152, 3, 2, 2, 2, 153, 154, 3, 2, 2, 2,
	154, 33, 3, 2, 2, 2, 155, 158, 5, 38, 20, 2, 156, 158, 5, 40, 21, 2, 157,
	155, 3, 2, 2, 2, 157, 156, 3, 2, 2, 2, 158, 160, 3, 2, 2, 2, 159, 161,
	5, 36, 19, 2, 160, 159, 3, 2, 2, 2, 160, 161, 3, 2, 2, 2, 161, 35, 3, 2,
	2, 2, 162, 163, 7, 16, 2, 2, 163, 164, 7, 17, 2, 2, 164, 37, 3, 2, 2, 2,
	165, 170, 7, 18, 2, 2, 166, 170, 7, 19, 2, 2, 167, 170, 7, 20, 2, 2, 168,
	170, 7, 21, 2, 2, 169, 165, 3, 2, 2, 2, 169, 166, 3, 2, 2, 2, 169, 167,
	3, 2, 2, 2, 169, 168, 3, 2, 2, 2, 170, 39, 3, 2, 2, 2, 171, 172, 7, 26,
	2, 2, 172, 41, 3, 2, 2, 2, 18, 46, 53, 67, 75, 83, 94, 99, 108, 116, 127,
	142, 150, 153, 157, 160, 169,
}
var literalNames = []string{
	"", "'module'", "'import'", "'interface'", "'{'", "'}'", "':'", "'('",
	"')'", "','", "'signal'", "'struct'", "'enum'", "'='", "'['", "']'", "'bool'",
	"'int'", "'float'", "'string'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "WHITESPACE", "INTEGER", "HEX", "TYPE_IDENTIFIER", "IDENTIFIER",
	"VERSION",
}

var ruleNames = []string{
	"documentRule", "headerRule", "moduleRule", "importRule", "declarationsRule",
	"interfaceRule", "interfaceMembersRule", "propertyRule", "methodRule",
	"outputRule", "inputRule", "signalRule", "structRule", "structFieldRule",
	"enumRule", "enumMemberRule", "schemaRule", "arrayRule", "primitiveSchema",
	"symbolSchema",
}

type ObjectApiParser struct {
	*antlr.BaseParser
}

// NewObjectApiParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *ObjectApiParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewObjectApiParser(input antlr.TokenStream) *ObjectApiParser {
	this := new(ObjectApiParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHeaderRuleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHeaderRuleContext)
}

func (s *DocumentRuleContext) AllDeclarationsRule() []IDeclarationsRuleContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDeclarationsRuleContext)(nil)).Elem())
	var tst = make([]IDeclarationsRuleContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDeclarationsRuleContext)
		}
	}

	return tst
}

func (s *DocumentRuleContext) DeclarationsRule(i int) IDeclarationsRuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationsRuleContext)(nil)).Elem(), i)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModuleRuleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModuleRuleContext)
}

func (s *HeaderRuleContext) AllImportRule() []IImportRuleContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IImportRuleContext)(nil)).Elem())
	var tst = make([]IImportRuleContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IImportRuleContext)
		}
	}

	return tst
}

func (s *HeaderRuleContext) ImportRule(i int) IImportRuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportRuleContext)(nil)).Elem(), i)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInterfaceRuleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInterfaceRuleContext)
}

func (s *DeclarationsRuleContext) StructRule() IStructRuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructRuleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructRuleContext)
}

func (s *DeclarationsRuleContext) EnumRule() IEnumRuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumRuleContext)(nil)).Elem(), 0)

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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInterfaceMembersRuleContext)(nil)).Elem())
	var tst = make([]IInterfaceMembersRuleContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInterfaceMembersRuleContext)
		}
	}

	return tst
}

func (s *InterfaceRuleContext) InterfaceMembersRule(i int) IInterfaceMembersRuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInterfaceMembersRuleContext)(nil)).Elem(), i)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyRuleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyRuleContext)
}

func (s *InterfaceMembersRuleContext) MethodRule() IMethodRuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodRuleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodRuleContext)
}

func (s *InterfaceMembersRuleContext) SignalRule() ISignalRuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISignalRuleContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISchemaRuleContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOutputRuleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOutputRuleContext)
}

func (s *MethodRuleContext) AllInputRule() []IInputRuleContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInputRuleContext)(nil)).Elem())
	var tst = make([]IInputRuleContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInputRuleContext)
		}
	}

	return tst
}

func (s *MethodRuleContext) InputRule(i int) IInputRuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInputRuleContext)(nil)).Elem(), i)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISchemaRuleContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISchemaRuleContext)(nil)).Elem(), 0)

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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInputRuleContext)(nil)).Elem())
	var tst = make([]IInputRuleContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInputRuleContext)
		}
	}

	return tst
}

func (s *SignalRuleContext) InputRule(i int) IInputRuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInputRuleContext)(nil)).Elem(), i)

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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStructFieldRuleContext)(nil)).Elem())
	var tst = make([]IStructFieldRuleContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStructFieldRuleContext)
		}
	}

	return tst
}

func (s *StructRuleContext) StructFieldRule(i int) IStructFieldRuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructFieldRuleContext)(nil)).Elem(), i)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISchemaRuleContext)(nil)).Elem(), 0)

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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEnumMemberRuleContext)(nil)).Elem())
	var tst = make([]IEnumMemberRuleContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEnumMemberRuleContext)
		}
	}

	return tst
}

func (s *EnumRuleContext) EnumMemberRule(i int) IEnumMemberRuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumMemberRuleContext)(nil)).Elem(), i)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitiveSchemaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitiveSchemaContext)
}

func (s *SchemaRuleContext) SymbolSchema() ISymbolSchemaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISymbolSchemaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISymbolSchemaContext)
}

func (s *SchemaRuleContext) ArrayRule() IArrayRuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayRuleContext)(nil)).Elem(), 0)

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
