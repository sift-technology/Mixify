package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/zmb3/spotify"
)

type Response struct {
	ID uuid.UUID   `json:"ID"`
	R1 json.Number `json:"R1"`
	R2 json.Number `json:"R2"`
	R3 json.Number `json:"R3"`
	R4 json.Number `json:"R4"`
	R5 json.Number `json:"R5"`
	R6 json.Number `json:"R6"`
	M  struct {
		Danceability float64 `json:"danceability"`
		Energy       float64 `json:"energy"`
		Popularity   int     `json:"popularity"`
		Acousticness float64 `json:"acousticness"`
	}
}

type Rec struct {
	ID   uuid.UUID             `json:"ID"`
	Recs []spotify.SimpleTrack `json:"tracks"`
}

type Server struct {
	*mux.Router

	Responses_DB []Response
	Rec_DB       []Rec
	Client       *spotify.Client
}

func NewServer() *Server {
	s := &Server{
		Router:       mux.NewRouter(),
		Responses_DB: []Response{},
		Rec_DB:       []Rec{},
		Client:       Authenticate(),
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
		var userRec Rec
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
		userRec.ID = user.ID
		userRec.Recs = Recommend(s.Client, &user)
		s.Rec_DB = append(s.Rec_DB, userRec)
		// create track attributes here
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
	R5, err5 := user.R5.Int64()
	R6, err6 := user.R6.Int64()

	if err1 != nil && err2 != nil && err3 != nil && err4 != nil && err5 != nil && err6 != nil {
		fmt.Print("can't convert Response")
	}

	DanceWeightR1 := [4]float64{0.25, 0.5, 0.75, 1.0}[R1-1]
	user.M.Danceability = DanceWeightR1 //average

	EnergyWeightR2 := [4]float64{0.25, 0.5, 0.75, 1.0}[R2-1]
	user.M.Energy = EnergyWeightR2 //scaled

	AccusticnessWeightR4 := [4]float64{0.25, 0.5, 0.75, 1.0}[R4-1]
	user.M.Acousticness = AccusticnessWeightR4

	PopularityWeightR6 := int(R6) * 100
	user.M.Popularity = PopularityWeightR6

}

func (s *Server) ListResponses() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(s.Rec_DB); err != nil {
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
