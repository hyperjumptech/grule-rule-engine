package ast

import "fmt"

func NewGrl() *Grl {
	return &Grl{
		RuleEntries: make(map[string]*RuleEntry, 0),
	}
}

type Grl struct {
	RuleEntries map[string]*RuleEntry
}

type GrlReceiver interface {
	AcceptGrl(grl *Grl) error
}

func (g *Grl) ReceiveRuleEntry(entry *RuleEntry) error {
	if g.RuleEntries == nil {
		g.RuleEntries = make(map[string]*RuleEntry)
	}
	if _, ok := g.RuleEntries[entry.RuleName.SimpleName]; ok {
		return fmt.Errorf("duplicate rule entry %w", entry.RuleName.SimpleName)
	}
	g.RuleEntries[entry.RuleName.SimpleName] = entry
	return nil
}
