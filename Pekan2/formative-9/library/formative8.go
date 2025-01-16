package library

import (
	"fmt"
	"math"
)

// Soal 1
type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

type Tabung struct {
	JariJari, Tinggi float64
}

type Balok struct {
	Panjang, Lebar, Tinggi int
}

type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

func (s SegitigaSamaSisi) Luas() int {
	return (s.Alas * s.Tinggi) / 2
}

func (s SegitigaSamaSisi) Keliling() int {
	return 3 * s.Alas
}

func (p PersegiPanjang) Luas() int {
	return p.Panjang * p.Lebar
}

func (p PersegiPanjang) Keliling() int {
	return 2 * (p.Panjang + p.Lebar)
}

func (t Tabung) Volume() float64 {
	return math.Pi * math.Pow(t.JariJari, 2) * t.Tinggi
}

func (t Tabung) LuasPermukaan() float64 {
	return 2 * math.Pi * t.JariJari * (t.JariJari + t.Tinggi)
}

func (b Balok) Volume() float64 {
	return float64(b.Panjang * b.Lebar * b.Tinggi)
}

func (b Balok) LuasPermukaan() float64 {
	return 2 * float64(b.Panjang*b.Lebar + b.Panjang*b.Tinggi + b.Lebar*b.Tinggi)
}

// Soal 2
type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

type Display interface {
	Show() string
}

func (p Phone) Show() string {
	return fmt.Sprintf("name: %s\nbrand : %s\nyear : %d\ncolors : %v", p.Name, p.Brand, p.Year, p.Colors)
}

// Soal 3
func LuasPersegi(sisi int, pesan bool) interface{} {
	if sisi == 0 {
		if pesan {
			return "Maaf anda belum menginput sisi dari persegi"
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
func Assertion() {
	var prefix interface{} = "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

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