//  Copyright hyperjumptech/grule-rule-engine Authors
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package ast

import (
	"github.com/hyperjumptech/grule-rule-engine/logger"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

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
	// astLogFields default fields for grule
	astLogFields = logger.Fields{
		"package": "ast",
	}

	// AstLog is a logger instance twith default fields for grule
	AstLog = logger.Log.WithFields(astLogFields)
)

// SetLogger changes default logger on external
func SetLogger(log interface{}) {
	var entry logger.LogEntry

	switch log.(type) {
	case *zap.Logger:
		log, ok := log.(*zap.Logger)
		if !ok {
			return
		}
		entry = logger.NewZap(log)
	case *logrus.Logger:
		log, ok := log.(*logrus.Logger)
		if !ok {
			return
		}
		entry = logger.NewLogrus(log)
	default:
		return
	}

	AstLog = entry.WithFields(astLogFields)
	GrlLogger = entry.WithFields(grlLoggerFields)
}

// Node defines interface to implement by all AST node models
type Node interface {
	GetAstID() string
	GetGrlText() string
	GetSnapshot() string
	SetGrlText(grlText string)
	MakeCatalog(cat *Catalog)
}
