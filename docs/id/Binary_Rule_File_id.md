# Menyimpan dan Memuat File GRB

---

:construction:
__THIS PAGE IS BEING TRANSLATED__
:construction:

:construction_worker: Contributors are invited. Please read [CONTRIBUTING](../../CONTRIBUTING.md) and [CONTRIBUTING TRANSLATION](../CONTRIBUTING_TRANSLATION.md) guidelines.

:vulcan_salute: Please remove this note once you're done translating.

---


[![Binary_Rule_File_cn](https://github.com/yammadev/flag-icons/blob/master/png/CN.png?raw=true)](../cn/Binary_Rule_File_cn.md)
[![Binary_Rule_File_de](https://github.com/yammadev/flag-icons/blob/master/png/DE.png?raw=true)](../de/Binary_Rule_File_de.md)
[![Binary_Rule_File_en](https://github.com/yammadev/flag-icons/blob/master/png/GB.png?raw=true)](../en/Binary_Rule_File_en.md)
[![Binary_Rule_File_id](https://github.com/yammadev/flag-icons/blob/master/png/ID.png?raw=true)](../id/Binary_Rule_File_id.md)
[![Binary_Rule_File_in](https://github.com/yammadev/flag-icons/blob/master/png/IN.png?raw=true)](../in/Binary_Rule_File_in.md)

[Tentang Grule](About_id.md) | [Tutorial](Tutorial_id.md) | [Rule Engine](RuleEngine_id.md) | [GRL](GRL_id.md) | [GRL JSON](GRL_JSON_id.md) | [Algoritma RETE](RETE_id.md) | [Fungsi-fungsi](Function_id.md) | [FAQ](FAQ_id.md) | [Benchmark](Benchmarking_id.md)

---

When you loading huge amount (hundreds) of rules in GRL script into `KnowledgeLibrary`, e.g. when you start 
the engine, you may notice that it may took some time to load, some time it could go up to a couple of minutes.
This is due to the syntax parsing done by ANTLR4. Don't get me wrong, ANTLR is a great tools and it done the job very well.
But obviously, tens of thousands of lines in a script file is no small task, for any parser tools.

So the idea is, to store all the rule in a binary file. So it load faster next time. Just like
a compiler. You compile your text GRL script and the result is a binary file (GRB) which load 10 time faster.

The workflow is as following : As a Rule author, you still work in the textual GRL script. When you want to release your rule set,
you can "compile" it into GRB binary file. The you ditribute that GRB into your server and have the server load
from GRB.

## Storing KnowledgeBase into GRB

First, you should have a `KnowledgeLibrary` containing the `KnowledgeBase` you want to store into GRB.
Normally you would load a GRL into your library as follows :

```go
	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err := rb.BuildRuleFromResource("HugeRuleSet", "0.0.1", pkg.NewFileResource("HugeRuleSet.grl"))
	assert.NoError(t, err)
```

Second, you can now save the knowledge base into GRB as follows:

```go
	f, err := os.OpenFile("HugeRuleSet.grb", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	assert.Nil(t, err)

	// Save the knowledge base into the file and close it.
	err = lib.StoreKnowledgeBaseToWriter(f, "HugeRuleSet", "0.0.1")
	assert.Nil(t, err)
	_ = f.Close()
```

Your GRB file is now contains all rules in the specified knowledge base
and ready for future loading.

Yes, the GRB file size is inflated like 10 times compared to the GRL one, 
But it, most of the time, like that when you compile some script into its 
compiled binary form. (.java vs .class, .c vs .exe, go vs executable)

## Loading GRB into KnowledgeLibrary

Loading GRB is much simpler. No need a builder.

```go
	lib := ast.NewKnowledgeLibrary()

	// Open the existing safe file
	f, err := os.Open("HugeRuleSet.grb")
	assert.Nil(t, err)

	// Load the file directly into the library and close the file
	// btw, you should not use the blueprint_kb in your engine execution.
	bluerint_kb, err := lib.LoadKnowledgeBaseFromReader(f2, true)
	assert.Nil(t, err)
	_ = f.Close()
```

There you go !!!, the GRB is loaded into `KnowledgeLibrary`
You can obtain the knowledge base normally.

```go
    kb := lib.NewKnowledgeBaseInstance("HugeRuleSet", "0.0.1")
```

One thing, if in your `KnowledgeLibrary` already contains the same `KnowledgeBase` name and version
to the one in the GRB, that `KnowledgeBase` in the library will be overwritten.