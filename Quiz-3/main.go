package main

import (
	"fmt"
	"Quiz-3/functions"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var (
	authUser  = []string{"admin", "editor", "trainer"}
	authPass  = []string{"password", "secret", "rahasia"}
)

func main() {
	router := httprouter.New()


	router.GET("/bangun-datar/segitiga-sama-sisi", functions.GetSegitiga)
	router.GET("/bangun-datar/persegi", functions.GetPersegi)
	router.GET("/bangun-datar/persegi-panjang", functions.GetPersegiPanjang)
	router.GET("/bangun-datar/lingkaran", functions.GetLingkaran)
	router.GET("/bangun-datar/jajar-genjang", functions.GetJajarGenjang)

	router.GET("/categories", functions.GetAllCategory)
	router.POST("/categories", auth(http.HandlerFunc(functions.PostCategory)))
	router.PUT("/categories/:id", auth(http.HandlerFunc(functions.UpdateCategory)))
	router.DELETE("/categories/:id", auth(http.HandlerFunc(functions.DeleteCategory)))
	router.GET("/categories/:id/books", functions.GetBooksByCategory)

	router.GET("/books", functions.GetAllBooks)
	router.POST("/books", auth(http.HandlerFunc(functions.PostBook)))
	router.PUT("/books/:id", auth(http.HandlerFunc(functions.UpdateBook)))
	router.DELETE("/books/:id", auth(http.HandlerFunc(functions.DeleteBook)))
	// serve
	fmt.Println("Server Running at Port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func auth(next http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		user, pass, ok := r.BasicAuth()

		password := ""
		isUserRegistered := false

		for i, u := range authUser {
			if u == user {
				password = authPass[i]
				isUserRegistered = true
			}
		}

		if r.Method != "POST" {
			if !ok {
				w.Write([]byte("Username atau Password tidak boleh kosong"))
				return
			}

			if !isUserRegistered || pass != password {
				w.Write([]byte("Username atau Password tidak sesuai"))
				return
			}
		}
		next.ServeHTTP(w, r)
	}
}




