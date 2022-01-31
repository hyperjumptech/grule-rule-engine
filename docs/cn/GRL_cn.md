# Grule Rule Language (GRL)

[![GRL_cn](https://github.com/yammadev/flag-icons/blob/master/png/CN.png?raw=true)](../cn/GRL_cn.md)
[![GRL_de](https://github.com/yammadev/flag-icons/blob/master/png/DE.png?raw=true)](../de/GRL_de.md)
[![GRL_en](https://github.com/yammadev/flag-icons/blob/master/png/GB.png?raw=true)](../en/GRL_en.md)
[![GRL_id](https://github.com/yammadev/flag-icons/blob/master/png/ID.png?raw=true)](../id/GRL_id.md)

[About](About_cn.md) | [Tutorial](Tutorial_cn.md) | [Rule Engine](RuleEngine_cn.md) | [GRL](GRL_cn.md) | [GRL JSON](GRL_JSON_cn.md) | [RETE Algorithm](RETE_cn.md) | [Functions](Function_cn.md) | [FAQ](GRL_cn.md) | [Benchmark](Benchmarking_cn.md)

---

**GRL**是为了Grule设计的DSL(Domain Specific Language). GRL是一种简单的语法，可以用来定义规则的条件和条件满足之后将执行哪些操作。

语法将按照如下的结构:

```Shell
rule <RuleName> <RuleDescription> [salience <priority>] {
    when
        <boolean expression>
    then
        <assignment or operation expression>
}
```

**RuleName**: 规则名称标识特定规则. 在整个knowledge base中，整个名字需要是唯一的。规则名字智能是一个单词而且不能包含空格。

**RuleDescription**: 关于规则的描述，便于理解。描述应该包在双引号内。

**Salience** (optional, default 0): 优先级，可选参数，默认为0. 定义了规则的重要性。值越低，优先级越低。优先级用于同时满足多个规则时的排序使用。优先级可以是负数，你可能希望去标识一些很少注意的规则。规则引擎是声明式的，所以你不能保证规则的执行顺序。因此，考虑使用` salience`去提示规则引擎当发生冲突的时候应该使用哪个规则。

**Boolean Expression**: 布尔表达式。一个断言表达式，规则引擎用来评估事实是否满足表达式，然后是否要执行操作。

**Assignment or Operation Expression**: 赋值或者操作表达式。当规则评估为`true`就会执行这里的操作。这里你可以写多个表达式，使用`;`分割。操作可以是改变事实Fact的变量，继续一些计算，打日志等等。

### 布尔表达式

布尔表达式对大多数人都很熟悉，不仅仅是研发。

```go
when
     contains(User.Name, "robert") &&
     User.Age > 35
then
     ...
```

### 字面值常量

| 类型    | 表述                                             | 示例                                              |
| ------- | ------------------------------------------------ | ------------------------------------------------- |
| String  | 存储了字符串，使用双引号 (&或者 (') 单引号包围。 | "This is a string" 或者 'this is a string'        |
| Integer | 整数，可以是负数                                 | `1` or `34` or `42344` or `-553`                  |
| Real    | 浮点数                                           | `234.4553`, `-234.3`, `314E-2`, `.32`, `12.32E12` |
| Boolean | 布尔值                                           | `true`, `TRUE`, `False`                           |

更多类型可以参考 [GRL Literals](GRL_Literals_cn.md).

注意: 如Go中的字符串，特殊字符应该使用转义.。不管怎么样，反引号字符串不支持。

### 操作符

| Type                 | Operator                         |
| -------------------- | -------------------------------- |
| Math                 | `+`, `-`, `/`, `*`, `%`          |
| Bit-wise operators   | `|`, `&`                         |
| Logical operators    | `&&`, `||`                       |
| Comparison operators | `<`, `<=`, `>`, `>=`, `==`, `!=` |

### 操作符优先级

Grule 的操作符优先级与Go保持一致。

| Precedence | Operator                         |
| ---------- | -------------------------------- |
|    5       | `*`, `/`, `%`, `&`               |
|    4       | `+`, `-`, `\|`                   |
|    3       | `==`, `!=`, `<`, `<=`, `>`, `>=` |
|    2       | `&&`                             |
|    1       | `\|\|`                           |

### 注释

注释也与Go语法相同。

```go
// This is a comment
// And this

/* And also this */

/*
   As well as this
*/
```

### 数组/切片，map

从版本1.6.0开始，Grule支持访问facts中的数组/切片或者map。

假设你的fact结构体如下：

```go
type MyFact struct {
    AnIntArray   []int
    AStringArray []string
    SubFacts     []*MyFact
    SubMaps      map[string]*MyFact
}
```

你可以在你的规则中访问你的切片和map。

```go
    when 
       Fact.AnIntArray[1] == 12 &&
       Fact.AStringArray[12] != "SomeText" &&
       Fact.SubFacts[1].SubFacts[2].AnIntArray[12] > 100 &&
       Fact.SubMaps["Key"].AnIntArray[0] == 1000
    then
       ...
```

如果数组越界了，规则执行将会导致panic。

#### 赋值给 Array/Slice 和 Map

只要索引有效，你可以给数组赋值。

```go
   then
      Fact.AnIntArray[10] = 12;
      Fact.SubMap["AKey"].AStringArray[1] = "New Value";
      Fact.AnotherMap[Fact.SomeFunction()] = "Another Value";
```

有一些内置的可以操作array/slice,map 的函数。可以参考[Function page](Function_cn.md).

### 否定

GRL支持一元否定操作符 `!` ，等同于 `!=` 符号.

可以在布尔表达式前使用。

表达式原子中使用

```go
when 
    !FunctionReturnTrue() ||
    !false
then
    ... 
```

或者在表达式中:

```go
when
    !(you.IsOk() || !today.isMonday())
then
    ...
```

### 函数调用

可见的函数而且返回一个或者0个值，就可以被调用。举例：

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

在1.6.0版本，Grule可以链式调用和值访问。比如：

```go
    when
        Fact.Function().StringField == "" ||
        Fact.Function("contant").ObjField.OtherFunction() &&
        ...
    then
        Fact.CallFunction().CallAnotherFunction();
        ...
```

同样是1.6.0中引入的，你可以对常量进行函数调用。举例：

```go
    when
        "AString   ".Trim().ToUpper().HasSuffix("ING")
    then
        Fact.Result = Fact.ReturnStringFunc().Trim().ToLower();
```

内置的函数列表参考 [Function Page](Function_cn.md).

#### 举例

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

### 调试 GRL 语法

 在你的应用中，你可以使用GRL脚本或者代码片段去测试GRL。

```go
        RuleWithError := `
        rule ErrorRule1 "Rule with error"  salience 10{
            when
              Pogo.Compare(User.Name, "Calo")  
            then
              User.Name = "Success";
              Log(User.Name)
              Retract("AgeNameCheck");
        }
        `

	// Build normally
	err := ruleBuilder.BuildRuleFromResource("Test", "0.1.1", pkg.NewBytesResource([]byte(RuleWithError)))

	// If the err != nil something is wrong.
	if err != nil {
		// Cast the error into pkg.GruleErrorReporter with typecast checking.
		// Typecast checking is necessary because the err might not only parsing error.
		if reporter, ok := err.(*pkg.GruleErrorReporter); ok {
			// Lets iterate all the error we get during parsing.
			for i, er := range reporter.Errors {
				fmt.Printf("detected error #%d : %s\n", i, er.Error())
			}
		} else {
			// Well, its an error but not GruleErrorReporter instance. could be IO error.
			t.Error("There should be GruleErrorReporter")
			t.FailNow()
		}
	}
```

将会打印出来

```txt
detected error #0 : grl error on 8:6 missing ';' at 'Retract'
```


### IDE 支持

支持Visual Studio Code插件: [https://marketplace.visualstudio.com/items?itemName=avisdsouza.grule-syntax](https://marketplace.visualstudio.com/items?itemName=avisdsouza.grule-syntax)
