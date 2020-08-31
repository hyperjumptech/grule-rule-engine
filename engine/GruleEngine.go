package engine

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/sirupsen/logrus"
)

var (
	// Logger is a logrus instance with default fields for grule
	log = logrus.WithFields(logrus.Fields{
		"lib":    "grule",
		"struct": "GruleEngineV2",
	})
)

// NewGruleEngine will create new instance of GruleEngine struct.
// It will set the max cycle to 5000
func NewGruleEngine() *GruleEngine {
	return &GruleEngine{
		MaxCycle: 5000,
	}
}

// GruleEngine is the engine structure. It has the Execute method to start the engine to work.
type GruleEngine struct {
	MaxCycle uint64
}

// Execute function is the same as ExecuteWithContext(context.Background())
func (g *GruleEngine) Execute(dataCtx ast.IDataContext, knowledge *ast.KnowledgeBase) error {
	return g.ExecuteWithContext(context.Background(), dataCtx, knowledge)
}

// ExecuteWithContext function will execute a knowledge evaluation and action against data context.
// The engine will evaluate context cancelation status in each cycle.
// The engine also do conflict resolution of which rule to execute.
func (g *GruleEngine) ExecuteWithContext(ctx context.Context, dataCtx ast.IDataContext, knowledge *ast.KnowledgeBase) error {
	var contextError error
	var contextCanceled bool

	go func() {
		select {
		case <-ctx.Done():
			contextError = ctx.Err()
			contextCanceled = true
		}
	}()

	log.Debugf("Starting rule execution using knowledge '%s' version %s. Contains %d rule entries", knowledge.Name, knowledge.Version, len(knowledge.RuleEntries))

	// Prepare the timer, we need to measure the processing time in debug mode.
	startTime := time.Now()

	// Prepare the build-in function and add to datacontext.
	defunc := &ast.BuiltInFunctions{
		Knowledge:     knowledge,
		WorkingMemory: knowledge.WorkingMemory,
		DataContext:   dataCtx,
	}
	dataCtx.Add("DEFUNC", defunc)

	// Working memory need to be resetted. all Expression will be set as not evaluated.
	log.Debugf("Resetting Working memory")
	knowledge.WorkingMemory.ResetAll()

	// Initialize all AST with datacontext and working memory
	log.Debugf("Initializing Context")
	knowledge.InitializeContext(dataCtx)

	var cycle uint64

	/*
		Un-limitted loop as long as there are rule to execute.
		We need to add safety mechanism to detect unlimitted loop as there are posibility executed rule are not changing
		data context which makes rules to get executed again and again.
	*/
	for {
		if contextCanceled {
			log.Error("Context canceled")
			return contextError
		}

		// Select all rule entry that can be executed.
		log.Tracef("Select all rule entry that can be executed.")
		runnable := make([]*ast.RuleEntry, 0)
		for _, v := range knowledge.RuleEntries {
			// test if this rule entry v can execute.
			can, err := v.Evaluate(dataCtx, knowledge.WorkingMemory)
			if err != nil {
				log.Errorf("Failed testing condition for rule : %s. Got error %v", v.RuleName.SimpleName, err)
				// No longer return error, since unavailability of variable or fact in context might be intentional.
			}
			// if can, add into runnable array
			if can {
				runnable = append(runnable, v)
			}
		}

		// disabled to test the rete's variable change detection.
		// knowledge.RuleContextReset()
		log.Tracef("Selected rules %d.", len(runnable))

		// If there are rules to execute, sort them by their Salience
		if len(runnable) > 0 {
			// add the cycle counter
			cycle++

			log.Debugf("Cycle #%d", cycle)
			// if cycle is above the maximum allowed cycle, returnan error indicated the cycle has ended.
			if cycle > g.MaxCycle {
				log.Error("Max cycle reached")
				return fmt.Errorf("the GruleEngine successfully selected rule candidate for execution after %d cycles, this could possibly caused by rule entry(s) that keep added into execution pool but when executed it does not change any data in context. Please evaluate your rule entries \"When\" and \"Then\" scope. You can adjust the maximum cycle using GruleEngine.MaxCycle variable", g.MaxCycle)
			}

			// execute the top most prioritized rule
			runner := runnable[0]

			// scan all runnables and pick the highest salience
			if len(runnable) > 1 {
				for idx, pr := range runnable {
					if idx > 0 && runner.Salience.SalienceValue < pr.Salience.SalienceValue {
						runner = pr
					}
				}
			}
			// execute the top most prioritized rule
			err := runner.Execute(dataCtx, knowledge.WorkingMemory)
			if err != nil {
				log.Errorf("Failed execution rule : %s. Got error %v", runner.RuleName.SimpleName, err)
				return err
			}

			if dataCtx.IsComplete() {
				break
			}
		} else {
			// No more rule can be executed, so we are done here.
			log.Debugf("No more rule to run")
			break
		}
	}
	log.Debugf("Finished Rules execution. With knowledge base '%s' version %s. Total #%d cycles. Duration %d ms.", knowledge.Name, knowledge.Version, cycle, time.Now().Sub(startTime).Nanoseconds()/1e6)
	return nil
}

// FetchMatchingRules function is responsible to fetch all the rules that matches to a fact against all rule entries
// Returns []*ast.RuleEntry order by salience
func (g *GruleEngine) FetchMatchingRules(dataCtx ast.IDataContext, knowledge *ast.KnowledgeBase) ([]*ast.RuleEntry, error) {
	log.Debugf("Starting rule matching using knowledge '%s' version %s. Contains %d rule entries", knowledge.Name, knowledge.Version, len(knowledge.RuleEntries))
	// Prepare the build-in function and add to datacontext.
	defunc := &ast.BuiltInFunctions{
		Knowledge:     knowledge,
		WorkingMemory: knowledge.WorkingMemory,
		DataContext:   dataCtx,
	}
	dataCtx.Add("DEFUNC", defunc)

	// Working memory need to be resetted. all Expression will be set as not evaluated.
	log.Debugf("Resetting Working memory")
	knowledge.WorkingMemory.ResetAll()
	// Initialize all AST with datacontext and working memory
	log.Debugf("Initializing Context")
	knowledge.InitializeContext(dataCtx)

	//Loop through all the rule entries available in the knowledge base and add to the response list if it is able to evaluate
	// Select all rule entry that can be executed.
	log.Tracef("Select all rule entry that can be executed.")
	runnable := make([]*ast.RuleEntry, 0)
	for _, v := range knowledge.RuleEntries {
		// test if this rule entry v can execute.
		can, err := v.Evaluate(dataCtx, knowledge.WorkingMemory)
		if err != nil {
			log.Errorf("Failed testing condition for rule : %s. Got error %v", v.RuleName.SimpleName, err)
			// No longer return error, since unavailability of variable or fact in context might be intentional.
		}
		// if can, add into runnable array
		if can {
			runnable = append(runnable, v)
		}
	}

	log.Debugf("Matching rules length %d.", len(runnable))
	if len(runnable) > 1 {
		sort.SliceStable(runnable, func(i, j int) bool {
			return runnable[i].Salience.SalienceValue > runnable[j].Salience.SalienceValue
		})
	}
	return runnable, nil
}
