# Listening to Grule Engine Events

[Tutorial](Tutorial_en.md) | [Rule Engine](RuleEngine_en.md) | [GRL](GRL_en.md) | [RETE Algorithm](RETE_en.md) | [Functions](Function_en.md) | [Grule Events](GruleEvent_en.md)

Sometime you want to know what happening when Rule engine is running. 
You may want to know when a certain rule is being executed, when the engine
is started, or ended, or to simply know what is the current execution cycle.

Grule is running an internal *EventBus* in it. All you have to do is to obtain
a *Subscriber* to its *Broker* by specifying function to capture the event.

```go

import (
    eventbus "github.com/hyperjumptech/grule-rule-engine/pkg/eventbus"
    events   "github.com/hyperjumptech/grule-rule-engine/events"
)

...
...


	ruleEntrySubscriber := eventbus.DefaultBrooker.GetSubscriber(events.RuleEntryEventTopic, func(i interface{}) error {
		if i != nil && getTypeOf(i) == "*RuleEntryEvent" {
			event := i.(*events.RuleEntryEvent)
			if event.EventType == events.RuleEntryExecuteStartEvent {
				log.Infof("Rule executed %s", event.RuleName)
			}
		} else if i != nil {
			log.Infof("RuleEntry Subscriber, Receive type is %s ", getTypeOf(i))
		}
		return nil
	})
	ruleEntrySubscriber.Subscribe()
```

Another example

```go
	ruleEngineSubscriber := eventbus.DefaultBrooker.GetSubscriber(events.RuleEngineEventTopic, func(i interface{}) error {
		if i != nil && getTypeOf(i) == "*RuleEngineEvent" {
			event := i.(*events.RuleEngineEvent)
			if event.EventType == events.RuleEngineEndEvent {
				log.Infof("Engine finished in %d cycles", event.Cycle)
			}
		} else if i != nil {
			log.Infof("RuleEngine Subscriber, Receive type is %s ", getTypeOf(i))
		}
		return nil
	})
	ruleEngineSubscriber.Subscribe()
```

Like a common *Messaging* infrastructure, bule is publishing *messages* on a certain topic.

Those topic are:

```go
package events
...
const (
	// RuleEntryEventTopic the event topic for rule entry
	RuleEntryEventTopic = "RuleEntryEvent"

	// RuleEngineEventTopic an event topic for rule engine
	RuleEngineEventTopic = "RuleEngineEvent"
)
``` 

The following are the struct get sents...

```go
package events
...

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
```

The following is list of `EventType` that explains further about the event.

```go
package events
...
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
)
```