[![Build Status](https://travis-ci.org/hyperjumptech/grule-rule-engine.svg?branch=master)](https://travis-ci.org/hyperjumptech/grule-rule-engine)
[![Build Status](https://dev.azure.com/hyperjumptech/grule-rule-engine/_apis/build/status/hyperjumptech.grule-rule-engine?branchName=master)](https://dev.azure.com/hyperjumptech/grule-rule-engine/_build/latest?definitionId=1&branchName=master)

__"Gopher Holds The Rules"__

# Grule

```
import "github.com/hyperjumptech/grule-rule-engine"
```

## Rule Engine for Go


**Grule** is a Rule Engine library for Golang programming language. Inspired by the acclaimed JBOSS Drools, done in a much simple manner.

Like **Drools**, **Grule** have its own *DSL* comparable as follows.

Drools's DRL be like :

```
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

```
rule SpeedUp "When testcar is speeding up we keep increase the speed." salience 10  {
    when
        TestCar.SpeedUp == true && TestCar.Speed < TestCar.MaxSpeed
    then
        TestCar.Speed = TestCar.Speed + TestCar.SpeedIncrement;
        DistanceRecord.TotalDistance = DistanceRecord.TotalDistance + TestCar.Speed;
}
```

# What is RuleEngine?
---

There are no better explanation compared to the article authored by Martin Fowler. You can read the article here ([RulesEngine by Martin Fowler](https://martinfowler.com/bliki/RulesEngine.html)).

Taken from **TutorialsPoint** website (with very slight modification), 

**Grule** is Rule Engine or a Production Rule System that uses the rule-based approach to implement and Expert System. Expert Systems are knowledge-based systems that use knowledge representation to process acquired knowledge into a knowledge base that can be used for reasoning.

A Production Rule System is Turing complete with a focus on knowledge representation to express propositional and first-order logic in a concise, non-ambiguous and declarative manner.

The brain of a Production Rules System is an *Inference Engine* that can scale to a large number of rules and facts. The Inference Engine matches facts and data against Production Rules – also called **Productions** or just **Rules** – to infer conclusions which result in actions.

A Production Rule is a two-part structure that uses first-order logic for reasoning over knowledge representation. A business rule engine is a software system that executes one or more business rules in a runtime production environment.

A Rule Engine allows you to define **“What to Do”** and not **“How to do it.”**

## What is a Rule?

*(also taken from TutorialsPoint)*

Rules are pieces of knowledge often expressed as, "When some conditions occur, then do some tasks."

```
When
   <Condition is true>
Then
   <Take desired Action>
```

The most important part of a Rule is its when part. If the **when** part is satisfied, the **then** part is triggered.

```
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

Rules make it easy to express solutions to difficult problems and get the solutions verified as well. Unlike codes, Rules are written in less complex language; Business Analysts can easily read and verify a set of rules.

### Logic and Data Separation

The data resides in the Domain Objects and the business logic resides in the Rules. Depending upon the kind of project, this kind of separation can be very advantageous.

### Centralization of Knowledge

By using Rules, you create a repository of knowledge (a knowledge base) which is executable. It is a single point of truth for business policy. Ideally, Rules are so readable that they can also serve as documentation.

### Agility To Change

Since business rules are actually treated as data. Adjusting the rule according to business dynamic nature become trivial. No need to re-build codes, deploy as normal software development do, you only need to roll out sets of rule and apply them to knowledge repository.

## Hello Grule

## Grule Rule Language (GRL)

The **GRL**  is a DSL (Domain Specific Language) designed for Grule. Its a simplified language
to be used for defining rule evaluation criteria and action to be executed if the criteria were met.

Generally, the language have the following structure :

```.go
rule <RuleName> <RuleDescription> [salience <priority>] {
    when
        <boolean expression>
    then
        <assignment or operation expression>
}
```

**RuleName** identify a specific rule. The name should be unique in the entire knowledge base, consist of one word thus 
it should not contains white-space. 

**RuleDescription** describes the rule. The description should be enclosed with a double-quote.

**Salience** defines the importance of the rule. Its an optional rule configuration, and by default, when you don't specify them, all rule have the salience of 0 (zero).
The lower the value, the less important the rule. Whenever multiple rule are a candidate for execution, highest salience rule will be executed first. You
may define negative value for the salience, to make the salience even lower. Like any implementation of Rule-Engine,
there are no definitive algorithm to specify which rule to be execute in case of conflicting candidate, the engine may run which ever they like.
Salience is one way of hinting the rule engine of which rule have more importance compared to the other.

**Boolean Expression** is an expression that will be used by rule engine to identify if that specific rule
are a candidate for execution for the current facts.

**Assignment or Operation Expression** contains list of expressions (each expression should be ended with ";" symbol.) 
The expression are designed to modify the current fact values, making calculation, make some logging, etc.

#### Boolean Expression 

Boolean expression comes natural for java or golang developer in GRL. 

```go
when
     contains(User.Name, "robert") &&
     User.Age > 35
then
     ...
```
#### Constants and Literals

| Literal | Description                                                            | Example                          |
| ------- | ---------------------------------------------------------------------- | -------------------------------- |
| String  | Hold string literal, enclosed a string with double quote symbol &quot; | "This is a string"               |
| Decimal | Hold a decimal value, may preceeded with negative symbol -             | `1` or `34` or `42344` or `-553` |
| Real    | Hold a real value                                                      | `234.4553`, `-234.3`             |
| Boolean | Hold a boolean value                                                   | `true`, `TRUE`, `False`          |

Math operator such as `+`, `-`, `/`, `*`; Logical `&&` and `||`; Comparison 
`<`,`<=`,`>`,`>=`,`==`,`!=` all are supported by the language.
#### Comments

You can always put a comment inside your GRL script. Such as :

```go
// This is a comment
// And this

/* And also this */

/*
   As well as this
*/
```

#### Examples

```go
rule SpeedUp "When testcar is speeding up we keep increase the speed."  {
    when
        TestCar.SpeedUp == true && TestCar.Speed < TestCar.MaxSpeed
    then
        TestCar.Speed = TestCar.Speed + TestCar.SpeedIncrement;
		DistanceRecord.TotalDistance = DistanceRecord.TotalDistance + TestCar.Speed;
}

rule StartSpeedDown "When testcar is speeding up and over max speed we change to speed down."  {
    when
        TestCar.SpeedUp == true && TestCar.Speed >= TestCar.MaxSpeed
    then
        TestCar.SpeedUp = false;
		log("Now we slow down");
}

rule SlowDown "When testcar is slowing down we keep decreasing the speed."  {
    when
        TestCar.SpeedUp == false && TestCar.Speed > 0
    then
        TestCar.Speed = TestCar.Speed - TestCar.SpeedIncrement;
		DistanceRecord.TotalDistance = DistanceRecord.TotalDistance + TestCar.Speed;
}

rule SetTime "When Distance Recorder time not set, set it." {
	when
		isNil(DistanceRecord.TestTime)
	then
		log("Set the test time");
		DistanceRecord.TestTime = now();
}
```

## Loading GRL on to Knowledge

One knowledge base may consist of many rules. Tens to hundreds of rules. 
They may be loaded from multiple sources. Those rules will go to the same knowledge
as long as you use the same knowledge when loading each of the resource.

Before you load any rule, you need to have your own knowledge

```go
knowledgeBase := model.NewKnowledgeBase()
ruleBuilder := builder.NewRuleBuilder(knowledgeBase)
```

### From File

```go
fileRes := NewFileResource("/path/to/rules.grl")
err := ruleBuilder.BuildRuleFromResource(fileRes)
if err != nil {
    panic(err)
}
```

### From String or ByteArray

```go
byteArr := NewBytesResource([]byte(rules))
err := ruleBuilder.BuildRuleFromResource(byteArr)
if err != nil {
    panic(err)
}
```

### From URL

```go
urlRes := NewUrlResource("http://host.com/path/to/rule.grl")
err := ruleBuilder.BuildRuleFromResource(urlRes)
if err != nil {
    panic(err)
}
```

## Preparing Facts

In Grule, fact is merely a simple `struct` instance.

Suppose you have the following `struct`.

```go
type TestCar struct {
	SpeedUp        bool
	Speed          int
	MaxSpeed       int
	SpeedIncrement int
}

type DistanceRecorder struct {
	TotalDistance int
	TestTime      time.Time
}
```

And then you instantiate those struct.

```go
tc := &TestCar{
    SpeedUp:        true,
    Speed:          0,
    MaxSpeed:       100,
    SpeedIncrement: 2,
}
dr := &DistanceRecorder{
    TotalDistance: 0,
}
```

Add those struct instances (fact) into data context.

```go
dataContext := context.NewDataContext()
dataContext.Add("TestCar", tc)
dataContext.Add("DistanceRecord", dr)
```

Now your fact is ready to be executed in the rule engine that already prepared with some knowledge.

## Executing A Knowledge On Facts and get result

You already know how to load rules into knowledge base, and you also know how to prepare
your fact. Now we can execute the knowledge base agains facts.

```go
engine := NewGruleEngine()
err = engine.Execute(dataContext, knowledgeBase)
if err != nil {
    t.Errorf("Got error : %v", err)
    t.FailNow()
} else {
    t.Log(dr.TotalDistance)
}
```
The rule engine will use loaded knowledge base to work upon sets of 
fact data in data context. 

## Calling Function in Grule

All invocable functions which are invocable from the DataContext is **Invocable** from within the rule,
both in the "When" scope and "Then" scope.

Assuming you have a struct with some functions.

```go
type MyPoGo struct {
}

func (p *MyPoGo) GetStringLength(sarg string) int {
    return len(sarg)
}

func (p *MyPoGo) AppendString(aString, subString string) string {
    return sprintf("%s%s", aString, subString)
}
```

And add the struct into knowledge base 

```go
dctx := context.NewDataContext()
dctx.Add("Pogo", &MyPoGo{})
```

You can call the fuction within the rule

```go
when
    Pogo.GetStringLength(some.variable) < 100
then
    some.variable = Pogo.AppendString(some.variable, "Groooling");
```

### Default Functions

All functions defined in `model/GruleFunctions.go` are built-in functions.
This means, that you can call the function straight away without having to mention the
instance name as it automatically added by the engine.

I don't maintain list of built in function as you can look into the Go source code directly.
(the `model/GruleFunctions.go`)

### Important Thing you must know about Custom Function in Grule

When you make your own function to be called from rule engine, you need to know the following rules.

1. The function must be visible. Public function convention should start with capital letter. Private functions cant be executed.
2. The function must only return 1 type. Returning multiple variable from function are not acceptable, the rule execution will fail if there are multiple return variable.
3. The way number literal were treated in Grule's DRL is; **decimal** will always taken as `int64` and **real** as `float64`, thus always consider to define your function argument and returns between `int64` and `float64` when your function works with numbers.

# Tasks and Help Wanted.

* Need to do more and more and more test.
* Better code coverage test.
* Better commenting for go doc best practice.
* Improve function argument handling to be more fluid and intuitive.
