# Grule's RETE Algorithm

[Tutorial](Tutorial_en.md) | [Rule Engine](RuleEngine_en.md) | [GRL](GRL_en.md) | [RETE Algorithm](RETE_en.md) | [Functions](Function_en.md) | [FAQ](FAQ_en.md)

From Wikipedia : The Rete algorithm (/ˈriːtiː/ REE-tee, /ˈreɪtiː/ RAY-tee, rarely /ˈriːt/ REET, /rɛˈteɪ/ reh-TAY) is a pattern matching algorithm for implementing rule-based systems. The algorithm was developed to efficiently apply many rules or patterns to many objects, or facts, in a knowledge base. It is used to determine which of the system's rules should fire based on its data store, its facts.

Some form of the RETE algorithm was implemented in `grule-rule-engine` starting from version `1.1.0`.
It replaces the __Naive__ approach when evaluating rules to add to `ConflictSet`.

`ExpressionAtom` in the DRL are compiled and will not be duplicated within the working memory of the engine.
This increased the engine performance significantly if you have many rules defined with lots of duplicated expressions
or lots of heavy function/method calls.

Grule's RETE implementation don't have `Class` selector as one expression may involve multiple class. For example an expression such as:

```.go
when
    ClassA.attr == ClassB.attr + ClassC.AFunc()
then
    ...
```

The expression above involve attribute/function result comparison and math operation from 3 different class. This makes
RETE's class separation of expression token difficult.

You can read about RETE algorithm here:

* https://en.wikipedia.org/wiki/Rete_algorithm
* https://www.drdobbs.com/architecture-and-design/the-rete-matching-algorithm/184405218
* https://www.sparklinglogic.com/rete-algorithm-demystified-part-2/ 

### Why Rete Algorithm is necessary

Suppose we have a fact.

```go
type Fact struct {
    StringValue string
}

func (f *Fact) VeryHeavyAndLongFunction() bool {
    ...
}
```

And add the fact to data contest

```go
f := &Fact{}
dctx := context.NewDataContext()
err := dctx.Add("Fact", f)
```

And we have DRL like ...

```go
rule ... {
    when
        Fact.VeryHeavyAndLongFunction() && Fact.StringValue == "Fish"
    then
        ...
}
rule ... {
    when
        Fact.VeryHeavyAndLongFunction() && Fact.StringValue == "Bird"
    then
        ...
}
rule ... {
    when
        Fact.VeryHeavyAndLongFunction() && Fact.StringValue == "Mammal"
    then
        ...
}
...
// and alot more of simillar rule
...
rule ... {
    when
        Fact.VeryHeavyAndLongFunction() && Fact.StringValue == "Insect"
    then
        ...
}
```

Executing the DRL above might "kill" the engine because when it tries to choose what rules to execute,
the engine will call the `Fact.VeryHeavyAndLongFunction` function in every rule's `when` scope.

Thus, instead of executing the `Fact.VeryHeavyAndLongFunction` while evaluating each
rule, Rete algorithm only evaluate them once (one the first encounter with the function), and remember the result
for the rest of the rules.

The same with `Fact.StringValue`. Rete algorithm will load the value from the object instance and
remember it. Until it got changed in the `then` scope, as in ...

```go
rule ... {
    when
        ...
    then
        Fact.StringValue = "something else";
}
```

### What is inside Grule's Working-Memory

Grule will try to remember all of the `Expression` defined within rule's `when` scope of all rules
in the KnowledgeBase.

First, It will try its best to make sure none of the AST (Abstract Syntax Tree) node get duplicated.

Second, each of this AST node can only be evaluated once, until it's relevant `variable` get changed. For example :

Boolean Expression :

```Shell
    when
    Fact.A == Fact.B + Fact.Func(Fact.C) - 20
```

This expression will be broken down into the following Expressions.

```Shell
Expression "Fact.A" --> A variable
Expression "Fact.B" --> A variable
Expression "Fact.C" --> A variable
Expression "Fact.Func(Fact.C)"" --> A function containing argument Fact.C
Expression "20" --> A constant
Expression "Fact.B + Fact.Func(Fact.C)" --> A math operation contains 2 variable; Fact.B and Fact.C
Expression "(Fact.B + Fact.Func(Fact.C))" - 20 -- A math operation also contains 2 variable.
```

Each of the above Expressions will be remembered for their underlaying values whenever
they get evaluated for the first time. So subsequent evaluation will not be evaluated
as their remembered value will immediately returned.

If one of this Variable got altered inside the rule's `then` scope, for example

```Shell
    then
        Fact.B = Fact.A * 20
```

We can see `Fact.B` value is changed, then all Expression containing `Fact.B` will
be removed from Working memory:

```Shell
Expression "Fact.B"
Expression "Fact.B + Fact.Func(Fact.C)" --> A math operation contains 2 variable; Fact.B and Fact.C
Expression "(Fact.B + Fact.Func(Fact.C))" - 20 -- A math operation also contains 2 variable. 
```

This makes those Expression removed from the working memory and get re-evaluated again on the next cycle.

### Known RETE issue with Functions or Methods

While Grule will try to remember any variable it evaluate within the `when` and `then` scope, if you change
the variable value from outside the rule engine, for example changed from within a function call,
Grule won't be able to see this change, thus Grule may mistakenly evaluate a variable which already changed.

Consider the following fact:

```go
type Fact struct {
    StringValue string
}

func (f *Fact) SetStringValue(newValue string) {
    f.StringValue = newValue
}
```

Then you instantiate your fact and add it into data context

```go
f := &Fact{
    StringValue: "One",
}
dctx := context.NewDataContext()
err := dctx.Add("Fact", f)
```

In your GRL you did something like this

```go
rule one "One" {
    when
        Fact.StringValue == "One"
        // here grule remembers that Fact.StringValue value is "One"
    then
        Fact.SetStringValue("Two");
        // here grule does not know if Fact.StringValue has changed inside the function.
        // What grule know is Fact.StringValue is still "One"
}

rule two "Two" {
    when
        Fact.StringValue == "Two"
        // Because of that, this will never evaluated true.
    then
        Fact.SetStringValue("Three");
}
```

Thus the engine will finish without error, but the expected result, where `Fact.StringValue` should be `Two`
is not met.

To overcome this, you should tell grule if the variable has changed using `Changed` function.

```go
rule one "One" {
    when 
        Fact.StringValue == "One"
        // here grule remember that Fact.StringValue value is "One"
    then
        Fact.SetStringValue("Two");
        // here grule does not know if Fact.StringValue has changed inside the function.
        // What grule know is Fact.StringValue is still "One"

        // We should tell Grule that the variable changed within the Fact
        Changed("Fact.StringValue")
}
```
