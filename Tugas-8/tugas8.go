package main

import (
	"fmt"	
	"math"
	"strconv"
)

type segitigaSamaSisi struct{
	alas, tinggi int
}

type persegiPanjang struct{
	panjang, lebar int
}

type tabung struct{
	jariJari, tinggi float64
}

type balok struct{
	panjang, lebar, tinggi int
}

type phone struct{
	name, brand string
	year int
	colors []string
}

type hitungBangunDatar interface{
	luas() int
	keliling() int
}

type hitungBangunRuang interface{
	volume() float64
	luasPermukaan() float64
}

type phoneOps interface{
	addPhone() string
}


func (s segitigaSamaSisi) luas() int {
	return s.alas / 2 * s.tinggi
}

func (s segitigaSamaSisi) keliling() int {
	return s.alas * 3
}

func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
	return 2 * (p.panjang + p.lebar)
}

func (b balok) volume() float64 {
	return float64(b.panjang) * float64(b.lebar) * float64(b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	pl := b.panjang * b.lebar
	pt := b.panjang * b.tinggi
	lt := b.lebar * b.tinggi

	return 2 * (float64(pl) + float64(pt) + float64(lt))
}

func (t tabung) volume() float64 {
	return 3.14 * math.Pow(t.jariJari, 2) * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
	return 2 * 3.14 * t.jariJari * (t.jariJari + t.tinggi)
}

func (ph phone) addPhone() string {
	name := ph.name
	brand := ph.brand
	year := strconv.Itoa(ph.year)
	ph.colors = append(ph.colors, "Mystic Bronze", "Mystic White", "Mystic Black")
	colorStr := ""
	for i := 0; i < len(ph.colors) - 1; i++{
		colorStr += ph.colors[i] + ", "
	}

	colorStr += ph.colors[len(ph.colors) - 1]

	return "name : " + name + "\nbrand : " + brand + "\nyear : " + year + "\ncolors : " + colorStr
}

func luasPersegi(num int, action bool) interface{} {
	var any interface{}
	if num > 0 && action == true{
		luas := num * num
		any = "Luas Persegi dengan sisi " + strconv.Itoa(num) + "cm adalah " + strconv.Itoa(luas) +"cm"
	} else if num > 0 && action == false{
		any = num
	} else if num == 0 && action == true {
		any = "Maaf anda belum menginput sisi dari persegi"
	} else {
		any = nil
	}

	return any

}

func main(){
	//soal 1
	fmt.Println("=======SOAL 1=========")
	var bangunDatar hitungBangunDatar
	var bangunRuang hitungBangunRuang
	bangunDatar = persegiPanjang{5, 9}
	fmt.Println("Luas Persegi Panjang :" , bangunDatar.luas())
	fmt.Println("Keliling Persegi Panjang :" , bangunDatar.keliling())

	bangunDatar = segitigaSamaSisi{4, 6}
	fmt.Println("===================")
	fmt.Println("Luas Segitiga Sama Sisi :" , bangunDatar.luas())
	fmt.Println("Keliling Segitiga Sama Sisi :" , bangunDatar.keliling())

	bangunRuang = balok{4, 6, 8}
	fmt.Println("===================")
	fmt.Println("Volume Balok :" , bangunRuang.volume())
	fmt.Println("Luas Permukaan Balok :" , bangunRuang.luasPermukaan())

	bangunRuang = tabung{3.0, 7.0}
	fmt.Println("===================")
	fmt.Println("Volume Tabung :" , bangunRuang.volume())
	fmt.Println("Luas Permukaan Tabung :" , bangunRuang.luasPermukaan())


	//soal 2
	fmt.Println("=======SOAL 2=========")
	var newPhone phoneOps
	newPhone = phone{name : "Samsung Galaxy Note 20", brand : "Samsung", year : 2020}

	fmt.Println(newPhone.addPhone())

	//soal 3
	fmt.Println("=======SOAL 3=========")
	fmt.Println(luasPersegi(4, true))

	fmt.Println(luasPersegi(8, false))

	fmt.Println(luasPersegi(0, true))

	fmt.Println(luasPersegi(0, false))

	//soal 4
	fmt.Println("=======SOAL 4=========")
	var prefix interface{}= "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6,8}
	var kumpulanAngkaKedua interface{} = []int{12,14}

	var pre string = prefix.(string)

	var number1 = kumpulanAngkaPertama.([]int)[0]
	var number2 = kumpulanAngkaPertama.([]int)[1]
	var number3 = kumpulanAngkaKedua.([]int)[0]
	var number4 = kumpulanAngkaKedua.([]int)[1]

	sum := number1 + number2 + number3 + number4

	finalSentence := pre + strconv.Itoa(number1) + " + " + strconv.Itoa(number2) + " + " + strconv.Itoa(number3) + " + " + strconv.Itoa(number4) + " = " + strconv.Itoa(sum)
	fmt.Println(finalSentence)
}