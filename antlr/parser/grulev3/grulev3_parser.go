// Code generated from /Users/ferdinandneman/Laboratory/Golang/src/github.com/newm4n/grule-rule-engine/antlr/grulev3.g4 by ANTLR 4.8. DO NOT EDIT.

package grulev3 // grulev3
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 52, 269, 
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9, 
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23, 
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4, 
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34, 
	9, 34, 3, 2, 7, 2, 70, 10, 2, 12, 2, 14, 2, 73, 11, 2, 3, 2, 3, 2, 3, 3, 
	3, 3, 3, 3, 5, 3, 80, 10, 3, 3, 3, 5, 3, 83, 10, 3, 3, 3, 3, 3, 3, 3, 3, 
	3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 
	8, 3, 8, 3, 8, 3, 9, 6, 9, 104, 10, 9, 13, 9, 14, 9, 105, 3, 10, 3, 10, 
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 118, 10, 
	10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 5, 12, 126, 10, 12, 3, 12, 
	3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 133, 10, 12, 3, 12, 3, 12, 3, 12, 3, 
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 155, 10, 12, 12, 12, 14, 
	12, 158, 11, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 
	3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 176, 10, 
	18, 3, 18, 3, 18, 3, 18, 3, 18, 7, 18, 182, 10, 18, 12, 18, 14, 18, 185, 
	11, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 192, 10, 19, 3, 20, 3, 
	20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 7, 20, 201, 10, 20, 12, 20, 14, 
	20, 204, 11, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 23, 
	3, 23, 3, 23, 5, 23, 216, 10, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 
	25, 3, 25, 3, 25, 7, 25, 226, 10, 25, 12, 25, 14, 25, 229, 11, 25, 3, 26, 
	3, 26, 5, 26, 233, 10, 26, 3, 27, 5, 27, 236, 10, 27, 3, 27, 3, 27, 3, 
	28, 5, 28, 241, 10, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 5, 29, 248, 
	10, 29, 3, 30, 5, 30, 251, 10, 30, 3, 30, 3, 30, 3, 31, 5, 31, 256, 10, 
	31, 3, 31, 3, 31, 3, 32, 5, 32, 261, 10, 32, 3, 32, 3, 32, 3, 33, 3, 33, 
	3, 34, 3, 34, 3, 34, 2, 5, 22, 34, 38, 35, 2, 4, 6, 8, 10, 12, 14, 16, 
	18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 
	54, 56, 58, 60, 62, 64, 66, 2, 8, 3, 2, 41, 42, 3, 2, 28, 32, 3, 2, 6, 
	8, 4, 2, 4, 5, 38, 39, 4, 2, 27, 27, 33, 37, 3, 2, 22, 23, 2, 269, 2, 71, 
	3, 2, 2, 2, 4, 76, 3, 2, 2, 2, 6, 89, 3, 2, 2, 2, 8, 92, 3, 2, 2, 2, 10, 
	94, 3, 2, 2, 2, 12, 96, 3, 2, 2, 2, 14, 99, 3, 2, 2, 2, 16, 103, 3, 2, 
	2, 2, 18, 117, 3, 2, 2, 2, 20, 119, 3, 2, 2, 2, 22, 132, 3, 2, 2, 2, 24, 
	159, 3, 2, 2, 2, 26, 161, 3, 2, 2, 2, 28, 163, 3, 2, 2, 2, 30, 165, 3, 
	2, 2, 2, 32, 167, 3, 2, 2, 2, 34, 175, 3, 2, 2, 2, 36, 191, 3, 2, 2, 2, 
	38, 193, 3, 2, 2, 2, 40, 205, 3, 2, 2, 2, 42, 209, 3, 2, 2, 2, 44, 212, 
	3, 2, 2, 2, 46, 219, 3, 2, 2, 2, 48, 222, 3, 2, 2, 2, 50, 232, 3, 2, 2, 
	2, 52, 235, 3, 2, 2, 2, 54, 240, 3, 2, 2, 2, 56, 247, 3, 2, 2, 2, 58, 250, 
	3, 2, 2, 2, 60, 255, 3, 2, 2, 2, 62, 260, 3, 2, 2, 2, 64, 264, 3, 2, 2, 
	2, 66, 266, 3, 2, 2, 2, 68, 70, 5, 4, 3, 2, 69, 68, 3, 2, 2, 2, 70, 73, 
	3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 74, 3, 2, 2, 2, 
	73, 71, 3, 2, 2, 2, 74, 75, 7, 2, 2, 3, 75, 3, 3, 2, 2, 2, 76, 77, 7, 17, 
	2, 2, 77, 79, 5, 8, 5, 2, 78, 80, 5, 10, 6, 2, 79, 78, 3, 2, 2, 2, 79, 
	80, 3, 2, 2, 2, 80, 82, 3, 2, 2, 2, 81, 83, 5, 6, 4, 2, 82, 81, 3, 2, 2, 
	2, 82, 83, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 85, 7, 11, 2, 2, 85, 86, 
	5, 12, 7, 2, 86, 87, 5, 14, 8, 2, 87, 88, 7, 12, 2, 2, 88, 5, 3, 2, 2, 
	2, 89, 90, 7, 26, 2, 2, 90, 91, 5, 56, 29, 2, 91, 7, 3, 2, 2, 2, 92, 93, 
	7, 40, 2, 2, 93, 9, 3, 2, 2, 2, 94, 95, 9, 2, 2, 2, 95, 11, 3, 2, 2, 2, 
	96, 97, 7, 18, 2, 2, 97, 98, 5, 22, 12, 2, 98, 13, 3, 2, 2, 2, 99, 100, 
	7, 19, 2, 2, 100, 101, 5, 16, 9, 2, 101, 15, 3, 2, 2, 2, 102, 104, 5, 18, 
	10, 2, 103, 102, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105, 103, 3, 2, 2, 2, 
	105, 106, 3, 2, 2, 2, 106, 17, 3, 2, 2, 2, 107, 108, 5, 20, 11, 2, 108, 
	109, 7, 10, 2, 2, 109, 118, 3, 2, 2, 2, 110, 111, 5, 44, 23, 2, 111, 112, 
	7, 10, 2, 2, 112, 118, 3, 2, 2, 2, 113, 114, 5, 38, 20, 2, 114, 115, 5, 
	46, 24, 2, 115, 116, 7, 10, 2, 2, 116, 118, 3, 2, 2, 2, 117, 107, 3, 2, 
	2, 2, 117, 110, 3, 2, 2, 2, 117, 113, 3, 2, 2, 2, 118, 19, 3, 2, 2, 2, 
	119, 120, 5, 38, 20, 2, 120, 121, 9, 3, 2, 2, 121, 122, 5, 22, 12, 2, 122, 
	21, 3, 2, 2, 2, 123, 125, 8, 12, 1, 2, 124, 126, 7, 25, 2, 2, 125, 124, 
	3, 2, 2, 2, 125, 126, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127, 128, 7, 13, 
	2, 2, 128, 129, 5, 22, 12, 2, 129, 130, 7, 14, 2, 2, 130, 133, 3, 2, 2, 
	2, 131, 133, 5, 34, 18, 2, 132, 123, 3, 2, 2, 2, 132, 131, 3, 2, 2, 2, 
	133, 156, 3, 2, 2, 2, 134, 135, 12, 9, 2, 2, 135, 136, 5, 24, 13, 2, 136, 
	137, 5, 22, 12, 10, 137, 155, 3, 2, 2, 2, 138, 139, 12, 8, 2, 2, 139, 140, 
	5, 26, 14, 2, 140, 141, 5, 22, 12, 9, 141, 155, 3, 2, 2, 2, 142, 143, 12, 
	7, 2, 2, 143, 144, 5, 28, 15, 2, 144, 145, 5, 22, 12, 8, 145, 155, 3, 2, 
	2, 2, 146, 147, 12, 6, 2, 2, 147, 148, 5, 30, 16, 2, 148, 149, 5, 22, 12, 
	7, 149, 155, 3, 2, 2, 2, 150, 151, 12, 5, 2, 2, 151, 152, 5, 32, 17, 2, 
	152, 153, 5, 22, 12, 6, 153, 155, 3, 2, 2, 2, 154, 134, 3, 2, 2, 2, 154, 
	138, 3, 2, 2, 2, 154, 142, 3, 2, 2, 2, 154, 146, 3, 2, 2, 2, 154, 150, 
	3, 2, 2, 2, 155, 158, 3, 2, 2, 2, 156, 154, 3, 2, 2, 2, 156, 157, 3, 2, 
	2, 2, 157, 23, 3, 2, 2, 2, 158, 156, 3, 2, 2, 2, 159, 160, 9, 4, 2, 2, 
	160, 25, 3, 2, 2, 2, 161, 162, 9, 5, 2, 2, 162, 27, 3, 2, 2, 2, 163, 164, 
	9, 6, 2, 2, 164, 29, 3, 2, 2, 2, 165, 166, 7, 20, 2, 2, 166, 31, 3, 2, 
	2, 2, 167, 168, 7, 21, 2, 2, 168, 33, 3, 2, 2, 2, 169, 170, 8, 18, 1, 2, 
	170, 176, 5, 36, 19, 2, 171, 176, 5, 44, 23, 2, 172, 176, 5, 38, 20, 2, 
	173, 174, 7, 25, 2, 2, 174, 176, 5, 34, 18, 3, 175, 169, 3, 2, 2, 2, 175, 
	171, 3, 2, 2, 2, 175, 172, 3, 2, 2, 2, 175, 173, 3, 2, 2, 2, 176, 183, 
	3, 2, 2, 2, 177, 178, 12, 5, 2, 2, 178, 182, 5, 46, 24, 2, 179, 180, 12, 
	4, 2, 2, 180, 182, 5, 42, 22, 2, 181, 177, 3, 2, 2, 2, 181, 179, 3, 2, 
	2, 2, 182, 185, 3, 2, 2, 2, 183, 181, 3, 2, 2, 2, 183, 184, 3, 2, 2, 2, 
	184, 35, 3, 2, 2, 2, 185, 183, 3, 2, 2, 2, 186, 192, 5, 64, 33, 2, 187, 
	192, 5, 56, 29, 2, 188, 192, 5, 50, 26, 2, 189, 192, 5, 66, 34, 2, 190, 
	192, 7, 24, 2, 2, 191, 186, 3, 2, 2, 2, 191, 187, 3, 2, 2, 2, 191, 188, 
	3, 2, 2, 2, 191, 189, 3, 2, 2, 2, 191, 190, 3, 2, 2, 2, 192, 37, 3, 2, 
	2, 2, 193, 194, 8, 20, 1, 2, 194, 195, 7, 40, 2, 2, 195, 202, 3, 2, 2, 
	2, 196, 197, 12, 4, 2, 2, 197, 201, 5, 42, 22, 2, 198, 199, 12, 3, 2, 2, 
	199, 201, 5, 40, 21, 2, 200, 196, 3, 2, 2, 2, 200, 198, 3, 2, 2, 2, 201, 
	204, 3, 2, 2, 2, 202, 200, 3, 2, 2, 2, 202, 203, 3, 2, 2, 2, 203, 39, 3, 
	2, 2, 2, 204, 202, 3, 2, 2, 2, 205, 206, 7, 15, 2, 2, 206, 207, 5, 22, 
	12, 2, 207, 208, 7, 16, 2, 2, 208, 41, 3, 2, 2, 2, 209, 210, 7, 9, 2, 2, 
	210, 211, 7, 40, 2, 2, 211, 43, 3, 2, 2, 2, 212, 213, 7, 40, 2, 2, 213, 
	215, 7, 13, 2, 2, 214, 216, 5, 48, 25, 2, 215, 214, 3, 2, 2, 2, 215, 216, 
	3, 2, 2, 2, 216, 217, 3, 2, 2, 2, 217, 218, 7, 14, 2, 2, 218, 45, 3, 2, 
	2, 2, 219, 220, 7, 9, 2, 2, 220, 221, 5, 44, 23, 2, 221, 47, 3, 2, 2, 2, 
	222, 227, 5, 22, 12, 2, 223, 224, 7, 3, 2, 2, 224, 226, 5, 22, 12, 2, 225, 
	223, 3, 2, 2, 2, 226, 229, 3, 2, 2, 2, 227, 225, 3, 2, 2, 2, 227, 228, 
	3, 2, 2, 2, 228, 49, 3, 2, 2, 2, 229, 227, 3, 2, 2, 2, 230, 233, 5, 52, 
	27, 2, 231, 233, 5, 54, 28, 2, 232, 230, 3, 2, 2, 2, 232, 231, 3, 2, 2, 
	2, 233, 51, 3, 2, 2, 2, 234, 236, 7, 5, 2, 2, 235, 234, 3, 2, 2, 2, 235, 
	236, 3, 2, 2, 2, 236, 237, 3, 2, 2, 2, 237, 238, 7, 43, 2, 2, 238, 53, 
	3, 2, 2, 2, 239, 241, 7, 5, 2, 2, 240, 239, 3, 2, 2, 2, 240, 241, 3, 2, 
	2, 2, 241, 242, 3, 2, 2, 2, 242, 243, 7, 45, 2, 2, 243, 55, 3, 2, 2, 2, 
	244, 248, 5, 58, 30, 2, 245, 248, 5, 60, 31, 2, 246, 248, 5, 62, 32, 2, 
	247, 244, 3, 2, 2, 2, 247, 245, 3, 2, 2, 2, 247, 246, 3, 2, 2, 2, 248, 
	57, 3, 2, 2, 2, 249, 251, 7, 5, 2, 2, 250, 249, 3, 2, 2, 2, 250, 251, 3, 
	2, 2, 2, 251, 252, 3, 2, 2, 2, 252, 253, 7, 47, 2, 2, 253, 59, 3, 2, 2, 
	2, 254, 256, 7, 5, 2, 2, 255, 254, 3, 2, 2, 2, 255, 256, 3, 2, 2, 2, 256, 
	257, 3, 2, 2, 2, 257, 258, 7, 48, 2, 2, 258, 61, 3, 2, 2, 2, 259, 261, 
	7, 5, 2, 2, 260, 259, 3, 2, 2, 2, 260, 261, 3, 2, 2, 2, 261, 262, 3, 2, 
	2, 2, 262, 263, 7, 49, 2, 2, 263, 63, 3, 2, 2, 2, 264, 265, 9, 2, 2, 2, 
	265, 65, 3, 2, 2, 2, 266, 267, 9, 7, 2, 2, 267, 67, 3, 2, 2, 2, 26, 71, 
	79, 82, 105, 117, 125, 132, 154, 156, 175, 181, 183, 191, 200, 202, 215, 
	227, 232, 235, 240, 247, 250, 255, 260,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "','", "'+'", "'-'", "'/'", "'*'", "'%'", "'.'", "';'", "'{'", "'}'", 
	"'('", "')'", "'['", "']'", "", "", "", "'&&'", "'||'", "", "", "", "'!'", 
	"", "'=='", "'='", "'+='", "'-='", "'/='", "'*='", "'>'", "'<'", "'>='", 
	"'<='", "'!='", "'&'", "'|'",
}
var symbolicNames = []string{
	"", "", "PLUS", "MINUS", "DIV", "MUL", "MOD", "DOT", "SEMICOLON", "LR_BRACE", 
	"RR_BRACE", "LR_BRACKET", "RR_BRACKET", "LS_BRACKET", "RS_BRACKET", "RULE", 
	"WHEN", "THEN", "AND", "OR", "TRUE", "FALSE", "NIL_LITERAL", "NEGATION", 
	"SALIENCE", "EQUALS", "ASSIGN", "PLUS_ASIGN", "MINUS_ASIGN", "DIV_ASIGN", 
	"MUL_ASIGN", "GT", "LT", "GTE", "LTE", "NOTEQUALS", "BITAND", "BITOR", 
	"SIMPLENAME", "DQUOTA_STRING", "SQUOTA_STRING", "DECIMAL_FLOAT_LIT", "DECIMAL_EXPONENT", 
	"HEX_FLOAT_LIT", "HEX_EXPONENT", "DEC_LIT", "HEX_LIT", "OCT_LIT", "SPACE", 
	"COMMENT", "LINE_COMMENT",
}

