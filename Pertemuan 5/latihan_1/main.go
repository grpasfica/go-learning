package main

import (
	"fmt"
)

// // Variable untuk konversi string ke byte
// var str string = "Aku adalah seorang programmer golang yang handal"
// var bytes []byte = []byte(str)

// var a int = 30
// var a_float float64 = float64(a)

// var c int = -45
// var c_float float64 = float64(c)

// // Hasil keuntungan
// var produk_a int = 100
// var harga_jual int = 150000
// var harga_beli int = 100000
// var biaya_ops int = 1000
// var diskon float64 = 0.15

func main() {

	// // Konversi tipe data dengan function
	// fmt.Println(str, bytes)

	// // Operasi aritmatika
	// var b float64 = 24.5
	// var d float64 = 0.67

	// var e float64 = (a_float + b) * c_float / d
	// fmt.Println("Nilai dari variable E : ", e)

	// Data
	// jumlah := 100
	// hargaJual := 150000.0
	// hargaBeli := 100000.0
	// biayaOperasional := 1000.0
	// diskon := 0.15

	// var jml_float float64 = float64(jumlah)

	// // Harga jual setelah diskon
	// hargaSetelahDiskon := hargaJual * (1 - diskon)

	// // Total pendapatan
	// totalPendapatan := hargaSetelahDiskon * jml_float

	// // Total biaya
	// totalBiaya := (hargaBeli * jml_float) + (biayaOperasional * jml_float)

	// // Total keuntungan
	// totalKeuntungan := totalPendapatan - totalBiaya

	// // Output
	// fmt.Printf("Harga Jual Setelah Diskon: Rp %.2f\n", hargaSetelahDiskon)
	// fmt.Printf("Total Pendapatan: Rp %.2f\n", totalPendapatan)
	// fmt.Printf("Total Biaya: Rp %.2f\n", totalBiaya)
	// fmt.Printf("Total Keuntungan: Rp %.2f\n", totalKeuntungan)

	angka := [...]int{5, 6, 7, 9, 15, 10, 3, 5, 6, 7, 8, 4, 8, 2, 1, 4, 9, 15, 27, 10, 1}
	indeks := angka[6:18]

	fmt.Println("Tampilan nilai dari index 6 sampai 18 ", indeks)
}
