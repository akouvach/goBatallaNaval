package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/akouvach/gobatallanaval/routes"
	"github.com/gorilla/mux"
)

// Search se utiliza para buscar algo
func Search(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Search \n")

}

// Load se utiliza para buscar algo
func Load(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Load \n")
}

// Canales se utiliza para buscar algo
func Canales(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	canal := vars["canal"]

	fmt.Fprintf(w, "Canal: %s \n", canal)

}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	// models.Init()
	r := routes.NewRouter()

	// r.Handle("/", http.FileServer(http.Dir("./static/")))
	// s := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	// r.PathPrefix("/static/").Handler(s)
	r.HandleFunc("/search/{searchTerm}", Search).Methods("GET")
	r.HandleFunc("/load/{dataId}", Load).Methods("GET")
	r.HandleFunc("/canales/{canal}", Canales).Methods("GET")

	http.Handle("/", r)

	http.ListenAndServe(":"+port, nil)
}
