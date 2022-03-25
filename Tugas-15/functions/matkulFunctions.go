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


func GetMk(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	nilai, err := query.GetMk(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, nilai, http.StatusOK)
	
}

func PostMk(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type json", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var matkul models.MataKuliah

	if err := json.NewDecoder(r.Body).Decode(&matkul); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}


	if err := query.InsertMk(ctx, matkul); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
 	utils.ResponseJSON(w, res, http.StatusCreated)
}

func UpdateMk(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var matkul models.MataKuliah

	if err := json.NewDecoder(r.Body).Decode(&matkul); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	idMk := ps.ByName("id")

	if err := query.UpdateMk(ctx, matkul, idMk); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)
}

func DeleteMk(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	ctx, cancel := context.WithCancel(context.Background())
  	defer cancel()

	idMk := ps.ByName("id")

	if err := query.DeleteMk(ctx, idMk); err != nil {
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