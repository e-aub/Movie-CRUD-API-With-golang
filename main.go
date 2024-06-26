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
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	for index, item := range movies {
		if item.ID == vars["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	for _, item := range movies {
		if item.ID == vars["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&movie)
	for index, item := range movies {
		if item.ID == params["id"] {
			movie.ID = params["id"]
			movies = append(movies[:index], movies[index+1:]...)
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movies)
			return
		}
	}
}
func main() {
	movies = append(movies, Movie{ID: "001", Isbn: "0115", Title: "Interstellar", Director: &Director{
		FirstName: "Christopher",
		LastName:  "Nolan",
	}})
	movies = append(movies, Movie{ID: "002", Isbn: "011543", Title: "Oppenheimer", Director: &Director{
		FirstName: "john",
		LastName:  "doo",
	}})
	r := mux.NewRouter()
	//defining routes
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
