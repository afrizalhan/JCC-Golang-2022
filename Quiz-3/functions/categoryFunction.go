package functions

import (
	"context"
	"encoding/json"
	"fmt"
	"Quiz-3/models"
	"Quiz-3/query"
	"Quiz-3/utils"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func GetAllCategory(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	cat, err := query.GetAllCategory(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, cat, http.StatusOK)
	
}

func PostCategory(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type json", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var category models.Category

	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}


	if err := query.InsertCategory(ctx, category); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
 	utils.ResponseJSON(w, res, http.StatusCreated)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var category models.Category

	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	id := strings.Split(r.URL.Path, "/")[2]

	if err := query.UpdateCategory(ctx, category, id); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request){
	ctx, cancel := context.WithCancel(context.Background())
  	defer cancel()

	  id := strings.Split(r.URL.Path, "/")[2]

	if err := query.DeleteCategory(ctx, id); err != nil {
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