package main

import (
	"fmt"
	"errors"
	"sort"
	"sync"
	"time"
)

// Soal 1
func print(kalimat string, tahun int) {
	fmt.Printf("Kalimat: %s\nTahun: %d\n", kalimat, tahun)
}

//Soal 2
func kelilingSegitigaSamaSisi(sisi int, tampilKalimat bool) (string, error) {
	if sisi == 0 {
		if tampilKalimat { 
			return "", errors.New("masukkan sisi terlebih dahulu")
		} else {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered from panic:", r)
				}
			}()
			panic("Input sisi terlebih dahulu")
		}
	}

	keliling := 3 * sisi

	if tampilKalimat { 
		return fmt.Sprintf("Keliling segitiga sama sisi dengan sisi %d cm = %d cm", sisi, keliling), nil
	} else {
		return fmt.Sprintf("%d", keliling), nil
	}
}

// Soal 3
func tambahAngka(nilai int, angka *int) {
	*angka += nilai
}

func cetak(angka *int) {
	fmt.Printf("Total angka: %d\n", *angka)
}

// Soal 4
func tambahData(phones *[]string) {
	*phones = append(*phones, "Xiaomi", "Asus", "IPhone", "Samsung", "Oppo", "Realme", "Vivo")
}

//Soal 5
func tampilPhone(phones []string, wg *sync.WaitGroup) {
	defer wg.Done()
	sort.Strings(phones)

	for i, phone := range phones {
		fmt.Printf("%d. %s\n", i+1, phone)
		time.Sleep(1 * time.Second)
	}
}

//Soal 6
func getMovies(moviesChannel chan<- string, movies ...string) {
	for _, movie := range movies {
		moviesChannel <- movie
		time.Sleep(1 * time.Second)
	}
	close(moviesChannel)
}

func main() {
	// Soal 3
	fmt.Println("\nSoal 3:")
	angka := 1

	fmt.Printf("Total angka: %d\n", angka)
	
	defer cetak(&angka)
	
	tambahAngka(7, &angka)
	tambahAngka(6, &angka)
	tambahAngka(-1, &angka)
	tambahAngka(9, &angka)

	// Soal 2
	fmt.Println("\nSoal 2:")
	fmt.Println(kelilingSegitigaSamaSisi(4, true))
	fmt.Println(kelilingSegitigaSamaSisi(8, false))
	fmt.Println(kelilingSegitigaSamaSisi(0, true))
	fmt.Println(kelilingSegitigaSamaSisi(0, false))

	// Soal 4
	fmt.Println("\nSoal 4:")
	var phones = []string{}

	tambahData(&phones)

	sort.Strings(phones)

	for i, phone := range phones {
		fmt.Printf("%d. %s\n", i+1, phone)
		time.Sleep(1 * time.Second)
	}

	// Soal 5
	fmt.Println("\nSoal 5:")
	var phone2 = []string{"Xiaomi", "Asus", "IPhone", "Samsung", "Oppo", "Realme", "Vivo"}

	var wg sync.WaitGroup
	wg.Add(1)

	go tampilPhone(phone2, &wg)

	wg.Wait()

	// Soal 6
	fmt.Println("\nSoal 6:")
	var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}

	moviesChannel := make(chan string)

	go getMovies(moviesChannel, movies...)

	for value := range moviesChannel {
		fmt.Println(value)
	}

	// Soal 1
	fmt.Println("\nSoal 1:")
	defer print("Golang backend Development", 2021)
}