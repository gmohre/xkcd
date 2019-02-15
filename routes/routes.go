package routes

import (
	"github.com/gorilla/mux"

	"github.com/gmohre/xkcd/api"
)

func NewRoutes(api *api.API) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// api
	a := router.PathPrefix("/api").Subrouter()
	c := a.PathPrefix("/xkcd").Subrouter()

	c.HandleFunc("/latest", api.GetLatestComic).Methods("GET")
	c.HandleFunc("/comic/{comicID}", api.GetComicByID).Methods("GET")
	c.HandleFunc("/fave", api.CreateFavorite).Methods("POST")
	c.HandleFunc("/fave", api.GetFavorites).Methods("GET")
	return router
}
