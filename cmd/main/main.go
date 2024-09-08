package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm/dialects/mysql"
	"github.com/singh-vinayak/go-movie-server/tree/mysql/pkg/routes"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterBookStoreRoutes(r)

	http.Handle("/", r )
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}