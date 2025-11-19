package main

import "fmt"

// Struct
type Buku struct {
	Judul       string
	Penulis     string
	TahunTerbit int
}

// Fungsi untuk memfilter buku berdasarkan tahun
func filterBuku(semuaBuku []Buku, tahunMinimal int) []Buku {
	var hasil []Buku

	for i := range semuaBuku {
		b := semuaBuku[i]

		if b.TahunTerbit >= tahunMinimal {
			hasil = append(hasil, b)
		}
	}
	return hasil
}

func main() {
	// inisiasi dengan menggunakan struct Buku
	perpustakaan := []Buku{
		{Judul: "Belajar Go", Penulis: "Budi", TahunTerbit: 2023},
		{Judul: "Sejarah Dunia", Penulis: "Siti", TahunTerbit: 1995},
		{Judul: "Resep Masakan", Penulis: "Rina", TahunTerbit: 2010},
		{Judul: "Teknologi AI", Penulis: "Joko", TahunTerbit: 2024},
		{Judul: "Novel", Penulis: "Andi", TahunTerbit: 2005},
	}

	fmt.Println("Daftar Buku (Tahun 2010 ke atas): ")

	// Memanggil function filterBuku dengan parameter semuaBuku = perpustakaan, tahunMinimal = 2010
	listBuku := filterBuku(perpustakaan, 2010)

	// Menampilkan list buku hasil filter
	for i := range listBuku {
		b := listBuku[i]
		// cetak data
		fmt.Println("Judul: ", b.Judul, ", Tahun: ", b.TahunTerbit, ", Penulis: ", b.Penulis)
	}
}
