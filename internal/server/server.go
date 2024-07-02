package server

import (
	"azuki774/sbiport-server/internal/usecase"
	"net"
	"net/http"

	"github.com/gorilla/mux"

	"go.uber.org/zap"
)

type Server struct {
	Host    string
	Port    string
	Logger  *zap.Logger
	Usecase *usecase.Usecase
}

func (s *Server) Start() error {
	router := mux.NewRouter()
	s.addRecordFunc(router)

	return http.ListenAndServe(net.JoinHostPort(s.Host, s.Port), router)
}

func (s *Server) addRecordFunc(r *mux.Router) {
	r.HandleFunc("/", s.rootHandler)
	r.HandleFunc("/regist/{categoryTag}/{date}", s.registHandler).Methods("POST")
	r.HandleFunc("/daily/{categoryTag}/{date}", s.getDailyHandler).Methods("GET")
	r.Use(s.middlewareLogging)
}

func (s *Server) middlewareLogging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.Logger.Info("access", zap.String("url", r.URL.Path), zap.String("X-Forwarded-For", r.Header.Get("X-Forwarded-For")))
		h.ServeHTTP(w, r)
	})
}
