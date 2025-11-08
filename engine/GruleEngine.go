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

package engine

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/rs/zerolog"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/logger"
)

const (
	DefaultCycleCount = 5000
)

var (
	// logFields default fields for grule
	logFields = logger.Fields{
		"package": "engine",
	}

	// Logger is a logger instance with default fields for grule
	log = logger.Log.WithFields(logFields)
)

// SetLogger changes default logger on external
func SetLogger(externalLog interface{}) {
	var entry logger.LogEntry

	switch externalLog.(type) {
	case *zap.Logger:
		log, ok := externalLog.(*zap.Logger)
		if !ok {

			return
		}
		entry = logger.NewZap(log)
	case *logrus.Logger:
		log, ok := externalLog.(*logrus.Logger)
		if !ok {

			return
		}
		entry = logger.NewLogrus(log)
	case *zerolog.Logger:
		log, ok := externalLog.(*zerolog.Logger)
		if !ok {
			return
		}
		entry = logger.NewZero(log)
	default:

		return
	}

	log = entry.WithFields(logFields)
}

// NewGruleEngine will create new instance of GruleEngine struct.
// It will set the max cycle to 5000
func NewGruleEngine() *GruleEngine {

	return &GruleEngine{
		MaxCycle: DefaultCycleCount,
	}
}

// GruleEngine is the engine structure. It has the Execute method to start the engine to work.
type GruleEngine struct {
	MaxCycle                        uint64
	ReturnErrOnFailedRuleEvaluation bool
	CompareNilValues                bool
	Listeners                       []GruleEngineListener
}

// Execute function is the same as ExecuteWithContext(context.Background())
func (g *GruleEngine) Execute(dataCtx ast.IDataContext, knowledge *ast.KnowledgeBase) error {

	return g.ExecuteWithContext(context.Background(), dataCtx, knowledge)
}

// notifyEvaluateRuleEntry will notify all registered listener that a rule is being evaluated.
func (g *GruleEngine) notifyEvaluateRuleEntry(ctx context.Context, cycle uint64, entry *ast.RuleEntry, candidate bool) {
	if g.Listeners != nil && len(g.Listeners) > 0 {
		for _, gl := range g.Listeners {
			gl.EvaluateRuleEntry(ctx, cycle, entry, candidate)
		}
	}
}

// notifyEvaluateRuleEntry will notify all registered listener that a rule is being executed.
func (g *GruleEngine) notifyExecuteRuleEntry(ctx context.Context, cycle uint64, entry *ast.RuleEntry) {
	if g.Listeners != nil && len(g.Listeners) > 0 {
		for _, gl := range g.Listeners {
			gl.ExecuteRuleEntry(ctx, cycle, entry)
		}
	}
}

// notifyEvaluateRuleEntry will notify all registered listener that a rule is being executed.
func (g *GruleEngine) notifyBeginCycle(ctx context.Context, cycle uint64) {
	if g.Listeners != nil && len(g.Listeners) > 0 {
		for _, gl := range g.Listeners {
			gl.BeginCycle(ctx, cycle)
		}
	}
}

