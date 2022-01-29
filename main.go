package main

import (
	"go-crud/controller"
	"go-crud/database"
	"go-crud/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting the HTTP Server")
	database.Open()
	database.Migrate()

	router := mux.NewRouter().StrictSlash(true)

	api := controller.PersonApi{
		Service: service.PersonService{
			Database: database.Connector,
		},
	}

	router.Use(securityFilter)

	for _, route := range api.GetAllRoutes() {
		router.HandleFunc(route.Path, route.Handle).Methods(route.Method...)
	}

	log.Fatal(http.ListenAndServe(":8090", router))
}

func securityFilter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		whiteList := []string{"/v1/authenticate"}

		for _, permitted := range whiteList {
			currUrl := r.URL.Path
			if permitted == currUrl && r.Method == "POST" {
				next.ServeHTTP(rw, r)
				return
			}
		}

		token := r.Header.Get("Authorization")

		if token == "" {
			http.Error(rw, "Forbidden", http.StatusForbidden)
		} else {
			next.ServeHTTP(rw, r)
		}

	})
}
