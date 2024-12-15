package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	//handle default route
	r.HandleFunc("/", HomeHandler)

	r.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("About Us"))
	})

	r.HandleFunc("/login", LoginPageHandler).Methods("GET")
	r.HandleFunc("/login", LoginHandler).Methods("POST")
	r.HandleFunc("/about", AboutHandler)
	r.HandleFunc("/dashboard", DashboardHandler).Methods("GET")

	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))


	r.HandleFunc("/search", SearchHandler).Methods("GET")

	http.Handle("/", r)
	fmt.Println("Server Ready")
	http.ListenAndServe(":8080", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/about.html")
	}

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/login.html")
	}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/dashboard.html")
	}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Cek jika metode request adalah POST
		// Parsing data dari form
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")

		// Validasi username dan password (contoh sederhana)
		if username == "admin" && password == "password123" {
			http.Redirect(w, r, "/dashboard", http.StatusFound)
			return
		} 
		
		errorMessage := "Username atau password salah"
		http.Error(w, errorMessage, http.StatusUnauthorized)
}


func SearchHandler( w http.ResponseWriter, r *http.Request){
	vars := r.URL.Query()
	query := vars.Get("q")

	aStr := vars.Get("a")
	bStr := vars.Get("b")
	a, errA := strconv.Atoi(aStr)
	b, errB := strconv.Atoi(bStr)

	if errA != nil || errB != nil {
		http.Error(w, "Parameter A dan B harus berupa bilangan", http.StatusBadRequest)
		return
	}

	c := a + b

	responseMessage := fmt.Sprintf("Hasil pencarian untuk: %s . Hasil Penjumlahan: %d + %d = %d", query,a,b,c)
	w.Write([]byte(responseMessage))
}