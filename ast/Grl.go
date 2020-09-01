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
	if _, ok := g.RuleEntries[entry.RuleName.SimpleName]; ok {
		return fmt.Errorf("duplicate rule entry %s", entry.RuleName.SimpleName)
	}
	g.RuleEntries[entry.RuleName.SimpleName] = entry
	return nil
}
