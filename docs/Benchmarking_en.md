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

Benchmark_Grule_Load_Rules/100_rules-12                       60          20968700 ns/op         8871574 B/op     216554 allocs/op
Benchmark_Grule_Load_Rules/100_rules#01-12                    60          20800060 ns/op         8871255 B/op     216552 allocs/op
Benchmark_Grule_Load_Rules/100_rules#02-12                    60          21284699 ns/op         8871410 B/op     216553 allocs/op
Benchmark_Grule_Load_Rules/100_rules#03-12                    61          20414968 ns/op         8871317 B/op     216552 allocs/op
Benchmark_Grule_Load_Rules/100_rules#04-12                    58          20618596 ns/op         8871612 B/op     216554 allocs/op
Benchmark_Grule_Load_Rules/100_rules#05-12                    60          21217303 ns/op         8871294 B/op     216552 allocs/op
Benchmark_Grule_Load_Rules/100_rules#06-12                    67          21312189 ns/op         8871592 B/op     216554 allocs/op
Benchmark_Grule_Load_Rules/100_rules#07-12                    61          20592475 ns/op         8871213 B/op     216552 allocs/op
Benchmark_Grule_Load_Rules/100_rules#08-12                    60          22628754 ns/op         8871388 B/op     216553 allocs/op
Benchmark_Grule_Load_Rules/100_rules#09-12                    68          21192157 ns/op         8871223 B/op     216552 allocs/op
Benchmark_Grule_Load_Rules/100_rules#10-12                    60          21242572 ns/op         8871226 B/op     216552 allocs/op

Benchmark_Grule_Load_Rules/1000_rules-12                       6         209761389 ns/op        88641262 B/op    2141287 allocs/op
Benchmark_Grule_Load_Rules/1000_rules#01-12                    6         204268674 ns/op        88644670 B/op    2141304 allocs/op
Benchmark_Grule_Load_Rules/1000_rules#02-12                    6         210895687 ns/op        88639476 B/op    2141278 allocs/op
Benchmark_Grule_Load_Rules/1000_rules#03-12                    6         214102248 ns/op        88642209 B/op    2141293 allocs/op
Benchmark_Grule_Load_Rules/1000_rules#04-12                    5         268977045 ns/op        88639793 B/op    2141279 allocs/op
Benchmark_Grule_Load_Rules/1000_rules#05-12                    5         211837045 ns/op        88641822 B/op    2141289 allocs/op
Benchmark_Grule_Load_Rules/1000_rules#06-12                    6         221863753 ns/op        88642209 B/op    2141293 allocs/op
Benchmark_Grule_Load_Rules/1000_rules#07-12                    6         223676073 ns/op        88643585 B/op    2141299 allocs/op
Benchmark_Grule_Load_Rules/1000_rules#08-12                    6         224317362 ns/op        88643070 B/op    2141297 allocs/op
Benchmark_Grule_Load_Rules/1000_rules#09-12                    5         241930711 ns/op        88641422 B/op    2141289 allocs/op
Benchmark_Grule_Load_Rules/1000_rules#10-12                    4         261857403 ns/op        88637592 B/op    2141269 allocs/op

```

To load `100` rules into knowledgeBase it took `22628754 ns/op` (took the highest value) that is equal to `22.628754ms` and (`8871388 B/op`) `8.8MB` memory per operation

To load `1000` rules into knowledgeBase it took `261857403 ns/op` (took the highest value) that is equal to `~261ms` and `88MB` memory per operation

### Test2 - Executing a fact against rules
Command to run: 
```go
Load 100 and 1000 Rules into Grule rule Engine
Benchmark_Grule_Load_Rules/100_rules-12                       12          96674568 ns/op        49297966 B/op     731119 allocs/op
Benchmark_Grule_Load_Rules/100_rules#01-12                    12          97915910 ns/op        49293839 B/op     731103 allocs/op
Benchmark_Grule_Load_Rules/100_rules#02-12                    12          97716674 ns/op        49293398 B/op     731129 allocs/op
Benchmark_Grule_Load_Rules/100_rules#03-12                    12          97227219 ns/op        49299542 B/op     731145 allocs/op
Benchmark_Grule_Load_Rules/100_rules#04-12                    12          99342047 ns/op        49295906 B/op     731131 allocs/op
Benchmark_Grule_Load_Rules/100_rules#05-12                    12          98636912 ns/op        49297570 B/op     731228 allocs/op
Benchmark_Grule_Load_Rules/100_rules#06-12                    12          98414282 ns/op        49297168 B/op     731122 allocs/op
Benchmark_Grule_Load_Rules/100_rules#07-12                    12          97733003 ns/op        49299440 B/op     731184 allocs/op
Benchmark_Grule_Load_Rules/100_rules#08-12                    12          98122635 ns/op        49297690 B/op     731132 allocs/op
Benchmark_Grule_Load_Rules/100_rules#09-12                    12          98451525 ns/op        49292262 B/op     731055 allocs/op

