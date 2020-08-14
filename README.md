
[![Gopheer Holds The Rule](https://github.com/hyperjumptech/grule-rule-engine/blob/master/gopher-grule.png?raw=true)](https://github.com/hyperjumptech/grule-rule-engine/blob/master/gopher-grule.png?raw=true)

[![Build Status](https://travis-ci.org/hyperjumptech/grule-rule-engine.svg?branch=master)](https://travis-ci.org/hyperjumptech/grule-rule-engine)
[![Build Status](https://circleci.com/gh/hyperjumptech/grule-rule-engine.svg?style=svg)](https://circleci.com/gh/hyperjumptech/grule-rule-engine)
[![Go Report Card](https://goreportcard.com/badge/github.com/hyperjumptech/grule-rule-engine)](https://goreportcard.com/report/github.com/hyperjumptech/grule-rule-engine)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

__"Gopher Holds The Rules"__

# Grule

```go
import "github.com/hyperjumptech/grule-rule-engine"
```

## Rule Engine for Go

**Grule** is a Rule Engine library for the Golang programming language. Inspired by the acclaimed JBOSS Drools, done in a much simple manner.

Like **Drools**, **Grule** have its own *DSL* comparable as follows.

Drools's DRL be like :

```drool
rule "SpeedUp"
    salience 10
    when
        $TestCar : TestCarClass( speedUp == true && speed < maxSpeed )
        $DistanceRecord : DistanceRecordClass()
    then
        $TestCar.setSpeed($TestCar.Speed + $TestCar.SpeedIncrement);
        update($TestCar);
        $DistanceRecord.setTotalDistance($DistanceRecord.getTotalDistance() + $TestCar.Speed)
        update($DistanceRecord)
end
```

And Grule's GRL be like :

```go
rule SpeedUp "When testcar is speeding up we keep increase the speed." salience 10  {
    when
        TestCar.SpeedUp == true && TestCar.Speed < TestCar.MaxSpeed
    then
        TestCar.Speed = TestCar.Speed + TestCar.SpeedIncrement;
        DistanceRecord.TotalDistance = DistanceRecord.TotalDistance + TestCar.Speed;
}
```

# What is RuleEngine

There isn't a better explanation than the article authored by Martin Fowler. You can read the article here ([RulesEngine by Martin Fowler](https://martinfowler.com/bliki/RulesEngine.html)).

Taken from **TutorialsPoint** website (with slight modifications),

**Grule** Rule Engine is a Production Rule System that uses the rule-based approach to implement an Expert System. Expert Systems are knowledge-based systems that use knowledge representations to process acquired knowledge into a knowledge base that can be used for reasoning.

A Production Rule System is Turing complete with a focus on knowledge representation to express propositional and first-order logic in a concise, non-ambiguous and declarative manner.

The brain of a Production Rules System is an *Inference Engine* that can scale to a large number of rules and facts. The Inference Engine matches facts and data against Production Rules – also called **Productions** or just **Rules** – to infer conclusions which result in actions.

A Production Rule is a two-part structure that uses first-order logic for reasoning over knowledge representation. A business rule engine is a software system that executes one or more business rules in a runtime production environment.

A Rule Engine allows you to define **“What to Do”** and not **“How to do it.”**

## What is a Rule

*(also taken from TutorialsPoint)*

Rules are pieces of knowledge often expressed as, "When some conditions occur, then do some tasks."

```go
When
   <Condition is true>
Then
   <Take desired Action>
```

The most important part of a Rule is its when part. If the **when** part is satisfied, the **then** part is triggered.

```go
rule  <rule_name> <rule_description>
   <attribute> <value> {
   when
      <conditions>

   then
      <actions>
}
```

## Advantages of a Rule Engine

### Declarative Programming

Rules make it easy to express solutions to difficult problems and get the verifications as well. Unlike codes, Rules are written with less complex language; Business Analysts can easily read and verify a set of rules.

### Logic and Data Separation

The data resides in the Domain Objects and the business logic resides in the Rules. Depending upon the kind of project, this kind of separation can be very advantageous.

### Centralization of Knowledge

By using Rules, you create a repository of knowledge (a knowledge base) which is executable. It is a single point of truth for business policy. Ideally, Rules are so readable that they can also serve as documentation.

### Agility To Change

Since business rules are actually treated as data. Adjusting the rule according to business dynamic nature become trivial. No need to re-build codes, deploy as normal software development do, you only need to roll out sets of rule and apply them to knowledge repository.

### Docs

Grule's Documentation now viewable in ViewDocs. [http://hyperjumptech.viewdocs.io](http://hyperjumptech.viewdocs.io/grule-rule-engine)

### Benchmark
`Loading rules into KnowledgeBase`:

* To load `100` rules into knowledgeBase it took `37549279 ns/op` (took the highest value) that is equal to `37.5ms` and (`8871417 B/op`) `8.8MB` memory per operation

* To load `1000` rules into knowledgeBase it took `211954473 ns/op` (took the highest value) that is equal to `~211ms` and `88MB` memory per operation

`Executing rules against a fact`:

* To execute a fact against 100 rules, Grule Engine took `~39917 ns/op` (took the highest value as base) that is hardly `~0.03917 ms` and `4377 B/op` which is pretty fast.

* To execute a fact against 1000 rules, Grule Engine took `~1420603 ns/op` (took the highest value as base) that is hardly `~1.420603 ms` and `137563B/op` which is also pretty fast.


You can read the [detail report here](docs/Benchmarking_en.md)




# Tasks and Help Wanted

Yes. We need contributors to make Grule even better and useful to the Open Source Community.

* Need to do more and more and more tests.
* Better code coverage test.
* Better commenting for go doc best practice.
* Improve function argument handling to be more fluid and intuitive.

If you really want to help us, simply `Fork` the project and apply for Pull Request.
Please read our [Contribution Manual](CONTRIBUTING.md) and [Code of Conduct](CODE_OF_CONDUCTS.md)