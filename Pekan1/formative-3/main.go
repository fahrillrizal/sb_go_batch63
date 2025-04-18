package main

import (
	"fmt"
	"strconv"
)

func main() {
	//soal 1
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"
	
	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"
	
	pPersegi, _ := strconv.Atoi(panjangPersegiPanjang)
	lPersegi, _ := strconv.Atoi(lebarPersegiPanjang)
	aSegitiga, _ := strconv.Atoi(alasSegitiga)
	tSegitiga, _ := strconv.Atoi(tinggiSegitiga)

	var lpp = pPersegi * lPersegi
	var kpp = 2 * (pPersegi + lPersegi)
	var ls = (aSegitiga * tSegitiga) / 2

	fmt.Println("soal 1:")
	fmt.Printf("Luas Persegi Panjang = %d x %d = %d\n", pPersegi, lPersegi, lpp)
	fmt.Printf("Keliling Persegi Panjang = 2 x (%d + %d) = %d\n", pPersegi, lPersegi, kpp)
	fmt.Printf("Luas Segitiga = 1/2 x (%d x %d) = %d\n", aSegitiga, tSegitiga, ls)

	//soal 2
	var nilaiJohn = 80
	var nilaiDoe = 50

	var indeksJohn, indeksDoe string

	if nilaiJohn >= 80 {
		indeksJohn = "A"
	} else if nilaiJohn >= 70 {
		indeksJohn = "B"
	} else if nilaiJohn >= 60 {
		indeksJohn = "C"
	} else if nilaiJohn >= 50 {
		indeksJohn = "D"
	} else {
		indeksJohn = "E"
	}

	if nilaiDoe >= 80 {
		indeksDoe = "A"
	} else if nilaiDoe >= 70 {
		indeksDoe = "B"
	} else if nilaiDoe >= 60 {
		indeksDoe = "C"
	} else if nilaiDoe >= 50 {
		indeksDoe = "D"
	} else {
		indeksDoe = "E"
	}

	fmt.Println("Soal 2:")
	fmt.Println("Indeks Nilai John:", indeksJohn)
	fmt.Println("Indeks Nilai Doe:", indeksDoe)

	//soal 3
	var tanggal = 14
	var bulan = 1
	var tahun = 2005

	namaBulan := ""
	switch bulan {
	case 1:
		namaBulan = "Januari"
	case 2:
		namaBulan = "Februari"
	case 3:
		namaBulan = "Maret"
	case 4:
		namaBulan ="April"
	case 5:
		namaBulan = "Mei"
	case 6:
		namaBulan = "Juni"
	case 7:
		namaBulan = "Juli"
	case 8:
		namaBulan = "Agustus"
	case 9:
		namaBulan = "September"
	case 10:
		namaBulan = "Oktober"
	case 11:
		namaBulan = "November"
	case 12:
		namaBulan = "Desember"
	default:
		namaBulan = "Bulan tidak valid"
	}

	fmt.Println("Soal 3:")
	fmt.Printf("%d %s %d\n", tanggal, namaBulan, tahun)

	//soal 4
	var tahunKelahiran = 2005

	generasi := ""
	if tahunKelahiran >= 1944 && tahunKelahiran <= 1964 {
		generasi = "Baby boomer"
	} else if tahunKelahiran >= 1965 && tahunKelahiran <= 1979 {
		generasi = "Gen X"
	} else if tahunKelahiran >= 1980 && tahunKelahiran <= 1994 {
		generasi = "Gen Y (Millenials)"
	} else if tahunKelahiran >= 1995 && tahunKelahiran <= 2015 {
		generasi = "Gen Z"
	} else {
		generasi = "Tidak termasuk dalam kategori generasi yang diketahui"
	}

	fmt.Println("Soal 4:")
	fmt.Println("Anda termasuk dalam generasi:", generasi)
}