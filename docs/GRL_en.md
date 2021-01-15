# Grule Rule Language (GRL)

[Tutorial](Tutorial_en.md) | [Rule Engine](RuleEngine_en.md) | [GRL](GRL_en.md) | [GRL JSON](GRL_JSON_en.md) | [RETE Algorithm](RETE_en.md) | [Functions](Function_en.md) | [FAQ](FAQ_en.md) | [Benchmark](Benchmarking_en.md)

---

The **GRL** is a DSL (Domain Specific Language) designed for Grule. It's a
simplified language to be used for defining rule condition criteria and actions
to be executed if the criteria are met.

The language has the following structure:

```Shell
rule <RuleName> <RuleDescription> [salience <priority>] {
    when
        <boolean expression>
    then
        <assignment or operation expression>
}
```

**RuleName**: Identifies a specific rule. The name must be unique in the entire
knowledge base, consist of one word and it must not contain white space.

**RuleDescription**: Describes the rule for human consumption. The description
should be enclosed in double quotes.

**Salience** (optional, default 0): Defines the importance of the rule. Lower
values indicate rules of lower priority. The salience value is used to specify a
priority-sorted order when multiple rules are encountered. Salience will accept
negative values, so you may wish to use those to mark rules that you barely even
care about. Rule engines are *declarative* so you can't be guaranteed in which
order your rules will be evaluated.  As such, consider `salience` to be a *hint*
to the engine that helps it decide what to do in the event of a conflict.

**Boolean Expression**: A predicate expression that will be evaluated by the
rule engine to identify whether or not a specific rule's action is a candidate
for execution with the current facts.

**Assignment or Operation Expression**: This is the action to be taken should
the rule evaluate to `true`. You are not limited to a single expression and can
supply a list of them, separated by the `;` character. The action statements are
meant to modify the current fact values, make calculations, log some statements,
etc...

### Boolean Expression

A boolean expression should be familiar to most, if not all programmers.

```go
when
     contains(User.Name, "robert") &&
     User.Age > 35
then
     ...
```

### Constants and Literals

| Literal | Description                                                                | Example                                            |
| ------- | -------------------------------------------------------------------------- | -------------------------------------------------- |
| String  | Holds a string literal, enclosed with double (&quot;) or single (') quotes | "This is a string" or 'this is a string'           |
| Integer | Holds an integer value and may preceded with negative symbol -             | `1` or `34` or `42344` or `-553`                   |
| Real    | Holds a real value                                                         | `234.4553`, `-234.3`, `314E-2`, `.32`, `12.32E12`  |
| Boolean | Holds a boolean value                                                      | `true`, `TRUE`, `False`                            |

More examples can be found at [GRL Literals](GRL_Literals_en.md).

Note: Special characters in strings must be escaped following the same rules
used for strings in Go.  However, backtick strings are not supported.

### Operators supported 

| Type                 | Operator                          |
| -------------------- | --------------------------------- |
| Math                 |  `+`, `-`, `/`, `*`, `%`          |
| Bit-wise operators   | `\|`, `&`                         |
| Logical operators    | `&&`, `\|\|`                      |
| Comparison operators | `<`, `<=`, `>`, `>=`, `==`, `!=`  |

### Operator precedence

Grule follows operator precedence in Go.

| Precedence | Operator                         |
| ---------- | -------------------------------- |
|    5       | `*`, `/`, `%`, `&`               |
|    4       | `+`, `-`, `\|`                   |
|    3       | `==`, `!=`, `<`, `<=`, `>`, `>=` |
|    2       | `&&`                             |
|    1       | `\|\|`                           |

### Comments

Comments also follow the standard Go format.

```go
// This is a comment
// And this

/* And also this */

/*
   As well as this
*/
```

### Array/Slice and Map

Since version 1.6.0, Grule supports accessing facts in array/slice or map.

Suppose you have a fact structure like the following:

```go
type MyFact struct {
    AnIntArray   []int
    AStringArray []string
    SubFacts     []*MyFact
    SubMaps      map[string]*MyFact
}
```

You can evaluate those slices and maps from your rule with:

```go
    when 
       Fact.AnIntArray[1] == 12 &&
       Fact.AStringArray[12] != "SomeText" &&
       Fact.SubFacts[1].SubFacts[2].AnIntArray[12] > 100 &&
       Fact.SubMaps["Key"].AnIntArray[0] == 1000
    then
       ...
```

Rule execution will panic if your rule tries to access an array element that
is out of bounds.

#### Assigning values into Array/Slice and Map

You can set an array value if the index you specify is valid.

```go
   then
      Fact.AnIntArray[10] = 12;
      Fact.SubMap["AKey"].AStringArray[1] = "New Value";
      Fact.AnotherMap[Fact.SomeFunction()] = "Another Value";
```

There are a couple of functions you can use to work with array/slice and map.
Those can be found at [Function page](Function_en.md).

### Negation

A unary negation symbol `!` is supported by GRL in addition to NEQ `!=` symbol.
It is to be used in front of a boolean expression or expression atom.

For example in expression atom:

```go
when 
    !FunctionReturnTrue() ||
    !false
then
    ... 
```

or in expression:

```go
when
    !(you.IsOk() || !today.isMonday())
then
    ...
```

### Function call

Any visible function can be called from your rule so long as they return 0 or 1
value.  For example:

```go
    when
        Fact.FunctionA() == "text" ||
        Fact.FunctionB("arg") == "text" ||
        Fact.FunctionC(Fact.Field, true)
    then
        Fact.CallFunction();
        Fact.Value = Fact.CallFunction();
        ...
```

In version 1.6.0, Grule can chain function calls and value accessors.  For
example;

```go
    when
        Fact.Function().StringField == "" ||
        Fact.Function("contant").ObjField.OtherFunction() &&
        ...
    then
        Fact.CallFunction().CallAnotherFunction();
        ...
```

Also introduced in 1.6.0, you can call functions on literatl constants.  For
example:

```go
    when
        "AString   ".Trim().ToUpper().HasSuffix("ING")
    then
        Fact.Result = Fact.ReturnStringFunc().Trim().ToLower();
```

For a list of available functions, consult [Function Page](Function_en.md).

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

### IDE Support

Visual Studio Code: [https://marketplace.visualstudio.com/items?itemName=avisdsouza.grule-syntax](https://marketplace.visualstudio.com/items?itemName=avisdsouza.grule-syntax)
