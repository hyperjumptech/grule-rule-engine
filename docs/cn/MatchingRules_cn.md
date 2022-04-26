## 获取匹配的规则 (按优先级排序)

[![MatchingRules_cn](https://github.com/yammadev/flag-icons/blob/master/png/CN.png?raw=true)](../cn/MatchingRules_cn.md)
[![MatchingRules_de](https://github.com/yammadev/flag-icons/blob/master/png/DE.png?raw=true)](../de/MatchingRules_de.md)
[![MatchingRules_en](https://github.com/yammadev/flag-icons/blob/master/png/GB.png?raw=true)](../en/MatchingRules_en.md)
[![MatchingRules_id](https://github.com/yammadev/flag-icons/blob/master/png/ID.png?raw=true)](../id/MatchingRules_id.md)
[![MatchingRules_pl](https://github.com/yammadev/flag-icons/blob/master/png/PL.png?raw=true)](../pl/MatchingRules_pl.md)

`GruleEngine.go`中`FetchMatchingRules`函数的将会获取所有能够满足事实的所有有效规则，并返回按优先级排序的  `ast.RuleEntry` 列表。

##### 规则:

```go
rule DuplicateRule1 "Duplicate Rule 1" salience 5 {
    when
        (Fact.Distance > 5000 && Fact.Duration > 120) && (Fact.Result == false)
    Then
        Fact.NetAmount=143.320007;
        Fact.Result=true;
}

rule DuplicateRule2 "Duplicate Rule 2" salience 6 {
    when
        (Fact.Distance > 5000 && Fact.Duration > 120) && (Fact.Result == false)
    Then
        Fact.NetAmount=143.320007;
        Fact.Result=true;
}

rule DuplicateRule3 "Duplicate Rule 3" salience 7 {
    when
        (Fact.Distance > 5000 && Fact.Duration > 120) && (Fact.Result == false)
    Then
        Fact.NetAmount=143.320007;
        Fact.Result=true;
}

rule DuplicateRule4 "Duplicate Rule 4" salience 8 {
    when
        (Fact.Distance > 5000 && Fact.Duration > 120) && (Fact.Result == false)
    Then
        Fact.NetAmount=143.320007;
        Fact.Result=true;
}

rule UniqueRule5 "Unique Rule 5" salience 0 {
    when
        (Fact.Distance > 5000 && Fact.Duration == 120) && (Fact.Result == false)
    Then
        Output.NetAmount=143.320007;
        Fact.Result=true;
}
```

除了 `UniqueRule5`规则之外，所有以上的规则除了优先级不一样之外都是重复的.  正如我们所了解的，有更高优先级的规则将会在发生冲突的时候会被优先执行。

```go
fact := &Fact{
		Distance: 6000,
		Duration: 121,
}
```

##### 调用FetchMatchingRules:

```go
engine := engine.NewGruleEngine()
ruleEntries, err := engine.FetchMatchingRules(dctx, kb)
if err != nil {
    panic(err)
}
```

#### 返回结果:

```go
Returns []*ast.RuleEntry (All Matching Rule Entries sorted by Salience)

rule DuplicateRule4 "Duplicate Rule 4" salience 8 {
    when
        (Fact.Distance > 5000 && Fact.Duration > 120) && (Fact.Result == false)
    Then
        Fact.NetAmount=143.320007;
        Fact.Result=true;
}

rule DuplicateRule3 "Duplicate Rule 3" salience 7 {
    when
        (Fact.Distance > 5000 && Fact.Duration > 120) && (Fact.Result == false)
    Then
        Fact.NetAmount=143.320007;
        Fact.Result=true;
}

rule DuplicateRule2 "Duplicate Rule 2" salience 6 {
    when
        (Fact.Distance > 5000 && Fact.Duration > 120) && (Fact.Result == false)
    Then
        Fact.NetAmount=143.320007;
        Fact.Result=true;
}

rule DuplicateRule1 "Duplicate Rule 1" salience 5 {
    when
        (Fact.Distance > 5000 && Fact.Duration > 120) && (Fact.Result == false)
    Then
        Fact.NetAmount=143.320007;
        Fact.Result=true;
}
```
