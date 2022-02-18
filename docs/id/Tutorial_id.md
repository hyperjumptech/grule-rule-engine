# Tutorial Singkat GRULE

---

:construction:
__THIS PAGE IS BEING TRANSLATED__
:construction:

:construction_worker: Contributors are invited. Please read [CONTRIBUTING](../../CONTRIBUTING.md) and [CONTRIBUTING TRANSLATION](../CONTRIBUTING_TRANSLATION.md) guidelines.

:vulcan_salute: Please remove this note once you're done translating.

---


[![Tutorial_cn](https://github.com/yammadev/flag-icons/blob/master/png/CN.png?raw=true)](../cn/Tutorial_cn.md)
[![Tutorial_de](https://github.com/yammadev/flag-icons/blob/master/png/DE.png?raw=true)](../de/Tutorial_de.md)
[![Tutorial_en](https://github.com/yammadev/flag-icons/blob/master/png/GB.png?raw=true)](../en/Tutorial_en.md)
[![Tutorial_id](https://github.com/yammadev/flag-icons/blob/master/png/ID.png?raw=true)](../id/Tutorial_id.md)

[Tentang Grule](About_id.md) | [Tutorial](Tutorial_id.md) | [Rule Engine](RuleEngine_id.md) | [GRL](GRL_id.md) | [GRL JSON](GRL_JSON_id.md) | [Algoritma RETE](RETE_id.md) | [Fungsi-fungsi](Function_id.md) | [FAQ](FAQ_id.md) | [Benchmark](Benchmarking_id.md)

---

## Persiapan

Mohon dicatat bahwa Grule menggunakan Go 1.13

Untuk menggunakan Grule didalam proyek anda, cukup dengan mudah masukannya.

```text
$go get github.com/hyperjumptech/grule-rule-engine
```

Dalam file `.go` anda,

```go
import (
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
) 
``` 

## Membuat struktur fakta

Sebuah `fakta` dalam grule adalah hanya sebuah **pointer** kepada sebua *instance* dari suatu `struktur`
Struktur ini bisa berisi properti-propert seperti halnya struktur normal golang biasa, dan juga berisi `fungsi`
atau yang daldam dunia *OOP* disebut juga dengan `method`

```go
type MyFact struct {
    IntAttribute       int64
    StringAttribute    string
    BooleanAttribute   bool
    FloatAttribute     float64
    TimeAttribute      time.Time
    WhatToSay          string
}

```

Seperti cara umum bahasa Golang, hanya atribut-atribut yang *terlihat* saja yang dapat di akses
dari dalam mesin Grule, hanya atribut dan fungsi yang berawalan huruf besar.

Grule juga dapa memanggil fungsi-fungsi dalam fakta.

```go
func (mf *MyFact) GetWhatToSay(sentence string) string {
    return fmt.Sprintf("Let say \"%s\"", sentence)
}
```

Mohon dicatat, ada beberapa persyaratan.

* Fungsi dalam fakta harus bisa **terlihat**, memiliki huruf besar di awal.
* Jika fungsinya ada balikan, hanya boleh ada 1 balikan saja.
* Dalam argumen dan balikan, jika itu adalah `int`, `uint` atau `float`, harus menggunakan varian 64-bit-nya. Seperti `int64`, `uint64`, `float64`.
* Fungsi yang dipangging tidak disarankan mengubah nilai data dari atribut-atribut dalam fakta, ini menyebabkan deteksi RETE pada *working memory* menjadi mustahil. Jika anda tetap **HARUS** merubah beberapa atribut, anda harus memberitahu Grule menggunakan fungsi `Changed(varname string)` yang sudah disiapkan.

## Menambahkan Fakta kedalam DataContext

Untui menambahkan fakta kedalam *DataContext* anda harus membuat sebuah *instance* dari `fakta` anda.

```go
myFact := &MyFact{
    IntAttribute: 123,
    StringAttribute: "Some string value",
    BooleanAttribute: true,
    FloatAttribute: 1.234,
    TimeAttribute: time.Now(),
}
```

Anda bisa membuat banyak fakta sesuai kebutuhan anda.

Berikutnya, anda hendak menyiapkan `DataContext` dan menambahkan `instance` fakta anda kedalamnya.

```go
dataCtx := ast.NewDataContext()
err := dataCtx.Add("MF", myFact)
if err != nil {
    panic(err)
}
```

## Membuat pustaka KnowledgeLibrary dan memasukan Rule kedalamnya.

Pada dasarnya, `KnowledgeLibrary` adalah sebuah pustaka berisi kumpulan cetak-biru `KnowledgeBase`.
Dan `KnowledgeBase` adalah kumpulan dari banyak *rule* yang berasal dari kumpulan definisi (GRL)
yang diambil dari banyak sumber.
Kita menggunakan `RuleBuilder` untuk membangun `KnowledgeBase` dan menambahkannya kedalam `KnowledgeLibrary`

Sebuah DRL, bisa dibuat dari sebuah string sederhana, disimpan didalam file atau tersimpan di suatu tempat di internet.
DRL ini digunakan untuk membangun satu atau lebih *rule*

Sekarang, mari kita buat `KnowledgeLibrary` dan `RuleBuilder` untuk membangun *rule* kedalam `KnowledgeLibrary` yang sudah disiapkan.

```go
knowledgeLibrary := ast.NewKnowledgeLibrary()
ruleBuilder := builder.NewRuleBuilder(knowledgeLibrary)
```

Sekaran kita bisa menambahkan sebuah *rule* (didefinisikan dalam GRL)

```go
// mari kita siapkan sebuah definisi rule
drls := `
rule CheckValues "Check the default values" salience 10 {
    when 
        MF.IntAttribute == 123 && MF.StringAttribute == "Some string value"
    then
        MF.WhatToSay = MF.GetWhatToSay("Hello Grule");
        Retract("CheckValues);
}
`

// Tambahkan definisi diatas kedalam pustakan dan kita namakan 'TutorialRules'  dengan versi '0.0.1'
byteArr := pkg.NewBytesResource([]byte(drls))
err := ruleBuilder.BuildRuleFromResource("TutorialRules", "0.0.1", byteArr)
if err != nil {
    panic(err)
}
```

### Sumber

Anda dapat memuat GRL dari berbagai dan banyak sumber.

#### Dari File

```go
fileRes := pkg.NewFileResource("/path/to/rules.grl")
err := ruleBuilder.BuildRuleFromResource("TutorialRules", "0.0.1", fileRes)
if err != nil {
    panic(err)
}
```

atau jika anda inginkan anda bisa memuat dari sumber file dengan pola

```go
bundle := pkg.NewFileResourceBundle("/path/to/grls", "/path/to/grls/**/*.grl")
resources := bundle.MustLoad()
for _, res := range resources {
    err := ruleBuilder.BuildRuleFromResource("TutorialRules", "0.0.1", res)
    if err != nil {
        panic(err)
    }
}
```

#### Dari String atau ByteArray

```go
byteArr := pkg.NewBytesResource([]byte(rules))
err := ruleBuilder.BuildRuleFromResource("TutorialRules", "0.0.1", byteArr)
if err != nil {
    panic(err)
}
```

#### Dari URL

```go
urlRes := pkg.NewUrlResource("http://host.com/path/to/rule.grl")
err := ruleBuilder.BuildRuleFromResource("TutorialRules", "0.0.1", urlRes)
if err != nil {
    panic(err)
}
```

#### Dari GIT

```go
bundle := pkg.NewGITResourceBundle("https://github.com/hyperjumptech/grule-rule-engine.git", "/**/*.grl")
resources := bundle.MustLoad()
for _, res := range resources {
    err := ruleBuilder.BuildRuleFromResource("TutorialRules", "0.0.1", res)
    if err != nil {
        panic(err)
    }
}
```

Sekarang, didalam `KnowledgeLibrary` kita memiliki sebuah `KnowledgeBase` dengan nama `TutorialRules` dengan versi `0.0.1`. Untuk menjalankan rule ini, anda harus membuat sebuah *instance* dari `KnowledgeBase` ini dan mengambilnya dari `KnowledgeLibrary`. Ini akan dijelaskan dalam sesi berikutnya.

## Menjalankan Mesin Rule Grule

Untuk menjalankan sebuah `KnowledgeBase`, kita perlu membuat sebuah *instance* dari `KnowledgeBase`, diambail dari `KnowledgeLibrary`

```go
knowledgeBase := knowledgeLibrary.NewKnowledgeBaseInstance("Tutorial", "0.0.1")
```

Setiap instance yang anda dapatkan dari `KnowledgeLibrary` adalah *tiruan* dari *cetak-biru* `KnowledgeBase`. Tiruan ini adalah *instance* yang berbeda dan membuat eksekusi terhadapnya bersifat *thread-safe*. Setiap tiruan *instance* ini juga membawa `WorkingMemory`-nya sendiri. Ini sangat berguna jika anda ingin menjalankan banyak eksekusi *multithread* dari rule-engine (contohnya dalam sebuah *web-server* yang melayani setiap permintaan menggunakan mesin rule).

Ini memberikan peningkatan performa yang cukup besar dikarenakan anda tidak perlu membangun ulang `KnowledgeBase` dari GRL setiap kali anda memulai thread baru. `KnowledgeLibrary` akan membuat tiruan dari struktur `AST` dalam cetak biru `KnowledgeBase` kedalam *instance* baru.

Ok, sekarang mari kita jalankan `KnowledgeBase` yang didapat dari `KnowlegeLibrary` dengan menggunakan `DataContext` yang sudah disiapkan.

```go
engine = engine.NewGruleEngine()
err = engine.Execute(dataCtx, knowledgeBase)
if err != nil {
    panic(err)
}
```

## Mengambil Hasil Eksekusi

Jika anda perhatikan, pada GRL diatas,

```go
rule CheckValues "Check the default values" salience 10 {
    when 
        MF.IntAttribute == 123 && MF.StringAttribute == "Some string value"
    then
        MF.WhatToSay = MF.GetWhatToSay("Hello Grule");
        Retract("CheckValues");
}
```

Rule ini mengubah atribut `MF.WhatToSay` dimana ini mengacu pada `WhatToSay` dalam struktur `MyFact`.
Jadi cukup mengakses variable tersebut untuk mengambil hasilnya.

```go
fmt.Println(myFact.WhatToSay)
// this should prints
// Lets Say "Hello Grule"

## Resources

GRLs can be stored in external files and there are many ways to obtain and load
the contents of those files.

### From File

```go
fileRes := pkg.NewFileResource("/path/to/rules.grl")
err := ruleBuilder.BuildRuleFromResource("TutorialRules", "0.0.1", fileRes)
if err != nil {
    panic(err)
}
```

You can also load multiple files into a bundle with paths and glob patterns:

```go
bundle := pkg.NewFileResourceBundle("/path/to/grls", "/path/to/grls/**/*.grl")
resources := bundle.MustLoad()
for _, res := range resources {
    err := ruleBuilder.BuildRuleFromResource("TutorialRules", "0.0.1", res)
    if err != nil {
        panic(err)
    }
}
```

### From String or ByteArray

```go
bs := pkg.NewBytesResource([]byte(rules))
err := ruleBuilder.BuildRuleFromResource("TutorialRules", "0.0.1", bs)
if err != nil {
    panic(err)
}
```

### From URL

```go
urlRes := pkg.NewUrlResource("http://host.com/path/to/rule.grl")
err := ruleBuilder.BuildRuleFromResource("TutorialRules", "0.0.1", urlRes)
if err != nil {
    panic(err)
}
```

#### With Headers

```go
headers := make(http.Header)
headers.Set("Authorization", "Basic YWxhZGRpbjpvcGVuc2VzYW1l")
urlRes := pkg.NewURLResourceWithHeaders("http://host.com/path/to/rule.grl", headers)
err := ruleBuilder.BuildRuleFromResource("TutorialRules", "0.0.1", urlRes)
if err != nil {
    panic(err)
}
```

### From GIT

```go
bundle := pkg.NewGITResourceBundle("https://github.com/hyperjumptech/grule-rule-engine.git", "/**/*.grl")
resources := bundle.MustLoad()
for _, res := range resources {
    err := ruleBuilder.BuildRuleFromResource("TutorialRules", "0.0.1", res)
    if err != nil {
        panic(err)
    }
}
```

### From JSON

You can now build rules from JSON! [Read how it works](GRL_JSON_id.md) 

## Compile GRL into GRB

If you want to have faster rule set loading performance (e.g. you have very
large rule sets and loading GRL is too slow), you can save your rule set
into GRB (Grules Rule Binary) file. [Read how to store and load GRB](Binary_Rule_File_id.md) 