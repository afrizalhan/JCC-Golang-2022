package main

import (
	"fmt"
	// "strconv"
)

func main(){
	
	//soal 1

	for i:= 1; i <= 20; i++{
		if i % 3 == 0 && i % 2 != 0{
			fmt.Printf("%d - I Love Coding \n", i)
		} else if i % 2 == 0 {
			fmt.Printf("%d - Candradimuka \n", i)
		} else {
			fmt.Printf("%d - JCC \n", i)
		}
	}

	//soal 2
	rows := 7
	hashes := ""

	for i := 1; i <= rows; i++{
		for j := 0; j < i; j++{
			hashes += "#"
		}
		hashes += "\n"
	}

	fmt.Printf(hashes)

	//soal 3
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	var newKalimat = kalimat[2:7]

	fmt.Println(newKalimat)

	//soal 4
	var sayuran = []string{}
	var sayuranS = append(sayuran, "Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun")

	for i := 0; i < len(sayuranS); i++{
		fmt.Printf("%d. %s \n", i + 1, sayuranS[i])
	}

	//soal 5
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}
	volume := 1

	for key, element := range satuan{
		fmt.Println(key, "=" , element)
		volume *= element
	}
	fmt.Printf("Volume Balok = %d \n", volume)
}