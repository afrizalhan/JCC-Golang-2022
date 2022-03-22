package main

import (
	"fmt"
	"sync"
	"time"
	"strconv"
)

func printPhones(number int, phone string, wg *sync.WaitGroup){
	fmt.Println(strconv.Itoa(number) + ". " + phone)
	time.Sleep(time.Second)
	wg.Done()
}

func getMovies(movieChannel chan string, movies ...string){
	fmt.Println("List Movies : ")
	for i := 0; i < len(movies); i++{
		movie := strconv.Itoa(i + 1) + ". " + movies[i]
		movieChannel <- movie
	}
	close(movieChannel)
}

func luasPersegiPanjang(ch chan int, panjang int, lebar int){
	ch <- panjang * lebar
}

func kelilingPersegiPanjang(ch chan int, panjang int, lebar int){
	ch <- 2 * (panjang + lebar)
}

func volumeBalok(ch chan int, panjang int, lebar int, tinggi int){
	ch <- panjang * lebar * tinggi
}

func luasLingkaran(ch chan float64, r float64) {
	ch <- 3.14 * r * r
}

func kelilingLingkaran(ch chan float64, r float64) {
	ch <- 3.14 * r * 2
}

func volumeTabung(ch chan float64, r float64, t float64) {
	ch <- 3.14 * r * r * t
}

func main(){

	//soal 1
	fmt.Println("=======SOAL 1=========")
	var phones = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}

	var wg sync.WaitGroup

	for i := 0; i < len(phones); i++{
		wg.Add(1)
		printPhones(i + 1, phones[i], &wg)
		wg.Wait()
	}

	// soal 2
	fmt.Println("=======SOAL 2=========")
	var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}

	moviesChannel := make(chan string)

	go getMovies(moviesChannel, movies...)

	for value := range moviesChannel {
		fmt.Println(value)
	}

	//soal 3
	fmt.Println("=======SOAL 3=========")

	chRound := make(chan float64)
	
	go luasLingkaran(chRound, 8)
	go kelilingLingkaran(chRound, 8)
	go luasLingkaran(chRound, 14)
	go kelilingLingkaran(chRound, 14)
	go luasLingkaran(chRound, 20)
	go kelilingLingkaran(chRound, 20)
	go volumeTabung(chRound, 8, 10)

	for value := range chRound {
		fmt.Println(value)
	}
	
	
}