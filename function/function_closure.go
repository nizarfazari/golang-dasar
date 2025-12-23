package function

import "fmt"

func FunctionCLosure() {
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	var numbers = []int{2, 3, 4, 3, 4, 2, 3}
	var min, max = getMinMax(numbers)
	fmt.Printf("data : %v\nmin  : %v\nmax  : %v\n", numbers, min, max)

	// Template %v digunakan untuk menampilkan data tanpa melihat tipe datanya. Jadi bisa digunakan untuk menampilkan data array, int, float, bool, dan lainnya.
	// Bisa dilihat di contoh statement, data bertipe array dan numerik ditampilkan menggunakan %v.
	fmt.Printf("data : %v\nmin  : %v\nmax  : %v\n", numbers, min, max)
}
