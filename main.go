package main

import (
	"log"
	"main/handlers"
	"net/http"
)

func main() {
	log.Println("Serveur Go en démarrage...")

<<<<<<< HEAD
	http.HandleFunc("/", handlers.HomeHandler)
=======
	http.HandleFunc("/contact", handlers.ContactHandler)
>>>>>>> templates/contact

	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
