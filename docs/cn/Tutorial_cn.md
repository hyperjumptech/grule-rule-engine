# Grule 简短教程

[![Tutorial_cn](https://github.com/yammadev/flag-icons/blob/master/png/CN.png?raw=true)](../cn/Tutorial_cn.md)
[![Tutorial_de](https://github.com/yammadev/flag-icons/blob/master/png/DE.png?raw=true)](../de/Tutorial_de.md)
[![Tutorial_en](https://github.com/yammadev/flag-icons/blob/master/png/GB.png?raw=true)](../en/Tutorial_en.md)
[![Tutorial_id](https://github.com/yammadev/flag-icons/blob/master/png/ID.png?raw=true)](../id/Tutorial_id.md)
[![Tutorial_pl](https://github.com/yammadev/flag-icons/blob/master/png/PL.png?raw=true)](../pl/Tutorial_pl.md)

[About](About_cn.md) | [Tutorial](Tutorial_cn.md) | [Rule Engine](RuleEngine_cn.md) | [GRL](GRL_cn.md) | [GRL JSON](GRL_JSON_cn.md) | [RETE Algorithm](RETE_cn.md) | [Functions](Function_cn.md) | [FAQ](Tutorial_cn.md) | [Benchmark](Benchmarking_cn.md)

---

## 准备

Grule 使用的Go 1.24.4版本。

以如下的方式在你的项目中引入Grule.

```Shell
$ go get github.com/hyperjumptech/grule-rule-engine
```

在使用之前请先import Grule.

```go
import (
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
) 
```

## 创建Fact结构体

在grule中事实`fact`，是一个`struct`的实例指针。结构可以像正常的Golang结构体一样包含属性，也可以有方法`method`, 当然在grule中`method`有一些限制。

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

正如Golang的理念，Grule只能访问可见的属性和方法，在Golang中，可见的属性和方法是大写字母开头。

```go
func (mf *MyFact) GetWhatToSay(sentence string) string {
    return fmt.Sprintf("Let say \"%s\"", sentence)
}
```

**注意**：成员函数有以下一些限制：

* 成员函数可见，即函数名的首字母大写。
* 成员函数只能返回0个或者1个返回值。大于1个返回值不支持。
*  所有的数字类型的参数和返回值都应该是64bit的。比如`int64`, `uint64`, `float64`。
* 成员函数不应该改变Fact的内部状态。算法不能自动监测到这些变动，这就使得查原因变得很复杂，而且会有隐含的bug。如果你一定要改变Fact的内部状态，你可以通过调用`Changed(varname string)`函数通知Grule。

## 添加 Fact 到 DataContext

添加 fact 到 `DataContext`，你需要创建 `fact`的实例。

```go
myFact := &MyFact{
    IntAttribute: 123,
    StringAttribute: "Some string value",
    BooleanAttribute: true,
    FloatAttribute: 1.234,
    TimeAttribute: time.Now(),
}
```

你可以创建任意多的fact。

在fact创建之后，你可以添加这些实例到`DataContext`:

```go
dataCtx := ast.NewDataContext()
err := dataCtx.Add("MF", myFact)
if err != nil {
    panic(err)
}
```

### 从json创建Fact

在 Grule 1.8.0版本之后，描述fact的json数据可以被加载到`DataContext`中。详情参考 [JSON as a Fact](../cn/JSON_Fact_cn.md).

## 创建一个知识库，添加规则进知识库

`KnowledgeLibrary`是`KnowledgeBase`的集合，而`KnowledgeBase`是各种来源规则的集合。

我们使用`RuleBuilder`去创建`KnowledgeBase`实例，然后添加到`KnowledgeLibrary`中。

 GRL 有很多来源:

* 原始字符串
* 文件内容
* http请求结果

使用`RuleBuilder`去填充我们的 `KnowledgeLibrary`.

```go
knowledgeLibrary := ast.NewKnowledgeLibrary()
ruleBuilder := builder.NewRuleBuilder(knowledgeLibrary)
```

接下来我们使用DSL的原始字符串去定义一个基本的规则：

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

最后我们使用builder将定义创建一个`resource`然后添加到`knowledgeLibrary`中。

```go
// Add the rule definition above into the library and name it 'TutorialRules'  version '0.0.1'
bs := pkg.NewBytesResource([]byte(drls))
err := ruleBuilder.BuildRuleFromResource("TutorialRules", "0.0.1", bs)
if err != nil {
    panic(err)
}
```

 `KnowledgeLibrary` 现在包含了一个名字为`TutorialRules`的而且版本是`0.0.1`的`KnowledgeBase`。为了执行规则，我们需要从`KnowledgeLibrary`中获取一个实例，将在下面的章节阐述。

## 执行 Grule 规则引擎

为了执行KnowledgeBase，我们需要先从`KnowledgeLibrary`获取这个`KnowledgeBase`。

```go
knowledgeBase := knowledgeLibrary.NewKnowledgeBaseInstance("TutorialRules", "0.0.1")
```

从`knowledgeLibrary`获得的每一个实例都是`KnowledgeBase`的唯一克隆。每个唯一的实例都有自己的独立的`WorkingMemory`。因为实例不会共享状态给其他的实例，所以你可以在多线程中随意使用，只要你不是在多线程中同时执行了同一个实例。

从`KnowledgeBase`蓝图构造也保证了我们不需要每次使用实例的都再重新计算一次。计算工作只需要做一次，从`AST`拷贝是极其高效的。

然后使用准备好的`DataContext`去执行`KnowledgeBase`实例。

```go
engine = engine.NewGruleEngine()
err = engine.Execute(dataCtx, knowledgeBase)
if err != nil {
    panic(err)
}
```

## 获取结果

下面是我们定义的规则：

```go
rule CheckValues "Check the default values" salience 10 {
    when 
        MF.IntAttribute == 123 && MF.StringAttribute == "Some string value"
    then
        MF.WhatToSay = MF.GetWhatToSay("Hello Grule");
        Retract("CheckValues");
}
```

假设满足了条件，所实施的动作是去改变`MF.WhatToSay`属性。为了保证规则不会立即被重新赋值，调用`Retract`去从集合中收回，即不再下次循环的时候执行这个规则。在这个特殊的实例中，如果规则没有执行，它将在下一个循环中继续匹配，然后一直重复。最终Grule将会以一个error终止，因为它无法收敛于终止结果。

在这个案例中，你可以通过查看`myFact`去获取规则的结果。

```go
fmt.Println(myFact.WhatToSay)
// this should prints
// Lets Say "Hello Grule"
```
## 资源

GRLs can be stored in external files and there are many ways to obtain and load
the contents of those files.

GRLs 规则列表可以存储在外部文件中，有很多方式去获取、加载这些文件内容。

### 从文件获取

```go
fileRes := pkg.NewFileResource("/path/to/rules.grl")
err := ruleBuilder.BuildRuleFromResource("TutorialRules", "0.0.1", fileRes)
if err != nil {
    panic(err)
}
```

通过指定多个路径或者模式，你可以一次加载多个文件。

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

### 从字符串或者ByteArray获取

```go
bs := pkg.NewBytesResource([]byte(rules))
err := ruleBuilder.BuildRuleFromResource("TutorialRules", "0.0.1", bs)
if err != nil {
    panic(err)
}
```

### 从URL获取

```go
urlRes := pkg.NewUrlResource("http://host.com/path/to/rule.grl")
err := ruleBuilder.BuildRuleFromResource("TutorialRules", "0.0.1", urlRes)
if err != nil {
    panic(err)
}
```

#### 携带Headers

```go
headers := make(http.Header)
headers.Set("Authorization", "Basic YWxhZGRpbjpvcGVuc2VzYW1l")
urlRes := pkg.NewURLResourceWithHeaders("http://host.com/path/to/rule.grl", headers)
err := ruleBuilder.BuildRuleFromResource("TutorialRules", "0.0.1", urlRes)
if err != nil {
    panic(err)
}
```

### 从Git获取

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

### 从 JSON获取

你可以从JSON构建规则! [请阅读这里](GRL_JSON_cn.md) 

## 编译 GRL 到 GRB

如果你想要更快的规则加载速度（比如你有很多特别大的规则集合，然后加载GRL很慢），你可以将这些规则集合存储为GRB（Grules Rule Binary）文件。详情参考[如何存储和加载GRB](Binary_Rule_File_cn.md)
