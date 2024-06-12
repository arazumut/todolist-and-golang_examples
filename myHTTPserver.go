package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, HTTPS!")
}

func main() {
	http.HandleFunc("/", helloHandler)

	// Sertifika ve anahtar dosyalarının yollarını belirtin
	certFile := "server.crt"
	keyFile := "server.key"

	// HTTPS sunucusunu başlatın
	log.Println("Starting HTTPS server on https://localhost:8081")
	err := http.ListenAndServeTLS(":8081", certFile, keyFile, nil)
	if err != nil {
		log.Fatalf("ListenAndServeTLS failed: %v", err)
	}
}
