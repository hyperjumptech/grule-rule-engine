# Fungsi-fungsi didalam Grule

---

:construction:
__THIS PAGE IS BEING TRANSLATED__
:construction:

:construction_worker: Contributors are invited. Please read [CONTRIBUTING](../../CONTRIBUTING.md) and [CONTRIBUTING TRANSLATION](../CONTRIBUTING_TRANSLATION.md) guidelines.

:vulcan_salute: Please remove this note once you're done translating.

---


[![Function_cn](https://github.com/yammadev/flag-icons/blob/master/png/CN.png?raw=true)](../cn/Function_cn.md)
[![Function_de](https://github.com/yammadev/flag-icons/blob/master/png/DE.png?raw=true)](../de/Function_de.md)
[![Function_en](https://github.com/yammadev/flag-icons/blob/master/png/GB.png?raw=true)](../en/Function_en.md)
[![Function_id](https://github.com/yammadev/flag-icons/blob/master/png/ID.png?raw=true)](../id/Function_id.md)
[![Function_in](https://github.com/yammadev/flag-icons/blob/master/png/IN.png?raw=true)](../in/Function_in.md)

[Tentang Grule](About_id.md) | [Tutorial](Tutorial_id.md) | [Rule Engine](RuleEngine_id.md) | [GRL](GRL_id.md) | [GRL JSON](GRL_JSON_id.md) | [Algoritma RETE](RETE_id.md) | [Fungsi-fungsi](Function_id.md) | [FAQ](FAQ_id.md) | [Benchmark](Benchmarking_id.md)

---

## Fungsi-fungsi Built-In

Fungsi-fungsi built-in semuanya dapat ditemukan di file `ast/BuildInFunctions.go`. Saat ini, ia berisi :

### MakeTime(year, month, day, hour, minute, second int64) time.Time

`MakeTime` akan membuat `time.Time` dengan `locale` yang lokal.

#### Penjelasan Argumen

* `year` Angka tahun.
* `month` Angka bulan, Januari = 1.
* `day` Angka tanggal dalam bulan.
* `hour` Angka jam dalam hari dimulai dari 0.
* `minute` Angka menit dalam jam dimulai dari 0.
* `second` Angka detik dalam menit dimulai dari 0.

#### Mengembalikan

* `time.Time` nilai yang merepresentasikan waktu dimana menggunakan zona waktu lokal.

#### Contoh

```text
rule SetExpire "Set the expire date for Fact created before 2020" {
    when
       Fact.CreateTime < MakeTime(2020,1,1,0,0,0)
    then
       Fact.ExpireTime = MakeTime(2021,1,1,0,0,0);
}
```

### Changed(variableName string)

`Changed` akan membuang nilai pada nama variabel yang ter-rekam dalam __Working Memory__ menjadi hilang, sehingga nilai tersebut akan diambil lagi dari konteks data.

#### Penjelasan Argumen

* `variableName` Nama variabel yang mana nilai yang tercatat dalam __working memory__ akan dibuang.

#### Contoh

```text
rule SetExpire "Set new expire date" {
    when
        IsZero(Fact.ExpireTime)
    then
        Fact.CalculateExpire(); // this function will internally change the ExpireTime variable
        Changed("Fact.ExpireTime")
}
```

### Now() time.Time

`Now` fungsi akan membuat nilai `time.Time` dengan nilai waktu saat fungsi ini dipanggil.

#### Mengembalikan

* `time.Time` nilai berisi waktu sekarang.

#### Contoh

```text
rule ResetTime "Reset the lastUpdate time" {
    when
        Fact.LastUpdate < Now()
    then
        Fact.LastUpdate = Now();
}
```

### Log(text string)

`Log` Akan mengeluarkan sebuah log dengan tingkat "Debug" dari dalam GRL

#### Penjelasan Argumen

* `text` Teks yang akan dikeluarkan dalam log

#### Contoh

```text
rule SomeRule "Log candidate name if he is bellow 17 years old" {
    when
        Candidate.Age < 17
    then
        Log("Under aged : " + Candidate.Name);
}
```

### IsNil(i interface{}) bool

`IsNil` melakukan pemeriksaan apakah nilai dalam argumen adalah `nil`.

#### Penjelasan Argumen

* `i` variabel yang akan diperiksa.

#### Mengembalikan

* `true` Jika argumen berisi `nil` atau memiliki nilai `ptr` yang tidak valid.
* `false` Jika argumen bukan `nil` atau memiliki nilai `ptr` yang valid.

#### Contoh

```text
rule CheckEducation "Check candidate's education fact" {
    when
        IsNil(Candidate.Education) == false &&
        Candidate.Education.Grade == "PHD"
    then
        Candidate.Onboard = true;
}
```

### IsZero(i interface{}) bool

`IsZero` akan memeriksa nilai argumen apakah berstatus `Zero`. Zero berarti
argumen tersebut baru saja dibuat dan belum diberi nilai apapun oleh program (menggunakan nilai default)
Ini biasanya berlaku pada beberapa tipe variabel seperti `string`, `int64`, `uint64`, `bool`, `time.Time`, etc.

#### Penjelasan Argumen

* `i` argumen yang akan diperiksa.

#### Mengembalikan

* `true` Jika argumen berstatus Zero.
* `false` Jika argumen tidak berstatus Zero.

#### Contoh

```text
rule CheckStartTime "Check device's starting time." {
    when
        IsZero(Device.StartTime) == true
    then
        Device.StartTime = Now();
}
```

### Retract(ruleName string)

`Retract` akan menarik __rule__ yang disebutkan dari evaluasi di siklus berikutnya.
Jika sebuah __rule__ ditarik, maka skop `when` nya tidak akan dievaluasi. Saat __rule engine__ ini dijalankan kembail
dari awal, maka semua __rule__ yang ditarik akan dikembalikan seperti sediakala.

#### Penjelasan Argumen

* `ruleName` Nama __rule__ yang akan ditarik.

#### Contoh

```text
rule CheckStartTime "Check device's starting time." salience 1000 {
    when
        IsZero(Device.StartTime) == true
    then
        Device.StartTime = Now();
        Retract("CheckStartTime");
}
```

### GetTimeYear(time time.Time) int

`GetTimeYear` Akan mengambil nilai tahun dari argumen.

#### Penjelasan Argumen

* `time` Variabel `time.Time`

#### Mengembalikan

* Nilai tahun

#### Contoh

```text
rule StartNewYearProcess "Check if its a new year to restart new FinancialYear." salience 1000 {
    when
        GetTimeYear(Now()) != GL.FinancialYear
    then
        GL.CloseYear(GL.FinancialYear)
}
```

### GetTimeMonth(time time.Time) int

`GetTimeMonth` Akan mengambil nilai bulan dari argumen.

#### Penjelasan Argumen

* `time` The time variable

#### Mengembalikan

* Month value of the time. 1 = January.

#### Contoh

```text
rule StartNewYearProcess "Check if its a new year to restart new FinancialYear." salience 1000 {
    when
        isZero(Process.Month)
    then
        Process.Month = GetTimeMonth(Process.Month);
}
```

### GetTimeDay(time time.Time) int

`GetTimeDay` Akan mengambil nilai tanggal dari argumen.

#### Penjelasan Argumen

* `time` The time variable

#### Mengembalikan

* Day of month value of the time.

#### Contoh

```text
rule GreetEveryDay "Log a greeting every day." salience 1000 {
    when
        Greeting.Day != GetTimeDay(Now())
    then
        Log("Its a new Day !!!")
        Retract("GreetEveryDay")
}
```

### GetTimeHour(time time.Time) int

`GetTimeHour` Akan mengambil nilai jam dari argumen.

#### Penjelasan Argumen

* `time` The time variable

#### Mengembalikan

* Hour value of the time. Is between 0 to 23

#### Contoh

```text
rule DailyCheckBuild "Execute build every 6AM and 6PM." {
    when
        GetTimeHour(Now()) == 6 || GetTimeHour(Now()) == 18
    then
        CiCd.BuildDaily();
        Retract("DailyCheckBuild");
}
```

### GetTimeMinute(time time.Time) int

`GetTimeMinute` Akan mengambil nilai menit dari argumen.

#### Penjelasan Argumen

* `time` The time variable

#### Mengembalikan

* Minute value of the time, between 0 to 59

#### Contoh

```text
rule DailyCheckBuild "Execute build every 6.30AM and 6.30PM." {
    when
        (GetTimeHour(Now()) == 6 || GetTimeHour(Now()) == 18) &&
        GetTimeMinute(Now()) == 30
    then
        CiCd.BuildDaily();
        Retract("DailyCheckBuild");
}
```

### GetTimeSecond(time time.Time) int

`GetTimeSecond` Akan mengambil nilai detik dari argumen.

#### Penjelasan Argumen

* `time` The time variable

#### Mengembalikan

* Second value of the time, between 0 to 59

#### Contoh

```text
rule DailyCheckBuild "Execute build every 6.30AM and 6.30PM." {
    when
        (GetTimeHour(Now()) == 6 || GetTimeHour(Now()) == 18) &&
        GetTimeMinute(Now()) == 30 && GetTimeSecond(Now()) == 0
    then
        CiCd.BuildDaily();
        Retract("DailyCheckBuild");
}
```

### IsTimeBefore(time, before time.Time) bool

`IsTimeBefore` akan memeriksa apakah sebuah waktu pada argumen 1 itu adalah waktu sebelum nilai waktu pada argumen 2.

#### Penjelasan Argumen

* `time` The time variable
* `before` Another time variable

#### Mengembalikan

* True if the `before` time value is before the `time` value.
* False if the `before` time value is not before the `time` value.

#### Contoh

```text
rule PromotionExpireCheck  "Apply a promotion if the promotion's expired date is not due." {
    when
        IsTimeBefore(Now(), Promotion.ExpireDateTime)
    then
        Promotion.Discount = 0.10;
        Retract("PromotionExpireCheck");
}
```

### IsTimeAfter(time, after time.Time) bool

`IsTimeAfter` akan memeriksa apakah sebuah waktu pada argumen 1 itu adalah waktu sesudah nilai waktu pada argumen 2.

#### Penjelasan Argumen

* `time` The time variable
* `after` Another time variable

#### Mengembalikan

* True if the `after` time value is after the `time` value.
* False if the `after` time value is not after the `time` value.

#### Contoh

```text
rule AdditionalTax  "Apply additional tax if purchase after date specified." {
    when
        IsTimeAfter(Purchase.TransactionTime, TaxRegulation.StartSince)
    then
        Purchase.Tax = PurchaseTax + 0.01;
}
```

### TimeFormat(time time.Time, layout string) string

`TimeFormat` akan mem-format nilai waktu pada argumen sesuai dengan format yang ditentukan pada argumen `layout`.

#### Penjelasan Argumen

* `time` The time variable
* `layout` String variable specifying the date format layout.

For layout format, you can [read this article](https://yourbasic.org/golang/format-parse-string-time-date-example/)

#### Mengembalikan

* String contains the formatted time

#### Contoh

```text
rule LogPurchaseDate  "Log the purchase date." {
    when
        IsZero(Purchase.TransactionDate) == false
    then
        Log(TimeFormat(Purchase.TransactionDate, "2006-01-02T15:04:05-0700");
}
```

### Complete()

`Complete` will cause the engine to stop processing further rules in its
current cycle. This is useful if you want to terminate further rule evaluation
under a set condition.

#### Contoh

```Shell
rule DailyCheckBuild "Execute build at 6.30AM and 6.30PM." {
    when
        (GetTimeHour(Now()) == 6 || GetTimeHour(Now()) == 18) &&
        GetTimeMinute(Now()) == 30 && GetTimeSecond(Now()) == 0
    then
        CiCd.BuildDaily();
        Complete();
}
```

## Math Functions

All the functions bellow is a wrapper to their golang math functions.
You should read Golang math page to know how to use each function.

Unlike go, you don't have to use the `math.` prefix
to use them in your GRL.

Use them like normal built in function.

```go
when 
    Max(Fact.A, Fact.C, Fact.B) > 10
then
    Fact.X = Acosh(Fact.C);
```

- Max(vals ...float64) float64
- Min(vals ...float64) float64
- Abs(x float64) float64
- Acos(x float64) float64
- Acosh(x float64) float64
- Asin(x float64) float64
- Asinh(x float64) float64
- Atan(x float64) float64
- Atan2(y, x float64) float64
- Atanh(x float64) float64
- Cbrt(x float64) float64
- Ceil(x float64) float64
- Copysign(x, y float64) float64
- Cos(x float64) float64
- Cosh(x float64) float64
- Dim(x, y float64) float64
- Erf(x float64) float64
- Erfc(x float64) float64
- Erfcinv(x float64) float64
- Erfinv(x float64) float64
- Exp(x float64) float64
- Exp2(x float64) float64
- Expm1(x float64) float64
- Float64bits(f float64) uint64
- Float64frombits(b uint64) float64
- Floor(x float64) float64
- Gamma(x float64) float64
- Hypot(p, q float64) float64
- Ilogb(x float64) int
- IsInf(f float64, sign int64) bool
- IsNaN(f float64) (is bool)
- J0(x float64) float64
- J1(x float64) float64
- Jn(n int64, x float64) float64
- Ldexp(frac float64, exp int64) float64
- MathLog(x float64) float64
- Log10(x float64) float64
- Log1p(x float64) float64
- Log2(x float64) float64
- Logb(x float64) float64
- Mod(x, y float64) float64
- NaN() float64
- Pow(x, y float64) float64
- Pow10(n int64) float64
- Remainder(x, y float64) float64
- Round(x float64) float64
- RoundToEven(x float64) float64
- Signbit(x float64) bool
- Sin(x float64) float64
- Sinh(x float64) float64
- Sqrt(x float64) float64
- Tan(x float64) float64
- Tanh(x float64) float64
- Trunc(x float64) float64


## Constant Functions

The following functions can be called from within GRL as long as the receiver
value type is correct.

### string.Len() int

`Len` will return string's length.

#### Mengembalikan

* The length of string's receiver

#### Contoh

```Shell
rule DoSomething "Do something when string length is sufficient" {
    when
        Fact.Name.Len() > "ATextConstant".Len()
    then
        Fact.DoSomething();
}
```

### string.Compare(string) int

`Compare` will compare the receiver string to the argument.

#### Penjelasan Argumen

* `string` The string to compare to

#### Mengembalikan

* `< 0` if receiver is less than the argument
* `0` if receiver is equal to the argument
* `> 0` if receiver is greater thant the argument

#### Contoh

```Shell
rule CompareString "Do something when Fact.Text is greater than A" {
    when
        Fact.Text.Compare("A") > 0
    then
        Fact.DoSomething();
}
```

### string.Contains(string) bool

`Contains` will check if its argument is contained within the receiver.

#### Penjelasan Argumen

* `string` The substring to check within the receiver

#### Mengembalikan

* `true` if the argument string is contained within the receiver.
* `false` if the argument string is not contained within the receiver.

#### Contoh

```Shell
rule ContainString "Do something when Fact.Text is contains XXX" {
    when
        Fact.Text.Contains("XXX")
    then
        Fact.DoSomething();
}
```

### string.Count(string) int

`Count` will count the number of occurences of argument in receiver string.

#### Penjelasan Argumen

* `string` The substring to count within the receiver

#### Mengembalikan

* number of occurences of the argument in the receiver.

#### Contoh

```Shell
rule CountString "Do something when Fact.Text contains 3 occurrences of 'ABC'" {
    when
        Fact.Text.Count("ABC") == 3
    then
        Fact.DoSomething();
}
```

### string.HasPrefix(string) bool

`HasPrefix` will check if the receiver string has a specific prefix.

#### Penjelasan Argumen

* `string` The expected prefix.

#### Mengembalikan

* `true` if the receiver has the argument as its prefix.
* `false` if the receiver does not have the argument as its prefix.

#### Contoh

```Shell
rule IsPrefixed "Do something when Fact.Text started with PREF" {
    when
        Fact.Text.HasPrefix("PREF")
    then
        Fact.DoSomething();
}
```

### string.HasSuffix(string) bool

`HasSuffix` will check if the receiver string has a specific suffix.

#### Penjelasan Argumen

* `string` The expected suffix.

#### Mengembalikan

* `true` if the receiver has the argument as its suffix.
* `false` if the receiver does not have the argument as its suffix.

#### Contoh

```Shell
rule IsSuffixed "Do something when Fact.Text ends with SUFF" {
    when
        Fact.Text.HasSuffix("SUFF")
    then
        Fact.DoSomething();
}
```

### string.Index(string) int

`Index` will return the index of the first occurrence of the argument in the receiver string.

#### Penjelasan Argumen

* `string` The substring to search for.

#### Mengembalikan

* The index value of the first occurrence of the argument.

#### Contoh

```Shell
rule IndexCheck "Do something when Fact.Text ABC occurs as specified" {
    when
        Fact.Text.Index("ABC") == "abABCabABC".Index("ABC")
    then
        Fact.DoSomething();
}
```

### string.LastIndex(string) int

`LastIndex` will return the index of last occurrence of the argument in the receiver string.

#### Penjelasan Argumen

* `string` The substring to search for.

#### Mengembalikan

* The index of the last occurrence of the argument.

#### Contoh

```Shell
rule LastIndexCheck "Do something when Fact.Text ABC occurs in the last position as specified" {
    when
        Fact.Text.LastIndex("ABC") == "abABCabABC".LastIndex("ABC")
    then
        Fact.DoSomething();
}
```

### string.Repeat(int64) string

`Repeat` will return a string containing `n` occurrences of the receiver string.

#### Penjelasan Argumen

* `int64` the repeat count

#### Mengembalikan

* A new string containing `n` occurrences of the receiver.

#### Contoh

```Shell
rule StringRepeat "Do something when Fact.Text contains ABCABCABC" {
    when
        Fact.Text == "ABC".Repeat(3)
    then
        Fact.DoSomething();
}
```

### string.Replace(old, new string) string

`Replace` will return a string with all occurrences of `old` replaced with `new`.

#### Penjelasan Argumen

* `old` the substring you wish to have replaced.
* `new` the string you wish to replace all occurrences of `old`.

#### Mengembalikan

* A string where all instances of `old` in the receiver have been replaced with `new`.

#### Contoh

```Shell
rule ReplaceString "Do something when Fact.Text contains replaced string" {
    when
        Fact.Text == "ABC123ABC".Replace("123","ABC")
    then
        Fact.DoSomething();
}
```

### string.Split(string) []string

`Split` will return a string slice whose elements are determined after
splitting the receiver by the string token argument.  The token will not be
present in the resulting slice elements.

#### Penjelasan Argumen

* `string` the token you wish to use to split the receiver.

#### Mengembalikan

* The string slice containing parts of the original string as split by the token.

#### Contoh

```Shell
rule SplitString "Do something when Fact.Text is prefixed by 'ABC,'" {
    when
        Fact.Text.Split(",")[0] == "ABC"
    then
        Fact.DoSomething();
}
```

### string.ToLower() string

`ToLower` will return a string whose contents are all lower case instances of
characters in the receiver.

#### Mengembalikan

* A new string that is a lower-cased version of the receiver.

#### Contoh

```Shell
rule LowerText "Do something when Fact.Text is equal to 'abc'" {
    when
        Fact.Text.ToLower() == "Abc".ToLower()
    then
        Fact.DoSomething();
}
```

### string.ToUpper() string

`ToUpper` will return a string whose contents are all upper case instances of
characters in the receiver.

#### Mengembalikan

* A new string that is an upper-cased version of the receiver.

#### Contoh

```Shell
rule UpperText "Do something when Fact.Text is equal to 'ABC'" {
    when
        Fact.Text.ToUpper() == "Abc".ToUpper()
    then
        Fact.DoSomething();
}
```

### string.Trim() string

`Trim` will return a string where the whitespace on either end of the string has been removed.

#### Mengembalikan

* A string with the whitespace removed from the beginning and end.

#### Contoh

```Shell
rule TrimText "Do something when Fact.Text is 'ABC'" {
    when
        Fact.Text == "  Abc   ".Trim().ToUpper()
    then
        Fact.DoSomething();
}
```

### array.Len() int

`Len` will return the length of the array/slice.

#### Mengembalikan

* The length of array/slice.

#### Contoh

```Shell
rule DoSomething "Do something when array length is sufficient" {
    when
        Fact.ChildrenArray.Len() > 2
    then
        Fact.DoSomething();
}
```

### array.Append(val)

`Append` will append `val` onto the end of the receiver array.

#### Penjelasan Argumen

* `val` value to have appended.

#### Contoh

```Shell
rule DoSomething "Add a new child when the array has less than 2 children" {
    when
        Fact.ChildrenArray.Len() < 2
    then
        Fact.ChildrenArray.Append(Fact.NewChild());
}
```

### map.Len() int

`Len` will return map's length.

#### Mengembalikan

* The length of map receiver.

#### Contoh

```Shell
rule DoSomething "Do something when map length is sufficient" {
   when
       Fact.ChildrenMap.Len() > 2
   then
       Fact.DoSomething();
}
```

## Menambah Fungsi Sendiri (Custom)

Semua fungsi yang mana bisa dijalankan dari `DataContext` bisa dijalankan dari dalam skrip __rule__,
baik di dalam skop "When" atau "Then".

Karenanya, anda dapat membuat sebuah fungsi dan menjadikan fungsi ini merupakan milik (method) dari data, maka fungsi/method itu
bisa dijalankan dari dalam GRL.

Diasumsikan anda memiliki sebuah `struct` dengan beberapa fungsi.

```go
type MyPoGo struct {
}

func (p *MyPoGo) GetStringLength(sarg string) int {
    return len(sarg)
}

func (p *MyPoGo) AppendString(aString, subString string) string {
    return sprintf("%s%s", aString, subString)
}
```

Dan `struct` tersebut ditambahkan kedalam konteks data

```go
dctx := grule.context.NewDataContext()
dctx.Add("Pogo", &MyPoGo{})
```

Anda dapat menjalan fungsi-fungsi tadi dalam __rule__

```go
when
    Pogo.GetStringLength(some.variable) < 100
then
    some.variable = Pogo.AppendString(some.variable, "Groooling");
```

### Argumen-argumen fungsi Variadic

Argumen-argumen Variadic dapat dipergunakan di dalam fungsi.

```go
func (p *MyPoGo) GetLongestString(strs... string) string {
    var longestStr string
    for _, s := range strs {
        if len(s) > len(longestStr) {
            longestStr = s
        }
    }
    return longestStr
}
```

This function can then be called from within a rule with zero or more values supplied for the variadic argument

```go
when
    Pogo.GetStringLength(some.variable) < 100
then
    some.longest = Pogo.GetLongestString(some.stringA, some.stringB, some.stringC);
```

Since it is possible to provide zero values to satisfy a variadic argument, they can also be used to simulate optional parameters.

```go
func (p *MyPoGo) AddTax(cost int64, optionalTaxRate... float64) int64 {
    var taxRate float64 = 0.2
    if len(optionalTaxRate) > 0 {
        taxRate = optionalTaxRate[0]
    }
    return cost * (1+taxRate)
}
```

```go
when
    Pogo.IsTaxApplied() == false
then
    some.cost = Pogo.AddTax(come.cost);

//or

when
    Pogo.IsTaxApplied() == false
then
    some.cost = Pogo.AddTax(come.cost, 0.15);
```

### Important Thing you must know about Custom Function in Grule

When you make your own function to be called from the rule engine, you need to know the following rules.

1. The function must be visible. The convention for public functions should start with a capital letter. Private functions cannot be executed.
2. The function must only return 1 type. Returning multiple variables from a function is not supported, the rule execution will fail if there are multiple return variables.
3. The way number literals are treated in Grule's DRL is such that a **decimal** will always be taken as an `int64` type and a **real** as `float64`, thus always consider to define your function arguments and returns from the types `int64` and `float64` when you work with numbers.