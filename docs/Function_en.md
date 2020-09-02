# Function in Grule

[Tutorial](Tutorial_en.md) | [Rule Engine](RuleEngine_en.md) | [GRL](GRL_en.md) | [RETE Algorithm](RETE_en.md) | [Functions](Function_en.md) | [FAQ](FAQ_en.md)

## Built-In Functions

Built-in functions are all defined within the `ast/BuiltInFunctions.go` file. As of now, they are :

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

`Changed` will make sure the specified variableName is removed from the working memory.

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

* `time.Time` value coontaining the current value

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
rule SomeRule "Log candidate name if he is bellow 17 years old" {
    when
        Candidate.Age < 17
    then
        Log("Under aged : " + Candidate.Name);
}
```

### IsNil(i interface{}) bool

`IsNil` will check any variable in the argument for `Nil` value.

#### Arguments

* `i` a variable to check.

#### Returns

* `true` if the specified argument contains `nil` or an invalid `ptr` value.
* `false` if the specified argument contains a valid `ptr` value.

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
that the variable is newly defined and have not been assigned with any value.
This is usually applied to types like `string`, `int64`, `uint64`, `bool`, `time.Time`, etc.

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

`Retract` will retract the specified rule from next cycle evaluation. If a rule
is retracted, its `when` scope will not be evaluated. When the rule engine execute
again, all the retracted status of all rules will be restored.

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
rule StartNewYearProcess "Check if its a new year to restart new FinancialYear." salience 1000 {
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

`IsTimeBefore` will check if a time value is before the other time value.

#### Arguments

* `time` The time variable
* `before` Another time variable

#### Returns

* True if the `before` time value is before the `time` value.
* False if the `before` time value is not before the `time` value.

#### Example

```Shell
rule PromotionExpireCheck  "Apply a promotion if the promotion's expired date is not due." {
    when
        IsTimeBefore(Now(), Promotion.ExpireDateTime)
    then
        Promotion.Discount = 0.10;
        Retract("PromotionExpireCheck");
}
```

### IsTimeAfter(time, after time.Time) bool

`IsTimeAfter` will check if a time value is after the other time value.

#### Arguments

* `time` The time variable
* `after` Another time variable

#### Returns

* True if the `after` time value is after the `time` value.
* False if the `after` time value is not after the `time` value.

#### Example

```Shell
rule AdditionalTax  "Apply additional tax if purchase after date specified." {
    when
        IsTimeAfter(Purchase.TransactionTime, TaxRegulation.StartSince)
    then
        Purchase.Tax = PurchaseTax + 0.01;
}
```

### TimeFormat(time time.Time, layout string) string

`TimeFormat` will format a time argument to a format specified by `layout` argument.

#### Arguments

* `time` The time variable
* `layout` String variable specifying the date format layout.

For layout format, you can [read this article](https://yourbasic.org/golang/format-parse-string-time-date-example/)

#### Returns

* String contains the formatted time

#### Example

```Shell
rule LogPurchaseDate  "Log the purchase date." {
    when
        IsZero(Purchase.TransactionDate) == false
    then
        Log(TimeFormat(Purchase.TransactionDate, "2006-01-02T15:04:05-0700");
}
```

### Complete()

`Complete` will cause the engine to stop processign further rules in its current cycle. This is useful if you want to terminate further rukle evaluation under a set condition.

#### Example

```Shell
rule DailyCheckBuild "Execute build every 6.30AM and 6.30PM." {
    when
        (GetTimeHour(Now()) == 6 || GetTimeHour(Now()) == 18) &&
        GetTimeMinute(Now()) == 30 && GetTimeSecond(Now()) == 0
    then
        CiCd.BuildDaily();
        Complete();
}
```

## Constant Functions

The following functions can be immediately called from within GRL
as long as the receiver value type is correct.

### string.Len() int

`Len` will return string's length.

#### Returns

* The length of string's receiver

#### Example

```Shell
rule DoSomething  "Do something when string length is sufficient" {
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
rule CompareString  "Do something when Fact.Text is greater than A" {
    when
        Fact.Text.Compare("A") > 0
    then
        Fact.DoSomething();
}
```

### string.Contains(string) bool

`Contains` will check if argument is contained within the receiver.

#### Arguments

* `string` The sub string to check

#### Returns

* `true` if receiver string is containing the argument
* `false` if receiver string do not contain the argument

#### Example

```Shell
rule ContainString  "Do something when Fact.Text is contains XXX" {
    when
        Fact.Text.Contains("XXX")
    then
        Fact.DoSomething();
}
```

### string.Count(string) int

`Count` will count the number of occurences of argument in receiver string.

#### Arguments

* `string` The sub string to check

#### Returns

* int count of occurences

#### Example

```Shell
rule CountString  "Do something when Fact.Text contains 3 ABC" {
    when
        Fact.Text.Count("ABC") == 3
    then
        Fact.DoSomething();
}
```

### string.HasPrefix(string) bool

`HasPrefix` will check if the receiver string is prefixed with the argument

#### Arguments

* `string` The sub string to check

#### Returns

* `true` if receiver have prefix of the argument
* `false` if receiver do not prefixed of the argument

#### Example

```Shell
rule IsPrefixed  "Do something when Fact.Text started with PREF" {
    when
        Fact.Text.HasPrefix("PREF")
    then
        Fact.DoSomething();
}
```

### string.HasSuffix(string) bool

`HasSuffix` will check if the receiver string is suffixed with the argument

#### Arguments

* `string` The sub string to check

#### Returns

* `true` if receiver have suffix of the argument
* `false` if receiver do not suffix of the argument

#### Example

```Shell
rule IsSuffixed  "Do something when Fact.Text ends with SUFF" {
    when
        Fact.Text.HasSuffix("SUFF")
    then
        Fact.DoSomething();
}
```

### string.Index(string) int

`Index` will return the offset of first occurrence of argument in receiver string

#### Arguments

* `string` The sub string to check

#### Returns

* int the offset of first occurrence of argument

#### Example

```Shell
rule IndexCheck  "Do something when Fact.Text ABC occurrence as specified" {
    when
        Fact.Text.Index("ABC") == "abABCabABC".Index("ABC")
    then
        Fact.DoSomething();
}
```

### string.LastIndex(string) int

`LastIndex` will return the offset of last occurrence of argument in receiver string

#### Arguments

* `string` The sub string to check

#### Returns

* int the offset of last occurrence of argument

#### Example

```Shell
rule LastIndexCheck  "Do something when Fact.Text ABC last occurrence as specified" {
    when
        Fact.Text.LastIndex("ABC") == "abABCabABC".LastIndex("ABC")
    then
        Fact.DoSomething();
}
```

### string.Repeat(int) string

`Repeat` will return a string containing repeated receiver string, n times

#### Arguments

* `int` the repeat count

#### Returns

* string contains repeated receiver string, n times

#### Example

```Shell
rule StringRepeat  "Do something when Fact.Text contains ABCABCABC" {
    when
        Fact.Text == "ABC".Repeat(3)
    then
        Fact.DoSomething();
}
```

### string.Replace(old, new string) string

`Replace` will return a string containing replaced receiver string

#### Arguments

* `old` the substring to search
* `new` the string replacement

#### Returns

* string contains receiver string after replace

#### Example

```Shell
rule ReplaceString  "Do something when Fact.Text contains replaced string" {
    when
        Fact.Text == "ABC123ABC".Replace("123","ABC")
    then
        Fact.DoSomething();
}
```

### string.Split(string) []string

`Split` will return a string slice contains splitted receiver string using argument as separator

#### Arguments

* `string` the separator

#### Returns

* string slice

#### Example

```Shell
rule SplitString  "Do something when Fact.Text split first index is ABC" {
    when
        Fact.Text.Split(",")[0] == "ABC"
    then
        Fact.DoSomething();
}
```

### string.ToLower() string

`ToLower` will return a string contains all lower case of the receiver

#### Returns

* string lower cased string

#### Example

```Shell
rule LowerText  "Do something when Fact.Text lower case is abc" {
    when
        Fact.Text.ToLower() == "Abc".ToLower()
    then
        Fact.DoSomething();
}
```

### string.ToUpper() string

`ToUpper` will return a string contains all upper case of the receiver

#### Returns

* string upper cased string

#### Example

```Shell
rule UpperText  "Do something when Fact.Text upper case is ABC" {
    when
        Fact.Text.ToUpper() == "Abc".ToUpper()
    then
        Fact.DoSomething();
}
```

### string.Trim() string

`Trim` will return a string contains trimmed version of receiver

#### Returns

* string trimmed string

#### Example

```Shell
rule TrimText  "Do something when Fact.Text upper case is ABC" {
    when
        Fact.Text == "  Abc   ".Trim().ToUpper()
    then
        Fact.DoSomething();
}
```

### array.Len() int

`Len` will return array/slice's length.

#### Returns

* The length of array/slice's receiver

#### Example

```Shell
rule DoSomething  "Do something when array length is sufficient" {
    when
        Fact.ChildrenArray.Len() > 2
    then
        Fact.DoSomething();
}
```

### array.Append(val) 

`Append` will append val into end of receiver array

#### Arguments

* `val` value to append


#### Example

```Shell
rule DoSomething  "Do something when array length is sufficient" {
    when
        Fact.ChildrenArray.Len() < 2
    then
        Fact.ChildrenArray.Append(Fact.NewChild());
}
```

### array.Clear() 

`Clear` will empty receiver array

#### Example

```Shell
rule DoSomething  "Do something when array length is sufficient" {
    when
        Fact.ChildrenArray.Len() > 2
    then
        Fact.ChildrenArray.Clear();
}
```

### map.Len() int
   
`Len` will return map's length.

#### Returns

* The length of map receiver

#### Example

```Shell
rule DoSomething  "Do something when map length is sufficient" {
   when
       Fact.ChildrenMap.Len() > 2
   then
       Fact.DoSomething();
}
```

### map.Clear() 

`Clear` will empty receiver map

#### Example

```Shell
rule DoSomething  "Do something when array length is sufficient" {
    when
        Fact.ChildrenMap.Len() > 2
    then
        Fact.ChildrenMap.Clear();
}
```

## Custom Functions

All functions which are invocable from the DataContext is **Invocable** from within the rule,
both in the "When" scope and the "Then" scope.

Thus, you can create functions and have your Fact as receiver, and those function can be called from GRL.

Assuming you have a struct with some functions.

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

And add the struct into knowledge base

```go
dctx := grule.context.NewDataContext()
dctx.Add("Pogo", &MyPoGo{})
```

You can call the function within the rule

```go
when
    Pogo.GetStringLength(some.variable) < 100
then
    some.variable = Pogo.AppendString(some.variable, "Groooling");
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
