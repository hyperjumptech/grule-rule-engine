# Literały Grule Rule Language (GRL)

[![GRL_Literals_cn](https://github.com/yammadev/flag-icons/blob/master/png/CN.png?raw=true)](../cn/GRL_Literals_cn.md)
[![GRL_Literals_de](https://github.com/yammadev/flag-icons/blob/master/png/DE.png?raw=true)](../de/GRL_Literals_de.md)
[![GRL_Literals_en](https://github.com/yammadev/flag-icons/blob/master/png/GB.png?raw=true)](../en/GRL_Literals_en.md)
[![GRL_Literals_id](https://github.com/yammadev/flag-icons/blob/master/png/ID.png?raw=true)](../id/GRL_Literals_id.md)
[![GRL_Literals_pl](https://github.com/yammadev/flag-icons/blob/master/png/PL.png?raw=true)](../pl/GRL_Literals_pl.md)

[About](About_pl.md) | [Tutorial](Tutorial_pl.md) | [Rule Engine](RuleEngine_pl.md) | [GRL](GRL_pl.md) | [GRL JSON](GRL_JSON_pl.md) | [RETE Algorithm](RETE_pl.md) | [Functions](Function_pl.md) | [FAQ](FAQ_pl.md) | [Benchmark](Benchmarking_pl.md)

---

## Literały łańcuchowe

W GRL ciąg znaków to dowolna sekwencja znaków otoczona pojedynczym `'` lub podwójnym `'` cudzysłowem.
Jeśli literał zaczyna się od pojedynczego cudzysłowu, to musi być zakończony pojedynczym cudzysłowem. To samo dotyczy cudzysłowów podwójnych.

Na przykład

```go
"a quick brown fox jumps over a lazy dog"
```

lub

```go
'a quick brown fox jumps over a lazy dog'
```

Literał łańcuchowy może zawierać białe znaki spacji, takie jak `space`, `tab` lub `carriage-return`

Na przykład

```go
"A quick brown fox
    Jumps
Over a lazy dog"
```

Aby zawrzeć znaki specjalne w łańcuchu, można je *escape*, tak jak to jest normalnie w Go

Na przykład

```go
"This string contains \" Double Quote"
```

## Literały liczbowe

GRL stosuje numerację literalną określoną przez specyfikację Golanga najlepiej jak potrafi. Rozumie różne notacje liczbowe, takie jak Base10 (dziesiętne), Base8 (ósemkowe) i Base16 (szesnastkowe). Base2 (Binary) nie jest jeszcze zaimplementowana.

### Literały Liczb Całkowitych

#### Liczby dziesiętne

W Base 10 - na przykład

```go
0
123
34592
-1
-47234
```

W Base 8 - na przykład

```go
01
07
010
017
-034
-045
04328 (error : invalid octal number)
```

W Base 16 - na przykład

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

### Liczby rzeczywiste / Literały zmiennoprzecinkowe

W Base 10 - na przykład

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

W Base 16 - na przykład

```go
0x1p-2 
0x2.p10
0x1.Fp+0
0X.8p-0
0X_1FFFP-16
0x15e-2
```

## Literały Boolean

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
