package main

import (
	"fmt"
	"strconv"
	"Tugas-9/bangun-datar"
	"Tugas-9/bangun-ruang"
	"Tugas-9/persegi"
	"Tugas-9/phone"
	"Tugas-9/soal-4"
)

func main(){
	//soal 1
	fmt.Println("=======SOAL 1=========")
	var bangunDatar bangundatar.HitungBangunDatar
	var bangunRuang bangunruang.HitungBangunRuang
	bangunDatar = bangundatar.PersegiPanjang{5, 9}
	fmt.Println("Luas Persegi Panjang :" , bangunDatar.Luas())
	fmt.Println("Keliling Persegi Panjang :" , bangunDatar.Keliling())

	bangunDatar = bangundatar.SegitigaSamaSisi{4, 6}
	fmt.Println("===================")
	fmt.Println("Luas Segitiga Sama Sisi :" , bangunDatar.Luas())
	fmt.Println("Keliling Segitiga Sama Sisi :" , bangunDatar.Keliling())

	bangunRuang = bangunruang.Balok{4, 6, 8}
	fmt.Println("===================")
	fmt.Println("Volume Balok :" , bangunRuang.Volume())
	fmt.Println("Luas Permukaan Balok :" , bangunRuang.LuasPermukaan())

	bangunRuang = bangunruang.Tabung{3.0, 7.0}
	fmt.Println("===================")
	fmt.Println("Volume Tabung :" , bangunRuang.Volume())
	fmt.Println("Luas Permukaan Tabung :" , bangunRuang.LuasPermukaan())


	//soal 2
	fmt.Println("=======SOAL 2=========")
	var newPhone phone.PhoneOps
	newPhone = phone.Phone{Name : "Samsung Galaxy Note 20", Brand : "Samsung", Year : 2020, Colors: []string{"Mystic Black", "Mystic White", "Mystic Black"}}

	fmt.Println(newPhone.AddPhone())

	//soal 3
	fmt.Println("=======SOAL 3=========")
	fmt.Println(persegi.LuasPersegi(4, true))

	fmt.Println(persegi.LuasPersegi(8, false))

	fmt.Println(persegi.LuasPersegi(0, true))

	fmt.Println(persegi.LuasPersegi(0, false))

	//soal 4
	fmt.Println("=======SOAL 4=========")
	
	var pre string = soal4.Prefix.(string)

	var number1 = soal4.KumpulanAngkaPertama.([]int)[0]
	var number2 = soal4.KumpulanAngkaPertama.([]int)[1]
	var number3 = soal4.KumpulanAngkaKedua.([]int)[0]
	var number4 = soal4.KumpulanAngkaKedua.([]int)[1]

	sum := number1 + number2 + number3 + number4

	finalSentence := pre + strconv.Itoa(number1) + " + " + strconv.Itoa(number2) + " + " + strconv.Itoa(number3) + " + " + strconv.Itoa(number4) + " = " + strconv.Itoa(sum)
	fmt.Println(finalSentence)

}