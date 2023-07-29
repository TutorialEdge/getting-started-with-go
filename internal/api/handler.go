package api

import "net/http"

func (h *RocketHandler) getAllRockets(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get all rockets"}`))
}

func (h *RocketHandler) getRocketByID(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get a single rocket"}`))
}

func (h *RocketHandler) createRocket(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "create a rocket"}`))
}

func (h *RocketHandler) updateRocket(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "update a rocket"}`))
}

func (h *RocketHandler) deleteRocket(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "delete a rocket"}`))
}
