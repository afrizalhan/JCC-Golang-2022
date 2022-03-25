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

func GetMahasiswa(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	mhs, err := query.GetMhs(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, mhs, http.StatusOK)
	
}

func PostMahasiswa(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type json", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var mahasiswa models.Mahasiswa

	if err := json.NewDecoder(r.Body).Decode(&mahasiswa); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}


	if err := query.InsertMhs(ctx, mahasiswa); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
 	utils.ResponseJSON(w, res, http.StatusCreated)
}

func UpdateMahasiswa(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var mahasiswa models.Mahasiswa

	if err := json.NewDecoder(r.Body).Decode(&mahasiswa); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	idMhs := ps.ByName("id")

	if err := query.UpdateMhs(ctx, mahasiswa, idMhs); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)
}

func DeleteMahasiswa(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	ctx, cancel := context.WithCancel(context.Background())
  	defer cancel()

	idMhs := ps.ByName("id")

	if err := query.DeleteMhs(ctx, idMhs); err != nil {
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