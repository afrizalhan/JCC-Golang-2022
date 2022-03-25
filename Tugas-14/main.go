package main

import (
	"context"
	"encoding/json"
	"fmt"
	"Tugas-14/models"
	"Tugas-14/nilai"
	"Tugas-14/utils"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)


func main() {
	router := httprouter.New()
	router.GET("/nilai", GetNilai)
	router.POST("/nilai/create", PostNilai)
	router.PUT("/nilai/:id/update", UpdateNilai)
	router.DELETE("/nilai/:id/delete", DeleteNilai)
	// serve
	fmt.Println("Server Running at Port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetNilai(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	nilai, err := nilai.GetAll(ctx)

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

	if err := nilai.Insert(ctx, nilaiMhs); err != nil {
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

	idMhs := ps.ByName("id")

	if err := nilai.Update(ctx, nilaiMhs, idMhs); err != nil {
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

	if err := nilai.Delete(ctx, idMhs); err != nil {
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