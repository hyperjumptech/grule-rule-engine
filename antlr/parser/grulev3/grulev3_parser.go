// Code generated from C:/Users/Ferdinand/WSL/workspace/golang/src/github.com/newm4n/grule-rule-engine/antlr\grulev3.g4 by ANTLR 4.10.1. DO NOT EDIT.

package grulev3 // grulev3
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

type grulev3Parser struct {
	*antlr.BaseParser
}

var grulev3ParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func grulev3ParserInit() {
	staticData := &grulev3ParserStaticData
	staticData.literalNames = []string{
		"", "','", "'+'", "'-'", "'/'", "'*'", "'%'", "'.'", "';'", "'{'", "'}'",
		"'('", "')'", "'['", "']'", "", "", "", "'&&'", "'||'", "", "", "",
		"'!'", "", "'=='", "'='", "'+='", "'-='", "'/='", "'*='", "'>'", "'<'",
		"'>='", "'<='", "'!='", "'&'", "'|'",
	}
	staticData.symbolicNames = []string{
		"", "", "PLUS", "MINUS", "DIV", "MUL", "MOD", "DOT", "SEMICOLON", "LR_BRACE",
		"RR_BRACE", "LR_BRACKET", "RR_BRACKET", "LS_BRACKET", "RS_BRACKET",
		"RULE", "WHEN", "THEN", "AND", "OR", "TRUE", "FALSE", "NIL_LITERAL",
		"NEGATION", "SALIENCE", "EQUALS", "ASSIGN", "PLUS_ASIGN", "MINUS_ASIGN",
		"DIV_ASIGN", "MUL_ASIGN", "GT", "LT", "GTE", "LTE", "NOTEQUALS", "BITAND",
		"BITOR", "SIMPLENAME", "DQUOTA_STRING", "SQUOTA_STRING", "DECIMAL_FLOAT_LIT",
		"DECIMAL_EXPONENT", "HEX_FLOAT_LIT", "HEX_EXPONENT", "DEC_LIT", "HEX_LIT",
		"OCT_LIT", "SPACE", "COMMENT", "LINE_COMMENT",
	}
	staticData.ruleNames = []string{
		"grl", "ruleEntry", "salience", "ruleName", "ruleDescription", "whenScope",
		"thenScope", "thenExpressionList", "thenExpression", "assignment", "expression",
		"mulDivOperators", "addMinusOperators", "comparisonOperator", "andLogicOperator",
		"orLogicOperator", "expressionAtom", "constant", "variable", "arrayMapSelector",
		"memberVariable", "functionCall", "methodCall", "argumentList", "floatLiteral",
		"decimalFloatLiteral", "hexadecimalFloatLiteral", "integerLiteral",
		"decimalLiteral", "hexadecimalLiteral", "octalLiteral", "stringLiteral",
		"booleanLiteral",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 50, 263, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 1, 0, 5, 0, 68, 8, 0, 10, 0, 12, 0, 71, 9, 0, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 1, 3, 1, 78, 8, 1, 1, 1, 3, 1, 81, 8, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 4, 7, 104, 8, 7, 11, 7, 12, 7,
		105, 1, 8, 1, 8, 3, 8, 110, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10,
		3, 10, 118, 8, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 125, 8, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10,
		147, 8, 10, 10, 10, 12, 10, 150, 9, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1,
		13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 3, 16, 168, 8, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 5,
		16, 176, 8, 16, 10, 16, 12, 16, 179, 9, 16, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 3, 17, 186, 8, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 5, 18, 195, 8, 18, 10, 18, 12, 18, 198, 9, 18, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 3, 21, 210, 8, 21, 1,
		21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 5, 23, 220, 8, 23,
		10, 23, 12, 23, 223, 9, 23, 1, 24, 1, 24, 3, 24, 227, 8, 24, 1, 25, 3,
		25, 230, 8, 25, 1, 25, 1, 25, 1, 26, 3, 26, 235, 8, 26, 1, 26, 1, 26, 1,
		27, 1, 27, 1, 27, 3, 27, 242, 8, 27, 1, 28, 3, 28, 245, 8, 28, 1, 28, 1,
		28, 1, 29, 3, 29, 250, 8, 29, 1, 29, 1, 29, 1, 30, 3, 30, 255, 8, 30, 1,
		30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 0, 3, 20, 32, 36, 33, 0,
		2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38,
		40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 0, 6, 1, 0, 39, 40,
		1, 0, 26, 30, 1, 0, 4, 6, 2, 0, 2, 3, 36, 37, 2, 0, 25, 25, 31, 35, 1,
		0, 20, 21, 263, 0, 69, 1, 0, 0, 0, 2, 74, 1, 0, 0, 0, 4, 87, 1, 0, 0, 0,
		6, 90, 1, 0, 0, 0, 8, 92, 1, 0, 0, 0, 10, 94, 1, 0, 0, 0, 12, 97, 1, 0,
		0, 0, 14, 103, 1, 0, 0, 0, 16, 109, 1, 0, 0, 0, 18, 111, 1, 0, 0, 0, 20,
		124, 1, 0, 0, 0, 22, 151, 1, 0, 0, 0, 24, 153, 1, 0, 0, 0, 26, 155, 1,
		0, 0, 0, 28, 157, 1, 0, 0, 0, 30, 159, 1, 0, 0, 0, 32, 167, 1, 0, 0, 0,
		34, 185, 1, 0, 0, 0, 36, 187, 1, 0, 0, 0, 38, 199, 1, 0, 0, 0, 40, 203,
		1, 0, 0, 0, 42, 206, 1, 0, 0, 0, 44, 213, 1, 0, 0, 0, 46, 216, 1, 0, 0,
		0, 48, 226, 1, 0, 0, 0, 50, 229, 1, 0, 0, 0, 52, 234, 1, 0, 0, 0, 54, 241,
		1, 0, 0, 0, 56, 244, 1, 0, 0, 0, 58, 249, 1, 0, 0, 0, 60, 254, 1, 0, 0,
		0, 62, 258, 1, 0, 0, 0, 64, 260, 1, 0, 0, 0, 66, 68, 3, 2, 1, 0, 67, 66,
		1, 0, 0, 0, 68, 71, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0,
		70, 72, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 72, 73, 5, 0, 0, 1, 73, 1, 1, 0,
		0, 0, 74, 75, 5, 15, 0, 0, 75, 77, 3, 6, 3, 0, 76, 78, 3, 8, 4, 0, 77,
		76, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 80, 1, 0, 0, 0, 79, 81, 3, 4, 2,
		0, 80, 79, 1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 81, 82, 1, 0, 0, 0, 82, 83,
		5, 9, 0, 0, 83, 84, 3, 10, 5, 0, 84, 85, 3, 12, 6, 0, 85, 86, 5, 10, 0,
		0, 86, 3, 1, 0, 0, 0, 87, 88, 5, 24, 0, 0, 88, 89, 3, 54, 27, 0, 89, 5,
		1, 0, 0, 0, 90, 91, 5, 38, 0, 0, 91, 7, 1, 0, 0, 0, 92, 93, 7, 0, 0, 0,
		93, 9, 1, 0, 0, 0, 94, 95, 5, 16, 0, 0, 95, 96, 3, 20, 10, 0, 96, 11, 1,
		0, 0, 0, 97, 98, 5, 17, 0, 0, 98, 99, 3, 14, 7, 0, 99, 13, 1, 0, 0, 0,
		100, 101, 3, 16, 8, 0, 101, 102, 5, 8, 0, 0, 102, 104, 1, 0, 0, 0, 103,
		100, 1, 0, 0, 0, 104, 105, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 105, 106,
		1, 0, 0, 0, 106, 15, 1, 0, 0, 0, 107, 110, 3, 18, 9, 0, 108, 110, 3, 32,
		16, 0, 109, 107, 1, 0, 0, 0, 109, 108, 1, 0, 0, 0, 110, 17, 1, 0, 0, 0,
		111, 112, 3, 36, 18, 0, 112, 113, 7, 1, 0, 0, 113, 114, 3, 20, 10, 0, 114,
		19, 1, 0, 0, 0, 115, 117, 6, 10, -1, 0, 116, 118, 5, 23, 0, 0, 117, 116,
		1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119, 120, 5, 11,
		0, 0, 120, 121, 3, 20, 10, 0, 121, 122, 5, 12, 0, 0, 122, 125, 1, 0, 0,
		0, 123, 125, 3, 32, 16, 0, 124, 115, 1, 0, 0, 0, 124, 123, 1, 0, 0, 0,
		125, 148, 1, 0, 0, 0, 126, 127, 10, 7, 0, 0, 127, 128, 3, 22, 11, 0, 128,
		129, 3, 20, 10, 8, 129, 147, 1, 0, 0, 0, 130, 131, 10, 6, 0, 0, 131, 132,
		3, 24, 12, 0, 132, 133, 3, 20, 10, 7, 133, 147, 1, 0, 0, 0, 134, 135, 10,
		5, 0, 0, 135, 136, 3, 26, 13, 0, 136, 137, 3, 20, 10, 6, 137, 147, 1, 0,
		0, 0, 138, 139, 10, 4, 0, 0, 139, 140, 3, 28, 14, 0, 140, 141, 3, 20, 10,
		5, 141, 147, 1, 0, 0, 0, 142, 143, 10, 3, 0, 0, 143, 144, 3, 30, 15, 0,
		144, 145, 3, 20, 10, 4, 145, 147, 1, 0, 0, 0, 146, 126, 1, 0, 0, 0, 146,
		130, 1, 0, 0, 0, 146, 134, 1, 0, 0, 0, 146, 138, 1, 0, 0, 0, 146, 142,
		1, 0, 0, 0, 147, 150, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 148, 149, 1, 0,
		0, 0, 149, 21, 1, 0, 0, 0, 150, 148, 1, 0, 0, 0, 151, 152, 7, 2, 0, 0,
		152, 23, 1, 0, 0, 0, 153, 154, 7, 3, 0, 0, 154, 25, 1, 0, 0, 0, 155, 156,
		7, 4, 0, 0, 156, 27, 1, 0, 0, 0, 157, 158, 5, 18, 0, 0, 158, 29, 1, 0,
		0, 0, 159, 160, 5, 19, 0, 0, 160, 31, 1, 0, 0, 0, 161, 162, 6, 16, -1,
		0, 162, 168, 3, 34, 17, 0, 163, 168, 3, 36, 18, 0, 164, 168, 3, 42, 21,
		0, 165, 166, 5, 23, 0, 0, 166, 168, 3, 32, 16, 1, 167, 161, 1, 0, 0, 0,
		167, 163, 1, 0, 0, 0, 167, 164, 1, 0, 0, 0, 167, 165, 1, 0, 0, 0, 168,
		177, 1, 0, 0, 0, 169, 170, 10, 4, 0, 0, 170, 176, 3, 44, 22, 0, 171, 172,
		10, 3, 0, 0, 172, 176, 3, 40, 20, 0, 173, 174, 10, 2, 0, 0, 174, 176, 3,
		38, 19, 0, 175, 169, 1, 0, 0, 0, 175, 171, 1, 0, 0, 0, 175, 173, 1, 0,
		0, 0, 176, 179, 1, 0, 0, 0, 177, 175, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0,
		178, 33, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0, 180, 186, 3, 62, 31, 0, 181,
		186, 3, 54, 27, 0, 182, 186, 3, 48, 24, 0, 183, 186, 3, 64, 32, 0, 184,
		186, 5, 22, 0, 0, 185, 180, 1, 0, 0, 0, 185, 181, 1, 0, 0, 0, 185, 182,
		1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 185, 184, 1, 0, 0, 0, 186, 35, 1, 0,
		0, 0, 187, 188, 6, 18, -1, 0, 188, 189, 5, 38, 0, 0, 189, 196, 1, 0, 0,
		0, 190, 191, 10, 3, 0, 0, 191, 195, 3, 40, 20, 0, 192, 193, 10, 2, 0, 0,
		193, 195, 3, 38, 19, 0, 194, 190, 1, 0, 0, 0, 194, 192, 1, 0, 0, 0, 195,
		198, 1, 0, 0, 0, 196, 194, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0, 197, 37, 1,
		0, 0, 0, 198, 196, 1, 0, 0, 0, 199, 200, 5, 13, 0, 0, 200, 201, 3, 20,
		10, 0, 201, 202, 5, 14, 0, 0, 202, 39, 1, 0, 0, 0, 203, 204, 5, 7, 0, 0,
		204, 205, 5, 38, 0, 0, 205, 41, 1, 0, 0, 0, 206, 207, 5, 38, 0, 0, 207,
		209, 5, 11, 0, 0, 208, 210, 3, 46, 23, 0, 209, 208, 1, 0, 0, 0, 209, 210,
		1, 0, 0, 0, 210, 211, 1, 0, 0, 0, 211, 212, 5, 12, 0, 0, 212, 43, 1, 0,
		0, 0, 213, 214, 5, 7, 0, 0, 214, 215, 3, 42, 21, 0, 215, 45, 1, 0, 0, 0,
		216, 221, 3, 20, 10, 0, 217, 218, 5, 1, 0, 0, 218, 220, 3, 20, 10, 0, 219,
		217, 1, 0, 0, 0, 220, 223, 1, 0, 0, 0, 221, 219, 1, 0, 0, 0, 221, 222,
		1, 0, 0, 0, 222, 47, 1, 0, 0, 0, 223, 221, 1, 0, 0, 0, 224, 227, 3, 50,
		25, 0, 225, 227, 3, 52, 26, 0, 226, 224, 1, 0, 0, 0, 226, 225, 1, 0, 0,
		0, 227, 49, 1, 0, 0, 0, 228, 230, 5, 3, 0, 0, 229, 228, 1, 0, 0, 0, 229,
		230, 1, 0, 0, 0, 230, 231, 1, 0, 0, 0, 231, 232, 5, 41, 0, 0, 232, 51,
		1, 0, 0, 0, 233, 235, 5, 3, 0, 0, 234, 233, 1, 0, 0, 0, 234, 235, 1, 0,
		0, 0, 235, 236, 1, 0, 0, 0, 236, 237, 5, 43, 0, 0, 237, 53, 1, 0, 0, 0,
		238, 242, 3, 56, 28, 0, 239, 242, 3, 58, 29, 0, 240, 242, 3, 60, 30, 0,
		241, 238, 1, 0, 0, 0, 241, 239, 1, 0, 0, 0, 241, 240, 1, 0, 0, 0, 242,
		55, 1, 0, 0, 0, 243, 245, 5, 3, 0, 0, 244, 243, 1, 0, 0, 0, 244, 245, 1,
		0, 0, 0, 245, 246, 1, 0, 0, 0, 246, 247, 5, 45, 0, 0, 247, 57, 1, 0, 0,
		0, 248, 250, 5, 3, 0, 0, 249, 248, 1, 0, 0, 0, 249, 250, 1, 0, 0, 0, 250,
		251, 1, 0, 0, 0, 251, 252, 5, 46, 0, 0, 252, 59, 1, 0, 0, 0, 253, 255,
		5, 3, 0, 0, 254, 253, 1, 0, 0, 0, 254, 255, 1, 0, 0, 0, 255, 256, 1, 0,
		0, 0, 256, 257, 5, 47, 0, 0, 257, 61, 1, 0, 0, 0, 258, 259, 7, 0, 0, 0,
		259, 63, 1, 0, 0, 0, 260, 261, 7, 5, 0, 0, 261, 65, 1, 0, 0, 0, 24, 69,
		77, 80, 105, 109, 117, 124, 146, 148, 167, 175, 177, 185, 194, 196, 209,
		221, 226, 229, 234, 241, 244, 249, 254,
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

// grulev3ParserInit initializes any static state used to implement grulev3Parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// Newgrulev3Parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Grulev3ParserInit() {
	staticData := &grulev3ParserStaticData
	staticData.once.Do(grulev3ParserInit)
}

// Newgrulev3Parser produces a new parser instance for the optional input antlr.TokenStream.
func Newgrulev3Parser(input antlr.TokenStream) *grulev3Parser {
	Grulev3ParserInit()
	this := new(grulev3Parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &grulev3ParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "grulev3.g4"

	return this
}

// grulev3Parser tokens.
const (
	grulev3ParserEOF               = antlr.TokenEOF
	grulev3ParserT__0              = 1
	grulev3ParserPLUS              = 2
	grulev3ParserMINUS             = 3
	grulev3ParserDIV               = 4
	grulev3ParserMUL               = 5
	grulev3ParserMOD               = 6
	grulev3ParserDOT               = 7
	grulev3ParserSEMICOLON         = 8
	grulev3ParserLR_BRACE          = 9
	grulev3ParserRR_BRACE          = 10
	grulev3ParserLR_BRACKET        = 11
	grulev3ParserRR_BRACKET        = 12
	grulev3ParserLS_BRACKET        = 13
	grulev3ParserRS_BRACKET        = 14
	grulev3ParserRULE              = 15
	grulev3ParserWHEN              = 16
	grulev3ParserTHEN              = 17
	grulev3ParserAND               = 18
	grulev3ParserOR                = 19
	grulev3ParserTRUE              = 20
	grulev3ParserFALSE             = 21
	grulev3ParserNIL_LITERAL       = 22
	grulev3ParserNEGATION          = 23
	grulev3ParserSALIENCE          = 24
	grulev3ParserEQUALS            = 25
	grulev3ParserASSIGN            = 26
	grulev3ParserPLUS_ASIGN        = 27
	grulev3ParserMINUS_ASIGN       = 28
	grulev3ParserDIV_ASIGN         = 29
	grulev3ParserMUL_ASIGN         = 30
	grulev3ParserGT                = 31
	grulev3ParserLT                = 32
	grulev3ParserGTE               = 33
	grulev3ParserLTE               = 34
	grulev3ParserNOTEQUALS         = 35
	grulev3ParserBITAND            = 36
	grulev3ParserBITOR             = 37
	grulev3ParserSIMPLENAME        = 38
	grulev3ParserDQUOTA_STRING     = 39
	grulev3ParserSQUOTA_STRING     = 40
	grulev3ParserDECIMAL_FLOAT_LIT = 41
	grulev3ParserDECIMAL_EXPONENT  = 42
	grulev3ParserHEX_FLOAT_LIT     = 43
	grulev3ParserHEX_EXPONENT      = 44
	grulev3ParserDEC_LIT           = 45
	grulev3ParserHEX_LIT           = 46
	grulev3ParserOCT_LIT           = 47
	grulev3ParserSPACE             = 48
	grulev3ParserCOMMENT           = 49
	grulev3ParserLINE_COMMENT      = 50
)

// grulev3Parser rules.
const (
	grulev3ParserRULE_grl                     = 0
	grulev3ParserRULE_ruleEntry               = 1
	grulev3ParserRULE_salience                = 2
	grulev3ParserRULE_ruleName                = 3
	grulev3ParserRULE_ruleDescription         = 4
	grulev3ParserRULE_whenScope               = 5
	grulev3ParserRULE_thenScope               = 6
	grulev3ParserRULE_thenExpressionList      = 7
	grulev3ParserRULE_thenExpression          = 8
	grulev3ParserRULE_assignment              = 9
	grulev3ParserRULE_expression              = 10
	grulev3ParserRULE_mulDivOperators         = 11
	grulev3ParserRULE_addMinusOperators       = 12
	grulev3ParserRULE_comparisonOperator      = 13
	grulev3ParserRULE_andLogicOperator        = 14
	grulev3ParserRULE_orLogicOperator         = 15
	grulev3ParserRULE_expressionAtom          = 16
	grulev3ParserRULE_constant                = 17
	grulev3ParserRULE_variable                = 18
	grulev3ParserRULE_arrayMapSelector        = 19
	grulev3ParserRULE_memberVariable          = 20
	grulev3ParserRULE_functionCall            = 21
	grulev3ParserRULE_methodCall              = 22
	grulev3ParserRULE_argumentList            = 23
	grulev3ParserRULE_floatLiteral            = 24
	grulev3ParserRULE_decimalFloatLiteral     = 25
	grulev3ParserRULE_hexadecimalFloatLiteral = 26
	grulev3ParserRULE_integerLiteral          = 27
	grulev3ParserRULE_decimalLiteral          = 28
	grulev3ParserRULE_hexadecimalLiteral      = 29
	grulev3ParserRULE_octalLiteral            = 30
	grulev3ParserRULE_stringLiteral           = 31
	grulev3ParserRULE_booleanLiteral          = 32
)

// IGrlContext is an interface to support dynamic dispatch.
type IGrlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGrlContext differentiates from other interfaces.
	IsGrlContext()
}

type GrlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGrlContext() *GrlContext {
	var p = new(GrlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = grulev3ParserRULE_grl
	return p
}

func (*GrlContext) IsGrlContext() {}

func NewGrlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GrlContext {
	var p = new(GrlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = grulev3ParserRULE_grl

	return p
}

func (s *GrlContext) GetParser() antlr.Parser { return s.parser }

func (s *GrlContext) EOF() antlr.TerminalNode {
	return s.GetToken(grulev3ParserEOF, 0)
}

func (s *GrlContext) AllRuleEntry() []IRuleEntryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRuleEntryContext); ok {
			len++
		}
	}

	tst := make([]IRuleEntryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRuleEntryContext); ok {
			tst[i] = t.(IRuleEntryContext)
			i++
		}
	}

	return tst
}