var ruleNames = []string{
	"grl", "ruleEntry", "salience", "ruleName", "ruleDescription", "whenScope", 
	"thenScope", "thenExpressionList", "thenExpression", "assignment", "expression", 
	"mulDivOperators", "addMinusOperators", "comparisonOperator", "andLogicOperator", 
	"orLogicOperator", "expressionAtom", "constant", "variable", "arrayMapSelector", 
	"memberVariable", "functionCall", "methodCall", "argumentList", "floatLiteral", 
	"decimalFloatLiteral", "hexadecimalFloatLiteral", "integerLiteral", "decimalLiteral", 
	"hexadecimalLiteral", "octalLiteral", "stringLiteral", "booleanLiteral",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type grulev3Parser struct {
	*antlr.BaseParser
}

func Newgrulev3Parser(input antlr.TokenStream) *grulev3Parser {
	this := new(grulev3Parser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "grulev3.g4"

	return this
}

// grulev3Parser tokens.
const (
	grulev3ParserEOF = antlr.TokenEOF
	grulev3ParserT__0 = 1
	grulev3ParserPLUS = 2
	grulev3ParserMINUS = 3
	grulev3ParserDIV = 4
	grulev3ParserMUL = 5
	grulev3ParserMOD = 6
	grulev3ParserDOT = 7
	grulev3ParserSEMICOLON = 8
	grulev3ParserLR_BRACE = 9
	grulev3ParserRR_BRACE = 10
	grulev3ParserLR_BRACKET = 11
	grulev3ParserRR_BRACKET = 12
	grulev3ParserLS_BRACKET = 13
	grulev3ParserRS_BRACKET = 14
	grulev3ParserRULE = 15
	grulev3ParserWHEN = 16
	grulev3ParserTHEN = 17
	grulev3ParserAND = 18
	grulev3ParserOR = 19
	grulev3ParserTRUE = 20
	grulev3ParserFALSE = 21
	grulev3ParserNIL_LITERAL = 22
	grulev3ParserNEGATION = 23
	grulev3ParserSALIENCE = 24
	grulev3ParserEQUALS = 25
	grulev3ParserASSIGN = 26
	grulev3ParserPLUS_ASIGN = 27
	grulev3ParserMINUS_ASIGN = 28
	grulev3ParserDIV_ASIGN = 29
	grulev3ParserMUL_ASIGN = 30
	grulev3ParserGT = 31
	grulev3ParserLT = 32
	grulev3ParserGTE = 33
	grulev3ParserLTE = 34
	grulev3ParserNOTEQUALS = 35
	grulev3ParserBITAND = 36
	grulev3ParserBITOR = 37
	grulev3ParserSIMPLENAME = 38
	grulev3ParserDQUOTA_STRING = 39
	grulev3ParserSQUOTA_STRING = 40
	grulev3ParserDECIMAL_FLOAT_LIT = 41
	grulev3ParserDECIMAL_EXPONENT = 42
	grulev3ParserHEX_FLOAT_LIT = 43
	grulev3ParserHEX_EXPONENT = 44
	grulev3ParserDEC_LIT = 45
	grulev3ParserHEX_LIT = 46
	grulev3ParserOCT_LIT = 47
	grulev3ParserSPACE = 48
	grulev3ParserCOMMENT = 49
	grulev3ParserLINE_COMMENT = 50
)

// grulev3Parser rules.
const (
	grulev3ParserRULE_grl = 0
	grulev3ParserRULE_ruleEntry = 1
	grulev3ParserRULE_salience = 2
	grulev3ParserRULE_ruleName = 3
	grulev3ParserRULE_ruleDescription = 4
	grulev3ParserRULE_whenScope = 5
	grulev3ParserRULE_thenScope = 6
	grulev3ParserRULE_thenExpressionList = 7
	grulev3ParserRULE_thenExpression = 8
	grulev3ParserRULE_assignment = 9
	grulev3ParserRULE_expression = 10
	grulev3ParserRULE_mulDivOperators = 11
	grulev3ParserRULE_addMinusOperators = 12
	grulev3ParserRULE_comparisonOperator = 13
	grulev3ParserRULE_andLogicOperator = 14
	grulev3ParserRULE_orLogicOperator = 15
	grulev3ParserRULE_expressionAtom = 16
	grulev3ParserRULE_constant = 17
	grulev3ParserRULE_variable = 18
	grulev3ParserRULE_arrayMapSelector = 19
	grulev3ParserRULE_memberVariable = 20
	grulev3ParserRULE_functionCall = 21
	grulev3ParserRULE_methodCall = 22
	grulev3ParserRULE_argumentList = 23
	grulev3ParserRULE_floatLiteral = 24
	grulev3ParserRULE_decimalFloatLiteral = 25
	grulev3ParserRULE_hexadecimalFloatLiteral = 26
	grulev3ParserRULE_integerLiteral = 27
	grulev3ParserRULE_decimalLiteral = 28
	grulev3ParserRULE_hexadecimalLiteral = 29
	grulev3ParserRULE_octalLiteral = 30
	grulev3ParserRULE_stringLiteral = 31
	grulev3ParserRULE_booleanLiteral = 32
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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRuleEntryContext)(nil)).Elem())
	var tst = make([]IRuleEntryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRuleEntryContext)
		}
	}

	return tst
}

