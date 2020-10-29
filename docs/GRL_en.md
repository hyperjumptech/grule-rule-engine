# Grule Rule Language (GRL)

[Tutorial](Tutorial_en.md) | [Rule Engine](RuleEngine_en.md) | [GRL](GRL_en.md) | [RETE Algorithm](RETE_en.md) | [Functions](Function_en.md) | [FAQ](FAQ_en.md)

The **GRL** is a DSL (Domain Specific Language) designed for Grule. It's a simplified language
to be used for defining rule evaluation criterias and actions to be executed if the criteria(s) are met.

Generally, the language have the following structure :

```Shell
rule <RuleName> <RuleDescription> [salience <priority>] {
    when
        <boolean expression>
    then
        <assignment or operation expression>
}
```

**RuleName** identifies a specific rule. The name should be unique in the entire knowledge base, consist of one word and
it should not contains a white-space.

**RuleDescription** describes the rule. The description should be enclosed with a double-quote.

**Salience** defines the importance of the rule. It's an optional rule configuration, and by default, when you don't specify them, all rules will have the salience value of 0 (zero).
The lower the value, the less important the rule. Whenever multiple rules become a candidate for execution, the highest salience rule will be executed first. You may define negative values for the salience to make the salience even lower. Like any implementation of Rule-Engine, there are no definitive algorithm to specify which rule to be execute in case of conflicting candidate, the engine may run which ever they like.
Salience is one way of hinting to the rule engine which rules have more importance relatively.

**Boolean Expression** is an expression that will be evaluated by the rule engine to identify if that specific rule
is a candidate for execution with the current facts.

**Assignment or Operation Expression** contains list of expressions (each expression should be ended with ";" character).
The expression are designed to modify the current fact values, making calculation, make some logging, etc.

### Boolean Expression

Boolean expression comes naturally for java or golang developer in GRL.

```go
when
     contains(User.Name, "robert") &&
     User.Age > 35
then
     ...
```

### Constants and Literals

| Literal | Description                                                            | Example                          |
| ------- | ---------------------------------------------------------------------- | -------------------------------- |
| String  | Hold string literal, enclosed a string with double quote symbol &quot; or a single quote ' | "This is a string" or 'this is a string' |
| Decimal | Hold a decimal value, may preceded with negative symbol -             | `1` or `34` or `42344` or `-553` |
| Real    | Hold a real value                                                      | `234.4553`, `-234.3` , `314E-2`, `.32` , `12.32E12`  |
| Boolean | Hold a boolean value                                                   | `true`, `TRUE`, `False`          |

Note: Strings are escaped following the same rules used for standard Go strings. Backtick strings are not supported.

### Operators supported 

| Type | Operator                                                            |
| ------- | ---------------------------------------------------------------------- |
| Math |  `+`, `-`, `/`, `*`, `%` |
| Bit-wise operators | `\|`, `&` |
| Logical operators | `&&` and `\|\|` |
| Comparison operators | `<`, `<=`, `>`, `>=`, `==`, `!=`  |

### Operator precedence

Grule follows operator precedence in Golang.

| Precedence |  Operator |
| ---------- | --------- |
|    5       |      `*` , `/` , `%` , `&` |
|    4       |      `+` , `-` , `\|`     |
|    3       |      `==` , `!=` , `<` , `<=` ,`>` , `>=`  |
|    2       |      `&&`  |
|    1       |      `\|\|`  |

### Comments

You can always put a comment inside your GRL script. Such as :

```go
// This is a comment
// And this

/* And also this */

/*
   As well as this
*/
```

### Array/Slice and Map

Since version 1.6.0, Grule support accessing fact in array/slice or map.

Suppose you have a fact structure like the following :

```go
type MyFact struct {
    AnIntArray   []int
    AStringArray []string
    SubFacts     []*MyFact
    SubMaps      map[string]*MyFact
}
```

You can always evaluate those slice and map from your rule such as

```go
    when 
       Fact.AnIntArray[1] == 12 &&
       Fact.AStringArray[12] != "SomeText" &&
       Fact.SubFacts[1].SubFacts[2].AnIntArray[12] > 100 &&
       Fact.SubMaps["Key"].AnIntArray[0] == 1000
    then
       ...
```

Rule execution will fail if your rule is trying to access array element
that beyond the fact's capacity.

#### Assigning values into Array/Slice and Map

You can always set an array value if the index you specify is valid.

```go
   then
      Fact.AnIntArray[10] = 12;
      Fact.SubMap["AKey"].AStringArray[1] = "New Value";
      Fact.AnotherMap[Fact.SomeFunction()] = "Another Value";
```

There are a couple functions you can use to work with array/slice and map in the [Function page](Function_en.md)

### Function call

From your Rule, you can always call any visible functions your fact have. As long as they're visible and have no or max 1 return value.
For example:

```go
    when
        Fact.Function() == "text" ||
        Fact.Function("arg") == "text" ||
        Fact.Function(Fact.Field, true)
    then
        Fact.CallFunction();
        Fact.Value = Fact.CallFunction();
        ...
```

In version 1.6.0, Grule can chain function return value.
For example;

```go
    when
        Fact.Function().StringField == "" ||
        Fact.Function("contant").ObjField.OtherFunction() &&
        ...
    then
        Fact.CallFunction().CallAnotherFunction();
        ...
```

Also introduced in 1.6.0, constants value may have built in functions to manipulated them.
For example;

```go
    when
        "AString   ".Trim().ToUpper().HasSuffix("ING")
    then
        Fact.Result = Fact.ReturnStringFunc().Trim().ToLower();
```

List of these constant functions can be found in the [Function Page](Function_en.md).

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
