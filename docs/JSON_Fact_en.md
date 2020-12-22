# JSON Fact

[Tutorial](Tutorial_en.md) | [Rule Engine](RuleEngine_en.md) | [GRL](GRL_en.md) | [GRL JSON](GRL_JSON_en.md) | [RETE Algorithm](RETE_en.md) | [Functions](Function_en.md) | [FAQ](FAQ_en.md) | [Benchmark](Benchmarking_en.md)

---

Using JSON straight away as fact in Grule is available starting on version 1.8.0. It enable user to treat JSON string as fact
and add it into `DataContext` just like how you previously add fact data into it. The loaded JSON fact are now "visible" the
the Grule scripts (the GRLs).

## Adding JSON as fact

Assuming you have a JSON as follow:

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

You put your JSON into a byte array.

```go
myJSON := []byte (...your JSON here...)
```

You simply add you JSON variable into `DataContext`

```go
// create new instance of DataContext
dataContext := ast.NewDataContext()

// add your JSON Fact into data cotnxt using AddJSON() function.
err := dataContext.AddJSON("MyJSON", myJSON)
```

Yes, you can add as many _facts_ as you wish into the context and you can mix between JSON facts
(using `AddJSON`) and normal Go fact (using `Add`)

 ## Evaluating (Reading) JSON Fact Values in GRL
 
 Inside GRL script, the fact is always visible through their label as you provide them 
 when adding to the `DataContext`. For example, the code bellow add your JSON and it will be 
 using label `MyJSON`.
 
 ```go
err := dataContext.AddJSON("MyJSON", myJSON)
```
 
 Yes, you can use any label as long as its a single word.
 
 ### Traversing member variable like a normal object
 
 Now. Using the JSON shown at the beginning, your GRL `when` scope can evaluate your json 
 like the following.
 
 ```text
when
    MyJSON.name == "John Doe"
``` 

or 

```text
when
    MyJSON.address.city.StrContains("ville")
```

or

```text
when
    MyJSON.age > 30 && MyJSON.height < 60
```

### Traversing member variable like a map

You can access JSON object's fields using `Map` like selector or like normal object.

 ```text
when
    MyJSON["name"] == "John Doe"
``` 

or 

```text
when
    MyJSON["address"].city.StrContains("ville")
```

or

```text
when
    MyJSON.age > 30 && MyJSON["HEIGHT".ToLower()] < 60
```

### Traversing array member variable

You can inspect JSON Array element just like a normal array

 ```text
when
    MyJSON.friends[3] == "Jake"
```

## Writing values into JSON Facts in GRL

Yes, you can write new values into you JSON facts in the `then` scope of your rules. Changing those values will
certainly evaluated on the following rule evaluation cycles. BUT, there are some caveat (read "Things you should know" bellow.)

### Writing member variable like a normal object
 
Now. Using the JSON shown at the beginning, your GRL `then` scope can modify your json 
**fact** like the following.
 
 ```text
then
    MyJSON.name = "Robert Woo";
``` 

or 

```text
then
    MyJSON.address.city = "Corruscant";
```

or

```text
then
    MyJSON.age = 30;
```

That's pretty straight forward. But there are some twist to this.

1. You can modify not only the value of member variable of your JSON object, you can also change the `type`.
   Assuming your rule can handle the next evaluation chain for the new type you can do this, otherwise we very strongly not recommended this.
   
   Example:
   
   You modify the `MyJSON.age` into string.
   
   ```text
    then
        MyJSON.age = "Thirty";
   ```
   
   This make the engine to panic when evaluating rule like.
   
   ```text
    when
        myJSON.age > 25
   ```
   
2. You can assign a value to non-existent member variable
 
   Example:
   
      ```text
       then
           MyJSON.category = "FAT";
      ```

    Where the `category` member is not existed in the original JSON.
    
### Writing member variable like a normal map
 
Now. Using the JSON shown at the beginning, your GRL `then` scope can modify your json 
**fact** like the following.
 
 ```text
then
    MyJSON["name"] = "Robert Woo";
``` 

or 

```text
then
    MyJSON["address"]["city"] = "Corruscant";
```

or

```text
then
    MyJSON["age"] = 30;
```

Like the object style, there are same exact twist to this.

1. You can modify not only the value of member variable of your JSON map, you can also change the `type`.
   Assuming your rule can handle the next evaluation chain for the new type you can do this, otherwise we very strongly not recommended this.
   
   Example:
   
   You modify the `MyJSON.age` into string.
   
   ```text
    then
        MyJSON["age"] = "Thirty";
   ```
   
   This make the engine to panic when evaluating rule like.
   
   ```text
    when
        myJSON.age > 25
   ```
   
2. You can assign a value to non-existent member variable
 
   Example:
   
      ```text
       then
           MyJSON["category"] = "FAT";
      ```

    Where the `category` member is not existed in the original JSON.

### Writing member array

In array you can simple replace array element by it's indices.

```text
then
   MyJSON.friends[3] == "Jake";
```

As long as that indice is a valid one. Grule will panic if the indices is out of bound.
Just like normal JSON, you can replace the value of any element with different type.
You can always inspect the array length. Like ...

```text
when
   MyJSON.friends.Length() > 4;
```

Yes, you can always append into Array using `Append` function. Append list of argument value of different types.

```text
then
   MyJSON.friends.Append("Rubby", "Anderson", "Smith", 12.3);
```

**Known Issue**

As of now, there are no built-in function to help user to inspect array element easily, such as `Contains(value) bool`

## Things you should know

1. After you add JSON fact into `DataContext`, the change to this string variable will not reflect the facts already in the `DataContext`. This is also
   applied in vice-versa, where changes in the fact within `DataContext` will not change the JSON string.
2. You can modify your JSON fact in the `then` scope, but unlike normal `Go` facts, these changes will not reflect to your original JSON string. If you want this to happen, 
   you should parse your JSON into a `struct` before hand, and add your `struct` into `DataContext` normally. 