func (s *GrlContext) RuleEntry(i int) IRuleEntryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRuleEntryContext)(nil)).Elem(), i)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRuleNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRuleNameContext)
}

func (s *RuleEntryContext) LR_BRACE() antlr.TerminalNode {
	return s.GetToken(grulev3ParserLR_BRACE, 0)
}

func (s *RuleEntryContext) WhenScope() IWhenScopeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhenScopeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhenScopeContext)
}

func (s *RuleEntryContext) ThenScope() IThenScopeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IThenScopeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IThenScopeContext)
}

func (s *RuleEntryContext) RR_BRACE() antlr.TerminalNode {
	return s.GetToken(grulev3ParserRR_BRACE, 0)
}

func (s *RuleEntryContext) RuleDescription() IRuleDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRuleDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRuleDescriptionContext)
}

func (s *RuleEntryContext) Salience() ISalienceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISalienceContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerLiteralContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IThenExpressionListContext)(nil)).Elem(), 0)

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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IThenExpressionContext)(nil)).Elem())
	var tst = make([]IThenExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IThenExpressionContext)
		}
	}

	return tst
}

func (s *ThenExpressionListContext) ThenExpression(i int) IThenExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IThenExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IThenExpressionContext)
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
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = _la == grulev3ParserSIMPLENAME {
		{
			p.SetState(100)
			p.ThenExpression()
		}


		p.SetState(103)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *ThenExpressionContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(grulev3ParserSEMICOLON, 0)
}

