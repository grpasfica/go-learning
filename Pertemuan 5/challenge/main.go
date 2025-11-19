package main

import "fmt"

func hitungParkir(jam int, isMember bool, isHoliday bool) int {
	// Inisiasi total
	total := 0

	// 2 jam pertama = Rp 5000
	if jam <= 2 {
		total = 5000
	} else {
		// Jam ke-3 dan seterusnya Rp 2000/jam
		total = 5000 + (jam-2)*2000
	}

	// Jika member
	if isMember {
		if jam <= 5 {
			total = total - (total * 50 / 100) // diskon 50%
		} else {
			total = total - (total * 30 / 100) // diskon 30%
		}
	}

	// Jika hari libur
	if isHoliday {
		total = total + 3000
	}

	return total
}

func main() {
	// 4 jam, bukan member, bukan hari libur
	total1 := hitungParkir(4, false, false)
	fmt.Println("Total biaya (4 jam, non-member, bukan libur): Rp", total1)

	// 2 jam, member, hari libur
	total2 := hitungParkir(2, true, true)
	fmt.Println("Total biaya (2 jam, member, hari libur): Rp", total2)
}
