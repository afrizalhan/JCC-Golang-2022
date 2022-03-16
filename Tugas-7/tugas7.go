package main

import "fmt"
import "strconv"

type buah struct{
	nama, warna string
	adaBijinya bool
	harga int
}

type segitiga struct{
	alas, tinggi int
}
  
type persegi struct{
	sisi int
}
  
type persegiPanjang struct{
	panjang, lebar int
}

type phone struct{
	name, brand string
	year int
	colors []string
}

type movie struct{
	title, genre string
	duration, year int
}

func luasPersegiPanjang(panjang int, lebar int) int {
	return panjang * lebar
}

func luasPersegi(sisi int) int {
	return sisi * sisi
}

func luasSegitiga(alas int, tinggi int) int {
	return alas / 2 * tinggi
}

func tambahWarna(warnaPointer *[]string, warna1 string, warna2 string, warna3 string){
	*warnaPointer = append(*warnaPointer, warna1, warna2, warna3)
}

func main(){
	//soal 1
	fmt.Println("=======SOAL 1=========")
	var nanas = buah{"Nanas", "Kuning", false, 9000}
	var jeruk = buah{"Jeruk", "Oranye", true, 8000}
	var semangka = buah{"Semangka", "Hijau & Merah", true, 10000}
	var pisang = buah{"Pisang", "Kuning", false, 15000}

	fmt.Println(nanas)
	fmt.Println(jeruk)
	fmt.Println(semangka)
	fmt.Println(pisang)

	//soal 2
	fmt.Println("=======SOAL 2=========")
	var segitiga = segitiga{6, 4}
	var persegi = persegi{7}
	var persegiPanjang = persegiPanjang{8, 9}

	fmt.Printf("Luas Segitiga : %d \n", luasSegitiga(segitiga.alas, segitiga.tinggi))
	fmt.Printf("Luas Persegi  : %d \n", luasPersegi(persegi.sisi))
	fmt.Printf("Luas Persegi Panjang : %d \n", luasPersegiPanjang(persegiPanjang.panjang, persegiPanjang.lebar))
	
	//soal 3
	fmt.Println("=======SOAL 3=========")
	var myPhone = phone{name : "Note 8", brand : "Redmi", year : 2020}

	tambahWarna(&myPhone.colors, "Merah", "Silver", "Navy")

	fmt.Println(myPhone)

	//soal 4
	fmt.Println("=======SOAL 4=========")
	var dataFilm = []movie{}

	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	for i, item := range dataFilm {
		durasi := item.duration / 60
		durasiStr := strconv.Itoa(durasi) + " jam"

		no := strconv.Itoa(i + 1)

		fmt.Println(no + ". \ttitle : " + item.title)
		fmt.Println("\tduration : " + durasiStr)
		fmt.Println("\tgenre : " + item.genre)
		fmt.Println("\tyear : " + strconv.Itoa(item.year))
	}
}

func tambahDataFilm(judul string, durasi int, genre string, tahunRilis int, filmPointer *[]movie){
	var film = movie{judul, genre, durasi, tahunRilis}

	*filmPointer = append(*filmPointer, film)
}

