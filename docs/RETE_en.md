# Grule's RETE Algorithm

[Tutorial](Tutorial_en.md) | [Rule Engine](RuleEngine_en.md) | [GRL](GRL_en.md) | [GRL JSON](GRL_JSON_en.md) | [RETE Algorithm](RETE_en.md) | [Functions](Function_en.md) | [FAQ](FAQ_en.md) | [Benchmark](Benchmarking_en.md)

---

From Wikipedia : The Rete algorithm (/ˈriːtiː/ REE-tee, /ˈreɪtiː/ RAY-tee, rarely /ˈriːt/ REET, /rɛˈteɪ/ reh-TAY) is a pattern matching algorithm for implementing rule-based systems. The algorithm was developed to efficiently apply many rules or patterns to many objects, or facts, in a knowledge base. It is used to determine which of the system's rules should fire based on its data store, its facts.

Some form of the RETE algorithm was implemented in `grule-rule-engine` starting from version `1.1.0`.
It replaces the __Naive__ approach when evaluating rules to add to `ConflictSet`.

The `ExpressionAtom` elements in the GRL are compiled and will not be duplicated within the working memory of the engine.
This increases the engine performance significantly when you have many rules defined with many duplicated expressions
or many heavy function/method calls.

Grule's RETE implementation does not have a `Class` selector, as one expression may involve multiple classes. For example an, expression such as:

```.go
when
    ClassA.attr == ClassB.attr + ClassC.AFunc()
then
    ...
```

The expression above involves attribute and function call result comparisons and math operations from 3 different classes. This makes
RETE's class separation of expression tokens difficult.

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

And we add the fact to the data context:

```go
f := &Fact{}
dctx := context.NewDataContext()
err := dctx.Add("Fact", f)
```

And we have GRL like:

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

// and many similar rules

rule ... {
    when
        Fact.VeryHeavyAndLongFunction() && Fact.StringValue == "Insect"
    then
        ...
}
```

Executing the GRL above might "kill" the engine because, when it tries to choose which rules to execute,
the engine will call the `Fact.VeryHeavyAndLongFunction` function in every rule's `when` scope.

Thus, instead of executing the `Fact.VeryHeavyAndLongFunction` while evaluating each
rule, the Rete algorithm only evaluates them once (when the function call is first encountered), and it then remembers the result
for the rest of the rules. (**Note** that this means your function call *must be referentially transparent* -- i.e. it must have no side effects)

The same with `Fact.StringValue`. The Rete algorithm will load the value from the object instance and
remember it until it gets changed in a `then` scope, such as:

```go
rule ... {
    when
        ...
    then
        Fact.StringValue = "something else";
}
```

### What is inside Grule's Working-Memory

Grule will try to remember all of the `Expression` elements defined within a rule's `when` scope of all rules
in the KnowledgeBase.

First, it will try its best to make sure that none of the AST (Abstract Syntax Tree) nodes are duplicated.

Second, each of these AST nodes can be evaluated only once, until it's relevant `variable` gets changed. For example:

```Shell
    when
        Fact.A == Fact.B + Fact.Func(Fact.C) - 20
```

This condition will be broken down into the following `Expression`s.

```Shell
Expression "Fact.A" --> A variable
Expression "Fact.B" --> A variable
Expression "Fact.C" --> A variable
Expression "Fact.Func(Fact.C)"" --> A function containing argument Fact.C
Expression "20" --> A constant
Expression "Fact.B + Fact.Func(Fact.C)" --> A math operation contains 2 variable; Fact.B and Fact.C
Expression "(Fact.B + Fact.Func(Fact.C))" - 20 -- A math operation also contains 2 variable.
```

The resulting values for each of the above `Expression`s will be remembered (memoized) upon their first invocation so that subsequent references to them will avoid a re-invocation of them, returning the remembered value immediately instead.

If one of these values is altered inside the rule's `then` scope, for example...

```Shell
    then
        Fact.B = Fact.A * 20
```

... then all Expression containing `Fact.B` will be removed from Working memory:

```Shell
Expression "Fact.B"
Expression "Fact.B + Fact.Func(Fact.C)" --> A math operation contains 2 variable; Fact.B and Fact.C
Expression "(Fact.B + Fact.Func(Fact.C))" - 20 -- A math operation also contains 2 variable. 
```

Those `Expression`s will be removed from the working memory so that they get re-evaluated on the next cycle.

### Known RETE issue with Functions or Methods

While Grule will try to remember any variable it evaluates within the `when`
and `then` scope, if you change the variable value from outside the rule
engine, for example changed from within a function call, Grule won't be able to
see this change. As a result, Grule may mistakenly use the old (memoized) value
for the variable, since it doesn't know that the value has changed.  You should
endeavour to ensure your functions are **referentially transparent** in order
to never have to deal with this issue.

Consider the following fact:

```go
type Fact struct {
    StringValue string
}

func (f *Fact) SetStringValue(newValue string) {
    f.StringValue = newValue
}
```

Then you instantiate your fact and add it into data context:

```go
f := &Fact{
    StringValue: "One",
}
dctx := context.NewDataContext()
err := dctx.Add("Fact", f)
```

In your GRL you then do something like this

```go
rule one "One" {
    when
        Fact.StringValue == "One"
        // Here grule remembers that Fact.StringValue value is "One"
    then
        Fact.SetStringValue("Two");
        // Here grule does not know that Fact.StringValue has changed inside the function.
        // What grule know is Fact.StringValue is still "One".
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
