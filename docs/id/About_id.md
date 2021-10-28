[![Gopher Holds The Rules](https://github.com/hyperjumptech/grule-rule-engine/blob/master/gopher-grule.png?raw=true)](https://github.com/hyperjumptech/grule-rule-engine/blob/master/gopher-grule.png?raw=true)

---

:construction:
__THIS PAGE IS BEING TRANSLATED__
:construction:

:construction_worker: Contributors are invited. Please read [CONTRIBUTING](../../CONTRIBUTING.md) and [CONTRIBUTING TRANSLATION](../CONTRIBUTING_TRANSLATION.md) guidelines.

:vulcan_salute: Please remove this note once you're done translating.

---


[![About_cn](https://github.com/yammadev/flag-icons/blob/master/png/CN.png?raw=true)](../cn/About_cn.md)
[![About_de](https://github.com/yammadev/flag-icons/blob/master/png/DE.png?raw=true)](../de/About_de.md)
[![About_en](https://github.com/yammadev/flag-icons/blob/master/png/GB.png?raw=true)](../en/About_en.md)
[![About_id](https://github.com/yammadev/flag-icons/blob/master/png/ID.png?raw=true)](../id/About_id.md)
[![About_in](https://github.com/yammadev/flag-icons/blob/master/png/IN.png?raw=true)](../in/About_in.md)

[Tentang Grule](About_id.md) | [Tutorial](Tutorial_id.md) | [Rule Engine](RuleEngine_id.md) | [GRL](GRL_id.md) | [GRL JSON](GRL_JSON_id.md) | [Algoritma RETE](RETE_id.md) | [Fungsi-fungsi](Function_id.md) | [FAQ](FAQ_id.md) | [Benchmark](Benchmarking_id.md)

# Grule

```go
import "github.com/hyperjumptech/grule-rule-engine"
```

## Rule Engine untuk Go

**Grule** adalah sebuah pustaka perangkat lunak berupa *Rule Engine* untuk bahasa pemrograman Golang. Di-inspirasi dari JBOSS Drools yang terkenal, tapi dibuat dengan lebih sederhana.

Seperti halnya **Drools**, **Grule** memiliki *DSL*-nya sendiri yang perbandingannya seperti berikut.

DRL dari Drools seperti :

```drool
rule "SpeedUp"
    salience 10
    when
        $TestCar : TestCarClass( speedUp == true && speed < maxSpeed )
        $DistanceRecord : DistanceRecordClass()
    then
        $TestCar.setSpeed($TestCar.Speed + $TestCar.SpeedIncrement);
        update($TestCar);
        $DistanceRecord.setTotalDistance($DistanceRecord.getTotalDistance() + $TestCar.Speed)
        update($DistanceRecord)
end
```

Dan DRL dari Grule akan seperti :

```go
rule SpeedUp "When testcar is speeding up we keep increase the speed." salience 10  {
    when
        TestCar.SpeedUp == true && TestCar.Speed < TestCar.MaxSpeed
    then
        TestCar.Speed = TestCar.Speed + TestCar.SpeedIncrement;
        DistanceRecord.TotalDistance = DistanceRecord.TotalDistance + TestCar.Speed;
}
```

# Apa itu RuleEngine

Tidak ada penjelasan yang lebih baik dari sebuah artikel yang ditulis oleh Martin Fowler. Anda dapat membaca artikelnya disini ([RulesEngine by Martin Fowler](https://martinfowler.com/bliki/RulesEngine.html)).

Diambil dari situs **TutorialsPoint** (dengan sedikit modifikasi),

**Grule** Rule Engine adalah sebuah *Production Rule* System yang menggunakan pendekatan berbasis *rule* untuk membuat sebuah *System Pakar*. System Pakar menggunakan sebuah system berdasarkan *pengetahuan*  dimana sistem ini akan memproses sebua representasi *pengetahuan* dan menambahkannya kedalam kumpulan basis pengetahuan didalamnya. Basis pengetahuan ini dapat dipergunakan untuk membuat *reasoning*.

Sebuah sistem *Production Rule* adalah *Turing Complete* yang ber-fokus pada representasi *pengetahuan* untuk menggambarkan logika yang proporsional dan logika *first-order* secara lengkap, jelas dan deklaratif.

Otak dari sebuah sistem *Production Rules* adalah sebuah mesin *Inference* yang dapat terbentuk dari sejumlah besar *rule* dan *fakta*. Mesin *Inference* akan mencocokan fakta dan data terhadap sekumpul *rule* untuk menentukan tidakan apa yang akan dilakukan berikutnya.

*Production Rule* adalah sebuah struktur yang terdiri dari dua bagian yang menggunakan logika *first-order* untuk mekanisme penentuan didalam representasi *pengetahuan*. Sebuah *business rule engine* adalah perangkat lunak yang menjalan satu atau lebih *business rule* dalam sebuah lingkungan *production*.

*Rule Engine* mengizinkan anda untuk mendefinisikan **“Apa yang harus dilakukan”** dan bukan **“Bagaimana melakukannya.”**

## Apa itu Rule

*(juga diambil dari TutorialsPoint)*

Rules adalah sebuah pengetahuan yang ditulis dalam bentuk "Saat (when) sebuah kondisi terjadi, Maka (then) lakukan sesuatu"

```grule
When
   <Condition is true>
Then
   <Take desired Action>
```

Bagian terpenting dalam sebuah Rule adalah pada bagian **when** nya. Jika bagian **when** terpenuhi, maka pagian **then** akan dieksekusi.

```grule
rule  <rule_name> <rule_description>
   <attribute> <value> {
   when
      <conditions>

   then
      <actions>
}
```

## Keuntungan dari Rule Engine

### Pemrograman Deklaratif

*Rules* mempermudah untuk mengekspresikan sebuah solusi untuk permasalahan uyang sulit dan juga mendapakan verifikasinya. Berbeda dengan kode program, *Rule* ditulis menggunakan bahasa yang sederhana; Bisnis analis dapat dengan mudah membaca dan memverifikasi sekumpulan *Rule*

### Pemisahan antara Logika dan Data

Data berada didalam *Domain Object* sementara logika bisnis akan berada didalam sekumpulan *Rule*. Tergantung dari jenis proyeknya, pemisahan ini akan sangat menguntungkan.

### Sentralisasi pengetahuan

Dengan menggunakan *Rule*, anda membuat sebuah penyimpanan pengetahuan (*knowledge base*). Penyimpanan ini menjadi sumber kebenaran atas aturan-aturan bisnis. Idealnya, aturan ini sangat mudah dibaca dan menjadikannya sebagai dokumentasi tersendiri.

### Kemudahan adaptasi terhadap perubahan

Karea aturan bisnis ini sebenarnya diperlakukan sebagai data. Mengubah aturan untuk menyusaikan dinamika bisnis menjadi mudah. Tidak perlu membangun ulang kode program atau melakukan *deployment* sebagaimana proses pembangunan perangkat lunak biasanya, yang perlu anda lakukan hanya melepas sekumpulan aturan dan memasukannya kedalam  penyimpanan *Rule*.