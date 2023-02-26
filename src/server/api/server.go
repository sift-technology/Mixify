package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Response struct {
	ID uuid.UUID   `json:"ID"`
	R1 json.Number `json:"R1"`
	R2 json.Number `json:"R2"`
	R3 json.Number `json:"R3"`
	R4 json.Number `json:"R4"`
}

type Server struct {
	*mux.Router

	Responses_DB []Response
}

func NewServer() *Server {
	s := &Server{
		Router:       mux.NewRouter(),
		Responses_DB: []Response{},
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.HandleFunc("/response", s.ListResponses()).Methods("GET")
	s.HandleFunc("/results", s.CreateResponse()).Methods("POST")
	s.HandleFunc("/response/{id}", s.removeResponse()).Methods("DELETE")
	s.PathPrefix("/").Handler(AngularHandler).Methods("GET")
}

func (s *Server) CreateResponse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user Response
		var i [4]json.Number
		fmt.Println("reached")
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		user.R1 = i[0]
		user.R2 = i[1]
		user.R3 = i[2]
		user.R4 = i[3]
		user.ID = uuid.New()
		s.Responses_DB = append(s.Responses_DB, user)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(i); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) ListResponses() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(s.Responses_DB); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) removeResponse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr, _ := mux.Vars(r)["id"]
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		for i, item := range s.Responses_DB {
			if item.ID == id {
				s.Responses_DB = append(s.Responses_DB[:i], s.Responses_DB[i+1:]...)
				break
			}
		}
	}
}
