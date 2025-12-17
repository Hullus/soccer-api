package main

import (
	"log"
	"net/http"
	"os"
	internalHttp "soccer-api/internal/http"
)

func main() {
	r := internalHttp.CreateRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, r))
}
