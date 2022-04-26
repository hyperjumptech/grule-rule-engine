# JSON Fact

[![JSON_Fact_cn](https://github.com/yammadev/flag-icons/blob/master/png/CN.png?raw=true)](../cn/JSON_Fact_cn.md)
[![JSON_Fact_de](https://github.com/yammadev/flag-icons/blob/master/png/DE.png?raw=true)](../de/JSON_Fact_de.md)
[![JSON_Fact_en](https://github.com/yammadev/flag-icons/blob/master/png/GB.png?raw=true)](../en/JSON_Fact_en.md)
[![JSON_Fact_id](https://github.com/yammadev/flag-icons/blob/master/png/ID.png?raw=true)](../id/JSON_Fact_id.md)
[![JSON_Fact_pl](https://github.com/yammadev/flag-icons/blob/master/png/PL.png?raw=true)](../pl/JSON_Fact_pl.md)

[About](About_pl.md) | [Tutorial](Tutorial_pl.md) | [Rule Engine](RuleEngine_pl.md) | [GRL](GRL_pl.md) | [GRL JSON](GRL_JSON_pl.md) | [RETE Algorithm](RETE_pl.md) | [Functions](Function_pl.md) | [FAQ](FAQ_pl.md) | [Benchmark](Benchmarking_pl.md)

---

Używanie danych JSON do reprezentowania faktów w Grule jest dostępne od wersji 1.8.0. Umożliwia ono użytkownikowi wyrażanie faktów w formacie JSON, a następnie dodawanie tych faktów do `DataContext`, tak jak to się normalnie robi w kodzie. Wczytane fakty JSON są teraz "widoczne" dla skryptów Grule (GRL).

## Dodawanie JSON jako faktu

Zakładając, że masz JSON w następującej postaci:

```json
{
  "name" : "John Doe",
  "age" : 24,
  "gender" : "M",
  "height" : 74.8,
  "married" : false,
  "address" : {
    "street" : "9886 2nd St.",
    "city" : "Carpentersville",
    "state" : "Illinois",
    "postal" : 60110
  },
  "friends" : [ "Roth", "Jane", "Jake" ]
}
```

Umieszczasz swój plik JSON w tablicy bajtów.

```go
myJSON := []byte (...your JSON here...)
```

Wystarczy dodać zmienną JSON do `DataContext`.

```go
// create new instance of DataContext
dataContext := ast.NewDataContext()

// add your JSON Fact into data context using AddJSON() function.
err := dataContext.AddJSON("MyJSON", myJSON)
```

Tak, możesz dodać tyle _faktów_ ile chcesz do kontekstu i możesz mieszać fakty JSON (używając `AddJSON`) z normalnymi faktami Go (używając `Add`)

## Ocenianie (odczytywanie) wartości faktów JSON w GRL
 
W skrypcie GRL, fakty są zawsze widoczne poprzez ich etykiety, które podajesz podczas dodawania do `DataContext`. Na przykład, poniższy kod dodaje twój JSON i będzie on używał etykiety `MyJSON`.
 
 ```go
err := dataContext.AddJSON("MyJSON", myJSON)
```
 
Tak, możesz użyć dowolnej etykiety, pod warunkiem że jest to pojedyncze słowo.
 
### Przemieszczanie się po zmiennych członkowskich jak po normalnym obiekcie
 
Używając JSON pokazanego na początku, zakres GRL `when` może obliczyć twój
json w następujący sposób.
 
 ```text
when
    MyJSON.name == "John Doe"
``` 

lub 

```text
when
    MyJSON.address.city.StrContains("ville")
```

lub

```text
when
    MyJSON.age > 30 && MyJSON.height < 60
```

### Przemierzanie zmiennych członkowskich jak mapa

Możesz uzyskać dostęp do pól obiektu JSON używając `Map` jak selektora lub jak normalnego obiektu.

 ```text
when
    MyJSON["name"] == "John Doe"
``` 

lub 

```text
when
    MyJSON["address"].city.StrContains("ville")
```

lub

```text
when
    MyJSON.age > 30 && MyJSON["HEIGHT".ToLower()] < 60
```

### Przemierzanie zmiennej członkowskiej tablicy

Element tablicy JSON można przeglądać tak samo jak zwykłą tablicę

 ```text
when
    MyJSON.friends[3] == "Jake"
```

## Zapisywanie wartości do faktów JSON w GRL

Tak, możesz zapisywać nowe wartości do faktów JSON w zakresie `then` swoich reguł. Te zmienione wartości będą wtedy dostępne w następnym cyklu oceny reguł. ALE, są pewne zastrzeżenia (przeczytaj "Rzeczy, które powinieneś wiedzieć" poniżej).

### Zapisywanie zmiennej członkowskiej jak normalnego obiektu
 
Używając JSON pokazanego na początku, twój GRL `then` może zmodyfikować twój json **fact** w następujący sposób.
 
 ```text
then
    MyJSON.name = "Robert Woo";
``` 

lub 

```text
then
    MyJSON.address.city = "Corruscant";
```

lub

```text
then
    MyJSON.age = 30;
```

To dość proste. Ale są też pewne utrudnienia.

1. Możesz modyfikować nie tylko wartość zmiennej członkowskiej obiektu JSON, możesz również zmienić jej `typ`.
   Zakładając, że twoja reguła jest w stanie obsłużyć kolejny łańcuch ewaluacji dla nowego typu, możesz to zrobić, w przeciwnym razie **bardzo mocno odradzamy to**.
   
   Przykład:
   
   Zmodyfikowałeś `MyJSON.age` na string.
   
   ```text
    then
        MyJSON.age = "Thirty";
   ```
   
   Ta zmiana sprawi, że silnik będzie wpadał w panikę podczas sprawdzania reguł typu:
   
   ```text
    when
        myJSON.age > 25
   ```
   
2. Można przypisać wartość do nieistniejącej zmiennej członkowskiej.
 
   Przykład:
   
      ```text
       then
           MyJSON.category = "FAT";
      ```

    Gdzie element `category` nie istnieje w oryginalnym JSON.
    
### Zapisywanie zmiennej członkowskiej jak normalnej mapy
 
Używając JSON-a pokazanego na początku, twój GRL `then` może zmodyfikować twój json **fact** w następujący sposób.
 
 ```text
then
    MyJSON["name"] = "Robert Woo";
``` 

lub 

```text
then
    MyJSON["address"]["city"] = "Corruscant";
```

lub

```text
then
    MyJSON["age"] = 30;
```

Podobnie jak w przypadku stylu obiektu, obowiązują te same zwroty.

1. Możesz modyfikować nie tylko wartość zmiennej członkowskiej swojej mapy JSON, możesz również zmienić jej `typ`.
   Zakładając, że Twoja reguła jest w stanie obsłużyć kolejny łańcuch ewaluacji dla nowego typu, możesz to zrobić, w przeciwnym razie **bardzo mocno odradzamy to**.
   
   Przykład:
   
   Zmodyfikowałeś `MyJSON.age` na string.
   
   ```text
    then
        MyJSON["age"] = "Thirty";
   ```
   
   Ta zmiana sprawi, że silnik będzie wpadał w panikę podczas sprawdzania reguł typu:
   
   ```text
    when
        myJSON.age > 25
   ```
   
2. Można przypisać wartość do nieistniejącej zmiennej członkowskiej
 
   Przykład:
   
      ```text
       then
           MyJSON["category"] = "FAT";
      ```

    Gdy element `category` nie istnieje w oryginalnym JSON.

### Zapisywanie tablicy członków

Element tablicy można zastąpić, używając jego indeksu.

```text
then
   MyJSON.friends[3] == "Jake";
```

Podany indeks musi być poprawny. Grule wpadnie w panikę, jeśli indeks będzie poza granicami.
Podobnie jak w przypadku zwykłego JSON, można zastąpić wartość dowolnego elementu innym typem.
Zawsze można też sprawdzić długość tablicy.

```text
when
   MyJSON.friends.Length() > 4;
```

Można również dołączać do tablicy za pomocą funkcji `Append`.  Append może również dołączać zmienną listę wartości argumentów do tablicy, używając różnych typów. (W przypadku zmiany typu danej wartości obowiązują te same zastrzeżenia).

```text
then
   MyJSON.friends.Append("Rubby", "Anderson", "Smith", 12.3);
```

**Znany problem**

Nie ma wbudowanych funkcji ułatwiających użytkownikowi sprawdzanie zawartości tablicy, takich jak `Contains(value) bool`.

## Rzeczy, które powinieneś wiedzieć

1. Po dodaniu faktu JSON do `DataContext`, zmiana w łańcuchu JSON nie będzie odzwierciedlać faktów już znajdujących się w `DataContext`. Jest to również stosowane w odwrotnym kierunku, gdzie zmiany w faktach w `DataContext` nie zmienią łańcucha JSON.
2. Możesz modyfikować swój fakt JSON w zakresie `then`, ale w przeciwieństwie do normalnych faktów `Go`, te zmiany nie będą miały odzwierciedlenia w oryginalnym łańcuchu JSON. Jeśli chcesz, aby tak się stało, powinieneś sparsować swój JSON do `struct` przedtem, i dodać `struct` do `DataContext` normalnie. 
