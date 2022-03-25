package models


type (
	NilaiMahasiswa struct {
		ID          uint   `json:"id"`
		Indeks string `json:"indeks"`
		Skor        uint `json:"skor"`
		MahasiswaId  uint `json:"mahasiswa_id"`
		MataKuliahId       uint   `json:"mata_kuliah_id"`
	}
)