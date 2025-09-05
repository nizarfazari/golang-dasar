package Variable

import "fmt"

func Variable() {

	// var <nama-variabel> <tipe-data> = <nilai>
	var firstName string = "nizar"
	var lastName string = "fazari"

	fmt.Printf("Halo %s %s \n", firstName, lastName)

	// type inference
	// metode deklarasi variabel yang tipe data-nya diketahui secara otomatis dari data/nilai

	var name string = "john"

	fmt.Printf("halo %s !\n", name)

	/*

		Deklarasi Variabel Menggunakan Keyword make
		Fungsi make() ini hanya bisa digunakan untuk pembuatan beberapa jenis variabel saja, yaitu:
		- channel
		- slice
		- map

	*/
}
