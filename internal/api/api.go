package api

import (
	"fmt"
	"net/http"
)

type API struct {
	mux *http.ServeMux
}

func New() *API {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homepage)
	mux.HandleFunc("/health", healthcheck)

	rocketHandler := &RocketHandler{}
	mux.Handle("/rocket",
		LoggingMiddleware(
			JSONMiddleware(rocketHandler),
		),
	)

	return &API{
		mux: mux,
	}
}

func (api *API) Start() error {
	return http.ListenAndServe(":8080", api.mux)
}

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
