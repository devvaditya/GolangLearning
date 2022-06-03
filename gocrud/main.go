package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"entgo.io/ent/entc/integration/gremlin/ent/item"
	"github.com/gogo/protobuf/test/indeximport-issue72/index"
	"github.com/gorilla/mux"
	"google.golang.org/genproto/googleapis/cloud/aiplatform/v1beta1/schema/predict/params"
)

// in this project we are not using db instead we will use struct
// struct  - an object contains key-value pairs

type Movie struct{
	ID string `json:"id"`
	// unique id 
	Isbn string `json:"isbn"`
	title string `json:"title"`
	// pointer same value in this director whatever we define in the type director
	// type of director
	Director *Director `json:"director"`


}

type Director struct{
	FirstName string `json:"firstname"`
	LastName string `json:lastname"`

}

// multiple movies in array
var movies[] movies


func getMovies(w http.ResponseWriter, r *http.Request){
	// set response in json type
	w.Header().Set("Content-Type","application/json")
	// return new encode that writes to w
	json.NewEncoder(w).Encode(movies)
}
// delete function where we will pass an id
func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	// all the movies are inside items now
	// item.id will give access to each of the movie id
	// iterating over the movies using foreach loop
	for index, item := range movies {
		if item.ID == params["id"]{
			// 
			movies = append(movies[:index],movies[index+1]...)
			break
	}
}

}

func main(){
	// adding router from the mux library
	r := mux.NewRouter()
	// at the beginnig there is no movies present so we will add some movies at the start using append
	movies = append(movies, Movie{ID:"1",Isbn:"152",title: "stranger4",Director: &Director{
		FirstName: "duffer",
		LastName: "Boys",
	}})
	movies = append(movies, Movie{ID:"2",Isbn: "321",title: "After we fell",Director: &Director{
		FirstName :"Duffer",
		LastName: "Boys",
	}})
	// 5 functions
	// 1 . getAll
	// 2 . getById
	// 3 . Create 
	// 4 . Update
	// 5 . Delete
	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/create",createMovie).Methods("POST")
	r.HandleFunc("/update/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/delete/{id}",deleteMovie).Methods("DELETE")
	// START THE SERVER
	fmt.Printf("starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080",r))
}