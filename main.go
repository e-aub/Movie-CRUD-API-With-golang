package main

import (
	"fmt"
	"log"
	"net/http"

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

func getMovie(w http.ResponseWriter, r *http.Request) {

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
