# Grule JSON Format

Grule rules can be represented in JSON form and translated for use by the rule engine into standard Grule syntax. The design of the JSON format is intended to offer a high level of flexability to suite the needs of the user.

The basic structure of a JSON rule is as follows:
```json
{
    "name": "SpeedUp",
    "desc": "When testcar is speeding up we keep increase the speed.",
    "salience": 10,
    "when": ...,
    "then": [
        ...
    ]
}
```

## Elements

| Name       | Description                                                                                                             |
| ---------- | ----------------------------------------------------------------------------------------------------------------------- |
| `name`     | The name of the rule. This field is required.                                                                           |
| `desc`     | The description for the rule. If not set the rules description will be set to `""`                                      |
| `salience` | The salience value for the rule. If this value is not set the saliance will default to 0                                |
| `when`     | The when conndition for the rule. This field can either be a plain string value or a condition object (described below) |
| `then`           | An array of then actions for the rule. Each aray element can be a plain string or a then object (described below)                                                                                                                        |

## Condition Object

In order to provide a great deal of flexability, the when condition of a rule can be broken down into individual components. This is particularly useful for structuring larger rules and supporting GUI applications used for rule editing and analysis.

Condition objects are parsed recursively, meaning that objects can be nested arbitarily to support even the most complex rules. Any time a condition object is expected by the parser, the user can choose to instead provide a constant string or numeric value which will be interpreted by the parser as raw input to be echoed into the output rule.

Each condition object takes the following format:
```json
{"operator":[x, y...]}
```
where operator is one of the operators described below and x and y two or more condition objects or constants.

### Operators

| Operator  | Description       |
| --------- | ----------------- |
| `"and"`   | GRL && operator   |
| `"or"`    | GRL \|\| operator |
| `"eq"`    | GRL == operator   |
| `"not"`   | GRL != operator   |
| `"gt"`    | GRL > operator    |
| `"gte"`   | GRL >= operator   |
| `"lt"`    | GRL < operator    |
| `"lte"`   | GRL <= operaotr   |
| `"bor"`   | GRL \| operator   |
| `"band"`  | GRL & operator    |
| `"plus"`  | GRL + operator    |
| `"minus"` | GRL - operator    |
| `"div"`   | GRL / operator    |
| `"mul"`   | GRL * operator    |
| `"mod"`   | GRL % operator    |

### Special Operators

The following operators have slightly different behaviours than the standard operators.

| Operator  | Description                                                                                                                                                                                                                |
| --------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `"set"`   | GRL = operator. This operator will set the value of the first operand to the output of the second. This can only be used in the `then` section of the rule.                                                                |
| `"call"`  | GRL function call. The operator will call the funtion name specified in the first operand. If more then one operand is specififed, the subsequent operands are interpreted as arguments to be passed to the funciton call. |
| `"obj"`   | Explicitly identifies a GRL object. Unkline other operators, this object takes the form of a simple key/value pair. For example: `{"obj": "TestCar.Speed"}`                                                                |
| `"const"` | Explicitly identified a GRL constant. This opertor takes the same form as the `obj` operator                                                                                                                               |

### Supported Constants

The following constant types are supported:

| Type      | Example                     |
| --------- | --------------------------- |
| `string`  | `{"const": "String Value"}` |
| `integer` | `{"const": 123}`            |
| `float`   | `{"const": 1.29738}`        |
| `bool`    | `{"const": true}`           |


## Then Actions

The `then` actions are formed in the same way as the condition objects. The primary difference is that the root element for each then condition should be either a `set` or `call` operator.

# Example

To demonstrate the JSON representation capabilities, the following example rule is to be converted to JSON syntax:

```
rule SpeedUp "When testcar is speeding up we keep increase the speed." salience 10 {
    when
        TestCar.SpeedUp == true && TestCar.Speed < TestCar.MaxSpeed
    then
        TestCar.Speed = TestCar.Speed + TestCar.SpeedIncrement;
        DistanceRecord.TotalDistance = DistanceRecord.TotalDistance + TestCar.Speed;
        Log("Speed increased");
}
```

## Basic Representation

The most basic representation of this rule in JSON is as follows:
```json
{
    "name": "SpeedUp",
    "desc": "When testcar is speeding up we keep increase the speed.",
    "salience": 10,
    "when": "TestCar.SpeedUp == true && TestCar.Speed < TestCar.MaxSpeed",
    "then": [
        "TestCar.Speed = TestCar.Speed + TestCar.SpeedIncrement",
        "DistanceRecord.TotalDistance = DistanceRecord.TotalDistance + TestCar.Speed",
        "Log(\"Speed increased\")"
    ]
}
```

