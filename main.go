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
	ID       string    `json:"ID"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movie []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(movie)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range movie {
		if item.ID == params["id"] {
			movie = append(movie[:index], movie[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movie)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Connten-type", "application/json")
	params := mux.Vars(r)
	for _, item := range movie {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var movies Movie
	json.NewDecoder(r.Body).Decode(&movies)
	movies.ID = strconv.Itoa(rand.Intn(10000000))
	movie=append(movie,movies)
	json.NewEncoder(w).Encode(movies)
}
func updateMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-type","application/json")
	params:= mux.Vars(r)

	for index,item:=range movie{
		if item.ID == params["id"]{
			movie=append(movie[:index],movie[index+1:]...)
			var movies Movie
			json.NewDecoder(r.Body).Decode(&movies)
			movies.ID=params["id"]
			movie=append(movie,movies)
			json.NewEncoder(w).Encode(movies)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()
	movie = append(movie, Movie{ID: "1", Isbn: "438227", Title: "Movie one", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movie = append(movie, Movie{ID: "2", Isbn: "5343247", Title: "Movie two", Director: &Director{Firstname: "Steve", Lastname: "smith"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/(id)", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/(id)", updateMovie).Methods("PUT")
	r.HandleFunc("/movies(id)", deleteMovie).Methods("DELETE")

	fmt.Printf("starting the server at port 3001\n ")
	log.Fatal(http.ListenAndServe(":3001", r))

}
