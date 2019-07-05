package home

import "net/http"

func (h *HomeHandlers) home(w http.ResponseWriter, r *http.Request) {

	// h.db.ExecContext(r.Context(), "")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello Go Microservice!!"))
}
