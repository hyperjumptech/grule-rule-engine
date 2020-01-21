package ast

import "github.com/sirupsen/logrus"

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
