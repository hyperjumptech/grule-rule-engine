# JSON 事实

[![JSON_Fact_cn](https://github.com/yammadev/flag-icons/blob/master/png/CN.png?raw=true)](../cn/JSON_Fact_cn.md)
[![JSON_Fact_de](https://github.com/yammadev/flag-icons/blob/master/png/DE.png?raw=true)](../de/JSON_Fact_de.md)
[![JSON_Fact_en](https://github.com/yammadev/flag-icons/blob/master/png/GB.png?raw=true)](../en/JSON_Fact_en.md)
[![JSON_Fact_id](https://github.com/yammadev/flag-icons/blob/master/png/ID.png?raw=true)](../id/JSON_Fact_id.md)
[![JSON_Fact_pl](https://github.com/yammadev/flag-icons/blob/master/png/PL.png?raw=true)](../pl/JSON_Fact_pl.md)

[About](About_cn.md) | [Tutorial](Tutorial_cn.md) | [Rule Engine](RuleEngine_cn.md) | [GRL](GRL_cn.md) | [GRL JSON](GRL_JSON_cn.md) | [RETE Algorithm](RETE_cn.md) | [Functions](Function_cn.md) | [FAQ](JSON_Fact_cn.md) | [Benchmark](Benchmarking_cn.md)

---

 从1.8.0版本开始，Grule支持使用JSON数据作为事实。能够使用户使用JSON格式表述他们的事实，并且能够像正常代码一样把这些事实加载到`DataContext`中。加载的JSON事实在Grule脚本现在是可见的了。

## 添加JSON事实

假设我们的JSON格式如下

```json
{
  "name" : "John Doe",
  "age" : 24,
  "gender" : "M",
  "height" : 74.8,
  "married" : false,
  "address" : {
    "street" : "9886 2nd St.",
    "city" : "Carpentersville",
    "state" : "Illinois",
    "postal" : 60110
  },
  "friends" : [ "Roth", "Jane", "Jake" ]
}
```

你可以把JSON存在字节数组中

```go
myJSON := []byte (...your JSON here...)
```

然后简单地添加JSON变量到 `DataContext`

```go
// create new instance of DataContext
dataContext := ast.NewDataContext()

// add your JSON Fact into data context using AddJSON() function.
err := dataContext.AddJSON("MyJSON", myJSON)
```

是的，你可以添加很多事实到上下文中，而且你可以混合JSON事实（使用AddJSON）和正常的Go事实（使用Add）。

## 在GRL中评估（读）JSON事实的值

在GRL较本周，当你添加到 `DataContext`，通过你提供的标签，事实是可见的。比如下面的代码是添加你的JSON，而且被打上`MyJSON`标签。

 ```go
err := dataContext.AddJSON("MyJSON", myJSON)
 ```

是的，你可以使用任何一个标签，只要他是一个词。

### 像正常对象一样访问成员变量

正如使用开头展示的JSON，你的GRL `when` 范围可以如下评估你的json。

 ```text
when
    MyJSON.name == "John Doe"
 ```

或者 

```text
when
    MyJSON.address.city.StrContains("ville")
```

或者

```text
when
    MyJSON.age > 30 && MyJSON.height < 60
```

### 像map一样访问成员变量

你可以像使用`Map`一样访问JSON对象，也可以像正常对象一样，也可以两者混用

 ```text
when
    MyJSON["name"] == "John Doe"
 ```

或者

```text
when
    MyJSON["address"].city.StrContains("ville")
```

或者

```text
when
    MyJSON.age > 30 && MyJSON["HEIGHT".ToLower()] < 60
```

### 访问数组成员变量

你可以像正常的数组一样访问JSON 数组元素。

 ```text
when
    MyJSON.friends[3] == "Jake"
 ```

## 在GRL中写入JSON事实

是的,你可以在`then`范围内写入你的值到JSON事实。这些变动的值将会在接下来的循环中被访问到。但是，有一些警告（参考下面的 你应该知道的）。 

### 像正常对象写入成员变量

正如开头展示的JSON，你的GRL `then`范围可以像如下修改你的试事实。

 ```text
then
    MyJSON.name = "Robert Woo";
 ```

或者

```text
then
    MyJSON.address.city = "Corruscant";
```

或者

```text
then
    MyJSON.age = 30;
```

这个是完美的直接方式。但是有一些别扭。

1. 你可以修改不仅仅是你JSON对象的成员变量，你可以修改类型。假设你的规则在接下来的评估过程中可以处理新的类型，否则强烈建议不要这么做。
   
   例子:
   
   你可以改变`MyJSON.age` 到 字符串.
   
   ```text
    then
        MyJSON.age = "Thirty";
   ```
   
   这个修改在遇到下面的规则时，会导致引擎panic。
   
   ```text
    when
        myJSON.age > 25
   ```
   
2. 你可以赋值给一个不存在的成员变量。

   比如:
   
      ```text
       then
           MyJSON.category = "FAT";
      ```

    其中  `category` 变量在原始JSON中不存在。
   
### 像map一样写入成员变量

正如开始所展示的JSON，在你的`then`范围内，你可以如下修改你的json事实。

 ```text
then
    MyJSON["name"] = "Robert Woo";
 ```

或者 

```text
then
    MyJSON["address"]["city"] = "Corruscant";
```

或者

```text
then
    MyJSON["age"] = 30;
```

正如对象格式的，同样适用。

1. 你可以修改不仅仅是你JSON对象的成员变量，你可以修改类型。假设你的规则在接下来的评估过程中可以处理新的类型，否则强烈建议不要这么做。
   
   举例
   
   你可以修改 `MyJSON.age` 到 字符串.
   
   ```text
    then
        MyJSON["age"] = "Thirty";
   ```
   
   这个修改在遇到下面的规则时，会导致引擎panic。
   
   ```text
    when
        myJSON.age > 25
   ```
   
2. 你可以赋值给一个不存在的成员变量。

   举例:
   
      ```text
       then
           MyJSON["category"] = "FAT";
      ```

   其中  `category` 变量在原始JSON中不存在。

### 写入数组成员变量

你可以使用索引修改一个数组元素。

```text
then
   MyJSON.friends[3] == "Jake";
```

指定的索引必须有效。如果越界，Grule将会panic。正如正常的JSON，你可以使用其他类型去替换任意元素。你也可以获取数组的长度。

```text
when
   MyJSON.friends.Length() > 4;
```

你可以通过 `Append` 函数往数组里面添加元素.  Append 也可以添加一个不同类型的元素到数组中。 (同样的警示 适用于 w.r.t. 修改指定类型)

```text
then
   MyJSON.friends.Append("Rubby", "Anderson", "Smith", 12.3);
```

**已知问题**

没有内置函数可以很容易地帮助使用者检查数组内容，比如Contains(value) bool。

## 你应该知道的

1. 在你添加一个JSON事实到 `DataContext`, JSON串的修改将不会影响在 `DataContext`中的事实。反方向同样适用， 在 `DataContext`中的修改将不会改变JSON串。
2. 你可以在`then`范围内修改你的JSON事实，但是不像正常的Go事实，这些修改将不会影响你的原始JSON字符串。如果你想实现这样的功能，你可以提前解析JSON到一个`struct` ，然后可以像正常一样添加你的`struct` 到`DataContext`。
