# Grule JSON 格式

[![GRL_JSON_cn](https://github.com/yammadev/flag-icons/blob/master/png/CN.png?raw=true)](../cn/GRL_JSON_cn.md)
[![GRL_JSON_de](https://github.com/yammadev/flag-icons/blob/master/png/DE.png?raw=true)](../de/GRL_JSON_de.md)
[![GRL_JSON_en](https://github.com/yammadev/flag-icons/blob/master/png/GB.png?raw=true)](../en/GRL_JSON_en.md)
[![GRL_JSON_id](https://github.com/yammadev/flag-icons/blob/master/png/ID.png?raw=true)](../id/GRL_JSON_id.md)
[![GRL_JSON_pl](https://github.com/yammadev/flag-icons/blob/master/png/PL.png?raw=true)](../pl/GRL_JSON_pl.md)

[About](About_cn.md) | [Tutorial](Tutorial_cn.md) | [Rule Engine](RuleEngine_cn.md) | [GRL](GRL_cn.md) | [GRL JSON](GRL_JSON_cn.md) | [RETE Algorithm](RETE_cn.md) | [Functions](Function_cn.md) | [FAQ](GRL_JSON_cn.md) | [Benchmark](Benchmarking_cn.md)

---

Grule规则可以表达成json格式，而且可以翻译成标准的Grule语法表达。为了满足用户的需求，JSON格式的设计提供一种更高程度的灵活度 。

JSON格式的规则的基本结构如下：

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

## 元素

| Name       | Description                                    |
| ---------- | ---------------------------------------------- |
| `name`     | rule的名称，必填                               |
| `desc`     | 规则的描述，非必填，默认是 `""`                |
| `salience` | 规则的优先级，非必填，默认为`0`                |
| `when`     | 规则的条件. 可以是字符串，也可以是条件对象。   |
| `then`     | 规则的操作数组. 每个元素是字符串或者操作对象。 |

## 条件对象

为了提供更高的灵活性，`when`条件可能会被分割成独立的组件。当组建更大的更大的规则，而且需要支持能够编辑和分析规则和GUI应用时，这一部分非常有用。

条件对象将会被递归的解析，这就意味着对象可以任意嵌套，从而可以支持更复杂的规则。任何时候一个条件对象是解析器所期望的，用户可以选择提供一个常量解析器将解释为原始输入的字符串或数值被回显到输出规则中。

每个条件对象都将是一下的格式：

```json
{"operator":[x, y, ...]}
```

其中 `operator` 是以下操作符中的一个。 `x` 和 `y` 是两个或者多个条件对象或者变量。

### 操作符

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

### 特殊的操作符

下面的操作与标准的操作符有稍微的不同。 

| Operator  | Description                                                  |
| --------- | ------------------------------------------------------------ |
| `"set"`   | GRL = 操作符. 这个操作符将给第二个操作数的值付给第一个操作数。只能使用在规则的`then`中。 |
| `"call"`  | GRL 函数调用。第一个操作数是函数名字，如果有大于一个操作数，后续的操作数是这个函数调用的参数。 |
| `"obj"`   | 显式指定GRL对象。不像其他操作符，这个操作符由key/value 对组成。 比如: `{"obj": "TestCar.Speed"}` |
| `"const"` | 显式指定GRL常量。这个操作符格式与 `obj` 操作符相同。         |

### 支持的常量

支持以下的常量类型:

| Type      | Example                     |
| --------- | --------------------------- |
| `string`  | `{"const": "String Value"}` |
| `integer` | `{"const": 123}`            |
| `float`   | `{"const": 1.29738}`        |
| `bool`    | `{"const": true}`           |


## Then 操作

 `then` 操作组成方式跟条件对象相同。最主要的区别是每个`then`条件的根元素要么是`set`要么是`call`操作符。

# 示例

为了标识JSON的表示能力，下面的规则将会展示使用JSON语法进行展示。

```
rule SpeedUp "When testcar is speeding up we increase the speed." salience 10 {
    when
        TestCar.SpeedUp == true && TestCar.Speed < TestCar.MaxSpeed
    then
        TestCar.Speed = TestCar.Speed + TestCar.SpeedIncrement;
        DistanceRecord.TotalDistance = DistanceRecord.TotalDistance + TestCar.Speed;
        Log("Speed increased");
}
```

## 基础表示形式

最基础的JSON表示方法:

```json
{
    "name": "SpeedUp",
    "desc": "When testcar is speeding up we increase the speed.",
    "salience": 10,
    "when": "TestCar.SpeedUp == true && TestCar.Speed < TestCar.MaxSpeed",
    "then": [
        "TestCar.Speed = TestCar.Speed + TestCar.SpeedIncrement",
        "DistanceRecord.TotalDistance = DistanceRecord.TotalDistance + TestCar.Speed",
        "Log(\"Speed increased\")"
    ]
}
```

在这个示例中，`when`和`then`都是原始输入。这种表达将会更高层次控制输出的规则。在大多数的情况下，翻译器可以很简单的输出规则的原始表达。但是在大多数情况下，在不依赖操作符的时候，翻译器将会插入在表达式周围插入括号。这个不应该影响规则的逻辑意思。

## 扩展表示形式

通过继续拆分`when`和`then`条件成全对象形式，可以使上述规则可以表述成更详细的格式。

```json
{
    "name": "SpeedUp",
    "desc": "When testcar is speeding up we increase the speed.",
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

翻译器将会解释上述规则，并且会产出一样的规则。这个规则将会更冗长，而且对于展示或者分析更加容易解析和格式化。

## 隐式表示和字符串转义

尽管更冗长上述规则仍然可以更隐式表示成对象和布尔常量。并不需要将每个规则中的对象和常量都使用`const`或者`obj`操作符去包装，因为翻译器将会隐式获取输入的类型。不管如何，在某些情况下，使用`const`或者`obj`包装反而会有优势， 最显著的是在需要转义字符串的地方将不会在需要转义。

比如在`call`操作中的转义字符串。一个字符串作为参数传递给函数，为了传递原始对象，需要使用双引号

```
{"call": ["Log", "\"Speed increased\""]}
```

虽然这给出了正确的输出，但输出中包含某种程度的转义，使得规则变得不清晰，从而在某种程度上混淆了规则。

如果字符串不能被正确转义，翻译器将会生成一个无效的规则输出。为了防止这种错误，可以使用`const`操作符包裹常量。

```
{"call": ["Log", {"const": "Speed increased"}]}
```

在这种情况下，翻译器将会正确转义常量字符串，而且也能够正确输出。

## 详细表达形式

最详细的JSON语法表达如下。（注意额外附加的`obj`和`const`操作符实际不是完全必须的，但是对于可视形式用来编辑或者解析规则的渲染引擎或者工作很有用。）

```json
{
    "name": "SpeedUp",
    "desc": "When testcar is speeding up we increase the speed.",
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

# 加载 JSON 规则

通过使用`NewJSONResourceFromResource` 和 `NewJSONResourceBundleFromBundle`函数，JSON规则可以从底层资源或者资源包提供者加载规则。当调用`Load()`函数时，底层资源被加载，JSON语法的规则将会被翻译成标准GRL语法规则。翻译函数的结果可以很容易集成进已有的代码。

```go
f, err := os.Open("rules.json")
if err != nil {
    panic(err)
}
underlying := pkg.NewReaderResource(f)
resource := pkg.NewJSONResourceFromResource(underlying)
...
```

通过`ParseJSONRuleset`函数调用，可以解析包含JSON规则的字节数组并转换成GRL语法的规则集合。

```go
jsonData, err := io.ReadFile("rules.json")
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
