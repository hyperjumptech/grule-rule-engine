# Benchmarking

[Tutorial](Tutorial_en.md) | [Rule Engine](RuleEngine_en.md) | [GRL](GRL_en.md) | [GRL JSON](GRL_JSON_en.md) | [RETE Algorithm](RETE_en.md) | [Functions](Function_en.md) | [FAQ](FAQ_en.md) | [Benchmark](Benchmarking_en.md)

---

For any library `Benchmarking` is very much required in terms of evaluating the performance and helps us to improve the performance better.

I have benchmarked two things:
* Loading 100 and 1000 rules into KnowledgeBase
* Executing a fact against rules against 100 and 1000 rules

All the tests will run for multiple times to see how it is performing by changing the value of N 

`N is b.N where b is an instance of *testing.B`

### Rules:
refer `100_rules.grl` and `1000_rules.grl` files under `examples/benchmark` directory

Command to run: 
---
```go
> go test -bench=. -benchmem
goos: darwin
goarch: amd64
Number of Cores - 6
Ram - 16 GB 2400 MHz DDR4
pkg: github.com/hyperjumptech/grule-rule-engine/examples/benchmark
```

Result 
-- 

### Test1 - Loading Rules into KnowledgeBase
Result:
```go
Load 100 and 1000 Rules into Grule rule Engine
Benchmark_Grule_Load_Rules/100_rules-12                    67137             17387 ns/op            5071 B/op        118 allocs/op
Benchmark_Grule_Load_Rules/100_rules#01-12                 67485             17447 ns/op            5071 B/op        118 allocs/op
Benchmark_Grule_Load_Rules/100_rules#02-12                 67332             17408 ns/op            5071 B/op        118 allocs/op
Benchmark_Grule_Load_Rules/100_rules#03-12                 67992             17436 ns/op            5071 B/op        118 allocs/op
Benchmark_Grule_Load_Rules/100_rules#04-12                 68170             17420 ns/op            5071 B/op        118 allocs/op
Benchmark_Grule_Load_Rules/100_rules#05-12                 67777             17645 ns/op            5071 B/op        118 allocs/op
Benchmark_Grule_Load_Rules/100_rules#06-12                 65100             17431 ns/op            5071 B/op        118 allocs/op
Benchmark_Grule_Load_Rules/100_rules#07-12                 67396             17396 ns/op            5071 B/op        118 allocs/op
Benchmark_Grule_Load_Rules/100_rules#08-12                 68132             17458 ns/op            5071 B/op        118 allocs/op
Benchmark_Grule_Load_Rules/100_rules#09-12                 67881             17399 ns/op            5071 B/op        118 allocs/op
Benchmark_Grule_Load_Rules/100_rules#10-12                 67216             17523 ns/op            5071 B/op        118 allocs/op
Benchmark_Grule_Load_Rules/1000_rules-12                   66828             17823 ns/op            5071 B/op        118 allocs/op
Benchmark_Grule_Load_Rules/1000_rules#01-12                69122             17581 ns/op            5071 B/op        118 allocs/op
Benchmark_Grule_Load_Rules/1000_rules#02-12                67815             17425 ns/op            5071 B/op        118 allocs/op
Benchmark_Grule_Load_Rules/1000_rules#03-12                67405             19681 ns/op            5070 B/op        118 allocs/op
Benchmark_Grule_Load_Rules/1000_rules#04-12                48511             21222 ns/op            5071 B/op        118 allocs/op
Benchmark_Grule_Load_Rules/1000_rules#05-12                67779             18999 ns/op            5070 B/op        118 allocs/op
Benchmark_Grule_Load_Rules/1000_rules#06-12                56694             17691 ns/op            5071 B/op        118 allocs/op
Benchmark_Grule_Load_Rules/1000_rules#07-12                69086             17641 ns/op            5071 B/op        118 allocs/op
Benchmark_Grule_Load_Rules/1000_rules#08-12                51638             19401 ns/op            5070 B/op        118 allocs/op
Benchmark_Grule_Load_Rules/1000_rules#09-12                58940             20498 ns/op            5071 B/op        118 allocs/op
Benchmark_Grule_Load_Rules/1000_rules#10-12                67411             19487 ns/op            5071 B/op        118 allocs/op


```

