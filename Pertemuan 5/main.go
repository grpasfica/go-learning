package main

import (
	"fmt"
	"strconv" // modul untuk konversi tipe data
)

// Deklarasi variable menggunakan var
// Deklarasi menggunakan var cukup menampilkan var nama_variable dan tipe data, namun bisa juga dengan menampilkan nilai sekaligus
var var_name string = "Gracia"
var var_age int = 25

// Deklarasi variabel dengan constant (Jika menyimpan nilai dengan variable constant maka nilai awal tidak dapat diubah)
const Pi float64 = 3.1459
const Truth bool = true

// Digunakan untuk percobaan konversi tipe data (int dan float)
var a int = 42
var b int64 = int64(a)
var c int8 = int8(a)

var d float64 = float64(a)
var e float64 = 3.14
var f int = int(e)

var str string = "Hallo, Go"
var bytes []byte = []byte(str)
var str2 string = string(bytes)

// Warning : hati-hati dengan kemungkinan overflow
/*
	Overflow terjadi saat nilai dari satu tipe data mencoba diubah ke tipe data lain,
	tetapi nilai tersebut berada di luar rentang yang dapat ditampung oleh tipe data tujuan.
	Hal ini menyebabkan hasil yang tidak terduga atau salah,
	karena nilai yang melebihi batas akan "meluap" dan menyebabkan wrapping atau kehilangan data.
*/

func main() {
	fmt.Println("Hello, World!")

	fmt.Println("Ini contoh deklarasi variable menggunakan var")
	fmt.Println("Name : ", var_name)
	fmt.Println("Age : ", var_age)

	// Inisialisasi variable yang tidak menampilkan tipe data, namun tipe data diinisiasikan secara otomatis berdasarkan nilai
	fmt.Println("Ini contoh deklarasi variable dengan inisialisasi tipe otomatis")
	name_otomatis := "John"
	age_otomatis := 25
	fmt.Println("Name : ", name_otomatis)
	fmt.Println("Age : ", age_otomatis)

	// Memanggil variable dengan tipe data const
	fmt.Println(Pi)
	fmt.Println(Truth)
	// Contoh mengubah nilai pada variable constant
	// Pi = 4.5 // Akan terjadi error karena variable Pi menggunakan const

	// Konversi tipe data : Teknik mengubah satu tipe data ke tipe data lain dengan aman dalam Golang.
	fmt.Println(a, b, c)
	// Konversi int ke float dan float ke int
	fmt.Println(d, f)
	// Konversi str, bytes, str2
	fmt.Println(str, bytes, str2)

	// Konversi tipe data dengan function

	// String ke int
	var str string = "123"
	var num, err = strconv.Atoi(str)
	if err == nil {
		fmt.Println(num)
	}

	// Konversi dari int ke string
	var num2 int = 462
	var str2 string = strconv.Itoa(num2)
	fmt.Println(str2)

	// Konversi dari string ke float64
	var str3 string = "3.14"
	var f, err2 = strconv.ParseFloat(str3, 64)
	if err2 == nil {
		fmt.Println(f)
	}

	// Konversi dari float64 ke string
	var f2 float64 = 1.618
	var str4 string = strconv.FormatFloat(f2, 'f', 6, 64)
	fmt.Println(str4)
}
