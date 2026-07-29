package main

import (
	"log"
	"net/http"
)

func main() {
	const port = "8080"

	fileServerHandler := http.FileServer(http.Dir("."))

	mux := http.NewServeMux()
	mux.Handle("/", fileServerHandler)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Server running on port: %v\n", port)
	log.Fatal(server.ListenAndServe())
}
