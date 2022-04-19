# Grule Short Tutorial

[![Tutorial_cn](https://github.com/yammadev/flag-icons/blob/master/png/CN.png?raw=true)](../cn/Tutorial_cn.md)
[![Tutorial_de](https://github.com/yammadev/flag-icons/blob/master/png/DE.png?raw=true)](../de/Tutorial_de.md)
[![Tutorial_en](https://github.com/yammadev/flag-icons/blob/master/png/GB.png?raw=true)](../en/Tutorial_en.md)
[![Tutorial_id](https://github.com/yammadev/flag-icons/blob/master/png/ID.png?raw=true)](../id/Tutorial_id.md)
[![Tutorial_pl](https://github.com/yammadev/flag-icons/blob/master/png/PL.png?raw=true)](../id/Tutorial_pl.md)

[About](About_pl.md) | [Tutorial](Tutorial_pl.md) | [Rule Engine](RuleEngine_pl.md) | [GRL](GRL_pl.md) | [GRL JSON](GRL_JSON_pl.md) | [RETE Algorithm](RETE_pl.md) | [Functions](Function_pl.md) | [FAQ](FAQ_pl.md) | [Benchmark](Benchmarking_pl.md)

---

## Przygotowanie

Należy pamiętać, że Grule używa Go 1.16.

Aby zaimportować Grule do swojego projektu:

```Shell
$ go get github.com/hyperjumptech/grule-rule-engine
```

Z Twojego `go` możesz zaimportować Grule.

```go
import (
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
) 
``` 

## Tworzenie struktury faktów

A `fact` w grule to **pointer** na instancję `struktury`.  Struktura może zawierać właściwości, tak jak każda normalna `struktura` Golanga, łącznie z każdą `metodą`, którą chcesz zdefiniować, pod warunkiem, że spełnia ona wymagania dla metod zdefiniowanych poniżej.  Na przykład:

```go
type MyFact struct {
    IntAttribute       int64
    StringAttribute    string
    BooleanAttribute   bool
    FloatAttribute     float64
    TimeAttribute      time.Time
    WhatToSay          string
}
```

Zgodnie z konwencją Golanga, Grule jest w stanie uzyskać dostęp tylko do tych **widocznych** atrybutów i metod, które są eksponowane z początkową wielką literą.

```go
func (mf *MyFact) GetWhatToSay(sentence string) string {
    return fmt.Sprintf("Let say \"%s\"", sentence)
}
```

**UWAGA:** Funkcje członkowskie podlegają następującym wymaganiom:

* Funkcja członkowska musi być **widoczna**; jej nazwa musi zaczynać się od dużej litery.
* Funkcja członkowska musi zwracać wartości `0` lub `1`. Nie jest obsługiwana więcej niż jedna wartość zwracana.
* Wszystkie liczbowe typy argumentów i zwracanych wartości muszą być ich 64-bitowymi wariantami, np. `int64`, `uint64`, `float64`.
* Funkcja członkowska nie powinna **zmieniać** wewnętrznego stanu Faktu. Algorytm nie jest w stanie automatycznie wykryć takich zmian, rzeczy stają się trudniejsze do zrozumienia i mogą wkradać się błędy.  Jeśli **MUSISZ** zmienić jakiś stan wewnętrzny Faktu, to możesz powiadomić o tym Grule, używając wbudowanej funkcji `Changed(varname string)`.

## Dodaj fakt do DataContext

Aby dodać fakt do `DataContext` musisz stworzyć instancję swojego `faktu`.

```go
myFact := &MyFact{
    IntAttribute: 123,
    StringAttribute: "Some string value",
    BooleanAttribute: true,
    FloatAttribute: 1.234,
    TimeAttribute: time.Now(),
}
```

Można utworzyć dowolną liczbę faktów.

Po utworzeniu faktów można dodać ich instancje do `DataContext`:

```go
dataCtx := ast.NewDataContext()
err := dataCtx.Add("MF", myFact)
if err != nil {
    panic(err)
}
```

### Tworzenie faktów z JSON

Dane JSON mogą być również używane do opisywania faktów w Grule od wersji 1.8.0.  Więcej szczegółów można znaleźć w [JSON as a Fact](JSON_Fact_pl.md).

## Tworzenie biblioteki wiedzy i dodawanie do niej reguł

`Knowledge Library` jest zbiorem niebieskich wydruków `KnowledgeBase`, a `KnowledgeBase` jest zbiorem wielu reguł pochodzących z definicji reguł załadowanych z wielu źródeł.  Używamy `RuleBuilder` do budowania instancji `KnowledgeBase`, a następnie dodawania ich do `KnowledgeLibrary`.

Źródłową postacią GRL może być:

* surowy łańcuch znaków
* zawartość pliku
* dokument w punkcie końcowym HTTP.

Użyjmy `RuleBuilder`, aby zacząć wypełniać naszą `KnowledgeBase`.

```go
knowledgeLibrary := ast.NewKnowledgeLibrary()
ruleBuilder := builder.NewRuleBuilder(knowledgeLibrary)
```

Następnie możemy zdefiniować podstawową regułę jako nieprzetworzony łańcuch w DSL:

```go
// lets prepare a rule definition
drls := `
rule CheckValues "Check the default values" salience 10 {
    when 
        MF.IntAttribute == 123 && MF.StringAttribute == "Some string value"
    then
        MF.WhatToSay = MF.GetWhatToSay("Hello Grule");
        Retract("CheckValues");
}
`
```

I wreszcie możemy użyć konstruktora, aby dodać definicję do `knowledgeLibrary` z zadeklarowanego `zasobu`:

```go
// Add the rule definition above into the library and name it 'TutorialRules'  version '0.0.1'
bs := pkg.NewBytesResource([]byte(drls))
err := ruleBuilder.BuildRuleFromResource("TutorialRules", "0.0.1", bs)
if err != nil {
    panic(err)
}
```

Biblioteka `KnowledgeLibrary` zawiera teraz bazę wiedzy `KnowledgeBase` o nazwie `TutorialRules` w wersji `0.0.1`. Aby wykonać tę konkretną regułę, musimy uzyskać jej instancję z `KnowledgeLibrary`. Zostanie to wyjaśnione w następnym rozdziale.

## Wykonywanie silnika reguł Grule

Aby uruchomić bazę wiedzy, musimy pobrać instancję tej bazy z `KnowledgeLibrary`.

```go
knowledgeBase := knowledgeLibrary.NewKnowledgeBaseInstance("TutorialRules", "0.0.1")
```

Każda instancja uzyskana z `knowledgeLibrary` jest unikalnym *klonem* bazowej `KnowledgeBase` *blueprint*.  Każda unikalna instancja posiada także swoją własną, odrębną `WorkingMemory`. Ponieważ żadna instancja nie dzieli stanu z inną instancją, można ich używać w dowolnym środowisku wielowątkowym, pod warunkiem, że nie wykonuje się pojedynczej instancji z wielu wątków jednocześnie.

Konstruowanie z wzorca `KnowledgeBase` zapewnia również, że nie wykonujemy ponownie obliczeń za każdym razem, gdy chcemy skonstruować instancję.  Praca obliczeniowa jest wykonywana tylko raz, co czyni klonowanie `AST` niezwykle wydajnym.

Teraz wykonajmy instancję `KnowledgeBase` używając przygotowanego `DataContext`.

```go
engine = engine.NewGruleEngine()
err = engine.Execute(dataCtx, knowledgeBase)
if err != nil {
    panic(err)
}
```

## Uzyskiwanie wyniku

Oto reguła, którą zdefiniowaliśmy powyżej, tak dla porównania:

```go
rule CheckValues "Check the default values" salience 10 {
    when 
        MF.IntAttribute == 123 && MF.StringAttribute == "Some string value"
    then
        MF.WhatToSay = MF.GetWhatToSay("Hello Grule");
        Retract("CheckValues");
}
```

Zakładając, że warunek jest spełniony (a jest), akcja zmodyfikuje atrybut `MF.WhatToSay`.  W celu zapewnienia, że reguła nie będzie natychmiast ponownie analizowana, jest ona `Rectract` ze zbioru.  W tym konkretnym przypadku, gdyby reguła nie zdołała tego zrobić, dopasowałaby się ponownie w następnym cyklu, i tak dalej, i tak dalej.  W końcu Grule zakończy pracę z błędem, ponieważ nie będzie w stanie dojść do ostatecznego wyniku.

W tym przypadku, wszystko co musisz zrobić, aby uzyskać wynik, to po prostu sprawdzić instancję `myFact` pod kątem modyfikacji, którą wprowadziła reguła:

```go
fmt.Println(myFact.WhatToSay)
// this should prints
// Lets Say "Hello Grule"
```
## Zasoby

GRL mogą być przechowywane w zewnętrznych plikach i istnieje wiele sposobów, aby uzyskać i załadować zawartość tych plików.

### Z pliku

```go
fileRes := pkg.NewFileResource("/path/to/rules.grl")
err := ruleBuilder.BuildRuleFromResource("TutorialRules", "0.0.1", fileRes)
if err != nil {
    panic(err)
}
```

Do pakietu można także wczytać wiele plików, używając ścieżek i wzorców globalnych:

```go
bundle := pkg.NewFileResourceBundle("/path/to/grls", "/path/to/grls/**/*.grl")
resources := bundle.MustLoad()
for _, res := range resources {
    err := ruleBuilder.BuildRuleFromResource("TutorialRules", "0.0.1", res)
    if err != nil {
        panic(err)
    }
}
```

### Z String lub ByteArray

```go
bs := pkg.NewBytesResource([]byte(rules))
err := ruleBuilder.BuildRuleFromResource("TutorialRules", "0.0.1", bs)
if err != nil {
    panic(err)
}
```

### Od URL

```go
urlRes := pkg.NewUrlResource("http://host.com/path/to/rule.grl")
err := ruleBuilder.BuildRuleFromResource("TutorialRules", "0.0.1", urlRes)
if err != nil {
    panic(err)
}
```

#### z nagłówkami

```go
headers := make(http.Header)
headers.Set("Authorization", "Basic YWxhZGRpbjpvcGVuc2VzYW1l")
urlRes := pkg.NewURLResourceWithHeaders("http://host.com/path/to/rule.grl", headers)
err := ruleBuilder.BuildRuleFromResource("TutorialRules", "0.0.1", urlRes)
if err != nil {
    panic(err)
}
```

### Z GIT

```go
bundle := pkg.NewGITResourceBundle("https://github.com/hyperjumptech/grule-rule-engine.git", "/**/*.grl")
resources := bundle.MustLoad()
for _, res := range resources {
    err := ruleBuilder.BuildRuleFromResource("TutorialRules", "0.0.1", res)
    if err != nil {
        panic(err)
    }
}
```

### Z JSON

Możesz teraz tworzyć reguły z JSON! [Przeczytaj, jak to działa](GRL_JSON_pl.md) 

## Skompiluj GRL do GRB

Jeśli chcesz uzyskać szybsze ładowanie zbioru reguł (np. masz bardzo duży zbiór reguł i ładowanie GRL jest zbyt wolne), możesz zapisać swój zbiór reguł do pliku GRB (Grules Rule Binary). [Przeczytaj, jak zapisywać i wczytywać GRB](Binary_Rule_File_pl.md) 