func (s *ThenExpressionContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *ThenExpressionContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ThenExpressionContext) MethodCall() IMethodCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodCallContext)
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

	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(105)
			p.Assignment()
		}
		{
			p.SetState(106)
			p.Match(grulev3ParserSEMICOLON)
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(108)
			p.FunctionCall()
		}
		{
			p.SetState(109)
			p.Match(grulev3ParserSEMICOLON)
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(111)
			p.variable(0)
		}
		{
			p.SetState(112)
			p.MethodCall()
		}
		{
			p.SetState(113)
			p.Match(grulev3ParserSEMICOLON)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *AssignmentContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

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
		p.SetState(117)
		p.variable(0)
	}
	{
		p.SetState(118)
		_la = p.GetTokenStream().LA(1)

		if !((((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << grulev3ParserASSIGN) | (1 << grulev3ParserPLUS_ASIGN) | (1 << grulev3ParserMINUS_ASIGN) | (1 << grulev3ParserDIV_ASIGN) | (1 << grulev3ParserMUL_ASIGN))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(119)
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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionAtomContext)
}

func (s *ExpressionContext) MulDivOperators() IMulDivOperatorsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMulDivOperatorsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMulDivOperatorsContext)
}

func (s *ExpressionContext) AddMinusOperators() IAddMinusOperatorsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAddMinusOperatorsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAddMinusOperatorsContext)
}