func (s *GrlContext) RuleEntry(i int) IRuleEntryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRuleEntryContext); ok {
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

	return t.(IRuleEntryContext)
}

func (s *GrlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GrlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GrlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.EnterGrl(s)
	}
}

func (s *GrlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.ExitGrl(s)
	}
}

func (s *GrlContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case grulev3Visitor:
		return t.VisitGrl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *grulev3Parser) Grl() (localctx IGrlContext) {
	this := p
	_ = this

	localctx = NewGrlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, grulev3ParserRULE_grl)
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
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == grulev3ParserRULE {
		{
			p.SetState(66)
			p.RuleEntry()
		}

		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(72)
		p.Match(grulev3ParserEOF)
	}

	return localctx
}

// IRuleEntryContext is an interface to support dynamic dispatch.
type IRuleEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRuleEntryContext differentiates from other interfaces.
	IsRuleEntryContext()
}

type RuleEntryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleEntryContext() *RuleEntryContext {
	var p = new(RuleEntryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = grulev3ParserRULE_ruleEntry
	return p
}

func (*RuleEntryContext) IsRuleEntryContext() {}

func NewRuleEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleEntryContext {
	var p = new(RuleEntryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = grulev3ParserRULE_ruleEntry

	return p
}

func (s *RuleEntryContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleEntryContext) RULE() antlr.TerminalNode {
	return s.GetToken(grulev3ParserRULE, 0)
}

func (s *RuleEntryContext) RuleName() IRuleNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRuleNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRuleNameContext)
}

