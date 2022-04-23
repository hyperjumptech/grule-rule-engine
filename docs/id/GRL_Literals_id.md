# Literal-literal Dalam Grule Rule Language (GRL)

---

:construction:
__THIS PAGE IS BEING TRANSLATED__
:construction:

:construction_worker: Contributors are invited. Please read [CONTRIBUTING](../../CONTRIBUTING.md) and [CONTRIBUTING TRANSLATION](../CONTRIBUTING_TRANSLATION.md) guidelines.

:vulcan_salute: Please remove this note once you're done translating.

---


[![GRL_Literals_cn](https://github.com/yammadev/flag-icons/blob/master/png/CN.png?raw=true)](../cn/GRL_Literals_cn.md)
[![GRL_Literals_de](https://github.com/yammadev/flag-icons/blob/master/png/DE.png?raw=true)](../de/GRL_Literals_de.md)
[![GRL_Literals_en](https://github.com/yammadev/flag-icons/blob/master/png/GB.png?raw=true)](../en/GRL_Literals_en.md)
[![GRL_Literals_id](https://github.com/yammadev/flag-icons/blob/master/png/ID.png?raw=true)](../id/GRL_Literals_id.md)
[![GRL_Literals_pl](https://github.com/yammadev/flag-icons/blob/master/png/PL.png?raw=true)](../pl/GRL_Literals_pl.md)

[Tentang Grule](About_id.md) | [Tutorial](Tutorial_id.md) | [Rule Engine](RuleEngine_id.md) | [GRL](GRL_id.md) | [GRL JSON](GRL_JSON_id.md) | [Algoritma RETE](RETE_id.md) | [Fungsi-fungsi](Function_id.md) | [FAQ](FAQ_id.md) | [Benchmark](Benchmarking_id.md)

---

## String Literals

In GRL, a string is any sequence of characters surrounded by either a single `'` or double `"` quotes.
If the literal is started with single quote, then it must be terminated by single quote. The same for double quote.

For example

```go
"a quick brown fox jumps over a lazy dog"
```

or

```go
'a quick brown fox jumps over a lazy dog'
```

A string literal may contain white space characters such as `space`, `tab` or a
`carriage-return`

For example

```go
"A quick brown fox
    Jumps
Over a lazy dog"
```

To include special characters in string, you can *escape* them as is normal in Go

For example

```go
"This string contains \" Double Quote"
```

## Number Literals

GRL follows literal numbering as specified by the Golang specification as best
as it can. It understands various numbers notation such as
Base10 (Decimals), Base8 (Octal) and Base16 (Hex). Base2 (Binary) is not yet implemented.

### Integer Literal

#### Decimals

In Base 10 - For Example

```go
0
123
34592
-1
-47234
```

In Base 8 - For Example

```go
01
07
010
017
-034
-045
04328 (error : invalid octal number)
```

In Base 16 - For Example

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

### Real Numbers / Float Literals

In Base 10 - For Example

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

In Base 16 - For Example

```go
0x1p-2 
0x2.p10
0x1.Fp+0
0X.8p-0
0X_1FFFP-16
0x15e-2
```

## Boolean Literal

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