func (s *ExpressionContext) ComparisonOperator() IComparisonOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonOperatorContext)
}

func (s *ExpressionContext) AndLogicOperator() IAndLogicOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAndLogicOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAndLogicOperatorContext)
}

func (s *ExpressionContext) OrLogicOperator() IOrLogicOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrLogicOperatorContext)(nil)).Elem(), 0)

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
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if _la == grulev3ParserNEGATION {
			{
				p.SetState(122)
				p.Match(grulev3ParserNEGATION)
			}

		}
		{
			p.SetState(125)
			p.Match(grulev3ParserLR_BRACKET)
		}
		{
			p.SetState(126)
			p.expression(0)
		}
		{
			p.SetState(127)
			p.Match(grulev3ParserRR_BRACKET)
		}


	case 2:
		{
			p.SetState(129)
			p.expressionAtom(0)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(152)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, grulev3ParserRULE_expression)
				p.SetState(132)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(133)
					p.MulDivOperators()
				}
				{
					p.SetState(134)
					p.expression(8)
				}


			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, grulev3ParserRULE_expression)
				p.SetState(136)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(137)
					p.AddMinusOperators()
				}
				{
					p.SetState(138)
					p.expression(7)
				}


			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, grulev3ParserRULE_expression)
				p.SetState(140)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(141)
					p.ComparisonOperator()
				}
				{
					p.SetState(142)
					p.expression(6)
				}


			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, grulev3ParserRULE_expression)
				p.SetState(144)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(145)
					p.AndLogicOperator()
				}
				{
					p.SetState(146)
					p.expression(5)
				}


			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, grulev3ParserRULE_expression)
				p.SetState(148)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(149)
					p.OrLogicOperator()
				}
				{
					p.SetState(150)
					p.expression(4)
				}

			}

		}
		p.SetState(156)
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
		p.SetState(157)
		_la = p.GetTokenStream().LA(1)

		if !((((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << grulev3ParserDIV) | (1 << grulev3ParserMUL) | (1 << grulev3ParserMOD))) != 0)) {
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
		p.SetState(159)
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
		p.SetState(161)
		_la = p.GetTokenStream().LA(1)

		if !(((((_la - 25)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 25))) & ((1 << (grulev3ParserEQUALS - 25)) | (1 << (grulev3ParserGT - 25)) | (1 << (grulev3ParserLT - 25)) | (1 << (grulev3ParserGTE - 25)) | (1 << (grulev3ParserLTE - 25)) | (1 << (grulev3ParserNOTEQUALS - 25)))) != 0)) {
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
		p.SetState(163)
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
		p.SetState(165)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ExpressionAtomContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *ExpressionAtomContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ExpressionAtomContext) NEGATION() antlr.TerminalNode {
	return s.GetToken(grulev3ParserNEGATION, 0)
}

