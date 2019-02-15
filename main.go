package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gmohre/xkcd/api"
	"github.com/gmohre/xkcd/models"
	"github.com/gmohre/xkcd/routes"

	"github.com/rs/cors"
)

func main() {
	db := models.NewSqliteDB("data.db")
	api := api.NewAPI(db)
	routes := routes.NewRoutes(api)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{getCorsAllowedOrigin()},
	})
	n := negroni.Classic()
	n.Use(c)
	n.UseHandler(routes)
	s := &http.Server{
		Addr:         ":8000",
		Handler:      n,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Serving on 8000")
	log.Fatal(s.ListenAndServe())
}

func getCorsAllowedOrigin() string {
	envContent := os.Getenv("CORS_ALLOWED_ORIGIN")
	if envContent == "" {
		envContent = "http://localhost:8080"
	}
	return envContent
}
