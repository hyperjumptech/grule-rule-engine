package ast

func NewGrl() *Grl {
	return &Grl{
		RuleEntries: make([]*RuleEntry, 0),
	}
}

type Grl struct {
	RuleEntries []*RuleEntry
}

type GrlReceiver interface {
	AcceptGrl(grl *Grl) error
}

func (g *Grl) ReceiveRuleEntry(entry *RuleEntry) error {
	if g.RuleEntries == nil {
		g.RuleEntries = make([]*RuleEntry, 0)
	}
	g.RuleEntries = append(g.RuleEntries, entry)
	return nil
}
