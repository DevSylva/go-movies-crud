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

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"Director"`
}

type Director struct{
	firstName string `json:"firstName"`
	lastName string `json:"lastName"`
}

var movies []Movie

// get movies function
func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// delete movie function
func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
		break
		}
	}
}

// get movie function
func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		for item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(movies)
}

// create movie function
func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_= json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa((rand.Intn(100000000000)))
	json.NewEncoder(w).Encode(movie)
}

// update movie function
func updateMovie(w http.ResponseWriter, r *http.Request){
	// set json content type
	w.Header().Set("Content-Type", "application/json")
	// params
	// loop ove the movies, range
	// delete the moview with the id
	// add a new movie
	params := mux.Vars(r)
	for index, item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "345542", Title: "Movie One", Director: &Director{firstName: "Johgn", lastName: "Doe"}})

	movies = append(movies, Movie{ID: "2", Isbn: "542342", Title: "Movie Two", Director: &Director{firstName: "Peter", lastName: "Don"}})

	movies = append(movies, Movie{ID: "3", Isbn: "347642", Title: "Movie Three", Director: &Director{firstName: "Barry", lastName: "Star"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movie{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}