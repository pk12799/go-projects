package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	Name     string    `json:"name"`
	Id       string    `json:id`
	Isbn     string    `json:"isbn"`
	Director *Director `json:"director"`
}
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func allmovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
func dmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func movie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
}

func cmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.Id = strconv.Itoa(rand.Intn(100000000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func umovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.Id = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movies)
		}
	}

}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{Id: "1", Isbn: "4578", Name: "ddlj", Director: &Director{Firstname: "parvez", Lastname: "khan"}})
	movies = append(movies, Movie{Id: "2", Isbn: "4586", Name: "sholey", Director: &Director{Firstname: "khan", Lastname: "Parvez"}})
	fmt.Println(movies)
	r.HandleFunc("/movies", allmovies).Methods("GET")
	r.HandleFunc("/movie/{id}", movie).Methods("GET")
	r.HandleFunc("/movies", cmovie).Methods("POST")
	r.HandleFunc("/movies/{id}", umovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", dmovie).Methods("DELETE")
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}
}
