package home

import (
	"log"
	"net/http"
	"time"
)

type HomeHandlers struct {
	logger *log.Logger
	// db     *sqlx.DB
}

// a pattern in go where functions are being used as constrouctors
func NewHomeHanders(logger *log.Logger) *HomeHandlers {
	return &HomeHandlers{
		logger: logger,
		// db:     db,
	}
}

// middleware in go
// we can add instrumentation
func (h *HomeHandlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer h.logger.Printf("request processed in %d ms \n", time.Now().Sub(startTime))
		next(w, r)
	}
}

// to have control over the package we will pass it the mux
// so that all the package logic is self contained
func (h *HomeHandlers) SetupRoutes(mux *http.ServeMux) {
	// and now here we can call our middleware
	mux.HandleFunc("/", h.Logger(h.home))
}
