package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	http.Handle("/", http.FileServer(http.Dir("./static")))

	fmt.Printf("Iniciando server en puerto %v \n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
