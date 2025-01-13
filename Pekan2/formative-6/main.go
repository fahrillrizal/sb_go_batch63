package main

import (
	"fmt"
	"math"
)

// Soal 1
func hitungO(luasO *float64, kelilingO *float64, r float64){
	*luasO = math.Pi * math.Pow(r, 2)
	*kelilingO = math.Pi * r
}

// Soal 2
func introduce(sentence *string, nama string, gender string, pekerjaan string, usia int) {
    if gender == "laki-laki" {
        if usia <= 22 {
            *sentence = fmt.Sprintf("Saudara %s adalah seorang %s yang berusia %d tahun", nama, pekerjaan, usia)
        } else if usia < 18 {
            *sentence = fmt.Sprintf("Kak %s adalah seorang %s yang berusia %d tahun", nama, pekerjaan, usia)
        } else {
            *sentence = fmt.Sprintf("Pak %s adalah seorang %s yang berusia %d tahun", nama, pekerjaan, usia)
        }
    } else if gender == "perempuan" {
        if usia <= 22 {
            *sentence = fmt.Sprintf("Saudari %s adalah seorang %s yang berusia %d tahun", nama, pekerjaan, usia)
        } else if usia < 18 {
            *sentence = fmt.Sprintf("Kak %s adalah seorang %s yang berusia %d tahun", nama, pekerjaan, usia)
        } else {
            *sentence = fmt.Sprintf("Bu %s adalah seorang %s yang berusia %d tahun", nama, pekerjaan, usia)
        }
    }
}

// Soal 3
func tambahBuah(buah *[]string, namaBuah string) {
	*buah = append(*buah,namaBuah)
}

// Soal 4
func tambahFilm(title string, durasi string, genre string, tahun string, film *[]map[string]string) {
	*film = append(*film, map[string]string{
		"title": title,
		"duration": durasi,
		"genre": genre,
		"year": tahun,
	})
}

func main() {
	// Soal 1
	var luasO float64
	var kelilingO float64

	hitungO(&luasO, &kelilingO, 7)

	fmt.Println("Soal 1:")
	fmt.Printf("Luas Lingkaran: %.2f\n", luasO)
	fmt.Printf("Keliling Lingkaran: %.2f\n", kelilingO)

	// Soal 2
	var sentence string

	introduce(&sentence, "John", "laki-laki", "penulis", 30)
	fmt.Println("\nSoal 2:")
	fmt.Println(sentence)

	introduce(&sentence, "Sarah", "perempuan", "model", 18)
	fmt.Println(sentence) 

	// Soal 3
	var buah = []string{}

	tambahBuah(&buah, "Jeruk")
	tambahBuah(&buah, "Semangka")
	tambahBuah(&buah, "Mangga")
	tambahBuah(&buah, "Strawberry")
	tambahBuah(&buah, "Durian")
	tambahBuah(&buah, "Manggis")
	tambahBuah(&buah, "Alpukat")

	fmt.Println("\nSoal 3:")
	for i, b := range buah {
		fmt.Printf("%d. %s\n", i+1, b)
	}

	// Soal 4
	var film = []map[string]string{}

	tambahFilm("LOTR", "2 jam", "action", "1999", &film)
	tambahFilm("Avenger", "2 jam", "action", "2019", &film)
	tambahFilm("Spiderman", "2 jam", "action", "2004", &film)
	tambahFilm("Juon", "2 jam", "horror", "2004", &film)

	fmt.Println("\nSoal 4:")
	for i, f := range film{
		fmt.Printf("%d. title : %s\n", i+1, f["title"])
		fmt.Printf("   duration : %s\n", f["duration"])
		fmt.Printf("   genre : %s\n", f["genre"])
		fmt.Printf("   year : %s\n", f["year"])
	}
}