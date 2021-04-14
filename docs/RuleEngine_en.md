# Introduction to Rule Engine

[Tutorial](Tutorial_en.md) | [Rule Engine](RuleEngine_en.md) | [GRL](GRL_en.md) | [GRL JSON](GRL_JSON_en.md) | [RETE Algorithm](RETE_en.md) | [Functions](Function_en.md) | [FAQ](FAQ_en.md) | [Benchmark](Benchmarking_en.md)

---

A Rule Engine, as Martin Fowler explained, is an alternative to the computational model.
Where you evaluate multiple conditions, by which you then select an appropriate action if certain
conditions are met. In the simplest explanation, each *Rule* depicts an *if-then* statement.

You feed a collection of rules into a **KnowledgeBase**, then the *engine* uses each 
rule inside the KnowledgeBase to evaluate some **Facts**. If a rule's requirements are met,
the **action** specified by that selected rule will be executed.

## Fact

A `fact` is a fact, as silly it may sound but that's what it is. Fact, in rule engine context,
is basically a piece of information that can be evaluated. Fact can be from any source, eg.
Database, triggered process, point of sales, reports, etc.

If you still can't picture it, it would be much easier to just look
at an example of a fact. Suppose we have this fact:

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

As you can see, a **Fact** is basically any information or data about something.

From this sample Purchasing fact, we know lots of information: the item being purchased, the quantity,
purchase date, etc. However, we don't know how much tax should be assigned to that purchase,
how much discount we can give, and the final price the buyer should pay.

## Rule

A rule is a specification about how to evaluate a **Fact**. If the rule's
conditions are met by the Fact, then the rule's action will be selected to be
executed.  Sometimes, multiple rules are selected because their specifications
all meet the Fact, which results in a conflict. The collection of all rules in
a conflict are called **Conflict Set**. To resolve this conflict set, we
specify a *strategy* that we will cover later.  

Back to our example, as in a simple *purchasing* system, some business rules should be established in order to
calculate the final price, like calculating the tax first and then the discount. When both tax and discount
are known, we can show the price.

Let's specify some rules.

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

If you examine the above rules, you should easily understand the meaning
of **rule**. These collection of rules will form a **Knowledge**. A knowledge of
**"How to calculate Item's final price"**.

## Cycle

A rule evaluation cycle will start with evaluating each of the rule's requirements (the **IFs**)
to select which rule to potentially execute. Every time the engine finds a satisfied
requirement, instead of executing the satisfied rule's action (the **THENs**), it adds
that rule into a list of rule candidate (Conflict Set).

When all rules' requirements have been evaluated,
will the engine execute the selected rule's action?  That depends on the contents of the conflict set:

* If there's no Rule inside, that means the Engine execution is finished.
* If theres only 1 rule inside the conflict set then that action is executed immediately.
* If there are multiple rules, the engine need to apply some strategy to select one rule and execute its action.

If an action gets executed, the cycle repeats again as long as there's an action that needs execution.
When no more actions get executed, this means that there are no more rule that are statisfied
by the fact, the cycle stops and the rule engine is finished.

The *Pseudo Code* for this depicted bellow

```text
Start Engine With a FACT using a KNOWLEDGE
BEGIN
    For Every RULE in KNOWLEDGE
        Check if RULE's Requirement is satisfied by FACT
            If RULE's Requirement is satisfied
                Add RULE into CONFLICT SET
            End If
        End Check
    End For
    If CONFLICT SET is EMPTY
        Finished
        END
    If CONFLICT SET has 1 RULE
        Execute the RULE's Action
        Clear CONFLICT SET
        Repeat cycle from BEGIN
    If CONFLICT SET has Many RULEs
        Apply Conflict resolution strategy to choose 1 RULE.
        Execute the chosen RULE's Action
        Clear CONFLICT SET
        Repeat cycle from BEGIN
END
```

Grule will keep track of how many cycles it has been repeating. If the rule evaluation and execution have
been repeated too many times, above a specific amount of cycles, the engine will terminate
and an error will be emitted.

## Conflict Set Resolution Strategy

As explained above, the rule engine will evaluate all rules' requirements and add
them into a list of conflicting rules called the **Conflict Set**. If only 1 rule is
inside the list, it means that are no rule(s) conflicting with that 1 rule. The engine
will immediately execute the rule's action.

If there are multiple rules, there maybe conflicts. There are many conflict resolution strategies that
can be implemented. The easiest way to do it is by specifying the rule's **salience** (also known as
**priority** or **importance**). We add some indicator into the rule definition such as:

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

...
```

By default, all rules are assigned a salience of `0`.

This way, it's easy for the engine to pick which rule to execute when there are multiple
conflicting rules. If there are multiple rules with similarly top priorities, the engine
will pick the first one found.

Salience can be a value below zero (negative) to ensure the rule has even lower
priority. This will to ensure that the rule's action will be executed last,
after all other rules are met.
