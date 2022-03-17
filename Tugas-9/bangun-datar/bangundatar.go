package bangundatar

type SegitigaSamaSisi struct{
	Alas, Tinggi int
}

type PersegiPanjang struct{
	Panjang, Lebar int
}

type HitungBangunDatar interface{
	Luas() int
	Keliling() int
}

func (s SegitigaSamaSisi) Luas() int {
	return s.Alas / 2 * s.Tinggi
}

func (s SegitigaSamaSisi) Keliling() int {
	return s.Alas * 3
}

func (p PersegiPanjang) Luas() int {
	return p.Panjang * p.Lebar
}

func (p PersegiPanjang) Keliling() int {
	return 2 * (p.Panjang + p.Lebar)
}