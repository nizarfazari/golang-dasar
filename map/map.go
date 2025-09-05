package map_go

import "fmt"

/*
Map adalah tipe data asosiatif yang ada di Go yang berbentuk key-value pair. Data/value yang disimpan di map selalu disertai dengan key.
Key sendiri harus unik, karena digunakan sebagai penanda (atau identifier) untuk pengaksesan value yang disimpan di map.
*/
func Map() {
	var chicken map[string]int = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"]) // januari 50
	fmt.Println("mei", chicken["mei"])         // mei 0

	// Zero value dari map adalah nil. Disarankan untuk menginisialisasi secara explisit nilai awalnya agar tidak nil.

	var data map[string]int
	// data["one"] = 1
	// akan muncul error!

	data = map[string]int{}
	data["one"] = 1
	// tidak ada error

	// cara horizontal
	var chicken1 = map[string]int{"januari": 50, "februari": 40}

	// cara vertical
	var chicken2 = map[string]int{
		"januari":  50,
		"februari": 40,
	}

	fmt.Println(chicken1)
	fmt.Println(chicken2)

	for key, val := range chicken {
		fmt.Println(key, "  \t:", val)
	}
}

func MapCheckKey() {
	var chicken = map[string]int{"januari": 50, "februari": 40}
	var value, isExist = chicken["mei"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}

}

func MapDelete() {
	var chicken = map[string]int{"januari": 50, "februari": 40}

	fmt.Println(len(chicken)) // 2
	fmt.Println(chicken)

	delete(chicken, "januari")

	fmt.Println(len(chicken)) // 1
	fmt.Println(chicken)
}

func MapCombineSlice() {
	var chickens = []map[string]string{
		map[string]string{"name": "chicken blue", "gender": "male"},
		map[string]string{"name": "chicken red", "gender": "male"},
		map[string]string{"name": "chicken yellow", "gender": "female"},
	}

	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}
}
