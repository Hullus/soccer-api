package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"soccer-api/internal/db"
	internalHttp "soccer-api/internal/http"
)

func main() {
	ctx := context.Background()

	pool, err := db.NewDBPool(ctx)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer pool.Close()

	r := internalHttp.CreateRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, r))
}
