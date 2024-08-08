// Code generated from pkg/idl/parser/ObjectApi.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ObjectApi

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ObjectApiParser struct {
	*antlr.BaseParser
}

var ObjectApiParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func objectapiParserInit() {
	staticData := &ObjectApiParserStaticData
	staticData.LiteralNames = []string{
		"", "'module'", "'import'", "'extern'", "'interface'", "'extends'",
		"'{'", "'}'", "'readonly'", "':'", "'('", "')'", "','", "'signal'",
		"'struct'", "'enum'", "'='", "'['", "']'", "'bool'", "'int'", "'int32'",
		"'int64'", "'float'", "'float32'", "'float64'", "'string'", "'bytes'",
		"", "", "", "", "", "", "", "", "'.'", "", "", "'_'", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "WHITESPACE", "INTEGER",
		"HEX", "IDENTIFIER", "VERSION", "DOCLINE", "TAGLINE", "COMMENT", "DOT",
		"LETTER", "DIGIT", "UNDERSCORE", "SEMICOLON",
	}
	staticData.RuleNames = []string{
		"documentRule", "headerRule", "moduleRule", "importRule", "declarationsRule",
		"externRule", "interfaceRule", "interfaceMembersRule", "propertyRule",
		"operationRule", "operationReturnRule", "operationParamRule", "signalRule",
		"structRule", "structFieldRule", "enumRule", "enumMemberRule", "schemaRule",
		"arrayRule", "primitiveSchema", "symbolSchema", "metaRule",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 40, 278, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 1, 0, 1, 0, 5, 0, 47, 8, 0, 10, 0, 12, 0, 50, 9, 0, 1, 1, 1,
		1, 5, 1, 54, 8, 1, 10, 1, 12, 1, 57, 9, 1, 1, 2, 5, 2, 60, 8, 2, 10, 2,
		12, 2, 63, 9, 2, 1, 2, 1, 2, 1, 2, 3, 2, 68, 8, 2, 1, 2, 3, 2, 71, 8, 2,
		1, 3, 1, 3, 1, 3, 3, 3, 76, 8, 3, 1, 3, 3, 3, 79, 8, 3, 1, 4, 1, 4, 1,
		4, 1, 4, 3, 4, 85, 8, 4, 1, 5, 5, 5, 88, 8, 5, 10, 5, 12, 5, 91, 9, 5,
		1, 5, 1, 5, 1, 5, 3, 5, 96, 8, 5, 1, 6, 5, 6, 99, 8, 6, 10, 6, 12, 6, 102,
		9, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 108, 8, 6, 1, 6, 1, 6, 5, 6, 112, 8,
		6, 10, 6, 12, 6, 115, 9, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 3, 7, 122, 8,
		7, 1, 8, 5, 8, 125, 8, 8, 10, 8, 12, 8, 128, 9, 8, 1, 8, 3, 8, 131, 8,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 137, 8, 8, 1, 9, 5, 9, 140, 8, 9, 10,
		9, 12, 9, 143, 9, 9, 1, 9, 1, 9, 1, 9, 5, 9, 148, 8, 9, 10, 9, 12, 9, 151,
		9, 9, 1, 9, 1, 9, 3, 9, 155, 8, 9, 1, 9, 3, 9, 158, 8, 9, 1, 10, 1, 10,
		1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 167, 8, 11, 1, 12, 5, 12, 170,
		8, 12, 10, 12, 12, 12, 173, 9, 12, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 179,
		8, 12, 10, 12, 12, 12, 182, 9, 12, 1, 12, 1, 12, 3, 12, 186, 8, 12, 1,
		13, 5, 13, 189, 8, 13, 10, 13, 12, 13, 192, 9, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 5, 13, 198, 8, 13, 10, 13, 12, 13, 201, 9, 13, 1, 13, 1, 13, 1,
		14, 5, 14, 206, 8, 14, 10, 14, 12, 14, 209, 9, 14, 1, 14, 3, 14, 212, 8,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 218, 8, 14, 1, 15, 5, 15, 221, 8,
		15, 10, 15, 12, 15, 224, 9, 15, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 230,
		8, 15, 10, 15, 12, 15, 233, 9, 15, 1, 15, 1, 15, 1, 16, 5, 16, 238, 8,
		16, 10, 16, 12, 16, 241, 9, 16, 1, 16, 1, 16, 1, 16, 3, 16, 246, 8, 16,
		1, 16, 3, 16, 249, 8, 16, 1, 17, 1, 17, 3, 17, 253, 8, 17, 1, 17, 3, 17,
		256, 8, 17, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 3, 19, 270, 8, 19, 1, 20, 1, 20, 1, 21, 1, 21,
		3, 21, 276, 8, 21, 1, 21, 0, 0, 22, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18,
		20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 0, 0, 304, 0, 44, 1, 0,
		0, 0, 2, 51, 1, 0, 0, 0, 4, 61, 1, 0, 0, 0, 6, 72, 1, 0, 0, 0, 8, 84, 1,
		0, 0, 0, 10, 89, 1, 0, 0, 0, 12, 100, 1, 0, 0, 0, 14, 121, 1, 0, 0, 0,
		16, 126, 1, 0, 0, 0, 18, 141, 1, 0, 0, 0, 20, 159, 1, 0, 0, 0, 22, 162,
		1, 0, 0, 0, 24, 171, 1, 0, 0, 0, 26, 190, 1, 0, 0, 0, 28, 207, 1, 0, 0,
		0, 30, 222, 1, 0, 0, 0, 32, 239, 1, 0, 0, 0, 34, 252, 1, 0, 0, 0, 36, 257,
		1, 0, 0, 0, 38, 269, 1, 0, 0, 0, 40, 271, 1, 0, 0, 0, 42, 275, 1, 0, 0,
		0, 44, 48, 3, 2, 1, 0, 45, 47, 3, 8, 4, 0, 46, 45, 1, 0, 0, 0, 47, 50,
		1, 0, 0, 0, 48, 46, 1, 0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 1, 1, 0, 0, 0,
		50, 48, 1, 0, 0, 0, 51, 55, 3, 4, 2, 0, 52, 54, 3, 6, 3, 0, 53, 52, 1,
		0, 0, 0, 54, 57, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56,
		3, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 58, 60, 3, 42, 21, 0, 59, 58, 1, 0,
		0, 0, 60, 63, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 64,
		1, 0, 0, 0, 63, 61, 1, 0, 0, 0, 64, 65, 5, 1, 0, 0, 65, 67, 5, 31, 0, 0,
		66, 68, 5, 32, 0, 0, 67, 66, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 70, 1,
		0, 0, 0, 69, 71, 5, 40, 0, 0, 70, 69, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71,
		5, 1, 0, 0, 0, 72, 73, 5, 2, 0, 0, 73, 75, 5, 31, 0, 0, 74, 76, 5, 32,
		0, 0, 75, 74, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 78, 1, 0, 0, 0, 77, 79,
		5, 40, 0, 0, 78, 77, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 7, 1, 0, 0, 0,
		80, 85, 3, 10, 5, 0, 81, 85, 3, 12, 6, 0, 82, 85, 3, 26, 13, 0, 83, 85,
		3, 30, 15, 0, 84, 80, 1, 0, 0, 0, 84, 81, 1, 0, 0, 0, 84, 82, 1, 0, 0,
		0, 84, 83, 1, 0, 0, 0, 85, 9, 1, 0, 0, 0, 86, 88, 3, 42, 21, 0, 87, 86,
		1, 0, 0, 0, 88, 91, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0,
		90, 92, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 92, 93, 5, 3, 0, 0, 93, 95, 5,
		31, 0, 0, 94, 96, 5, 40, 0, 0, 95, 94, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0,
		96, 11, 1, 0, 0, 0, 97, 99, 3, 42, 21, 0, 98, 97, 1, 0, 0, 0, 99, 102,
		1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 103, 1, 0,
		0, 0, 102, 100, 1, 0, 0, 0, 103, 104, 5, 4, 0, 0, 104, 107, 5, 31, 0, 0,
		105, 106, 5, 5, 0, 0, 106, 108, 5, 31, 0, 0, 107, 105, 1, 0, 0, 0, 107,
		108, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109, 113, 5, 6, 0, 0, 110, 112,
		3, 14, 7, 0, 111, 110, 1, 0, 0, 0, 112, 115, 1, 0, 0, 0, 113, 111, 1, 0,
		0, 0, 113, 114, 1, 0, 0, 0, 114, 116, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0,
		116, 117, 5, 7, 0, 0, 117, 13, 1, 0, 0, 0, 118, 122, 3, 16, 8, 0, 119,
		122, 3, 18, 9, 0, 120, 122, 3, 24, 12, 0, 121, 118, 1, 0, 0, 0, 121, 119,
		1, 0, 0, 0, 121, 120, 1, 0, 0, 0, 122, 15, 1, 0, 0, 0, 123, 125, 3, 42,
		21, 0, 124, 123, 1, 0, 0, 0, 125, 128, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0,
		126, 127, 1, 0, 0, 0, 127, 130, 1, 0, 0, 0, 128, 126, 1, 0, 0, 0, 129,
		131, 5, 8, 0, 0, 130, 129, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131, 132,
		1, 0, 0, 0, 132, 133, 5, 31, 0, 0, 133, 134, 5, 9, 0, 0, 134, 136, 3, 34,
		17, 0, 135, 137, 5, 40, 0, 0, 136, 135, 1, 0, 0, 0, 136, 137, 1, 0, 0,
		0, 137, 17, 1, 0, 0, 0, 138, 140, 3, 42, 21, 0, 139, 138, 1, 0, 0, 0, 140,
		143, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 144,
		1, 0, 0, 0, 143, 141, 1, 0, 0, 0, 144, 145, 5, 31, 0, 0, 145, 149, 5, 10,
		0, 0, 146, 148, 3, 22, 11, 0, 147, 146, 1, 0, 0, 0, 148, 151, 1, 0, 0,
		0, 149, 147, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 152, 1, 0, 0, 0, 151,
		149, 1, 0, 0, 0, 152, 154, 5, 11, 0, 0, 153, 155, 3, 20, 10, 0, 154, 153,
		1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 157, 1, 0, 0, 0, 156, 158, 5, 40,
		0, 0, 157, 156, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 19, 1, 0, 0, 0,
		159, 160, 5, 9, 0, 0, 160, 161, 3, 34, 17, 0, 161, 21, 1, 0, 0, 0, 162,
		163, 5, 31, 0, 0, 163, 164, 5, 9, 0, 0, 164, 166, 3, 34, 17, 0, 165, 167,
		5, 12, 0, 0, 166, 165, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 23, 1, 0,
		0, 0, 168, 170, 3, 42, 21, 0, 169, 168, 1, 0, 0, 0, 170, 173, 1, 0, 0,
		0, 171, 169, 1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172, 174, 1, 0, 0, 0, 173,
		171, 1, 0, 0, 0, 174, 175, 5, 13, 0, 0, 175, 176, 5, 31, 0, 0, 176, 180,
		5, 10, 0, 0, 177, 179, 3, 22, 11, 0, 178, 177, 1, 0, 0, 0, 179, 182, 1,
		0, 0, 0, 180, 178, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181, 183, 1, 0, 0,
		0, 182, 180, 1, 0, 0, 0, 183, 185, 5, 11, 0, 0, 184, 186, 5, 40, 0, 0,
		185, 184, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0, 186, 25, 1, 0, 0, 0, 187, 189,
		3, 42, 21, 0, 188, 187, 1, 0, 0, 0, 189, 192, 1, 0, 0, 0, 190, 188, 1,
		0, 0, 0, 190, 191, 1, 0, 0, 0, 191, 193, 1, 0, 0, 0, 192, 190, 1, 0, 0,
		0, 193, 194, 5, 14, 0, 0, 194, 195, 5, 31, 0, 0, 195, 199, 5, 6, 0, 0,
		196, 198, 3, 28, 14, 0, 197, 196, 1, 0, 0, 0, 198, 201, 1, 0, 0, 0, 199,
		197, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200, 202, 1, 0, 0, 0, 201, 199,
		1, 0, 0, 0, 202, 203, 5, 7, 0, 0, 203, 27, 1, 0, 0, 0, 204, 206, 3, 42,
		21, 0, 205, 204, 1, 0, 0, 0, 206, 209, 1, 0, 0, 0, 207, 205, 1, 0, 0, 0,
		207, 208, 1, 0, 0, 0, 208, 211, 1, 0, 0, 0, 209, 207, 1, 0, 0, 0, 210,
		212, 5, 8, 0, 0, 211, 210, 1, 0, 0, 0, 211, 212, 1, 0, 0, 0, 212, 213,
		1, 0, 0, 0, 213, 214, 5, 31, 0, 0, 214, 215, 5, 9, 0, 0, 215, 217, 3, 34,
		17, 0, 216, 218, 5, 40, 0, 0, 217, 216, 1, 0, 0, 0, 217, 218, 1, 0, 0,
		0, 218, 29, 1, 0, 0, 0, 219, 221, 3, 42, 21, 0, 220, 219, 1, 0, 0, 0, 221,
		224, 1, 0, 0, 0, 222, 220, 1, 0, 0, 0, 222, 223, 1, 0, 0, 0, 223, 225,
		1, 0, 0, 0, 224, 222, 1, 0, 0, 0, 225, 226, 5, 15, 0, 0, 226, 227, 5, 31,
		0, 0, 227, 231, 5, 6, 0, 0, 228, 230, 3, 32, 16, 0, 229, 228, 1, 0, 0,
		0, 230, 233, 1, 0, 0, 0, 231, 229, 1, 0, 0, 0, 231, 232, 1, 0, 0, 0, 232,
		234, 1, 0, 0, 0, 233, 231, 1, 0, 0, 0, 234, 235, 5, 7, 0, 0, 235, 31, 1,
		0, 0, 0, 236, 238, 3, 42, 21, 0, 237, 236, 1, 0, 0, 0, 238, 241, 1, 0,
		0, 0, 239, 237, 1, 0, 0, 0, 239, 240, 1, 0, 0, 0, 240, 242, 1, 0, 0, 0,
		241, 239, 1, 0, 0, 0, 242, 245, 5, 31, 0, 0, 243, 244, 5, 16, 0, 0, 244,
		246, 5, 29, 0, 0, 245, 243, 1, 0, 0, 0, 245, 246, 1, 0, 0, 0, 246, 248,
		1, 0, 0, 0, 247, 249, 5, 12, 0, 0, 248, 247, 1, 0, 0, 0, 248, 249, 1, 0,
		0, 0, 249, 33, 1, 0, 0, 0, 250, 253, 3, 38, 19, 0, 251, 253, 3, 40, 20,
		0, 252, 250, 1, 0, 0, 0, 252, 251, 1, 0, 0, 0, 253, 255, 1, 0, 0, 0, 254,
		256, 3, 36, 18, 0, 255, 254, 1, 0, 0, 0, 255, 256, 1, 0, 0, 0, 256, 35,
		1, 0, 0, 0, 257, 258, 5, 17, 0, 0, 258, 259, 5, 18, 0, 0, 259, 37, 1, 0,
		0, 0, 260, 270, 5, 19, 0, 0, 261, 270, 5, 20, 0, 0, 262, 270, 5, 21, 0,
		0, 263, 270, 5, 22, 0, 0, 264, 270, 5, 23, 0, 0, 265, 270, 5, 24, 0, 0,
		266, 270, 5, 25, 0, 0, 267, 270, 5, 26, 0, 0, 268, 270, 5, 27, 0, 0, 269,
		260, 1, 0, 0, 0, 269, 261, 1, 0, 0, 0, 269, 262, 1, 0, 0, 0, 269, 263,
		1, 0, 0, 0, 269, 264, 1, 0, 0, 0, 269, 265, 1, 0, 0, 0, 269, 266, 1, 0,
		0, 0, 269, 267, 1, 0, 0, 0, 269, 268, 1, 0, 0, 0, 270, 39, 1, 0, 0, 0,
		271, 272, 5, 31, 0, 0, 272, 41, 1, 0, 0, 0, 273, 276, 5, 34, 0, 0, 274,
		276, 5, 33, 0, 0, 275, 273, 1, 0, 0, 0, 275, 274, 1, 0, 0, 0, 276, 43,
		1, 0, 0, 0, 39, 48, 55, 61, 67, 70, 75, 78, 84, 89, 95, 100, 107, 113,
		121, 126, 130, 136, 141, 149, 154, 157, 166, 171, 180, 185, 190, 199, 207,
		211, 217, 222, 231, 239, 245, 248, 252, 255, 269, 275,
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
	staticData := &ObjectApiParserStaticData
	staticData.once.Do(objectapiParserInit)
}

// NewObjectApiParser produces a new parser instance for the optional input antlr.TokenStream.
func NewObjectApiParser(input antlr.TokenStream) *ObjectApiParser {
	ObjectApiParserInit()
	this := new(ObjectApiParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ObjectApiParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "ObjectApi.g4"

	return this
}

// ObjectApiParser tokens.
const (
	ObjectApiParserEOF        = antlr.TokenEOF
	ObjectApiParserT__0       = 1
	ObjectApiParserT__1       = 2
	ObjectApiParserT__2       = 3
	ObjectApiParserT__3       = 4
	ObjectApiParserT__4       = 5
	ObjectApiParserT__5       = 6
	ObjectApiParserT__6       = 7
	ObjectApiParserT__7       = 8
	ObjectApiParserT__8       = 9
	ObjectApiParserT__9       = 10
	ObjectApiParserT__10      = 11
	ObjectApiParserT__11      = 12
	ObjectApiParserT__12      = 13
	ObjectApiParserT__13      = 14
	ObjectApiParserT__14      = 15
	ObjectApiParserT__15      = 16
	ObjectApiParserT__16      = 17
	ObjectApiParserT__17      = 18
	ObjectApiParserT__18      = 19
	ObjectApiParserT__19      = 20
	ObjectApiParserT__20      = 21
	ObjectApiParserT__21      = 22
	ObjectApiParserT__22      = 23
	ObjectApiParserT__23      = 24
	ObjectApiParserT__24      = 25
	ObjectApiParserT__25      = 26
	ObjectApiParserT__26      = 27
	ObjectApiParserWHITESPACE = 28
	ObjectApiParserINTEGER    = 29
	ObjectApiParserHEX        = 30
	ObjectApiParserIDENTIFIER = 31
	ObjectApiParserVERSION    = 32
	ObjectApiParserDOCLINE    = 33
	ObjectApiParserTAGLINE    = 34
	ObjectApiParserCOMMENT    = 35
	ObjectApiParserDOT        = 36
	ObjectApiParserLETTER     = 37
	ObjectApiParserDIGIT      = 38
	ObjectApiParserUNDERSCORE = 39
	ObjectApiParserSEMICOLON  = 40
)

// ObjectApiParser rules.
const (
	ObjectApiParserRULE_documentRule         = 0
	ObjectApiParserRULE_headerRule           = 1
	ObjectApiParserRULE_moduleRule           = 2
	ObjectApiParserRULE_importRule           = 3
	ObjectApiParserRULE_declarationsRule     = 4
	ObjectApiParserRULE_externRule           = 5
	ObjectApiParserRULE_interfaceRule        = 6
	ObjectApiParserRULE_interfaceMembersRule = 7
	ObjectApiParserRULE_propertyRule         = 8
	ObjectApiParserRULE_operationRule        = 9
	ObjectApiParserRULE_operationReturnRule  = 10
	ObjectApiParserRULE_operationParamRule   = 11
	ObjectApiParserRULE_signalRule           = 12
	ObjectApiParserRULE_structRule           = 13
	ObjectApiParserRULE_structFieldRule      = 14
	ObjectApiParserRULE_enumRule             = 15
	ObjectApiParserRULE_enumMemberRule       = 16
	ObjectApiParserRULE_schemaRule           = 17
	ObjectApiParserRULE_arrayRule            = 18
	ObjectApiParserRULE_primitiveSchema      = 19
	ObjectApiParserRULE_symbolSchema         = 20
	ObjectApiParserRULE_metaRule             = 21
)

// IDocumentRuleContext is an interface to support dynamic dispatch.
type IDocumentRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	HeaderRule() IHeaderRuleContext
	AllDeclarationsRule() []IDeclarationsRuleContext
	DeclarationsRule(i int) IDeclarationsRuleContext

	// IsDocumentRuleContext differentiates from other interfaces.
	IsDocumentRuleContext()
}

type DocumentRuleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDocumentRuleContext() *DocumentRuleContext {
	var p = new(DocumentRuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_documentRule
	return p
}

func InitEmptyDocumentRuleContext(p *DocumentRuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_documentRule
}

func (*DocumentRuleContext) IsDocumentRuleContext() {}

func NewDocumentRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentRuleContext {
	var p = new(DocumentRuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewDocumentRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ObjectApiParserRULE_documentRule)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		p.HeaderRule()
	}
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&25769852952) != 0 {
		{
			p.SetState(45)
			p.DeclarationsRule()
		}

		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHeaderRuleContext is an interface to support dynamic dispatch.
type IHeaderRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ModuleRule() IModuleRuleContext
	AllImportRule() []IImportRuleContext
	ImportRule(i int) IImportRuleContext

	// IsHeaderRuleContext differentiates from other interfaces.
	IsHeaderRuleContext()
}

type HeaderRuleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHeaderRuleContext() *HeaderRuleContext {
	var p = new(HeaderRuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_headerRule
	return p
}

func InitEmptyHeaderRuleContext(p *HeaderRuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_headerRule
}

func (*HeaderRuleContext) IsHeaderRuleContext() {}

func NewHeaderRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HeaderRuleContext {
	var p = new(HeaderRuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewHeaderRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ObjectApiParserRULE_headerRule)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		p.ModuleRule()
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ObjectApiParserT__1 {
		{
			p.SetState(52)
			p.ImportRule()
		}

		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	AllMetaRule() []IMetaRuleContext
	MetaRule(i int) IMetaRuleContext
	SEMICOLON() antlr.TerminalNode
	VERSION() antlr.TerminalNode

	// IsModuleRuleContext differentiates from other interfaces.
	IsModuleRuleContext()
}

type ModuleRuleContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	name    antlr.Token
	version antlr.Token
}

func NewEmptyModuleRuleContext() *ModuleRuleContext {
	var p = new(ModuleRuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_moduleRule
	return p
}

func InitEmptyModuleRuleContext(p *ModuleRuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_moduleRule
}

func (*ModuleRuleContext) IsModuleRuleContext() {}

func NewModuleRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleRuleContext {
	var p = new(ModuleRuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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

func (s *ModuleRuleContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ObjectApiParserSEMICOLON, 0)
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
	localctx = NewModuleRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ObjectApiParserRULE_moduleRule)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ObjectApiParserDOCLINE || _la == ObjectApiParserTAGLINE {
		{
			p.SetState(58)
			p.MetaRule()
		}

		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(64)
		p.Match(ObjectApiParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(65)

		var _m = p.Match(ObjectApiParserIDENTIFIER)

		localctx.(*ModuleRuleContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ObjectApiParserVERSION {
		{
			p.SetState(66)

			var _m = p.Match(ObjectApiParserVERSION)

			localctx.(*ModuleRuleContext).version = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ObjectApiParserSEMICOLON {
		{
			p.SetState(69)
			p.Match(ObjectApiParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode
	VERSION() antlr.TerminalNode

	// IsImportRuleContext differentiates from other interfaces.
	IsImportRuleContext()
}

type ImportRuleContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	name    antlr.Token
	version antlr.Token
}

func NewEmptyImportRuleContext() *ImportRuleContext {
	var p = new(ImportRuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_importRule
	return p
}

func InitEmptyImportRuleContext(p *ImportRuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_importRule
}

func (*ImportRuleContext) IsImportRuleContext() {}

func NewImportRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportRuleContext {
	var p = new(ImportRuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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

func (s *ImportRuleContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ObjectApiParserSEMICOLON, 0)
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
	localctx = NewImportRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ObjectApiParserRULE_importRule)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.Match(ObjectApiParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(73)

		var _m = p.Match(ObjectApiParserIDENTIFIER)

		localctx.(*ImportRuleContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ObjectApiParserVERSION {
		{
			p.SetState(74)

			var _m = p.Match(ObjectApiParserVERSION)

			localctx.(*ImportRuleContext).version = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ObjectApiParserSEMICOLON {
		{
			p.SetState(77)
			p.Match(ObjectApiParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclarationsRuleContext is an interface to support dynamic dispatch.
type IDeclarationsRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExternRule() IExternRuleContext
	InterfaceRule() IInterfaceRuleContext
	StructRule() IStructRuleContext
	EnumRule() IEnumRuleContext

	// IsDeclarationsRuleContext differentiates from other interfaces.
	IsDeclarationsRuleContext()
}

type DeclarationsRuleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationsRuleContext() *DeclarationsRuleContext {
	var p = new(DeclarationsRuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_declarationsRule
	return p
}

func InitEmptyDeclarationsRuleContext(p *DeclarationsRuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_declarationsRule
}

func (*DeclarationsRuleContext) IsDeclarationsRuleContext() {}

func NewDeclarationsRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationsRuleContext {
	var p = new(DeclarationsRuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectApiParserRULE_declarationsRule

	return p
}

func (s *DeclarationsRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationsRuleContext) ExternRule() IExternRuleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExternRuleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExternRuleContext)
}

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
	localctx = NewDeclarationsRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ObjectApiParserRULE_declarationsRule)
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(80)
			p.ExternRule()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(81)
			p.InterfaceRule()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(82)
			p.StructRule()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(83)
			p.EnumRule()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExternRuleContext is an interface to support dynamic dispatch.
type IExternRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	AllMetaRule() []IMetaRuleContext
	MetaRule(i int) IMetaRuleContext
	SEMICOLON() antlr.TerminalNode

	// IsExternRuleContext differentiates from other interfaces.
	IsExternRuleContext()
}

type ExternRuleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyExternRuleContext() *ExternRuleContext {
	var p = new(ExternRuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_externRule
	return p
}

func InitEmptyExternRuleContext(p *ExternRuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_externRule
}

func (*ExternRuleContext) IsExternRuleContext() {}

func NewExternRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternRuleContext {
	var p = new(ExternRuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectApiParserRULE_externRule

	return p
}

func (s *ExternRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *ExternRuleContext) GetName() antlr.Token { return s.name }

func (s *ExternRuleContext) SetName(v antlr.Token) { s.name = v }

func (s *ExternRuleContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ObjectApiParserIDENTIFIER, 0)
}

func (s *ExternRuleContext) AllMetaRule() []IMetaRuleContext {
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

func (s *ExternRuleContext) MetaRule(i int) IMetaRuleContext {
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

func (s *ExternRuleContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ObjectApiParserSEMICOLON, 0)
}

func (s *ExternRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExternRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExternRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.EnterExternRule(s)
	}
}

func (s *ExternRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectApiListener); ok {
		listenerT.ExitExternRule(s)
	}
}

func (p *ObjectApiParser) ExternRule() (localctx IExternRuleContext) {
	localctx = NewExternRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ObjectApiParserRULE_externRule)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ObjectApiParserDOCLINE || _la == ObjectApiParserTAGLINE {
		{
			p.SetState(86)
			p.MetaRule()
		}

		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(92)
		p.Match(ObjectApiParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(93)

		var _m = p.Match(ObjectApiParserIDENTIFIER)

		localctx.(*ExternRuleContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ObjectApiParserSEMICOLON {
		{
			p.SetState(94)
			p.Match(ObjectApiParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInterfaceRuleContext is an interface to support dynamic dispatch.
type IInterfaceRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// GetExtends returns the extends token.
	GetExtends() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// SetExtends sets the extends token.
	SetExtends(antlr.Token)

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	AllMetaRule() []IMetaRuleContext
	MetaRule(i int) IMetaRuleContext
	AllInterfaceMembersRule() []IInterfaceMembersRuleContext
	InterfaceMembersRule(i int) IInterfaceMembersRuleContext

	// IsInterfaceRuleContext differentiates from other interfaces.
	IsInterfaceRuleContext()
}

type InterfaceRuleContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	name    antlr.Token
	extends antlr.Token
}

func NewEmptyInterfaceRuleContext() *InterfaceRuleContext {
	var p = new(InterfaceRuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_interfaceRule
	return p
}

func InitEmptyInterfaceRuleContext(p *InterfaceRuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_interfaceRule
}

func (*InterfaceRuleContext) IsInterfaceRuleContext() {}

func NewInterfaceRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InterfaceRuleContext {
	var p = new(InterfaceRuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectApiParserRULE_interfaceRule

	return p
}

func (s *InterfaceRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *InterfaceRuleContext) GetName() antlr.Token { return s.name }

func (s *InterfaceRuleContext) GetExtends() antlr.Token { return s.extends }

func (s *InterfaceRuleContext) SetName(v antlr.Token) { s.name = v }

func (s *InterfaceRuleContext) SetExtends(v antlr.Token) { s.extends = v }

func (s *InterfaceRuleContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(ObjectApiParserIDENTIFIER)
}

func (s *InterfaceRuleContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(ObjectApiParserIDENTIFIER, i)
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
	localctx = NewInterfaceRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ObjectApiParserRULE_interfaceRule)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ObjectApiParserDOCLINE || _la == ObjectApiParserTAGLINE {
		{
			p.SetState(97)
			p.MetaRule()
		}

		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(103)
		p.Match(ObjectApiParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(104)

		var _m = p.Match(ObjectApiParserIDENTIFIER)

		localctx.(*InterfaceRuleContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ObjectApiParserT__4 {
		{
			p.SetState(105)
			p.Match(ObjectApiParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)

			var _m = p.Match(ObjectApiParserIDENTIFIER)

			localctx.(*InterfaceRuleContext).extends = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(109)
		p.Match(ObjectApiParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&27917295872) != 0 {
		{
			p.SetState(110)
			p.InterfaceMembersRule()
		}

		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(116)
		p.Match(ObjectApiParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInterfaceMembersRuleContext is an interface to support dynamic dispatch.
type IInterfaceMembersRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PropertyRule() IPropertyRuleContext
	OperationRule() IOperationRuleContext
	SignalRule() ISignalRuleContext

	// IsInterfaceMembersRuleContext differentiates from other interfaces.
	IsInterfaceMembersRuleContext()
}

type InterfaceMembersRuleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInterfaceMembersRuleContext() *InterfaceMembersRuleContext {
	var p = new(InterfaceMembersRuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_interfaceMembersRule
	return p
}

func InitEmptyInterfaceMembersRuleContext(p *InterfaceMembersRuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_interfaceMembersRule
}

func (*InterfaceMembersRuleContext) IsInterfaceMembersRuleContext() {}

func NewInterfaceMembersRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InterfaceMembersRuleContext {
	var p = new(InterfaceMembersRuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewInterfaceMembersRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ObjectApiParserRULE_interfaceMembersRule)
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(118)
			p.PropertyRule()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(119)
			p.OperationRule()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(120)
			p.SignalRule()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPropertyRuleContext is an interface to support dynamic dispatch.
type IPropertyRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetReadonly returns the readonly token.
	GetReadonly() antlr.Token

	// GetName returns the name token.
	GetName() antlr.Token

	// SetReadonly sets the readonly token.
	SetReadonly(antlr.Token)

	// SetName sets the name token.
	SetName(antlr.Token)

	// GetSchema returns the schema rule contexts.
	GetSchema() ISchemaRuleContext

	// SetSchema sets the schema rule contexts.
	SetSchema(ISchemaRuleContext)

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	SchemaRule() ISchemaRuleContext
	AllMetaRule() []IMetaRuleContext
	MetaRule(i int) IMetaRuleContext
	SEMICOLON() antlr.TerminalNode

	// IsPropertyRuleContext differentiates from other interfaces.
	IsPropertyRuleContext()
}

type PropertyRuleContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	readonly antlr.Token
	name     antlr.Token
	schema   ISchemaRuleContext
}

func NewEmptyPropertyRuleContext() *PropertyRuleContext {
	var p = new(PropertyRuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_propertyRule
	return p
}

func InitEmptyPropertyRuleContext(p *PropertyRuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_propertyRule
}

func (*PropertyRuleContext) IsPropertyRuleContext() {}

func NewPropertyRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyRuleContext {
	var p = new(PropertyRuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectApiParserRULE_propertyRule

	return p
}

func (s *PropertyRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyRuleContext) GetReadonly() antlr.Token { return s.readonly }

func (s *PropertyRuleContext) GetName() antlr.Token { return s.name }

func (s *PropertyRuleContext) SetReadonly(v antlr.Token) { s.readonly = v }

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

func (s *PropertyRuleContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ObjectApiParserSEMICOLON, 0)
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
	localctx = NewPropertyRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ObjectApiParserRULE_propertyRule)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ObjectApiParserDOCLINE || _la == ObjectApiParserTAGLINE {
		{
			p.SetState(123)
			p.MetaRule()
		}

		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ObjectApiParserT__7 {
		{
			p.SetState(129)

			var _m = p.Match(ObjectApiParserT__7)

			localctx.(*PropertyRuleContext).readonly = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(132)

		var _m = p.Match(ObjectApiParserIDENTIFIER)

		localctx.(*PropertyRuleContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(133)
		p.Match(ObjectApiParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(134)

		var _x = p.SchemaRule()

		localctx.(*PropertyRuleContext).schema = _x
	}
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ObjectApiParserSEMICOLON {
		{
			p.SetState(135)
			p.Match(ObjectApiParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	AllMetaRule() []IMetaRuleContext
	MetaRule(i int) IMetaRuleContext
	OperationReturnRule() IOperationReturnRuleContext
	SEMICOLON() antlr.TerminalNode
	AllOperationParamRule() []IOperationParamRuleContext
	OperationParamRule(i int) IOperationParamRuleContext

	// IsOperationRuleContext differentiates from other interfaces.
	IsOperationRuleContext()
}

type OperationRuleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	params IOperationParamRuleContext
}

func NewEmptyOperationRuleContext() *OperationRuleContext {
	var p = new(OperationRuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_operationRule
	return p
}

func InitEmptyOperationRuleContext(p *OperationRuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_operationRule
}

func (*OperationRuleContext) IsOperationRuleContext() {}

func NewOperationRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationRuleContext {
	var p = new(OperationRuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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

func (s *OperationRuleContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ObjectApiParserSEMICOLON, 0)
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
	localctx = NewOperationRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ObjectApiParserRULE_operationRule)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ObjectApiParserDOCLINE || _la == ObjectApiParserTAGLINE {
		{
			p.SetState(138)
			p.MetaRule()
		}

		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(144)

		var _m = p.Match(ObjectApiParserIDENTIFIER)

		localctx.(*OperationRuleContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(145)
		p.Match(ObjectApiParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ObjectApiParserIDENTIFIER {
		{
			p.SetState(146)

			var _x = p.OperationParamRule()

			localctx.(*OperationRuleContext).params = _x
		}

		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(152)
		p.Match(ObjectApiParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ObjectApiParserT__8 {
		{
			p.SetState(153)
			p.OperationReturnRule()
		}

	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ObjectApiParserSEMICOLON {
		{
			p.SetState(156)
			p.Match(ObjectApiParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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

	// Getter signatures
	SchemaRule() ISchemaRuleContext

	// IsOperationReturnRuleContext differentiates from other interfaces.
	IsOperationReturnRuleContext()
}

type OperationReturnRuleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	schema ISchemaRuleContext
}

func NewEmptyOperationReturnRuleContext() *OperationReturnRuleContext {
	var p = new(OperationReturnRuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_operationReturnRule
	return p
}

func InitEmptyOperationReturnRuleContext(p *OperationReturnRuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_operationReturnRule
}

func (*OperationReturnRuleContext) IsOperationReturnRuleContext() {}

func NewOperationReturnRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationReturnRuleContext {
	var p = new(OperationReturnRuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewOperationReturnRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ObjectApiParserRULE_operationReturnRule)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(159)
		p.Match(ObjectApiParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(160)

		var _x = p.SchemaRule()

		localctx.(*OperationReturnRuleContext).schema = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	SchemaRule() ISchemaRuleContext

	// IsOperationParamRuleContext differentiates from other interfaces.
	IsOperationParamRuleContext()
}

type OperationParamRuleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	schema ISchemaRuleContext
}

func NewEmptyOperationParamRuleContext() *OperationParamRuleContext {
	var p = new(OperationParamRuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_operationParamRule
	return p
}

func InitEmptyOperationParamRuleContext(p *OperationParamRuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_operationParamRule
}

func (*OperationParamRuleContext) IsOperationParamRuleContext() {}

func NewOperationParamRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationParamRuleContext {
	var p = new(OperationParamRuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewOperationParamRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ObjectApiParserRULE_operationParamRule)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(162)

		var _m = p.Match(ObjectApiParserIDENTIFIER)

		localctx.(*OperationParamRuleContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(163)
		p.Match(ObjectApiParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(164)

		var _x = p.SchemaRule()

		localctx.(*OperationParamRuleContext).schema = _x
	}
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ObjectApiParserT__11 {
		{
			p.SetState(165)
			p.Match(ObjectApiParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	AllMetaRule() []IMetaRuleContext
	MetaRule(i int) IMetaRuleContext
	SEMICOLON() antlr.TerminalNode
	AllOperationParamRule() []IOperationParamRuleContext
	OperationParamRule(i int) IOperationParamRuleContext

	// IsSignalRuleContext differentiates from other interfaces.
	IsSignalRuleContext()
}

type SignalRuleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	params IOperationParamRuleContext
}

func NewEmptySignalRuleContext() *SignalRuleContext {
	var p = new(SignalRuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_signalRule
	return p
}

func InitEmptySignalRuleContext(p *SignalRuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_signalRule
}

func (*SignalRuleContext) IsSignalRuleContext() {}

func NewSignalRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SignalRuleContext {
	var p = new(SignalRuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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

func (s *SignalRuleContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ObjectApiParserSEMICOLON, 0)
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
	localctx = NewSignalRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ObjectApiParserRULE_signalRule)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ObjectApiParserDOCLINE || _la == ObjectApiParserTAGLINE {
		{
			p.SetState(168)
			p.MetaRule()
		}

		p.SetState(173)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(174)
		p.Match(ObjectApiParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(175)

		var _m = p.Match(ObjectApiParserIDENTIFIER)

		localctx.(*SignalRuleContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(176)
		p.Match(ObjectApiParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ObjectApiParserIDENTIFIER {
		{
			p.SetState(177)

			var _x = p.OperationParamRule()

			localctx.(*SignalRuleContext).params = _x
		}

		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(183)
		p.Match(ObjectApiParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ObjectApiParserSEMICOLON {
		{
			p.SetState(184)
			p.Match(ObjectApiParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	AllMetaRule() []IMetaRuleContext
	MetaRule(i int) IMetaRuleContext
	AllStructFieldRule() []IStructFieldRuleContext
	StructFieldRule(i int) IStructFieldRuleContext

	// IsStructRuleContext differentiates from other interfaces.
	IsStructRuleContext()
}

type StructRuleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyStructRuleContext() *StructRuleContext {
	var p = new(StructRuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_structRule
	return p
}

func InitEmptyStructRuleContext(p *StructRuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_structRule
}

func (*StructRuleContext) IsStructRuleContext() {}

func NewStructRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructRuleContext {
	var p = new(StructRuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewStructRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ObjectApiParserRULE_structRule)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ObjectApiParserDOCLINE || _la == ObjectApiParserTAGLINE {
		{
			p.SetState(187)
			p.MetaRule()
		}

		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(193)
		p.Match(ObjectApiParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(194)

		var _m = p.Match(ObjectApiParserIDENTIFIER)

		localctx.(*StructRuleContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(195)
		p.Match(ObjectApiParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&27917287680) != 0 {
		{
			p.SetState(196)
			p.StructFieldRule()
		}

		p.SetState(201)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(202)
		p.Match(ObjectApiParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructFieldRuleContext is an interface to support dynamic dispatch.
type IStructFieldRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetReadonly returns the readonly token.
	GetReadonly() antlr.Token

	// GetName returns the name token.
	GetName() antlr.Token

	// SetReadonly sets the readonly token.
	SetReadonly(antlr.Token)

	// SetName sets the name token.
	SetName(antlr.Token)

	// GetSchema returns the schema rule contexts.
	GetSchema() ISchemaRuleContext

	// SetSchema sets the schema rule contexts.
	SetSchema(ISchemaRuleContext)

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	SchemaRule() ISchemaRuleContext
	AllMetaRule() []IMetaRuleContext
	MetaRule(i int) IMetaRuleContext
	SEMICOLON() antlr.TerminalNode

	// IsStructFieldRuleContext differentiates from other interfaces.
	IsStructFieldRuleContext()
}

type StructFieldRuleContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	readonly antlr.Token
	name     antlr.Token
	schema   ISchemaRuleContext
}

func NewEmptyStructFieldRuleContext() *StructFieldRuleContext {
	var p = new(StructFieldRuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_structFieldRule
	return p
}

func InitEmptyStructFieldRuleContext(p *StructFieldRuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_structFieldRule
}

func (*StructFieldRuleContext) IsStructFieldRuleContext() {}

func NewStructFieldRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructFieldRuleContext {
	var p = new(StructFieldRuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectApiParserRULE_structFieldRule

	return p
}

func (s *StructFieldRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *StructFieldRuleContext) GetReadonly() antlr.Token { return s.readonly }

func (s *StructFieldRuleContext) GetName() antlr.Token { return s.name }

func (s *StructFieldRuleContext) SetReadonly(v antlr.Token) { s.readonly = v }

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

func (s *StructFieldRuleContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ObjectApiParserSEMICOLON, 0)
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
	localctx = NewStructFieldRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ObjectApiParserRULE_structFieldRule)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ObjectApiParserDOCLINE || _la == ObjectApiParserTAGLINE {
		{
			p.SetState(204)
			p.MetaRule()
		}

		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ObjectApiParserT__7 {
		{
			p.SetState(210)

			var _m = p.Match(ObjectApiParserT__7)

			localctx.(*StructFieldRuleContext).readonly = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(213)

		var _m = p.Match(ObjectApiParserIDENTIFIER)

		localctx.(*StructFieldRuleContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(214)
		p.Match(ObjectApiParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(215)

		var _x = p.SchemaRule()

		localctx.(*StructFieldRuleContext).schema = _x
	}
	p.SetState(217)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ObjectApiParserSEMICOLON {
		{
			p.SetState(216)
			p.Match(ObjectApiParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	AllMetaRule() []IMetaRuleContext
	MetaRule(i int) IMetaRuleContext
	AllEnumMemberRule() []IEnumMemberRuleContext
	EnumMemberRule(i int) IEnumMemberRuleContext

	// IsEnumRuleContext differentiates from other interfaces.
	IsEnumRuleContext()
}

type EnumRuleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyEnumRuleContext() *EnumRuleContext {
	var p = new(EnumRuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_enumRule
	return p
}

func InitEmptyEnumRuleContext(p *EnumRuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_enumRule
}

func (*EnumRuleContext) IsEnumRuleContext() {}

func NewEnumRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumRuleContext {
	var p = new(EnumRuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewEnumRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ObjectApiParserRULE_enumRule)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ObjectApiParserDOCLINE || _la == ObjectApiParserTAGLINE {
		{
			p.SetState(219)
			p.MetaRule()
		}

		p.SetState(224)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(225)
		p.Match(ObjectApiParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(226)

		var _m = p.Match(ObjectApiParserIDENTIFIER)

		localctx.(*EnumRuleContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(227)
		p.Match(ObjectApiParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&27917287424) != 0 {
		{
			p.SetState(228)
			p.EnumMemberRule()
		}

		p.SetState(233)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(234)
		p.Match(ObjectApiParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	AllMetaRule() []IMetaRuleContext
	MetaRule(i int) IMetaRuleContext
	INTEGER() antlr.TerminalNode

	// IsEnumMemberRuleContext differentiates from other interfaces.
	IsEnumMemberRuleContext()
}

type EnumMemberRuleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	value  antlr.Token
}

func NewEmptyEnumMemberRuleContext() *EnumMemberRuleContext {
	var p = new(EnumMemberRuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_enumMemberRule
	return p
}

func InitEmptyEnumMemberRuleContext(p *EnumMemberRuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_enumMemberRule
}

func (*EnumMemberRuleContext) IsEnumMemberRuleContext() {}

func NewEnumMemberRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumMemberRuleContext {
	var p = new(EnumMemberRuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewEnumMemberRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ObjectApiParserRULE_enumMemberRule)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ObjectApiParserDOCLINE || _la == ObjectApiParserTAGLINE {
		{
			p.SetState(236)
			p.MetaRule()
		}

		p.SetState(241)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(242)

		var _m = p.Match(ObjectApiParserIDENTIFIER)

		localctx.(*EnumMemberRuleContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ObjectApiParserT__15 {
		{
			p.SetState(243)
			p.Match(ObjectApiParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(244)

			var _m = p.Match(ObjectApiParserINTEGER)

			localctx.(*EnumMemberRuleContext).value = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ObjectApiParserT__11 {
		{
			p.SetState(247)
			p.Match(ObjectApiParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISchemaRuleContext is an interface to support dynamic dispatch.
type ISchemaRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PrimitiveSchema() IPrimitiveSchemaContext
	SymbolSchema() ISymbolSchemaContext
	ArrayRule() IArrayRuleContext

	// IsSchemaRuleContext differentiates from other interfaces.
	IsSchemaRuleContext()
}

type SchemaRuleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySchemaRuleContext() *SchemaRuleContext {
	var p = new(SchemaRuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_schemaRule
	return p
}

func InitEmptySchemaRuleContext(p *SchemaRuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_schemaRule
}

func (*SchemaRuleContext) IsSchemaRuleContext() {}

func NewSchemaRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SchemaRuleContext {
	var p = new(SchemaRuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewSchemaRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ObjectApiParserRULE_schemaRule)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ObjectApiParserT__18, ObjectApiParserT__19, ObjectApiParserT__20, ObjectApiParserT__21, ObjectApiParserT__22, ObjectApiParserT__23, ObjectApiParserT__24, ObjectApiParserT__25, ObjectApiParserT__26:
		{
			p.SetState(250)
			p.PrimitiveSchema()
		}

	case ObjectApiParserIDENTIFIER:
		{
			p.SetState(251)
			p.SymbolSchema()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ObjectApiParserT__16 {
		{
			p.SetState(254)
			p.ArrayRule()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayRuleContext() *ArrayRuleContext {
	var p = new(ArrayRuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_arrayRule
	return p
}

func InitEmptyArrayRuleContext(p *ArrayRuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_arrayRule
}

func (*ArrayRuleContext) IsArrayRuleContext() {}

func NewArrayRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayRuleContext {
	var p = new(ArrayRuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewArrayRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ObjectApiParserRULE_arrayRule)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(257)
		p.Match(ObjectApiParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(258)
		p.Match(ObjectApiParserT__17)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyPrimitiveSchemaContext() *PrimitiveSchemaContext {
	var p = new(PrimitiveSchemaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_primitiveSchema
	return p
}

func InitEmptyPrimitiveSchemaContext(p *PrimitiveSchemaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_primitiveSchema
}

func (*PrimitiveSchemaContext) IsPrimitiveSchemaContext() {}

func NewPrimitiveSchemaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitiveSchemaContext {
	var p = new(PrimitiveSchemaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewPrimitiveSchemaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ObjectApiParserRULE_primitiveSchema)
	p.SetState(269)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ObjectApiParserT__18:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(260)

			var _m = p.Match(ObjectApiParserT__18)

			localctx.(*PrimitiveSchemaContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ObjectApiParserT__19:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(261)

			var _m = p.Match(ObjectApiParserT__19)

			localctx.(*PrimitiveSchemaContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ObjectApiParserT__20:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(262)

			var _m = p.Match(ObjectApiParserT__20)

			localctx.(*PrimitiveSchemaContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ObjectApiParserT__21:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(263)

			var _m = p.Match(ObjectApiParserT__21)

			localctx.(*PrimitiveSchemaContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ObjectApiParserT__22:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(264)

			var _m = p.Match(ObjectApiParserT__22)

			localctx.(*PrimitiveSchemaContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ObjectApiParserT__23:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(265)

			var _m = p.Match(ObjectApiParserT__23)

			localctx.(*PrimitiveSchemaContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ObjectApiParserT__24:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(266)

			var _m = p.Match(ObjectApiParserT__24)

			localctx.(*PrimitiveSchemaContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ObjectApiParserT__25:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(267)

			var _m = p.Match(ObjectApiParserT__25)

			localctx.(*PrimitiveSchemaContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ObjectApiParserT__26:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(268)

			var _m = p.Match(ObjectApiParserT__26)

			localctx.(*PrimitiveSchemaContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsSymbolSchemaContext differentiates from other interfaces.
	IsSymbolSchemaContext()
}

type SymbolSchemaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptySymbolSchemaContext() *SymbolSchemaContext {
	var p = new(SymbolSchemaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_symbolSchema
	return p
}

func InitEmptySymbolSchemaContext(p *SymbolSchemaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_symbolSchema
}

func (*SymbolSchemaContext) IsSymbolSchemaContext() {}

func NewSymbolSchemaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SymbolSchemaContext {
	var p = new(SymbolSchemaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewSymbolSchemaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ObjectApiParserRULE_symbolSchema)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(271)

		var _m = p.Match(ObjectApiParserIDENTIFIER)

		localctx.(*SymbolSchemaContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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

	// Getter signatures
	TAGLINE() antlr.TerminalNode
	DOCLINE() antlr.TerminalNode

	// IsMetaRuleContext differentiates from other interfaces.
	IsMetaRuleContext()
}

type MetaRuleContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	tagLine antlr.Token
	docLine antlr.Token
}

func NewEmptyMetaRuleContext() *MetaRuleContext {
	var p = new(MetaRuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_metaRule
	return p
}

func InitEmptyMetaRuleContext(p *MetaRuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectApiParserRULE_metaRule
}

func (*MetaRuleContext) IsMetaRuleContext() {}

func NewMetaRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MetaRuleContext {
	var p = new(MetaRuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewMetaRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ObjectApiParserRULE_metaRule)
	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ObjectApiParserTAGLINE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(273)

			var _m = p.Match(ObjectApiParserTAGLINE)

			localctx.(*MetaRuleContext).tagLine = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ObjectApiParserDOCLINE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(274)

			var _m = p.Match(ObjectApiParserDOCLINE)

			localctx.(*MetaRuleContext).docLine = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
