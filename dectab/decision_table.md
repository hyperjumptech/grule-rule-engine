# Decision Table

Status : DRAFT

Decision table is one of the Rule Engine modeling approach. With decision table approach
its easy to model rule criteria in evaluating facts and also easy to define
action when a fact matched the criteria. 

With decision table approach, we can create a simple user-interface to be used by 
end user to create and modify rules as needed. It also can serve as a template rule model
which later can be translated into a more elaborated, fine-grained and flexible rule definition
like a GRL.

in-fact, this is the proposed approach as a step before running the
decision table in Grule engine as depicted in the following flow.

```text
+-------------------+ 
| Rule Table Editor |
+-------------------|
   |             ^                                                         ( Fact )
  save         load                                                            |
   V             |                                                             V
+-------------------+              +------------+         +------------------------+
|  Grule DMN JSON   |--translate-->| GRL script |--load-->| Grule Engine & Execute |
+-------------------+              +------------+         +------------------------+
```

## Standards and Roadmap

This document should adhere to implementing the [DMN 1.3 standard](https://www.omg.org/spec/DMN/1.3/PDF) whenever appropriate, doable and 
compatible with GRL Engine. Because of the wide coverage of aspect in Decision Model, not all of the
specification described there in, this DMN implementation in Grule Engine will not implemented all points in DMN 1.3 specification.

### Phase 1 - MVP (Minimum Viable Product) - Simple Decision Table

The implementation of basic DMN capability as depicted in "DMN 1.3 standard - 5.3.1 Decision requirements level - Figure 5.3"

- Decision table JSON representation.
- GRL Expression based information binding for input.
- GRL Expression based invocation for output.
- Decision Table structure for Rules "DMN 1.3 standard - 8.1 Introduction - Figure 8.1, 8.2, 8.3, 8.4"
- Translation from Decision Table JSON to GRL.
- Ability to validate the Decision Table correctness.

### Phase 2 - MLP (Minimum Likeable Product) - Editor for Grule's Decision Table

- WEB user interface to work with decision table
- UI Ability to specify inputs and outputs
- UI Ability to specify types, labels, allowed values and default values for each inputs and outputs.
- Ability to save a Decision Table into DMN 1.3 styled JSON
- Ability to load a Decision DMN 1.3 styled JSON into Decision Table UI

## Decision Table

### Decision Table metamodel

#### preferredOrientation

The Grule implementation for DMN 1.3 will always have the _Rule-as-Row_ orientation.

#### hitPolicy

The Grule implementation for DMN 1.3 in this document will have a _RULE ORDER_ policy. 
But this capability may be expanded into other hitpolicy capability.

#### inputExpression

The input expressions will always logically evaluated with AND logical expression.
This will ensure the following clauses "The i-th inputExpression must satisfy the i-th
input Entry for all inputEntrys in order for the DecisionRule to match as defined in section 8.1"
- DMN 1.3 standard - 8.3.3 Decision Rule metamodel - pg.78

The following table are a simple decision table.

| No | Description | Information Item 1 | Information Item 2 |
|----|----| ------ |:--------------------------:|
| -  | _name_ | type  | grade                       |
| -  | _function_ | input  | output                       |
| -  | _type_ | string  | string                       |
| -  | _labels_ | "Goods Type"  |  "Good Grade"                       |
| -  | _allowed values_ | "electronic","machine","electric appliance"   |  "A", "B","C","D" |
| -  | _default values_ | any  |  "C"  |
| 1  | Electronic and machinery are all A grade | in("electronic","machine") | "A" |
| 2  | House electric powered appliances are all B grade  | "electric appliance" | "B" |
| 3  | All other items are C grade | any | "C" |

As you may've guessed, the decision table above speaks about mapping from "Good Type" to "Good Grade".
Its a straight forward rule to decide, if a good if of type "X" than it should be mapped as grade "Xa".

**"any" - keyword**

The "any" keyword specified in the input expression means that the input should be ignored in the evaluation. Thus
in the matching algorithm, the variable will be ignored during evaluation.

The "any" keyword specified in the output expression means that the output value should be equals to the default value. 

The use of "any" must be used alone in the expression. (e.g. `any > 200` is not allowed)

### name 

This is a fact's name property, accessible by the rule engine. 
For example, consider the following JSON fact :

```json
{
  "type": "value 1",
  "grade": ""
}
```

then we can see at least 2 possible name with native type values.

- `type`
- `grade`

The table define this using the `name` definition

| No | Description | Information Item 1 | Information Item 2 |
|----|----| ------ |:--------------------------:|
| -  | _name_ | type  | grade                       |

### function

"function" specifies if a certain variable is used in the rule evaluation "when" scope, or to be 
assigned when the rule match (to be changed in the "then" scope). 

For example

| No | Description | Information Item 1 | Information Item 2 |
|----|----| ------ |:--------------------------:|
| -  | _name_ | type  | grade                       |
| -  | **function** | **input**  | **output**                       |
| 1  | Electronic and machinery are all A grade | In("electronic","machine") | "A" |

Above you can see that fact `type` have an `input` function and `grade` have an `output` function.
Then the GRL would look something as follows:

```
rule Rule_1 "Electronic and machinery are all A grade" salience 1 {
when
    type.In("electronic","machine")
then
    grade = "A";
    Complete();
}
```

As you can see, `type` is the fact name to be used in evaluation `when` scope and
`grade` to be assigned in the `then` scope. 

### type

"type" specify the fact's item golang data type. This used as a hint to the engine on how to evaluate
the fact. For example:

| No | Description | Information Item 1 | Information Item 2 |
|----|----| ------ |:--------------------------:|
| -  | _name_ | type  | grade                       |
| -  | **type** | **string**  | **string**                       |

Here you can see that the fact `type` have variable type of `string`. The same with
fact `grade` which also a `string`.

The valid datatype supported for Decision table would be `string`, `datetime`, `int`, `float`, `bool` 

### label

"label" is an information to be displayed in the Decision Table UI Designed.

| No | Description | Information Item 1 | Information Item 2 |
|----|----| ------ |:--------------------------:|
| -  | _name_ | type  | grade                       |
| -  | _labels_ | "Goods Type"  |  "Good Grade"     |

### allowed_values

"allowed_values" defines all possible values for every inputs and outputs facts.

| No | Description | Information Item 1 | Information Item 2 |
|----|----| ------ |:--------------------------:|
| -  | _name_ | type  | grade                       |
| -  | _allowed values_ | "electronic","machine","electric appliance"   |  "A", "B","C","D"  |

Here you can see that for fact `type`, all possible values are "electronic","machine","electric appliance", any
The same with `grade` which can only have one of "A", "B","C" or "D"

In the input type information item, the use of `any` keyword is used to accept any input
as long as the data value type equals to required input type.

In the output type information item, the user of `any` keyword is used to signify that 
it would returned what ever the default value is specified.

**Set of possible values**

For `string`, `int`, `float` type, you can supply a speciffic set of possible values.

- "A", "B", "Z"
- 1, 4, 6
- 0.2, 1.23, 23.45

These values must conform to the `type` format.

**Range of possible values**

For `numeric` type (`int` and `float`) and `date-time`, you can specify a range.
You have to specify the lower and upper limit for a range.

- 2..3
- -23..23
- -34.56..78.9
- "2007-01-01T13:00:00Z".."2009-12-31T13:00:00Z"

These values must conform to the `type` format.

**Interval of possible values**

For `numeric` type (`int` and `float`) and `date-time`, you can specify an interval
where you combine the `set` and range `rage`.

- 2..3,14,25,36..50
- -23..23,60..90
- -34.56..78.9,93.2,120.3..150.0
- "2007-01-01T13:00:00Z".."2009-12-31T13:00:00Z", "2012-12-31T13:00:00Z"

These values must conform to the `type` format.

### default_value

| No | Description | Information Item 1 | Information Item 2 |
|----|----| ------ |:--------------------------:|
| -  | _name_ | type  | grade                       |
| -  | _default values_ | any  |  any  |

As the name implies, `default_value` specify a value for the specified fact
if the fact value is:

- For input, the value is not supplied (or empty) during evaluation operation.
- For input, the value is not within the `allowed_value`
- For output, the value would be equals to the input value.

The `default_value` must exist in the `allowed_value` unless the use of 'any' which means that
it can accept any input as long as the type is correct.


### Rule Order / Hit Order - Salience

Within the table, you will see evaluation order. This is a positive integer value and it started from number 1.
The value denotes evaluation order of the rule row. 

"hit policy (H) and rule numbers as indicated in Figure 8-5, Figure 8-7 and
Figure 8-9. Rule numbers are consecutive natural numbers starting at 1. Rule numbering is required for tables
with hit indicator F (first) or R (rule order), because the meaning depends on rule sequence. Crosstab tables have
no rule numbers." - DMN 1.3 standard - 8.2 Notation pg.67


The table will have the evaluation order (the "No" column), optional descrition,
input columns (those remarked with "&lt;in&gt;") and an output (remarked with "&lt;out&gt;")

If you're familiar with Grule's GRL script, the inputs would be variables to be evaluated
in the `when` scope, and the outputs are variables to be set in the `then` scopes.

### Decision Table's Fact Item Evaluation

### Decision Table Errors

## Examples

### Applicant Risk Rating

| No | Description | Information Item 1   | Information Item 2  | Information Item 3 |
|----|----|------|----------------------------|--------------------|
| -  | _name_ | age  | history     | rating|
| -  | _function_ | input  | input  | output|
| -  | _type_ | int  | string    | string|
| -  | _labels_ | "Applicant Age"  |  "Medical History"         | "Applicant Risk Rating" |
| -  | _allowed values_ |  0..200  | "good", "bad" | "high", "medium", "low" |
| -  | _default values_ | 30  |  "good"  | "medium" | 
| 1  | Old man with good medical history | &gt; 60      | "good"  | "medium"          |
| 2  | Old man with bad medical history  | &gt; 60      | "bad"   | "high"            |
| 3  | Adult productive age | [25..60]     | any   | "medium"               |
| 4  | Youngster with good medical history  | &lt; 25      | "good"  | "low"               |
| 5  | Youngster with bad medical history  | &lt; 25      | "bad"   | "medium"               |

**Evaluation Sample**

| age | history | rating | note |
|-----|---------|--------|------|
| 20  | "good"  | "low"  | rule 4 |
| 30  | "ugly"  | error  | not conform to __allowed values__ |
| 300 | "bad"   | error  | not conform to __allowed values__ |
| 60  | "bad"   | "medium" | rule 3 |

---

### Flow Throttle

| No | Description | Information Item 1   | Information Item 2|
|----|-------------|----------------------|--------------------|
| -  | _name_ | intake  | throughput |
| -  | _function_ | input  |  output|
| -  | _type_ | int  | int |
| -  | _labels_ | "Water Intake L/s"  |  "Water Througput"         |
| -  | _allowed values_ |  any  | any |
| -  | _default values_ |  30   | any | 
| 1  | Under flow | &lt; 20      | 0      |
| 2  | Normal flow  | [20..80]   | intake |
| 3  | Over flow    | &gt; 80    | 80     |

**Evaluation Sample**

| intake | throughput |  note |
|-----|---------|----------|
| 10  | 0       | rule 1 |
| 20  | 20      | rule 2 |
| 50  | 50      | rule 2 |
| 80  | 80      | rule 2 |
| -60 | 0       | rule 1 |
| 81  | 80      | rule 3 |
| n/a | 30      | intake default value = 30 -> rule 2 |

---

### Person Loan Compliance

| No | Description | Information Item 1   | Information Item 2  | Information Item 3 | Information Item 4 |
|----|----|------|----------------------------|--------------------|----|
| -  | _name_ | rating  | cc_balance     | loan_balance | compliance |
| -  | _function_ | input  | input  | input | output|
| -  | _type_ | string  | int    | int | string |
| -  | _labels_ | "Persons Credit Rating from Bureau"  |  "Person Credit Card Balance" | "Persons Education Loan Balance" | "Person Loan Compliance" |
| -  | _allowed values_ |  "A", "B", "C", "D" | &gt;=0 | &gt;=0 | "Compliant","Not Compliant" |
| -  | _default values_ | any  |  any  | any | "Not Compliant" |
| 1  | A grade Student with low CC debt and low loan balance is comply | "A"      | &lt; 10000  | &gt; 50000          | "Compliant" |
| 2  | Other than A grade student not comply  | !="A" | any   | any            | "Not Compliant" |
| 3  | Any grade with lots of CC debt not comply | any     | &gt;= 10000   | any               | "Not Compliant" |
| 4  | Any grade with high loan balance not comply   | any     | any  | &gt;= 50000               |"Not Compliant" |

---

### Special Discount


| No | Description | Information Item 1   | Information Item 2  | Information Item 3 | Information Item 4 |
|----|----|------|----------------------------|--------------------|----:|
| -  | _name_ | order  | location     | type | discount |
| -  | _function_ | input  | input  | input | output|
| -  | _type_ | string  | string    | string | int |
| -  | _labels_ | "Type of Order"  |  "Customer Location" | "Type of Customer" | "Specioal Discount %" |
| -  | _allowed values_ |  "Web", "Phone", "Whatsapp", "Email" | "US", "DE", "CN" | "Retailer", "Wholesaler", "Personal" | [0..100] |
| -  | _default values_ | any  |  any  | any | 0 |
| 1  | US wholesaler order from WEB | "Web"      | "US"  | "Wholesaler"          | 10 |
| 2  | Any phone order   | "Phone"     | any  | any               |0 |
| 3  | Non US customer  | any | !="US"   | any            | 0 |
| 4  | Any Retailer | any     | &gt;= 10000   | any               | 5 |

---

### Holidays

| No | Description | Information Item 1   | Information Item 2  | Information Item 3|
|----|----|------|----------------------------|--------------------|
| -  | _name_ | age  | service_year     | Holiday |
| -  | _function_ | input  | input  | output |
| -  | _type_ | int  | int    | int|
| -  | _labels_ | "Age"  |  "Years of Service" | "Holidays"|
| -  | _allowed values_ | [0..200] | [0..200] | 22,5,3,2 |
| -  | _default values_ | 200  |  200  | 22 |
| 1  |  | any      | any  | 22 |
| 2  |  | &gt;=60     | any  | 3 |
| 3  |  | any | &gt;=30   | 3 |
| 4  |  | &lt;18    | any  | 5 |
| 5  |  | &gt;60    | any  | 5 |
| 6  |  | any    | &gt;=30  | 5 |
| 7  |  | [18..60]    | [15..30]  | 2 |
| 4  |  | [45..60]    | &lt;30  | 2 |

---

### Insurance Based on Goods Grade and Price

| No | Description | Information Item 1   | Information Item 2  | Information Item 3  | Information Item 3  |
|----|----| :------: |----------------------------:|--------------------|---------------:|
| -  | _name_ | grade  | amount     | insurance | rate |
| -  | _function_ | input  | input  | output | output |
| -  | _type_ | string  | int    | bool | float |
| -  | _labels_ | "Grade"  |  "Amount of Loan"         | "Insurance Required" | "Insurance Rate |
| -  | _allowed values_ |  "A", "B","C","D"  | 0..999999999 | true, false | 0..1.0
| -  | _default values_ | any  |  0  | true | 0.002 | 
| 1  | Anything bellow 100000 do not need insurance  | any      | &lt; 100000                 | false              | 0              |
| 2  | Grade A with price between 100000 to 300000 will have 0.001 insurance rate | A        | [100000..299999]   | true               | 0.001          |
| 3  | Grade A with price between 300000 to 600000 will have 0.003 insurance rate | A        | [300000..599999]   | true               | 0.003          |
| 4  | Any other grade between 100000 to 600000 will have 0.002 insurance rate | any      | [100000..599999]   | true               | 0.002          |
| 5  | Price above 600000 will have 0.005 insurance rate flat | any      | &gt; 600000                 | true               | 0.005          |

```json
{
  "table_version": "1.0",
  "name": "InsuranceAmountRule",
  "description": "Insurance Based on Goods Grade and Price",
  "version": "1.2.3",
  "items": [
    {
      "name": "grade",
      "function": "input",
      "label": "Grade",
      "type": "string",
      "allowed": [
        {
          "set": [
            "A",
            "B",
            "C",
            "D"
          ]
        }
      ],
      "default": "any"
    },{
      "name": "amount",
      "function": "input",
      "label": "Loan Amount",
      "type": "int",
      "allowed": {
          "ranges": [
            {
              "min": 0,
              "max": 999999999
            }
          ]
        },
      "default": 0
    },{
      "name": "insurance",
      "function": "output",
      "label": "Insurance Required",
      "type": "bool",
      "default": false
    },{
      "name": "rate",
      "function": "output",
      "label": "Insurance Rate",
      "type": "float",
      "allowed": {
        "ranges": [
          {
            "min": 0.0,
            "max": 1.0
          }
        ]
      },
      "default": 0.0
    }
  ],
  "decision_rows": [
    {
      "hit": 1,
      "description": "Anything bellow 100000 do not need insurance",
      "input" : {
        "grade": "any",
        "amount": "< 100000"
      }
    }, {
      "hit": 2,
      "description": "Grade A with price between 100000 to 300000 will have 0.001 insurance rate",
      "input" : {
        "grade": "A",
        "amount": "100000..299999"
      },
      "output" : {
        "insurance": true,
        "rate": 0.001
      }
    }, {
      "hit": 3,
      "description": "Grade A with price between 300000 to 600000 will have 0.003 insurance rate",
      "input" : {
        "grade":  "A",
        "amount": "[300000..599999]"
      },
      "output" : {
        "insurance": true,
        "rate": 0.003
      }
    }, {
      "hit": 4,
      "description": "Any other grade between 100000 to 600000 will have 0.002 insurance rate",
      "input" : {
        "grade": "any",
        "amount": "[100000..599999]"
      },
      "output" : {
        "insurance": true,
        "rate": 0.002
      }
    }, {
      "hit": 5,
      "description": "Price above 600000 will have 0.005 insurance rate flat",
      "input" : {
        "grade":"any",
        "amount": ">=600000"
      },
      "output" : {
        "insurance": true,
        "rate": 0.003
      }
    }
  ]
}
```