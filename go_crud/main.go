package main

import (
	"encoding/json"
	"fmt"
	// "math/rand"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Movie struct {
	ID       string    `json:"id`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(movies)

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	num := vars["id"]
	id, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println("Conversion failed:", err)
		return
	}
	movies = append(movies[:id-1], movies[id:]...)

	fmt.Fprintf(w, "Successfully deleted")
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	params := mux.Vars(r)
	id := params["id"]
	id_num, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Conversion failed:", err)
		return
	}

	json.NewEncoder(w).Encode(movies[id_num-1])

}

func createMovie(w http.ResponseWriter, r *http.Request) {
	var new_movie Movie

	err := json.NewDecoder(r.Body).Decode(&new_movie)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	movies = append(movies, new_movie)
	fmt.Fprintf(w, "Received movie: %+v", new_movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	id_num, err := strconv.Atoi(id)

	if err != nil {
		fmt.Println("Conversion failed:", err)
		return
	}

	var update_movie Movie

	json.NewDecoder(r.Body).Decode(&update_movie)

	movies[id_num-1] = update_movie

	fmt.Fprintf(w, "Updated successfully")

}

func main() {
	router := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "438277", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "432789", Title: "Movie Two", Director: &Director{Firstname: "Alice", Lastname: "Wonderland"}})
	movies = append(movies, Movie{ID: "3", Isbn: "429309", Title: "Movie Three", Director: &Director{Firstname: "Mohan", Lastname: "Shyamlan"}})
	movies = append(movies, Movie{ID: "4", Isbn: "432902", Title: "Movie Four", Director: &Director{Firstname: "James", Lastname: "Cameron"}})
	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	port := ":8080"
	println("Server running at http://localhost" + port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		panic(err)
	}

}
