package events

import "fmt"

const (
	// RuleEntryAddedEvent an event when Rule Entry get added into knowledge base
	RuleEntryAddedEvent int = iota
	// RuleEntryRemovedEvent an event when Rule Entry get removed from knowledge base
	RuleEntryRemovedEvent
	// RuleEntryRetractedEvent an event when Rule Entry get retracted from rule engine next cycle execution
	RuleEntryRetractedEvent
	// RuleEntryResetEvent an event when Rule Entry get restored to next rule engine cycle execution
	RuleEntryResetEvent
	// RuleEntryExecuteStartEvent an event when Rule Entry about to be executed
	RuleEntryExecuteStartEvent
	// RuleEntryExecuteEndEvent an event when Rule Entry finish execution
	RuleEntryExecuteEndEvent
	// RuleEngineStartEvent an event when Rule Engine is about to start
	RuleEngineStartEvent
	// RuleEngineEndEvent an event when Rule Engine is just finished
	RuleEngineEndEvent
	// RuleEngineCycleEvent an event when Rule Engine start a cycle
	RuleEngineCycleEvent
	// RuleEngineErrorEvent an event when Rule Engine encounter an error in it's execution
	RuleEngineErrorEvent

	// RuleEntryEventTopic the event topic for rule entry
	RuleEntryEventTopic = "RuleEntryEvent"
	// RuleEngineEventTopic an event topic for rule engine
	RuleEngineEventTopic = "RuleEngineEvent"
)

// RuleEntryEvent is an event data to be sent for RuleEntryEvent topic
type RuleEntryEvent struct {
	fmt.Stringer
	// Name of the rule entry involved in this event.
	RuleName string
	//The event type to further categorized what is actually happening
	EventType int
}

// RuleEngineEvent is an event data to be sent for RuleEngineEvent topic
type RuleEngineEvent struct {
	fmt.Stringer
	//The event type to further categorized what is actually happening
	EventType int
	// The current execution cycle when this event is emitted
	Cycle uint64
	// An error, if this event is emitted because of an error
	Error error
}