func (s *RuleEntryContext) LR_BRACE() antlr.TerminalNode {
	return s.GetToken(grulev3ParserLR_BRACE, 0)
}

func (s *RuleEntryContext) WhenScope() IWhenScopeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhenScopeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhenScopeContext)
}

func (s *RuleEntryContext) ThenScope() IThenScopeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IThenScopeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IThenScopeContext)
}

func (s *RuleEntryContext) RR_BRACE() antlr.TerminalNode {
	return s.GetToken(grulev3ParserRR_BRACE, 0)
}

func (s *RuleEntryContext) RuleDescription() IRuleDescriptionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRuleDescriptionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRuleDescriptionContext)
}

func (s *RuleEntryContext) Salience() ISalienceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISalienceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISalienceContext)
}

func (s *RuleEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleEntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.EnterRuleEntry(s)
	}
}

func (s *RuleEntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.ExitRuleEntry(s)
	}
}

func (s *RuleEntryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case grulev3Visitor:
		return t.VisitRuleEntry(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *grulev3Parser) RuleEntry() (localctx IRuleEntryContext) {
	this := p
	_ = this

	localctx = NewRuleEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, grulev3ParserRULE_ruleEntry)
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
		p.SetState(74)
		p.Match(grulev3ParserRULE)
	}
	{
		p.SetState(75)
		p.RuleName()
	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == grulev3ParserDQUOTA_STRING || _la == grulev3ParserSQUOTA_STRING {
		{
			p.SetState(76)
			p.RuleDescription()
		}

	}
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == grulev3ParserSALIENCE {
		{
			p.SetState(79)
			p.Salience()
		}

	}
	{
		p.SetState(82)
		p.Match(grulev3ParserLR_BRACE)
	}
	{
		p.SetState(83)
		p.WhenScope()
	}
	{
		p.SetState(84)
		p.ThenScope()
	}
	{
		p.SetState(85)
		p.Match(grulev3ParserRR_BRACE)
	}

	return localctx
}

// ISalienceContext is an interface to support dynamic dispatch.
type ISalienceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSalienceContext differentiates from other interfaces.
	IsSalienceContext()
}

type SalienceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySalienceContext() *SalienceContext {
	var p = new(SalienceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = grulev3ParserRULE_salience
	return p
}

func (*SalienceContext) IsSalienceContext() {}

func NewSalienceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SalienceContext {
	var p = new(SalienceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = grulev3ParserRULE_salience

	return p
}

func (s *SalienceContext) GetParser() antlr.Parser { return s.parser }

func (s *SalienceContext) SALIENCE() antlr.TerminalNode {
	return s.GetToken(grulev3ParserSALIENCE, 0)
}

func (s *SalienceContext) IntegerLiteral() IIntegerLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerLiteralContext)
}

func (s *SalienceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SalienceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SalienceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.EnterSalience(s)
	}
}

func (s *SalienceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.ExitSalience(s)
	}
}

func (s *SalienceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case grulev3Visitor:
		return t.VisitSalience(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *grulev3Parser) Salience() (localctx ISalienceContext) {
	this := p
	_ = this

	localctx = NewSalienceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, grulev3ParserRULE_salience)

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
		p.Match(grulev3ParserSALIENCE)
	}
	{
		p.SetState(88)
		p.IntegerLiteral()
	}

	return localctx
}

// IRuleNameContext is an interface to support dynamic dispatch.
type IRuleNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRuleNameContext differentiates from other interfaces.
	IsRuleNameContext()
}

type RuleNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleNameContext() *RuleNameContext {
	var p = new(RuleNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = grulev3ParserRULE_ruleName
	return p
}

func (*RuleNameContext) IsRuleNameContext() {}

func NewRuleNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleNameContext {
	var p = new(RuleNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = grulev3ParserRULE_ruleName

	return p
}

func (s *RuleNameContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleNameContext) SIMPLENAME() antlr.TerminalNode {
	return s.GetToken(grulev3ParserSIMPLENAME, 0)
}

func (s *RuleNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.EnterRuleName(s)
	}
}

func (s *RuleNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.ExitRuleName(s)
	}
}

func (s *RuleNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case grulev3Visitor:
		return t.VisitRuleName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *grulev3Parser) RuleName() (localctx IRuleNameContext) {
	this := p
	_ = this

	localctx = NewRuleNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, grulev3ParserRULE_ruleName)

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
		p.SetState(90)
		p.Match(grulev3ParserSIMPLENAME)
	}

	return localctx
}

// IRuleDescriptionContext is an interface to support dynamic dispatch.
type IRuleDescriptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRuleDescriptionContext differentiates from other interfaces.
	IsRuleDescriptionContext()
}

type RuleDescriptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleDescriptionContext() *RuleDescriptionContext {
	var p = new(RuleDescriptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = grulev3ParserRULE_ruleDescription
	return p
}

func (*RuleDescriptionContext) IsRuleDescriptionContext() {}

func NewRuleDescriptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleDescriptionContext {
	var p = new(RuleDescriptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = grulev3ParserRULE_ruleDescription

	return p
}

func (s *RuleDescriptionContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleDescriptionContext) DQUOTA_STRING() antlr.TerminalNode {
	return s.GetToken(grulev3ParserDQUOTA_STRING, 0)
}

func (s *RuleDescriptionContext) SQUOTA_STRING() antlr.TerminalNode {
	return s.GetToken(grulev3ParserSQUOTA_STRING, 0)
}

func (s *RuleDescriptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleDescriptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleDescriptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.EnterRuleDescription(s)
	}
}

func (s *RuleDescriptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.ExitRuleDescription(s)
	}
}

