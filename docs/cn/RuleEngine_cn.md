# 规则引擎简介

[![RuleEngine_cn](https://github.com/yammadev/flag-icons/blob/master/png/CN.png?raw=true)](../cn/RuleEngine_cn.md)
[![RuleEngine_de](https://github.com/yammadev/flag-icons/blob/master/png/DE.png?raw=true)](../de/RuleEngine_de.md)
[![RuleEngine_en](https://github.com/yammadev/flag-icons/blob/master/png/GB.png?raw=true)](../en/RuleEngine_en.md)
[![RuleEngine_id](https://github.com/yammadev/flag-icons/blob/master/png/ID.png?raw=true)](../id/RuleEngine_id.md)
[![RuleEngine_pl](https://github.com/yammadev/flag-icons/blob/master/png/PL.png?raw=true)](../pl/RuleEngine_pl.md)

[About](About_cn.md) | [Tutorial](Tutorial_cn.md) | [Rule Engine](RuleEngine_cn.md) | [GRL](GRL_cn.md) | [GRL JSON](GRL_JSON_cn.md) | [RETE Algorithm](RETE_cn.md) | [Functions](Function_cn.md) | [FAQ](RuleEngine_cn.md) | [Benchmark](Benchmarking_cn.md)

正如Martin Fowler 阐述的，规则引擎是计算模型的替代方案，而不是评估了多个条件，如果满足了，然后选择适当的操作。

你可以传给**KnowledgeBase**一些规则集合，然后引擎使用这些每条规则去评估一些**Facts**。如果一个规则满足了条件，将会进行该规则指定的操作。

## Fact 事实

`fact`是事实，虽然听上去有点点蠢，但是就是如此。存在规则引擎的上下文中的事实，是用来评估的基础信息。事实可以由多种来源，比如数据库，触发流程，销售系统一个点，报告等等。

举一个事实的例子可能更容易理解。假设我们有如下的一个事实

```Text
Purchase Transaction
    Item Name     : Computer Monitor
    Quantity      : 10
    Purchase Date : 12 Dec 2019
    Item Price    : 150 USD
    Total Price   : 1500 USD
    Tax           : ?
    Discount      : ?
    Final Price   : ?
```

事实基本上是任何信息或者收集的数据。

就Purchasing Fact这个例子来说，我们可以获取很多信息：要购买的是什么东西，数量，日期等等。但是我们不知道应该交多少税，我们可以拿到多少这块，以及购买者最终需要支付多少钱。

## Rule 规则

规则是一种关于如果评估**Fact**的规格。如果Fact可以满足规则的条件，规则指定的操作将会被执行。有时候，一个Fact可以满足多个规则，这将会导致冲突。在冲突中的规则集合被叫做冲突集合。为了解决冲突集合的问题，我们需要指定一种策略，将会在后面进行讨论。

回到我们这个简单的购物系统案例：制定一些商业规则去计算最终的价格，可能需要先计算税，然后是折扣。当税和折扣都知道了，我们就可以知道价格。

以下是我们使用伪代码制定的一些规则。

```text
Rule 1
   IF
   - the Item's Tax is not known AND
   - the Item's Name is "Computer CPU"
   THEN
   - Item's Tax is 10%

Rule 2
   IF
   - the Item's Tax is not known AND
   - the Item's Name is "Computer Monitor"
   THEN
   - Item's Tax is 7%

Rule 3
   IF
   - the Item's Discount is not known AND
   - the Item's Price After Tax is Less Than 1000 USD
   THEN
   - Item's Discount is 0%

Rule 4
   IF
   - the Item's Discount is not known AND
   - the Item's Price After Tax is Less Than 1500 USD AND
   - the Item's Price After Tax is Greater Than or Equal To 1000 USD
   THEN
   - Item's Discount is 3%

Rule 5
   IF
   - the Item's Discount is not known AND
   - the Item's Price After Tax is Less Than 2000 USD AND
   - the Item's Price After Tax is Greater Than or Equal To 1500 USD
   THEN
   - Item's Discount is 5%

Rule 6
   IF
   - the Item's Discount is not known AND
   - the Item's Price After Tax is More Than 2000 USD
   THEN
   - Item's Discount is 10%

Rule 7
   IF
   - the Item's Total Price is known AND
   - the Item's Discount is known AND
   - the Item's Tax is known AND
   - the Item's Final Price is not known
   THEN
   - Item's Final Price is calculate price from Total Price
     with given Tax and Discount
```

如果你仔细检查以上的规则，你可以轻松的理解规则引擎中的规则这个概念。这些规则集合将会组成知识。在这个案例中，就组成了一个知识集合叫做**如果计算商品的最终价格**。

## Cycle 循环

一个规则评估循环从评估每个规则的条件(the **IFs**)去选择可能要执行的规则开始。每一次引擎找到了一个满足的要求，不是立马执行满足条件的规则的操作，而是将这个规则加入到规则候选集合，也叫冲突集合。

当所有的规则要求都被评估过了，是立即执行满足条件的规则的操作吗？这依赖于冲突集合的内容。

* 如果没有规则能够满足**IF**条件，引擎将会立即结束。
* 如果冲突集合中只有一个规则，在结束之前将会执行这个规则的操作。
* 如果冲突集合中有多个规则，引擎将会采取一种策略去优先选择一个规则，然后执行他的操作。

如果一个操作被执行，只要有一个操作需要执行，循环将会再次重复。如果没有操作可以执行，说明fact满足不了任何规则了，即匹配不上IF了。循环将会结束, 规则引擎将会完成评估。

冲突解决策略的伪代码如下所示：

```text
Start Engine With a FACT Using a KNOWLEDGE
BEGIN
    For Every RULE in KNOWLEDGE
        Check if RULE's Requirement is Satisfied by FACT
            If RULE's Requirement is Satisfied
                Add RULE into CONFLICT SET
            End If
        End Check
    End For
    If CONFLICT SET is EMPTY
        Finished
        END
    If CONFLICT SET Has 1 RULE
        Execute the RULE's Action
        Clear CONFLICT SET
        Repeat Cycle from BEGIN
    If CONFLICT SET has Many RULEs
        Apply Conflict Resolution Strategy to Choose 1 RULE.
        Execute the Chosen RULE's Action
        Clear CONFLICT SET
        Repeat Cycle from BEGIN
END

```

Grule将会记录在一次评估中运行了多少个循环。如果一个规则评估和执行重复了很多次，超过了实例化Grule引擎实例时指定的数据量，引擎将终止并返回错误。

##  冲突集合解决策略

正如上面描述的，规则引擎将会评估所有规则的要求，然后将冲突规则放到一个列表中，叫冲突集合。如果集合中只有一个规则，意味着不会和其他规则冲突。引擎将会立即执行这个规则的操作。

如果集合中有多个规则，将会产生冲突。有很多种冲突解决策略。最简单的一种是通过指定规则的优先级**salience**也可以叫priority，importance。

如以下伪代码所示，我们可以给一个规则定义一个优先级salience.

```text
Rule 1 - Priority 1
   IF
   - the Item's Tax is not known AND
   - the Item's Name is "Computer CPU"
   THEN
   - Item's Tax is 10%

Rule 2 - Priority 10
   IF
   - the Item's Tax is not known AND
   - the Item's Name is "Computer Monitor"
   THEN
   - Item's Tax is 7%
```

如果不指定，规则的优先级为0.

因为所有不指定的规则优先级都为0，引擎很容易从冲突集合中选取一个规则去执行。如果有多个规则满足了优先级，则选择第一个。因为Go的map是无序的，不能保证按照输入的顺序进行排序，所以假设会按照输入顺序去执行规则是不安全的。

Grule中的规则的优先级可以是负数，意思是比默认的优先级还低。这将会保证在所有其他规则都执行玩之后，这个规则会最后执行。
