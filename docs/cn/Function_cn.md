# Grule 中的函数

[![Function_cn](https://github.com/yammadev/flag-icons/blob/master/png/CN.png?raw=true)](../cn/Function_cn.md)
[![Function_de](https://github.com/yammadev/flag-icons/blob/master/png/DE.png?raw=true)](../de/Function_de.md)
[![Function_en](https://github.com/yammadev/flag-icons/blob/master/png/GB.png?raw=true)](../en/Function_en.md)
[![Function_id](https://github.com/yammadev/flag-icons/blob/master/png/ID.png?raw=true)](../id/Function_id.md)
[![Function_pl](https://github.com/yammadev/flag-icons/blob/master/png/PL.png?raw=true)](../pl/Function_pl.md)

[About](About_cn.md) | [Tutorial](Tutorial_cn.md) | [Rule Engine](RuleEngine_cn.md) | [GRL](GRL_cn.md) | [GRL JSON](GRL_JSON_cn.md) | [RETE Algorithm](RETE_cn.md) | [Functions](Function_cn.md) | [FAQ](Function_cn.md) | [Benchmark](Benchmarking_cn.md)

---

## 内置函数

内置函数在 `ast/BuiltInFunctions.go` 文件中。

### MakeTime(year, month, day, hour, minute, second int64) time.Time

`MakeTime` 将会使用local `locale`去创建 `time.Time`

#### 参数

* `year` 年
* `month` 月, January = 1.
* `day`日.
* `hour` 从0开始的小时.
* `minute` 从0开始的分钟.
* `second` 从 0开始的秒.

#### 发挥

* `time.Time` 返回使用local `locale`的时间。

#### 示例

```Shell
rule SetExpire "Set the expire date for Fact created before 2020" {
    when
       Fact.CreateTime < MakeTime(2020,1,1,0,0,0)
    then
       Fact.ExpireTime = MakeTime(2021,1,1,0,0,0);
}
```

### Changed(variableName string)

`Changed` 将会保证指定的 `variableName` 在下一个循环中从工作内存中移除。

#### 参数

* `variableName` 将从工作变量中移除的变量名称.

#### 示例

```Shell
rule SetExpire "Set new expire date" {
    when
        IsZero(Fact.ExpireTime)
    then
        Fact.CalculateExpire(); // this function will internally change the ExpireTime variable
        Changed("Fact.ExpireTime")
}
```

### Now() time.Time

`Now` 将会返回当前时间 `time.Time` .

#### 返回

* `time.Time` 代表着当前时间

#### 示例

```Shell
rule ResetTime "Reset the lastUpdate time" {
    when
        Fact.LastUpdate < Now()
    then
        Fact.LastUpdate = Now();
}
```

### Log(text string)

`Log` 将会打印规则的日志.

#### 参数

* `text` 要打印的日志文本

#### 示例

```Shell
rule SomeRule "Log candidate name if he is below 17 years old" {
    when
        Candidate.Age < 17
    then
        Log("Under aged: " + Candidate.Name);
}
```

### IsNil(i interface{}) bool

`IsNil` 将会检查参数是否是 `nil` .

#### 参数

* `i` 将要检查的变量.

#### 返回

* `true` 指定的参数是`nil` 或者一个无效的 `ptr` 值.
* `false`指定的参数是一个有效的 `ptr` 值.

#### 示例

```Shell
rule CheckEducation "Check candidate's education fact" {
    when
        IsNil(Candidate.Education) == false &&
        Candidate.Education.Grade == "PHD"
    then
        Candidate.Onboard = true;
}
```

### IsZero(i interface{}) bool

`IsZero`将会检查参数的变量是否是  `Zero` . Zero 意味这个这个变量是新定义的，还没有赋予初始值。经常被用来检查`string`, `int64`, `uint64`, `bool`,`time.Time`, 等等.

#### 参数

* `i` 将要检查的变量

#### 返回

* `true` 如果指定的参数是 Zero.
* `false` 如果指定的参数不是 Zero.

#### 示例

```Shell
rule CheckStartTime "Check device's starting time." {
    when
        IsZero(Device.StartTime) == true
    then
        Device.StartTime = Now();
}
```

### Retract(ruleName string)

`Retract` 将会在下一个循环评估中排查指定的规则.如果一个规则，在调用`Retract`之后在下一个循环中，在`when`中被排除。当引擎从头开始再次启动时，它将会自动将所有规则重置放回原处。

#### 参数

* `ruleName` 规则名称.

#### 示例

```Shell
rule CheckStartTime "Check device's starting time." salience 1000 {
    when
        IsZero(Device.StartTime) == true
    then
        Device.StartTime = Now();
        Retract("CheckStartTime");
}
```

### GetTimeYear(time time.Time) int

`GetTimeYear` 将会从指定的日期中提取年份

#### 参数

* `time` 时间变量

#### 返回

* 对应的年份

#### 示例

```Shell
rule StartNewYearProcess "Check if it's a new year to restart new FinancialYear." salience 1000 {
    when
        GetTimeYear(Now()) != GL.FinancialYear
    then
        GL.CloseYear(GL.FinancialYear)
}
```

### GetTimeMonth(time time.Time) int

`GetTimeMonth` 获取时间的月份.

#### 参数

* `time` 时间变量

#### 返回

* 月份. 1 = January.

#### 示例

```Shell
// TODO: something's not right here. The description is copy/pasted from above
// but the condition/action doesn't make sense to me
rule StartNewYearProcess "Check if its a new year to restart new FinancialYear." salience 1000 {
    when
        isZero(Process.Month)
    then
        Process.Month = GetTimeMonth(Process.Month);
}
```

### GetTimeDay(time time.Time) int

`GetTimeDay` 获取时间的天.

#### 参数

* `time` 时间

#### 返回

* 时间月的天.

#### 示例

```Shell
rule GreetEveryDay "Log a greeting every day." salience 1000 {
    when
        Greeting.Day != GetTimeDay(Now())
    then
        Log("Its a new Day !!!")
        Retract("GreetEveryDay")
}
```

### GetTimeHour(time time.Time) int

`GetTimeHour` 获取时间的小时数.

#### 参数

* `time` 时间变量

#### 返回

* 小时. 0 to 23

#### 示例

```Shell
rule DailyCheckBuild "Execute build every 6AM and 6PM." {
    when
        GetTimeHour(Now()) == 6 || GetTimeHour(Now()) == 18
    then
        CiCd.BuildDaily();
        Retract("DailyCheckBuild");
}
```

### GetTimeMinute(time time.Time) int

`GetTimeMinute` 获取分钟数.

#### 参数

* `time` 时间

#### 返回

* 时间的分钟,  0 to 59

#### 示例

```Shell
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

`GetTimeSecond` 时间的秒.

#### 参数

* `time` 时间

#### 返回

* 时间的秒,  0 to 59

#### 示例

```Shell
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

`IsTimeBefore` 将会检查一个时间是否在另一个时间的前面

#### 参数

* `time` 你希望检查的时间
* `before` 与上面参数对比的时间

#### 返回

* True 如果 `before` 领先 `time` 变量.
* False如果 `before` 不领先 `time` 变量.

#### Example

```Shell
rule PromotionExpireCheck "Apply a promotion if promotion hasn't yet expired." {
    when
        IsTimeBefore(Now(), Promotion.ExpireDateTime)
    then
        Promotion.Discount = 0.10;
        Retract("PromotionExpireCheck");
}
```

### IsTimeAfter(time, after time.Time) bool

`IsTimeAfter` 检查一个时间是否在另一个时间之后

#### 参数

* `time` 希望检查的时间
* `after` 以上述参数对应的时间

#### 返回

* True 如果 `after` 在 `time` 之后.
* False如果`after` 不在 `time` 之后.

#### 示例

```Shell
rule AdditionalTax "Apply additional tax if new tax rules are in effect." {
    when
        IsTimeAfter(Purchase.TransactionTime, TaxRegulation.StartSince)
    then
        Purchase.Tax = Purchase.Tax + 0.01;
}
```

### TimeFormat(time time.Time, layout string) string

`TimeFormat` 将会返回`layout` 指定的时间格式

#### 参数

* `time` 将被格式化的时间
* `layout` 时间格式

有关时间格式，参考 [read this article](https://yourbasic.org/golang/format-parse-string-time-date-example/)

#### 返回

* 指定格式的时间字符串

#### 示例

```Shell
rule LogPurchaseDate "Log the purchase date." {
    when
        IsZero(Purchase.TransactionDate) == false
    then
        Log(TimeFormat(Purchase.TransactionDate, "2006-01-02T15:04:05-0700");
}
```

### Complete()

`Complete` 将会在当前循环中使引擎停止执行后面的规则。如果你想停止执行后面的规则评估，这个很有用

#### 举例

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

## Math Functions数学函数

所有的函数都是golang数学函数的封装。你可以阅读Golang 相关页面去了解每一个函数。
不像go，你不需要在你的GRL中使用 `math.` 前缀。

.就像正常的内置函数一样使用他们

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


## Constant Functions常量函数

只要保证接收者的数据类型是正确的，接下来的函数都可以在GRL中调用。

### string.Len() int

`Len` 返回字符串的长度.

#### 犯规

* 字符串接收者的长度

#### 示例

```Shell
rule DoSomething "Do something when string length is sufficient" {
    when
        Fact.Name.Len() > "ATextConstant".Len()
    then
        Fact.DoSomething();
}
```

### string.Compare(string) int

`Compare` 对比接收者和参数。

#### 参数

* `string` 将要对比的字符串

#### 返回

* `< 0` 如果接收者比参数小 
* `0` 如果接收者和参数相同
* `> 0` 如果接收者比参数大

#### 示例

```Shell
rule CompareString "Do something when Fact.Text is greater than A" {
    when
        Fact.Text.Compare("A") > 0
    then
        Fact.DoSomething();
}
```

### string.Contains(string) bool

`Contains` 检查接收者是否包含参数指定的字符串

#### 参数

* `string` 将被检查的子串

#### 返回

* `true` 如果接收者包含参数字符串
* `false` 如果接收者不包含参数字符串

#### 示例

```Shell
rule ContainString "Do something when Fact.Text is contains XXX" {
    when
        Fact.Text.Contains("XXX")
    then
        Fact.DoSomething();
}
```

### string.In(string ...) bool

`In` 检查任一参数等于接收者

#### 参数

* `string` 变长字符参数

#### 返回

* boolean `true` 如果任一参数等于接收者, 反正是 `false`.

#### 示例

```Shell
rule CheckArgumentIn "Do something when Fact.Text is equals to 'ABC' or 'BCD' or 'CDE' " {
    when
        Fact.Text.In("ABC", "BCD", "CDE")
    then
        Fact.DoSomething();
}
```

### string.Count(string) int

`Count` 统计接收者中参数字符串出现的次数

#### 参数

* `string`将要检查的子串

#### 返回

* 子串在接收者出现的次数

#### 示例

```Shell
rule CountString "Do something when Fact.Text contains 3 occurrences of 'ABC'" {
    when
        Fact.Text.Count("ABC") == 3
    then
        Fact.DoSomething();
}
```

### string.HasPrefix(string) bool

`HasPrefix` 是否有前缀

#### 参数

* `string` 期望的前缀

#### 返回

* `true` 如果接收者有前缀
* `false` 接收者没有前缀

#### 示例

```Shell
rule IsPrefixed "Do something when Fact.Text started with PREF" {
    when
        Fact.Text.HasPrefix("PREF")
    then
        Fact.DoSomething();
}
```

### string.HasSuffix(string) bool

`HasSuffix` 后缀.

#### 参数

* `string` 要检查的后缀

#### 返回

* `true`接收者有后缀.
* `false` 接收者没有后缀.

#### 示例

```Shell
rule IsSuffixed "Do something when Fact.Text ends with SUFF" {
    when
        Fact.Text.HasSuffix("SUFF")
    then
        Fact.DoSomething();
}
```

### string.Index(string) int

`Index` 将会返回参数第一次出现的索引

#### 参数

* `string` 将要搜索的子串

#### 返回

* 参数第一次出现的索引

#### 示例

```Shell
rule IndexCheck "Do something when Fact.Text ABC occurs as specified" {
    when
        Fact.Text.Index("ABC") == "abABCabABC".Index("ABC")
    then
        Fact.DoSomething();
}
```

### string.LastIndex(string) int

`LastIndex` 最后一次出现的子串索引

#### 参数

* `string` 要检索的子串

#### 返回

* 最后一次出现的子串索引

#### 示例

```Shell
rule LastIndexCheck "Do something when Fact.Text ABC occurs in the last position as specified" {
    when
        Fact.Text.LastIndex("ABC") == "abABCabABC".LastIndex("ABC")
    then
        Fact.DoSomething();
}
```

### string.Repeat(int64) string

`Repeat` 返回一个接收者重复`n`次的字符串

#### 参数

* `int64` 重复次数

#### 返回

* 重复 `n` 次的字符串

#### 示例

```Shell
rule StringRepeat "Do something when Fact.Text contains ABCABCABC" {
    when
        Fact.Text == "ABC".Repeat(3)
    then
        Fact.DoSomething();
}
```

### string.Replace(old, new string) string

`Replace` 返回一个用`new`替换`old`的字符串

#### 参数

* `old` 希望被替换的子串
* `new` 要替换 `old`的新子串.

#### 返回

* 一个被替换过的字符串

#### 示例

```Shell
rule ReplaceString "Do something when Fact.Text contains replaced string" {
    when
        Fact.Text == "ABC123ABC".Replace("123","ABC")
    then
        Fact.DoSomething();
}
```

### string.Split(string) []string

`Split` 将会把字符串切分成slice，分割符由参数指定。分隔符将不会出现在结果slice元素里面。

#### 参数

* `string`分隔符.

#### 返回

* 被分割完成之后的字符串切片

#### 示例

```Shell
rule SplitString "Do something when Fact.Text is prefixed by 'ABC,'" {
    when
        Fact.Text.Split(",")[0] == "ABC"
    then
        Fact.DoSomething();
}
```

### string.ToLower() string

`ToLower` 字符串转小写

#### 返回

* 一个新的接收者转完小写的字符串

#### 示例

```Shell
rule LowerText "Do something when Fact.Text is equal to 'abc'" {
    when
        Fact.Text.ToLower() == "Abc".ToLower()
    then
        Fact.DoSomething();
}
```

### string.ToUpper() string

`ToUpper` 字符串转大写

#### 返回

* 一个接收者大写的字符串

#### 示例

```Shell
rule UpperText "Do something when Fact.Text is equal to 'ABC'" {
    when
        Fact.Text.ToUpper() == "Abc".ToUpper()
    then
        Fact.DoSomething();
}
```

### string.Trim() string

`Trim` 去除字符串前后空格

#### 返回

* 去除过前后空格的字符串

#### 示例

```Shell
rule TrimText "Do something when Fact.Text is 'ABC'" {
    when
        Fact.Text == "  Abc   ".Trim().ToUpper()
    then
        Fact.DoSomething();
}
```


### string.MatchString() string

`MatchString` MatchString reports whether the string s contains any match of the regular expression pattern. Similar to golang [MatchString](https://pkg.go.dev/regexp#MatchString)

#### Returns

* True if the `regexPattern` matches the string s
* False if the `regexPattern` doesn't match the string s

#### Example

```Shell
rule MatchStringText "Return true when regex pattern matches the string"  {
	when
	  Fact.Text.MatchString("B([a-z]+)ck")
	then
	  Fact.DoSomething();
}
```

### array.Len() int

`Len` 返回数据或者切片的长度.

#### 返回

* 数组或者切片的长度

#### 示例

```Shell
rule DoSomething "Do something when array length is sufficient" {
    when
        Fact.ChildrenArray.Len() > 2
    then
        Fact.DoSomething();
}
```

### array.Append(val) 

`Append` 添加 `val` 到接收者切片

#### 参数

* `val` 将要被添加的元素

#### 示例

```Shell
rule DoSomething "Add a new child when the array has less than 2 children" {
    when
        Fact.ChildrenArray.Len() < 2
    then
        Fact.ChildrenArray.Append(Fact.NewChild());
}
```

### map.Len() int

`Len` 返回map的长度

#### 返回

* 接收者map的长度

#### 示例

```Shell
rule DoSomething "Do something when map length is sufficient" {
   when
       Fact.ChildrenMap.Len() > 2
   then
       Fact.DoSomething();
}
```

## 自定义函数

可从 DataContext 访问的所有函数都可以从规则内调用，无论是在“When”范围还是“Then”范围。

你创建了一些函数，然后Fact是接收者，这些函数就可以在GRL中被调用。

举例:

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

你可以调用定义好的方法

```go
dctx := grule.context.NewDataContext()
dctx.Add("Pogo", &MyPoGo{})

rule "If it's possible to Groool, Groool" {
    when
        Pogo.GetStringLength(some.variable) < 100
    then
        some.variable = Pogo.AppendString(some.variable, "Groooling");
}
```

### 可变函数参数

自定义函数是支持可变长度的参数的

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

这个函数可以被规则传递0个或者多个值。

```go
when
    Pogo.GetStringLength(some.variable) < 100
then
    some.longest = Pogo.GetLongestString(some.stringA, some.stringB, some.stringC);
```

因为在可变参数中可以提供0个值，你可以模拟optional参数。

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

### Grule中的自定义函数法则

当你在规则引擎中调用自己的函数，你需要知道以下的法则：

1. 函数必须是可见的，意味着函数名称的第一个字母是大写字母。私有函数不能被执行。
2. 函数只能返回一个值。多个函数返回值不支持，如果函数有多个返回值将会导致规则执行失败。
3. Grule中GRL中的字面变量如果是整型，则应该被看做是`int64`，如果是有理数，则是`float64`，因此你对应的修改参数数字的定义。
