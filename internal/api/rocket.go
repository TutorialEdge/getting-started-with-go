package api

import "net/http"

type RocketHandler struct{}

func NewRocketHandler() *RocketHandler {
	return &RocketHandler{}
}

func (h *RocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// we can attempt to extract the rocketID from the incoming request
	rocketID := r.URL.Query().Get("id")

	switch {
	case r.Method == http.MethodGet && rocketID == "":
		h.getAllRockets(w, r)
		return
	case r.Method == http.MethodGet && rocketID != "":
		h.getRocketByID(w, r)
		return
	case r.Method == http.MethodPost:
		h.createRocket(w, r)
		return
	case r.Method == http.MethodPut:
		h.updateRocket(w, r)
		return
	case r.Method == http.MethodDelete:
		h.deleteRocket(w, r)
		return
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "rocket not found"}`))
		// oh no, nothing has matched!
		return
	}
}
