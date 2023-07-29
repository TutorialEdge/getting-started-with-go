package main

import (
	"fmt"
	"getting-started-with-go/internal/api"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "I'm alive!"}`))
}

func main() {
	fmt.Println("I'm officially a gopher!")
	mux := http.NewServeMux()

	mux.HandleFunc("/", homepage)
	mux.HandleFunc("/health", healthcheck)

	rocketHandler := &api.RocketHandler{}
	mux.Handle("/rocket",
		api.LoggingMiddleware(
			api.JSONMiddleware(rocketHandler),
		),
	)

	http.ListenAndServe(":8080", mux)
}
