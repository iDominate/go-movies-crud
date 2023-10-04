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
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, movie := range movies {
		if params["id"] == movie.Id {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, movie := range movies {
		if params["id"] == movie.Id {
			json.NewEncoder(w).Encode(movie)
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "apllication/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	id, err := strconv.Atoi(movies[len(movies)-1].Id)

	if err != nil {
		log.Fatal("An Error has occured")
	}
	movie.Id = strconv.Itoa((id + 1))
	movie.Isbn = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "apllication/json")
	params := mux.Vars(r)

	for index, movie := range movies {
		if params["id"] == movie.Id {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.Id = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return

		}
	}

}

func main() {
	router := mux.NewRouter()

	movies = append(movies, Movie{
		Id:    "1",
		Isbn:  "485767",
		Title: "Movie 1",
		Director: &Director{
			FirstName: "John",
			LastName:  "Doe",
		},
	})

	movies = append(movies, Movie{
		Id:    "2",
		Isbn:  "657767",
		Title: "Movie 2",
		Director: &Director{
			FirstName: "Anne",
			LastName:  "Smith",
		},
	})

	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting Server on port 8000")

	log.Fatal(http.ListenAndServe(":8000", router))

}
