# Grule Rule Language (GRL)

[![GRL_cn](https://github.com/yammadev/flag-icons/blob/master/png/CN.png?raw=true)](../cn/GRL_cn.md)
[![GRL_de](https://github.com/yammadev/flag-icons/blob/master/png/DE.png?raw=true)](../de/GRL_de.md)
[![GRL_en](https://github.com/yammadev/flag-icons/blob/master/png/GB.png?raw=true)](../en/GRL_en.md)
[![GRL_id](https://github.com/yammadev/flag-icons/blob/master/png/ID.png?raw=true)](../id/GRL_id.md)
[![GRL_pl](https://github.com/yammadev/flag-icons/blob/master/png/PL.png?raw=true)](../id/GRL_pl.md)

[About](About_pl.md) | [Tutorial](Tutorial_pl.md) | [Rule Engine](RuleEngine_pl.md) | [GRL](GRL_pl.md) | [GRL JSON](GRL_JSON_pl.md) | [RETE Algorithm](RETE_pl.md) | [Functions](Function_pl.md) | [FAQ](FAQ_pl.md) | [Benchmark](Benchmarking_pl.md)

---

**GRL** jest językiem DSL (Domain Specific Language) zaprojektowanym dla Grule. Jest to uproszczony język służący do definiowania kryteriów warunków reguł oraz akcji, które zostaną wykonane, jeśli kryteria te zostaną spełnione.

Język ten ma następującą strukturę:

```Shell
rule <RuleName> <RuleDescription> [salience <priority>] {
    when
        <boolean expression>
    then
        <assignment or operation expression>
}
```

**RuleName**: Identyfikuje konkretną regułę. Nazwa musi być unikalna w całej bazie wiedzy, składać się z jednego słowa i nie może zawierać białych spacji.

**RuleDescription**: Opisuje regułę przeznaczoną do spożycia przez ludzi. Opis powinien być ujęty w cudzysłów.

**Salience** (opcjonalne, domyślnie 0): Określa ważność reguły. Niższe wartości oznaczają reguły o niższym priorytecie. Wartość ważności jest używana do określenia kolejności sortowania według priorytetu w przypadku napotkania wielu reguł. Salience przyjmuje wartości ujemne, więc można je wykorzystać do oznaczania reguł, które nas w ogóle nie obchodzą. Silniki reguł są *deklaratywne*, więc nie można zagwarantować, w jakiej kolejności będą przetwarzane reguły.  W związku z tym należy traktować `salience` jako *podpowiedź* dla silnika, która pomoże mu zdecydować, co zrobić w przypadku konfliktu.

**Wyrażenie boolowskie**: Wyrażenie predykatowe, które zostanie obliczone przez silnik reguł w celu określenia, czy dana akcja reguły jest kandydatem do wykonania przy obecnym stanie faktycznym.

**Wyrażenie przypisania lub operacji**: Jest to akcja, która zostanie podjęta, jeśli reguła zostanie oceniona jako `true`. Nie jesteś ograniczony do pojedynczego wyrażenia i możesz podać ich listę, oddzielając je znakiem `;`. Wyrażenia akcji służą do modyfikowania bieżących wartości faktów, wykonywania obliczeń, rejestrowania pewnych stwierdzeń itd...

### Wyrażenie booleańskie

Wyrażenie boolean powinno być znane większości, jeśli nie wszystkim programistom.

```go
when
     contains(User.Name, "robert") &&
     User.Age > 35
then
     ...
```

### Stałe i znaki literowe

| Literal | Opis                                                                                       | Przykład                                          |
| ------- | ------------------------------------------------------------------------------------------ | ------------------------------------------------- |
| String  | Przechowuje literalny ciąg znaków, ujęty w podwójny (&quot;) lub pojedynczy (') cudzysłów. | "To jest ciąg znaków" lub "to jest ciąg znaków"   |
| Integer | Przechowuje wartość całkowitą i może być poprzedzona symbolem ujemnym -.                   | `1` lub `34` lub `42344` lub `-553`               |
| Real    | Przechowuje wartość rzeczywistą                                                            | `234.4553`, `-234.3`, `314E-2`, `.32`, `12.32E12` |
| Boolean | Przechowuje wartość typu boolean                                                           | `true`, `TRUE`, `False`                           |

Więcej przykładów można znaleźć na stronie [GRL Literals](GRL_Literals_pl.md).

Uwaga: Znaki specjalne w łańcuchach muszą być usuwane zgodnie z tymi samymi zasadami, które są stosowane dla łańcuchów w języku Go.  Łańcuchy z backtickiem nie są jednak obsługiwane.

### Obsługiwane operatory 

| Typ                  | Operator                          |
| -------------------- | --------------------------------- |
| Matematyka           | `+`, `-`, `/`, `*`, `%`           |
| Operatory bit-wise   | `\|`, `&`                         |
| Operatory logiczne   | `&&`, `\|\|`                      |
| Operatory porównania | `<`, `<=`, `>`, `>=`, `==`, `!=`  |

### Pierwszeństwo operatora

Grule jest zgodne z pierwszeństwem operatorów w Go.

| Pierwszeństwo | Operator                         |
| ------------- | -------------------------------- |
|    5          | `*`, `/`, `%`, `&`               |
|    4          | `+`, `-`, `\|`                   |
|    3          | `==`, `!=`, `<`, `<=`, `>`, `>=` |
|    2          | `&&`                             |
|    1          | `\|\|`                           |

### Komentarze

Komentarze również są zgodne ze standardowym formatem Go.

```go
// This is a comment
// And this

/* And also this */

/*
   As well as this
*/
```

### Tablica/plaster i mapa

Od wersji 1.6.0 Grule obsługuje dostęp do faktów w postaci tablicy/plastra lub mapy.

Załóżmy, że masz strukturę faktów jak poniżej:

```go
type MyFact struct {
    AnIntArray   []int
    AStringArray []string
    SubFacts     []*MyFact
    SubMaps      map[string]*MyFact
}
```

Te plastry i mapy można oceniać za pomocą reguł:

```go
    when 
       Fact.AnIntArray[1] == 12 &&
       Fact.AStringArray[12] != "SomeText" &&
       Fact.SubFacts[1].SubFacts[2].AnIntArray[12] > 100 &&
       Fact.SubMaps["Key"].AnIntArray[0] == 1000
    then
       ...
```

Jeśli reguła spróbuje uzyskać dostęp do elementu tablicy, który znajduje się poza jej granicami, spowoduje to przerwanie wykonywania reguły.

#### Przypisywanie wartości do tablic, plasterków i map

Wartość tablicy można ustawić, jeśli podany indeks jest poprawny.

```go
   then
      Fact.AnIntArray[10] = 12;
      Fact.SubMap["AKey"].AStringArray[1] = "New Value";
      Fact.AnotherMap[Fact.SomeFunction()] = "Another Value";
```

Istnieje kilka funkcji, których można użyć do pracy z tablicami/plastrami i mapami.
Można je znaleźć na stronie [Function page](Function_pl.md).

### Negacja

Symbol negacji jednoargumentowej `!` jest obsługiwany przez GRL jako dodatek do symbolu NEQ `!=`.
Należy go używać przed wyrażeniem booleańskim lub atomem wyrażenia.

Na przykład w wyrażeniu atom:

```go
when 
    !FunctionReturnTrue() ||
    !false
then
    ... 
```

lub w wyrażaniu:

```go
when
    !(you.IsOk() || !today.isMonday())
then
    ...
```

### Wywołanie funkcji

Z reguły można wywołać dowolną widoczną funkcję, o ile zwraca ona wartość 0 lub 1.  Na przykład:

```go
    when
        Fact.FunctionA() == "text" ||
        Fact.FunctionB("arg") == "text" ||
        Fact.FunctionC(Fact.Field, true)
    then
        Fact.CallFunction();
        Fact.Value = Fact.CallFunction();
        ...
```

W wersji 1.6.0 Grule może tworzyć łańcuchy wywołań funkcji i akcesorów wartości.  Na przykład;

```go
    when
        Fact.Function().StringField == "" ||
        Fact.Function("contant").ObjField.OtherFunction() &&
        ...
    then
        Fact.CallFunction().CallAnotherFunction();
        ...
```

W wersji 1.6.0 wprowadzono również możliwość wywoływania funkcji na stałych literalnych.  Na przykład:

```go
    when
        "AString   ".Trim().ToUpper().HasSuffix("ING")
    then
        Fact.Result = Fact.ReturnStringFunc().Trim().ToLower();
```

Listę dostępnych funkcji zawiera sekcja [Strona funkcji](Function_pl.md).

#### Przykłady

```go
rule SpeedUp "When testcar is speeding up we keep increase the speed."  {
    when
        TestCar.SpeedUp == true && TestCar.Speed < TestCar.MaxSpeed
    then
        TestCar.Speed = TestCar.Speed + TestCar.SpeedIncrement;
            DistanceRecord.TotalDistance = DistanceRecord.TotalDistance + TestCar.Speed;
}

rule StartSpeedDown "When testcar is speeding up and over max speed we change to speed down."  {
    when
        TestCar.SpeedUp == true && TestCar.Speed >= TestCar.MaxSpeed
    then
        TestCar.SpeedUp = false;
            log("Now we slow down");
}

rule SlowDown "When testcar is slowing down we keep decreasing the speed."  {
    when
        TestCar.SpeedUp == false && TestCar.Speed > 0
    then
        TestCar.Speed = TestCar.Speed - TestCar.SpeedIncrement;
        DistanceRecord.TotalDistance = DistanceRecord.TotalDistance + TestCar.Speed;
}

rule SetTime "When Distance Recorder time not set, set it." {
    when
        isNil(DistanceRecord.TestTime)
    then
        log("Set the test time");
        DistanceRecord.TestTime = now();
}
```

### Debugowanie składni GRL

W swojej aplikacji możesz sprawdzić, czy skrypt lub fragment GRL zawiera błąd składni GRL.

```go
        RuleWithError := `
        rule ErrorRule1 "Rule with error"  salience 10{
            when
              Pogo.Compare(User.Name, "Calo")  
            then
              User.Name = "Success";
              Log(User.Name)
              Retract("AgeNameCheck");
        }
        `

	// Build normally
	err := ruleBuilder.BuildRuleFromResource("Test", "0.1.1", pkg.NewBytesResource([]byte(RuleWithError)))

	// If the err != nil something is wrong.
	if err != nil {
		// Cast the error into pkg.GruleErrorReporter with typecast checking.
		// Typecast checking is necessary because the err might not only parsing error.
		if reporter, ok := err.(*pkg.GruleErrorReporter); ok {
			// Lets iterate all the error we get during parsing.
			for i, er := range reporter.Errors {
				fmt.Printf("detected error #%d : %s\n", i, er.Error())
			}
		} else {
			// Well, its an error but not GruleErrorReporter instance. could be IO error.
			t.Error("There should be GruleErrorReporter")
			t.FailNow()
		}
	}
```

Spowoduje to wyprintowanie

```txt
detected error #0 : grl error on 8:6 missing ';' at 'Retract'
```


### Obsługa IDE

Visual Studio Code: [https://marketplace.visualstudio.com/items?itemName=avisdsouza.grule-syntax](https://marketplace.visualstudio.com/items?itemName=avisdsouza.grule-syntax)
