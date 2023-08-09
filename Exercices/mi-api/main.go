package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hola, mundo")
}

func main() {
	// Configurar enrutador
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)

	// Iniciar servidor HTTP
	log.Println("Servidor iniciado en http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", mux))
}
