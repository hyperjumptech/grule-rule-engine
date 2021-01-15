## Fetch Matching Rules (Order by Salience)

`FetchMatchingRules` in `GruleEngine.go` fetches all the rules valid for a given fact and returns a list of `ast.RuleEntry` values ordered by salience property.

##### Rules:

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

All the above rules are duplicate ones except for a differing salience value, except `UniqueRule5`.  As we all know, a rule with higher salience has a higher priority and will get executed before a rule with lower salience if there is a conflict.

```go
fact := &Fact{
		Distance: 6000,
		Duration: 121,
}
```

##### FetchMatchingRules:

```go
engine := engine.NewGruleEngine()
ruleEntries, err := engine.FetchMatchingRules(dctx, kb)
if err != nil {
    panic(err)
}
```

#### result:

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
