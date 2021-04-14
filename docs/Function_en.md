# Function in Grule

[Tutorial](Tutorial_en.md) | [Rule Engine](RuleEngine_en.md) | [GRL](GRL_en.md) | [GRL JSON](GRL_JSON_en.md) | [RETE Algorithm](RETE_en.md) | [Functions](Function_en.md) | [FAQ](FAQ_en.md) | [Benchmark](Benchmarking_en.md)

---

## Built-In Functions

Built-in functions are all defined within the `ast/BuiltInFunctions.go` file. As of now, they are:

### MakeTime(year, month, day, hour, minute, second int64) time.Time

`MakeTime` will create a `time.Time` with local `locale`.

#### Arguments

* `year` is the Year number.
* `month` is the Month number, January = 1.
* `day` is the day number in a month.
* `hour` the hour of the day starting from 0.
* `minute` the minute of the hour starting from 0.
* `second` the second of the minute starting from 0.

#### Returns

* `time.Time` value representing the time as specified in the argument in `local` locale.

#### Example

```Shell
rule SetExpire "Set the expire date for Fact created before 2020" {
    when
       Fact.CreateTime < MakeTime(2020,1,1,0,0,0)
    then
       Fact.ExpireTime = MakeTime(2021,1,1,0,0,0);
}
```

### Changed(variableName string)

`Changed` will ensure the specified `variableName` is removed from the working
memory before the next cycle.

#### Arguments

* `variableName` the variable name to be removed from working memory.

#### Example

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

`Now` function will create a new `time.Time` value containing the current time.

#### Returns

* `time.Time` value representing the current value

#### Example

```Shell
rule ResetTime "Reset the lastUpdate time" {
    when
        Fact.LastUpdate < Now()
    then
        Fact.LastUpdate = Now();
}
```

### Log(text string)

`Log` will emit a log-debug string from within the rule.

#### Arguments

* `text` The text to emit into the Log-Debug

#### Example

```Shell
rule SomeRule "Log candidate name if he is below 17 years old" {
    when
        Candidate.Age < 17
    then
        Log("Under aged: " + Candidate.Name);
}
```

### IsNil(i interface{}) bool

`IsNil` will check if the argument is a `nil` value.

#### Arguments

* `i` a variable to check.

#### Returns

* `true` if the specified argument is`nil` or an invalid `ptr` value.
* `false` if the specified argument is a valid `ptr` value.

#### Example

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

`IsZero` will check any variable in the argument for its `Zero` status value. Zero means
that the variable is newly defined and has not been assigned an initial value.
This is usually applied to types like `string`, `int64`, `uint64`, `bool`,
`time.Time`, etc.

#### Arguments

* `i` a variable to check.

#### Returns

* `true` if the specified argument is Zero.
* `false` if the specified argument not Zero.

#### Example

```Shell
rule CheckStartTime "Check device's starting time." {
    when
        IsZero(Device.StartTime) == true
    then
        Device.StartTime = Now();
}
```

### Retract(ruleName string)

`Retract` will remove the specified rule from the next cycle evaluation. If a
rule is retracted its `when` scope will not be evaluated on the next immediate
cycle after the call to `Retract`. Before the following cycle, all retracted
rules will be restored.

#### Arguments

* `ruleName` name of the rule to retract.

#### Example

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

`GetTimeYear` will extract the Year value of the time argument.

#### Arguments

* `time` The time variable

#### Returns

* Year value of the time.

#### Example

```Shell
rule StartNewYearProcess "Check if it's a new year to restart new FinancialYear." salience 1000 {
    when
        GetTimeYear(Now()) != GL.FinancialYear
    then
        GL.CloseYear(GL.FinancialYear)
}
```

### GetTimeMonth(time time.Time) int

`GetTimeMonth` will extract the Month value of the time argument.

#### Arguments

* `time` The time variable

#### Returns

* Month value of the time. 1 = January.

#### Example

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

`GetTimeDay` will extract the Day of the month value of the time argument.

#### Arguments

* `time` The time variable

#### Returns

* Day of month value of the time.

#### Example

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

`GetTimeHour` will extract the Hour value of the time argument.

#### Arguments

* `time` The time variable

#### Returns

* Hour value of the time. Is between 0 to 23

#### Example

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

`GetTimeMinute` will extract the Minute value of the time argument.

#### Arguments

* `time` The time variable

#### Returns

* Minute value of the time, between 0 to 59

#### Example

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

`GetTimeSecond` will extract the Second value of the time argument.

#### Arguments

* `time` The time variable

#### Returns

