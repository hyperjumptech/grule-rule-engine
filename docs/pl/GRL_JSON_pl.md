# Format Grule JSON

[![GRL_JSON_cn](https://github.com/yammadev/flag-icons/blob/master/png/CN.png?raw=true)](../cn/GRL_JSON_cn.md)
[![GRL_JSON_de](https://github.com/yammadev/flag-icons/blob/master/png/DE.png?raw=true)](../de/GRL_JSON_de.md)
[![GRL_JSON_en](https://github.com/yammadev/flag-icons/blob/master/png/GB.png?raw=true)](../en/GRL_JSON_en.md)
[![GRL_JSON_id](https://github.com/yammadev/flag-icons/blob/master/png/ID.png?raw=true)](../id/GRL_JSON_id.md)
[![GRL_JSON_pl](https://github.com/yammadev/flag-icons/blob/master/png/PL.png?raw=true)](../pl/GRL_JSON_pl.md)

[About](About_pl.md) | [Tutorial](Tutorial_pl.md) | [Rule Engine](RuleEngine_pl.md) | [GRL](GRL_pl.md) | [GRL JSON](GRL_JSON_pl.md) | [RETE Algorithm](RETE_pl.md) | [Functions](Function_pl.md) | [FAQ](FAQ_pl.md) | [Benchmark](Benchmarking_pl.md)

---

Reguły Grule mogą być reprezentowane w formacie JSON i tłumaczone przez silnik reguł na standardową składnię Grule. Format JSON ma oferować wysoki poziom elastyczności, aby dostosować się do potrzeb użytkownika.

Podstawowa struktura reguły JSON jest następująca:
```json
{
    "name": "SpeedUp",
    "desc": "When testcar is speeding up we keep increase the speed.",
    "salience": 10,
    "when": ...,
    "then": [
        ...
    ]
}
```

## Elementy

| Nazwa      | Opis                                                                                                                 |
| ---------- | -------------------------------------------------------------------------------------------------------------------- |
| `name`     | Nazwa reguły. **Wymagane**.                                                                                          |
| `desc`     | Opis reguły. **Opcjonalne**, domyślnie jest to `""`.                                                                 |
| `salience` | Wartość ważności dla reguły. **Opcjonalne**, domyślnie `0`.                                                          |
| `when`     | Warunek dla reguły. To pole może być albo zwykłym łańcuchem znaków, albo obiektem warunku (opisanym poniżej)         |
| `then`     | Tablica akcji dla danej reguły. Każdy element może być zwykłym łańcuchem znaków lub obiektem akcji (opisanym poniżej)|

## Warunek Obiekt

W celu zapewnienia dużej elastyczności, warunek `when` reguły może być rozbity na poszczególne komponenty. Jest to szczególnie użyteczne przy tworzeniu większych reguł i wspieraniu aplikacji GUI służących do edycji i analizy reguł.

Obiekty warunków są przetwarzane rekursywnie, co oznacza, że mogą być dowolnie zagnieżdżane w celu obsługi nawet najbardziej złożonych reguł. Za każdym razem, gdy parser oczekuje obiektu warunku, użytkownik może zamiast niego podać stały łańcuch znaków lub wartość liczbową, która zostanie zinterpretowana przez parser jako surowe dane wejściowe, które zostaną powtórzone w regule wyjściowej.

Każdy obiekt warunku ma następujący format:

```json
{"operator":[x, y, ...]}
```

gdzie `operator` jest jednym z operatorów opisanych poniżej, a `x` i `y` są dwoma lub więcej obiektami warunku lub stałymi.

### Operators

| Operator  | Opis              |
| --------- | ----------------- |
| `"and"`   | GRL && operator   |
| `"or"`    | GRL \|\| operator |
| `"eq"`    | GRL == operator   |
| `"not"`   | GRL != operator   |
| `"gt"`    | GRL > operator    |
| `"gte"`   | GRL >= operator   |
| `"lt"`    | GRL < operator    |
| `"lte"`   | GRL <= operator   |
| `"bor"`   | GRL \| operator   |
| `"band"`  | GRL & operator    |
| `"plus"`  | GRL + operator    |
| `"minus"` | GRL - operator    |
| `"div"`   | GRL / operator    |
| `"mul"`   | GRL * operator    |
| `"mod"`   | GRL % operator    |

### Operatory specjalne

Poniższe operatory zachowują się nieco inaczej niż operatory standardowe.

| Operator  | Opis                                                                                                                                                                                                                |
| --------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `"set"`   | GRL = operator. Ten operator ustawi wartość pierwszego argumentu na wartość wyjściową drugiego. Może być użyty tylko w sekcji `then` reguły.                                                                |
| `"call"`  | Wywołanie funkcji GRL. Operator wywołuje nazwę funkcji podaną w pierwszym operandzie. Jeżeli podano więcej niż jeden operand, to kolejne operandy są interpretowane jako argumenty przekazywane do wywołania funkcji. |
| `"obj"`   | Jednoznacznie identyfikuje obiekt GRL. W przeciwieństwie do innych operatorów, obiekt ten ma postać prostej pary klucz/wartość. Na przykład: `{"obj": "TestCar.Speed"}`                                                                 |
| `"const"` | Wyraźnie identyfikuje stałą GRL. Operator ten ma taką samą postać jak operator `obj`.                                                                                                                               |

### Obsługiwane stałe

Obsługiwane są następujące typy stałych:

| Typ      | Przykład                     |
| --------- | --------------------------- |
| `string`  | `{"const": "String Value"}` |
| `integer` | `{"const": 123}`            |
| `float`   | `{"const": 1.29738}`        |
| `bool`    | `{"const": true}`           |


## Akcje Then

Akcje `then` są tworzone w taki sam sposób jak obiekty warunków. Główną różnicą jest to, że elementem głównym każdego warunku `then` powinien być operator `set` lub `call`.

# Przykład

Aby zademonstrować możliwości reprezentacji JSON, należy przekonwertować poniższą przykładową regułę na składnię JSON:

```
rule SpeedUp "When testcar is speeding up we increase the speed." salience 10 {
    when
        TestCar.SpeedUp == true && TestCar.Speed < TestCar.MaxSpeed
    then
        TestCar.Speed = TestCar.Speed + TestCar.SpeedIncrement;
        DistanceRecord.TotalDistance = DistanceRecord.TotalDistance + TestCar.Speed;
        Log("Speed increased");
}
```

## Podstawowa reprezentacja

Najbardziej podstawowa reprezentacja tej reguły w JSON jest następująca:

```json
{
    "name": "SpeedUp",
    "desc": "When testcar is speeding up we increase the speed.",
    "salience": 10,
    "when": "TestCar.SpeedUp == true && TestCar.Speed < TestCar.MaxSpeed",
    "then": [
        "TestCar.Speed = TestCar.Speed + TestCar.SpeedIncrement",
        "DistanceRecord.TotalDistance = DistanceRecord.TotalDistance + TestCar.Speed",
        "Log(\"Speed increased\")"
    ]
}
```

W tym przykładzie warunki `when` i `then` są przedstawione jako surowe obiekty wejściowe. Daje to największy poziom kontroli nad regułą wyjściową. W większości przypadków translator wyprowadzi regułę, która jest dokładnym dopasowaniem do oryginalnej reprezentacji.  Jednakże w niektórych przypadkach tłumacz może wstawić nawiasy wokół wyrażeń, w których nie jest to wymagane, w zależności od obecności operatora. Nie powinno to mieć wpływu na logiczne znaczenie reguły.

## Rozszerzona reprezentacja

Powyższa reguła może być również przedstawiona w bardziej jednoznacznym formacie poprzez rozbicie warunków `when` i `then` na pełną reprezentację obiektową:

```json
{
    "name": "SpeedUp",
    "desc": "When testcar is speeding up we increase the speed.",
    "salience": 10,
    "when": {
       "and": [
           {"eq": ["TestCar.SpeedUp", true]},
           {"lt": ["TestCar.Speed", "TestCar.MaxSpeed"]}
       ]
    },
    "then": [
        {"set": ["TestCar.Speed", {"plus": ["TestCar.Speed", "TestCar.SpeedIncrement"]}]},
        {"set": ["DistanceRecord.TotalDistance", {"plus": ["DistanceRecord.TotalDistance", "TestCar.Speed"]}]},
        {"call": ["Log", {"const": "Speed increased"}]}
    ]
}
```

Tłumacz zinterpretuje powyższą regułę i wygeneruje taki sam wynik, jak w przypadku reguły przykładowej. Reguła ta jest znacznie bardziej rozbudowana i może być łatwo przetworzona i sformatowana do celów wyświetlania lub analizy.

## Niejawna reprezentacja i ucieczka od łańcuchów

Mimo że powyższa reguła jest bardziej rozbudowana, nadal reprezentuje obiekty i stałe boolean implicite. Nie jest konieczne owijanie każdego obiektu i stałej wewnątrz operatora `const` lub `obj`, ponieważ translator domyślnie zinterpretuje typ danych wejściowych. Jednakże, w niektórych przypadkach może być korzystne użycie operatorów `const` lub `obj`, z których najbardziej znaczącym jest wymuszenie reguł ucieczki na stałych łańcuchowych, które nie byłyby zastosowane do implikowanego łańcucha.

Przykład zachowania ucieczki z łańcuchów, pokazany w powyższej regule, występuje w akcji `call`. Jako argument funkcji przekazywany jest łańcuch znaków, więc aby przekazać argument jako obiekt surowy, użytkownik musi zawinąć stałą w cudzysłów i uciec od niej ręcznie:

```
{"call": ["Log", "\"Speed increased\""]}
```

Chociaż daje to poprawne wyniki, to jednak w pewnym sensie zaciemnia regułę, ponieważ nie jest jasne, jaki poziom ucieczki jest stosowany do danych wyjściowych.

Jeśli ucieczka z łańcucha zostanie wykonana niepoprawnie, tłumacz wygeneruje niepoprawne dane wyjściowe reguły. Aby zapobiec tego typu błędom, lepiej jest zawinąć stałą w operator `const`:

```
{"call": ["Log", {"const": "Speed increased"}]}
```

W takim przypadku translator dokona odpowiedniego escape'owania ciągu stałego i poprawnie wyprowadzi stałą do reguły wyjściowej.

## Przedstawienie słownikowe

Najbardziej dosłowna możliwa wersja przykładowej reguły w składni JSON wygląda następująco. (Zwróć uwagę, że dodatkowe operatory `obj` i `const` są tutaj zupełnie niepotrzebne, ale mogą być przydatne dla silników renderujących lub narzędzi przeznaczonych do edycji lub analizy reguł w formie wizualnej).

```json
{
    "name": "SpeedUp",
    "desc": "When testcar is speeding up we increase the speed.",
    "salience": 10,
    "when": {
       "and": [
           {"eq": [{"obj": "TestCar.SpeedUp"}, {"const": true}]},
           {"lt": [{"obj": "TestCar.Speed"}, {"obj": "TestCar.MaxSpeed"}]}
       ]
    },
    "then": [
        {"set": [{"obj": "TestCar.Speed"}, {"plus": [{"obj": "TestCar.Speed"}, {"obj": "TestCar.SpeedIncrement"}]}]},
        {"set": [{"obj": "DistanceRecord.TotalDistance"}, {"plus": [{"obj": "DistanceRecord.TotalDistance"}, {"obj": "TestCar.Speed"}]}]},
        {"call": ["Log", {"const": "Speed increased"}]}
    ]
}
```

# Wczytywanie reguł JSON

Reguły JSON mogą być ładowane z bazowego dostawcy zasobów lub dostawcy ResourceBundle za pomocą funkcji odpowiednio `NewJSONResourceFromResource` i `NewJSONResourceBundleFromBundle`. Kiedy wywołana jest funkcja `Load()`, ładowany jest bazowy zasób, a reguły są tłumaczone ze składni JSON na standardową składnię GRL. Dzięki temu funkcje translacji mogą być bardzo łatwo zintegrowane z istniejącym kodem.

```go
f, err := os.Open("rules.json")
if err != nil {
    panic(err)
}
underlying := pkg.NewReaderResource(f)
resource := pkg.NewJSONResourceFromResource(underlying)
...
```

Możliwe jest również parsowanie tablicy bajtów zawierającej reguły JSON bezpośrednio do zbioru reguł składni GRL przez wywołanie funkcji `ParseJSONRuleset`.

```go
jsonData, err := io.ReadFile("rules.json")
if err != nil {
    panic(err)
}
ruleset, err := pkg.ParseJSONRuleset(jsonData)
if err != nil {
    panic(err)
}
fmt.Println("Parsed ruleset: ")
fmt.Println(ruleset)
```
