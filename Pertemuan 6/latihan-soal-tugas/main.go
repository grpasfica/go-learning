package main

import (
	"fmt"
	"math" // library untuk mendefinisikan nilai Pi
)

// Interface
type Shape interface {
	Area() float64
	Perimeter() float64
	Name() string
}

// Struct square
type Square struct {
	Sisi float64
}

func (s Square) Area() float64      { return s.Sisi * s.Sisi }
func (s Square) Perimeter() float64 { return 4 * s.Sisi }
func (s Square) Name() string       { return "Persegi" }

// Struct rectangle
type Rectangle struct {
	Panjang float64
	Lebar   float64
}

func (r Rectangle) Area() float64      { return r.Panjang * r.Lebar }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Panjang + r.Lebar) }
func (r Rectangle) Name() string       { return "Persegi Panjang" }

// struct triangle
type Triangle struct {
	Alas   float64
	Tinggi float64
	Miring float64
}

func (t Triangle) Area() float64      { return 0.5 * t.Alas * t.Tinggi }
func (t Triangle) Perimeter() float64 { return t.Alas + t.Tinggi + t.Miring }
func (t Triangle) Name() string       { return "Segitiga" }

// struct lingkaran
type Circle struct {
	JariJari float64
}

func (c Circle) Area() float64      { return math.Pi * c.JariJari * c.JariJari }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.JariJari }
func (c Circle) Name() string       { return "Lingkaran" }

// struct trapesium
type Trapezoid struct {
	Atas   float64
	Bawah  float64
	Tinggi float64
	Miring float64
}

func (tr Trapezoid) Area() float64      { return 0.5 * (tr.Atas + tr.Bawah) * tr.Tinggi }
func (tr Trapezoid) Perimeter() float64 { return tr.Atas + tr.Bawah + (2 * tr.Miring) }
func (tr Trapezoid) Name() string       { return "Trapesium" }

// Processing
func ProcessShape(sh Shape) {
	fmt.Println("--- Hasil Perhitungan ---")
	fmt.Println("Bangun:", sh.Name())

	// %.2f artinya: menampilkan angka desimal dengan 2 angka di belakang koma
	fmt.Printf("Luas    : %.2f\n", sh.Area())
	fmt.Printf("Keliling: %.2f\n", sh.Perimeter())
	fmt.Println()
}

func main() {
	// Inisiasi variable
	kotak := Square{Sisi: 5}
	persegiPanjang := Rectangle{Panjang: 10, Lebar: 5}
	segitiga := Triangle{Alas: 3, Tinggi: 4, Miring: 5}
	lingkaran := Circle{JariJari: 7}
	trapesium := Trapezoid{Atas: 6, Bawah: 10, Tinggi: 4, Miring: 5}

	// call function
	ProcessShape(kotak)
	ProcessShape(persegiPanjang)
	ProcessShape(segitiga)
	ProcessShape(lingkaran)
	ProcessShape(trapesium)
}
