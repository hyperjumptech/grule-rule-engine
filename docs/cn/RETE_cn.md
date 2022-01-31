# Grule's RETE 算法

[![RETE_cn](https://github.com/yammadev/flag-icons/blob/master/png/CN.png?raw=true)](../cn/RETE_cn.md)
[![RETE_de](https://github.com/yammadev/flag-icons/blob/master/png/DE.png?raw=true)](../de/RETE_de.md)
[![RETE_en](https://github.com/yammadev/flag-icons/blob/master/png/GB.png?raw=true)](../en/RETE_en.md)
[![RETE_id](https://github.com/yammadev/flag-icons/blob/master/png/ID.png?raw=true)](../id/RETE_id.md)

[About](About_cn.md) | [Tutorial](Tutorial_cn.md) | [Rule Engine](RuleEngine_cn.md) | [GRL](GRL_cn.md) | [GRL JSON](GRL_JSON_cn.md) | [RETE Algorithm](RETE_cn.md) | [Functions](Function_cn.md) | [FAQ](RETE_cn.md) | [Benchmark](Benchmarking_cn.md)

---

来自维基百科 :  Rete  (/ˈriːtiː/ REE-tee, /ˈreɪtiː/ RAY-tee, rarely /ˈriːt/ REET, /rɛˈteɪ/ reh-TAY)算法是一种模式匹配算法，用来实现基于规则的系统。这个算法是为了在知识库中有效得把很多规则或者模式应用到多个对象或者事实上而开发的。它被用来针对一些数据事实来决定使用系统的哪个规则。 

在`grule-rule-engine`的版本`1.1.0` 中开始引入RETE算法。
当需要评估规则并加入`ConflictSet`时，它替代了稚嫩的解决方案

在GRL中的 `ExpressionAtom` 元素会被编译而且将不会被在引擎的工作内存中复制。当你有很多规则定义了重复的表达式或者很多重的函数调用时，这将会有效提升引擎的性能，

Grule RETE 实现不需要一个`Class`选择器，因为一个表达式可以被多个类需要。比如一个表达式如下：

```.go
when
    ClassA.attr == ClassB.attr + ClassC.AFunc()
then
    ...
```

上面的表达式需要属性和函数调用结果对比以及来自三个不同类的数学操作。这将会使得RETE的类和表达式分离边等很困难。

你可以参考 RETE 算法:

* https://en.wikipedia.org/wiki/Rete_algorithm
* https://www.drdobbs.com/architecture-and-design/the-rete-matching-algorithm/184405218
* https://www.sparklinglogic.com/rete-algorithm-demystified-part-2/ 

### 为什么需要Rete算法

假设我们有一个事实如下

```go
type Fact struct {
    StringValue string
}

func (f *Fact) VeryHeavyAndLongFunction() bool {
    ...
}
```

我们把事实添加到数据上下文中：

```go
f := &Fact{}
dctx := context.NewDataContext()
err := dctx.Add("Fact", f)
```

我们的GRL如下：

```go
rule ... {
    when
        Fact.VeryHeavyAndLongFunction() && Fact.StringValue == "Fish"
    then
        ...
}

rule ... {
    when
        Fact.VeryHeavyAndLongFunction() && Fact.StringValue == "Bird"
    then
        ...
}

rule ... {
    when
        Fact.VeryHeavyAndLongFunction() && Fact.StringValue == "Mammal"
    then
        ...
}

// and many similar rules

rule ... {
    when
        Fact.VeryHeavyAndLongFunction() && Fact.StringValue == "Insect"
    then
        ...
}
```

执行上面的GRL将会杀掉引擎，因为当它要选择要执行的规则，引擎会调用每一个规则里面的`when`里面对应的`Fact.VeryHeavyAndLongFunction` 函数。

因此，替代在评估每个规则的时候都去执行 `Fact.VeryHeavyAndLongFunction` ,  Rete 算法将会评估他们一次（即在第一次调用这个函数的时候）, 然后它会记住结果，在剩余的规则中可以直接拿来使用。（注意这就意味着你的函数是参考透明的，比如他不会产生副作用）。

对`Fact.StringValue`也是 一样的. Rete 算法将会从对象实例加载这个值，然后记住它，直到它在`then`里面被改变。 比如

```go
rule ... {
    when
        ...
    then
        Fact.StringValue = "something else";
}
```

### Grule工作内存里面有什么

Grule 将会尝试保存所有在KnowledgeBase的规则`when`中的  `Expression` 元素 。

首先，会努力尝试确实AST (Abstract Syntax Tree) 节点不会重复。

其次，每个语法树节点只会被评估一次，直到相关变量被修改。举例：

```Shell
    when
        Fact.A == Fact.B + Fact.Func(Fact.C) - 20
```

条件将会被解析成如下的表达式`Expression`.

```Shell
Expression "Fact.A" --> A variable
Expression "Fact.B" --> A variable
Expression "Fact.C" --> A variable
Expression "Fact.Func(Fact.C)"" --> A function containing argument Fact.C
Expression "20" --> A constant
Expression "Fact.B + Fact.Func(Fact.C)" --> A math operation contains 2 variable; Fact.B and Fact.C
Expression "(Fact.B + Fact.Func(Fact.C))" - 20 -- A math operation also contains 2 variable.
```

上述每个 `Expression`的结果值都会在第一次调用的时候被保存到内存中，这样后续的访问都避免了重复调用，只需要返回保存好的值就可以了。

如果其中的一个值在`then`中被修改了，比如：

```Shell
    then
        Fact.B = Fact.A * 20
```

然后所有包含 `Fact.B`的表达式都将从工作内存中移除：

```Shell
Expression "Fact.B"
Expression "Fact.B + Fact.Func(Fact.C)" --> A math operation contains 2 variable; Fact.B and Fact.C
Expression "(Fact.B + Fact.Func(Fact.C))" - 20 -- A math operation also contains 2 variable. 
```

这些表达式 `Expression`都将从工作内存中移除，方便他们在下一循环中被再次评估。

### 已知的 RETE 关于函数的问题

尽管Grule将会记住在`when`和`then`中要评估的每个变量，如果你在引擎之外修改了这个变量的值，比如在函数调用中修改了这个变量，Grule将看不到这个修改。最终将会导致，Grule会错误使用旧的变量的值，因为它不知道这个变量被修改了。你应该尝试保证你的函数是参考透明的以保证永远不会处理这种问题。

考虑以下的事实:

```go
type Fact struct {
    StringValue string
}

func (f *Fact) SetStringValue(newValue string) {
    f.StringValue = newValue
}
```

当你实例化你的事实，然后添加到数据上下文中:

```go
f := &Fact{
    StringValue: "One",
}
dctx := context.NewDataContext()
err := dctx.Add("Fact", f)
```

在你的GRL中，你可能会这么做：

```go
rule one "One" {
    when
        Fact.StringValue == "One"
        // Here grule remembers that Fact.StringValue value is "One"
    then
        Fact.SetStringValue("Two");
        // Here grule does not know that Fact.StringValue has changed inside the function.
        // What grule know is Fact.StringValue is still "One".
}

rule two "Two" {
    when
        Fact.StringValue == "Two"
        // Because of that, this will never evaluated true.
    then
        Fact.SetStringValue("Three");
}
```

因此引擎虽然不会报错，但是 `Fact.StringValue` 预期的应该是 `Two`，但是实际不是。

为了克服这个问题，如果变量发生了变化，你应该使用调用`Changed`函数告诉grule这个变量发生了变化。

```go
rule one "One" {
    when 
        Fact.StringValue == "One"
        // here grule remember that Fact.StringValue value is "One"
    then
        Fact.SetStringValue("Two");
        // here grule does not know if Fact.StringValue has changed inside the function.
        // What grule know is Fact.StringValue is still "One"

        // We should tell Grule that the variable changed within the Fact
        Changed("Fact.StringValue")
}
```
