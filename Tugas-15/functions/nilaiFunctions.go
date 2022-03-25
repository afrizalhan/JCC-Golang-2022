package functions

import (
	"context"
	"encoding/json"
	"fmt"
	"Tugas-15/models"
	"Tugas-15/query"
	"Tugas-15/utils"
	"net/http"

	"github.com/julienschmidt/httprouter"
)


func getIndeksNilai(nilai uint) string {
	switch {
	case nilai >= 80:
		return "A"
	case nilai >= 70 && nilai < 80:
		return "B"
	case nilai >= 60 && nilai < 70:
		return "C"
	case nilai >= 50 && nilai < 60:
		return "D"
	default:
		return "E"
	}
}


func GetNilai(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	nilai, err := query.GetNilai(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, nilai, http.StatusOK)
	
}

func PostNilai(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type json", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var nilaiMhs models.NilaiMahasiswa

	if err := json.NewDecoder(r.Body).Decode(&nilaiMhs); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if nilaiMhs.Skor > 100 {
		http.Error(w, "Nilai tidak boleh diinput lebih dari 100", http.StatusBadRequest)
		return
	}

	nilaiMhs.Indeks = getIndeksNilai(nilaiMhs.Skor)

	if err := query.InsertNilai(ctx, nilaiMhs); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
 	utils.ResponseJSON(w, res, http.StatusCreated)
}

func UpdateNilai(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var nilaiMhs models.NilaiMahasiswa

	if err := json.NewDecoder(r.Body).Decode(&nilaiMhs); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if nilaiMhs.Skor > 100 {
		http.Error(w, "Nilai tidak boleh diinput lebih dari 100", http.StatusBadRequest)
		return
	}

	nilaiMhs.Indeks = getIndeksNilai(nilaiMhs.Skor)

	idMhs := ps.ByName("id")

	if err := query.UpdateNilai(ctx, nilaiMhs, idMhs); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)
}

func DeleteNilai(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	ctx, cancel := context.WithCancel(context.Background())
  	defer cancel()

	idMhs := ps.ByName("id")

	if err := query.DeleteNilai(ctx, idMhs); err != nil {
		kesalahan := map[string]string{
		  "error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusOK)
}