package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// NewRouter crea un router de gorilla mux
func NewRouter() *mux.Router {

	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/search/{searchTerm}", Search).Methods("GET")
	r.HandleFunc("/load/{dataId}", Load).Methods("GET")
	r.HandleFunc("/canales/{canal}", Canales).Methods("GET")

	return r
}

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
	r := NewRouter()

	// r.Handle("/", http.FileServer(http.Dir("./static/")))
	// s := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	// r.PathPrefix("/static/").Handler(s)

	// Choose the folder to serve
	staticDir := "/"

	// Create the route
	r.PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("./static"+staticDir))))

	http.Handle("/", r)

	fmt.Printf("Iniciando servidor en el puerto %v", port)
	http.ListenAndServe(":"+port, nil)
}
