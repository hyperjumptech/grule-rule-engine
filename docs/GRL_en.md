# Grule Rule Language (GRL)

[Tutorial](Tutorial_en.md) | [Rule Engine](RuleEngine_en.md) | [GRL](GRL_en.md) | [RETE Algorithm](RETE_en.md) | [Functions](Function_en.md) | [Grule Events](GruleEvent_en.md)

The **GRL**  is a DSL (Domain Specific Language) designed for Grule. Its a simplified language
to be used for defining rule evaluation criteria and action to be executed if the criteria were met.

Generally, the language have the following structure :

```text
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
| String  | Hold string literal, enclosed a string with double quote symbol &quot; or a single quote | "This is a string" or 'this is a string' |
| Decimal | Hold a decimal value, may preceeded with negative symbol -             | `1` or `34` or `42344` or `-553` |
| Real    | Hold a real value                                                      | `234.4553`, `-234.3` , `314E-2`, `.32` , `12.32E12`  |
| Boolean | Hold a boolean value                                                   | `true`, `TRUE`, `False`          |

Operators supported :

* Math operators : `+`, `-`, `/`, `*`, `%`
* Bit-wise operators : `|`, `&`
* Logical operators : `&&` and `||`
* Comparison operators : `<`,`<=`,`>`,`>=`,`==`,`!=` 

Operator precedence : 

Grule follows operator precedence in Golang.

| Precedence |  Operator |
| ---------- | --------- |
|    5       |      `*` , `/` , `%` , `&` |
|    4       |      `+` , `-` , `\|`     |
|    3       |      `==` , `!=` , `<` , `<=` ,`>` , `>=`  |
|    2       |      `&&`  |
|    1       |      `\|\|`  |

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
