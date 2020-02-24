# Function in Grule

[Tutorial](Tutorial_en.md) | [Rule Engine](RuleEngine_en.md) | [GRL](GRL_en.md) | [RETE Algorithm](RETE_en.md) | [Functions](Function_en.md) | [Grule Events](GruleEvent_en.md)


## Build-In Functions

All build-in function are all defined within the `ast/BuildInFunctions.go` File. As of now, they are :

### MakeTime(year, month, day, hour, minute, second int64) time.Time 

`MakeTime` will create a `time.Time` with local `locale`.

#### Arguments:

* `year` is the Year number.
* `month` is the Month number, January = 1.
* `day` is the day number in a month.
* `hour` the hour of the day starting from 0.
* `minute` the minute of the hour starting from 0.
* `second` the second of the minute starting from 0.

#### Returns

* `time.Time` value representing the time as specified in the argument in `local` locale.

#### Example

```text
rule SetExpire "Set the expire date for Fact created before 2020" {
    when 
       Fact.CreateTime < MakeTime(2020,1,1,0,0,0)
    then
       Fact.ExpireTime = MakeTime(2021,1,1,0,0,0);
}
``` 

### Changed(variableName string)

`Changed` will make sure the specified variable name is removed from working memory.

#### Arguments:

* `variableName` the variable name to be removed from working memory.

#### Example

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

`Now` function will create a new `time.Time` value containing the current time.

#### Returns

* `time.Time` value coontaining the current value

#### Example

```text
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

* `text` The text to emit into Log-Debug

#### Example

```text
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

`IsZero` will check any variable in the argument for its `Zero` status value. Zero means
that the variable is newly defined and have not assigned with any value.
This is usually applied to types like `string`, `int64`, `uint64`, `bool`, `time.Time`, etc. 

#### Arguments

* `i` a variable to check.

#### Returns

* `true` if the specified argument is Zero.
* `false` if the specified argument not Zero.

#### Example

```text
rule CheckStartTime "Check device's starting time." {
    when
        IsZero(Device.StartTime) == true
    then
        Device.StartTime = Now();
}
```

### Retract(ruleName string)

`Retract` will retract the specified rule from next cycle evaluation. If a rule
is retracted, it's `when` scope will not be evaluated. When the rule engine execute
again, all retract status of all rules will be restored.

#### Arguments

* `ruleName` name of the rule to retract.

#### Example

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

`GetTimeYear` will extract the Year value of the time argument.

#### Arguments

* `time` The time variable

#### Returns

* Year value of the time.

#### Example

```text
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

```text
rule StartNewYearProcess "Check if its a new year to restart new FinancialYear." salience 1000 {
    when
        isZero(Process.Month)        
    then
        Process.Month = GetTimeMonth(Process.Month);
}
```

### GetTimeDay(time time.Time) int

`GetTimeDay` will extract the Day of Month value of the time argument.

#### Arguments

* `time` The time variable

#### Returns

* Day of month value of the time. 

#### Example

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


`GetTimeHour` will extract the Hour value of the time argument.

#### Arguments

* `time` The time variable

#### Returns

* Hour value of the time. Is between 0 to 23 

#### Example

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


`GetTimeMinute` will extract the Minute value of the time argument.

#### Arguments

* `time` The time variable

#### Returns

* Minute value of the time. Is between 0 to 59 

#### Example

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

`GetTimeSecond` will extract the Second value of the time argument.

#### Arguments

* `time` The time variable

#### Returns

* Second value of the time. Is between 0 to 59 

#### Example

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


`IsTimeBefore` will check if a time value is before the other time value.

#### Arguments

* `time` The time variable
* `before` Another time variable

#### Returns

* True if the `before` time value is before the `time` value.
* False if the `before` time value is not before the `time` value.
  
#### Example

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


`IsTimeAfter` will check if a time value is after the other time value.

#### Arguments

* `time` The time variable
* `after` Another time variable

#### Returns

* True if the `after` time value is after the `time` value.
* False if the `after` time value is not after the `time` value.
  
#### Example

```text
rule AdditionalTax  "Apply additional tax if purchase after date specified." {
    when
        IsTimeAfter(Purchase.TransactionTime, TaxRegulation.StartSince)
    then
        Purchase.Tax = PurchaseTax + 0.01;
}
```

### TimeFormat(time time.Time, layout string) string


`TimeFormat` will check if a time value is after the other time value.

#### Arguments

* `time` The time variable
* `layout` String variable specifying the date format layout.

For layout format, you can [read this article](https://yourbasic.org/golang/format-parse-string-time-date-example/)

#### Returns

* String contains the formatted time
  
#### Example

```text
rule LogPurchaseDate  "Log the purchase date." {
    when
        IsZero(Purchase.TransactionDate) == false
    then
        Log(TimeFormat(Purchase.TransactionDate, "2006-01-02T15:04:05-0700");
}
```

## Custom Functions

All invocable functions which are invocable from the DataContext is **Invocable** from within the rule,
both in the "When" scope and "Then" scope.

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

You can call the fuction within the rule

```go
when
    Pogo.GetStringLength(some.variable) < 100
then
    some.variable = Pogo.AppendString(some.variable, "Groooling");
```

### Important Thing you must know about Custom Function in Grule

When you make your own function to be called from rule engine, you need to know the following rules.

1. The function must be visible. Public function convention should start with capital letter. Private functions cant be executed.
2. The function must only return 1 type. Returning multiple variable from function are not acceptable, the rule execution will fail if there are multiple return variable.
3. The way number literal were treated in Grule's DRL is; **decimal** will always taken as `int64` and **real** as `float64`, thus always consider to define your function argument and returns between `int64` and `float64` when your function works with numbers.