Benchmark_Grule_Load_Rules/1000_rules-12                       2         933617752 ns/op        488126636 B/op   7239752 allocs/op
Benchmark_Grule_Load_Rules/1000_rules#01-12                    2         926896605 ns/op        488120920 B/op   7239869 allocs/op
Benchmark_Grule_Load_Rules/1000_rules#02-12                    2         928509980 ns/op        488118076 B/op   7239757 allocs/op
Benchmark_Grule_Load_Rules/1000_rules#03-12                    2         926093793 ns/op        488119492 B/op   7239927 allocs/op
Benchmark_Grule_Load_Rules/1000_rules#04-12                    2         924214904 ns/op        488154840 B/op   7240215 allocs/op
Benchmark_Grule_Load_Rules/1000_rules#05-12                    2         928009912 ns/op        488078180 B/op   7239902 allocs/op
Benchmark_Grule_Load_Rules/1000_rules#06-12                    2         925822584 ns/op        488082700 B/op   7239303 allocs/op
Benchmark_Grule_Load_Rules/1000_rules#07-12                    2         923116273 ns/op        488088032 B/op   7239301 allocs/op
Benchmark_Grule_Load_Rules/1000_rules#08-12                    2         924545950 ns/op        488103888 B/op   7240207 allocs/op
Benchmark_Grule_Load_Rules/1000_rules#09-12                    2         930476936 ns/op        488166652 B/op   7240389 allocs/op


```

To load `100` rules into knowledgeBase it took `99342047 ns/op` (took the highest value) that is equal to `~99.342047ms` and (`49295906 B/op`) `~49.295906MB` memory per operation

To load `1000` rules into knowledgeBase it took `933617752 ns/op` (took the highest value) that is equal to `~933.617752ms` and (`488126636 B/op`) `~488.126636` memory per operation

### Test2 - Executing a fact against rules
Result:
```go
Benchmark_Grule_Execution_Engine/100_rules-12             140134              8175 ns/op            3939 B/op         59 allocs/op
Benchmark_Grule_Execution_Engine/100_rules#01-12          140442              8240 ns/op            3939 B/op         59 allocs/op
Benchmark_Grule_Execution_Engine/100_rules#02-12          141249              8151 ns/op            3937 B/op         59 allocs/op
Benchmark_Grule_Execution_Engine/100_rules#03-12          142011              8191 ns/op            3935 B/op         59 allocs/op
Benchmark_Grule_Execution_Engine/100_rules#04-12          137010              8226 ns/op            3947 B/op         59 allocs/op
Benchmark_Grule_Execution_Engine/100_rules#05-12          122870              9112 ns/op            3989 B/op         59 allocs/op
Benchmark_Grule_Execution_Engine/100_rules#06-12          133470              9697 ns/op            3957 B/op         59 allocs/op
Benchmark_Grule_Execution_Engine/100_rules#07-12          135206              8210 ns/op            3952 B/op         59 allocs/op
Benchmark_Grule_Execution_Engine/100_rules#08-12          139328              8213 ns/op            3941 B/op         59 allocs/op
Benchmark_Grule_Execution_Engine/100_rules#09-12          136437              8287 ns/op            3949 B/op         59 allocs/op

Benchmark_Grule_Execution_Engine/1000_rules-12              1912            525881 ns/op          273244 B/op       3843 allocs/op
Benchmark_Grule_Execution_Engine/1000_rules#01-12           2014            508415 ns/op          260310 B/op       3651 allocs/op
Benchmark_Grule_Execution_Engine/1000_rules#02-12           1770            568959 ns/op          293710 B/op       4147 allocs/op
Benchmark_Grule_Execution_Engine/1000_rules#03-12           1984            513188 ns/op          263958 B/op       3706 allocs/op
Benchmark_Grule_Execution_Engine/1000_rules#04-12           1771            566971 ns/op          293550 B/op       4145 allocs/op
Benchmark_Grule_Execution_Engine/1000_rules#05-12           1858            541169 ns/op          280695 B/op       3954 allocs/op
Benchmark_Grule_Execution_Engine/1000_rules#06-12           1896            530956 ns/op          275395 B/op       3875 allocs/op
Benchmark_Grule_Execution_Engine/1000_rules#07-12           1939            522682 ns/op          269694 B/op       3790 allocs/op
Benchmark_Grule_Execution_Engine/1000_rules#08-12           1851            545408 ns/op          281652 B/op       3968 allocs/op
Benchmark_Grule_Execution_Engine/1000_rules#09-12           1844            543697 ns/op          282657 B/op       3983 allocs/op



```

To execute a fact against 100 rules, Grule Engine took `~9697 ns/op` (took the highest value as base) that is hardly `~0.009697ms` and `3957 B/op` which is pretty fast.

To execute a fact against 1000 rules, Grule Engine took `~568959 ns/op` (took the highest value as base) that is hardly `~0.568959ms` and `293710 B/op` which is also pretty fast.