func (s *ExpressionAtomContext) ExpressionAtom() IExpressionAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionAtomContext)
}

func (s *ExpressionAtomContext) MethodCall() IMethodCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodCallContext)
}

func (s *ExpressionAtomContext) MemberVariable() IMemberVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemberVariableContext)
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
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(168)
			p.Constant()
		}


	case 2:
		{
			p.SetState(169)
			p.FunctionCall()
		}


	case 3:
		{
			p.SetState(170)
			p.variable(0)
		}


	case 4:
		{
			p.SetState(171)
			p.Match(grulev3ParserNEGATION)
		}
		{
			p.SetState(172)
			p.expressionAtom(1)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(179)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionAtomContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, grulev3ParserRULE_expressionAtom)
				p.SetState(175)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(176)
					p.MethodCall()
				}


			case 2:
				localctx = NewExpressionAtomContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, grulev3ParserRULE_expressionAtom)
				p.SetState(177)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(178)
					p.MemberVariable()
				}

			}

		}
		p.SetState(183)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *ConstantContext) IntegerLiteral() IIntegerLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerLiteralContext)
}

func (s *ConstantContext) FloatLiteral() IFloatLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFloatLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFloatLiteralContext)
}

func (s *ConstantContext) BooleanLiteral() IBooleanLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanLiteralContext)(nil)).Elem(), 0)

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

	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(184)
			p.StringLiteral()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(185)
			p.IntegerLiteral()
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(186)
			p.FloatLiteral()
		}


	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(187)
			p.BooleanLiteral()
		}


	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(188)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *VariableContext) MemberVariable() IMemberVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemberVariableContext)
}

