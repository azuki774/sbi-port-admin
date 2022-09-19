package server

import (
	"azuki774/sbiport-server/internal/usecase"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func init() {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	time.Local = jst
}

func (s *Server) rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "It is the root page.\n")
}

func (s *Server) registHandler(w http.ResponseWriter, r *http.Request) {
	pathParam := mux.Vars(r)
	date := pathParam["date"]
	t, err := time.ParseInLocation("20060102", date, time.Local)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "date parse error: %v\n", err)
	}

	// Get body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "internal error: %v\n", err)
		return
	}
	defer r.Body.Close()

	result, err := s.Usecase.RegistDailyRecords(context.Background(), string(body), t)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "internal error: %v\n", err)
		return
	}

	outputJson, err := json.Marshal(&result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "internal error: %v\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(outputJson))
}

func (s *Server) getDailyHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	pathParam := mux.Vars(r)
	date := pathParam["date"]

	results, err := s.Usecase.GetDailyRecords(context.Background(), date)
	if err != nil {
		if errors.Is(err, usecase.ErrInvalidDate) {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "bad request: %v\n", err)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "internal error: %v\n", err)
		}
		return
	}

	if len(results) == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "record is not regisited")
		return
	}

	outputJson, err := json.Marshal(&results)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "internal error: %v\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(outputJson))
}
