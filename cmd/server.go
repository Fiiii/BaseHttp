package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type apiFunc func(http.ResponseWriter, *http.Request) error

type Config struct {
	ListenAddr string
}

type APIServer struct {
	*Config
}

type ApiError struct {
	Error string `json:"error"`
}

func NewServer(cfg *Config) (*APIServer, error) {
	return &APIServer{
		Config: cfg,
	}, nil
}

func (s *APIServer) Start() {
	router := mux.NewRouter()

	router.HandleFunc("/hello", makeHTTPHandleFunc(s.hello))

	log.Println("JSON API server running on port: ", s.Config.ListenAddr)

	err := http.ListenAndServe(s.Config.ListenAddr, router)
	if err != nil {
		log.Fatal(err)
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func (s *APIServer) hello(w http.ResponseWriter, r *http.Request) error {

	return fmt.Errorf("method not allowed %s", r.Method)
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}
