package main

import (
	"fmt"
	// "strconv"
)

//fungsi luas untuk soal no 1
func luasPersegiPanjang(panjang int, lebar int) int {
	return panjang * lebar
}

//fungsi keliling untuk soal no 1
func kelilingPersegiPanjang(panjang int, lebar int) int {
	return 2 * (panjang + lebar)
}

//fungsi volume untuk soal no 1
func volumeBalok(panjang int, lebar int, tinggi int) int {
	return panjang * lebar * tinggi
}

//fungsi introduce untuk soal no 2
func introduce(nama string, jk string, profesi string, umur string) (kalimat string){
	if jk == "perempuan"{
		kalimat = "Bu " + nama + " adalah seorang " + profesi + " yang berusia " + umur + " tahun"
		return kalimat
	}
	kalimat = "Pak " + nama + " adalah seorang " + profesi + " yang berusia " + umur + " tahun"
	return kalimat
}

//fungsi buah favorit untuk soal no 3
func buahFavorit(nama string, buahan ...string)(kalimat string){
	kalimat = "halo nama saya " + nama + " dan buah favorit saya adalah "

	for i := 0; i < len(buahan) - 1; i++{
		kalimat +=  buahan[i] + ", "
	}

	kalimat += buahan[len(buahan) - 1]

	return kalimat
}


func main(){
	
	//soal 1
	panjang := 12
	lebar := 4
	tinggi := 8
	
	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas) 
	fmt.Println(keliling)
	fmt.Println(volume)

	//soal 2
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	//soal 3
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}

	var buahFavoritJohn = buahFavorit("john", buah...)

	fmt.Println(buahFavoritJohn)
	// halo nama saya john dan buah favorit saya adalah "semangka", "jeruk", "melon", "pepaya"

	//soal 4
	var dataFilm = []map[string]string{}
	// buatlah closure function disini
	var tambahDataFilm = func(judul string, durasi string, genre string, tahunRilis string){
		film := map[string]string{
			"genre" : genre,
			"jam" : durasi,
			"tahun" : tahunRilis,
			"title" : judul,
		}
		dataFilm = append(dataFilm, film)
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}
}