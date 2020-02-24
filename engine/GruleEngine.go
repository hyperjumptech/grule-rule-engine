package engine

import (
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/events"
	"github.com/hyperjumptech/grule-rule-engine/pkg/eventbus"
	"github.com/juju/errors"
	"github.com/sirupsen/logrus"
	"sort"
	"time"
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

// Execute function will execute a knowledge evaluation and action against data context.
// The engine also do conflict resolution of which rule to execute.
func (g *GruleEngine) Execute(dataCtx *ast.DataContext, knowledge *ast.KnowledgeBase, memory *ast.WorkingMemory) error {
	RuleEnginePublisher := eventbus.DefaultBrooker.GetPublisher(events.RuleEngineEventTopic)
	RuleEntryPublisher := eventbus.DefaultBrooker.GetPublisher(events.RuleEntryEventTopic)

	// emit engine start event
	RuleEnginePublisher.Publish(&events.RuleEngineEvent{
		EventType: events.RuleEngineStartEvent,
		Cycle:     0,
	})

	log.Debugf("Starting rule execution using knowledge '%s' version %s. Contains %d rule entries", knowledge.Name, knowledge.Version, len(knowledge.RuleEntries))

	knowledge.WorkingMemory = memory

	// Prepare the timer, we need to measure the processing time in debug mode.
	startTime := time.Now()

	// Prepare the build-in function and add to datacontext.
	defunc := &ast.BuildInFunctions{
		Knowledge:     knowledge,
		WorkingMemory: memory,
	}
	dataCtx.Add("DEFUNC", defunc)

	// Working memory need to be resetted. all Expression will be set as not evaluated.
	log.Debugf("Resetting Working memory")
	knowledge.WorkingMemory.ResetAll()

	// Initialize all AST with datacontext and working memory
	log.Debugf("Initializing Context")
	knowledge.InitializeContext(dataCtx, memory)

	var cycle uint64

	/*
		Un-limitted loop as long as there are rule to execute.
		We need to add safety mechanism to detect unlimitted loop as there are posibility executed rule are not changing
		data context which makes rules to get executed again and again.
	*/
	for {

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
			err := errors.Errorf("GruleEngine successfully selected rule candidate for execution after %d cycles, this could possibly caused by rule entry(s) that keep added into execution pool but when executed it does not change any data in context. Please evaluate your rule entries \"When\" and \"Then\" scope. You can adjust the maximum cycle using GruleEngine.MaxCycle variable.", g.MaxCycle)

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
				dataCtx.VariableChangeCount = 0
				log.Debugf("Executing rule : %s. Salience %d", r.Name, r.Salience)

				// emit rule execute start event
				RuleEntryPublisher.Publish(&events.RuleEntryEvent{
					EventType: events.RuleEntryExecuteStartEvent,
					RuleName:  r.Name,
				})

				err := r.Execute()
				if err != nil {
					log.Errorf("Failed execution rule : %s. Got error %v", r.Name, err)
					return errors.Trace(err)
				}

				// emit rule execute end event
				RuleEntryPublisher.Publish(&events.RuleEntryEvent{
					EventType: events.RuleEntryExecuteEndEvent,
					RuleName:  r.Name,
				})

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
	log.Debugf("Finished Rules execution. With knowledge base '%s' version %s. Total #%d cycles. Duration %d ms.", knowledge.Name, knowledge.Version, cycle, time.Now().Sub(startTime).Milliseconds())

	// emit engine finish event
	RuleEnginePublisher.Publish(&events.RuleEngineEvent{
		EventType: events.RuleEngineEndEvent,
		Cycle:     cycle,
	})

	return nil
}
