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

func validateImageUrl(s string) bool {
	return strings.HasPrefix(s, "http")
}

func validateYear(year int) bool {
	isValid := false
	if year > 1979 && year < 2022 {
		isValid = true
	}
	return isValid
}


func GetThickness(n int) string {
	result := ""
	switch {
	case n >= 201:
		result = "Tebal"
	case n >= 101:
		result = "Sedang"
	default:
		result = "Tipis"
	}
	return result
}


func GetAllBooks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	books, err := query.GetAllBooks(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, books, http.StatusOK)
	
}

func PostBook(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type json", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var book models.Books

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	isYearValid := validateYear(book.ReleaseYear)

	errMsg := ""
	if !isYearValid {
		errMsg += "masukkan tahun rilis yang benar (1980 - 2021)"
	}

	if !validateImageUrl(book.ImageUrl){
		errMsg += "masukkan url yang benar"
	}

	if errMsg != ""{
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	book.Thickness = GetThickness(book.TotalPage)

	if err := query.InsertBook(ctx, book); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
 	utils.ResponseJSON(w, res, http.StatusCreated)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var book models.Books

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	isYearValid := validateYear(book.ReleaseYear)

	errMsg := ""
	if !isYearValid {
		errMsg += "masukkan tahun rilis yang benar (1980 - 2021)\n"
	}

	if !validateImageUrl(book.ImageUrl){
		errMsg += "masukkan url yang benar"
	}

	if errMsg != ""{
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	book.Thickness = GetThickness(book.TotalPage)


	id := strings.Split(r.URL.Path, "/")[2]

	if err := query.UpdateBook(ctx, book, id); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	ctx, cancel := context.WithCancel(context.Background())
  	defer cancel()

	id := strings.Split(r.URL.Path, "/")[2]

	if err := query.DeleteBook(ctx, id); err != nil {
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

func GetBooksByCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	ctx, cancel := context.WithCancel(context.Background())
  	defer cancel()

	id := ps.ByName("id")

	books, err := query.GetBooksByCategory(ctx, id)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, books, http.StatusOK)
}

