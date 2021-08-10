# Introduction to Rule Engine

[Tutorial](Tutorial_en.md) | [Rule Engine](RuleEngine_en.md) | [GRL](GRL_en.md) | [GRL JSON](GRL_JSON_en.md) | [RETE Algorithm](RETE_en.md) | [Functions](Function_en.md) | [FAQ](FAQ_en.md) | [Benchmark](Benchmarking_en.md)

---

A Rule Engine, as Martin Fowler explained, is an alternative to the computational model, instead
evaluating multiple conditions, by which an appropriate action is selected if certain
conditions are met. In the simplest explanation, each *Rule* depicts an *if-then* statement.

You feed a collection of rules into a **KnowledgeBase**, and then the *engine* uses each 
rule inside the KnowledgeBase to evaluate some **Facts**. If a rule's requirements are met,
the **action** specified by the selected rule will be executed.

## Fact

A `fact` is a fact, as silly it may sound, but that's what it is. A Fact, in rule engine context,
is basically a piece of information that can be evaluated. Facts can be from any source, eg. a
database, triggered process, point of sale system, report, etc.

It might be much easier to just look at an example of a Fact. Suppose we have this Fact:

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

A **Fact** is basically any piece of information or collected data. 

From this sample Purchasing Fact, we know lots of information: the item being purchased, the quantity,
the purchase date, etc. However, we don't know how much tax should be assigned to that purchase,
how much discount we can give, and the final price the buyer should pay.

## Rule

A Rule is a specification about how to evaluate a **Fact**. If a Rule's
conditions are met by a Fact, then the Rule's action will be selected to be
executed. Sometimes, multiple Rules are selected because their specifications
all apply to a Fact, which results in a conflict. The collection of all Rules in
a conflict are called **Conflict Set**. To resolve this conflict set, we
specify a *strategy* (covered later, below).  

Back to our example of a simple purchasing system: some business Rules should be established in order to
calculate the final price, probably calculating the tax first, and then the discount. If both the tax and 
discount are known, we can show the price.

Let's specify some Rules (in psuedocode).

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

If you examine the above Rules, you should easily understand the of **Rule** concept for Rule Engines. 
These collection of rules will form a set of **Knowledge**. In this case, they form a Knowledge set of
**"how to calculate Item's final price"**.

## Cycle

A Rule evaluation cycle starts by evaluating each Rule's requirements (the **IFs**)
to select which Rules to potentially execute. Every time the engine finds a satisfied
requirement, instead of executing the satisfied Rule's action (the **THENs**), it adds
that Rule into a list of Rule candidates (called the Conflict Set).

When all Rules' requirements have been evaluated, does the engine execute the selected Rules' actions?  
That depends on the contents of the conflict set:

* If there's no Rule with a matching **IF** condition, the Engine execution can immediately finish.
* If there's only one Rule inside the Conflict Set, then that Rule's action is executed before finishing.
* If there are multiple Rules in the Conflict Set, the engine must apply a strategy to prioritize one Rule and execute its action.

If an action gets executed, the cycle repeats again, as long as there's an action that needs execution.
When no more actions get executed, this indicates that there are no more Rules that are statisfied
by the fact (no more matching **IF** statements), and the cycle stops, letting the Rule engine finish evalutation.

The pseudocode for this conflict resolution strategy is depicted below:

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

Grule will keep track of how many cycles it performs in a single evaluation of a ruleset. 
If the Rule evaluation and execution have been repeated too many times, over and above an 
amount specified upon instantiating a Grule engine instance, the engine will terminate and 
an error will be returned.

## Conflict Set Resolution Strategy

As explained above, the Rule engine will evaluate all Rules' requirements and add
them into a list of conflicting Rules called the **Conflict Set**. If only one Rule is
inside the list, it means that are no Rule(s) conflicting with that noe Rule. The engine
will immediately execute the Rule's action.

If there are multiple Rules inside the set, there maybe conflicts. There are many conflict resolution 
strategies that can be implemented for this type of Ruel conflict resolution. The easiest way to resolve
is through specification of a Rule's **salience** (also known as **priority** or **importance**). 
We can add an indicator of Rule **salience** or importance into a Rule definition like in the pseudocode below:

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

By default, all Rules are assigned a salience of `0`.

Because all nonspecified Rules are salience 0, it's easy for the engine to pick which Rule 
to execute when there are multiple Rules in the Conflict Set. If there are multiple Rules 
with matching priorities, the engine will choose the first one found. Because Go's map types 
are not guaranteed to preserve input order, it is not safe to assume that rule evaluation order 
will match the order in which Rules are added to a Grule knowledge instance. 

Salience for Grule Rules can be a value below zero (reaching into the negative) to ensure a 
Rule has even lower priority than the default. This will ensure that a Rule's action will be 
executed last, after all other Rules are evaluated.