func (s *RuleDescriptionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case grulev3Visitor:
		return t.VisitRuleDescription(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *grulev3Parser) RuleDescription() (localctx IRuleDescriptionContext) {
	this := p
	_ = this

	localctx = NewRuleDescriptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, grulev3ParserRULE_ruleDescription)
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
		p.SetState(92)
		_la = p.GetTokenStream().LA(1)

		if !(_la == grulev3ParserDQUOTA_STRING || _la == grulev3ParserSQUOTA_STRING) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IWhenScopeContext is an interface to support dynamic dispatch.
type IWhenScopeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhenScopeContext differentiates from other interfaces.
	IsWhenScopeContext()
}

type WhenScopeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhenScopeContext() *WhenScopeContext {
	var p = new(WhenScopeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = grulev3ParserRULE_whenScope
	return p
}

func (*WhenScopeContext) IsWhenScopeContext() {}

func NewWhenScopeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhenScopeContext {
	var p = new(WhenScopeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = grulev3ParserRULE_whenScope

	return p
}

func (s *WhenScopeContext) GetParser() antlr.Parser { return s.parser }

func (s *WhenScopeContext) WHEN() antlr.TerminalNode {
	return s.GetToken(grulev3ParserWHEN, 0)
}

func (s *WhenScopeContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WhenScopeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhenScopeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhenScopeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.EnterWhenScope(s)
	}
}

func (s *WhenScopeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.ExitWhenScope(s)
	}
}

func (s *WhenScopeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case grulev3Visitor:
		return t.VisitWhenScope(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *grulev3Parser) WhenScope() (localctx IWhenScopeContext) {
	this := p
	_ = this

	localctx = NewWhenScopeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, grulev3ParserRULE_whenScope)

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
		p.SetState(94)
		p.Match(grulev3ParserWHEN)
	}
	{
		p.SetState(95)
		p.expression(0)
	}

	return localctx
}

// IThenScopeContext is an interface to support dynamic dispatch.
type IThenScopeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsThenScopeContext differentiates from other interfaces.
	IsThenScopeContext()
}

type ThenScopeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyThenScopeContext() *ThenScopeContext {
	var p = new(ThenScopeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = grulev3ParserRULE_thenScope
	return p
}

func (*ThenScopeContext) IsThenScopeContext() {}

func NewThenScopeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ThenScopeContext {
	var p = new(ThenScopeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = grulev3ParserRULE_thenScope

	return p
}

func (s *ThenScopeContext) GetParser() antlr.Parser { return s.parser }

func (s *ThenScopeContext) THEN() antlr.TerminalNode {
	return s.GetToken(grulev3ParserTHEN, 0)
}

func (s *ThenScopeContext) ThenExpressionList() IThenExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IThenExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IThenExpressionListContext)
}

func (s *ThenScopeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThenScopeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ThenScopeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.EnterThenScope(s)
	}
}

func (s *ThenScopeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.ExitThenScope(s)
	}
}

func (s *ThenScopeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case grulev3Visitor:
		return t.VisitThenScope(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *grulev3Parser) ThenScope() (localctx IThenScopeContext) {
	this := p
	_ = this

	localctx = NewThenScopeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, grulev3ParserRULE_thenScope)

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
		p.SetState(97)
		p.Match(grulev3ParserTHEN)
	}
	{
		p.SetState(98)
		p.ThenExpressionList()
	}

	return localctx
}

// IThenExpressionListContext is an interface to support dynamic dispatch.
type IThenExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsThenExpressionListContext differentiates from other interfaces.
	IsThenExpressionListContext()
}

type ThenExpressionListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyThenExpressionListContext() *ThenExpressionListContext {
	var p = new(ThenExpressionListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = grulev3ParserRULE_thenExpressionList
	return p
}

func (*ThenExpressionListContext) IsThenExpressionListContext() {}

func NewThenExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ThenExpressionListContext {
	var p = new(ThenExpressionListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = grulev3ParserRULE_thenExpressionList

	return p
}

func (s *ThenExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ThenExpressionListContext) AllThenExpression() []IThenExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IThenExpressionContext); ok {
			len++
		}
	}

	tst := make([]IThenExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IThenExpressionContext); ok {
			tst[i] = t.(IThenExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ThenExpressionListContext) ThenExpression(i int) IThenExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IThenExpressionContext); ok {
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

	return t.(IThenExpressionContext)
}

func (s *ThenExpressionListContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(grulev3ParserSEMICOLON)
}

func (s *ThenExpressionListContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(grulev3ParserSEMICOLON, i)
}

func (s *ThenExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThenExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ThenExpressionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.EnterThenExpressionList(s)
	}
}

func (s *ThenExpressionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.ExitThenExpressionList(s)
	}
}

func (s *ThenExpressionListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case grulev3Visitor:
		return t.VisitThenExpressionList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *grulev3Parser) ThenExpressionList() (localctx IThenExpressionListContext) {
	this := p
	_ = this

	localctx = NewThenExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, grulev3ParserRULE_thenExpressionList)
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
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((int64(1)<<uint(_la))&((int64(1)<<grulev3ParserMINUS)|(int64(1)<<grulev3ParserTRUE)|(int64(1)<<grulev3ParserFALSE)|(int64(1)<<grulev3ParserNIL_LITERAL)|(int64(1)<<grulev3ParserNEGATION))) != 0) || (((_la-38)&-(0x1f+1)) == 0 && ((int64(1)<<uint((_la-38)))&((int64(1)<<(grulev3ParserSIMPLENAME-38))|(int64(1)<<(grulev3ParserDQUOTA_STRING-38))|(int64(1)<<(grulev3ParserSQUOTA_STRING-38))|(int64(1)<<(grulev3ParserDECIMAL_FLOAT_LIT-38))|(int64(1)<<(grulev3ParserHEX_FLOAT_LIT-38))|(int64(1)<<(grulev3ParserDEC_LIT-38))|(int64(1)<<(grulev3ParserHEX_LIT-38))|(int64(1)<<(grulev3ParserOCT_LIT-38)))) != 0) {
		{
			p.SetState(100)
			p.ThenExpression()
		}
		{
			p.SetState(101)
			p.Match(grulev3ParserSEMICOLON)
		}

		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IThenExpressionContext is an interface to support dynamic dispatch.
type IThenExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsThenExpressionContext differentiates from other interfaces.
	IsThenExpressionContext()
}

type ThenExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyThenExpressionContext() *ThenExpressionContext {
	var p = new(ThenExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = grulev3ParserRULE_thenExpression
	return p
}

func (*ThenExpressionContext) IsThenExpressionContext() {}

func NewThenExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ThenExpressionContext {
	var p = new(ThenExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = grulev3ParserRULE_thenExpression

	return p
}

func (s *ThenExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ThenExpressionContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *ThenExpressionContext) ExpressionAtom() IExpressionAtomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionAtomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionAtomContext)
}

func (s *ThenExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThenExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ThenExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.EnterThenExpression(s)
	}
}

func (s *ThenExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.ExitThenExpression(s)
	}
}

func (s *ThenExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case grulev3Visitor:
		return t.VisitThenExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *grulev3Parser) ThenExpression() (localctx IThenExpressionContext) {
	this := p
	_ = this

	localctx = NewThenExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, grulev3ParserRULE_thenExpression)

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

	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(107)
			p.Assignment()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(108)
			p.expressionAtom(0)
		}

	}

	return localctx
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = grulev3ParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = grulev3ParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *AssignmentContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(grulev3ParserASSIGN, 0)
}

func (s *AssignmentContext) PLUS_ASIGN() antlr.TerminalNode {
	return s.GetToken(grulev3ParserPLUS_ASIGN, 0)
}

func (s *AssignmentContext) MINUS_ASIGN() antlr.TerminalNode {
	return s.GetToken(grulev3ParserMINUS_ASIGN, 0)
}

func (s *AssignmentContext) DIV_ASIGN() antlr.TerminalNode {
	return s.GetToken(grulev3ParserDIV_ASIGN, 0)
}

func (s *AssignmentContext) MUL_ASIGN() antlr.TerminalNode {
	return s.GetToken(grulev3ParserMUL_ASIGN, 0)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case grulev3Visitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *grulev3Parser) Assignment() (localctx IAssignmentContext) {
	this := p
	_ = this

	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, grulev3ParserRULE_assignment)
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
		p.SetState(111)
		p.variable(0)
	}
	{
		p.SetState(112)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((int64(1)<<uint(_la))&((int64(1)<<grulev3ParserASSIGN)|(int64(1)<<grulev3ParserPLUS_ASIGN)|(int64(1)<<grulev3ParserMINUS_ASIGN)|(int64(1)<<grulev3ParserDIV_ASIGN)|(int64(1)<<grulev3ParserMUL_ASIGN))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(113)
		p.expression(0)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = grulev3ParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = grulev3ParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(grulev3ParserLR_BRACKET, 0)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *ExpressionContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(grulev3ParserRR_BRACKET, 0)
}

func (s *ExpressionContext) NEGATION() antlr.TerminalNode {
	return s.GetToken(grulev3ParserNEGATION, 0)
}

func (s *ExpressionContext) ExpressionAtom() IExpressionAtomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionAtomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionAtomContext)
}

func (s *ExpressionContext) MulDivOperators() IMulDivOperatorsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMulDivOperatorsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMulDivOperatorsContext)
}

func (s *ExpressionContext) AddMinusOperators() IAddMinusOperatorsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddMinusOperatorsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddMinusOperatorsContext)
}

func (s *ExpressionContext) ComparisonOperator() IComparisonOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparisonOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparisonOperatorContext)
}

func (s *ExpressionContext) AndLogicOperator() IAndLogicOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAndLogicOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAndLogicOperatorContext)
}

