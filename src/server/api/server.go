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
	M  struct {
		Danceability float32 `json:"danceability"`
		Energy       float32 `json:"energy"`
		Popularity   int64   `json:"popularity"`
		Acousticness float32 `json:"acousticness"`
	}
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
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		user.R1 = i[0]
		user.R2 = i[1]
		user.R3 = i[2]
		user.R4 = i[3]
		user.ID = uuid.New()
		Weights(&user)
		s.Responses_DB = append(s.Responses_DB, user)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(i); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func Weights(user *Response) {
	// metric = {danceability, energy, popularity, acousticness}
	R1, err1 := user.R1.Int64()
	R2, err2 := user.R2.Int64()
	R3, err3 := user.R3.Int64()
	R4, err4 := user.R4.Int64()

	if err1 != nil && err2 != nil && err3 != nil && err4 != nil {
		fmt.Print("can't convert Response")
	}

	DanceWeightR1 := [4]float32{0.5, 0.2, 0.4, 0.1}[R1-1]
	DanceWeightR3 := float32(R3 / 100)
	user.M.Danceability = (DanceWeightR1 + DanceWeightR3) / 2 //average

	EnergyWeightR1 := ([4]float32{0.2, 0.4, 0.1, 0.3}[R1-1]) * 0.2
	EnergyWeightR4 := ([4]float32{0.8, 0.4, 0.3, 0.7}[R4-1]) * 0.8
	user.M.Energy = EnergyWeightR1 + EnergyWeightR4 //scaled

	PopularityWeightR1 := [4]int64{20, 60, 30, 80}[R1-1]
	PopularityWeightR2 := [4]int64{40, 20, 60, 70}[R2-1]
	PopularityWeightR3 := R3
	PopularityWeightR4 := [4]int64{80, 50, 30, 60}[R4-1]
	user.M.Popularity = (PopularityWeightR1 + PopularityWeightR2 + PopularityWeightR3 + PopularityWeightR4) / 4

	AccusticnessWeightR4 := [4]float32{0.4, 0.8, 0.3, 0.9}[R4-1]
	user.M.Acousticness = AccusticnessWeightR4
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
