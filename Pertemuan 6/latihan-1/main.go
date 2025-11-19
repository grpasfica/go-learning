package main

import "fmt"

type Team struct {
	Name  string
	Score int
}

func (t *Team) UpdateScore(newScore int) {
	fmt.Printf("Mencoba memperbarui skor menjadi: %d\n", newScore)

	if newScore > t.Score {
		t.Score = newScore
		fmt.Println("Skor diperbarui.")
	} else {
		fmt.Println("Skor baru tidak lebih besar dari skor saat ini.")
	}
}

func main() {
	// Inisialisasi Data Awal
	myTeam := Team{
		Name:  "Garuda",
		Score: 10,
	}

	// Skenario A: Mencoba update dengan skor lebih KECIL (Seharusnya Gagal)
	myTeam.UpdateScore(5)
	fmt.Println("Output Score Terupdate: ", myTeam.Score)

	// Skenario B: Mencoba update dengan skor lebih BESAR (Seharusnya Berhasil)
	myTeam.UpdateScore(20)
	fmt.Println("Output Score Terupdate: ", myTeam.Score)
}
