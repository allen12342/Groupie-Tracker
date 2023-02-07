package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Artist est la structure pour stocker les informations sur les artistes
type Artist struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Group    string `json:"group"`
	Location string `json:"location"`
	Albums   []Album
	Dates    []string
}

// Album est la structure pour stocker les informations sur les albums
type Album struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Artist string `json:"artist"`
}

// Location est la structure pour stocker les informations sur les lieux
type Location struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getLocations() []Location {
	return []Location{
		{ID: 1, Name: "Lieu 1"},
		{ID: 2, Name: "Lieu 2"},
	}
}

func getArtists() []Artist {
	return []Artist{
		{
			ID:       1,
			Name:     "Artiste 1",
			Group:    "Groupe 1",
			Location: "Lieu 1",
			Albums: []Album{
				{ID: 1, Name: "Album 1", Artist: "Artiste 1"},
				{ID: 2, Name: "Album 2", Artist: "Artiste 1"},
			},
			Dates: []string{"2022-05-01", "2022-06-02"},
		},
		{
			ID:       2,
			Name:     "Artiste 2",
			Group:    "Groupe 2",
			Location: "Lieu 2",
			Albums: []Album{
				{ID: 1, Name: "Album 1", Artist: "Artiste 2"},
				{ID: 2, Name: "Album 2", Artist: "Artiste 2"},
			},
			Dates: []string{"2022-07-01", "2022-08-02"},
		},
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("https://groupietrackers.herokuapp.com/api/artists", func(w http.ResponseWriter, r *http.Request) {
		artists := getArtists()
		response, err := json.Marshal(artists)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}).Methods("GET")

	router := mux.NewRouter()
	router.HandleFunc("https://groupietrackers.herokuapp.com/api/locations", func(w http.ResponseWriter, r *http.Request) {
		locations := getLocations()
		response, err := json.Marshal(locations)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}).Methods("GET")

	router := mux.NewRouter()
	router.HandleFunc("https://groupietrackers.herokuapp.com/api/dates", func(w http.ResponseWriter, r *http.Request) {
		dates := getDates()
		response, err := json.Marshal(dates)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}).Methods("GET")

	router := mux.NewRouter()
	router.HandleFunc("https://groupietrackers.herokuapp.com/api/relation", func(w http.ResponseWriter, r *http.Request) {
		relation := getRelation()
		response, err := json.Marshal(lrelation)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
