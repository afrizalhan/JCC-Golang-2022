package functions

import (
	"Quiz-3/models"	
	"Quiz-3/utils"	
	"net/http"
	"strconv"
	"sync"

	"github.com/julienschmidt/httprouter"
)

func GetSegitiga(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var bangun models.BangunDatar

	hitung := r.URL.Query().Get("hitung")
	alas, err := strconv.Atoi(r.URL.Query().Get("alas"))
	if err != nil {
		alas = 0
	}
	tinggi, err := strconv.Atoi(r.URL.Query().Get("tinggi"))
	if err != nil {
		tinggi = 0
	}

	switch hitung {
	case "luas":
		bangun.Result = float64(alas * tinggi / 2)
	case "keliling":
		bangun.Result = float64(alas * 3)
	default:
		bangun.Result = 0
	}

	utils.ResponseJSON(w, bangun, http.StatusOK)
}

func GetPersegi(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var wg sync.WaitGroup
	var bangun models.BangunDatar

	printHasil := func() {
		utils.ResponseJSON(w, bangun, http.StatusOK)
		wg.Done()
	}

	hitung := r.URL.Query().Get("hitung")
	sisi, err := strconv.Atoi(r.URL.Query().Get("sisi"))
	if err != nil {
		sisi = 0
	}

	switch hitung {
	case "luas":
		bangun.Result = float64(sisi * sisi)
	case "keliling":
		bangun.Result = float64(sisi * 4)
	default:
		bangun.Result = 0
	}

	wg.Add(1)
	go printHasil()
	wg.Wait()
}

func GetPersegiPanjang(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var bangun models.BangunDatar

	hitung := r.URL.Query().Get("hitung")
	panjang, err := strconv.Atoi(r.URL.Query().Get("panjang"))
	if err != nil {
		panjang = 0
	}
	lebar, err := strconv.Atoi(r.URL.Query().Get("lebar"))
	if err != nil {
		lebar = 0
	}

	switch hitung {
	case "luas":
		bangun.Result = float64(panjang * lebar)
	case "keliling":
		bangun.Result = float64(2 * (panjang + lebar))
	default:
		bangun.Result = 0
	}

	utils.ResponseJSON(w, bangun, http.StatusOK)
}

func GetLingkaran(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var bangun models.BangunDatar

	hitung := r.URL.Query().Get("hitung")
	jari, err := strconv.Atoi(r.URL.Query().Get("jariJari"))
	if err != nil {
		jari = 0
	}

	inputCh := func(ch chan int, rs int) {
		ch <- rs
		close(ch)
	}

	jariChannel := make(chan int)
	go inputCh(jariChannel, jari)

	switch hitung {
	case "luas":
		bangun.Result = 3.14 * float64(<-jariChannel) * float64(<-jariChannel)
	case "keliling":
		bangun.Result = 2 * 3.14 * float64(<-jariChannel)
	default:
		bangun.Result = 0
	}

	utils.ResponseJSON(w, bangun, http.StatusOK)
}

func GetJajarGenjang(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var bangun models.BangunDatar

	hitung := r.URL.Query().Get("hitung")
	alas, err := strconv.Atoi(r.URL.Query().Get("alas"))
	if err != nil {
		alas = 0
	}
	tinggi, err := strconv.Atoi(r.URL.Query().Get("tinggi"))
	if err != nil {
		tinggi = 0
	}
	sisi, err := strconv.Atoi(r.URL.Query().Get("sisi"))
	if err != nil {
		sisi = 0
	}

	switch hitung {
	case "luas":
		bangun.Result = float64(alas * tinggi)
	case "keliling":
		bangun.Result = float64((2 * alas) + (2 * sisi))
	default:
		bangun.Result = 0
	}

	utils.ResponseJSON(w, bangun, http.StatusOK)
}