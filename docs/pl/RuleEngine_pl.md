# Introduction to Rule Engine

[![RuleEngine_cn](https://github.com/yammadev/flag-icons/blob/master/png/CN.png?raw=true)](../cn/RuleEngine_cn.md)
[![RuleEngine_de](https://github.com/yammadev/flag-icons/blob/master/png/DE.png?raw=true)](../de/RuleEngine_de.md)
[![RuleEngine_en](https://github.com/yammadev/flag-icons/blob/master/png/GB.png?raw=true)](../en/RuleEngine_en.md)
[![RuleEngine_id](https://github.com/yammadev/flag-icons/blob/master/png/ID.png?raw=true)](../id/RuleEngine_id.md)
[![RuleEngine_pl](https://github.com/yammadev/flag-icons/blob/master/png/PL.png?raw=true)](../id/RuleEngine_pl.md)

[About](About_pl.md) | [Tutorial](Tutorial_pl.md) | [Rule Engine](RuleEngine_pl.md) | [GRL](GRL_pl.md) | [GRL JSON](GRL_JSON_pl.md) | [RETE Algorithm](RETE_pl.md) | [Functions](Function_pl.md) | [FAQ](FAQ_pl.md) | [Benchmark](Benchmarking_pl.md)

---

Silnik reguł, jak wyjaśnił Martin Fowler, jest alternatywą dla modelu obliczeniowego, w którym zamiast oceny wielu warunków wybierana jest odpowiednia akcja, jeśli spełnione są określone warunki. W najprostszym ujęciu każda *Rule* przedstawia instrukcję *if-then*.

Zbiór reguł jest wprowadzany do **KnowledgeBase**, a następnie *engine* wykorzystuje każdą regułę w bazie wiedzy do oceny pewnych **Facts**. Jeśli wymagania danej reguły zostaną spełnione, zostanie wykonana **action** określona przez wybraną regułę.

## Fakt

`fact` jest faktem, jakkolwiek głupio by to nie brzmiało, ale tak właśnie jest. Fakt, w kontekście silnika reguł, jest po prostu informacją, która może być analizowana. Fakty mogą pochodzić z dowolnego źródła, np. z bazy danych, wywołanego procesu, systemu sprzedaży, raportu itp.

Dużym ułatwieniem może być przyjrzenie się przykładowi takiego faktu. Załóżmy, że mamy taki fakt:

```Text
Purchase Transaction
    Item Name     : Computer Monitor
    Quantity      : 10
    Purchase Date : 12 Dec 2019
    Item Price    : 150 USD
    Total Price   : 1500 USD
    Tax           : ?
    Discount      : ?
    Final Price   : ?
```

**Fact** to w zasadzie każda informacja lub zebrane dane. 

Z tego przykładowego Faktu zakupu znamy wiele informacji: przedmiot zakupu, ilość, datę zakupu itd. Nie wiemy jednak, ile podatku należy przypisać do tego zakupu, jak dużego rabatu możemy udzielić oraz jaką cenę końcową powinien zapłacić kupujący.

## Reguła

Reguła to specyfikacja określająca, jak należy oceniać **Fact**. Jeśli fakt spełni warunki reguły, to zostanie wybrana akcja, która zostanie wykonana w ramach reguły. Zdarza się, że wybieranych jest wiele reguł, ponieważ wszystkie ich specyfikacje dotyczą jednego faktu, co prowadzi do konfliktu. Zbiór wszystkich Reguł będących w konflikcie nazywamy **Conflict Set**. Aby rozwiązać ten zbiór, określamy *strategy* (opisaną w dalszej części instrukcji).  

Wracając do naszego przykładu prostego systemu zakupów: aby obliczyć ostateczną cenę, należy ustalić pewne Reguły biznesowe, prawdopodobnie najpierw obliczając podatek, a następnie rabat. Jeśli zarówno podatek, jak i rabat są znane, możemy wyświetlić cenę.

Określmy kilka reguł (w pseudokodzie).

```text
Rule 1
   IF
   - the Item's Tax is not known AND
   - the Item's Name is "Computer CPU"
   THEN
   - Item's Tax is 10%

Rule 2
   IF
   - the Item's Tax is not known AND
   - the Item's Name is "Computer Monitor"
   THEN
   - Item's Tax is 7%

Rule 3
   IF
   - the Item's Discount is not known AND
   - the Item's Price After Tax is Less Than 1000 USD
   THEN
   - Item's Discount is 0%

Rule 4
   IF
   - the Item's Discount is not known AND
   - the Item's Price After Tax is Less Than 1500 USD AND
   - the Item's Price After Tax is Greater Than or Equal To 1000 USD
   THEN
   - Item's Discount is 3%

Rule 5
   IF
   - the Item's Discount is not known AND
   - the Item's Price After Tax is Less Than 2000 USD AND
   - the Item's Price After Tax is Greater Than or Equal To 1500 USD
   THEN
   - Item's Discount is 5%

Rule 6
   IF
   - the Item's Discount is not known AND
   - the Item's Price After Tax is More Than 2000 USD
   THEN
   - Item's Discount is 10%

Rule 7
   IF
   - the Item's Total Price is known AND
   - the Item's Discount is known AND
   - the Item's Tax is known AND
   - the Item's Final Price is not known
   THEN
   - Item's Final Price is calculate price from Total Price
     with given Tax and Discount
```

Jeśli przeanalizujesz powyższe Reguły, powinieneś łatwo zrozumieć koncepcję **Rules** dla silników reguł. Te zbiory reguł tworzą zestaw **Knowledge**. W tym przypadku tworzą one zestaw wiedzy o tym, jak **"how to calculate Item's final price"**.

## Cykl

Cykl oceny Reguły rozpoczyna się od oceny wymagań każdej z Reguł (**IFs**) w celu wybrania Reguł, które potencjalnie mogą zostać wykonane. Za każdym razem, gdy silnik znajduje spełnione wymaganie, zamiast wykonywać akcję spełnionej Reguły (**THEN**), dodaje tę Regułę do listy kandydatów na Regułę (zwanej Zbiorem Konfliktów).

Czy po obliczeniu wszystkich wymagań Reguł silnik wykonuje akcje wybranych Reguł?  
To zależy od zawartości zbioru konfliktów:

* Jeśli nie ma żadnej Reguły z pasującym warunkiem **IF**, wykonanie silnika może się natychmiast zakończyć.
* Jeśli w zbiorze konfliktów znajduje się tylko jedna Reguła, to akcja tej Reguły jest wykonywana przed zakończeniem.
* Jeśli w zbiorze konfliktów znajduje się wiele Reguł, silnik musi zastosować strategię w celu nadania priorytetu jednej z nich i wykonania jej akcji.

Jeśli jakaś akcja zostanie wykonana, cykl powtarza się tak długo, jak długo istnieje akcja, która wymaga wykonania. Gdy nie zostanie wykonana żadna akcja, oznacza to, że nie ma już więcej Reguł, które są spełnione przez dany fakt (nie ma już pasujących instrukcji **IF**), a cykl zatrzymuje się, pozwalając silnikowi reguł na zakończenie oceny.

Pseudokod dla tej strategii rozwiązywania konfliktów jest przedstawiony poniżej:

```text
Start Engine With a FACT Using a KNOWLEDGE
BEGIN
    For Every RULE in KNOWLEDGE
        Check if RULE's Requirement is Satisfied by FACT
            If RULE's Requirement is Satisfied
                Add RULE into CONFLICT SET
            End If
        End Check
    End For
    If CONFLICT SET is EMPTY
        Finished
        END
    If CONFLICT SET Has 1 RULE
        Execute the RULE's Action
        Clear CONFLICT SET
        Repeat Cycle from BEGIN
    If CONFLICT SET has Many RULEs
        Apply Conflict Resolution Strategy to Choose 1 RULE.
        Execute the Chosen RULE's Action
        Clear CONFLICT SET
        Repeat Cycle from BEGIN
END
```

Grule śledzi, ile cykli wykonuje podczas pojedynczej oceny zestawu reguł. 
Jeśli ocena i wykonanie reguły zostaną powtórzone zbyt wiele razy, ponad ilość określoną przy tworzeniu instancji silnika Grule, silnik zakończy pracę i zostanie zwrócony błąd.

## Strategia rozwiązywania zbioru konfliktów

Jak wyjaśniono powyżej, silnik Reguł ocenia wymagania wszystkich Reguł i dodaje je do listy sprzecznych Reguł zwanej **Conflict Set**. Jeśli na liście znajduje się tylko jedna Reguła, oznacza to, że nie ma żadnej Reguły (żadnych Reguł) będącej w konflikcie z tą jedną Regułą. Silnik natychmiast wykona działanie tej Reguły.

Jeśli w zestawie znajduje się wiele Reguł, mogą występować konflikty. Istnieje wiele strategii rozwiązywania konfliktów, które mogą być zaimplementowane dla tego typu konfliktów Reguł. Najprostszym sposobem jest określenie **salience** Reguły (znanej również jako **priority** lub **importance**). Możemy dodać do definicji Reguły wskaźnik **salience** lub ważności Reguły, tak jak w poniższym pseudokodzie:

```text
Rule 1 - Priority 1
   IF
   - the Item's Tax is not known AND
   - the Item's Name is "Computer CPU"
   THEN
   - Item's Tax is 10%

Rule 2 - Priority 10
   IF
   - the Item's Tax is not known AND
   - the Item's Name is "Computer Monitor"
   THEN
   - Item's Tax is 7%
```

Domyślnie, wszystkie Reguły mają przypisaną saliencję `0`.

Ponieważ wszystkie nieokreślone Reguły mają wartość 0, silnik może łatwo wybrać, która z nich ma zostać wykonana, jeśli w zbiorze konfliktów znajduje się wiele Reguł. Jeśli istnieje wiele reguł o identycznych priorytetach, silnik wybierze pierwszą znalezioną. Ponieważ typy map Go nie gwarantują zachowania kolejności danych wejściowych, nie można bezpiecznie zakładać, że kolejność wykonywania reguł będzie zgodna z kolejnością dodawania reguł do instancji wiedzy Grule. 

Salience dla Reguł Go może mieć wartość poniżej zera (ujemną), aby zapewnić Regule jeszcze niższy priorytet niż domyślny. Zapewnia to, że działanie reguły zostanie wykonane jako ostatnie, po tym jak wszystkie inne reguły zostaną ocenione.