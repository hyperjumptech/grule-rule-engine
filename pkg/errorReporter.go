//  Copyright kalyan-arepalle/grule-rule-engine Authors
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

package pkg

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// GruleErrorReporter is an implementation of ErrorListener interface by antlr. The purpose is to capture errors during lexer tokenization and parsing.
type GruleErrorReporter struct {
	*antlr.DefaultErrorListener // Embed default which ensures we fit the interface
	Errors                      []error
}

// AddError simply add an error into this reporter
func (c *GruleErrorReporter) AddError(err error) {
	c.Errors = append(c.Errors, err)
}

// SyntaxError call back which will be called upon parsing error
func (c *GruleErrorReporter) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	c.Errors = append(c.Errors, fmt.Errorf("grl error on %d:%d %s", line, column, msg))
}

// HasError check if this reporter has an error
func (c *GruleErrorReporter) HasError() bool {
	return c.Errors != nil && len(c.Errors) > 0
}

// Error return an error text. This function is there for compatibility reason.
func (c *GruleErrorReporter) Error() string {
	if c.HasError() {
		return fmt.Sprintf("got %d error(s) in grl the script", len(c.Errors))
	}
	return fmt.Sprintf("no error in grl script")
}
