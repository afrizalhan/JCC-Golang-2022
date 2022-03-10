package main

import (
	"fmt"
	"strconv"
)

func main(){
	//soal 1 : convert type + hitung keliling dan luas
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	panjangPersegiInt, _ := strconv.ParseInt(panjangPersegiPanjang, 10, 64)
	lebarPersegiInt, _ := strconv.ParseInt(lebarPersegiPanjang, 10, 64)
	alasSegitigaInt, _ := strconv.ParseInt(alasSegitiga, 10, 64)
	tinggiSegitigaInt, _ := strconv.ParseInt(tinggiSegitiga, 10, 64)

	var kelilingPersegiPanjang int
	var luasSegitiga int

	kelilingPersegiPanjang = int(2 * (panjangPersegiInt + lebarPersegiInt))
	luasSegitiga = int(alasSegitigaInt / 2 * tinggiSegitigaInt)

	fmt.Printf("Keliling Persegi Panjang : %dcm \n", kelilingPersegiPanjang)
	fmt.Printf("Luas Segitiga : %dcm \n", luasSegitiga)

	//soal 2 : tentukan indeks nilai

	var nilaiJohn = 80
	var nilaiDoe = 50

	switch {
	case nilaiJohn >= 80:
		fmt.Println("Nilai John A")
	case (nilaiJohn >= 70 && nilaiJohn < 80) || (nilaiDoe >= 70 && nilaiDoe < 80):
		fmt.Println("Nilai John B")
	case (nilaiJohn >= 60 && nilaiJohn < 70) || (nilaiDoe >= 60 && nilaiDoe < 70):
		fmt.Println("Nilai John C")
	case (nilaiJohn >= 50 && nilaiJohn < 60) || (nilaiDoe >= 50 && nilaiDoe < 60):
		fmt.Println("Nilai John D")
	case nilaiJohn < 50 || nilaiDoe < 50:
		fmt.Println("Nilai John E")
	}

	switch {
	case nilaiDoe >= 80:
		fmt.Println("Nilai Doe A")
	case nilaiDoe >= 70 && nilaiDoe < 80:
		fmt.Println("Nilai Doe B")
	case nilaiDoe >= 60 && nilaiDoe < 70:
		fmt.Println("Nilai Doe C")
	case nilaiDoe >= 50 && nilaiDoe < 60:
		fmt.Println("Nilai Doe D")
	case nilaiDoe < 50:
		fmt.Println("Nilai Doe E")

	}

	//soal 3 : ganti ttl pribadi lalu switch case bulan
	var tanggal = 26;
	var bulan = 4; 
	var tahun = 2001;

	tanggalStr := strconv.Itoa(tanggal)
	tahunStr := strconv.Itoa(tahun)
	var bulanStr string

	switch bulan {
	case 1:
		bulanStr = "Januari"
	case 2:
		bulanStr = "Februari"
	case 3:
		bulanStr = "Maret"
	case 4:
		bulanStr = "April"
	case 5:
		bulanStr = "Mei"
	case 6:
		bulanStr = "Juni"
	case 7:
		bulanStr = "Juli"
	case 8:
		bulanStr = "Agustus"
	case 9:
		bulanStr = "September"
	case 10:
		bulanStr = "Oktober"
	case 11:
		bulanStr = "November"
	case 12:
		bulanStr = "Desember"
	default:
		fmt.Println("Bulan tidak ditemukan")
	}

	var tglLahir = tanggalStr + " " + bulanStr + " " + tahunStr

	fmt.Println(tglLahir)

	//soal 4 : conditional tahun lahir
	tahunLahir := 2001

	if 1944 <= tahunLahir && tahunLahir <= 1964{
		fmt.Println("Baby boomer")
	} else if 1965 <= tahunLahir && tahunLahir <= 1979{
		fmt.Println("Generasi X")
	} else if 1980 <= tahunLahir && tahunLahir <= 1994{
		fmt.Println("Generasi Y (Millenials)")
	} else if 1995 <= tahunLahir && tahunLahir <= 2015{
		fmt.Println("Generasi Z")
	} else {
		fmt.Println("Istilah tidak ditemukan")
	}

}