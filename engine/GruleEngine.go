package engine

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/events"
	"github.com/hyperjumptech/grule-rule-engine/pkg/eventbus"
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
	RuleEnginePublisher := eventbus.DefaultBrooker.GetPublisher(events.RuleEngineEventTopic)
	RuleEntryPublisher := eventbus.DefaultBrooker.GetPublisher(events.RuleEntryEventTopic)

	var contextError error
	var contextCanceled bool

	go func() {
		select {
		case <-ctx.Done():
			contextError = ctx.Err()
			contextCanceled = true
		}
	}()

	// emit engine start event
	RuleEnginePublisher.Publish(&events.RuleEngineEvent{
		EventType: events.RuleEngineStartEvent,
		Cycle:     0,
	})

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
			return contextError
		}

		// add the cycle counter
		cycle++

		// emit engine cycle event
		RuleEnginePublisher.Publish(&events.RuleEngineEvent{
			EventType: events.RuleEngineCycleEvent,
			Cycle:     cycle,
		})

		log.Debugf("Cycle #%d", cycle)
		// if cycle is above the maximum allowed cycle, returnan error indicated the cycle has ended.
		if cycle > g.MaxCycle {

			// create the error
			err := fmt.Errorf("the GruleEngine successfully selected rule candidate for execution after %d cycles, this could possibly caused by rule entry(s) that keep added into execution pool but when executed it does not change any data in context. Please evaluate your rule entries \"When\" and \"Then\" scope. You can adjust the maximum cycle using GruleEngine.MaxCycle variable", g.MaxCycle)

			// emit engine error event
			RuleEnginePublisher.Publish(&events.RuleEngineEvent{
				EventType: events.RuleEngineErrorEvent,
				Cycle:     cycle,
				Error:     err,
			})

			return err
		}

		// Select all rule entry that can be executed.
		log.Tracef("Select all rule entry that can be executed.")
		runnable := make([]*ast.RuleEntry, 0)
		for _, v := range knowledge.RuleEntries {
			// test if this rule entry v can execute.
			can, err := v.Evaluate()
			if err != nil {
				log.Errorf("Failed testing condition for rule : %s. Got error %v", v.Name, err)
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
				dataCtx.ResetVariableChangeCount()
				log.Debugf("Executing rule : %s. Salience %d", r.Name, r.Salience)

				// emit rule execute start event
				RuleEntryPublisher.Publish(&events.RuleEntryEvent{
					EventType: events.RuleEntryExecuteStartEvent,
					RuleName:  r.Name,
				})

				err := r.Execute()
				if err != nil {
					log.Errorf("Failed execution rule : %s. Got error %v", r.Name, err)
					return err
				}

				// emit rule execute end event
				RuleEntryPublisher.Publish(&events.RuleEntryEvent{
					EventType: events.RuleEntryExecuteEndEvent,
					RuleName:  r.Name,
				})

				if dataCtx.IsComplete() {
					cycleDone = true
					break
				}

				//if there is a variable change, restart the cycle.
				if dataCtx.HasVariableChange() {
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
	log.Debugf("Finished Rules execution. With knowledge base '%s' version %s. Total #%d cycles. Duration %d ms.", knowledge.Name, knowledge.Version, cycle, time.Now().Sub(startTime).Nanoseconds()/1e6)

	// emit engine finish event
	RuleEnginePublisher.Publish(&events.RuleEngineEvent{
		EventType: events.RuleEngineEndEvent,
		Cycle:     cycle,
	})

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

	runnable := make([]*ast.RuleEntry, 0)

	//Loop through all the rule entries available in the knowledge base and add to the response list if it is able to evaluate
	for _, v := range knowledge.RuleEntries {
		// test if this rule entry v can execute.
		can, err := v.Evaluate()
		if err != nil {
			log.Errorf("Failed testing condition for rule : %s. Got error %v", v.Name, err)
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
			return runnable[i].Salience > runnable[j].Salience
		})
	}
	return runnable, nil
}
