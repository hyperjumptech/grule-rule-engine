# 基准测试

[![Benchmarking_cn](https://github.com/yammadev/flag-icons/blob/master/png/CN.png?raw=true)](../cn/Benchmarking_cn.md)
[![Benchmarking_de](https://github.com/yammadev/flag-icons/blob/master/png/DE.png?raw=true)](../de/Benchmarking_de.md)
[![Benchmarking_en](https://github.com/yammadev/flag-icons/blob/master/png/GB.png?raw=true)](../en/Benchmarking_en.md)
[![Benchmarking_id](https://github.com/yammadev/flag-icons/blob/master/png/ID.png?raw=true)](../id/Benchmarking_id.md)
[![Benchmarking_pl](https://github.com/yammadev/flag-icons/blob/master/png/PL.png?raw=true)](../pl/Benchmarking_pl.md)

[About](About_cn.md) | [Tutorial](Tutorial_cn.md) | [Rule Engine](RuleEngine_cn.md) | [GRL](GRL_cn.md) | [GRL JSON](GRL_JSON_cn.md) | [RETE Algorithm](RETE_cn.md) | [Functions](Function_cn.md) | [FAQ](FAQ_cn.md) | [Benchmark](Benchmarking_cn.md)

---

对于任意的库，`Benchmarking`在评估性能和提升性能的时候都是必须的。

我基准测试了两个事情：

*  加载100和1000个规则到KnowledgeBase
* 针对100个规则和1000个规则，使用一个fact去执行规则。

通过改变N值，所有的测试都将跑很多次去测试性能。

`N is b.N where b is an instance of *testing.B`

### 规则:
 `100_rules.grl` 和 `1000_rules.grl` 文件 在 `examples/benchmark` 目录

运行命令: 
---
```go
> go test -bench=. -benchmem
goos: darwin
goarch: amd64
Number of Cores - 6
Ram - 16 GB 2400 MHz DDR4
pkg: github.com/kalyan-arepalle/grule-rule-engine/examples/benchmark

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

加载`100`规则到knowledgeBase耗费`22628754 ns/op`（取最高值），等于`22.628754ms`，同时需要(`8871388 B/op`) `8.8MB` 内存。

加载`1000`规则到knowledgeBase耗费`261857403 ns/op`（取最高值），等于`~261ms`，同时需要 `88MB` 内存。

### 测试2 - 在规则上执行 fact
运行命令: 
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

加载`100`规则到knowledgeBase耗费`99342047 ns/op`（取最高值），等于`~99.342047ms`，同时需要(`49295906 B/op`) `~49.295906MB` 内存。

加载`1000`规则到knowledgeBase耗费`933617752 ns/op`（取最高值），等于`~933.617752ms`，同时需要`488126636 B/op` `~488.126636MB` 内存。

### 测试2 - 在规则上执行 fact
结果:
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

在100个规则上执行fact，Grule引擎耗费`~9697 ns/op`（取最大值）, 约等于`~0.009697ms`, 和`3957 B/op`内存，这是相当快的了。

在1000个规则上执行fact，Grule引擎耗费`~568959 ns/op`（取最大值）, 约等于`~0.568959ms`, 和`293710 B/op`内存，这是相当快的了。

