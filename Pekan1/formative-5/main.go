package main

import "fmt"

//soal 1
func luasPersegiPanjang(panjang int, lebar int) int {
	return panjang * lebar
}

func kelilingPersegiPanjang(panjang int, lebar int) int {
	return 2 * (panjang + lebar)
}

func volumeBalok(panjang int, lebar int, tinggi int) int {
	return panjang * lebar * tinggi
}

//soal 2
func introduce(nama string, gender string, pekerjaan string, usia int) (kalimat string) {
    if gender == "laki-laki" {
        if usia <= 22 {
            kalimat = fmt.Sprintf("Saudara %s adalah seorang %s yang berusia %d tahun", nama, pekerjaan, usia)
        } else if usia < 18 {
            kalimat = fmt.Sprintf("Kak %s adalah seorang %s yang berusia %d tahun", nama, pekerjaan, usia)
        } else {
            kalimat = fmt.Sprintf("Pak %s adalah seorang %s yang berusia %d tahun", nama, pekerjaan, usia)
        }
    } else if gender == "perempuan" {
        if usia <= 22 {
            kalimat = fmt.Sprintf("Saudari %s adalah seorang %s yang berusia %d tahun", nama, pekerjaan, usia)
        } else if usia < 18 {
            kalimat = fmt.Sprintf("Kak %s adalah seorang %s yang berusia %d tahun", nama, pekerjaan, usia)
        } else {
            kalimat = fmt.Sprintf("Bu %s adalah seorang %s yang berusia %d tahun", nama, pekerjaan, usia)
        }
    }
    return
}

//soal 3
func buahFavorit(nama string, buah ...string) string {
	buahFavorit := fmt.Sprintf("Halo nama saya %s dan buah favorit saya adalah ", nama)
	for i, b := range buah {
		if i == len(buah)-1 {
			buahFavorit += fmt.Sprintf("\"%s\"", b)
		} else {
			buahFavorit += fmt.Sprintf("\"%s\", ", b)
		}
	}
	return buahFavorit
}

//soal 4
var dataFilm = []map[string]string{}

func tambahFilm(title string, durasi string, genre string, tahun string) {
	dataFilm = append(dataFilm, map[string]string{
		"title": title,
		"durasi": durasi,
		"genre": genre,
		"tahun": tahun,
	})
}

func main() {
    //Soal 1
    panjang := 12
    lebar := 4
    tinggi := 8

    luas := luasPersegiPanjang(panjang, lebar)
    keliling := kelilingPersegiPanjang(panjang, lebar)
    volume := volumeBalok(panjang, lebar, tinggi)

    fmt.Println("Soal 1:")
    fmt.Println("Luas Persegi Panjang:", luas)
    fmt.Println("Keliling Persegi Panjang:", keliling)
    fmt.Println("Volume Balok:", volume)

    //Soal 2
    john := introduce("John", "laki-laki", "penulis", 30)
    sarah := introduce("Sarah", "perempuan", "model", 28)

    fmt.Println("\nSoal 2:")
    fmt.Println(john)
    fmt.Println(sarah)

    //Soal 3
    var buah = []string{"semangka", "jeruk", "melon", "pepaya"}
    var buahFavoritJohn = buahFavorit("John", buah...)

    fmt.Println("\nSoal 3:")
    fmt.Println(buahFavoritJohn)

    //Soal 4
    tambahFilm("LOTR", "2 jam", "action", "1999")
    tambahFilm("avenger", "2 jam", "action", "2019")
    tambahFilm("spiderman", "2 jam", "action", "2004")
    tambahFilm("juon", "2 jam", "horror", "2004")

    fmt.Println("\nSoal 4:")
    for _, item := range dataFilm {
        fmt.Printf("map[genre:%s durasi:%s tahun:%s title:%s]\n", item["genre"], item["durasi"], item["tahun"], item["title"])
    }
}