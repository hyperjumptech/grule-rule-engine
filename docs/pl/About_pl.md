[![Gopher Holds The Rules](https://github.com/hyperjumptech/grule-rule-engine/blob/master/gopher-grule.png?raw=true)](https://github.com/hyperjumptech/grule-rule-engine/blob/master/gopher-grule.png?raw=true)


__"Gopher Trzyma Się Zasad"__

[![About_cn](https://github.com/yammadev/flag-icons/blob/master/png/CN.png?raw=true)](../cn/About_cn.md)
[![About_de](https://github.com/yammadev/flag-icons/blob/master/png/DE.png?raw=true)](../de/About_de.md)
[![About_en](https://github.com/yammadev/flag-icons/blob/master/png/GB.png?raw=true)](../en/About_en.md)
[![About_id](https://github.com/yammadev/flag-icons/blob/master/png/ID.png?raw=true)](../id/About_id.md)
[![About_pl](https://github.com/yammadev/flag-icons/blob/master/png/PL.png?raw=true)](../pl/About_pl.md)

[About](About_pl.md) | [Tutorial](Tutorial_pl.md) | [Rule Engine](RuleEngine_pl.md) | [GRL](GRL_pl.md) | [GRL JSON](GRL_JSON_pl.md) | [RETE Algorithm](RETE_pl.md) | [Functions](Function_pl.md) | [FAQ](FAQ_pl.md) | [Benchmark](Benchmarking_pl.md)

# Grule

```go
import "github.com/hyperjumptech/grule-rule-engine"
```

## Silnik reguł dla Go

**Grule** to biblioteka silnika reguł dla języka programowania Golang. Zainspirowana uznanym JBOSS Drools, ale wykonana w dużo prostszy sposób.

Podobnie jak **Droole**, **Grule** ma swoje własne *DSL*, które można porównać w następujący sposób.

Drools's DRL jest jak :

```go
rule "SpeedUp"
    salience 10
    when
        $TestCar : TestCarClass( speedUp == true && speed < maxSpeed )
        $DistanceRecord : DistanceRecordClass()
    then
        $TestCar.setSpeed($TestCar.Speed + $TestCar.SpeedIncrement);
        update($TestCar);
        $DistanceRecord.setTotalDistance($DistanceRecord.getTotalDistance() + $TestCar.Speed)
        update($DistanceRecord)
end
```

A Grule's GRL jest jak :

```go
rule SpeedUp "When testcar is speeding up we increase the speed." salience 10  {
    when
        TestCar.SpeedUp == true && TestCar.Speed < TestCar.MaxSpeed
    then
        TestCar.Speed = TestCar.Speed + TestCar.SpeedIncrement;
        DistanceRecord.TotalDistance = DistanceRecord.TotalDistance + TestCar.Speed;
}
```

# Co to jest RuleEngine

Nie ma lepszego wyjaśnienia niż artykuł autorstwa Martina Fowlera. Artykuł można przeczytać tutaj ([RulesEngine by Martin Fowler](https://martinfowler.com/bliki/RulesEngine.html)).

Zaczerpnięte z witryny **TutorialsPoint** (z niewielkimi modyfikacjami),

Silnik regułowy **Grule** to produkcyjny system regułowy, który wykorzystuje podejście oparte na regułach do implementacji systemu ekspertowego. Systemy ekspertowe to systemy oparte na wiedzy, które wykorzystują reprezentacje wiedzy do przetwarzania zdobytej wiedzy w bazę wiedzy, która może być wykorzystywana do wnioskowania.

System reguł produkcyjnych jest systemem Turinga kompletnym, w którym nacisk kładzie się na reprezentację wiedzy w celu wyrażenia logiki propozycjonalnej i pierwszego rzędu w sposób zwięzły, jednoznaczny i deklaratywny.

Mózgiem systemu reguł produkcyjnych jest *Interference Engine* [silnik wnioskowania], który można rozbudować do dużej liczby reguł i faktów. Silnik wnioskowania porównuje fakty i dane z Regułami Produkcji - zwanymi również **Produkcjami** lub po prostu **Regułami** - w celu wyciągnięcia wniosków, które skutkują działaniami.

Reguła produkcyjna jest dwuczęściową strukturą, która wykorzystuje logikę pierwszego rzędu do wnioskowania na temat reprezentacji wiedzy. Silnik reguł biznesowych to system oprogramowania, który wykonuje jedną lub więcej reguł biznesowych w środowisku produkcyjnym w trybie runtime.

Silnik reguł pozwala zdefiniować **"Co zrobić"**, a nie **"Jak to zrobić"**.

## Co to jest reguła

*(również zaczerpnięte z TutorialsPoint)*

Reguły to fragmenty wiedzy często wyrażane słowami: "Gdy wystąpią pewne warunki, wykonaj pewne zadania".

```go
When
   <Condition is true>
Then
   <Take desired Action>
```

Najważniejszą częścią Reguły jest jej część `when`. Jeśli część **when** jest spełniona, uruchamiana jest część **then**.

```go
rule  <rule_name> <rule_description>
   <attribute> <value> {
   when
      <conditions>

   then
      <actions>
}
```

## Zalety silnika reguł

### Programowanie deklaratywne

Reguły ułatwiają wyrażanie rozwiązań trudnych problemów, a także ich weryfikację. W przeciwieństwie do kodów, reguły są napisane mniej skomplikowanym językiem, dzięki czemu analitycy biznesowi mogą łatwo odczytać i zweryfikować zestaw reguł.

### Rozdzielenie logiki i danych

Dane są przechowywane w obiektach domeny, a logika biznesowa w regułach. W zależności od rodzaju projektu taki podział może być bardzo korzystny.

### Centralizacja wiedzy

Używając reguł, tworzy się repozytorium wiedzy (bazę wiedzy), która jest wykonywalna. Jest to pojedynczy punkt prawdy dla polityki biznesowej. Najlepiej, jeśli Reguły są tak czytelne, że mogą służyć również jako dokumentacja.

### Zdolność do zmian

Ponieważ reguły biznesowe są w rzeczywistości traktowane jak dane. Dostosowanie reguł do dynamicznej natury biznesu staje się banalnie proste. Nie ma potrzeby ponownego budowania kodu, wdrażania go tak, jak w przypadku normalnego tworzenia oprogramowania, wystarczy tylko opracować zestawy reguł i zastosować je w repozytorium wiedzy.
