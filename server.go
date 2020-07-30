package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/akouvach/gobatallanaval/routes"
)

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

	// Choose the folder to serve
	staticDir := "/"

	// Create the route
	r.PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("./static"+staticDir))))

	http.Handle("/", r)

	fmt.Printf("Iniciando servidor en el puerto %v", port)
	http.ListenAndServe(":"+port, nil)
}
