package main

import (
	"fmt"
	"strings"
)

func main() {
	//soal 1
	fmt.Println("Soal 1:")
	for i := 1; i <= 20; i++ {
		if i % 3 == 0 && i % 2 != 0 {
			fmt.Printf("%d - I Love Coding \n", i)
		} else if i % 2 == 0 {
			fmt.Printf("%d - Berkualitas \n", i)
		} else {
			fmt.Printf("%d - Santai \n", i)
		}
	}

	//soal 2
	fmt.Println("\nSoal 2:")
	for i := 1; i <= 7; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}

	//soal 3
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	hasil := strings.Join(kalimat[2:], " ")

	fmt.Println("\nSoal 3:")
	fmt.Printf("[%s]\n", hasil)

	//soal 4
	var sayuran = []string{}
	sayuran = append(sayuran, "Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun")
	
	fmt.Println("\nSoal 4:")
	for i, sayur := range sayuran {
		fmt.Printf("%d. %s\n", i+1, sayur)
	}

	//soal 5
	var satuan = map[string] int {
		"panjang": 7,
		"lebar": 4,
		"tinggi": 6,
	}
	fmt.Println("\nSoal 5:")
	volume := 1
	for key, value := range satuan {
		fmt.Printf("%s = %d\n", key, value)
		volume *= value
	}

	fmt.Println("Volume Balok=", volume)
}