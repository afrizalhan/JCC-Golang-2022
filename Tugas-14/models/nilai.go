package models


type (
	NilaiMahasiswa struct {
		ID          uint   `json:"id"`
		Nama        string `json:"nama"`
		MataKuliah  string `json:"mata_kuliah"`
		Nilai       uint   `json:"nilai"`
		IndeksNilai string `json:"indeks_nilai"`
	}
)