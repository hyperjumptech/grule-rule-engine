package engine

import (
	"github.com/hyperjumptech/grule-rule-engine/context"
	"github.com/hyperjumptech/grule-rule-engine/model"
	"github.com/juju/errors"
	"github.com/sirupsen/logrus"
	"sort"
	"time"
)

var (
	// Logger is a logrus instance twith default fields for grule
	log = logrus.WithFields(logrus.Fields{
		"lib":    "grule",
		"struct": "GruleEngine",
	})
)

// NewGruleEngine will create new instance of GruleEngine struct.
// It will set the max cycle to 5000
func NewGruleEngine() *GruleEngine {
	return &GruleEngine{
		MaxCycle:    5000,
		subscribers: make([]func(*model.RuleEntry), 0),
	}
}

// GruleEngine is the engine structure. It has the Execute method to start the engine to work.
type GruleEngine struct {
	MaxCycle    uint64
	subscribers []func(*model.RuleEntry)
}

// Subscribe adds custom func to subscribers slice
func (g *GruleEngine) Subscribe(f func(*model.RuleEntry)) {
	g.subscribers = append(g.subscribers, f)
}

// Notify all subscribers
func (g *GruleEngine) notifySubscribers(r *model.RuleEntry) {
	for _, f := range g.subscribers {
		go f(r)
	}
}

// Execute function will execute a knowledge evaluation and action against data context.
// The engine also do conflict resolution of which rule to execute.
func (g *GruleEngine) Execute(dataCtx *context.DataContext, knowledge *model.KnowledgeBase) error {
	log.Infof("Starting rule execution using knowledge '%s' version %s", knowledge.Name, knowledge.Version)
	startTime := time.Now()
	defunc := &model.GruleFunction{
		Knowledge: knowledge,
	}
	kctx := &context.KnowledgeContext{}
	rctx := &context.RuleContext{}
	dataCtx.Add("DEFUNC", defunc)

	knowledge.Reset()

	for _, v := range knowledge.RuleEntries {
		v.Initialize(kctx, rctx, dataCtx)
	}

	var cycle uint64

	/*
		Un-limitted loop as long as there are rule to exsecute.
		We need to add safety mechanism to detect unlimitted loop as there are posibility executed rule are not changing
		data context which makes rules to get executed again and again.
	*/
	for {
		cycle++

		if cycle > g.MaxCycle {
			return errors.Errorf("GruleEngine successfully selected rule candidate for execution after %d cycles, this could possibly caused by rule entry(s) that keep added into execution pool but when executed it does not change any data in context. Please evaluate your rule entries \"When\" and \"Then\" scope. You can adjust the maximum cycle using GruleEngine.MaxCycle variable.", g.MaxCycle)
		}

		// Select all rule entry that can be executed.
		runnable := make([]*model.RuleEntry, 0)
		for _, v := range knowledge.RuleEntries {
			// test if this rule entry v can execute.
			can, err := v.CanExecute()
			if err != nil {
				log.Errorf("Failed testing condition for rule : %s. Got error %v", v.RuleName, err)
				// No longer return error, since unavailability of variable or fact in context might be intentional.
			}
			// if can, add into runnable array
			if can {
				runnable = append(runnable, v)
			}
		}

		// If there are rules to execute, sort them by their Salience
		if len(runnable) > 0 {
			if len(runnable) > 1 {
				sort.SliceStable(runnable, func(i, j int) bool {
					return runnable[i].Salience > runnable[j].Salience
				})
			}
			// Start rule execution cycle.
			// We assume that none of the runnable rule will change variable so we set it to true.
			cycleDone := true

			for _, r := range runnable {
				// reset the counter to 0 to detect if there are variable change.
				dataCtx.VariableChangeCount = 0
				log.Debugf("Executing rule : %s. Salience %d", r.RuleName, r.Salience)
				err := r.Execute()
				if err != nil {
					log.Errorf("Failed execution rule : %s. Got error %v", r.RuleName, err)
					return errors.Trace(err)
				}

				// notify subscribers about executed rule
				g.notifySubscribers(r)

				//if there is a variable change, restart the cycle.
				if dataCtx.VariableChangeCount > 0 {
					cycleDone = false
					break
				}
				// this point means no variable change, so we move to the next rule entry.
			}
			// if cycleDone is true, we are done.
			if cycleDone {
				break
			}
		} else {
			// No more rule can be executed, so we are done here.
			break
		}
	}
	log.Infof("Finished Rules execution. With knowledge base '%s' version %s. Total #%d cycles. Duration %d microsec.", knowledge.Name, knowledge.Version, cycle, time.Now().Sub(startTime))
	return nil
}
