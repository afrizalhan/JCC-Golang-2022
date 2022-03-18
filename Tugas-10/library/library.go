package library

import (
	"fmt"
	"math"
)

func TambahPhone(phoneArr *[]string, brandList ...string){
	*phoneArr = append(*phoneArr, brandList...)
}

func HitungLingkaran(r float64) {
	luas := math.Round(3.14 * r * r)
	keliling := math.Round(3.14 * r * 2)

	fmt.Println("Luas Lingkaran :", luas)
	fmt.Println("Keliling Lingkaran :", keliling)
}

func HitungPersegiPanjang(p int, l int) {
	luas := p * l
	keliling := 2 * (p + l)

	fmt.Println("Luas Persegi Panjang :", luas)
	fmt.Println("Keliling Persegi Panjang :", keliling)
}