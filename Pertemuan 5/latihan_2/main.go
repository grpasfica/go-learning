package main

import "fmt"

// Hanya 1 function yang diperbolehkan
func ubahNilai(nama string) string {
	return nama + "Rizka " + "Pasfica" // mengubah nilai tanpa membuat variabel baru
}

func main() {
	var username string = "Gracia " // hanya 1 variabel yang dipakai

	fmt.Println("Nama Lengkap: ", ubahNilai(username))
}
