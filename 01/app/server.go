package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	srv    *http.Server
	router *mux.Router
}

func newServer() *Server {
	s := &Server{
		router: mux.NewRouter(),
	}
	s.configureRouter()
	return s
}

func (s *Server) Start(addr string) error {
	s.srv = &http.Server{
		Addr:         addr,
		Handler:      s.router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	err := s.srv.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}
	return err
}

func (s *Server) Stop(ctx context.Context) error {
	err := s.srv.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("server shutdown: %w", err)
	}
	return nil
}

func (s *Server) configureRouter() {
	s.router.HandleFunc("/health", s.handleHello).Methods(http.MethodGet)
	s.router.HandleFunc("/healthz", s.healthz).Methods(http.MethodGet)
	s.router.HandleFunc("/readyz", s.readyz).Methods(http.MethodGet)
}

func (s *Server) handleHello(w http.ResponseWriter, r *http.Request) {
	log.Printf("Get request %s\n", r.URL)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "OK"}`))
}

func (s *Server) healthz(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *Server) readyz(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
