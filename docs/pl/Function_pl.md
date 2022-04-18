# Funkcja w Grule

[![Function_cn](https://github.com/yammadev/flag-icons/blob/master/png/CN.png?raw=true)](../cn/Function_cn.md)
[![Function_de](https://github.com/yammadev/flag-icons/blob/master/png/DE.png?raw=true)](../de/Function_de.md)
[![Function_en](https://github.com/yammadev/flag-icons/blob/master/png/GB.png?raw=true)](../en/Function_en.md)
[![Function_id](https://github.com/yammadev/flag-icons/blob/master/png/ID.png?raw=true)](../id/Function_id.md)
[![Function_pl](https://github.com/yammadev/flag-icons/blob/master/png/PL.png?raw=true)](../id/Function_pl.md)

[About](About_pl.md) | [Tutorial](Tutorial_pl.md) | [Rule Engine](RuleEngine_pl.md) | [GRL](GRL_pl.md) | [GRL JSON](GRL_JSON_pl.md) | [RETE Algorithm](RETE_pl.md) | [Functions](Function_pl.md) | [FAQ](FAQ_pl.md) | [Benchmark](Benchmarking_pl.md)

---

## Funkcje wbudowane

Wszystkie wbudowane funkcje są zdefiniowane w pliku `ast/BuiltInFunctions.go`. Na chwilę obecną są to:

### MakeTime(year, month, day, hour, minute, second int64) time.Time

`MakeTime` spowoduje utworzenie `time.Time` z lokalnymi `locale`.

#### Argumenty

* `year` to numer roku.
* `month` to numer miesiąca, Styczeń = 1.
* `day` to numer dnia w miesiącu.
* `hour` godzina w ciągu dnia, zaczynając od 0.
* `minute` minuta w godzinie, zaczynając od 0.
* `second` sekunda w minucie, zaczynając od 0.

#### Zwraca

* `time.Time` wartość reprezentująca czas określony w argumencie w `local`.

#### Przykład

```Shell
rule SetExpire "Set the expire date for Fact created before 2020" {
    when
       Fact.CreateTime < MakeTime(2020,1,1,0,0,0)
    then
       Fact.ExpireTime = MakeTime(2021,1,1,0,0,0);
}
```

### Changed(variableName string)

`Changed` sprawi, że podana `variableName` zostanie usunięta z pamięci roboczej przed następnym cyklem.

#### Argumenty

* `variableName` nazwa zmiennej, która ma być usunięta z pamięci roboczej.

#### Przykład

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

Funkcja `Now` utworzy nową wartość `time.Time` zawierającą aktualny czas.

#### Zwraca

* `time.Time` wartość reprezentującą bieżącą wartość.

#### Przykład

```Shell
rule ResetTime "Reset the lastUpdate time" {
    when
        Fact.LastUpdate < Now()
    then
        Fact.LastUpdate = Now();
}
```

### Log(text string)

`Log` spowoduje wyemitowanie łańcucha log-debug z wnętrza reguły.

#### Argumenty

* `text` Tekst do wysłania do Log-Debug

#### Przykład

```Shell
rule SomeRule "Log candidate name if he is below 17 years old" {
    when
        Candidate.Age < 17
    then
        Log("Under aged: " + Candidate.Name);
}
```

### IsNil(i interface{}) bool

`IsNil` sprawdzi, czy argument jest wartością `nil`.

#### Argumenty

* `i` zmienna do sprawdzenia.

#### Zwraca.

* `true` jeśli podany argument jest `nil` lub niepoprawną wartością `ptr`.
* `false` jeśli podany argument jest poprawną wartością `ptr`.

#### Przykład

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

`IsZero` sprawdzi każdą zmienną w argumencie pod kątem jej wartości `Zero`. Zero oznacza, że zmienna jest nowo zdefiniowana i nie ma przypisanej wartości początkowej. Zwykle stosuje się to do typów takich jak `string`, `int64`, `uint64`, `bool`, `time.Time`, itp.

#### Argumenty

* `i` zmienna do sprawdzenia.

#### Zwraca.

* `true` jeśli podany argument jest zerem.
* `false` jeśli podany argument nie jest zerem.

#### Przykład

```Shell
rule CheckStartTime "Check device's starting time." {
    when
        IsZero(Device.StartTime) == true
    then
        Device.StartTime = Now();
}
```

### Retract(ruleName string)

`Retract` wykluczy podaną regułę z ewaluacji w kolejnych cyklach. Jeśli reguła jest wycofana, to jej zakres `when` nie będzie oceniany w następnych cyklach po wywołaniu `Retract`. Silnik automatycznie przywróci wszystkie reguły na swoje miejsce, gdy zacznie ponownie od początku.

#### Argumenty

* `ruleName` nazwa reguły do wycofania.

#### Przykład

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

Funkcja `GetTimeYear` wyodrębni wartość Year z argumentu time.

#### Argumenty

* `time` Zmienna czasowa

#### Zwraca

* Wartość roku w czasie.

#### Przykład

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

#### Argumenty

* `time` Zmienna czasowa

#### Zwraca.

* Wartość miesiąca czasu. 1 = styczeń.

#### Przykład

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

Polecenie `GetTimeDay` wyodrębni dzień miesiąca z argumentu time.

#### Argumenty

* `time` Zmienna czasowa

#### Zwraca

* Wartość dnia miesiąca dla czasu.

#### Przykład

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

Polecenie `GetTimeHour` wyodrębni wartość godziny z argumentu time.

#### Argumenty

* `time` Zmienna czasowa

#### Zwraca

* Wartość godziny czasu. Mieści się w przedziale od 0 do 23

#### Przykład

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

Polecenie `GetTimeMinute` wyodrębni wartość minutową argumentu time.

#### Argumenty

* `time` Zmienna czasowa

#### Zwraca

* Minutowa wartość czasu, od 0 do 59.

#### Przykład

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

Polecenie `GetTimeSecond` wyodrębni drugą wartość argumentu time.

#### Argumenty

* `time` Zmienna czasowa

#### Zwraca

* Drugą wartość czasu, z zakresu od 0 do 59.

#### Przykład

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

`IsTimeBefore` sprawdza, czy wartość czasowa poprzedza inną wartość czasową.

#### Argumenty

* `time` Wartość czasowa, która ma zostać sprawdzona
* `before` Wartość czasowa, względem której sprawdzana jest powyższa wartość.

#### Zwraca

* True jeśli wartość czasowa `before` poprzedza wartość `time`.
* False jeśli wartość czasu `before` nie poprzedza wartości `time`.

#### Przykład

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

`IsTimeAfter` sprawdza, czy wartość czasowa następuje po innej wartości czasowej.

#### Argumenty

* `time` Wartość czasowa, która ma zostać sprawdzona
* `after` Wartość czasowa, względem której sprawdzana jest powyższa wartość.

#### Zwraca

* True jeśli wartość czasu `after` podąża za wartością `time`.
* False jeśli wartość czasu `after` nie podąża za wartością `time`.

#### Przykład

```Shell
rule AdditionalTax "Apply additional tax if new tax rules are in effect." {
    when
        IsTimeAfter(Purchase.TransactionTime, TaxRegulation.StartSince)
    then
        Purchase.Tax = Purchase.Tax + 0.01;
}
```

### TimeFormat(time time.Time, layout string) string

Funkcja `TimeFormat` sformatuje argument czasu w sposób określony przez argument `layout`.

#### Argumenty

* `time` Wartość czasu, która ma zostać sformatowana.
* `layout` Zmienna łańcuchowa określająca układ formatu daty.

Aby uzyskać informacje na temat formatu układu, można [przeczytać ten artykuł](https://yourbasic.org/golang/format-parse-string-time-date-example/)

#### Zwraca

* Łańcuch znaków sformatowany w określony sposób.

#### Przykład

```Shell
rule LogPurchaseDate "Log the purchase date." {
    when
        IsZero(Purchase.TransactionDate) == false
    then
        Log(TimeFormat(Purchase.TransactionDate, "2006-01-02T15:04:05-0700");
}
```

### Complete()

Polecenie `Complete` spowoduje, że silnik przestanie przetwarzać kolejne reguły w bieżącym cyklu. Jest to przydatne, gdy chcemy zakończyć dalsze przetwarzanie reguł pod określonym warunkiem.

#### Przykład

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

## Funkcje matematyczne

Wszystkie poniższe funkcje są opakowaniem ich funkcji matematycznych Golanga.
Powinieneś przeczytać stronę Golanga o matematyce, aby dowiedzieć się, jak używać każdej z funkcji.

W przeciwieństwie do go, nie musisz używać prefiksu `math.`, aby użyć ich w swoim GRL.

Używaj ich jak normalnych wbudowanych funkcji.

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


## Funkcje stałe

Poniższe funkcje mogą być wywoływane z poziomu GRL, o ile odbiornik typ wartości jest poprawny.

### string.Len() int

`Len` zwróci długość łańcucha.

#### Zwraca

* Długość odbiornika łańcucha znaków.

#### Przykład

```Shell
rule DoSomething "Do something when string length is sufficient" {
    when
        Fact.Name.Len() > "ATextConstant".Len()
    then
        Fact.DoSomething();
}
```

### string.Compare(string) int

Polecenie `Compare` porówna łańcuch odbiorcy z argumentem.

#### Argumenty

* `string` Łańcuch do porównania

#### Zwraca

* `< 0` jeśli odbiornik jest mniejszy od argumentu
* `0` jeśli odbiornik jest równy argumentowi
* `> 0` jeśli odbiornik jest większy od argumentu

#### Przykład

```Shell
rule CompareString "Do something when Fact.Text is greater than A" {
    when
        Fact.Text.Compare("A") > 0
    then
        Fact.DoSomething();
}
```

### string.Contains(string) bool

`Contains` sprawdzi, czy jego argument jest zawarty w odbiorniku.

#### Argumenty

* `string` Podłańcuch do sprawdzenia w odbiorniku

#### Zwraca

* `true` jeśli argument jest zawarty w odbiorniku.
* `false` jeśli argument nie jest zawarty w odbiorniku.

#### Przykład

```Shell
rule ContainString "Do something when Fact.Text is contains XXX" {
    when
        Fact.Text.Contains("XXX")
    then
        Fact.DoSomething();
}
```

### string.In(string ...) bool

`In` sprawdzi, czy którykolwiek z argumentów jest równy odbiornikowi.

#### Argumenty

* `string` Zmienny argument łańcuchowy do sprawdzenia

#### Zwraca

* bolean `true` jeśli którykolwiek z argumentów jest równy odbiornikowi, lub `false` w przeciwnym przypadku.

#### Przykład

```Shell
rule CheckArgumentIn "Do something when Fact.Text is equals to 'ABC' or 'BCD' or 'CDE' " {
    when
        Fact.Text.In("ABC", "BCD", "CDE")
    then
        Fact.DoSomething();
}
```

### string.Count(string) int

`Count` policzy liczbę wystąpień argumentu w łańcuchu odbiornika.

#### Argumenty

* `string` Podłańcuch do policzenia w odbiorniku

#### Zwraca

* liczba wystąpień argumentu w odbiorniku.

#### Przykład

```Shell
rule CountString "Do something when Fact.Text contains 3 occurrences of 'ABC'" {
    when
        Fact.Text.Count("ABC") == 3
    then
        Fact.DoSomething();
}
```

### string.HasPrefix(string) bool

`HasPrefix` sprawdzi, czy łańcuch odbiornika ma określony prefiks.

#### Argumenty

* `string` Oczekiwany prefiks.

#### Zwraca

* `true` jeśli odbiornik ma dany argument jako swój przedrostek.
* `false` jeśli odbiornik nie ma tego argumentu jako swojego przedrostka.

#### Przykład

```Shell
rule IsPrefixed "Do something when Fact.Text started with PREF" {
    when
        Fact.Text.HasPrefix("PREF")
    then
        Fact.DoSomething();
}
```

### string.HasSuffix(string) bool

`HasSuffix` sprawdzi, czy łańcuch odbiornika ma określony sufiks.

#### Argumenty

* `string` Oczekiwany sufiks.

#### Zwraca

* `true` jeśli odbiornik ma dany argument jako swój sufiks.
* `false`, jeśli odbiornik nie ma tego argumentu jako przyrostka.

#### Przykład

```Shell
rule IsSuffixed "Do something when Fact.Text ends with SUFF" {
    when
        Fact.Text.HasSuffix("SUFF")
    then
        Fact.DoSomething();
}
```

### string.Index(string) int

`Index` zwróci indeks pierwszego wystąpienia argumentu w łańcuchu odbiornika.

#### Argumenty

* `string` Podłańcuch do wyszukania.

#### Zwraca

* Wartość indeksu pierwszego wystąpienia argumentu.

#### Przykład

```Shell
rule IndexCheck "Do something when Fact.Text ABC occurs as specified" {
    when
        Fact.Text.Index("ABC") == "abABCabABC".Index("ABC")
    then
        Fact.DoSomething();
}
```

### string.LastIndex(string) int

`LastIndex` zwróci indeks ostatniego wystąpienia argumentu w łańcuchu odbiornika.

#### Argumenty

* `string` Podłańcuch do wyszukania.

#### Zwraca

* Indeks ostatniego wystąpienia argumentu.

#### Przykład

```Shell
rule LastIndexCheck "Do something when Fact.Text ABC occurs in the last position as specified" {
    when
        Fact.Text.LastIndex("ABC") == "abABCabABC".LastIndex("ABC")
    then
        Fact.DoSomething();
}
```

### string.Repeat(int64) string

Polecenie `Powtórz` zwróci łańcuch zawierający `n` wystąpień łańcucha odbiorcy.

#### Argumenty

* `int64` liczba powtórzeń

#### Zwraca

* Nowy łańcuch zawierający `n` wystąpień łańcucha odbiorcy.

#### Przykład

```Shell
rule StringRepeat "Do something when Fact.Text contains ABCABCABC" {
    when
        Fact.Text == "ABC".Repeat(3)
    then
        Fact.DoSomething();
}
```

### string.Replace(old, new string) string

Polecenie `Replace` zwróci łańcuch ze wszystkimi wystąpieniami `starego` zastąpionymi `nowym`.

#### Argumenty

* `old` podłańcuch, który ma zostać zastąpiony.
* `new` łańcuch, który chcesz zastąpić wszystkimi wystąpieniami `old`.

#### Zwraca

* Łańcuch, w którym wszystkie wystąpienia `old` w odbiorniku zostały zastąpione przez `new`.

#### Przykład

```Shell
rule ReplaceString "Do something when Fact.Text contains replaced string" {
    when
        Fact.Text == "ABC123ABC".Replace("123","ABC")
    then
        Fact.DoSomething();
}
```

### string.Split(string) []string

`Split` zwróci wycinek łańcucha, którego elementy są określone po podziale odbiornika przez argument string token.  Token nie będzie występował w elementach wynikowego wycinka.

#### Argumenty

* `string` token, którego chcesz użyć do podziału odbiornika.

#### Zwraca.

* Plaster zawierający fragmenty oryginalnego łańcucha podzielone przez token.

#### Przykład

```Shell
rule SplitString "Do something when Fact.Text is prefixed by 'ABC,'" {
    when
        Fact.Text.Split(",")[0] == "ABC"
    then
        Fact.DoSomething();
}
```

### string.ToLower() string

`ToLower` zwróci łańcuch, którego zawartość stanowią wszystkie małe litery znaków w odbiorniku.

#### Zwraca

* Nowy łańcuch będący wersją odbiornika napisaną małymi literami.

#### Przykład

```Shell
rule LowerText "Do something when Fact.Text is equal to 'abc'" {
    when
        Fact.Text.ToLower() == "Abc".ToLower()
    then
        Fact.DoSomething();
}
```

### string.ToUpper() string

`ToUpper` zwróci łańcuch, którego zawartość stanowią wszystkie duże litery znaków znaków w odbiorniku.

#### Zwraca

* Nowy łańcuch będący wersją odbiornika napisaną dużymi literami.

#### Przykład

```Shell
rule UpperText "Do something when Fact.Text is equal to 'ABC'" {
    when
        Fact.Text.ToUpper() == "Abc".ToUpper()
    then
        Fact.DoSomething();
}
```

### string.Trim() string

`Trim` zwróci łańcuch, z którego usunięto białe spacje na obu końcach.

#### Zwraca

* Łańcuch znaków z usuniętymi białymi spacjami z początku i końca.

#### Przykład

```Shell
rule TrimText "Do something when Fact.Text is 'ABC'" {
    when
        Fact.Text == "  Abc   ".Trim().ToUpper()
    then
        Fact.DoSomething();
}
```


### string.MatchString() string

`MatchString` MatchString informuje, czy łańcuch s zawiera dowolne dopasowanie wzorca wyrażenia regularnego. Podobne do golang [MatchString](https://pkg.go.dev/regexp#MatchString)

#### Zwraca

* True jeśli `regexPattern` pasuje do łańcucha s
* False, jeśli `regexPattern` nie pasuje do łańcucha s.

#### Przykład

```Shell
rule MatchStringText "Return true when regex pattern matches the string"  {
	when
	  Fact.Text.MatchString("B([a-z]+)ck")
	then
	  Fact.DoSomething();
}
```

### array.Len() int

`Len` zwróci długość tablicy/plastra.

#### Zwraca

* Długość tablicy/układu.

#### Przykład

```Shell
rule DoSomething "Do something when array length is sufficient" {
    when
        Fact.ChildrenArray.Len() > 2
    then
        Fact.DoSomething();
}
```

### array.Append(val) 

Polecenie `Append` powoduje dołączenie `val` na końcu tablicy odbiorników.

#### Argumenty

* `val` wartość, która ma być dołączona.

#### Przykład

```Shell
rule DoSomething "Add a new child when the array has less than 2 children" {
    when
        Fact.ChildrenArray.Len() < 2
    then
        Fact.ChildrenArray.Append(Fact.NewChild());
}
```

### map.Len() int
   
`Len` zwróci długość mapy.

#### Zwraca

* Długość odbiornika mapy.

#### Przykład

```Shell
rule DoSomething "Do something when map length is sufficient" {
   when
       Fact.ChildrenMap.Len() > 2
   then
       Fact.DoSomething();
}
```

## Funkcje niestandardowe

Wszystkie funkcje dostępne z poziomu DataContext są **nieodwołalne** z poziomu reguły, zarówno w zakresie "When" jak i "Then".

Możesz tworzyć funkcje, których odbiorcą jest Twój Fakt, a funkcje te mogą być wywoływane z poziomu GRL.

Na przykład. Podane:

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

Użytkownik może wywoływać zdefiniowane metody:

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

### Argumenty funkcji zmiennopozycyjnych

W funkcjach niestandardowych obsługiwane są argumenty zmienne.

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

Funkcję tę można następnie wywołać z wnętrza reguły, podając zero lub więcej wartości dla argumentu variadic.

```go
when
    Pogo.GetStringLength(some.variable) < 100
then
    some.longest = Pogo.GetLongestString(some.stringA, some.stringB, some.stringC);
```

Ponieważ możliwe jest podanie wartości zerowych w celu spełnienia argumentu variadic, można je również wykorzystać do symulacji parametrów opcjonalnych.

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

### Prawa funkcji specjalnych w Grule

Gdy tworzysz własną funkcję, która ma być wywoływana z silnika reguł, musisz znać następujące prawa:

1. Funkcja musi być widoczna, co oznacza, że funkcje muszą zaczynać się od dużej litery. Funkcje prywatne nie mogą być wykonywane.
2. Funkcja musi zwracać tylko jeden typ wartości. Zwracanie wielu wartości z funkcji nie jest obsługiwane, a wykonanie reguły nie powiedzie się, jeśli zwróconych zostanie wiele wartości.
3. Sposób traktowania literałów liczbowych w GRL Grule'a jest taki, że **integer** zawsze będzie traktowany jako typ `int64`, a **real** jako `float64`, dlatego zawsze musisz odpowiednio zdefiniować swoje typy liczbowe.
