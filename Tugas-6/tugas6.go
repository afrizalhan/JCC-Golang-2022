package main

import "fmt"
import "strconv"

//function hitungLingkaran untuk soal no 1
func hitungLingkaran(r float64)(float64, float64){
	luasLigkaran := 3.14 * r * r
	kelilingLingkaran := 2 * 3.14 * r

	return luasLigkaran, kelilingLingkaran
}

func introduce(sentencePointer *string, nama string, jk string, profesi string, umur string) string {
	var panggilan string
	if jk == "perempuan"{
		panggilan = "Bu "
	} else {
		panggilan = "Pak "
	}
	*sentencePointer  = panggilan + nama + " adalah seorang " + profesi + " yang berusia " + umur + " tahun"
	return *sentencePointer
}

func tambahBuah(buahPointer *[]string, buahan ...string)([]string){
	for i := 0; i < len(buahan); i++{
		*buahPointer = append(*buahPointer, buahan[i])
	}
	return *buahPointer
}

func main(){
	//soal 1
	fmt.Println("=======SOAL 1=========")
	var luasLingkaran float64 
	var kelilingLingkaran float64

	luasPointer := &luasLingkaran
	kelilingPointer := &kelilingLingkaran

	*luasPointer, *kelilingPointer = hitungLingkaran(7)

	fmt.Printf("Luas Lingkaran : %f \n", luasLingkaran)
	fmt.Printf("Keliling Lingkaran : %f \n", kelilingLingkaran)

	//soal 2
	fmt.Println("=======SOAL 2=========")
	var sentence string 
	introduce(&sentence, "John", "laki-laki", "penulis", "30")

	fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"
	introduce(&sentence, "Sarah", "perempuan", "model", "28")

	fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	//soal 3
	fmt.Println("=======SOAL 3=========")
	var buah = []string{}

	var buahan = []string{"Jeruk", "Semangka", "Mangga", "Strawberry", "Durian", "Manggis", "Alpukat"}
	
	tambahBuah(&buah, buahan...)

	for i := 0; i < len(buah); i++{
		fmt.Printf("%d. " + buah[i] +"\n", i + 1)
	}

	//soal 4
	fmt.Println("=======SOAL 4=========")
	var dataFilm = []map[string]string{}

	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	i := 1
	for _, item := range dataFilm {
		fmt.Print(strconv.Itoa(i) + ".")
		for key, val := range item {
			fmt.Print("\t")
			fmt.Println(key,":",val)
		}
		i++
	}
	
}

func tambahDataFilm(judul string, durasi string, genre string, tahunRilis string, filmPointer *[]map[string]string){
	film := map[string]string{
		"title" : judul,
		"duration" : durasi,
		"genre" : genre,
		"year" : tahunRilis,
	}

	*filmPointer = append(*filmPointer, film)
}