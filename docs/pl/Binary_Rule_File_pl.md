# Przechowywanie i ładowanie pliku GRB.

[![Binary_Rule_File_cn](https://github.com/yammadev/flag-icons/blob/master/png/CN.png?raw=true)](../cn/Binary_Rule_File_cn.md)
[![Binary_Rule_File_de](https://github.com/yammadev/flag-icons/blob/master/png/DE.png?raw=true)](../de/Binary_Rule_File_de.md)
[![Binary_Rule_File_en](https://github.com/yammadev/flag-icons/blob/master/png/GB.png?raw=true)](../en/Binary_Rule_File_en.md)
[![Binary_Rule_File_id](https://github.com/yammadev/flag-icons/blob/master/png/ID.png?raw=true)](../id/Binary_Rule_File_id.md)
[![Binary_Rule_File_pl](https://github.com/yammadev/flag-icons/blob/master/png/PL.png?raw=true)](../pl/Binary_Rule_File_pl.md)

[About](About_pl.md) | [Tutorial](Tutorial_pl.md) | [Rule Engine](RuleEngine_pl.md) | [GRL](GRL_pl.md) | [GRL JSON](GRL_JSON_pl.md) | [RETE Algorithm](RETE_pl.md) | [Functions](Function_pl.md) | [FAQ](FAQ_pl.md) | [Benchmark](Benchmarking_pl.md)

---

Podczas ładowania dużej ilości (setek) reguł w skrypcie GRL do `KnowledgeLibrary`, np. po uruchomieniu silnika, można zauważyć, że ładowanie może zająć trochę czasu, czasami nawet kilka minut. Wynika to z parsowania składni przez ANTLR4. Nie zrozum mnie źle, ANTLR jest świetnym narzędziem i bardzo dobrze wykonuje swoją pracę. Ale oczywiście dziesiątki tysięcy linii w pliku skryptowym to nie lada wyzwanie dla każdego narzędzia parsującego.

Chodzi więc o to, aby wszystkie reguły przechowywać w pliku binarnym. Dzięki temu następnym razem załaduje się on szybciej. Zupełnie jak kompilator. Kompilujesz swój tekstowy skrypt GRL, a rezultatem jest plik binarny (GRB), który ładuje się 10 razy szybciej.

Przebieg pracy jest następujący: Jako autor reguł nadal pracujesz w tekstowym skrypcie GRL. Gdy chcesz opublikować swój zbiór reguł, możesz go "skompilować" do pliku binarnego GRB. Następnie przesyłasz GRB na swój serwer, a serwer ładuje się z GRB.

## Przechowywanie bazy wiedzy w GRB

Po pierwsze, powinieneś mieć `KnowledgeLibrary` zawierającą `KnowledgeBase`, którą chcesz przechowywać w GRB.
Zwykle ładuje się GRL do biblioteki w następujący sposób :

```go
	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err := rb.BuildRuleFromResource("HugeRuleSet", "0.0.1", pkg.NewFileResource("HugeRuleSet.grl"))
	assert.NoError(t, err)
```

Po drugie, można teraz zapisać bazę wiedzy w GRB w następujący sposób:

```go
	f, err := os.OpenFile("HugeRuleSet.grb", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	assert.Nil(t, err)

	// Save the knowledge base into the file and close it.
	err = lib.StoreKnowledgeBaseToWriter(f, "HugeRuleSet", "0.0.1")
	assert.Nil(t, err)
	_ = f.Close()
```

Plik GRB zawiera teraz wszystkie reguły z określonej bazy wiedzy i jest gotowy do załadowania w przyszłości.

Tak, rozmiar pliku GRB jest 10-krotnie większy niż GRL, ale najczęściej dzieje się tak, gdy kompilujesz jakiś skrypt do jego postaci binarnej. (.java vs .class, .c vs .exe, go vs executable)

## Ładowanie GRB do Biblioteki Wiedzy

Ładowanie GRB jest znacznie prostsze. Nie potrzeba konstruktora.

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

Proszę bardzo!!!, GRB jest załadowana do `KnowledgeLibrary` Możesz normalnie uzyskać bazę wiedzy.

```go
    kb := lib.NewKnowledgeBaseInstance("HugeRuleSet", "0.0.1")
```

Jedna rzecz, jeśli `KnowledgeLibrary` zawiera już taką samą bazę wiedzy jak ta w GRB, ta `KnowledgeBase` w bibliotece zostanie nadpisana.