func (s *ExpressionContext) OrLogicOperator() IOrLogicOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrLogicOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrLogicOperatorContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case grulev3Visitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *grulev3Parser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *grulev3Parser) expression(_p int) (localctx IExpressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, grulev3ParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == grulev3ParserNEGATION {
			{
				p.SetState(116)
				p.Match(grulev3ParserNEGATION)
			}

		}
		{
			p.SetState(119)
			p.Match(grulev3ParserLR_BRACKET)
		}
		{
			p.SetState(120)
			p.expression(0)
		}
		{
			p.SetState(121)
			p.Match(grulev3ParserRR_BRACKET)
		}

	case 2:
		{
			p.SetState(123)
			p.expressionAtom(0)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(146)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, grulev3ParserRULE_expression)
				p.SetState(126)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(127)
					p.MulDivOperators()
				}
				{
					p.SetState(128)
					p.expression(8)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, grulev3ParserRULE_expression)
				p.SetState(130)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(131)
					p.AddMinusOperators()
				}
				{
					p.SetState(132)
					p.expression(7)
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, grulev3ParserRULE_expression)
				p.SetState(134)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(135)
					p.ComparisonOperator()
				}
				{
					p.SetState(136)
					p.expression(6)
				}

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, grulev3ParserRULE_expression)
				p.SetState(138)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(139)
					p.AndLogicOperator()
				}
				{
					p.SetState(140)
					p.expression(5)
				}

			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, grulev3ParserRULE_expression)
				p.SetState(142)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(143)
					p.OrLogicOperator()
				}
				{
					p.SetState(144)
					p.expression(4)
				}

			}

		}
		p.SetState(150)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}

	return localctx
}

// IMulDivOperatorsContext is an interface to support dynamic dispatch.
type IMulDivOperatorsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMulDivOperatorsContext differentiates from other interfaces.
	IsMulDivOperatorsContext()
}

type MulDivOperatorsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMulDivOperatorsContext() *MulDivOperatorsContext {
	var p = new(MulDivOperatorsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = grulev3ParserRULE_mulDivOperators
	return p
}

func (*MulDivOperatorsContext) IsMulDivOperatorsContext() {}

func NewMulDivOperatorsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MulDivOperatorsContext {
	var p = new(MulDivOperatorsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = grulev3ParserRULE_mulDivOperators

	return p
}

func (s *MulDivOperatorsContext) GetParser() antlr.Parser { return s.parser }

func (s *MulDivOperatorsContext) MUL() antlr.TerminalNode {
	return s.GetToken(grulev3ParserMUL, 0)
}

func (s *MulDivOperatorsContext) DIV() antlr.TerminalNode {
	return s.GetToken(grulev3ParserDIV, 0)
}

func (s *MulDivOperatorsContext) MOD() antlr.TerminalNode {
	return s.GetToken(grulev3ParserMOD, 0)
}

func (s *MulDivOperatorsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivOperatorsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MulDivOperatorsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.EnterMulDivOperators(s)
	}
}

func (s *MulDivOperatorsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.ExitMulDivOperators(s)
	}
}

func (s *MulDivOperatorsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case grulev3Visitor:
		return t.VisitMulDivOperators(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *grulev3Parser) MulDivOperators() (localctx IMulDivOperatorsContext) {
	this := p
	_ = this

	localctx = NewMulDivOperatorsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, grulev3ParserRULE_mulDivOperators)
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
		p.SetState(151)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((int64(1)<<uint(_la))&((int64(1)<<grulev3ParserDIV)|(int64(1)<<grulev3ParserMUL)|(int64(1)<<grulev3ParserMOD))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAddMinusOperatorsContext is an interface to support dynamic dispatch.
type IAddMinusOperatorsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAddMinusOperatorsContext differentiates from other interfaces.
	IsAddMinusOperatorsContext()
}

type AddMinusOperatorsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddMinusOperatorsContext() *AddMinusOperatorsContext {
	var p = new(AddMinusOperatorsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = grulev3ParserRULE_addMinusOperators
	return p
}

func (*AddMinusOperatorsContext) IsAddMinusOperatorsContext() {}

func NewAddMinusOperatorsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddMinusOperatorsContext {
	var p = new(AddMinusOperatorsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = grulev3ParserRULE_addMinusOperators

	return p
}

func (s *AddMinusOperatorsContext) GetParser() antlr.Parser { return s.parser }

func (s *AddMinusOperatorsContext) PLUS() antlr.TerminalNode {
	return s.GetToken(grulev3ParserPLUS, 0)
}

func (s *AddMinusOperatorsContext) MINUS() antlr.TerminalNode {
	return s.GetToken(grulev3ParserMINUS, 0)
}

func (s *AddMinusOperatorsContext) BITAND() antlr.TerminalNode {
	return s.GetToken(grulev3ParserBITAND, 0)
}

func (s *AddMinusOperatorsContext) BITOR() antlr.TerminalNode {
	return s.GetToken(grulev3ParserBITOR, 0)
}

func (s *AddMinusOperatorsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddMinusOperatorsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddMinusOperatorsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.EnterAddMinusOperators(s)
	}
}

func (s *AddMinusOperatorsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.ExitAddMinusOperators(s)
	}
}

func (s *AddMinusOperatorsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case grulev3Visitor:
		return t.VisitAddMinusOperators(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *grulev3Parser) AddMinusOperators() (localctx IAddMinusOperatorsContext) {
	this := p
	_ = this

	localctx = NewAddMinusOperatorsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, grulev3ParserRULE_addMinusOperators)
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
		p.SetState(153)
		_la = p.GetTokenStream().LA(1)

		if !(_la == grulev3ParserPLUS || _la == grulev3ParserMINUS || _la == grulev3ParserBITAND || _la == grulev3ParserBITOR) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IComparisonOperatorContext is an interface to support dynamic dispatch.
type IComparisonOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparisonOperatorContext differentiates from other interfaces.
	IsComparisonOperatorContext()
}

type ComparisonOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonOperatorContext() *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = grulev3ParserRULE_comparisonOperator
	return p
}

func (*ComparisonOperatorContext) IsComparisonOperatorContext() {}

func NewComparisonOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = grulev3ParserRULE_comparisonOperator

	return p
}

func (s *ComparisonOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonOperatorContext) GT() antlr.TerminalNode {
	return s.GetToken(grulev3ParserGT, 0)
}

func (s *ComparisonOperatorContext) LT() antlr.TerminalNode {
	return s.GetToken(grulev3ParserLT, 0)
}

func (s *ComparisonOperatorContext) GTE() antlr.TerminalNode {
	return s.GetToken(grulev3ParserGTE, 0)
}

func (s *ComparisonOperatorContext) LTE() antlr.TerminalNode {
	return s.GetToken(grulev3ParserLTE, 0)
}

func (s *ComparisonOperatorContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(grulev3ParserEQUALS, 0)
}

func (s *ComparisonOperatorContext) NOTEQUALS() antlr.TerminalNode {
	return s.GetToken(grulev3ParserNOTEQUALS, 0)
}

func (s *ComparisonOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.EnterComparisonOperator(s)
	}
}

func (s *ComparisonOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.ExitComparisonOperator(s)
	}
}

func (s *ComparisonOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case grulev3Visitor:
		return t.VisitComparisonOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *grulev3Parser) ComparisonOperator() (localctx IComparisonOperatorContext) {
	this := p
	_ = this

	localctx = NewComparisonOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, grulev3ParserRULE_comparisonOperator)
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
		p.SetState(155)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-25)&-(0x1f+1)) == 0 && ((int64(1)<<uint((_la-25)))&((int64(1)<<(grulev3ParserEQUALS-25))|(int64(1)<<(grulev3ParserGT-25))|(int64(1)<<(grulev3ParserLT-25))|(int64(1)<<(grulev3ParserGTE-25))|(int64(1)<<(grulev3ParserLTE-25))|(int64(1)<<(grulev3ParserNOTEQUALS-25)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAndLogicOperatorContext is an interface to support dynamic dispatch.
type IAndLogicOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAndLogicOperatorContext differentiates from other interfaces.
	IsAndLogicOperatorContext()
}

type AndLogicOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAndLogicOperatorContext() *AndLogicOperatorContext {
	var p = new(AndLogicOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = grulev3ParserRULE_andLogicOperator
	return p
}

func (*AndLogicOperatorContext) IsAndLogicOperatorContext() {}

func NewAndLogicOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndLogicOperatorContext {
	var p = new(AndLogicOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = grulev3ParserRULE_andLogicOperator

	return p
}

func (s *AndLogicOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *AndLogicOperatorContext) AND() antlr.TerminalNode {
	return s.GetToken(grulev3ParserAND, 0)
}

func (s *AndLogicOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndLogicOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AndLogicOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.EnterAndLogicOperator(s)
	}
}

func (s *AndLogicOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.ExitAndLogicOperator(s)
	}
}

func (s *AndLogicOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case grulev3Visitor:
		return t.VisitAndLogicOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *grulev3Parser) AndLogicOperator() (localctx IAndLogicOperatorContext) {
	this := p
	_ = this

	localctx = NewAndLogicOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, grulev3ParserRULE_andLogicOperator)

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
		p.SetState(157)
		p.Match(grulev3ParserAND)
	}

	return localctx
}

// IOrLogicOperatorContext is an interface to support dynamic dispatch.
type IOrLogicOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrLogicOperatorContext differentiates from other interfaces.
	IsOrLogicOperatorContext()
}

type OrLogicOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrLogicOperatorContext() *OrLogicOperatorContext {
	var p = new(OrLogicOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = grulev3ParserRULE_orLogicOperator
	return p
}

func (*OrLogicOperatorContext) IsOrLogicOperatorContext() {}

func NewOrLogicOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrLogicOperatorContext {
	var p = new(OrLogicOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = grulev3ParserRULE_orLogicOperator

	return p
}

func (s *OrLogicOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *OrLogicOperatorContext) OR() antlr.TerminalNode {
	return s.GetToken(grulev3ParserOR, 0)
}

func (s *OrLogicOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrLogicOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrLogicOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.EnterOrLogicOperator(s)
	}
}

func (s *OrLogicOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.ExitOrLogicOperator(s)
	}
}

func (s *OrLogicOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case grulev3Visitor:
		return t.VisitOrLogicOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *grulev3Parser) OrLogicOperator() (localctx IOrLogicOperatorContext) {
	this := p
	_ = this

	localctx = NewOrLogicOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, grulev3ParserRULE_orLogicOperator)

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
		p.SetState(159)
		p.Match(grulev3ParserOR)
	}

	return localctx
}

// IExpressionAtomContext is an interface to support dynamic dispatch.
type IExpressionAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionAtomContext differentiates from other interfaces.
	IsExpressionAtomContext()
}

type ExpressionAtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionAtomContext() *ExpressionAtomContext {
	var p = new(ExpressionAtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = grulev3ParserRULE_expressionAtom
	return p
}

func (*ExpressionAtomContext) IsExpressionAtomContext() {}

func NewExpressionAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionAtomContext {
	var p = new(ExpressionAtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = grulev3ParserRULE_expressionAtom

	return p
}

func (s *ExpressionAtomContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionAtomContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ExpressionAtomContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ExpressionAtomContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *ExpressionAtomContext) NEGATION() antlr.TerminalNode {
	return s.GetToken(grulev3ParserNEGATION, 0)
}

func (s *ExpressionAtomContext) ExpressionAtom() IExpressionAtomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionAtomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionAtomContext)
}

func (s *ExpressionAtomContext) MethodCall() IMethodCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethodCallContext)
}

func (s *ExpressionAtomContext) MemberVariable() IMemberVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMemberVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMemberVariableContext)
}

func (s *ExpressionAtomContext) ArrayMapSelector() IArrayMapSelectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayMapSelectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayMapSelectorContext)
}

func (s *ExpressionAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionAtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.EnterExpressionAtom(s)
	}
}

func (s *ExpressionAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.ExitExpressionAtom(s)
	}
}

func (s *ExpressionAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case grulev3Visitor:
		return t.VisitExpressionAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *grulev3Parser) ExpressionAtom() (localctx IExpressionAtomContext) {
	return p.expressionAtom(0)
}

func (p *grulev3Parser) expressionAtom(_p int) (localctx IExpressionAtomContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionAtomContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionAtomContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 32
	p.EnterRecursionRule(localctx, 32, grulev3ParserRULE_expressionAtom, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(162)
			p.Constant()
		}

	case 2:
		{
			p.SetState(163)
			p.variable(0)
		}

	case 3:
		{
			p.SetState(164)
			p.FunctionCall()
		}

	case 4:
		{
			p.SetState(165)
			p.Match(grulev3ParserNEGATION)
		}
		{
			p.SetState(166)
			p.expressionAtom(1)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(175)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionAtomContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, grulev3ParserRULE_expressionAtom)
				p.SetState(169)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(170)
					p.MethodCall()
				}

			case 2:
				localctx = NewExpressionAtomContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, grulev3ParserRULE_expressionAtom)
				p.SetState(171)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(172)
					p.MemberVariable()
				}

			case 3:
				localctx = NewExpressionAtomContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, grulev3ParserRULE_expressionAtom)
				p.SetState(173)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(174)
					p.ArrayMapSelector()
				}

			}

		}
		p.SetState(179)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}

	return localctx
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = grulev3ParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = grulev3ParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) StringLiteral() IStringLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *ConstantContext) IntegerLiteral() IIntegerLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerLiteralContext)
}

func (s *ConstantContext) FloatLiteral() IFloatLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFloatLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFloatLiteralContext)
}

func (s *ConstantContext) BooleanLiteral() IBooleanLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooleanLiteralContext)
}

func (s *ConstantContext) NIL_LITERAL() antlr.TerminalNode {
	return s.GetToken(grulev3ParserNIL_LITERAL, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.ExitConstant(s)
	}
}

func (s *ConstantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case grulev3Visitor:
		return t.VisitConstant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *grulev3Parser) Constant() (localctx IConstantContext) {
	this := p
	_ = this

	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, grulev3ParserRULE_constant)

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

	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(180)
			p.StringLiteral()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(181)
			p.IntegerLiteral()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(182)
			p.FloatLiteral()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(183)
			p.BooleanLiteral()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(184)
			p.Match(grulev3ParserNIL_LITERAL)
		}

	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = grulev3ParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = grulev3ParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) SIMPLENAME() antlr.TerminalNode {
	return s.GetToken(grulev3ParserSIMPLENAME, 0)
}

func (s *VariableContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *VariableContext) MemberVariable() IMemberVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMemberVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMemberVariableContext)
}

func (s *VariableContext) ArrayMapSelector() IArrayMapSelectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayMapSelectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayMapSelectorContext)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.ExitVariable(s)
	}
}

func (s *VariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case grulev3Visitor:
		return t.VisitVariable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *grulev3Parser) Variable() (localctx IVariableContext) {
	return p.variable(0)
}

func (p *grulev3Parser) variable(_p int) (localctx IVariableContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewVariableContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IVariableContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 36
	p.EnterRecursionRule(localctx, 36, grulev3ParserRULE_variable, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(188)
		p.Match(grulev3ParserSIMPLENAME)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(194)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
			case 1:
				localctx = NewVariableContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, grulev3ParserRULE_variable)
				p.SetState(190)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(191)
					p.MemberVariable()
				}

			case 2:
				localctx = NewVariableContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, grulev3ParserRULE_variable)
				p.SetState(192)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(193)
					p.ArrayMapSelector()
				}

			}

		}
		p.SetState(198)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
	}

	return localctx
}

// IArrayMapSelectorContext is an interface to support dynamic dispatch.
type IArrayMapSelectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayMapSelectorContext differentiates from other interfaces.
	IsArrayMapSelectorContext()
}

type ArrayMapSelectorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayMapSelectorContext() *ArrayMapSelectorContext {
	var p = new(ArrayMapSelectorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = grulev3ParserRULE_arrayMapSelector
	return p
}

func (*ArrayMapSelectorContext) IsArrayMapSelectorContext() {}

func NewArrayMapSelectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayMapSelectorContext {
	var p = new(ArrayMapSelectorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = grulev3ParserRULE_arrayMapSelector

	return p
}

func (s *ArrayMapSelectorContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayMapSelectorContext) LS_BRACKET() antlr.TerminalNode {
	return s.GetToken(grulev3ParserLS_BRACKET, 0)
}

func (s *ArrayMapSelectorContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArrayMapSelectorContext) RS_BRACKET() antlr.TerminalNode {
	return s.GetToken(grulev3ParserRS_BRACKET, 0)
}

func (s *ArrayMapSelectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayMapSelectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayMapSelectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.EnterArrayMapSelector(s)
	}
}

func (s *ArrayMapSelectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.ExitArrayMapSelector(s)
	}
}

func (s *ArrayMapSelectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case grulev3Visitor:
		return t.VisitArrayMapSelector(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *grulev3Parser) ArrayMapSelector() (localctx IArrayMapSelectorContext) {
	this := p
	_ = this

	localctx = NewArrayMapSelectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, grulev3ParserRULE_arrayMapSelector)

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
		p.SetState(199)
		p.Match(grulev3ParserLS_BRACKET)
	}
	{
		p.SetState(200)
		p.expression(0)
	}
	{
		p.SetState(201)
		p.Match(grulev3ParserRS_BRACKET)
	}

	return localctx
}

// IMemberVariableContext is an interface to support dynamic dispatch.
type IMemberVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMemberVariableContext differentiates from other interfaces.
	IsMemberVariableContext()
}

type MemberVariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMemberVariableContext() *MemberVariableContext {
	var p = new(MemberVariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = grulev3ParserRULE_memberVariable
	return p
}

func (*MemberVariableContext) IsMemberVariableContext() {}

func NewMemberVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemberVariableContext {
	var p = new(MemberVariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = grulev3ParserRULE_memberVariable

	return p
}

func (s *MemberVariableContext) GetParser() antlr.Parser { return s.parser }

func (s *MemberVariableContext) DOT() antlr.TerminalNode {
	return s.GetToken(grulev3ParserDOT, 0)
}

func (s *MemberVariableContext) SIMPLENAME() antlr.TerminalNode {
	return s.GetToken(grulev3ParserSIMPLENAME, 0)
}

func (s *MemberVariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberVariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MemberVariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.EnterMemberVariable(s)
	}
}

func (s *MemberVariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.ExitMemberVariable(s)
	}
}

func (s *MemberVariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case grulev3Visitor:
		return t.VisitMemberVariable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *grulev3Parser) MemberVariable() (localctx IMemberVariableContext) {
	this := p
	_ = this

	localctx = NewMemberVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, grulev3ParserRULE_memberVariable)

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
		p.SetState(203)
		p.Match(grulev3ParserDOT)
	}
	{
		p.SetState(204)
		p.Match(grulev3ParserSIMPLENAME)
	}

	return localctx
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = grulev3ParserRULE_functionCall
	return p
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = grulev3ParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) SIMPLENAME() antlr.TerminalNode {
	return s.GetToken(grulev3ParserSIMPLENAME, 0)
}

