# 经常问的问题

[![FAQ_cn](https://github.com/yammadev/flag-icons/blob/master/png/CN.png?raw=true)](../cn/FAQ_cn.md)
[![FAQ_de](https://github.com/yammadev/flag-icons/blob/master/png/DE.png?raw=true)](../de/FAQ_de.md)
[![FAQ_en](https://github.com/yammadev/flag-icons/blob/master/png/GB.png?raw=true)](../en/FAQ_en.md)
[![FAQ_id](https://github.com/yammadev/flag-icons/blob/master/png/ID.png?raw=true)](../id/FAQ_id.md)
[![FAQ_pl](https://github.com/yammadev/flag-icons/blob/master/png/PL.png?raw=true)](../pl/FAQ_pl.md)

[About](About_cn.md) | [Tutorial](Tutorial_cn.md) | [Rule Engine](RuleEngine_cn.md) | [GRL](GRL_cn.md) | [GRL JSON](GRL_JSON_cn.md) | [RETE Algorithm](RETE_cn.md) | [Functions](Function_cn.md) | [FAQ](FAQ_cn.md) | [Benchmark](Benchmarking_cn.md)

---

## 1. Grule 在最大循环引发Panic

**问题**: 在Grule引擎执行的时候获取的如下的panic信息。

```Shell
panic: GruleEngine successfully selected rule candidate for execution after 5000 cycles, this could possibly caused by rule entry(s) that keep added into execution pool but when executed it does not change any data in context. Please evaluate your rule entries "When" and "Then" scope. You can adjust the maximum cycle using GruleEngine.MaxCycle variable.
```

**回答**:  这个报错说明了你要评估的的规则有潜在的问题。Grule持续在内存中执行RETE网络，直到冲突集合中没有操作可以执行，这种情况我们叫做自然终止状态。如果你的规则不能使得RETE网络到达自然终止状态，程序将会永远执行下去。`GruleEngine.MaxCycle`的默认配置是`5000`，这会防止程序会无限循环执行下去。

如果你觉得你系统中的规则需要更多循环才能终止，你可以增加这个值。但是你不相信是这样的，你可能有一个非终止规则集。

考虑一下事实:

```go
type Fact struct {
   Payment int
   Cashback int
}
```

一下是定义好的规则:

```Shell
rule GiveCashback "Give cashback if payment is above 100" {
    When 
         F.Payment > 100
    Then
         F.Cashback = 10;
}

rule LogCashback "Emit log if cashback is given" {
    When 
         F.Cashback > 5
    Then
         Log("Cashback given :" + F.Cashback);
}
```

用如下的Fact实例去执行规则

```go
&Fact {
     Payment: 500,
}
```

将不会停止下来。 

```
Cycle 1: Execute "GiveCashback" .... because when F.Payment > 100 is a valid condition
Cycle 2: Execute "GiveCashback" .... because when F.Payment > 100 is a valid condition
Cycle 3: Execute "GiveCashback" .... because when F.Payment > 100 is a valid condition
...
Cycle 5000: Execute "GiveCashback" .... because when F.Payment > 100 is still a valid condition
panic
```

Grule会在同一个规则上一直执行，因为**When**条件将持续产生一个有效的结果。

一直有效解决方案是如下修改GiveCashback规则：

```Shell
rule GiveCashback "Give cashback if payment is above 100" {
    When 
         F.Payment > 100 &&
         F.Cashback == 0
    Then
         F.Cashback = 10;
}
```

`GiveCashback`这个规则定义将会考虑状态改变。初始`Cashback`值为0，但是操作将会这个值，在下一个循环中将不会匹配到这个规则，从而可以进入最终状态。

上述方法有点自然，因为他是规则条件管理了终止。但是，如果你不能以自然方式终止执行，可以在操作里面改变规则引擎的状态，如下：

```Shell
rule GiveCashback "Give cashback if payment is above 100" {
    When 
         F.Payment > 100
    Then
         F.Cashback = 10;
         Retract("GiveCashback");
}
```

`Retract`函数将会在下一个循环从knowledge base移除GiveCashback规则。因为它不在存在了，这个规则将不会在下个循环中被评估。主要注意的是，这只会在`Retract`调用后紧接着的信息中发生。在后续循环中将会重新引入调用。

---

## 2. 保存 Rule 入口到数据库

**问题**: 是否有计划将 Grule 集成到一个数据库?

**回答**: 没有. 尽管存储你的规则到一个数据库中是一个很好的想法，Grule将不会适配任何的数据库适配器去自动存储和获取规则。你可以很容易自己实现这样子的适配器，你可以实现Knowledgebase的通用接口：*Reader*, *File*, *Byte Array*, *String*
and *Git*。字符串也是很容易插入数据库和查询数据库，你可以加载他们到Grule的knowledgebase。

我们不想绑定Grule到任意一个特殊的数据库。

---

## 3. 在一个 knowledge-base最大的规则数

**问题**:  knowledgebase可以保存多少个规则?

**回答**: 你可以有任意多的规则，但是最少保证有一个。

---

## 4. 获取给定fact匹配到所有规则

**问题**: 给定facts，我怎么测试我的规则有效性?

**回答**: 你可以使用 `engine.FetchMatchingRule` 函数. 更多信息参考[Matching Rules Doc](MatchingRules_cn.md) 

---

## 5. Rule Engine use-case

**问题**: 我已经了解了规则引擎，但是我可以从中受益什么？可以举一些例子吗？

**回答**: 以我个人意见来看，以下的情况规则 引擎是一个比较好的解决方案。

1. 一个专家系统，必须通过评估事实，然后得到一些现实中的结论。如果不使用RETE格式的规则引擎，一个人将会写出来一堆`if/else`表达式和组合去很快评估一些事情，但是很难管理。一个基于表格的规则引擎可以很有效，但是依然很难修改，而且不太容易编码。一个像Grule一样的系统可以使你描述规则和你系统里的事实，把你从规则描述和事实分离开，隐藏了大量的复杂度。
2. 一个投票系统。比如，银行系统需要根据消费者的交易记录，对每个消费者创建一个积分。我们会根据他们与银行交易程度，买入卖出的金额，支付速度，利息收入等等，可以看到他们的分数变化。开发者提供的规则引擎，事实规范，规则，可以提供给主题专家去做银行的客户业务。分离这些团队，每个团队单独负责自己的业务。
3. 游戏。玩家的状态，酬金，处罚，伤害，分数和概率体统是几乎所有计算机游戏中规则发挥重要作用的不同例子。这些规则可以以一种复杂的方式相互影响，常常是开发者想象不到的方式。针对这种动态场景使用脚本语言（比如LUA）可能会变得特别复杂，规则引擎可以极大地有效简化这部分工作。
4. 分类系统。这实际是上面介绍的投票系统的一般化。使用规则引擎，我们可以划分为信用资格，生物识别特征，保险产品的风险评估，潜在的安全威胁等等。
5. Advice/Suggestion system. A "rule" is simply another kind of data, which
   makes it a prime candidate for definition by another program.  This program
   can be another expert system or artificial intelligence.  Rules can be
   manipulated by another system in order to deal with new types of facts or
   newly discovered information about the domain which the rule set is intending
   to model.建议系统。一个”规则“是一个简单的其他类型的数据，可以作为其他程序的基本定义。这个程序可以是一个专家系统，也可以是人工智能。规则可以由另一个系统操纵，以便处理新类型的事实或新发现的有关规则集打算建模的域的信息。

还有很多其他的可以从规则引擎中收益使用案例。上面的只是很少的一部分潜在代表。

无论如何，需要强调的是规则引擎不是银弹。很多其他已有的方案可以去解决知识问题，并在合适的时候采用这些替代方案。举例，当只需要简单的 `if` / `else`就足够的时候就不要使用规则引擎。

还有一些其他的事情需要注意：有一些规则引擎的实现是很昂贵的，但是很多企业从中获得了如此多的价值，以至于运行他们的成本很容易被该价值抵消。即使对于中等复杂的用例，强大的规则引擎可以解耦团队和业务，复杂性可以看起来更明显。

---

## 6. 日志

**问题**: Grule的日志太冗长了，我可以关掉Grule的日志吗?

**回答**: 可以的.通过增加日志级别， 你可以减少 (或者完全关闭) Grule日志.

```go
import (
    "github.com/kalyan-arepalle/grule-rule-engine/logger"
    "github.com/sirupsen/logrus"
)
...
...
logger.SetLogLevel(logrus.PanicLevel)
```

上面代码将会将Grule的日志级别设置成 `Panic` 级别, 只有当panic发生的时候才会记录日志.

当然，修改日志级别可以减少你debug系统的能力，所以我们建议在生产环境才使用更高级别的日志级别。
