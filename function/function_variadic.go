package function

// function variadic
// Deklarasi parameter variadic sama dengan cara deklarasi variabel biasa, pembedanya adalah pada parameter jenis ini ditambahkan tanda titik tiga kali (...)
// tepat setelah penulisan variabel, sebelum tipe data. Nantinya semua nilai yang disisipkan sebagai parameter akan ditampung oleh variabel tersebut.
func CalculateVariadic(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))
	return avg
}
