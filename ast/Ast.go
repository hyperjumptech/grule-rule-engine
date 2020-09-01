package ast

import "github.com/sirupsen/logrus"

const (
	// ARGUMENTLIST signature for argument list snapshot
	ARGUMENTLIST = "AL"
	// MAPARRAYSELECTOR signature for map array snapshot
	MAPARRAYSELECTOR = "MAS"
	// ASSIGMENT signature for assignment snapshot
	ASSIGMENT = "AS"
	// CONSTANT signature for constant snapshot
	CONSTANT = "C"
	// EXPRESSION signature for expression snapshot
	EXPRESSION = "E"
	// EXPRESSIONATOM signature for expression atom snapshot
	EXPRESSIONATOM = "A"
	// FUNCTIONCALL signature for function call snapshot
	FUNCTIONCALL = "F"
	// RULEENTRY signature for rule entry snapshot
	RULEENTRY = "R"
	// THENEXPRESSION signature for then expression snapshot
	THENEXPRESSION = "TE"
	// THENEXPRESSIONLIST signature for then expression list snapshot
	THENEXPRESSIONLIST = "TEL"
	// THENSCOPE signature for then scope snapshot
	THENSCOPE = "TS"
	// WHENSCOPE signature for when scope snapshot
	WHENSCOPE = "WS"
	// VARIABLE signature for variable snapshot
	VARIABLE = "V"
)

var (
	// AstLog is a logrus instance twith default fields for grule
	AstLog = logrus.WithFields(logrus.Fields{
		"lib":     "grule",
		"package": "AST",
	})
)

// Node defines interface to implement by all AST node models
type Node interface {
	GetAstID() string
	GetGrlText() string
	GetSnapshot() string
	SetGrlText(grlText string)
}
