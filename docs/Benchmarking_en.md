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
Benchmark_Grule_Load_Rules/100_rules-12                       44          23445211 ns/op         8871465 B/op     216553 allocs/op
Benchmark_Grule_Load_Rules/100_rules#01-12                    51          24634072 ns/op         8871551 B/op     216554 allocs/op
Benchmark_Grule_Load_Rules/100_rules#02-12                    46          25222948 ns/op         8871361 B/op     216553 allocs/op
Benchmark_Grule_Load_Rules/100_rules#03-12                    44          28146387 ns/op         8871486 B/op     216553 allocs/op
Benchmark_Grule_Load_Rules/100_rules#04-12                    45          30634115 ns/op         8871631 B/op     216554 allocs/op
Benchmark_Grule_Load_Rules/100_rules#05-12                    49          22554323 ns/op         8871228 B/op     216552 allocs/op
Benchmark_Grule_Load_Rules/100_rules#06-12                    46          22489783 ns/op         8871493 B/op     216554 allocs/op
Benchmark_Grule_Load_Rules/100_rules#07-12                    54          25210478 ns/op         8871489 B/op     216554 allocs/op
Benchmark_Grule_Load_Rules/100_rules#08-12                    54          23163462 ns/op         8871450 B/op     216553 allocs/op
Benchmark_Grule_Load_Rules/100_rules#09-12                    51          37621286 ns/op         8871515 B/op     216554 allocs/op
Benchmark_Grule_Load_Rules/100_rules#10-12                    33          37549279 ns/op         8871417 B/op     216553 allocs/op



Benchmark_Grule_Load_Rules/1000_rules-12                       4         345123608 ns/op        88642916 B/op    2141300 allocs/op
Benchmark_Grule_Load_Rules/1000_rules#01-12                    3         358250329 ns/op        88638602 B/op    2141279 allocs/op
Benchmark_Grule_Load_Rules/1000_rules#02-12                    3         347296778 ns/op        88646288 B/op    2141318 allocs/op
Benchmark_Grule_Load_Rules/1000_rules#03-12                    5         283385379 ns/op        88643715 B/op    2141304 allocs/op
Benchmark_Grule_Load_Rules/1000_rules#04-12                    3         372053420 ns/op        88643888 B/op    2141303 allocs/op
Benchmark_Grule_Load_Rules/1000_rules#05-12                    5         213746759 ns/op        88643011 B/op    2141299 allocs/op
Benchmark_Grule_Load_Rules/1000_rules#06-12                    6         230968014 ns/op        88641652 B/op    2141293 allocs/op
Benchmark_Grule_Load_Rules/1000_rules#07-12                    5         216604105 ns/op        88645020 B/op    2141310 allocs/op
Benchmark_Grule_Load_Rules/1000_rules#08-12                    5         213267279 ns/op        88640585 B/op    2141289 allocs/op
Benchmark_Grule_Load_Rules/1000_rules#09-12                    5         214347871 ns/op        88641289 B/op    2141292 allocs/op
Benchmark_Grule_Load_Rules/1000_rules#10-12                    5         211954473 ns/op        88642294 B/op    2141297 allocs/op
```

To load `100` rules into knowledgeBase it took `37549279 ns/op` (took the highest value) that is equal to `37.5ms` and (`8871417 B/op`) `8.8MB` memory per operation

To load `1000` rules into knowledgeBase it took `211954473 ns/op` (took the highest value) that is equal to `~211ms` and `88MB` memory per operation

### Test2 - Executing a fact against rules
Command to run: 
```go
go test -bench=. -benchmem
```

Result:
```go
Benchmark_Grule_Execution_Engine/100_rules-12              35340             33921 ns/op            4391 B/op         78 allocs/op
Benchmark_Grule_Execution_Engine/100_rules#01-12           29650             34346 ns/op            4446 B/op         79 allocs/op
Benchmark_Grule_Execution_Engine/100_rules#02-12           29587             34380 ns/op            4429 B/op         79 allocs/op
Benchmark_Grule_Execution_Engine/100_rules#03-12           31029             34342 ns/op            4423 B/op         78 allocs/op
Benchmark_Grule_Execution_Engine/100_rules#04-12           29646             35943 ns/op            4451 B/op         79 allocs/op
Benchmark_Grule_Execution_Engine/100_rules#05-12           35835             33039 ns/op            4402 B/op         78 allocs/op
Benchmark_Grule_Execution_Engine/100_rules#06-12           29305             34495 ns/op            4390 B/op         79 allocs/op
Benchmark_Grule_Execution_Engine/100_rules#07-12           28704             34857 ns/op            4397 B/op         79 allocs/op
Benchmark_Grule_Execution_Engine/100_rules#08-12           34936             34349 ns/op            4448 B/op         78 allocs/op
Benchmark_Grule_Execution_Engine/100_rules#09-12           36352             33935 ns/op            4332 B/op         77 allocs/op
Benchmark_Grule_Execution_Engine/100_rules#10-12           30698             39917 ns/op            4377 B/op         79 allocs/op
``

Benchmark_Grule_Execution_Engine/1000_rules-12              3478            317176 ns/op           43912 B/op        690 allocs/op
Benchmark_Grule_Execution_Engine/1000_rules#01-12           3434            319312 ns/op           44239 B/op        698 allocs/op
Benchmark_Grule_Execution_Engine/1000_rules#02-12           3565            321366 ns/op           43288 B/op        675 allocs/op
Benchmark_Grule_Execution_Engine/1000_rules#03-12           3385            330214 ns/op           44611 B/op        707 allocs/op
Benchmark_Grule_Execution_Engine/1000_rules#04-12           2544            454601 ns/op           53267 B/op        916 allocs/op
Benchmark_Grule_Execution_Engine/1000_rules#05-12            744           1420603 ns/op          137563 B/op       2953 allocs/op
Benchmark_Grule_Execution_Engine/1000_rules#06-12           1472            824068 ns/op           78644 B/op       1529 allocs/op
Benchmark_Grule_Execution_Engine/1000_rules#07-12           2671            408376 ns/op           51611 B/op        876 allocs/op
Benchmark_Grule_Execution_Engine/1000_rules#08-12           3524            359907 ns/op           43579 B/op        682 allocs/op
Benchmark_Grule_Execution_Engine/1000_rules#09-12           2970            357360 ns/op           48268 B/op        795 allocs/op
Benchmark_Grule_Execution_Engine/1000_rules#10-12           3511            367940 ns/op           43670 B/op        684 allocs/op

```

To execute a fact against 100 rules, Grule Engine took `~39917 ns/op` (took the highest value as base) that is hardly `~0.03917 ms` and `4377 B/op` which is pretty fast.

To execute a fact against 1000 rules, Grule Engine took `~1420603 ns/op` (took the highest value as base) that is hardly `~1.420603 ms` and `137563B/op` which is also pretty fast.


