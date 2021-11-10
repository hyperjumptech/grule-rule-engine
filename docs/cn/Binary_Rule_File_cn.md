# 恢复和加载 GRB 文件




[![Binary_Rule_File_cn](https://github.com/yammadev/flag-icons/blob/master/png/CN.png?raw=true)](../cn/Binary_Rule_File_cn.md)
[![Binary_Rule_File_de](https://github.com/yammadev/flag-icons/blob/master/png/DE.png?raw=true)](../de/Binary_Rule_File_de.md)
[![Binary_Rule_File_en](https://github.com/yammadev/flag-icons/blob/master/png/GB.png?raw=true)](../en/Binary_Rule_File_en.md)
[![Binary_Rule_File_id](https://github.com/yammadev/flag-icons/blob/master/png/ID.png?raw=true)](../id/Binary_Rule_File_id.md)

[About](About_cn.md) | [Tutorial](Tutorial_cn.md) | [Rule Engine](RuleEngine_cn.md) | [GRL](GRL_cn.md) | [GRL JSON](GRL_JSON_cn.md) | [RETE Algorithm](RETE_cn.md) | [Functions](Function_cn.md) | [FAQ](FAQ_cn.md) | [Benchmark](Benchmarking_cn.md)

---

当你从GRL脚本中加载大量的(几百以上)规则到`KnowledgeLibrary`, 比如当你启动引擎的时候，你可能会注意到会花费很长的时候去加载，有时候可能会花几分钟。这个主要归结于ANTLR4语法解析。别误会，ANTLR是一个伟大的工具，能够很好的完成工作。但是明显的，对于任意解析工具，数十个几千行的脚本文件都不是一个小工作。

所以解决方案是，把所有的规则存储到二进制文件。这样下次的加载的时候会更快。就如编译器的工作原理一样。你可以编译你的GRL脚本，从二进制（GRB）中加载结果要快10倍。

工作流将会编程：作为规则的作者，你依然可以编辑文本的GRL脚本。当你想要发布你的规则集合时，你可以编译成GRB二进制文件。你发布GRB到你的服务，然后服务从GRB加载。

## 存储 KnowledgeBase 到 GRB

First, you should have a `KnowledgeLibrary` containing the `KnowledgeBase` you want to store into GRB.
Normally you would load a GRL into your library as follows :

首先你得有一个包含想要存储到GRB的`KnowledgeBase`的`KnowledgeLibrary`。如下，你可以正常加载GRL到你的知识库。

```go
	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err := rb.BuildRuleFromResource("HugeRuleSet", "0.0.1", pkg.NewFileResource("HugeRuleSet.grl"))
	assert.NoError(t, err)
```

然后,你可以如下存储knowledge base到GRB：

```go
	f, err := os.OpenFile("HugeRuleSet.grb", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	assert.Nil(t, err)

	// Save the knowledge base into the file and close it.
	err = lib.StoreKnowledgeBaseToWriter(f, "HugeRuleSet", "0.0.1")
	assert.Nil(t, err)
	_ = f.Close()
```

你的GRB文件现在包含了knowledge base的所有规则，然后已经可以被用来加载。

是的，相对于GRL来说，GRB文件大小将会膨胀大概10倍。但是正如大部分时候，其他编译程序将脚本编译成二进制格式。(.java vs .class, .c vs .exe, go vs executable)

## 加载 GRB 到 KnowledgeLibrary

加载GRB更简单，不需要builder。

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

恭喜你,  GRB 会被加载到 `KnowledgeLibrary`，你可以像往常 一样获取knowledge base。

```go
    kb := lib.NewKnowledgeBaseInstance("HugeRuleSet", "0.0.1")
```

还有一件事，如果你的`KnowledgeLibrary`包含了名字和版本号和GRB中一样的`KnowledgeBase`，`KnowledgeBase`将会被覆盖。