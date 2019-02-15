package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/gmohre/xkcd/api"
)

func NewRoutes(api *api.API) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// client static files
	router.Handle("/", http.FileServer(http.Dir("./client/dist/"))).Methods("GET")
	router.PathPrefix("/static/js").Handler(http.StripPrefix("/static/js/", http.FileServer(http.Dir("./client/dist/static/js/"))))

	// api
	a := router.PathPrefix("/api").Subrouter()
	c := a.PathPrefix("/xkcd").Subrouter()

	c.HandleFunc("/latest", api.GetLatestComic).Methods("GET")
	c.HandleFunc("/comic/{comicID}", api.GetComicByID).Methods("GET")
	c.HandleFunc("/fave", api.CreateFavorite).Methods("POST")
	c.HandleFunc("/fave", api.GetFavorites).Methods("GET")
	return router
}