This example presents the when and then conditions as raw input objects. This gives the greatest level of control over the output rule. In most cases the translator will output a rule which is an exact match to the original representaion, however in some cases the translator may insert brackets around expressions where it is not required depending on operator presedence. This should not affect the logical meaning of the rule.

## Expanded Representation

The above rule can also be represented in a more explicit format by breaking down the when and then conditions to full object representation:
```json
{
    "name": "SpeedUp",
    "desc": "When testcar is speeding up we keep increase the speed.",
    "salience": 10,
    "when": {
       "and": [
           {"eq": ["TestCar.SpeedUp", true]},
           {"lt": ["TestCar.Speed", "TestCar.MaxSpeed"]}
       ]
    },
    "then": [
        {"set": ["TestCar.Speed", {"plus": ["TestCar.Speed", "TestCar.SpeedIncrement"]}]},
        {"set": ["DistanceRecord.TotalDistance", {"plus": ["DistanceRecord.TotalDistance", "TestCar.Speed"]}]},
        {"call": ["Log", {"const": "Speed increased"}]}
    ]
}
```

The translator will interpret the above rule and produce the same output as the example rule. This rule is much more verbose and can be easily parsed and formated for display or analysis purposes.

## Implicit Representation and String Escaping

Despite being more verbose, the rule above still represent objects and boolean constants implictly. It is not necessary to wrap each object and constant inside of a `const` or `obj` operator because the translator will implicitly interpret the type of input. However, in some cases it can be advantagious to use the `const` or `obj` wrappers, the most notable of which is to enforce escaping rules on string constants that would not be applied to an implicilt string.

An example of string escaping behaviour shown in the rule above is in the `call` then action. A string is being passed as an argument to the function so in order to pass the argument as a raw object it is necessary for the user to wrap the constant in double quotes and escape them manually:
`{"call": ["Log", "\"Speed increased\""]}`

While this gives the correct output, it somewhat obfuscates the rule by making it unclear which level of escaping is applied to the output.

If the escaping of the string is done incorrectly, the translator will generate an invalid rule output. In order to prevent this kind of error, it is preferable to wrap the constant in a `const` operator:
`{"call": ["Log", {"const": "Speed increased"}]}`

I this case, the translator will escape the constant string appropriately and ouput the constant correctly to the output rule.

## Verbose Representation

The most verbose possible version of the example rule in JSON syntax is as follows. Note that the additional `obj` and `const` operators here are completely unnecessary but can be useful for rendering engines or tools designed to edit or analyse rules in a visual form.

```json
{
    "name": "SpeedUp",
    "desc": "When testcar is speeding up we keep increase the speed.",
    "salience": 10,
    "when": {
       "and": [
           {"eq": [{"obj": "TestCar.SpeedUp"}, {"const": true}]},
           {"lt": [{"obj": "TestCar.Speed"}, {"obj": "TestCar.MaxSpeed"}]}
       ]
    },
    "then": [
        {"set": [{"obj": "TestCar.Speed"}, {"plus": [{"obj": "TestCar.Speed"}, {"obj": "TestCar.SpeedIncrement"}]}]},
        {"set": [{"obj": "DistanceRecord.TotalDistance"}, {"plus": [{"obj": "DistanceRecord.TotalDistance"}, {"obj": "TestCar.Speed"}]}]},
        {"call": ["Log", {"const": "Speed increased"}]}
    ]
}
```

# Loading JSON Rules

JSON rules can be loaded from an underlying Resource or ResourceBundle provider using the functions `NewJSONResourceFromResource` and `NewJSONResourceBundleFromBundle` respectively. When the `Load()` function is called the underlying Resource is loaded and the rules are translated from JSON syntax to standard GRL syntax. As a result the translation functions can very easily be integrated into existing code.

```go
f, err := os.Open("rules.json")
if err != nil {
    panic(err)
}
underlying := pkg.NewReaderResource(f)
resource := pkg.NewJSONResourceFromResource(underlying)
...
```

It is also possible to parse a byte array containing JSON rules directly into a GRL syntax ruleset by calling the `ParseJSONRuleset` function.

```go
jsonData, err := ioutil.ReadFile("rules.json")
if err != nil {
    panic(err)
}
ruleset, err := pkg.ParseJSONRuleset(jsonData)
if err != nil {
    panic(err)
}
fmt.Println("Parsed ruleset: ")
fmt.Println(ruleset)
```