* Second value of the time, between 0 to 59

#### Example

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

`IsTimeBefore` will check if a time value precedes another time value.

#### Arguments

* `time` The time value you wish to have checked
* `before` The time value against which the above is checked

#### Returns

* True if the `before` time value precedes the `time` value.
* False if the `before` time value does not precede the `time` value.

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

`IsTimeAfter` will check if a time value follows another time value.

#### Arguments

* `time` The time value you wish to have checked
* `after` The time value against which the above is checked

#### Returns

* True if the `after` time value follows `time` value.
* False if the `after` time value does not follow the `time` value.

#### Example

```Shell
rule AdditionalTax "Apply additional tax if new tax rules are in effect." {
    when
        IsTimeAfter(Purchase.TransactionTime, TaxRegulation.StartSince)
    then
        Purchase.Tax = Purchase.Tax + 0.01;
}
```

### TimeFormat(time time.Time, layout string) string

`TimeFormat` will format a time argument as specified by `layout` argument.

#### Arguments

* `time` The time value you wish to have formatted.
* `layout` String variable specifying the date format layout.

For the layout format, you can [read this article](https://yourbasic.org/golang/format-parse-string-time-date-example/)

#### Returns

* A string formatted as specified.

#### Example

```Shell
rule LogPurchaseDate "Log the purchase date." {
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

#### Example

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

#### Returns

* The length of string's receiver

#### Example

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

#### Arguments

* `string` The string to compare to

#### Returns

* `< 0` if receiver is less than the argument
* `0` if receiver is equal to the argument
* `> 0` if receiver is greater thant the argument

#### Example

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

#### Arguments

* `string` The substring to check within the receiver

#### Returns

* `true` if the argument string is contained within the receiver.
* `false` if the argument string is not contained within the receiver.

#### Example

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

#### Arguments

* `string` The substring to count within the receiver

#### Returns

* number of occurences of the argument in the receiver.

#### Example

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

#### Arguments

* `string` The expected prefix.

#### Returns

* `true` if the receiver has the argument as its prefix.
* `false` if the receiver does not have the argument as its prefix.

#### Example

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

#### Arguments

* `string` The expected suffix.

#### Returns

* `true` if the receiver has the argument as its suffix.
* `false` if the receiver does not have the argument as its suffix.

#### Example

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

#### Arguments

* `string` The substring to search for.

#### Returns

* The index value of the first occurrence of the argument.

#### Example

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

#### Arguments

* `string` The substring to search for.

#### Returns

* The index of the last occurrence of the argument.

#### Example

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

#### Arguments

* `int64` the repeat count

#### Returns

* A new string containing `n` occurrences of the receiver.

#### Example

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

#### Arguments

* `old` the substring you wish to have replaced.
* `new` the string you wish to replace all occurrences of `old`.

#### Returns

* A string where all instances of `old` in the receiver have been replaced with `new`.

#### Example

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

#### Arguments

* `string` the token you wish to use to split the receiver.

#### Returns

* The string slice containing parts of the original string as split by the token.

#### Example

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

#### Returns

* A new string that is a lower-cased version of the receiver.

#### Example

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

#### Returns

* A new string that is an upper-cased version of the receiver.

#### Example

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

#### Returns

* A string with the whitespace removed from the beginning and end.

#### Example

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

#### Returns

* The length of array/slice.

#### Example

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

#### Arguments

* `val` value to have appended.

#### Example

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

#### Returns

* The length of map receiver.

#### Example

```Shell
rule DoSomething "Do something when map length is sufficient" {
   when
       Fact.ChildrenMap.Len() > 2
   then
       Fact.DoSomething();
}
```

## Custom Functions

All functions that are acessible from the DataContext are **Invocable** from
within the rule, both in the "When" scope and the "Then" scope.

You can create functions and have your Fact as receiver, and those functions
can be called from GRL.

For example. Given:

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

You can make calls to the defined methods:

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

### Variadic Function Arguments

Variadic arguments are supported for custom functions.

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

This function can then be called from within a rule with zero or more values supplied for the variadic argument.

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

### The Laws of Custom Function in Grule

When you make your own function to be called from the rule engine, you need to know the following laws:

1. The function must be visible, meaning that functions must start with a
   capital letter. Private functions cannot be executed.
2. The function must only return one value type. Returning multiple values from
   a function is not supported and the rule execution will fail if there are
   multiple return values.
3. The way number literals are treated in Grule's GRL is such that a
   **integer** will always be taken as an `int64` type and a **real** as
   `float64`, thus you must always define your numeric types accordingly.
