# Pertanyaan yang sering ditanyakan

[![FAQ_cn](https://github.com/yammadev/flag-icons/blob/master/png/CN.png?raw=true)](../cn/FAQ_cn.md)
[![FAQ_de](https://github.com/yammadev/flag-icons/blob/master/png/DE.png?raw=true)](../de/FAQ_de.md)
[![FAQ_en](https://github.com/yammadev/flag-icons/blob/master/png/GB.png?raw=true)](../en/FAQ_en.md)
[![FAQ_id](https://github.com/yammadev/flag-icons/blob/master/png/ID.png?raw=true)](../id/FAQ_id.md)
[![FAQ_pl](https://github.com/yammadev/flag-icons/blob/master/png/PL.png?raw=true)](../pl/FAQ_pl.md)

[Tentang Grule](About_id.md) | [Tutorial](Tutorial_id.md) | [Rule Engine](RuleEngine_id.md) | [GRL](GRL_id.md) | [GRL JSON](GRL_JSON_id.md) | [Algoritma RETE](RETE_id.md) | [Fungsi-fungsi](Function_id.md) | [FAQ](FAQ_id.md) | [Benchmark](Benchmarking_id.md)

---

## 1. Grule Panik pada Siklus Maksimum

**Pertanyaan**: Saya mendapat pesan panik ini saat Grule engine dijalankan.

```Shell
panic: GruleEngine successfully selected rule candidate for execution after 5000 cycles, this could possibly caused by rule entry(s) that keep added into execution pool but when executed it does not change any data in context. Please evaluate your rule entries "When" and "Then" scope. You can adjust the maximum cycle using GruleEngine.MaxCycle variable.
```

**Jawaban**: Error ini mengindikasikan masalah yang ada pada __rule__ yang anda buat
dan dievaluasi oleh engine. Grule akan terus menjalankan __jaringan RETE__ didalam
__working memory__ hingga tidak ada lagi tindakan yang bisa dilakukan dalam __conflict set__,
yang mana disebut sebagai kondisi terminasi yang natural/normal. Jika dalam kumpulan rule tidak pernah
mengizinkan __jaringan RETE__ untuk mencapai kondisi akhir ini, maka eksekusi akan terus terjadi selamanya.
Secara __default__ konfigurasi untuk `GruleEngine.MaxCycle` adalah `5000`, dimana nilai ini untuk
melindungi eksekusi tidak berujung karena tidak pernah mencapai kondisi terminasi.

Anda dapat meningkatkan nilai ini jika menurut anda sistem __rule__ anda membutuhkan siklus lebih
banyak untuk bisa mencapai terminasi, tapi jika anda merasa ragu jika menambah nilai ini
akan menghentikan pesan panik, maka kemungkinan anda memiliki kumpulan __rule__ yang tidak
punya kondisi akhir.

Asumsikan __fact__ berikut ini:

```go
type Fact struct {
   Payment int
   Cashback int
}
```

Dan __rule-rule__ seperti berikut:

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

Kita akan menjalankan __rule__ tadi pada sebuah turunan fakta ...

```go
&Fact {
     Payment: 500,
}
```

... eksekusi ini tidak akan mencapai terminasi.

```
Siklus 1: Menjalankan "GiveCashback" .... karena F.Payment > 100 adalah kondisi yang valid
Siklus 2: Menjalankan "GiveCashback" .... karena F.Payment > 100 adalah kondisi yang valid
Siklus 3: Menjalankan "GiveCashback" .... karena F.Payment > 100 adalah kondisi yang valid
...
Siklus 5000: Menjalankan "GiveCashback" .... karena F.Payment > 100 adalah tetap kondisi yang valid
panik
```

Grule menjalankan __rule__ yang sama lagi dan lagi karena kondisi pada **WHEN**
terus menerus memberikan hasil yang valid.

Satu cara untuk memecahkan masalah ini adalah merubah __rule__ "GiveCashback" menjadi seperti:

```Shell
rule GiveCashback "Give cashback if payment is above 100" {
    When 
         F.Payment > 100 &&
         F.Cashback == 0
    Then
         F.Cashback = 10;
}
```

Dengan demikian, __rule__ `GiveCashback` turut memperhitungkan perubahan nilai yang terjadi.
Yang tadinya nilai variabel `Cashback` adalah 0, dikarenakan perubahan yang terjadi membuat
evaluasi ini menjadi tidak valid lagi pada siklus berikutnya, hingga menyebabkan evaluasi
pindah pada __rule__ yang lain hingga selesai.

Cara diatas adalah cara untuk mengotrol eksekusi __rule__ secara "natural" hingga setelah
serangkaian siklus engine akan berhenti secara normal karena tidak ada lagi __rule__ yang bisa
di eksekusi. Namun, ada kalanya anda tidak bisa menghetikan eksekusi dengan cara seperti ini.
Alternatif lain adalah untuk mengubah rule menjadi seperti berikut:

```Shell
rule GiveCashback "Give cashback if payment is above 100" {
    When 
         F.Payment > 100
    Then
         F.Cashback = 10;
         Retract("GiveCashback");
}
```

Fungsi `Retract` akan menghilangkan sementara __rule__ "GiveCashback" dari dalam __knowledge base__
hingga siklus berakhir. Karena __rule__ ini tidak lagi tersedia dalam siklus ini, maka __rule__
tersebut tidak dapat lagi dievaluasi hingga akhir. Perlu anda ketahui, bahwa __rule__ akan hilang
sementara saja setelah `Retract` dipanggil. Pada siklus-siklus setelahnya, rule tersebut akan tersedia
kembali.

---

## 2. Menyimpan Rule kedalam Database

**Pertanyaan**: Apakah ada rencana untuk mengintegrasikan Grule dengan penyimpanan di Database?

**Jawaban**: Tidak. Walaupun ini adalah ide yang baik untuk menyimpan __rule__ kedalam
database, Grule tidak akan membuat sebuah koneksi kepada sebuah database untuk secara otomatis menyimpan
dan mengambil __rule__. Anda dapat dengan mudah membuat mekanisme ini sendiri menggunakan
cara-cara yang sudah ada: menggunakan *Reader*, *File*, *Byte Array*, *String* dan *Git*.
Sebuah string dapat dengan mudah dimasukan dan baca dari database, untuk menyimpan/mengambil
__rule__ dan memasukannya kedalam __Knowledgebase__ dalam Grule.

Kami tidak ingin membuat keterikatan pada database apapun.

---

## 3. Jumlah maksimal rule dalam satu Knowledgebase

**Pertanyaan**: Berapa banyak __rule entry__ uang bisa dimasukan kedalam __knowledgebase__?

**Jawaban**: Anda dapat menambahkan berapapun __rule__ yang anda perlukan, selama minimal ada 1 rule dalam
sebuah __knowledgebase__

---

## 4. Mengetahui __rule__ apa saja yang valid untuk sebuah __fact__

**Pertanyaan**: Bagaimana saya mengetahui __rule__ - __rule__ mana saja yang valid terhadap sebuah __fact__?

**Jawaban**: Anda dapat menggunakan fungsi `engine.FetchMatchingRule`. Silahkan merujuk pada
[Matching Rules Doc](MatchingRules_id.md) untuk informasi lebih lengkap.

---

## 5. Use-Case untuk Rule Engine

**Pertanyaan**: Saya sudah membaca-baca tentang __rule engine__, tapi apa sebenarnya keuntungan yang didapat? Berikan kami contoh Use-Case.

**Jawaban**: Berikut ini adalah contoh situasi yang sebaiknya diselesaikan menggunakan solusi __rule-engine__ menurut hemat kami.

1. Sebuah sistem pakar yang harus mengevaluasi fakta-fakta guna memberikan sebuah kesimpulan yang nyata.
   Jika tidak menggunakan model RETE dan __rule-engine__, seorang developer akan membuat kode program
   yang berisi `if`/`else` yang beranak pinak dan permutasi terhadap kombinasi kondisi-kondisi yang ada
   membuat manajemen kode menjadi mustahil. Pendekatan __rule engine__ menggunakan tabel mungkin bisa
   memecahkan masalah, namun pendekatan ini menjadikan solusinya kaku dan tidak begitu mudah di
   buat kode program nya. Sistem seperti Grule ini memudahkan anda untuk mendeskripsikan peraturan terhadap
   data yang dipergunakan dalam sistem, dan melepaskan anda dari kebutuhan untuk mengimplementasi bagaimana
   sebenarnya evaluasi logika peraturan itu terlaksana, menyebunyikan banyak kompleksitas dari anda.

2. Sistem pemberian Rating atau Skor. Sebagai contoh, sebuah sistem perbankan ingin
   memberikan "skor" untuk setiap nasabah berdasarkan rekam jejak transaksi nasabah tersebut (fakta).
   Kita dapat melihat bagaimana skor nasabah berubah mengikuti seberapa sering mereka berinteraksi
   dengan bank, berapa bayak dana yang keluar dan masuk kedalam akun nasabah, seberapa cepat dan rajin
   seorang nasabah membayar tagihan hutang, total pendapatan nasabah dari bunga bank, dan seterusnya.
   __Rule engine__ di siapkan oleh teknisi IT dan spesifikasi __rule__ dan data disediakan langsung
   oleh mereka yang lebih mengerti mengenai sistem finansial dan analis keuangan para nasabah.
   Dengan demikian, menempatkan keahlian dan disiplin ilmu pada orang yang tepat.

3. Permainan Komputer (games). Status pemain, penghargaan, penalti, penilaian, kerusakan (damage)
   dan penghitungan probabilitas adalah beberapa dari banyak contoh dimana sebuah sistem __rule__
   sangat berperan penting pada hampir semua permainan komputer. __Rule-rule__ ini dapat menentukan
   interaksi dengan mekanisme permainan dengan cara yang sangat rumit, bahkan terlalu rumit sampai diluar
   imajinasi sang pembuatnya. Membuat peraturan-peraturan dalam permainan yang dinamis bisa saja di
   lakuan pada pemrograman skrip seperti LUA, namun logika bisa menjadi sangat rumit dan kompleks,
   dan dengan menggunakan sebuah __rule-engine__ dapat menurunkan kompleksitas cukup besar.

4. Sistem klasifikasi. Ini sebenarnya suatu bentuk umum dari sistem rating yang sudah dijelaskan sebelumnya.
   Dengan menggunakan __rule-engine__, kita bisa melakukan kalsifikasi terhadap hak tanggungan kredit,
   identifikasi kimia biologi, kategori resiko atas produk-produk asuransi, potensi resiko keamanan, dan
   banyak lagi.

5. Sistem pemberian saran. Sebuah "rule" sebenarnya adalah suatu bentuk data, dimana
   sebagai data, ia sendiri bisa merupakan hasil dari program yang lain. Program tersebut bisa jadi sebuah
   sistem pakar atau kecerdasan tiruan. __Rule__ bisa dibuat dan dimanipulasi oleh program lain
   agar secara dinamis mengikuti kondisi-kondisi perubahan fakta yang bersifat dinamis.

Ada sangat banyak contoh __use-case__ yang lain, yang akan mendapat keuntungan dari penggunaan
sebuah __Rule-Engine__. Contoh-contoh diatas hanya menunjukan sedikit potensi yang bisa didapat.

Walaupun demikian, perlu disebutkan bahwa __Rule-Engine__ tentu saja bukan jaminan untuk dapat
menyelesaikan semua masalah komputasi. Banyak alternatif lain yang memberikan solusi pada seputar
problem basis pengetahuan "knowledge base" dalam sebuah perangkat lunak, dan solusi-solusi tersebut
sebaiknya dipergunakan apabila lebih pantas. Contohnya, Seseorang tidak perlu menggunakan __rule-engine__
untuk sebuah masalah sederhana yang bisa dipecahkan dengan sebuah `if` dan `else` saja.

Ada hal lain yang menjadi catatan: Beberapa implementasi __rule engine__ adalah merupakan produk yang
sangat mahal harganya untuk dibeli atau disewa. Walaupun demikian, banyak bisnis yang mendapatkan
keuntungan berarti dari menggunakan produk-produk tersebut, dimana dengan menggunakannya, biaya yang
timbul dengan menggunakan produk-produk tersebut dengan mudah ditutupi dari keuntungan bisnis yang didapat.
Salah satu keuntungan bisnis yang sangat jelas, dimana penggunaan __Rule-Engine__ yang bisa
memutus keterikatan antara developer dan bisnis __user__ mempercepat pembangunan solusi dan
melunakkan kompleksitas bisnis itu sendiri.

---

## 6. Logging

**Pertanyaan**: Log yang dihasilakn oleh Grule terlalu banyak dan agak mengganggu. Bagaimana cara mengurangi / mematikan log ini?

**Jawaban**: Ya. Anda dapat mengurangi (atau bahkan menghilangkan) log yang dihasilan oleh Grule dengan cara merubah peringkat LOG nya.

```go
import (
    "github.com/DataWiseHQ/grule-rule-engine/logger"
    "github.com/sirupsen/logrus"
)
...
...
logger.SetLogLevel(logrus.PanicLevel)
```

Cara ini akan membuat Grule hanya mengeluarkan Log apabila iya panik.

Tentu saja, mengubah peringkat log ini mengurangi kemampuan anda untuk melakukan debugging,
karenanya, kami sarankan agar anda meningkatkan peringat log seperti ini hanya pada sistem
produksi saja (production environment)