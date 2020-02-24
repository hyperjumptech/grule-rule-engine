# Grule Short Tutorial

[Tutorial](Tutorial_en.md) | [Rule Engine](RuleEngine_en.md) | [GRL](GRL_en.md) | [RETE Algorithm](RETE_en.md) | [Functions](Function_en.md) | [Grule Events](GruleEvent_en.md)

## Preparation

Please note that Grule is using Go 1.13

To import Grule into your project you can simply import it.

```text
$ go get github.com/hyperjumptech/grule-rule-engine
```

From your `go` you can import Grule.

```go
import grule "github.com/hyperjumptech/grule-rule-engine"
``` 

## Creating Fact Structure

A `fact` in grule is a mere **pointer** to an instance of a `struct`.
This struct may contains properties just like any normal Golang struct, as well
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

Just like normal Golang convention, only **visible** attribute can be accessible
from Grule rule engine. those with capital first letter.

Grule can also call Fact's functions.

```go
func (mf *MyFact) GetWhatToSay(sentence string) string {
    return fmt.Sprintf("Let say \"%s\"", sentence)
}
```

Please note, that there are some convention though.

* Member function must be **Visible**, have capital first letter.
* If the function have returns, there must be only 1 return.
* Arguments and return, if its an `int`, `uint` or `float`, must be on their 64 bit variant. Eg. `int64`, `uint64`, `float64`.
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

Now, you have prepare a `DataContext` and add your instance(s) of fact into it.

```go
dataCtx := grule.ast.NewDataContext()
err := dataCtx.Add("MF", myFact)
if err != nil {
    panic(err)
}
```

## Creating KnowledgeBase and Adding Rules

A `KnowledgeBase` is basically collection of many rules sourced from rule definitions
loaded from multiple sources.

The DRL, can be in the form of simple string, stored on a file or some where on the internet, are
used to build those rules.

Now lets create the `KnowledgeBase`, `WorkingMemory` and then create new `RuleBuilder` to build the rule into prepared `KnowledgeBase`

```go
workingMemory := grule.ast.NewWorkingMemory()
knowledgeBase := grule.ast.NewKnowledgeBase("tutorial", "1.0.0")
ruleBuilder := grule.builder.NewRuleBuilder(knowledgeBase, workingMemory)
```

Now we can add rules (defined within a GRL)

```go
drls := `
rule CheckValues "Check the default values" salience 10 {
    when 
        MF.IntAttribute == 123 && MF.StringAttribute == "Some string value"
    then
        MF.WhatToSay = MF.GetWhatToSay("Hello Grule");
        Retract("CheckValues);
}
`
byteArr := grule.pkg.NewBytesResource([]byte(drls))
err := ruleBuilder.BuildRuleFromResource(byteArr)
if err != nil {
    panic(err)
}
```

### Resources

You can always load a GRL from multiple sources.


#### From File

```go
fileRes := grule.pkg.NewFileResource("/path/to/rules.grl")
err := ruleBuilder.BuildRuleFromResource(fileRes)
if err != nil {
    panic(err)
}
```

or if you want to get file resource by their pattern

```go
bundle := grule.pkg.NewFileResourceBundle("/path/to/grls", "/path/to/grls/**/*.grl")
resources := bundle.MustLoad()
for _, res := range resources {
    err := ruleBuilder.BuildRuleFromResource(res)
    if err != nil {
        panic(err)
    }
}
```

#### From String or ByteArray

```go
byteArr := grule.pkg.NewBytesResource([]byte(rules))
err := ruleBuilder.BuildRuleFromResource(byteArr)
if err != nil {
    panic(err)
}
```

#### From URL

```go
urlRes := grule.pkg.NewUrlResource("http://host.com/path/to/rule.grl")
err := ruleBuilder.BuildRuleFromResource(urlRes)
if err != nil {
    panic(err)
}
```

#### From GIT

```go
bundle := grule.pkg.NewGITResourceBundle("https://github.com/hyperjumptech/grule-rule-engine.git", "/**/*.grl")
resources := bundle.MustLoad()
for _, res := range resources {
    err := ruleBuilder.BuildRuleFromResource(res)
    if err != nil {
        panic(err)
    }
}
```

## Executing Grule Rule Engine

To execute the rules, we need to create an instance of `GruleEngine` and with it,
we execute evaluate our `KnowledgeBase` upon the facts in `DataContext`

```go
engine = grule.engine.NewGruleEngine()
err = engine.Execute(dataCtx, knowledgeBase, workingMemory)
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
        Retract("CheckValues);
}
```

The rule modify attribute `MF.WhatToSay` where this is revering to the `WhatToSay` in the 
`MyFact` struct. So simply access that member variable to get the result.

```go
fmt.Println(myFact.WhatToSay)
// this should prints
// Lets Say "Hello Grule"
```
