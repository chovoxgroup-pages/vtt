//go:build ignore
// +build ignore

package main

import (
	"log"
	"net/http"
)

func main() {
	port := "8080"
	log.Printf("Starting pure Go dev server on http://localhost:%s", port)
	log.Printf("Serving Firebase 'public' directory...")

	// Serve the public directory where index.html and vtt.wasm live
	err := http.ListenAndServe(":"+port, http.FileServer(http.Dir("public")))
	if err != nil {
		log.Fatal(err)
	}
}
