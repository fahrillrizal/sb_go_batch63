package main

import (
	"fmt"
	"sb_go_batch63/Pekan2/formative-9/library"
)

func main() {
	// Soal 1
	segitiga := library.SegitigaSamaSisi{Alas: 4, Tinggi: 5}
	persegi := library.PersegiPanjang{Panjang: 6, Lebar: 8}
	tabung := library.Tabung{JariJari: 3.5, Tinggi: 10}
	balok := library.Balok{Panjang: 6, Lebar: 8, Tinggi: 10}

	fmt.Println("Soal 1:")
	fmt.Printf("Luas Segitiga: %d\n", segitiga.Luas())
	fmt.Printf("Keliling Segitiga: %d\n", segitiga.Keliling())
	fmt.Printf("Luas Persegi Panjang: %d\n", persegi.Luas())
	fmt.Printf("Keliling Persegi Panjang: %d\n", persegi.Keliling())
	fmt.Printf("Volume Tabung: %.2f\n", tabung.Volume())
	fmt.Printf("Luas Permukaan Tabung: %.2f\n", tabung.LuasPermukaan())
	fmt.Printf("Volume Balok: %.2f\n", balok.Volume())
	fmt.Printf("Luas Permukaan Balok: %.2f\n", balok.LuasPermukaan())

	// Soal 2
	phone := library.Phone{
		Name:   "Samsung Galaxy Note 20",
		Brand:  "Samsung",
		Year:   2020,
		Colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}

	fmt.Println("\nSoal 2:")
	fmt.Println(phone.Show())

	// Soal 3
	fmt.Println("\nSoal 3:")
	fmt.Println(library.LuasPersegi(4, true))
	fmt.Println(library.LuasPersegi(8, false))
	fmt.Println(library.LuasPersegi(0, true))
	fmt.Println(library.LuasPersegi(0, false))

	// Soal 4
	fmt.Println("\nSoal 4:")
	library.Assertion()
}
