# Grule Rule Language (GRL) Literals

[![GRL_Literals_cn](https://github.com/yammadev/flag-icons/blob/master/png/CN.png?raw=true)](../cn/GRL_Literals_cn.md)
[![GRL_Literals_de](https://github.com/yammadev/flag-icons/blob/master/png/DE.png?raw=true)](../de/GRL_Literals_de.md)
[![GRL_Literals_en](https://github.com/yammadev/flag-icons/blob/master/png/GB.png?raw=true)](../en/GRL_Literals_en.md)
[![GRL_Literals_id](https://github.com/yammadev/flag-icons/blob/master/png/ID.png?raw=true)](../id/GRL_Literals_id.md)
[![GRL_Literals_pl](https://github.com/yammadev/flag-icons/blob/master/png/PL.png?raw=true)](../pl/GRL_Literals_pl.md)

[About](About_cn.md) | [Tutorial](Tutorial_cn.md) | [Rule Engine](RuleEngine_cn.md) | [GRL](GRL_cn.md) | [GRL JSON](GRL_JSON_cn.md) | [RETE Algorithm](RETE_cn.md) | [Functions](Function_cn.md) | [FAQ](GRL_Literals_cn.md) | [Benchmark](Benchmarking_cn.md)

---

## 字符串字面变量

在GRL中，字符串是一个被单引号 `'` 或者双引号`"`包围的连续的字符.

如果是以单引号开始的，必须以单引号结束。对双引号也是如此。

举例

```go
"a quick brown fox jumps over a lazy dog"
```

或者

```go
'a quick brown fox jumps over a lazy dog'
```

字符串字面变量可以包含空格字符，比如 `space`, `tab` 或者`carriage-return`

举例

```go
"A quick brown fox
    Jumps
Over a lazy dog"
```

字符串中为了包含特殊字符，你需要跟Go一样进行转义。

举例

```go
"This string contains \" Double Quote"
```

## 数字字面变量

GRL中数字字面变量 跟Golang指定的尽可能相同。它可以理解各种各样的数字格式，比如10进制，8进制，和16进制。二进制目前还没实现。

### 整型字面变量

#### 十进制

十进制，举例

```go
0
123
34592
-1
-47234
```

8进制，举例

```go
01
07
010
017
-034
-045
04328 (error : invalid octal number)
```

16进制，举例

```go
0x1
0xF
0x10
0x1F
0xFF00
-0x12
-0x00ABCD
-0x890AbCdEf
```

### 实数或者浮点数字面变量

十进制，举例

```go
0.
72.40
072.40
2.71828
1.e+0
6.67428e-11
1E6
.25
.12345E+5
-072.40
-2.71828
-1.e+0
```

16进制举例

```go
0x1p-2 
0x2.p10
0x1.Fp+0
0X.8p-0
0X_1FFFP-16
0x15e-2
```

## 布尔字面变量

```go
true
TRUE
True
TrUe
false
False
FALSE
FaLsE
```
