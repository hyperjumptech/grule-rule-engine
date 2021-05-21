
[![Gopheer Holds The Rule](https://github.com/hyperjumptech/grule-rule-engine/blob/master/gopher-grule.png?raw=true)](https://github.com/hyperjumptech/grule-rule-engine/blob/master/gopher-grule.png?raw=true)

[![Build Status](https://travis-ci.org/hyperjumptech/grule-rule-engine.svg?branch=master)](https://travis-ci.org/hyperjumptech/grule-rule-engine)
[![Build Status](https://circleci.com/gh/hyperjumptech/grule-rule-engine.svg?style=svg)](https://circleci.com/gh/hyperjumptech/grule-rule-engine)
[![Go Report Card](https://goreportcard.com/badge/github.com/hyperjumptech/grule-rule-engine)](https://goreportcard.com/report/github.com/hyperjumptech/grule-rule-engine)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

__"Gopher Holds The Rules"__

# Grule-Rule-Engine

```go
import "github.com/hyperjumptech/grule-rule-engine"
```

## Rule Engine for Go

**Grule** is a Rule Engine library for the Golang programming language. Inspired by the acclaimed JBOSS Drools, done in a much simple manner.

Like **Drools**, **Grule** have its own *DSL* comparable as follows.

Drools's DRL be like :

```drool
rule "SpeedUp"
    salience 10
    when
        $TestCar : TestCarClass( speedUp == true && speed < maxSpeed )
        $DistanceRecord : DistanceRecordClass()
    then
        $TestCar.setSpeed($TestCar.Speed + $TestCar.SpeedIncrement);
        update($TestCar);
        $DistanceRecord.setTotalDistance($DistanceRecord.getTotalDistance() + $TestCar.Speed);
        update($DistanceRecord);
end
```

And Grule's GRL be like :

```go
rule SpeedUp "When testcar is speeding up we keep increase the speed." salience 10  {
    when
        TestCar.SpeedUp == true && TestCar.Speed < TestCar.MaxSpeed
    then
        TestCar.Speed = TestCar.Speed + TestCar.SpeedIncrement;
        DistanceRecord.TotalDistance = DistanceRecord.TotalDistance + TestCar.Speed;
}
```

# What is RuleEngine

There isn't a better explanation than the article authored by Martin Fowler. You can read the article here ([RulesEngine by Martin Fowler](https://martinfowler.com/bliki/RulesEngine.html)).

Taken from **TutorialsPoint** website (with slight modifications),

**Grule** Rule Engine is a Production Rule System that uses the rule-based approach to implement an Expert System. Expert Systems are knowledge-based systems that use knowledge representations to process acquired knowledge into a knowledge base that can be used for reasoning.

A Production Rule System is Turing complete with a focus on knowledge representation to express propositional and first-order logic in a concise, non-ambiguous and declarative manner.

The brain of a Production Rules System is an *Inference Engine* that can scale to a large number of rules and facts. The Inference Engine matches facts and data against Production Rules – also called **Productions** or just **Rules** – to infer conclusions which result in actions.

A Production Rule is a two-part structure that uses first-order logic for reasoning over knowledge representation. A business rule engine is a software system that executes one or more business rules in a runtime production environment.

A Rule Engine allows you to define **“What to Do”** and not **“How to do it.”**

## What is a Rule

*(also taken from TutorialsPoint)*

Rules are pieces of knowledge often expressed as, "When some conditions occur, then do some tasks."

```go
When
   <Condition is true>
Then
   <Take desired Action>
```

The most important part of a Rule is its when part. If the **when** part is satisfied, the **then** part is triggered.

```go
rule  <rule_name> <rule_description>
   <attribute> <value> {
   when
      <conditions>

   then
      <actions>
}
```

## Advantages of a Rule Engine

### Declarative Programming

Rules make it easy to express solutions to difficult problems and get the verifications as well. Unlike codes, Rules are written with less complex language; Business Analysts can easily read and verify a set of rules.

### Logic and Data Separation

The data resides in the Domain Objects and the business logic resides in the Rules. Depending upon the kind of project, this kind of separation can be very advantageous.

### Centralization of Knowledge

By using Rules, you create a repository of knowledge (a knowledge base) which is executable. It is a single point of truth for business policy. Ideally, Rules are so readable that they can also serve as documentation.

### Agility To Change

Since business rules are actually treated as data. Adjusting the rule according to business dynamic nature become trivial. No need to re-build codes, deploy as normal software development do, you only need to roll out sets of rule and apply them to knowledge repository.

## Use Cases

The following cases are better solved with a rule-engine.

1. An expert system that must evaluate facts to provide some sort of real-world
   conclusion. If not using a a RETE-style rule engine, one would code up a
   cascading set of `if`/`else` statements and the permutations of the
   combinations of how those might be evaluated would quickly become impossible
   to manage.  A table-based rule engine might suffice but it is still more
   brittle against change and is not terribly easy to code. A system like Grule
   allows you to describe the rules and facts of your system, releasing you from
   the need to describe how the rules are evaluated against those facts, hiding
   the bulk of that complexity from you.

2. A rating system. For example, a bank system may want to create a "score" for
   each customer based on the customer's transaction records (facts).  We could
   see their score change based on how often they interact with the bank, how
   much money they transfer in and out, how quickly they pay their bills, how
   much interest they accrue earn for themselves or for the bank, and so on. A
   rule engine is provided by the developer and the specification of the facts
   and rules can then be supplied by subject matter experts in the bank's
   customer business. Decoupling these different teams puts the responsbilities
   where they should be.

3. Computer games. Player status, rewards, penalties, damage, scores and
   probability systems are many different examples of where rule play a
   significant part of nearly all computer games. These rules can interact in
   very complex ways, often times in ways that the developer didn't imagine.
   Coding these dynamic situations in a scripting language (e.g. LUA) can get
   quite complex, and a rule engine can help simplify the work tremendously.

4. Classification systems. This is actually a generalization of the rating
   system described above.  Using a rule engine, we can classify things such as
   credit eligibility, bio chemical identification, risk assessment for
   insurance products, potential security threats, and many more.

5. Advice/Suggestion system. A "rule" is simply another kind of data, which
   makes it a prime candidate for definition by another program.  This program
   can be another expert system or artificial intelligence.  Rules can be
   manipulated by another system in order to deal with new types of facts or
   newly discovered information about the domain which the rule set is intending
   to model.

There are so many other use-cases that would benefit from the use of
Rule-Engine. The above cases represent only a small number of the potential.

However it is important to state that a Rule-Engine not a silver bullet, of
course.  Many alternatives exist to solve "knowledge" problems in software and
those should be employed when they are most appropriate. One would not employ a
rule engine where a simple `if` / `else` branch would suffice, for instance.

Theres's someting else to note: some rule engine implementations are extremely
expensive yet many businesses gain so much value from them that the cost of
running them is easily offset by that value.  For even moderately complex use
cases, the benefit of a strong rule engine that can decouple teams and tame
business complexity seems to be quite clear.


## Docs

Grule's Documentation now viewable in ViewDocs. [http://hyperjumptech.viewdocs.io](http://hyperjumptech.viewdocs.io/grule-rule-engine)

### Benchmark
`Loading rules into KnowledgeBase`:

* To load `100` rules into knowledgeBase it took `99342047 ns/op` (took the highest value) that is equal to `~99.342047ms` and (`49295906 B/op`) `~49.295906MB` memory per operation

* To load `1000` rules into knowledgeBase it took `933617752 ns/op` (took the highest value) that is equal to `~933.617752ms` and (`488126636 B/op`) `~488.126636` memory per operation

`Executing rules against a fact`:

* To execute a fact against 100 rules, Grule Engine took `~9697 ns/op` (took the highest value as base) that is hardly `~0.009697ms` and `3957 B/op` which is pretty fast.

* To execute a fact against 1000 rules, Grule Engine took `~568959 ns/op` (took the highest value as base) that is hardly `~0.568959ms` and `293710 B/op` which is also pretty fast.


You can read the [detail report here](docs/Benchmarking_en.md)

# Our Contributors


<table width="100%">
<tr><td align="center"><a href="https://github.com/newm4n"><img width="80px" height="80px" src="https://avatars3.githubusercontent.com/u/3471399?v=4"><br><br>newm4n</a><br><br></td>
<td align="center"><a href="https://github.com/jinagamvasubabu"><img width="80px" height="80px" src="https://avatars1.githubusercontent.com/u/8560620?v=4"><br><br>jinagamvasubabu</a><br><br></td>
<td align="center"><a href="https://github.com/niallnsec"><img width="80px" height="80px" src="https://avatars3.githubusercontent.com/u/21335031?v=4"><br><br>niallnsec</a><br><br></td>
<td align="center"><a href="https://github.com/inhuman"><img width="80px" height="80px" src="https://avatars0.githubusercontent.com/u/2518263?v=4"><br><br>inhuman</a><br><br></td>
<td align="center"><a href="https://github.com/ariya"><img width="80px" height="80px" src="https://avatars1.githubusercontent.com/u/7288?v=4"><br><br>ariya</a><br><br></td>
<td align="center"><a href="https://github.com/sapiderman"><img width="80px" height="80px" src="https://avatars1.githubusercontent.com/u/964106?v=4"><br><br>sapiderman</a><br><br></td>
</tr>
<tr>
<td align="center"><a href="https://github.com/jtr860830"><img width="80px" height="80px" src="https://avatars1.githubusercontent.com/u/13183797?v=4"><br><br>jtr860830</a><br><br></td>
<td align="center"><a href="https://github.com/trancee"><img width="80px" height="80px" src="https://avatars0.githubusercontent.com/u/1520623?v=4"><br><br>trancee</a><br><br></td>
<td align="center"><a href="https://github.com/liouxiao"><img width="80px" height="80px" src="https://avatars2.githubusercontent.com/u/3435699?v=4"><br><br>liouxiao</a><br><br></td>
<td align="center"><a href="https://github.com/Troush"><img width="80px" height="80px" src="https://avatars0.githubusercontent.com/u/1163074?v=4"><br><br>Troush</a><br><br></td>
<td align="center"><a href="https://github.com/shanhuhai5739"><img width="80px" height="80px" src="https://avatars3.githubusercontent.com/u/3794113?v=4"><br><br>shanhuhai5739</a><br><br></td>
<td align="center"><a href="https://github.com/derekwyatt"><img width="80px" height="80px" src="https://avatars3.githubusercontent.com/u/62324?v=4"><br><br>derekwyatt</a><br><br></td>
</tr>
<tr>
<td align="center"><a href="https://github.com/garychristianto"><img width="80px" height="80px" src="https://avatars1.githubusercontent.com/u/50298986?v=4"><br><br>garychristianto</a><br><br></td>
<td align="center"><a href="https://github.com/sourcesoft"><img width="80px" height="80px" src="https://avatars2.githubusercontent.com/u/608906?v=4"><br><br>sourcesoft</a><br><br></td>
<td align="center"><a href="https://github.com/sdowding-koho"><img width="80px" height="80px" src="https://avatars3.githubusercontent.com/u/62896133?v=4"><br><br>sdowding-koho</a><br><br></td>
<td align="center"><a href="https://github.com/yomashExpel"><img width="80px" height="80px" src="https://avatars3.githubusercontent.com/u/25300754?v=4"><br><br>yomashExpel</a><br><br></td>
<td align="center"><a href="https://github.com/avisdsouza"><img width="80px" height="80px" src="https://avatars1.githubusercontent.com/u/8979874?v=4"><br><br>avisdsouza</a><br><br></td>
<td align="center"><a href="https://github.com/zct"><img width="80px" height="80px" src="https://avatars3.githubusercontent.com/u/4023051?v=4"><br><br>zct</a><br><br></td>
</tr>
<tr>
<td align="center"><a href="https://github.com/enricoojf"><img width="80px" height="80px" src="https://avatars2.githubusercontent.com/u/17194541?v=4"><br><br>enricoojf</a><br><br></td>
<td align="center"><a href="https://github.com/vlean"><img width="80px" height="80px" src="https://avatars1.githubusercontent.com/u/7309530?v=4"><br><br>vlean</a><br><br></td>
</tr>
</table>





# Tasks and Help Wanted

Yes. We need contributors to make Grule even better and useful to the Open Source Community.

* Need to do more and more and more tests.
* Better code coverage test.
* Better commenting for go doc best practice.
* Improve function argument handling to be more fluid and intuitive.

If you really want to help us, simply `Fork` the project and apply for Pull Request.
Please read our [Contribution Manual](CONTRIBUTING.md) and [Code of Conduct](CODE_OF_CONDUCTS.md)