func (s *FunctionCallContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(grulev3ParserLR_BRACKET, 0)
}

func (s *FunctionCallContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(grulev3ParserRR_BRACKET, 0)
}

func (s *FunctionCallContext) ArgumentList() IArgumentListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentListContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case grulev3Visitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *grulev3Parser) FunctionCall() (localctx IFunctionCallContext) {
	this := p
	_ = this

	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, grulev3ParserRULE_functionCall)
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
		p.SetState(206)
		p.Match(grulev3ParserSIMPLENAME)
	}
	{
		p.SetState(207)
		p.Match(grulev3ParserLR_BRACKET)
	}
	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((int64(1)<<uint(_la))&((int64(1)<<grulev3ParserMINUS)|(int64(1)<<grulev3ParserLR_BRACKET)|(int64(1)<<grulev3ParserTRUE)|(int64(1)<<grulev3ParserFALSE)|(int64(1)<<grulev3ParserNIL_LITERAL)|(int64(1)<<grulev3ParserNEGATION))) != 0) || (((_la-38)&-(0x1f+1)) == 0 && ((int64(1)<<uint((_la-38)))&((int64(1)<<(grulev3ParserSIMPLENAME-38))|(int64(1)<<(grulev3ParserDQUOTA_STRING-38))|(int64(1)<<(grulev3ParserSQUOTA_STRING-38))|(int64(1)<<(grulev3ParserDECIMAL_FLOAT_LIT-38))|(int64(1)<<(grulev3ParserHEX_FLOAT_LIT-38))|(int64(1)<<(grulev3ParserDEC_LIT-38))|(int64(1)<<(grulev3ParserHEX_LIT-38))|(int64(1)<<(grulev3ParserOCT_LIT-38)))) != 0) {
		{
			p.SetState(208)
			p.ArgumentList()
		}

	}
	{
		p.SetState(211)
		p.Match(grulev3ParserRR_BRACKET)
	}

	return localctx
}

// IMethodCallContext is an interface to support dynamic dispatch.
type IMethodCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMethodCallContext differentiates from other interfaces.
	IsMethodCallContext()
}

type MethodCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodCallContext() *MethodCallContext {
	var p = new(MethodCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = grulev3ParserRULE_methodCall
	return p
}

func (*MethodCallContext) IsMethodCallContext() {}

func NewMethodCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodCallContext {
	var p = new(MethodCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = grulev3ParserRULE_methodCall

	return p
}

func (s *MethodCallContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodCallContext) DOT() antlr.TerminalNode {
	return s.GetToken(grulev3ParserDOT, 0)
}

func (s *MethodCallContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *MethodCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.EnterMethodCall(s)
	}
}

func (s *MethodCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.ExitMethodCall(s)
	}
}

func (s *MethodCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case grulev3Visitor:
		return t.VisitMethodCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *grulev3Parser) MethodCall() (localctx IMethodCallContext) {
	this := p
	_ = this

	localctx = NewMethodCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, grulev3ParserRULE_methodCall)

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
		p.SetState(213)
		p.Match(grulev3ParserDOT)
	}
	{
		p.SetState(214)
		p.FunctionCall()
	}

	return localctx
}

// IArgumentListContext is an interface to support dynamic dispatch.
type IArgumentListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentListContext differentiates from other interfaces.
	IsArgumentListContext()
}

type ArgumentListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentListContext() *ArgumentListContext {
	var p = new(ArgumentListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = grulev3ParserRULE_argumentList
	return p
}

func (*ArgumentListContext) IsArgumentListContext() {}

func NewArgumentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentListContext {
	var p = new(ArgumentListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = grulev3ParserRULE_argumentList

	return p
}

func (s *ArgumentListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *ArgumentListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.EnterArgumentList(s)
	}
}

func (s *ArgumentListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.ExitArgumentList(s)
	}
}

func (s *ArgumentListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case grulev3Visitor:
		return t.VisitArgumentList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *grulev3Parser) ArgumentList() (localctx IArgumentListContext) {
	this := p
	_ = this

	localctx = NewArgumentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, grulev3ParserRULE_argumentList)
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
		p.SetState(216)
		p.expression(0)
	}
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == grulev3ParserT__0 {
		{
			p.SetState(217)
			p.Match(grulev3ParserT__0)
		}
		{
			p.SetState(218)
			p.expression(0)
		}

		p.SetState(223)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFloatLiteralContext is an interface to support dynamic dispatch.
type IFloatLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFloatLiteralContext differentiates from other interfaces.
	IsFloatLiteralContext()
}

type FloatLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatLiteralContext() *FloatLiteralContext {
	var p = new(FloatLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = grulev3ParserRULE_floatLiteral
	return p
}

func (*FloatLiteralContext) IsFloatLiteralContext() {}

func NewFloatLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatLiteralContext {
	var p = new(FloatLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = grulev3ParserRULE_floatLiteral

	return p
}

func (s *FloatLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatLiteralContext) DecimalFloatLiteral() IDecimalFloatLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDecimalFloatLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDecimalFloatLiteralContext)
}

func (s *FloatLiteralContext) HexadecimalFloatLiteral() IHexadecimalFloatLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHexadecimalFloatLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHexadecimalFloatLiteralContext)
}

func (s *FloatLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.EnterFloatLiteral(s)
	}
}

func (s *FloatLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.ExitFloatLiteral(s)
	}
}

func (s *FloatLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case grulev3Visitor:
		return t.VisitFloatLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *grulev3Parser) FloatLiteral() (localctx IFloatLiteralContext) {
	this := p
	_ = this

	localctx = NewFloatLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, grulev3ParserRULE_floatLiteral)

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

	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(224)
			p.DecimalFloatLiteral()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(225)
			p.HexadecimalFloatLiteral()
		}

	}

	return localctx
}

// IDecimalFloatLiteralContext is an interface to support dynamic dispatch.
type IDecimalFloatLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDecimalFloatLiteralContext differentiates from other interfaces.
	IsDecimalFloatLiteralContext()
}

type DecimalFloatLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecimalFloatLiteralContext() *DecimalFloatLiteralContext {
	var p = new(DecimalFloatLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = grulev3ParserRULE_decimalFloatLiteral
	return p
}

func (*DecimalFloatLiteralContext) IsDecimalFloatLiteralContext() {}

func NewDecimalFloatLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DecimalFloatLiteralContext {
	var p = new(DecimalFloatLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = grulev3ParserRULE_decimalFloatLiteral

	return p
}

func (s *DecimalFloatLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *DecimalFloatLiteralContext) DECIMAL_FLOAT_LIT() antlr.TerminalNode {
	return s.GetToken(grulev3ParserDECIMAL_FLOAT_LIT, 0)
}

func (s *DecimalFloatLiteralContext) MINUS() antlr.TerminalNode {
	return s.GetToken(grulev3ParserMINUS, 0)
}

func (s *DecimalFloatLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecimalFloatLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DecimalFloatLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.EnterDecimalFloatLiteral(s)
	}
}

func (s *DecimalFloatLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.ExitDecimalFloatLiteral(s)
	}
}

func (s *DecimalFloatLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case grulev3Visitor:
		return t.VisitDecimalFloatLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *grulev3Parser) DecimalFloatLiteral() (localctx IDecimalFloatLiteralContext) {
	this := p
	_ = this

	localctx = NewDecimalFloatLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, grulev3ParserRULE_decimalFloatLiteral)
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
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == grulev3ParserMINUS {
		{
			p.SetState(228)
			p.Match(grulev3ParserMINUS)
		}

	}
	{
		p.SetState(231)
		p.Match(grulev3ParserDECIMAL_FLOAT_LIT)
	}

	return localctx
}

// IHexadecimalFloatLiteralContext is an interface to support dynamic dispatch.
type IHexadecimalFloatLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHexadecimalFloatLiteralContext differentiates from other interfaces.
	IsHexadecimalFloatLiteralContext()
}

type HexadecimalFloatLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHexadecimalFloatLiteralContext() *HexadecimalFloatLiteralContext {
	var p = new(HexadecimalFloatLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = grulev3ParserRULE_hexadecimalFloatLiteral
	return p
}

func (*HexadecimalFloatLiteralContext) IsHexadecimalFloatLiteralContext() {}

func NewHexadecimalFloatLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HexadecimalFloatLiteralContext {
	var p = new(HexadecimalFloatLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = grulev3ParserRULE_hexadecimalFloatLiteral

	return p
}

func (s *HexadecimalFloatLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *HexadecimalFloatLiteralContext) HEX_FLOAT_LIT() antlr.TerminalNode {
	return s.GetToken(grulev3ParserHEX_FLOAT_LIT, 0)
}

func (s *HexadecimalFloatLiteralContext) MINUS() antlr.TerminalNode {
	return s.GetToken(grulev3ParserMINUS, 0)
}

func (s *HexadecimalFloatLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HexadecimalFloatLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HexadecimalFloatLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.EnterHexadecimalFloatLiteral(s)
	}
}

func (s *HexadecimalFloatLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.ExitHexadecimalFloatLiteral(s)
	}
}

func (s *HexadecimalFloatLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case grulev3Visitor:
		return t.VisitHexadecimalFloatLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *grulev3Parser) HexadecimalFloatLiteral() (localctx IHexadecimalFloatLiteralContext) {
	this := p
	_ = this

	localctx = NewHexadecimalFloatLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, grulev3ParserRULE_hexadecimalFloatLiteral)
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
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == grulev3ParserMINUS {
		{
			p.SetState(233)
			p.Match(grulev3ParserMINUS)
		}

	}
	{
		p.SetState(236)
		p.Match(grulev3ParserHEX_FLOAT_LIT)
	}

	return localctx
}