// ExecuteWithContext function will execute a knowledge evaluation and action against data context.
// The engine will evaluate context cancelation status in each cycle.
// The engine also do conflict resolution of which rule to execute.
func (g *GruleEngine) ExecuteWithContext(ctx context.Context, dataCtx ast.IDataContext, knowledge *ast.KnowledgeBase) error {
	if knowledge == nil || dataCtx == nil {

		return fmt.Errorf("nil KnowledgeBase or DataContext is not allowed")
	}

	log.Debugf("Starting rule execution using knowledge '%s' version %s. Contains %d rule entries", knowledge.Name, knowledge.Version, len(knowledge.RuleEntries))

	// Prepare the timer, we need to measure the processing time in debug mode.
	startTime := time.Now()

	// Prepare the build-in function and add to datacontext.
	defunc := &ast.BuiltInFunctions{
		Knowledge:     knowledge,
		WorkingMemory: knowledge.WorkingMemory,
		DataContext:   dataCtx,
	}
	err := dataCtx.Add("DEFUNC", defunc)
	if err != nil {
		log.Error("DEFUNC add err")

		return err
	}

	err = dataCtx.Add("COMPARE_NILS", g.CompareNilValues)
	if err != nil {
		log.Error("COMPARE_NILS add err")

		return err
	}

	// Working memory need to be resetted. all Expression will be set as not evaluated.
	log.Debugf("Resetting Working memory")
	knowledge.WorkingMemory.ResetAll()
	knowledge.Reset()

	// Initialize all AST with datacontext and working memory
	log.Debugf("Initializing Context")
	knowledge.InitializeContext(dataCtx)

	var cycle uint64

	/*
		Un-limited loop as long as there are rule to execute.
		We need to add safety mechanism to detect unlimited loop as there are possibility executed rule are not changing
		data context which makes rules to get executed again and again.
	*/
	for {
		if ctx.Err() != nil {
			log.Error("Context canceled")

			return ctx.Err()
		}

		g.notifyBeginCycle(ctx, cycle+1)

		// Select all rule entry that can be executed.
		log.Tracef("Select all rule entry that can be executed.")
		runnable := make([]*ast.RuleEntry, 0)
		for _, ruleEntry := range knowledge.RuleEntries {
			if ctx.Err() != nil {
				log.Error("Context canceled")

				return ctx.Err()
			}
			if !ruleEntry.Retracted && !ruleEntry.Deleted {
				// test if this rule entry v can execute.
				can, err := ruleEntry.Evaluate(ctx, dataCtx, knowledge.WorkingMemory)
				if err != nil {
					log.Errorf("Failed testing condition for rule : %s. Got error %v", ruleEntry.RuleName, err)
					if g.ReturnErrOnFailedRuleEvaluation {

						return err
					}
				}
				// if can, add into runnable array
				if can {
					runnable = append(runnable, ruleEntry)
				}
				// notify all listeners that a rule's when scope is been evaluated.
				g.notifyEvaluateRuleEntry(ctx, cycle+1, ruleEntry, can)
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

			runner := runnable[0]

			// scan all runnables and pick the highest salience
			if len(runnable) > 1 {
				for idx, pr := range runnable {
					if idx > 0 && runner.Salience < pr.Salience {
						runner = pr
					}
				}
			}
			// set the current rule entry to run. This is for trace ability purpose
			dataCtx.SetRuleEntry(runner)
			// notify listeners that we are about to execute a rule entry then scope
			g.notifyExecuteRuleEntry(ctx, cycle, runner)
			// execute the top most prioritized rule
			err := runner.Execute(ctx, dataCtx, knowledge.WorkingMemory)
			if err != nil {
				log.Errorf("Failed execution rule : %s. Got error %v", runner.RuleName, err)

				return fmt.Errorf("error while executing rule %s. got %w", runner.RuleName, err)
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
	if knowledge == nil || dataCtx == nil {

		return nil, fmt.Errorf("nil KnowledgeBase or DataContext is not allowed")
	}

	log.Debugf("Starting rule matching using knowledge '%s' version %s. Contains %d rule entries", knowledge.Name, knowledge.Version, len(knowledge.RuleEntries))
	// Prepare the build-in function and add to datacontext.
	defunc := &ast.BuiltInFunctions{
		Knowledge:     knowledge,
		WorkingMemory: knowledge.WorkingMemory,
		DataContext:   dataCtx,
	}
	err := dataCtx.Add("DEFUNC", defunc)
	if err != nil {
		log.Error("DEFUNC add err")

		return nil, err
	}

	err = dataCtx.Add("COMPARE_NILS", g.CompareNilValues)
	if err != nil {
		log.Error("COMPARE_NILS add err")

		return nil, err
	}
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
	for _, entries := range knowledge.RuleEntries {
		if !entries.Deleted {
			// test if this rule entry v can execute.
			can, err := entries.Evaluate(context.Background(), dataCtx, knowledge.WorkingMemory)
			if err != nil {
				log.Errorf("Failed testing condition for rule : %s. Got error %v", entries.RuleName, err)
				if g.ReturnErrOnFailedRuleEvaluation {
					return nil, err
				}
			}
			// if can, add into runnable array
			if can {
				runnable = append(runnable, entries)
			}
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
