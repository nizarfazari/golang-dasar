package array

import "fmt"

/*
Array adalah kumpulan data bertipe sama, yang disimpan dalam sebuah variabel.
Array memiliki kapasitas yang nilainya ditentukan pada saat pembuatan,
menjadikan elemen/data yang disimpan di array tersebut jumlahnya tidak boleh melebihi yang sudah dialokasikan.
*/
func Array() {
	var names [4]string
	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"

	fmt.Println(names[0], names[1], names[2], names[3])
}

func ArrayInitFirst() {
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)
}

func ArrayPatter() {
	var fruits [4]string

	// cara horizontal
	fruits = [4]string{"apple", "grape", "banana", "melon"}

	// cara vertikal
	fruits = [4]string{
		"apple",
		"grape",
		"banana",
		"melon",
	}

	fmt.Print(fruits)
}

func ArrayMultiDimension() {
	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)
}

func ArrayForLoop() {
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	for i := 0; i < len(fruits); i++ {
		fmt.Printf("elemen %d : %s\n", i, fruits[i])
	}
}

func ArrayForRange() {
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	for i, fruit := range fruits {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}
}

func ArrayMake() {
	{
		var fruits = make([]string, 2)
		fruits[0] = "apple"
		fruits[1] = "manggo"

		fmt.Println(fruits) // [apple manggo]

	}
}
