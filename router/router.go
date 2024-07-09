package router

import (
	"github.com/Pallav46/mongoapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// Define the route to get all movies
	router.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")

	// Define the route to create a new movie
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")

	// Define the route to mark a movie as watched
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")

	// Define the route to delete a specific movie by ID
	router.HandleFunc("/api/movie/{id}", controller.DeleteMovie).Methods("DELETE")

	// Define the route to delete all movies
	router.HandleFunc("/api/movies", controller.DeleteAllMovies).Methods("DELETE")

	return router
}