// IIntegerLiteralContext is an interface to support dynamic dispatch.
type IIntegerLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegerLiteralContext differentiates from other interfaces.
	IsIntegerLiteralContext()
}

type IntegerLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerLiteralContext() *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = grulev3ParserRULE_integerLiteral
	return p
}

func (*IntegerLiteralContext) IsIntegerLiteralContext() {}

func NewIntegerLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = grulev3ParserRULE_integerLiteral

	return p
}

func (s *IntegerLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerLiteralContext) DecimalLiteral() IDecimalLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDecimalLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDecimalLiteralContext)
}

func (s *IntegerLiteralContext) HexadecimalLiteral() IHexadecimalLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHexadecimalLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHexadecimalLiteralContext)
}

func (s *IntegerLiteralContext) OctalLiteral() IOctalLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOctalLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOctalLiteralContext)
}

func (s *IntegerLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.EnterIntegerLiteral(s)
	}
}

func (s *IntegerLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.ExitIntegerLiteral(s)
	}
}

func (s *IntegerLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case grulev3Visitor:
		return t.VisitIntegerLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *grulev3Parser) IntegerLiteral() (localctx IIntegerLiteralContext) {
	this := p
	_ = this

	localctx = NewIntegerLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, grulev3ParserRULE_integerLiteral)

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

	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(238)
			p.DecimalLiteral()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(239)
			p.HexadecimalLiteral()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(240)
			p.OctalLiteral()
		}

	}

	return localctx
}

// IDecimalLiteralContext is an interface to support dynamic dispatch.
type IDecimalLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDecimalLiteralContext differentiates from other interfaces.
	IsDecimalLiteralContext()
}

type DecimalLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecimalLiteralContext() *DecimalLiteralContext {
	var p = new(DecimalLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = grulev3ParserRULE_decimalLiteral
	return p
}

func (*DecimalLiteralContext) IsDecimalLiteralContext() {}

func NewDecimalLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DecimalLiteralContext {
	var p = new(DecimalLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = grulev3ParserRULE_decimalLiteral

	return p
}

func (s *DecimalLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *DecimalLiteralContext) DEC_LIT() antlr.TerminalNode {
	return s.GetToken(grulev3ParserDEC_LIT, 0)
}

func (s *DecimalLiteralContext) MINUS() antlr.TerminalNode {
	return s.GetToken(grulev3ParserMINUS, 0)
}

func (s *DecimalLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecimalLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DecimalLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.EnterDecimalLiteral(s)
	}
}

func (s *DecimalLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.ExitDecimalLiteral(s)
	}
}

func (s *DecimalLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case grulev3Visitor:
		return t.VisitDecimalLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *grulev3Parser) DecimalLiteral() (localctx IDecimalLiteralContext) {
	this := p
	_ = this

	localctx = NewDecimalLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, grulev3ParserRULE_decimalLiteral)
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
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == grulev3ParserMINUS {
		{
			p.SetState(243)
			p.Match(grulev3ParserMINUS)
		}

	}
	{
		p.SetState(246)
		p.Match(grulev3ParserDEC_LIT)
	}

	return localctx
}

// IHexadecimalLiteralContext is an interface to support dynamic dispatch.
type IHexadecimalLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHexadecimalLiteralContext differentiates from other interfaces.
	IsHexadecimalLiteralContext()
}

type HexadecimalLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHexadecimalLiteralContext() *HexadecimalLiteralContext {
	var p = new(HexadecimalLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = grulev3ParserRULE_hexadecimalLiteral
	return p
}

func (*HexadecimalLiteralContext) IsHexadecimalLiteralContext() {}

func NewHexadecimalLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HexadecimalLiteralContext {
	var p = new(HexadecimalLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = grulev3ParserRULE_hexadecimalLiteral

	return p
}

func (s *HexadecimalLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *HexadecimalLiteralContext) HEX_LIT() antlr.TerminalNode {
	return s.GetToken(grulev3ParserHEX_LIT, 0)
}

func (s *HexadecimalLiteralContext) MINUS() antlr.TerminalNode {
	return s.GetToken(grulev3ParserMINUS, 0)
}

func (s *HexadecimalLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HexadecimalLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HexadecimalLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.EnterHexadecimalLiteral(s)
	}
}

func (s *HexadecimalLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.ExitHexadecimalLiteral(s)
	}
}

func (s *HexadecimalLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case grulev3Visitor:
		return t.VisitHexadecimalLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *grulev3Parser) HexadecimalLiteral() (localctx IHexadecimalLiteralContext) {
	this := p
	_ = this

	localctx = NewHexadecimalLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, grulev3ParserRULE_hexadecimalLiteral)
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
	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == grulev3ParserMINUS {
		{
			p.SetState(248)
			p.Match(grulev3ParserMINUS)
		}

	}
	{
		p.SetState(251)
		p.Match(grulev3ParserHEX_LIT)
	}

	return localctx
}

// IOctalLiteralContext is an interface to support dynamic dispatch.
type IOctalLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOctalLiteralContext differentiates from other interfaces.
	IsOctalLiteralContext()
}

type OctalLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOctalLiteralContext() *OctalLiteralContext {
	var p = new(OctalLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = grulev3ParserRULE_octalLiteral
	return p
}

func (*OctalLiteralContext) IsOctalLiteralContext() {}

func NewOctalLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OctalLiteralContext {
	var p = new(OctalLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = grulev3ParserRULE_octalLiteral

	return p
}

func (s *OctalLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *OctalLiteralContext) OCT_LIT() antlr.TerminalNode {
	return s.GetToken(grulev3ParserOCT_LIT, 0)
}

func (s *OctalLiteralContext) MINUS() antlr.TerminalNode {
	return s.GetToken(grulev3ParserMINUS, 0)
}

func (s *OctalLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OctalLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OctalLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.EnterOctalLiteral(s)
	}
}

func (s *OctalLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.ExitOctalLiteral(s)
	}
}

func (s *OctalLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case grulev3Visitor:
		return t.VisitOctalLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *grulev3Parser) OctalLiteral() (localctx IOctalLiteralContext) {
	this := p
	_ = this

	localctx = NewOctalLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, grulev3ParserRULE_octalLiteral)
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
	p.SetState(254)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == grulev3ParserMINUS {
		{
			p.SetState(253)
			p.Match(grulev3ParserMINUS)
		}

	}
	{
		p.SetState(256)
		p.Match(grulev3ParserOCT_LIT)
	}

	return localctx
}

// IStringLiteralContext is an interface to support dynamic dispatch.
type IStringLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringLiteralContext differentiates from other interfaces.
	IsStringLiteralContext()
}

type StringLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringLiteralContext() *StringLiteralContext {
	var p = new(StringLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = grulev3ParserRULE_stringLiteral
	return p
}

func (*StringLiteralContext) IsStringLiteralContext() {}

func NewStringLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringLiteralContext {
	var p = new(StringLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = grulev3ParserRULE_stringLiteral

	return p
}

func (s *StringLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *StringLiteralContext) DQUOTA_STRING() antlr.TerminalNode {
	return s.GetToken(grulev3ParserDQUOTA_STRING, 0)
}

func (s *StringLiteralContext) SQUOTA_STRING() antlr.TerminalNode {
	return s.GetToken(grulev3ParserSQUOTA_STRING, 0)
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

func (s *StringLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case grulev3Visitor:
		return t.VisitStringLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *grulev3Parser) StringLiteral() (localctx IStringLiteralContext) {
	this := p
	_ = this

	localctx = NewStringLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, grulev3ParserRULE_stringLiteral)
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
		p.SetState(258)
		_la = p.GetTokenStream().LA(1)

		if !(_la == grulev3ParserDQUOTA_STRING || _la == grulev3ParserSQUOTA_STRING) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBooleanLiteralContext is an interface to support dynamic dispatch.
type IBooleanLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBooleanLiteralContext differentiates from other interfaces.
	IsBooleanLiteralContext()
}

type BooleanLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanLiteralContext() *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = grulev3ParserRULE_booleanLiteral
	return p
}

func (*BooleanLiteralContext) IsBooleanLiteralContext() {}

func NewBooleanLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = grulev3ParserRULE_booleanLiteral

	return p
}

func (s *BooleanLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanLiteralContext) TRUE() antlr.TerminalNode {
	return s.GetToken(grulev3ParserTRUE, 0)
}

func (s *BooleanLiteralContext) FALSE() antlr.TerminalNode {
	return s.GetToken(grulev3ParserFALSE, 0)
}

func (s *BooleanLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.EnterBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(grulev3Listener); ok {
		listenerT.ExitBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case grulev3Visitor:
		return t.VisitBooleanLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *grulev3Parser) BooleanLiteral() (localctx IBooleanLiteralContext) {
	this := p
	_ = this

	localctx = NewBooleanLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, grulev3ParserRULE_booleanLiteral)
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
		p.SetState(260)
		_la = p.GetTokenStream().LA(1)

		if !(_la == grulev3ParserTRUE || _la == grulev3ParserFALSE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *grulev3Parser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 10:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 16:
		var t *ExpressionAtomContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionAtomContext)
		}
		return p.ExpressionAtom_Sempred(t, predIndex)

	case 18:
		var t *VariableContext = nil
		if localctx != nil {
			t = localctx.(*VariableContext)
		}
		return p.Variable_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *grulev3Parser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *grulev3Parser) ExpressionAtom_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *grulev3Parser) Variable_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 8:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
