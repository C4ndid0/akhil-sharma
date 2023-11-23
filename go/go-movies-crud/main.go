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
  Director *Director `json:"director"`
}

type Director struct{
  FirstName string `json:"firstname"`
  LastName string `json:"lastname"`

}

var movies []Movie


func getMovies(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-type", "application/json") 
  json.NewEncoder(w).Encode(movies)
}


func getMovie(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-type", "application/json")
  params := mux.Vars(r)
  for _, item := range movies {
    if item.ID == params["id"] {
      json.NewEncoder(w).Encode(item)
      return 
    }
  }
}


func deleteMovie(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-type", "application/json")
  params := mux.Vars(r)
  for index, item := range movies {
    if item.ID == params["id"]{
      movies = append(movies[:index], movies[index+1:]... )
      break
    }
  }
  json.NewEncoder(w).Encode(movies)
}


func createMovie(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-type", "application/json")
  var movie Movie
  _ = json.NewDecoder(r.Body).Decode(&movie)
  movie.ID = strconv.Itoa(rand.Intn(10000000))
  movies = append(movies, movie)
  json.NewEncoder(w).Encode(movie)
}   


func updateMovie(w http.ResponseWriter, r *http.Request){
  w.Header().Set("content-type", "application/json")
  params := mux.Vars(r)

  for index, item := range movies {
    if item.ID == params["id"]{
      movies = append(movies[:index], movies[index+1:]...)
      var movie Movie
      _ = json.NewDecoder(r.Body).Decode(&movie)
      movie.ID = params["id"]
      movies = append(movies, movie)
      json.NewEncoder(w).Encode(movie)
    }
  }


}



func main(){
  r :=  mux.NewRouter()
  
  movies = append(movies, Movie{ID:"1",Isbn:"32112",Title:"Movie one", Director: &Director{FirstName:"John", LastName:"Doe"}}) 
  movies = append(movies, Movie{ID:"2",Isbn:"43255",Title:"Movie two", Director: &Director{FirstName:"Steve", LastName:"Duranth"}})
  movies = append(movies, Movie{ID:"3",Isbn:"02932",Title:"Movie tree", Director: &Director{FirstName:"Macus", LastName:"Castro"}})

  r.HandleFunc("/movies", getMovies).Methods("GET")
  r.HandleFunc("/movie{id}", getMovie).Methods("GET")
  r.HandleFunc("/movie", createMovie).Methods("POST")
  r.HandleFunc("/movie{id}", updateMovie).Methods("PUT")
  r.HandleFunc("/movie{id}", deleteMovie).Methods("DELETE")

  fmt.Printf("start server on port 8000 \n")
  log.Fatal(http.ListenAndServe(":8000", r))  

}
