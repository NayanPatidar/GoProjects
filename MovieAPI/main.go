package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Movie struct {
	Name   string  `json:"MovieName"`
	ID     int     `json:"MovieID"`
	Genre  string  `json:"Genre"`
	Rating float32 `json:"Rating"`
}

func (c *Movie) isEmpty() bool {
	return c.Name == ""
}

var Data []Movie

func main() {
	route := mux.NewRouter()

	fmt.Println("Running Movie Server")

	//adding data
	Data = append(Data, Movie{Name: "Avengers:InfinityWar", ID: 1, Genre: "Science Fiction", Rating: 4.5})
	Data = append(Data, Movie{Name: "Avengers:Endgame", ID: 3, Genre: "Science Fiction", Rating: 4})

	route.HandleFunc("/GetAllMovies", GetAllMovies).Methods("GET")
	route.HandleFunc("/WatchMovie/{id}", WatchMovie).Methods("GET")
	route.HandleFunc("/AddMovie", AddMovie).Methods("POST")
	route.HandleFunc("/DeleteMovie/{id}", DeleteMovie).Methods("DELETE")
	route.HandleFunc("/UpdateMovie/{id}", UpdateMovie).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", route))

}

// DeleteMovie, AddMovie, UpdateMovie, GetAllMovies, WatchMovie
func GetAllMovies(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Getting All Movies.")

	jsonData, err := json.MarshalIndent(Data, "", "\t")
	if err != nil {
		panic(err)
	}
	w.Write(jsonData)
}

func WatchMovie(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Watching Movie.")

	//getting ID
	params := mux.Vars(req)
	id := params["id"]

	//Converting id to int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}

	//Checking for movie
	for _, v := range Data {
		if v.ID == idInt {
			json.NewEncoder(w).Encode(v)
			return
		}
	}
	json.NewEncoder(w).Encode("No movie found with the id")
	return
}

func AddMovie(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Adding Movie")

	w.Header().Set("Content-Type", "application/json")

	if req.Body == nil {
		json.NewEncoder(w).Encode("Send some data.")
	}

	var MovieData Movie
	json.NewDecoder(req.Body).Decode(&MovieData)
	if MovieData.isEmpty() {
		json.NewEncoder(w).Encode("Body is empty")
		return
	}

	// Adding random ID to the moive

	rand.Seed(time.Now().UnixNano())
	MovieData.ID = rand.Intn(10)
	Data = append(Data, MovieData)
	json.NewEncoder(w).Encode(MovieData)
	return
}

func DeleteMovie(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Deleting Movie")

	params := mux.Vars(req)
	id := params["id"]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}

	for index, dataPerMovie := range Data {
		if dataPerMovie.ID == idInt {
			Data = append(Data[:index], Data[index+1:]...)
			fmt.Println("Movie deleted")
			return
		}
	}

}


// To update a Movie first store the id of the movie and then remove the movie by using append
// Then before making the request decode the body and add id to it.
// After adding id to the movie append it to the Data. Finally again encode the body

func UpdateMovie(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Updating the Movie")

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	id := params["id"]

	// Converting id in string to id in int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}

	for index, dataMovie := range Data {
		if dataMovie.ID == idInt {
			Data = append(Data[:index], Data[index+1:]...)

			// Adding updated movie to the id
			var MovieData Movie
			_ = json.NewDecoder(req.Body).Decode(&MovieData)
			MovieData.ID = idInt
			Data = append(Data, MovieData)
			json.NewEncoder(w).Encode(MovieData)
			return
		}
	}
}