func (s *VariableContext) ArrayMapSelector() IArrayMapSelectorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayMapSelectorContext)(nil)).Elem(), 0)

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
		p.SetState(192)
		p.Match(grulev3ParserSIMPLENAME)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(198)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
			case 1:
				localctx = NewVariableContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, grulev3ParserRULE_variable)
				p.SetState(194)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(195)
					p.MemberVariable()
				}


			case 2:
				localctx = NewVariableContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, grulev3ParserRULE_variable)
				p.SetState(196)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(197)
					p.ArrayMapSelector()
				}

			}

		}
		p.SetState(202)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

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
		p.SetState(203)
		p.Match(grulev3ParserLS_BRACKET)
	}
	{
		p.SetState(204)
		p.expression(0)
	}
	{
		p.SetState(205)
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
		p.SetState(207)
		p.Match(grulev3ParserDOT)
	}
	{
		p.SetState(208)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentListContext)(nil)).Elem(), 0)

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
		p.SetState(210)
		p.Match(grulev3ParserSIMPLENAME)
	}
	{
		p.SetState(211)
		p.Match(grulev3ParserLR_BRACKET)
	}
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << grulev3ParserMINUS) | (1 << grulev3ParserLR_BRACKET) | (1 << grulev3ParserTRUE) | (1 << grulev3ParserFALSE) | (1 << grulev3ParserNIL_LITERAL) | (1 << grulev3ParserNEGATION))) != 0) || ((((_la - 38)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 38))) & ((1 << (grulev3ParserSIMPLENAME - 38)) | (1 << (grulev3ParserDQUOTA_STRING - 38)) | (1 << (grulev3ParserSQUOTA_STRING - 38)) | (1 << (grulev3ParserDECIMAL_FLOAT_LIT - 38)) | (1 << (grulev3ParserHEX_FLOAT_LIT - 38)) | (1 << (grulev3ParserDEC_LIT - 38)) | (1 << (grulev3ParserHEX_LIT - 38)) | (1 << (grulev3ParserOCT_LIT - 38)))) != 0) {
		{
			p.SetState(212)
			p.ArgumentList()
		}

	}
	{
		p.SetState(215)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

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
		p.SetState(217)
		p.Match(grulev3ParserDOT)
	}
	{
		p.SetState(218)
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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ArgumentListContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

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
		p.SetState(220)
		p.expression(0)
	}
	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == grulev3ParserT__0 {
		{
			p.SetState(221)
			p.Match(grulev3ParserT__0)
		}
		{
			p.SetState(222)
			p.expression(0)
		}


		p.SetState(227)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecimalFloatLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDecimalFloatLiteralContext)
}

