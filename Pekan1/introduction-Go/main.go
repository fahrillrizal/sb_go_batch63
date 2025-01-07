package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// soal 1
	var kata1 = "Bootcamp"
	var kata2 = "Digital"
	var kata3 = "Skill"
	var kata4 = "Sanbercode"
	var kata5 = "Golang"

	kalimatSoal1 := fmt.Sprintf("%s %s %s %s %s", kata1, kata2, kata3, kata4, kata5)
	fmt.Println("Soal 1:", kalimatSoal1)

	// soal 2
	halo := "Halo Dunia"
	halo = strings.Replace(halo, "Dunia", "Golang", -1)
	fmt.Println("Soal 2:", halo)

	// soal 3
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"

	outputSoal3 := fmt.Sprintf("%s %s %s %s",
		kataPertama,
		strings.Title(kataKedua),
		strings.ToUpper(kataKetiga[:len(kataKetiga)-1])+string(kataKetiga[len(kataKetiga)-1]),
		strings.ToUpper(kataKeempat),
	)
	fmt.Println("Soal 3:", outputSoal3)

	// soal 4
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	angka1, _ := strconv.Atoi(angkaPertama)
	angka2, _ := strconv.Atoi(angkaKedua)
	angka3, _ := strconv.Atoi(angkaKetiga)
	angka4, _ := strconv.Atoi(angkaKeempat)

	total := angka1 + angka2 + angka3 + angka4
	fmt.Println("Soal 4:", total)

	// soal 5
	kalimat := "halo halo bandung"
	angka := 2021

	kalimat = strings.Replace(kalimat, "halo", "Hi", -1)
	kalimat = strings.Replace(kalimat, "Hi bandung", "Hi Bandung", -1)

	fmt.Printf("Soal 5: \"%s\" - %d\n", kalimat, angka)
}