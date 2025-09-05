package slice

import "fmt"

/*
Slice adalah reference elemen array. Slice bisa dibuat, atau bisa juga dihasilkan dari manipulasi sebuah array ataupun slice lainnya.
Karena slice merupakan data reference, menjadikan perubahan data di tiap elemen slice akan berdampak pada slice lain
yang memiliki alamat memori yang sama.
*/
func Slice() {
	// var fruits = []string{"apple", "grape", "banana", "melon"}
	// fmt.Println(fruits[0])

	var fruits = []string{"apple", "grape", "banana", "melon"}
	var newFruits = fruits[0:2]

	fmt.Println(newFruits)
}

func SliceReference() {
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]

	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]

	fmt.Println(fruits)   // [apple grape banana melon]
	fmt.Println(aFruits)  // [apple grape banana]
	fmt.Println(bFruits)  // [grape banana melon]
	fmt.Println(aaFruits) // [grape]
	fmt.Println(baFruits) // [grape]

	// Buah "grape" diubah menjadi "pinnaple"
	baFruits[0] = "pinnaple"

	fmt.Println(fruits)   // [apple pinnaple banana melon]
	fmt.Println(aFruits)  // [apple pinnaple banana]
	fmt.Println(bFruits)  // [pinnaple banana melon]
	fmt.Println(aaFruits) // [pinnaple]
	fmt.Println(baFruits) // [pinnaple]

}

func ArrayLen() {
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(len(fruits)) // 4

}

func ArrayCap() {
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(len(fruits)) // len: 4
	fmt.Println(cap(fruits)) // cap: 4

	var aFruits = fruits[0:3]
	fmt.Println(len(aFruits)) // len: 3
	fmt.Println(cap(aFruits)) // cap: 4

	var bFruits = fruits[1:4]
	fmt.Println(len(bFruits)) // len: 3
	fmt.Println(cap(bFruits)) // cap: 3

}
