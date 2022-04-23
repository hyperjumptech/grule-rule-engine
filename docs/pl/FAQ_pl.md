# FAQ

[![FAQ_cn](https://github.com/yammadev/flag-icons/blob/master/png/CN.png?raw=true)](../cn/FAQ_cn.md)
[![FAQ_de](https://github.com/yammadev/flag-icons/blob/master/png/DE.png?raw=true)](../de/FAQ_de.md)
[![FAQ_en](https://github.com/yammadev/flag-icons/blob/master/png/GB.png?raw=true)](../en/FAQ_en.md)
[![FAQ_id](https://github.com/yammadev/flag-icons/blob/master/png/ID.png?raw=true)](../id/FAQ_id.md)
[![FAQ_pl](https://github.com/yammadev/flag-icons/blob/master/png/PL.png?raw=true)](../pl/FAQ_pl.md)

[About](About_pl.md) | [Tutorial](Tutorial_pl.md) | [Rule Engine](RuleEngine_pl.md) | [GRL](GRL_pl.md) | [GRL JSON](GRL_JSON_pl.md) | [RETE Algorithm](RETE_pl.md) | [Functions](Function_pl.md) | [FAQ](FAQ_pl.md) | [Benchmark](Benchmarking_pl.md)

---

## 1. Grule spanikował na maksymalnym cyklu

**Pytanie**: Podczas wykonywania silnika Grule otrzymałem następujący komunikat o panice.

```Shell
panic: GruleEngine successfully selected rule candidate for execution after 5000 cycles, this could possibly caused by rule entry(s) that keep added into execution pool but when executed it does not change any data in context. Please evaluate your rule entries "When" and "Then" scope. You can adjust the maximum cycle using GruleEngine.MaxCycle variable.
```

**Odpowiedź**: Ten błąd wskazuje na potencjalny problem z regułami, które zleciłeś silnikowi do oceny. Grule kontynuuje wykonywanie sieci RETE w pamięci roboczej do momentu, gdy nie pozostaną żadne akcje do wykonania w zbiorze konfliktów, co będziemy nazywać naturalnym stanem końcowym.  Jeśli Twój zestaw reguł nigdy nie pozwoli sieci osiągnąć tego stanu końcowego, to będzie ona działać wiecznie.  Domyślną konfiguracją dla `GruleEngine.MaxCycle` jest `5000`, co jest używane do ochrony przed nieskończonym cyklem przebiegów w niekończącym się zestawie reguł.

Możesz zwiększyć tę wartość, jeśli uważasz, że twój system reguł potrzebuje więcej cykli, by się zakończyć, ale jeśli nie sądzisz, że tak jest, to prawdopodobnie masz niekończący się zestaw reguł.

Rozważmy ten fakt:

```go
type Fact struct {
   Payment int
   Cashback int
}
```

Zdefiniowane są następujące zasady:

```Shell
rule GiveCashback "Give cashback if payment is above 100" {
    When 
         F.Payment > 100
    Then
         F.Cashback = 10;
}

rule LogCashback "Emit log if cashback is given" {
    When 
         F.Cashback > 5
    Then
         Log("Cashback given :" + F.Cashback);
}
```

Wykonanie tych reguł na następującej instancji faktu...

```go
&Fact {
     Payment: 500,
}
```

... nigdy się nie kończy.

```
Cycle 1: Execute "GiveCashback" .... because when F.Payment > 100 is a valid condition
Cycle 2: Execute "GiveCashback" .... because when F.Payment > 100 is a valid condition
Cycle 3: Execute "GiveCashback" .... because when F.Payment > 100 is a valid condition
...
Cycle 5000: Execute "GiveCashback" .... because when F.Payment > 100 is still a valid condition
panic
```

Grule wykonuje tę samą regułę wielokrotnie, ponieważ warunek **WHEN** nadal daje prawidłowy wynik.

Jednym ze sposobów rozwiązania tego problemu jest zmiana reguły "GiveCashback" na coś w rodzaju na przykład:

```Shell
rule GiveCashback "Give cashback if payment is above 100" {
    When 
         F.Payment > 100 &&
         F.Cashback == 0
    Then
         F.Cashback = 10;
}
```

Ta definicja reguły `GiveCashback` bierze pod uwagę zmieniający się stan.  Początkowo element `Cashback` będzie miał wartość `0`, ale ponieważ akcja modyfikuje ten stan, nie uda się jej dopasować w następnym cyklu i osiągnięty zostanie stan końcowy.

Powyższa metoda jest w pewnym sensie "naturalna", gdyż to warunki reguł regulują zakończenie. Jeśli jednak nie można zakończyć wykonywania akcji w ten naturalny sposób, można zmodyfikować stan silnika w akcji, stosując poniższe rozwiązanie:

```Shell
rule GiveCashback "Give cashback if payment is above 100" {
    When 
         F.Payment > 100
    Then
         F.Cashback = 10;
         Retract("GiveCashback");
}
```

Funkcja `Retract` usuwa regułę "GiveCashback" z bazy wiedzy dla następnego cyklu. Ponieważ nie jest ona już obecna, nie może być ponownie oceniona w następnym cyklu. Należy jednak pamiętać, że dzieje się to tylko w cyklu następującym bezpośrednio po wywołaniu funkcji `Retract`.  W kolejnym cyklu wywołanie to zostanie ponownie wprowadzone.

---

## 2. Zapisywanie wpisu reguły w bazie danych

**Pytanie**: Czy planowana jest integracja Grule z systemem przechowywania baz danych?

**Odpowiedź**: Nie. Chociaż dobrym pomysłem jest przechowywanie reguł w jakiejś bazie danych, Grule nie stworzy adaptera do bazy danych, który automatycznie przechowywałby i pobierał reguły.  Możesz w prosty sposób stworzyć taki adapter samodzielnie, używając wspólnych interfejsów z bazy wiedzy: *Reader*, *File*, *Byte Array*, *String* i *Git*. Łańcuchy mogą być łatwo wstawiane i wybierane z bazy danych podczas ładowania ich do bazy wiedzy Grule'a. 

Nie chcemy przywiązywać Grule'a do żadnej konkretnej implementacji bazy danych.

---

## 3. Maksymalna liczba reguł w jednej bazie wiedzy

**Pytanie**: Ile wpisów z regułami można umieścić w bazie wiedzy?

**Odpowiedź**: Możesz mieć tyle wpisów z regułami, ile chcesz, ale przynajmniej jeden powinien być minimalny.

---

## 4. Pobierz wszystkie reguły obowiązujące dla danego faktu

**Pytanie**: Jak mogę przetestować poprawność moich reguł dla podanych faktów?

**Odpowiedź**: Można użyć funkcji `engine.FetchMatchingRule`. Odwołaj się do tego
[Matching Rules Doc](MatchingRules_pl.md), aby uzyskać więcej informacji

---

## 5. Przypadek użycia silnika reguł

**Pytanie**: Czytałem o silniku reguł, ale jakie realne korzyści może on przynieść? Podaj nam kilka przypadków użycia.

**Odpowiedź**: Moim skromnym zdaniem następujące przypadki są lepiej rozwiązywane za pomocą silnika reguł.

1. System ekspertowy, który musi oceniać fakty, aby wyciągnąć jakieś wnioski w świecie rzeczywistym. Jeśli nie używamy silnika regułowego w stylu RETE, tworzylibyśmy kaskadowy zestaw instrukcji `if`/`else`, a permutacje kombinacji sposobów ich oceny szybko stałyby się niemożliwe do opanowania.  Silnik reguł oparty na tabelach mógłby wystarczyć, ale jest on bardziej odporny na zmiany i nie jest zbyt łatwy do zakodowania. System taki jak Grule pozwala na opisanie reguł i faktów w systemie, zwalniając użytkownika z konieczności opisania, w jaki sposób reguły są obliczane względem tych faktów, ukrywając przed nim większość tej złożoności.

2. System ratingowy. Na przykład system bankowy może tworzyć "ocenę" dla każdego klienta na podstawie zapisów transakcji klienta (faktów).  Ocena może się zmieniać w zależności od tego, jak często klient kontaktuje się z bankiem, ile pieniędzy wpłaca i wypuszcza, jak szybko płaci rachunki, ile odsetek zarabia dla siebie lub dla banku itd. Silnik reguł jest dostarczany przez programistę, a specyfikację faktów i reguł mogą dostarczyć eksperci merytoryczni zajmujący się klientami banku. Rozdzielenie tych różnych zespołów sprawia, że odpowiedzialność spoczywa tam, gdzie powinna.

3. Gry komputerowe. Status gracza, nagrody, kary, uszkodzenia, wyniki i systemy prawdopodobieństwa to wiele przykładów sytuacji, w których zasady odgrywają istotną rolę w niemal wszystkich grach komputerowych. Reguły te mogą wchodzić w bardzo złożone interakcje, często w sposób, którego twórca gry sobie nie wyobrażał. Zakodowanie tych dynamicznych sytuacji w języku skryptowym (np. LUA) może być dość skomplikowane, a silnik reguł może znacznie uprościć pracę.

4. Systemy klasyfikacji. Jest to właściwie uogólnienie systemu oceniania opisanego powyżej.  Używając silnika regułowego, możemy klasyfikować takie rzeczy, jak zdolność kredytowa, identyfikacja biochemiczna, ocena ryzyka dla produktów ubezpieczeniowych, potencjalne zagrożenia bezpieczeństwa i wiele innych.

5. System porad/sugestii. Reguła" to po prostu inny rodzaj danych, co czyni ją idealnym kandydatem do zdefiniowania przez inny program.  Programem tym może być inny system ekspercki lub sztuczna inteligencja.  Reguły mogą być manipulowane przez inny system w celu uwzględnienia nowych rodzajów faktów lub nowo odkrytych informacji o dziedzinie, którą zbiór reguł ma modelować.

Istnieje wiele innych przypadków użycia, w których można by skorzystać z silnika reguł. Powyższe przypadki stanowią jedynie niewielką część potencjalnych możliwości. 

Należy jednak podkreślić, że Rule-Engine nie jest oczywiście "srebrną kulą".  Istnieje wiele alternatywnych sposobów rozwiązywania problemów z "wiedzą" w oprogramowaniu i powinny one być stosowane wtedy, gdy są najbardziej odpowiednie. Nie stosuje się silnika reguł tam, gdzie wystarczyłaby na przykład prosta gałąź `if` / `else`.

Jest jeszcze jedna rzecz, na którą warto zwrócić uwagę: niektóre implementacje mechanizmów regułowych są niezwykle kosztowne, ale wiele firm czerpie z nich tak wiele korzyści, że koszty ich działania są łatwo rekompensowane przez tę wartość.  W przypadku nawet średnio złożonych przypadków użycia korzyści z silnego silnika reguł, który może rozdzielić zespoły i okiełznać złożoność biznesową, wydają się być dość oczywiste.

---

## 6. Rejestrowanie

**Pytanie**: Logi Grule'a są bardzo obszerne. Czy mogę wyłączyć logger Grule'a?

**Odpowiedź**: Tak. Można ograniczyć (lub całkowicie wyłączyć) logowanie Grule'a poprzez zwiększenie poziomu logowania.

```go
import (
    "github.com/hyperjumptech/grule-rule-engine/logger"
    "github.com/sirupsen/logrus"
)
...
...
logger.SetLogLevel(logrus.PanicLevel)
```

To ustawi log Grule'a na poziom `Panic`, gdzie będzie on wysyłał log tylko wtedy, gdy wpadnie w panikę.

Oczywiście, modyfikacja poziomu logów zmniejsza możliwości debugowania systemu, dlatego sugerujemy, aby wyższy poziom logów był ustawiany tylko w środowiskach produkcyjnych.
