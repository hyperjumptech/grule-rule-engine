# Grule Short Tutorial

---

:construction:
__THIS PAGE IS BEING TRANSLATED__
:construction:

:construction_worker: Contributors are invited. Please read [CONTRIBUTING](../../CONTRIBUTING.md) and [CONTRIBUTING TRANSLATION](../CONTRIBUTING_TRANSLATION.md) guidelines.

:vulcan_salute: Please remove this note once you're done translating.

---


[![Tutorial_cn](https://github.com/yammadev/flag-icons/blob/master/png/CN.png?raw=true)](../cn/Tutorial_cn.md)
[![Tutorial_de](https://github.com/yammadev/flag-icons/blob/master/png/DE.png?raw=true)](../de/Tutorial_de.md)
[![Tutorial_en](https://github.com/yammadev/flag-icons/blob/master/png/GB.png?raw=true)](../en/Tutorial_en.md)
[![Tutorial_id](https://github.com/yammadev/flag-icons/blob/master/png/ID.png?raw=true)](../id/Tutorial_id.md)
[![Tutorial_pl](https://github.com/yammadev/flag-icons/blob/master/png/PL.png?raw=true)](../pl/Tutorial_pl.md)

[About](About_de.md) | [Tutorial](Tutorial_de.md) | [Rule Engine](RuleEngine_de.md) | [GRL](GRL_de.md) | [GRL JSON](GRL_JSON_de.md) | [RETE Algorithm](RETE_de.md) | [Functions](Function_de.md) | [FAQ](FAQ_de.md) | [Benchmark](Benchmarking_de.md)

---

## Preparation

Please note that Grule is using Go 1.13.

To import Grule into your project:

```Shell
$ go get github.com/hyperjumptech/grule-rule-engine
```

From your `go` you can import Grule.

```go
import (
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
) 
``` 

## Creating Fact Structure

A `fact` in grule is a **pointer** to an instance of a `struct`.  The struct
may also contain properties just as any normal Golang `struct`, including any
`method` you wish to define, provided it adheres to the requirements for
methods defined below.  As an example:

```go
type MyFact struct {
    IntAttribute       int64
    StringAttribute    string
    BooleanAttribute   bool
    FloatAttribute     float64
    TimeAttribute      time.Time
    WhatToSay          string
}
```

As with normal Golang conventions, Grule is only able to access those
**visible** attributes and methods exposed with an initial capital letter.

```go
func (mf *MyFact) GetWhatToSay(sentence string) string {
    return fmt.Sprintf("Let say \"%s\"", sentence)
}
```

**NOTE:** Member functions are subject to the following requirements:

* The member function must be **visible**; it's name must start with a capital
  letter.
* The member function must return `0` or `1` values. More than one return value
  is not supported.
* All numerical argument and return types must be their 64 bit variant. i.e.
  `int64`, `uint64`, `float64`.
* The member function **should not** change the Fact's internal state. The
  algorithm cannot automatically detect these changes, things become more
  difficult to reason about, and bugs can creep in.  If you **MUST** change
  some internal state of the Fact, then you can notify Grule using
  `Changed(varname string)` built-in function.

## Add Fact Into DataContext

To add a fact into `DataContext` you have to create an instance of your `fact`

```go
myFact := &MyFact{
    IntAttribute: 123,
    StringAttribute: "Some string value",
    BooleanAttribute: true,
    FloatAttribute: 1.234,
    TimeAttribute: time.Now(),
}
```

You can create as many facts as you wish.

After the fact(s) have been created, you can then add those instances into the
`DataContext`:

```go
dataCtx := ast.NewDataContext()
err := dataCtx.Add("MF", myFact)
if err != nil {
    panic(err)
}
```

### Creating a Fact from JSON

JSON data can also be used to describe facts in Grule as of version 1.8.0.  For
more detail, see [JSON as a Fact](../en/JSON_Fact_en.md).

## Creating a KnowledgeLibrary and Adding Rules Into It

A `KnowledgeLibrary` is a collection of `KnowledgeBase` blue prints and a
`KnowledgeBase` is a collection of many rules sourced from rule definitions
loaded from multiple sources.  We use `RuleBuilder` to build `KnowledgeBase`
instances and then add them to the `KnowledgeLibrary`.

The source form of a GRL can be:

* a raw string
* contents of a file
* a document at an HTTP endpoint

Lets use the `RuleBuilder` to start populating our `KnowledgeLibrary`.

```go
knowledgeLibrary := ast.NewKnowledgeLibrary()
ruleBuilder := builder.NewRuleBuilder(knowledgeLibrary)
```

Next we can define a basic rule as a raw string in the DSL:

```go
// lets prepare a rule definition
drls := `
rule CheckValues "Check the default values" salience 10 {
    when 
        MF.IntAttribute == 123 && MF.StringAttribute == "Some string value"
    then
        MF.WhatToSay = MF.GetWhatToSay("Hello Grule");
        Retract("CheckValues");
}
`
```

And finally we can use the builder to add the definition to the
`knowledgeLibrary` from a declared `resource`:

```go
// Add the rule definition above into the library and name it 'TutorialRules'  version '0.0.1'
bs := pkg.NewBytesResource([]byte(drls))
err := ruleBuilder.BuildRuleFromResource("TutorialRules", "0.0.1", bs)
if err != nil {
    panic(err)
}
```

The `KnowledgeLibrary` now contains a `KnowledgeBase` named `TutorialRules`
with version `0.0.1`. To execute this particular rule we must obtain an
instance from the `KnowledgeLibrary`. This will be explained on the next
section.

## Executing Grule Rule Engine

To execute a KnowledgeBase, we need to get an instance of this `KnowledgeBase`
from `KnowledgeLibrary` 

```go
knowledgeBase := knowledgeLibrary.NewKnowledgeBaseInstance("TutorialRules", "0.0.1")
```

Each instance you obtain from the `knowledgeLibrary` is a unique *clone* from
the underlying `KnowledgeBase` *blueprint*.  Each unique instance also carries
its own distinct `WorkingMemory`. As no instance shares any state with any
other instance, you are free to use them in any multithreaded environment
provided you aren't executing any single instance from multiple threads
simultaneously.

Constructing from the `KnowledgeBase` blueprint also ensures that we aren't
recomputing work every time we want to construct an instance.  The
computational work is only done once, making the work of cloning the `AST`
extremely efficient.

Now lets execute the `KnowledgeBase` instance using the prepared `DataContext`.

```go
engine = engine.NewGruleEngine()
err = engine.Execute(dataCtx, knowledgeBase)
if err != nil {
    panic(err)
}
```

## Obtaining Result

Here's the rule we defined above, just for reference:

```go
rule CheckValues "Check the default values" salience 10 {
    when 
        MF.IntAttribute == 123 && MF.StringAttribute == "Some string value"
    then
        MF.WhatToSay = MF.GetWhatToSay("Hello Grule");
        Retract("CheckValues");
}
```

Assuming the condition is matched (which it is) the action will modify the
`MF.WhatToSay` attribute.  In order to ensure that the rule is not then
immediately re-evaluted, the rule is `Retract`ed from the set.  In this
particular instance, if the rule failed to do this then it would match again on
the next cycle, and again, and again.  Eventually Grule would terminate with an
error, since it would be unable to converge on a terminal result.

In this case, all you have to do in order to obtain the result is just examine
your `myFact` instance for the modification your rule made:

```go
fmt.Println(myFact.WhatToSay)
// this should prints
// Lets Say "Hello Grule"
```
## Resources

GRLs can be stored in external files and there are many ways to obtain and load
the contents of those files.

### From File

```go
fileRes := pkg.NewFileResource("/path/to/rules.grl")
err := ruleBuilder.BuildRuleFromResource("TutorialRules", "0.0.1", fileRes)
if err != nil {
    panic(err)
}
```

You can also load multiple files into a bundle with paths and glob patterns:

```go
bundle := pkg.NewFileResourceBundle("/path/to/grls", "/path/to/grls/**/*.grl")
resources := bundle.MustLoad()
for _, res := range resources {
    err := ruleBuilder.BuildRuleFromResource("TutorialRules", "0.0.1", res)
    if err != nil {
        panic(err)
    }
}
```

### From String or ByteArray

```go
bs := pkg.NewBytesResource([]byte(rules))
err := ruleBuilder.BuildRuleFromResource("TutorialRules", "0.0.1", bs)
if err != nil {
    panic(err)
}
```

### From URL

```go
urlRes := pkg.NewUrlResource("http://host.com/path/to/rule.grl")
err := ruleBuilder.BuildRuleFromResource("TutorialRules", "0.0.1", urlRes)
if err != nil {
    panic(err)
}
```

#### With Headers

```go
headers := make(http.Header)
headers.Set("Authorization", "Basic YWxhZGRpbjpvcGVuc2VzYW1l")
urlRes := pkg.NewURLResourceWithHeaders("http://host.com/path/to/rule.grl", headers)
err := ruleBuilder.BuildRuleFromResource("TutorialRules", "0.0.1", urlRes)
if err != nil {
    panic(err)
}
```

### From GIT

```go
bundle := pkg.NewGITResourceBundle("https://github.com/hyperjumptech/grule-rule-engine.git", "/**/*.grl")
resources := bundle.MustLoad()
for _, res := range resources {
    err := ruleBuilder.BuildRuleFromResource("TutorialRules", "0.0.1", res)
    if err != nil {
        panic(err)
    }
}
```

### From JSON

You can now build rules from JSON! [Read how it works](GRL_JSON_de.md) 

## Compile GRL into GRB

If you want to have faster rule set loading performance (e.g. you have very
large rule sets and loading GRL is too slow), you can save your rule set
into GRB (Grules Rule Binary) file. [Read how to store and load GRB](Binary_Rule_File_de.md) 