To load `100` rules into knowledgeBase it took `17523 ns/op` (took the highest value) that is equal to `~0.017523ms` and (`5071 B/op`) `0.005071MB` memory per operation

To load `1000` rules into knowledgeBase it took `21222 ns/op` (took the highest value) that is equal to `~0.021222ms` and (`5071 B/op`) `0.005071MB` memory per operation

### Test2 - Executing a fact against rules
Result:
```go
Executing a fact against 100 and 1000 rules
Benchmark_Grule_Execution_Engine/100_rules-12            2055945               574 ns/op             512 B/op          9 allocs/op
Benchmark_Grule_Execution_Engine/100_rules#01-12         2048078               570 ns/op             512 B/op          9 allocs/op
Benchmark_Grule_Execution_Engine/100_rules#02-12         2086953               572 ns/op             512 B/op          9 allocs/op
Benchmark_Grule_Execution_Engine/100_rules#03-12         2094231               571 ns/op             512 B/op          9 allocs/op
Benchmark_Grule_Execution_Engine/100_rules#04-12         2078065               576 ns/op             512 B/op          9 allocs/op
Benchmark_Grule_Execution_Engine/100_rules#05-12         2028356               642 ns/op             512 B/op          9 allocs/op
Benchmark_Grule_Execution_Engine/100_rules#06-12         2002248               628 ns/op             512 B/op          9 allocs/op
Benchmark_Grule_Execution_Engine/100_rules#07-12         1850121               703 ns/op             512 B/op          9 allocs/op
Benchmark_Grule_Execution_Engine/100_rules#08-12         1761343               585 ns/op             512 B/op          9 allocs/op
Benchmark_Grule_Execution_Engine/100_rules#09-12         2080953               594 ns/op             512 B/op          9 allocs/op
Benchmark_Grule_Execution_Engine/100_rules#10-12         2082880               573 ns/op             512 B/op          9 allocs/op
Benchmark_Grule_Execution_Engine/1000_rules-12           2082183               575 ns/op             512 B/op          9 allocs/op
Benchmark_Grule_Execution_Engine/1000_rules#01-12        2098585               568 ns/op             512 B/op          9 allocs/op
Benchmark_Grule_Execution_Engine/1000_rules#02-12        2090640               570 ns/op             512 B/op          9 allocs/op
Benchmark_Grule_Execution_Engine/1000_rules#03-12        2109938               587 ns/op             512 B/op          9 allocs/op
Benchmark_Grule_Execution_Engine/1000_rules#04-12        2045216               576 ns/op             512 B/op          9 allocs/op
Benchmark_Grule_Execution_Engine/1000_rules#05-12        2092534               575 ns/op             512 B/op          9 allocs/op
Benchmark_Grule_Execution_Engine/1000_rules#06-12        1994415               579 ns/op             512 B/op          9 allocs/op
Benchmark_Grule_Execution_Engine/1000_rules#07-12        2098788               599 ns/op             512 B/op          9 allocs/op
Benchmark_Grule_Execution_Engine/1000_rules#08-12        2092808               573 ns/op             512 B/op          9 allocs/op
Benchmark_Grule_Execution_Engine/1000_rules#09-12        2085716               609 ns/op             512 B/op          9 allocs/op
Benchmark_Grule_Execution_Engine/1000_rules#10-12        1864302               576 ns/op             512 B/op          9 allocs/op

```

To execute a fact against 100 rules, Grule Engine took `~703 ns/op` (took the highest value as base) that is hardly `~0.000703ms` and (`512 B/op`) `0.000512MB` memory per operation which is pretty fast.

To execute a fact against 1000 rules, Grule Engine took `~587 ns/op` (took the highest value as base) that is hardly `~0.000587ms` and (`512 B/op`) `0.000512MB` memory per operation  which is also pretty fast.


