package main

import (
	"fmt"
	"math"	
)

// Soal 1
type segitigaSamaSisi struct{
	alas, tinggi int
}

type persegiPanjang struct{
	panjang, lebar int
}

type tabung struct{
	jariJari, tinggi float64
}

type balok struct{
	panjang, lebar, tinggi int
}

type hitungBangunDatar interface{
	luas() int
	keliling() int
}

type hitungBangunRuang interface{
	volume() float64
	luasPermukaan() float64
}

func (s segitigaSamaSisi) luas() int {
	return (s.alas * s.tinggi) / 2
}

func (s segitigaSamaSisi) keliling() int {
	return 3 * s.alas
}

func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
	return 2 * (p.panjang + p.lebar)
}

func (t tabung) volume() float64 {
	return math.Pi * math.Pow(t.jariJari, 2) * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
	return 2 * math.Pi * t.jariJari * (t.jariJari + t.tinggi)
}

func (b balok) volume() float64 {
	return float64(b.panjang * b.lebar * b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	return 2 * float64(b.panjang * b.lebar + b.panjang * b.tinggi + b.lebar * b.tinggi)
}

// Soal 2
type phone struct{
	name, brand string
	year int
	colors []string
}

type display interface {
	show() string
}

func (p phone) show() string {
	return fmt.Sprintf("name: %s\nbrand : %s\nyear : %d\ncolors : %v", p.name, p.brand, p.year, p.colors)
}

// Soal 3
func luasPersegi(sisi int, pesan bool) interface{} {
	if sisi == 0 {
		if pesan {
			return "Input sisi terlebih dahulu"
		} else {
			return nil
		}
	}

	luas := sisi * sisi

	if pesan {
		return fmt.Sprintf("Luas persegi dengan sisi %d cm adalah %d cm", sisi, luas)
	} else {
		return luas
	}
}

// Soal 4
func assertion() {
	var prefix interface{}= "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6,8}
	var kumpulanAngkaKedua interface{} = []int{12,14}

	kalimat := prefix.(string)
	angkaPertama := kumpulanAngkaPertama.([]int)
	angkaKedua := kumpulanAngkaKedua.([]int)

	total := 0
	for _, angka := range angkaPertama {
		total += angka
	}
	for _, angka := range angkaKedua {
		total += angka
	}

	fmt.Printf("%s%d + %d + %d + %d = %d\n", kalimat, angkaPertama[0], angkaPertama[1], angkaKedua[0], angkaKedua[1], total)
}

func main() {
	// Soal 1
	segitiga := segitigaSamaSisi{alas: 4, tinggi: 5}
	persegi := persegiPanjang{panjang: 6, lebar: 8}
	tabung := tabung{jariJari: 3.5, tinggi: 10}
	balok := balok{panjang: 6, lebar: 8, tinggi: 10}

	fmt.Println("Soal 1:")
	fmt.Printf("Luas Segitiga: %d\n", segitiga.luas())
	fmt.Printf("Keliling Segitiga: %d\n", segitiga.keliling())
	fmt.Printf("Luas Persegi Panjang: %d\n", persegi.luas())
	fmt.Printf("Keliling Persegi Panjang: %d\n", persegi.keliling())
	fmt.Printf("Volume Tabung: %.2f\n", tabung.volume())
	fmt.Printf("Luas Permukaan Tabung: %.2f\n", tabung.luasPermukaan())
	fmt.Printf("Volume Balok: %.2f\n", balok.volume())
	fmt.Printf("Luas Permukaan Balok: %.2f\n", balok.luasPermukaan())

	// Soal 2
	phone := phone{
		name: "Samsung Galaxy Note 20",
		brand: "Samsung",
		year: 2020,
		colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}

	fmt.Println("\nSoal 2:")
	fmt.Println(phone.show())

	// Soal 3
	fmt.Println("\nSoal 3:")
	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))

	fmt.Println("\nSoal 4:")
	assertion()
}