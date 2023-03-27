package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/zmb3/spotify"
)

var badBunny spotify.ID = spotify.ID("1IHWl5LamUGEuP4ozKQSXZ")
var tSwift spotify.ID = spotify.ID("1vrd6UOGamcKNGnSHJQlSt")
var nirvana spotify.ID = spotify.ID("4P5KoWXOxwuobLmHXLMobV")
var oneHeart spotify.ID = spotify.ID("4xF4ZBGPZKxECeDFrqSAG4")
var fOcean spotify.ID = spotify.ID("3xKsf9qdS1CyvXSMEid6g8")
var tImpala spotify.ID = spotify.ID("52ojopYMUzeNcudsoz7O9D")
var cKeef spotify.ID = spotify.ID("01Lr5YepbgjXAWR9iOEyH1")
var iceSpice spotify.ID = spotify.ID("6AQbmUe0Qwf5PZnt4HmTXv")

type Response struct {
	ID uuid.UUID   `json:"ID"`
	R1 json.Number `json:"R1"`
	R2 json.Number `json:"R2"`
	R3 json.Number `json:"R3"`
	R4 json.Number `json:"R4"`
	R5 json.Number `json:"R4"`
	R6 json.Number `json:"R4"`
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

	if err1 != nil && err2 != nil && err3 != nil && err4 != nil {
		fmt.Print("can't convert Response")
	}

	DanceWeightR1 := [4]float64{0.5, 0.2, 0.4, 0.1}[R1-1]
	DanceWeightR3 := float64(R3 / 100)
	user.M.Danceability = (DanceWeightR1 + DanceWeightR3) / 2 //average

	EnergyWeightR1 := ([4]float64{0.2, 0.4, 0.1, 0.3}[R1-1]) * 0.2
	EnergyWeightR4 := ([4]float64{0.8, 0.4, 0.3, 0.7}[R4-1]) * 0.8
	user.M.Energy = EnergyWeightR1 + EnergyWeightR4 //scaled

	PopularityWeightR1 := [4]int{20, 60, 30, 80}[R1-1]
	PopularityWeightR2 := [4]int{40, 20, 60, 70}[R2-1]
	PopularityWeightR3 := int(R3)
	PopularityWeightR4 := [4]int{80, 50, 30, 60}[R4-1]
	user.M.Popularity = (PopularityWeightR1 + PopularityWeightR2 + PopularityWeightR3 + PopularityWeightR4) / 4

	AccusticnessWeightR4 := [4]float64{0.4, 0.8, 0.3, 0.9}[R4-1]
	user.M.Acousticness = AccusticnessWeightR4
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
