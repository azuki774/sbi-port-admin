package server

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

func (s *Server) rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "It is the root page.\n")
}

func (s *Server) registHandler(w http.ResponseWriter, r *http.Request) {
	// Get body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "internal error: %v\n", err)
		return
	}
	defer r.Body.Close()
	s.Usecase.RegistDailyRecords(context.Background(), string(body))
	fmt.Fprintf(w, "OK")
}
