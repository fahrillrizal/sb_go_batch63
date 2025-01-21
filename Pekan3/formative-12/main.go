package main

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
)
// Soal 1
type NilaiMahasiswa struct{
	Nama, MataKuliah, IndeksNilai string
	Nilai, ID uint
}

var nilaiNilaiMahasiswa = []NilaiMahasiswa{}
var nextID uint = 1

func getIndexNilai(nilai uint) string {
	switch {
	case nilai >= 80:
		return "A"
	case nilai < 80 && nilai >= 70:
		return "B"
	case nilai < 70 && nilai >= 60:
		return "C"
	case nilai < 60 && nilai >= 50:
		return "D"
	default:
		return "E"
	}
}

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok || username != "admin" || password != "admin123" {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

func postNilai(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Nama, Matkul string
		Nilai uint
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if input.Nilai > 100 || input.Nilai < 0{
		http.Error(w, "Nilai tidak valid", http.StatusBadRequest)
		return
	}

	index := getIndexNilai(input.Nilai)

	data := NilaiMahasiswa{
		Nama: input.Nama,
		MataKuliah: input.Matkul,
		Nilai: input.Nilai,
		IndeksNilai: index,
		ID: nextID,
	}

	nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, data)
	nextID++

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}
// Soal 2
func getNilai(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nilaiNilaiMahasiswa)
}

func main() {
	http.HandleFunc("/nilai", Auth(postNilai))// Post nilai

	http.HandleFunc("/nilai/all", getNilai)// Get nilai semua mahasiswa

	fmt.Println("Server berjalan di http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}