# Slice
Slice adalah reference elemen array. Slice bisa dibuat, atau bisa juga dihasilkan dari manipulasi sebuah array ataupun slice lainnya. Karena slice merupakan data reference, menjadikan perubahan data di tiap elemen slice akan berdampak pada slice lain yang memiliki alamat memori yang sama.


```bash
var fruitsA = []string{"apple", "grape"}      // slice
var fruitsB = [2]string{"banana", "melon"}    // array
var fruitsC = [...]string{"papaya", "grape"}  // array

```


## Hubungan Slice Dengan Array & Operasi Slice
Kalau perbedannya hanya di penentuan alokasi pada saat inisialisasi, kenapa tidak menggunakan satu istilah saja? atau adakah perbedaan lainnya?

Sebenarnya slice dan array tidak bisa dibedakan karena merupakan sebuah kesatuan. Array adalah kumpulan nilai atau elemen, sedang slice adalah referensi tiap elemen tersebut.

Slice bisa dibentuk dari array yang sudah didefinisikan, caranya dengan memanfaatkan teknik 2 index untuk mengambil elemen-nya. Contoh bisa dilihat pada kode berikut.

```bash
    var fruits = []string{"apple", "grape", "banana", "melon"}
    var newFruits = fruits[0:2]

    fmt.Println(newFruits) // ["apple", "grape"]
```

Kode fruits[0:2] maksudnya adalah pengaksesan elemen dalam slice fruits yang dimulai dari indeks ke-0, hingga elemen sebelum indeks ke-2. Elemen yang memenuhi kriteria tersebut akan didapat, untuk kemudian disimpan pada variabel lain sebagai slice baru. Pada contoh di atas, newFruits adalah slice baru yang tercetak dari slice fruits, dengan isi 2 elemen, yaitu "apple" dan "grape".


