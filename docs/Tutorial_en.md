# Grule Short Tutorial

[Tutorial](Tutorial_en.md) | [Rule Engine](RuleEngine_en.md) | [GRL](GRL_en.md) | [RETE Algorithm](RETE_en.md) | [Functions](Function_en.md) | [FAQ](FAQ_en.md)

## Preparation

Please note that Grule is using Go 1.13

To import Grule into your project you can simply import it.

```Shell
$go get github.com/hyperjumptech/grule-rule-engine
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

A `fact` in grule is a mere **pointer** to an instance of a `struct`.
This struct may contains properties like any normal Golang struct, as well
as `function` or usually, in *OOP*, we call them `method`

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

Just like normal Golang conventions, only **visible** attributes can be accessible
from the Grule rule engine, those with capital first letter.

Grule can also call Fact's functions.

```go
func (mf *MyFact) GetWhatToSay(sentence string) string {
    return fmt.Sprintf("Let say \"%s\"", sentence)
}
```

Please note, that there are some conventions.

* Member function must be **Visible**, have capital first letter.
* If the function have returns, there must be only 1 return.
* Arguments and return, if it is an `int`, `uint` or `float`, must be their 64-bit variant. Eg. `int64`, `uint64`, `float64`.
* The function are not encouraged to change the Fact's member attribute, this cause RETE's working memory detection impossible.
If you **MUST** change some variable, you have to notify Grule using `Changed(varname string)` built-in function.

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

You can create as many fact as you wish.

Next, you have to prepare a `DataContext` and add your instance(s) of facts into it.

```go
dataCtx := ast.NewDataContext()
err := dataCtx.Add("MF", myFact)
if err != nil {
    panic(err)
}
```

## Creating KnowledgeLibrary and Add Rules Into It

A `KnowledgeLibrary` is basically collection of `KnowledgeBase` blue prints. 
And `KnowledgeBase` is a collection of many rules sourced from rule definitions
loaded from multiple sources.
We use `RuleBuilder` to build `KnowledgeBase` and add it into `KnowledgeLibrary`

The DRL, can be in the form of simple string, stored in a file or somewhere from the internet can each be
used to build those rules.

Now lets create the `KnowledgeLibrary` and `RuleBuilder` to build the rule into prepared `KnowledgeLibrary`

```go
knowledgeLibrary := ast.NewKnowledgeLibrary()
ruleBuilder := builder.NewRuleBuilder(knowledgeLibrary)
```

Now we can add rules (defined within a GRL)

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

// Add the rule definition above into the library and name it 'TutorialRules'  version '0.0.1'
byteArr := pkg.NewBytesResource([]byte(drls))
err := ruleBuilder.BuildRuleFromResource("TutorialRules", "0.0.1", byteArr)
if err != nil {
    panic(err)
}
```

### Resources

You can always load a GRL from multiple sources.

#### From File

```go
fileRes := pkg.NewFileResource("/path/to/rules.grl")
err := ruleBuilder.BuildRuleFromResource("TutorialRules", "0.0.1", fileRes)
if err != nil {
    panic(err)
}
```

or if you want to get file resource by their pattern

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

#### From String or ByteArray

```go
byteArr := pkg.NewBytesResource([]byte(rules))
err := ruleBuilder.BuildRuleFromResource("TutorialRules", "0.0.1", byteArr)
if err != nil {
    panic(err)
}
```

#### From URL

```go
urlRes := pkg.NewUrlResource("http://host.com/path/to/rule.grl")
err := ruleBuilder.BuildRuleFromResource("TutorialRules", "0.0.1", urlRes)
if err != nil {
    panic(err)
}
```

#### From GIT

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

#### From JSON

You can now build rules from JSON!,  [Read how it works](GRL_JSON_en.md) 

Now, in the `KnowledgeLibrary` we have a `KnowledgeBase` named `TutorialRules` with version `0.0.1`. To execute this particular rule, you have to obtain an instance of it from the `KnowledgeLibrary`. This will be explained on the next section.

## Executing Grule Rule Engine

To execute a KnowledgeBase, we need to get an instance of this `KnowledgeBase` from `KnowledgeLibrary` 

```go
knowledgeBase := knowledgeLibrary.NewKnowledgeBaseInstance("TutorialRules", "0.0.1")
```

Each instance you obtained from knowledgeLibrary is a *clone* from the underlying `KnowledgeBase` *blue-print*. Its entirely different instance that makes it *thread-safe* for execution. Each *instance* also carries its own `WorkingMemory`. This is very useful when you want to have a multithreaded execution of rule engine (eg. In a web-server to serve each request using a rule).

Its a significant performance improvement boost since you don't have to "recreate" a `KnowledgeBase` from GRL everytime you start another thread. The `KnowledgeLibrary` will clone the `AST` structure of `KnowledgeBase` into a new instance.

Ok, now lets execute the `KnowledgeBase` instance using the prepared `DataContext`.

```go
engine = engine.NewGruleEngine()
err = engine.Execute(dataCtx, knowledgeBase)
if err != nil {
    panic(err)
}
```

## Obtaining Result

If you see, on the rule's GRL above,

```go
rule CheckValues "Check the default values" salience 10 {
    when 
        MF.IntAttribute == 123 && MF.StringAttribute == "Some string value"
    then
        MF.WhatToSay = MF.GetWhatToSay("Hello Grule");
        Retract("CheckValues");
}
```

The rule modify attribute `MF.WhatToSay` where this is revering to the `WhatToSay` in the
`MyFact` struct. So simply access that member variable to get the result.

```go
fmt.Println(myFact.WhatToSay)
// this should prints
// Lets Say "Hello Grule"
```