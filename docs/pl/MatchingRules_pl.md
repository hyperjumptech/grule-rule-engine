## Wyszukaj reguły dopasowania (kolejność według atrakcyjności)

[![MatchingRules_cn](https://github.com/yammadev/flag-icons/blob/master/png/CN.png?raw=true)](../cn/MatchingRules_cn.md)
[![MatchingRules_de](https://github.com/yammadev/flag-icons/blob/master/png/DE.png?raw=true)](../de/MatchingRules_de.md)
[![MatchingRules_en](https://github.com/yammadev/flag-icons/blob/master/png/GB.png?raw=true)](../en/MatchingRules_en.md)
[![MatchingRules_id](https://github.com/yammadev/flag-icons/blob/master/png/ID.png?raw=true)](../id/MatchingRules_id.md)
[![MatchingRules_pl](https://github.com/yammadev/flag-icons/blob/master/png/PL.png?raw=true)](../pl/MatchingRules_pl.md)

`FetchMatchingRules` w `GruleEngine.go` pobiera wszystkie reguły ważne dla danego faktu i zwraca listę wartości `ast.RuleEntry` uporządkowanych według właściwości salience.

##### Zasady:

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

Wszystkie powyższe reguły są duplikatami z wyjątkiem reguły `UniqueRule5`, która ma inną wartość ważności.  Jak wiadomo, reguła o wyższej ważności ma wyższy priorytet i zostanie wykonana przed regułą o niższej ważności, jeśli wystąpi konflikt.

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

#### wynik:

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
