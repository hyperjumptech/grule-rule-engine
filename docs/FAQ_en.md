# FAQ

[Tutorial](Tutorial_en.md) | [Rule Engine](RuleEngine_en.md) | [GRL](GRL_en.md) | [GRL JSON](GRL_JSON_en.md) | [RETE Algorithm](RETE_en.md) | [Functions](Function_en.md) | [FAQ](FAQ_en.md) | [Benchmark](Benchmarking_en.md)

---

## 1. Grule Panicked on Maximum Cycle

**Question** : I got the following panic message when Grule engine is executed.

```Shell
panic: GruleEngine successfully selected rule candidate for execution after 5000 cycles, this could possibly caused by rule entry(s) that keep added into execution pool but when executed it does not change any data in context. Please evaluate your rule entries "When" and "Then" scope. You can adjust the maximum cycle using GruleEngine.MaxCycle variable.
```

**Answer** : The rule engine is done evaluating rule entries for choosing which one to execute in the 5000th time (5000 it the maximum execution cycle). You can change specify any positive number but if you doubt, you can leave it to 5000. When the rule set were evaluated that many times more to this number (the max cycle) it will panic. To fix this issue, have to understand how rule engine works. The following simulation should help you understand the problem and know how to mitigate it.

Consider this fact.

```go
type Fact struct {
   Payment int
   Cashback int
}
```

The following rules defined.

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

Then you execute the rule with Fact instance of

```go
&Fact {
     Payment: 500,
}
```

Now when the engine executes....

Cycle 1 : Execute "GiveCashback" .... because when F.Payment > 100 is a valid condition<br>
Cycle 2 : Execute "GiveCashback" .... because when F.Payment > 100 is a valid condition<br>
Cycle 3 : Execute "GiveCashback" .... because when F.Payment > 100 is a valid condition<br>
Cycle 4 : Execute "GiveCashback" .... because when F.Payment > 100 is a valid condition<br>
Cycle n : Execute "GiveCashback" .... because when F.Payment > 100 is a valid condition<br>
Cycle 5000 : Execute "GiveCashback" .... because when F.Payment > 100 is still a valid condition<br>
panic

You should notice Grule execute the same rule again and again because the **WHEN** condition keep yielding a valid result.

One way for this solution is to change "GiveCashback" rule to something like :

```Shell
rule GiveCashback "Give cashback if payment is above 100" {
    When 
         F.Payment > 100 &&
         F.Cashback == 0
    Then
         F.Cashback = 10;
}
```

This way, after the 1st execution, the rule's WHEN is become invalid and not get executed again.
Or ...

```Shell
rule GiveCashback "Give cashback if payment is above 100" {
    When 
         F.Payment > 100
    Then
         F.Cashback = 10;
         Retract("GiveCashback");
}
```

"Retract" function will pull the "GiveCashback" rule from knowledge base so it will not be evaluated again
in the next cycle. When you execute the engine again, all retracted rules will be reset.

---

## 2. Saving Rule Entry to database

**Question** : Will there be a feature that enable Grule to load/save rules from database ?

**Answer** : No. While it is a good idea to store your rule entries into a database, Grule will not create any database adapter to automaticaly store and retrieve rule from database.
You can easily create such adapter your self. Grule have provided a common way to load rules into Knowledgebase from *Reader*, *File*, *Byte Array*, *String* and *Git*. Strings can be easily inserted and selected from database, as you load them into Grule's knowledgebase. 

There are so many database can potentially store Rule Entries. Creating adapter for those databases means we need to get committed to their driver updates, every single one of them. So we decide that, if you want to store rule in database, you can create such mechanism your self.

---

## 3. Maximum number of rule in one knowledge-base

**Question** : How many rule entry can be inserted into knowledgebase ?

**Answer** : You can have as many rule entries you want. But there should be at least one minimum.

---

## 4. Fetch all rules valid for a given fact

**Question** : How can I test which of rules I define is valid for a given Facts?

**Answer** : You can use `engine.FetchMatchingRule` function, refer this [Matching Rules Doc](MatchingRules_en.md) for more info

---

## 5. Rule Engine use-case

**Question** : I have read about rule engine, but what real benefit it can bring ? Give us some use-cases.

**Answer** : The following cases are better solved with a rule-engine in my humble opinion.

1. An expert system that evaluate lots of fact factor to conclude something. Instead of coding lots of if-else in your code, you put the expert system knowledge in form of a rule-set. You supply a fact, and the engine will check them against your rules. It is even more important for such expert system to utilize rule engine if it need to add fast growing sets of facts and knowledge over time. Table based rule engine would require development effort if the fact/rule structure changes, thus not so adaptive.
2. Rating system. Bank system would create a score out of their customer record (facts), and this rating are very dynamic and they change all the time. We don't want those banker to disturb our developer every time they change their rating rules. So we just told them to define all of them in a rule set for our engine. Simply provide the customer facts, and the engine will evaluate them by the rule they define. No need or a development cycle (requirement-build-test-deploy).
3. Computer Games. Player statuses, reward syster, gamifications, etc often requires rules to establish. And these rules is dynamic. Game developer employs scripting language (eg. LUA) to make those dynamic things to be externalize. A fast rule-engine also have room for this case. 
4. Classification systems. Actually, Rating System falls into this category, but so much cases we can take in this category that very suitable for RuleEngine. Eg; classification of credit eligibility, of bio chemical, of risks, of security threat, more and more.
5. Advice/Suggestion system. Since "rule", by virtue, is another kind of data, it can be defined by another analytic or some form of AI program. 

Well, there are so many other use-cases that would benefit from the use of Rule-Engine. The above cases are only small of them that came first on my mind. 

It's important to know that Rule-Engine not a silver-bullet for everything. And one should know when should/shouldn't use a Rule-Engine in their own case.

Another note, some Rule-Engine implementation are very-very expensive and still lots of businesses willing to pay to get their benefit. This shows some prove
that rule-engine really gives them lots of benefit in their business use-cases.

---

## 6. Logging

**Question** : Grule is flooding my log, its so noisy. Can I somehow turn-off Grule's logger ?

**Answer** : Yes. You can reduce (or completely stop) Grule's logging by increasing it's log level.

```go
import (
    "github.com/hyperjumptech/grule-rule-engine/logger"
    "github.com/sirupsen/logrus"
)
...
...
logger.SetLogLevel(logrus.PanicLevel)
```

This will set Grule's log level to `Panic` level, where it will only emits log when it panicked.

Please note that turning Grule's log this way will makes debugging effort become difficult if you find your self
an issue/bug with the library. Do this only in your production environment where performance is paramount.