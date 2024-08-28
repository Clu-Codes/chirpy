package main

import (
	"html/template"
	"log"
	"net/http"
)

func (cfg *apiConfig) requestCountHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	t, err := template.New("metrics").Parse(`
	<html>
		<body>
			<h1>Welcome, Chirpy Admin</h1>
			<p>Chirpy has been visited {{.}} times!</p>
		</body>
	</html>
	`)
	if err != nil {
		log.Printf("An error occurred: %v", err)
	}
	t.Execute(w, cfg.fileserverHits)
}

func (cfg *apiConfig) middlewareMetrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits++
		next.ServeHTTP(w, r)
	})
}
