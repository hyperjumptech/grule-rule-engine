[![Gopher Holds The Rules](https://github.com/hyperjumptech/grule-rule-engine/blob/master/gopher-grule.png?raw=true)](https://github.com/hyperjumptech/grule-rule-engine/blob/master/gopher-grule.png?raw=true)


__"Gopher 遵守规则"__



[![About_cn](https://github.com/yammadev/flag-icons/blob/master/png/CN.png?raw=true)](../cn/About_cn.md)
[![About_de](https://github.com/yammadev/flag-icons/blob/master/png/DE.png?raw=true)](../de/About_de.md)
[![About_en](https://github.com/yammadev/flag-icons/blob/master/png/GB.png?raw=true)](../en/About_en.md)
[![About_id](https://github.com/yammadev/flag-icons/blob/master/png/ID.png?raw=true)](../id/About_id.md)

[About](About_cn.md) | [Tutorial](Tutorial_cn.md) | [Rule Engine](RuleEngine_cn.md) | [GRL](GRL_cn.md) | [GRL JSON](GRL_JSON_cn.md) | [RETE Algorithm](RETE_cn.md) | [Functions](Function_cn.md) | [FAQ](FAQ_cn.md) | [Benchmark](Benchmarking_cn.md)

# Grule

```go
import "github.com/hyperjumptech/grule-rule-engine"
```

## Go 规则引擎

Grule是Golang实现的规则引擎库。受业内称赞的JBOSS Drools的启发，我们实现了一种简单的规则引擎。

正如**Drools**,**Grule**也有自己的*DSL*，对比如下。

Drool的DRL如下：

```go
rule "SpeedUp"
    salience 10
    when
        $TestCar : TestCarClass( speedUp == true && speed < maxSpeed )
        $DistanceRecord : DistanceRecordClass()
    then
        $TestCar.setSpeed($TestCar.Speed + $TestCar.SpeedIncrement);
        update($TestCar);
        $DistanceRecord.setTotalDistance($DistanceRecord.getTotalDistance() + $TestCar.Speed)
        update($DistanceRecord)
end
```

同时Grule的GRL如下：

```go
rule SpeedUp "When testcar is speeding up we increase the speed." salience 10  {
    when
        TestCar.SpeedUp == true && TestCar.Speed < TestCar.MaxSpeed
    then
        TestCar.Speed = TestCar.Speed + TestCar.SpeedIncrement;
        DistanceRecord.TotalDistance = DistanceRecord.TotalDistance + TestCar.Speed;
}
```

# 什么是 RuleEngine



对于 Martin Fowler的文章，这不是一个比较好的诠释。你可以看原文[RulesEngine by Martin Fowler](https://martinfowler.com/bliki/RulesEngine.html)。

以下内容摘自 **TutorialsPoint** 网站（做了一些小改动）

**Grule**规则引擎是一种生产规则系统，使用了基于规则的方案去实现一个专家系统。专家系统是一种基于知识的系统，可以使用知识描述将获取的知识处理成可以用来推理的知识库。

生产规则系统是图灵完备的，专注于使用知识描述以一种简洁、明确和声明式的方式去表达一阶的命题逻辑。

生产规则系统的大脑是推理引擎，可以扩展到大量的规则(rule)和事实(fact)。推理引擎用事实和数据去匹配生产规则（也称为生产或者规则），从而推理出一个可以产生行动的结论。

生产规则系统是一个两部分结构，使用了一阶逻辑对知识描述进行一个推理。业务规则引擎是在运行时生产环境中执行一个或多个业务规则的软件系统。

规则引擎允许你定义 做什么(what)而不是怎么做(How)。

## 什么是规则(Rule)

规则是知识片段，经常被描述成 "当满足某些条件时，然后做一些动作"。

```go
When
   <Condition is true>
Then
   <Take desired Action>
```

规则最重要的部分是`when`部分。如果`when`中的条件被满足了，则可以触发`then`操作。

```go
rule  <rule_name> <rule_description>
   <attribute> <value> {
   when
      <conditions>

   then
      <actions>
}
```

## 规则引擎的优势

### 声明式编程

规则可以很容易地表达对困难问题的解决方案并获得验证。不像代码，规则可以使用不复杂的语言描述。业务分析师可以轻松阅读和验证一组规则。

### 逻辑与数据分离

数据驻留在域对象中，业务逻辑驻留在规则中。 根据项目的类型，这种分离可能非常有利。

### 知识集中化

通过使用规则，您可以创建一个可执行的知识库（知识库）。 这是商业政策的一个真理。 理想情况下，规则具有可读性，它们也可以用作文档。

### 变化敏捷

由于业务规则实际上被视为数据。 根据业务动态性质调整规则变得很容易。 无需像普通软件开发那样重新构建代码、部署，您只需要推出规则集并将它们应用到知识库中。

