package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Serveur Go en démarrage...")

	http.HandleFunc("/", HomeHandler)

	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bonjour depuis HomeHandler")
}
