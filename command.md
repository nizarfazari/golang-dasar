# Command

## Go Mod Init

`go mod init` digunakan untuk inisialisasi project Go yang menggunakan **Go Modules**.  

> Nama module bisa apa saja, tetapi biasanya disamakan dengan nama direktori/folder project.  
> Nama ini akan mempengaruhi **import path** sub-packages dalam project.

## Langkah-langkah

```bash
# 1. Buat direktori project
mkdir <nama-project>

# 2. Masuk ke direktori project
cd <nama-project>

# 3. Inisialisasi module Go
go mod init <nama-project>

```

## Go Run

`go run` digunakan untuk eksekusi file program, yaitu file yang ber-ekstensi .go. Cara penggunaannya dengan menuliskan command tersebut diikuti argumen nama file.

```bash
cd project-pertama
go run main.go
```

## Go Test
Go menyediakan package testing, berguna untuk keperluan pembuatan file test. Pada penerapannya, ada aturan yang wajib diikuti yaitu nama file test harus berakhiran `_test.go`.

Berikut adalah contoh penggunaan command go test untuk testing file main_test.go.
```bash 
go test main_test.go
```
## Go Build

`go build` digunakan untuk mengkompilasi file program menjadi file **executable (binary)**.  

### Perbedaan dengan `go run`
- `go run` → melakukan kompilasi **sementara** di folder **temporary**, lalu langsung mengeksekusi program.  
- `go build` → menghasilkan file **executable permanen** di folder aktif.  

### Contoh
Jika project bernama `project-pertama`, maka setelah di-build akan muncul file baru bernama:  

- **Windows**: `project-pertama.exe`  
- **Linux/MacOS**: `project-pertama`  

File executable tersebut kemudian bisa langsung dijalankan.

### Mengubah Nama File Binary
Secara default, nama binary mengikuti nama project.  
Untuk mengubah nama file, gunakan flag `-o`:

```bash
go build -o <nama-executable>
```

## Go Get
`go get` digunakan untuk men-download package atau dependency. Sebagai contoh, penulis ingin men-download package Kafka driver untuk Go pada project project-pertama, maka command-nya kurang lebih seperti berikut:


```bash
cd project-pertama
go get github.com/segmentio/kafka-go
```

Untuk mengunduh package/dependency versi terbaru, gunakan flag -u pada command go get, contohnya:

```bash
go get -u github.com/segmentio/kafka-go
```
