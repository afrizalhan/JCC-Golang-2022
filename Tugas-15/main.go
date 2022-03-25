package main

import (
	"fmt"
	"Tugas-15/functions"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)


func main() {
	router := httprouter.New()
	router.GET("/nilai", functions.GetNilai)
	router.POST("/nilai/create", functions.PostNilai)
	router.PUT("/nilai/:id/update", functions.UpdateNilai)
	router.DELETE("/nilai/:id/delete", functions.DeleteNilai)

	router.GET("/mahasiswa", functions.GetMahasiswa)
	router.POST("/mahasiswa/create", functions.PostMahasiswa)
	router.PUT("/mahasiswa/:id/update", functions.UpdateMahasiswa)
	router.DELETE("/mahasiswa/:id/delete", functions.DeleteMahasiswa)

	router.GET("/matkul", functions.GetMk)
	router.POST("/matkul/create", functions.PostMk)
	router.PUT("/matkul/:id/update", functions.UpdateMk)
	router.DELETE("/matkul/:id/delete", functions.DeleteMk)
	// serve
	fmt.Println("Server Running at Port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}




