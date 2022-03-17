package bangunruang

type Tabung struct{
	JariJari, Tinggi float64
}

type Balok struct{
	Panjang, Lebar, Tinggi int
}

type HitungBangunRuang interface{
	Volume() float64
	LuasPermukaan() float64
}

func (b Balok) Volume() float64 {
	return float64(b.Panjang) * float64(b.Lebar) * float64(b.Tinggi)
}

func (b Balok) LuasPermukaan() float64 {
	pl := b.Panjang * b.Lebar
	pt := b.Panjang * b.Tinggi
	lt := b.Lebar * b.Tinggi

	return 2 * (float64(pl) + float64(pt) + float64(lt))
}

func (t Tabung) Volume() float64 {
	return 3.14 * t.JariJari * t.JariJari * t.Tinggi
}

func (t Tabung) LuasPermukaan() float64 {
	return 2 * 3.14 * t.JariJari * (t.JariJari + t.Tinggi)
}