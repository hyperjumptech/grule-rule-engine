# Algorytm RETE w Grule

[![RETE_cn](https://github.com/yammadev/flag-icons/blob/master/png/CN.png?raw=true)](../cn/RETE_cn.md)
[![RETE_de](https://github.com/yammadev/flag-icons/blob/master/png/DE.png?raw=true)](../de/RETE_de.md)
[![RETE_en](https://github.com/yammadev/flag-icons/blob/master/png/GB.png?raw=true)](../en/RETE_en.md)
[![RETE_id](https://github.com/yammadev/flag-icons/blob/master/png/ID.png?raw=true)](../id/RETE_id.md)
[![RETE_pl](https://github.com/yammadev/flag-icons/blob/master/png/PL.png?raw=true)](../id/RETE_pl.md)

[About](About_pl.md) | [Tutorial](Tutorial_pl.md) | [Rule Engine](RuleEngine_pl.md) | [GRL](GRL_pl.md) | [GRL JSON](GRL_JSON_pl.md) | [RETE Algorithm](RETE_pl.md) | [Functions](Function_pl.md) | [FAQ](FAQ_pl.md) | [Benchmark](Benchmarking_pl.md)

---

Z Wikipedii : Algorytm Rete (/ˈriːtiː/ REE-tee, /ˈreɪtiː/ RAY-tee, rzadko /ˈriːt/ REET, /rɛˈteɪ/ reh-TAY) to algorytm dopasowywania wzorców służący do implementacji systemów opartych na regułach. Algorytm ten został opracowany w celu efektywnego stosowania wielu reguł lub wzorców do wielu obiektów lub faktów w bazie wiedzy. Służy on do określania, które z reguł systemu powinny zostać uruchomione na podstawie zbioru danych, czyli faktów.

Pewna forma algorytmu RETE została zaimplementowana w `grule-rule-engine` począwszy od wersji `1.1.0`.
Zastępuje on podejście __Naive__ przy ocenie reguł, które mają być dodane do `ConflictSet`.

Elementy `ExpressionAtom` w GRL są kompilowane i nie będą powielane w pamięci roboczej silnika.
Zwiększa to znacznie wydajność silnika, gdy mamy zdefiniowanych wiele reguł z wieloma zduplikowanymi wyrażeniami lub wieloma ciężkimi wywołaniami funkcji/metod.

Implementacja RETE w Grule nie posiada selektora `Klasa`, ponieważ jedno wyrażenie może dotyczyć wielu klas. Na przykład, wyrażenie takie jak:

```.go
when
    ClassA.attr == ClassB.attr + ClassC.AFunc()
then
    ...
```

Powyższe wyrażenie obejmuje porównywanie atrybutów i wyników wywołań funkcji oraz operacje matematyczne z trzech różnych klas. Utrudnia to stosowane przez RETE rozdzielanie tokenów wyrażeń na klasy.

O algorytmie RETE można przeczytać tutaj:

* https://en.wikipedia.org/wiki/Rete_algorithm
* https://www.drdobbs.com/architecture-and-design/the-rete-matching-algorithm/184405218
* https://www.sparklinglogic.com/rete-algorithm-demystified-part-2/ 

### Dlaczego algorytm Rete jest potrzebny

Załóżmy, że mamy pewien fakt.

```go
type Fact struct {
    StringValue string
}

func (f *Fact) VeryHeavyAndLongFunction() bool {
    ...
}
```

Następnie dodajemy ten fakt do kontekstu danych:

```go
f := &Fact{}
dctx := context.NewDataContext()
err := dctx.Add("Fact", f)
```

Mamy też GRL jak:

```go
rule ... {
    when
        Fact.VeryHeavyAndLongFunction() && Fact.StringValue == "Fish"
    then
        ...
}

rule ... {
    when
        Fact.VeryHeavyAndLongFunction() && Fact.StringValue == "Bird"
    then
        ...
}

rule ... {
    when
        Fact.VeryHeavyAndLongFunction() && Fact.StringValue == "Mammal"
    then
        ...
}

// and many similar rules

rule ... {
    when
        Fact.VeryHeavyAndLongFunction() && Fact.StringValue == "Insect"
    then
        ...
}
```

Wykonanie powyższego GRL może "zabić" silnik, ponieważ, gdy będzie on próbował wybrać, które reguły wykonać, silnik będzie wywoływał funkcję `Fact.VeryHeavyAndLongFunction` w zakresie `when` każdej reguły.

Tak więc, zamiast wykonywać funkcję `Fact.VeryHeavyAndLongFunction` podczas ewaluacji każdej regułę, algorytm Rete oblicza je tylko raz (gdy wywołanie funkcji jest napotkane po raz pierwszy), a następnie zapamiętuje wynik dla pozostałych reguł. (**Uwaga**, że oznacza to, że wywołanie funkcji *musi być referencyjnie przezroczyste* - tzn. nie może mieć żadnych efektów ubocznych).

Podobnie jest z `Fact.StringValue`. Algorytm Rete wczyta wartość z instancji obiektu i i zapamięta ją, dopóki nie zostanie ona zmieniona w zakresie `then`, np:

```go
rule ... {
    when
        ...
    then
        Fact.StringValue = "something else";
}
```

### Co znajduje się w pamięci roboczej Grule'a

Grule będzie próbował zapamiętać wszystkie elementy `Expression` zdefiniowane w zakresie `when` reguły we wszystkich regułach w Bazie Wiedzy.

Po pierwsze, będzie się starał jak najlepiej upewnić się, że żaden z węzłów AST (Abstract Syntax Tree) nie jest zduplikowany.

Po drugie, każdy z tych węzłów AST może być analizowany tylko raz, aż do momentu, gdy odpowiednia `zmienna` zostanie zmieniona. Na przykład:

```Shell
    when
        Fact.A == Fact.B + Fact.Func(Fact.C) - 20
```

Ten warunek zostanie podzielony na następujące `wyrażenia`.

```Shell
Expression "Fact.A" --> A variable
Expression "Fact.B" --> A variable
Expression "Fact.C" --> A variable
Expression "Fact.Func(Fact.C)"" --> A function containing argument Fact.C
Expression "20" --> A constant
Expression "Fact.B + Fact.Func(Fact.C)" --> A math operation contains 2 variable; Fact.B and Fact.C
Expression "(Fact.B + Fact.Func(Fact.C))" - 20 -- A math operation also contains 2 variable.
```

Wynikowe wartości dla każdego z powyższych `wyrażeń` zostaną zapamiętane (memoized) przy ich pierwszym wywołaniu, tak że kolejne odwołania do nich będą unikały ponownego ich wywoływania, zwracając natychmiast zapamiętaną wartość.

Jeśli jedna z tych wartości zostanie zmieniona wewnątrz zakresu `then` reguły, na przykład...

```Shell
    then
        Fact.B = Fact.A * 20
```

... to wszystkie Wyrażenia zawierające `Fact.B` zostaną usunięte z Pamięci roboczej:

```Shell
Expression "Fact.B"
Expression "Fact.B + Fact.Func(Fact.C)" --> A math operation contains 2 variable; Fact.B and Fact.C
Expression "(Fact.B + Fact.Func(Fact.C))" - 20 -- A math operation also contains 2 variable. 
```

Te `Expression` zostaną usunięte z pamięci roboczej, aby można je było ponownie ocenić w następnym cyklu.

### Znany problem z RETE dla funkcji i metod

Podczas gdy Grule będzie próbował zapamiętać każdą zmienną, którą ocenia w zakresie `when` i `then`, to jeśli zmienisz wartość zmiennej spoza silnika reguł, na przykład zmienionej w wywołaniu funkcji, Grule nie będzie w stanie zobaczyć tej zmiany. W rezultacie, Grule może błędnie użyć starej (zapamiętanej) wartości zmiennej, ponieważ nie wie, że wartość uległa zmianie.  Powinieneś starać się, aby Twoje funkcje były **referencyjnie przezroczyste**, aby nigdy nie mieć do czynienia z tym problemem.

Rozważmy następujący fakt:

```go
type Fact struct {
    StringValue string
}

func (f *Fact) SetStringValue(newValue string) {
    f.StringValue = newValue
}
```

Następnie tworzy się instancję faktu i dodaje go do kontekstu danych:

```go
f := &Fact{
    StringValue: "One",
}
dctx := context.NewDataContext()
err := dctx.Add("Fact", f)
```

W GRL należy wykonać następującą czynność

```go
rule one "One" {
    when
        Fact.StringValue == "One"
        // Here grule remembers that Fact.StringValue value is "One"
    then
        Fact.SetStringValue("Two");
        // Here grule does not know that Fact.StringValue has changed inside the function.
        // What grule know is Fact.StringValue is still "One".
}

rule two "Two" {
    when
        Fact.StringValue == "Two"
        // Because of that, this will never evaluated true.
    then
        Fact.SetStringValue("Three");
}
```

W ten sposób silnik zakończy pracę bez błędu, ale oczekiwany wynik, gdzie `Fact.StringValue` powinno być `Two`. nie zostanie osiągnięty.

Aby temu zaradzić, należy powiedzieć grule, czy zmienna się zmieniła, używając funkcji `Changed`.

```go
rule one "One" {
    when 
        Fact.StringValue == "One"
        // here grule remember that Fact.StringValue value is "One"
    then
        Fact.SetStringValue("Two");
        // here grule does not know if Fact.StringValue has changed inside the function.
        // What grule know is Fact.StringValue is still "One"

        // We should tell Grule that the variable changed within the Fact
        Changed("Fact.StringValue")
}
```
