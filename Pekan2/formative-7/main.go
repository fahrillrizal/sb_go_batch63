package main

import "fmt"

// Soal 1
type Buah struct {
	Nama string
	Warna string
	adaBijinya bool
	Harga int
}

// Soal 2
type segitiga struct{
	alas, tinggi int
}

type persegi struct{
	sisi int
}

type persegiPanjang struct{
	panjang, lebar int
}

func (s segitiga) Luas() float64 {
	return 0.5 *float64(s.alas) * float64(s.tinggi)
}

func (p persegi) Luas() int {
	return p.sisi * p.sisi
}

func (pp persegiPanjang) Luas() int {
	return pp.panjang * pp.lebar
}

// Soal 3
type phone struct{
	name, brand string
	year int
	colors []string
}

func (p *phone) TambahWarna(warna string){
	p.colors = append(p.colors, warna)
}

type Film struct{
	title string
	Genre string
	durasi int
	year int
}

func tambahFilm(judul string, durasi int, genre string, tahun int, dataFilm *[]Film) {
	film := Film{title: judul, Genre: genre, durasi: durasi, year: tahun}
	*dataFilm = append(*dataFilm, film)
}

func main() {
	// Soal 1
	nanas := Buah{Nama: "Nanas", Warna: "Kuning", adaBijinya: false, Harga: 9000}
	jeruk := Buah{Nama: "Jeruk", Warna: "Oranye", adaBijinya: true, Harga: 8000}
	semangka := Buah{Nama: "Semangka", Warna: "Hijau & Merah", adaBijinya: true, Harga: 5000}
	pisang := Buah{Nama: "Pisang", Warna: "Kuning", adaBijinya: false, Harga: 5000}

	fmt.Println("Soal 1:")
	fmt.Printf("%+v\n", nanas)
	fmt.Printf("%+v\n", jeruk)
	fmt.Printf("%+v\n", semangka)
	fmt.Printf("%+v\n", pisang)

	// Soal 2
	segitiga := segitiga{alas: 4, tinggi: 5}
	persegi := persegi{sisi: 6}
	persegiPanjang := persegiPanjang{panjang: 7, lebar: 8}

	fmt.Println("\nSoal 2:")
	fmt.Printf("Luas Segitiga: %.2f\n", segitiga.Luas())
	fmt.Printf("Luas Persegi: %d\n", persegi.Luas())
	fmt.Printf("Luas Persegi Panjang: %d\n", persegiPanjang.Luas())

	// Soal 3
	phone := phone{name: "Galaxy S21", brand: "Samsung", year: 2021}
	phone.TambahWarna("Hitam")
	phone.TambahWarna("Putih")

	fmt.Println("\nSoal 3:")
	fmt.Printf("%+v\n", phone)

	// Soal 4
	var dataFilm []Film

	tambahFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahFilm("juon", 120, "horror", 2004, &dataFilm)

	fmt.Println("\nSoal 4:")
	for i, film := range dataFilm {
		fmt.Printf("%d. title : %s\n", i+1, film.title)
		fmt.Printf("   duration : %d jam\n", film.durasi/60)
		fmt.Printf("   genre : %s\n", film.Genre)
		fmt.Printf("   year : %d\n", film.year)
	}
	
}