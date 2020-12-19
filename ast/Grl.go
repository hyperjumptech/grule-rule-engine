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

import "fmt"

// NewGrl creates new GRL instance
func NewGrl() *Grl {
	return &Grl{
		RuleEntries: make(map[string]*RuleEntry, 0),
	}
}

// Grl will contains multiple RuleEntries
type Grl struct {
	RuleEntries map[string]*RuleEntry
}

// GrlReceiver is interface for objects that should hold a GRL, will be called by ANTLR walker.
type GrlReceiver interface {
	AcceptGrl(grl *Grl) error
}

// ReceiveRuleEntry will make this GRL to accept rule entries created by ANTLR walker
func (g *Grl) ReceiveRuleEntry(entry *RuleEntry) error {
	if g.RuleEntries == nil {
		g.RuleEntries = make(map[string]*RuleEntry)
	}
	if _, ok := g.RuleEntries[entry.RuleName]; ok {
		return fmt.Errorf("duplicate rule entry %s", entry.RuleName)
	}
	g.RuleEntries[entry.RuleName] = entry
	return nil
}
