package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main(){
	
	//soal 1
	var jabar = "Jabar "
	var coding = "Coding "
	var camp = "Camp "
	var golang = "Golang "
	var year = "2022"

	var sentence = jabar + coding + camp + golang + year

	fmt.Println(sentence)

	//soal 2
	halo := "Halo Dunia"
	var find = "Dunia"

	var txt = "Golang"

	halo = strings.Replace(halo, find, txt, 1);

	fmt.Println(halo)

	//soal 3
	var kataPertama = "saya";
	var kataKedua = "senang";
	var kataKetiga = "belajar";
	var kataKeempat = "golang";

	kataKedua = strings.Replace(kataKedua, "s", "S", 1)
	kataKetiga = strings.Replace(kataKetiga, "r", "R", 1)
	kataKeempat = strings.ToUpper(kataKeempat)

	var modifiedSentence = kataPertama + " " + kataKedua + " " +kataKetiga + " " +kataKeempat
	fmt.Println(modifiedSentence)

	//soal 4

	var angkaPertama= "8";
	var angkaKedua= "5";
	var angkaKetiga= "6";
	var angkaKeempat = "7";

	angka1, _ := strconv.ParseInt(angkaPertama, 10, 64)
	angka2, _ := strconv.ParseInt(angkaKedua, 10, 64)
	angka3, _ := strconv.ParseInt(angkaKetiga, 10, 64)
	angka4, _ := strconv.ParseInt(angkaKeempat, 10, 64)

	jumlah := angka1 + angka2 + angka3 + angka4

	fmt.Println(jumlah)

	//soal 5
	kalimat := "halo halo bandung"
	angka := 2022

	kalimat = strings.Replace(kalimat, "halo", "hi", -1)
	kalimat = strings.Replace(kalimat, "h", "H", -1)
	angkaStr := strconv.Itoa(angka)

	kalimat = kalimat + " - " + angkaStr

	fmt.Println(kalimat)

}