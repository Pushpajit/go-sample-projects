package router

import (
	"github.com/Pushpajit/go-api/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", controller.Home).Methods("GET")
	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.InsertOneMovie).Methods("POST")

	return router
}
