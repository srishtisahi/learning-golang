package gomoviescrud

package main

import (
	"fmt"
	"log"
	"encoding/json" // encode data into json for postman
	"math/rand" // new movie ids will be random
	"net/http"
	"strconv" // random movie ids will be converted to string
	"github.com/gorilla/mux" // external library
)

// `` is below the esc button on laptop, not ''
// structs will be sent to postman as json

type Movie struct {
	ID string `json:"id"`
	Isbn string `json: "isbn"` // unique movie id
	Title string `json: "title"`
	Director *Director `json: "director"` // pointer to director struct
}

// every movie has a director - associated to movie struct 
type Director struct {
	Firstname string `json: "firstname"`
	Lastname string `json: "lastname"`
}

var movies []Movie // slice of movie structs

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // say that we're working in json
	json.NewEncoder(w).Encode(movies) // encode response that will be sent, pass the complete movies slice
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // pass an id of a movie as a parameter to delete it
	for index, item := range movies { // going over all movies to find a particular movie
		
		if item.ID == params["id"] { // access each id of the movies, check for a match with the one to delete
			movies = append(movies[:index], movies[index+1:]...) // movies[index+1:]... will append all of the data except for the movie id movies[:index] which will be deleted
			break
		} 
	}
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header.Set("Content-Type", "application/json")
	params := mux.Vars(r) // get request parameters with this mux function
	for _, item :=  
}

func main() {
	r := mux.NewRouter() // function from gorilla mux library

	movies = append(movies, Movies(ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}))
	movies = append(movies, Movie(ID: "2", Isbn: "45455", Title: "Movie Two", Director: &Director {Firstname: "Steve", Lastname: "Smith"}))

	// we will have 5 functions as per table in project-notes.md
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r)) // log out if server doesn't start
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000)) // value between zero and this number, converted to string
	movies = append(movies, movie) // movies (json) <--append-- movie (golang readable) 
	json.NewEncoder(w).Encode(movie)
}

func updateMovie() {
	// set json content type
	w.Header().Set("Content-Type", "application/json")

	// access params
	params := mux.Vars(r)

	// loop over the movies

	// delete the movie with the id that we sent

	// add a new movie - the one that you send in the body of postman
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1]...)
		}
	}

}