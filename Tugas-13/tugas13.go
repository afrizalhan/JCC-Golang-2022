package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type NilaiMahasiswa struct {
	ID          uint   `json:"id"`
	Nama        string `json:"nama"`
	MataKuliah  string `json:"mata_kuliah"`
	Nilai       uint   `json:"nilai"`
	IndeksNilai string `json:"indeks_nilai"`
}

var nilaiNilaiMahasiswa = []NilaiMahasiswa{}

func getIndeksNilai(nilai uint) string {
	switch {
	case nilai >= 80:
		return "A"
	case nilai >= 70 && nilai < 80:
		return "B"
	case nilai >= 60 && nilai < 70:
		return "C"
	case nilai >= 50 && nilai < 60:
		return "D"
	default:
		return "E"
	}
}

func getNilaiMahasiswa(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		dataNilaiMahasiswa, err := json.Marshal(nilaiNilaiMahasiswa)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(dataNilaiMahasiswa)
		return
	}
	http.Error(rw, "method not allowed", http.StatusMethodNotAllowed)
	return
}

func postNilaiMahasiswa(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		var nilaiMhs = NilaiMahasiswa{
			ID: uint(len(nilaiNilaiMahasiswa) + 1),
		}
		if r.Header.Get("Content-Type") == "application/json" {
			decodeJson := json.NewDecoder(r.Body)
			if err := decodeJson.Decode(&nilaiMhs); err != nil {
				log.Fatal(err)
			}
		} else {
			nama := r.FormValue("nama")
			mataKuliah := r.FormValue("mata_kuliah")
			getNilai := r.FormValue("nilai")
			nilai, _ := strconv.Atoi(getNilai)

			nilaiMhs.Nama = nama
			nilaiMhs.MataKuliah = mataKuliah
			nilaiMhs.Nilai = uint(nilai)
		}
		if nilaiMhs.Nilai > 100 {
			http.Error(rw, "Nilai tidak boleh diinput lebih dari 100", http.StatusBadRequest)
			return
		}

		nilaiMhs.IndeksNilai = getIndeksNilai(nilaiMhs.Nilai)

		dataNilai, _ := json.Marshal(nilaiMhs)

		nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, nilaiMhs)
		rw.Write(dataNilai)
		return
	}
	http.Error(rw, "method not allowed", http.StatusMethodNotAllowed)
	return
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()

		// cek username dan password kosong
		if !ok {
			rw.Write([]byte("username dan password tidak boleh kosong"))
			return
		}

		// cek username dan password valid
		if username == "admin" && password == "admin" {
			next.ServeHTTP(rw, r)
			return
		}

		rw.Write([]byte("username atau password salah"))
		return
	})
}

func main() {
	http.HandleFunc("/nilai-mahasiswa", getNilaiMahasiswa)
	http.Handle("/nilai-mahasiswa/create", Auth(http.HandlerFunc(postNilaiMahasiswa)))

	fmt.Println("Server Running at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
