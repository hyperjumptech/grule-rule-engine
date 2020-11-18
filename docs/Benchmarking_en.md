Benchmarking:
--
For any library `Benchmarking` is very much required in terms of evaluating the performance and helps us to improve the performance better.

I have benchmarked two things:
* Loading rules into KnowledgeBase
* Executing a fact against rules

All the tests will run for multiple times to see how it is performing by changing the value of N 

`N is b.N where b is an instance of *testing.B`

### Rules:
refer `100_rules.grl` and `1000_rules.grl` files under `examples/benchmark` directory


Result 
-- 

### Test1 - Loading Rules into KnowledgeBase
Command to run: 
```go
go test -bench=. -benchmem
```

Result:
```go
goos: darwin
goarch: amd64
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

To load `1000` rules into knowledgeBase it took `261857403 ns/op` (took the highest value) that is equal to `~211ms` and `88MB` memory per operation

### Test2 - Executing a fact against rules
Command to run: 
```go
go test -bench=. -benchmem
```

Result:
```go
Benchmark_Grule_Execution_Engine/100_rules-12              36260             33049 ns/op            4374 B/op         77 allocs/op
Benchmark_Grule_Execution_Engine/100_rules#01-12           31179             33153 ns/op            4428 B/op         78 allocs/op
Benchmark_Grule_Execution_Engine/100_rules#02-12           30547             34040 ns/op            4418 B/op         79 allocs/op
Benchmark_Grule_Execution_Engine/100_rules#03-12           31520             33168 ns/op            4417 B/op         78 allocs/op
Benchmark_Grule_Execution_Engine/100_rules#04-12           30784             34392 ns/op            4438 B/op         79 allocs/op
Benchmark_Grule_Execution_Engine/100_rules#05-12           31651             33713 ns/op            4368 B/op         78 allocs/op
Benchmark_Grule_Execution_Engine/100_rules#06-12           31008             33189 ns/op            4374 B/op         78 allocs/op
Benchmark_Grule_Execution_Engine/100_rules#07-12           31280             35847 ns/op            4466 B/op         78 allocs/op
Benchmark_Grule_Execution_Engine/100_rules#08-12           32546             34744 ns/op            4360 B/op         78 allocs/op
Benchmark_Grule_Execution_Engine/100_rules#09-12           32839             35002 ns/op            4471 B/op         78 allocs/op
Benchmark_Grule_Execution_Engine/100_rules#10-12           31684             32847 ns/op            4368 B/op         78 allocs/op

Benchmark_Grule_Execution_Engine/1000_rules-12              3590            339138 ns/op           43116 B/op        671 allocs/op
Benchmark_Grule_Execution_Engine/1000_rules#01-12           3368            318206 ns/op           44745 B/op        710 allocs/op
Benchmark_Grule_Execution_Engine/1000_rules#02-12           3685            314999 ns/op           42480 B/op        656 allocs/op
Benchmark_Grule_Execution_Engine/1000_rules#03-12           3620            332785 ns/op           42911 B/op        666 allocs/op
Benchmark_Grule_Execution_Engine/1000_rules#04-12           3426            475685 ns/op           44299 B/op        700 allocs/op
Benchmark_Grule_Execution_Engine/1000_rules#05-12           3464            383093 ns/op           44014 B/op        693 allocs/op
Benchmark_Grule_Execution_Engine/1000_rules#06-12           3337            326025 ns/op           44987 B/op        716 allocs/op
Benchmark_Grule_Execution_Engine/1000_rules#07-12           3637            324763 ns/op           42797 B/op        663 allocs/op
Benchmark_Grule_Execution_Engine/1000_rules#08-12           2763            366498 ns/op           50507 B/op        849 allocs/op
Benchmark_Grule_Execution_Engine/1000_rules#09-12           3164            421806 ns/op           46439 B/op        751 allocs/op
Benchmark_Grule_Execution_Engine/1000_rules#10-12           3648            323597 ns/op           42722 B/op        661 allocs/op


```

To execute a fact against 100 rules, Grule Engine took `~35847 ns/op` (took the highest value as base) that is hardly `~0.035847ms` and `4466B/op` which is pretty fast.

To execute a fact against 1000 rules, Grule Engine took `~475685 ns/op` (took the highest value as base) that is hardly `~0.475685ms` and `44299B/op` which is also pretty fast.


