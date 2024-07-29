package main

import (
	"log"
	"net/http"
)

func readyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	const port = "8080"

	mux := http.NewServeMux()
	mux.Handle("/app/*", http.FileServer(http.Dir(".")))
	mux.Handle("/healthz", http.HandlerFunc(readyHandler))

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
