package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/singh-vinayak/go-movie-server/tree/mysql/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	http.Handle("/",r)
	fmt.Println("server stated on port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080",r))
}