func (s *FloatLiteralContext) HexadecimalFloatLiteral() IHexadecimalFloatLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHexadecimalFloatLiteralContext)(nil)).Elem(), 0)

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

	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(228)
			p.DecimalFloatLiteral()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(229)
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
	p.SetState(233)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == grulev3ParserMINUS {
		{
			p.SetState(232)
			p.Match(grulev3ParserMINUS)
		}

	}
	{
		p.SetState(235)
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
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == grulev3ParserMINUS {
		{
			p.SetState(237)
			p.Match(grulev3ParserMINUS)
		}

	}
	{
		p.SetState(240)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecimalLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDecimalLiteralContext)
}

func (s *IntegerLiteralContext) HexadecimalLiteral() IHexadecimalLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHexadecimalLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHexadecimalLiteralContext)
}

func (s *IntegerLiteralContext) OctalLiteral() IOctalLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOctalLiteralContext)(nil)).Elem(), 0)

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

	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(242)
			p.DecimalLiteral()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(243)
			p.HexadecimalLiteral()
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(244)
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
	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == grulev3ParserMINUS {
		{
			p.SetState(247)
			p.Match(grulev3ParserMINUS)
		}

	}
	{
		p.SetState(250)
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
	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == grulev3ParserMINUS {
		{
			p.SetState(252)
			p.Match(grulev3ParserMINUS)
		}

	}
	{
		p.SetState(255)
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
	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == grulev3ParserMINUS {
		{
			p.SetState(257)
			p.Match(grulev3ParserMINUS)
		}

	}
	{
		p.SetState(260)
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
		p.SetState(262)
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
		p.SetState(264)
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
			if localctx != nil { t = localctx.(*ExpressionContext) }
			return p.Expression_Sempred(t, predIndex)

	case 16:
			var t *ExpressionAtomContext = nil
			if localctx != nil { t = localctx.(*ExpressionAtomContext) }
			return p.ExpressionAtom_Sempred(t, predIndex)

	case 18:
			var t *VariableContext = nil
			if localctx != nil { t = localctx.(*VariableContext) }
			return p.Variable_Sempred(t, predIndex)


	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *grulev3Parser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
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
	switch predIndex {
	case 5:
			return p.Precpred(p.GetParserRuleContext(), 3)

	case 6:
			return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *grulev3Parser) Variable_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 7:
			return p.Precpred(p.GetParserRuleContext(), 2)

	case 8:
			return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

