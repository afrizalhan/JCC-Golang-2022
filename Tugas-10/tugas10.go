package main

import (
	"fmt"
	"strconv"
	"errors"
	"time"
	"Tugas-10/library"
	// "math"
	"flag"
)

func str(kalimat string, tahun int){
	fmt.Println("=======SOAL 1=========")
	fmt.Println(kalimat, tahun)
}

func endApp(){
	fmt.Println("End App")
	message := recover()
	fmt.Println(message, "Maaf anda belum menginput sisi dari segitiga sama sisi")
}

func runApp(error bool){
	defer endApp()
	if error{
		panic("ERROR")
	}
}

func kelilingSegitigaSamaSisi(alas int, action bool) interface{} {
	var any interface{}
	if alas > 0 && action == true{
		keliling := 3 * alas
		any = "Keliling segitiga sama sisi dengan sisi " + strconv.Itoa(alas) + "cm adalah " + strconv.Itoa(keliling) +"cm"
	} else if alas > 0 && action == false{
		any = alas
	} else if alas == 0 && action == true {
		any = errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
	} else {
		runApp(true)
	}

	return any
}

func tambahAngka(num int, sum *int){
	*sum += num
}

func cetakAngka(sum *int){
	fmt.Println("=======SOAL 3=========")
	fmt.Println("Jumlah Angka :", *sum) 
}



func main(){
	//soal 3
	angka := 1
	defer cetakAngka(&angka)
	//soal 1
	defer str("Golang Backend Development", 2021)

	//soal 2
	fmt.Println("=======SOAL 2=========")
	fmt.Println(kelilingSegitigaSamaSisi(4, true))

	fmt.Println(kelilingSegitigaSamaSisi(8, false))

	fmt.Println(kelilingSegitigaSamaSisi(0, true))

	fmt.Println(kelilingSegitigaSamaSisi(0, false))

	//soal 3
	tambahAngka(7, &angka)

	tambahAngka(6, &angka)

	tambahAngka(-1, &angka)

	tambahAngka(9, &angka)

	//soal 4
	fmt.Println("=======SOAL 4=========")
	var phones = []string{}

	library.TambahPhone(&phones, "Xiaomi", "Asus", "IPhone", "Samsung", "Oppo", "Realme", "Vivo")

	for index, item := range phones {
		fmt.Print(index + 1)
		fmt.Println(".", item)
		time.Sleep(time.Second)
	}

	//soal 5
	fmt.Println("=======SOAL 5=========")
	library.HitungLingkaran(7)
	library.HitungLingkaran(5)
	library.HitungLingkaran(10)

	//soal 6
	fmt.Println("=======SOAL 6=========")
	panjangflag := flag.Int("panjang", 6, "masukkan panjang")
	lebarflag := flag.Int("lebar", 8, "masukkan lebar")

	flag.Parse()
	panjang := *panjangflag
	lebar := *lebarflag

	library.HitungPersegiPanjang(panjang, lebar)
}