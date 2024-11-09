package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

// Corregir el nombre de la función y la declaración de parámetros
func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World")
}

func main() {
	// Corregir la obtención del puerto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Valor por defecto si no se establece un puerto
	}

	// Corregir la declaración de la ruta
	http.HandleFunc("/", hello)

	// Corregir la forma en que se concatena el mensaje en el log
	log.Printf("Listening on :%s", port)

	// Corregir el llamado a ListenAndServe
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
